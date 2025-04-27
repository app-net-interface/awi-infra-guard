// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package aws

import (
	"context"
	"fmt"
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) GetSubnet(ctx context.Context, params *infrapb.GetSubnetRequest) (types.Subnet, error) {
	if params.GetVpcId() == "" || params.GetId() == "" {
		return types.Subnet{}, fmt.Errorf("both vpcID and ID must be provided for GetSubnet function")
	}
	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	builder.withSubnet(params.GetId())
	input := &ec2.DescribeSubnetsInput{
		Filters: builder.build(),
	}
	client, err := c.getEC2Client(ctx, params.GetAccountId(), params.GetRegion())
	if err != nil {
		return types.Subnet{}, err
	}
	out, err := client.DescribeSubnets(ctx, input)
	if err != nil {
		return types.Subnet{}, fmt.Errorf("could not get AWS subnets: %v", err)
	}
	subnets := c.convertSubnets(ctx, client, c.defaultAccountID, c.defaultRegion, params.GetAccountId(), params.GetRegion(), out.Subnets)
	if len(subnets) == 0 {
		return types.Subnet{}, fmt.Errorf("couldn't find subnet with ID %s", params.GetId())
	}
	if len(subnets) > 1 {
		return types.Subnet{}, fmt.Errorf("more than one matching subnet, id: %s", params.GetId())
	}
	return subnets[0], nil
}

func (c *Client) ListSubnets(ctx context.Context, params *infrapb.ListSubnetsRequest) ([]types.Subnet, error) {
	c.logger.Infof("List Subnets")
	c.creds = params.Creds
	c.accountID = params.AccountId
	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	if params.GetZone() != "" {
		builder.withAvailabilityZone(params.GetZone())
	}
	if params.GetCidr() != "" {
		builder.withCIDR(params.GetCidr())
	}
	filters := builder.build()

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg         sync.WaitGroup
			allSubnets []types.Subnet
			allErrors  []error

			resultChannel = make(chan regionResult)
		)

		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}

		for _, region := range regions {
			wg.Add(1)
			go func(regionName string) {
				defer wg.Done()
				subnets, err := c.getSubnetsForRegion(ctx, *region.RegionName, filters)
				resultChannel <- regionResult{
					region:  *region.RegionName,
					subnets: subnets,
					err:     err,
				}
			}(*region.RegionName)
		}
		c.logger.Infof("In account %s Found %d subnets across %d regions", c.accountID, len(allSubnets), len(regions))
		go func() {
			wg.Wait()
			close(resultChannel)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("Error in region %s: %v", result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allSubnets = append(allSubnets, result.subnets...)
			}
		}

		if len(allErrors) > 0 {
			return allSubnets, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allSubnets, nil
	}
	return c.getSubnetsForRegion(ctx, params.Region, filters)
}

func (c *Client) getSubnetsForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.Subnet, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeSubnets operation
	resp, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	// Pass client and ctx to convertSubnets
	// Make convertSubnets a method of Client to access logger
	return c.convertSubnets(ctx, client, c.defaultAccountID, c.defaultRegion, c.accountID, regionName, resp.Subnets), nil
}

// Note: Changed to a method on *Client to access logger easily.
// Added ctx and client parameters.
func (c *Client) convertSubnets(ctx context.Context, client *ec2.Client, defaultAccount, defaultRegion, account, region string, subnets []awstypes.Subnet) []types.Subnet {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}

	result := make([]types.Subnet, 0, len(subnets))
	for _, subnet := range subnets {
		subnetId := aws.ToString(subnet.SubnetId)
		vpcId := aws.ToString(subnet.VpcId)

		// Find the associated route table ID
		routeTableId, errRT := findAssociatedRouteTableID(ctx, client, subnetId, vpcId)
		if errRT != nil {
			c.logger.Warnf("Could not determine route table for subnet %s: %v", subnetId, errRT)
			routeTableId = "" // Set to empty if lookup failed
		}

		// Find the associated network ACL ID
		networkAclId, errAcl := findAssociatedNetworkAclID(ctx, client, subnetId, vpcId)
		if errAcl != nil {
			c.logger.Warnf("Could not determine network ACL for subnet %s: %v", subnetId, errAcl)
			networkAclId = "" // Set to empty if lookup failed
		}

		subnetLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#SubnetDetails:subnetId=%s", region, region, subnetId)

		// Populate the first element of the slices as per types.Subnet definition
		rtIds := []string{}
		if routeTableId != "" {
			rtIds = append(rtIds, routeTableId)
		}
		aclIds := []string{}
		if networkAclId != "" {
			aclIds = append(aclIds, networkAclId)
		}

		result = append(result, types.Subnet{
			Zone:          convertString(subnet.AvailabilityZone),
			SubnetId:      subnetId,
			Name:          convertString(getTagName(subnet.Tags)),
			VpcId:         vpcId,
			CidrBlock:     convertString(subnet.CidrBlock),
			Labels:        convertTags(subnet.Tags),
			Region:        region,
			AccountID:     aws.ToString(subnet.OwnerId), // Use OwnerId for AccountID
			Provider:      providerName,
			SelfLink:      subnetLink,
			RouteTableIds: rtIds,  // Populate the slice
			NetworkAclIds: aclIds, // Populate the slice
			// CreatedAt/UpdatedAt might need fetching if required by types.Subnet
		})
	}

	return result
}

// findAssociatedRouteTableID finds the route table explicitly or implicitly associated with a subnet.
func findAssociatedRouteTableID(ctx context.Context, client *ec2.Client, subnetId string, vpcId string) (string, error) {
	// Describe route tables for the VPC
	rtInput := &ec2.DescribeRouteTablesInput{
		Filters: []awstypes.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcId},
			},
		},
	}
	respRTs, err := client.DescribeRouteTables(ctx, rtInput)
	if err != nil {
		return "", fmt.Errorf("failed to describe route tables for VPC %s: %w", vpcId, err)
	}

	var defaultRTId string
	for _, rt := range respRTs.RouteTables {
		//isExplicitlyAssociated := false
		// Check if this route table is explicitly associated with the target subnet
		for _, assoc := range rt.Associations {
			if aws.ToString(assoc.SubnetId) == subnetId {
				// Found explicit association
				return aws.ToString(rt.RouteTableId), nil
				// Check if this association marks the route table as main
				// Check length before accessing index
			}
			if assoc.Main != nil && *assoc.Main {
				// This route table is the main one for the VPC
				defaultRTId = aws.ToString(rt.RouteTableId)
			}
		}
		// If we found an explicit association in the inner loop, we would have returned already.
		// If we reach here, there was no explicit association for *this* route table.
		// We continue checking other route tables.
	}

	// After checking all route tables:
	// If we found an explicit association, we would have returned.
	// If we found the main route table ID, return it.
	if defaultRTId != "" {
		return defaultRTId, nil
	}

	// If we reach here, no explicit association was found, and we couldn't identify a main route table.
	// This might indicate an issue or an edge case (e.g., VPC exists but has no main RT somehow).
	return "", fmt.Errorf("could not find explicit or main route table for VPC %s / subnet %s", vpcId, subnetId)
}

// findAssociatedNetworkAclID finds the network ACL explicitly or implicitly associated with a subnet.
func findAssociatedNetworkAclID(ctx context.Context, client *ec2.Client, subnetId string, vpcId string) (string, error) {
	// Describe network ACLs for the VPC
	naclInput := &ec2.DescribeNetworkAclsInput{
		Filters: []awstypes.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcId},
			},
		},
	}
	respNacls, err := client.DescribeNetworkAcls(ctx, naclInput)
	if err != nil {
		return "", fmt.Errorf("failed to describe network ACLs for VPC %s: %w", vpcId, err)
	}

	var defaultNaclId string
	for _, nacl := range respNacls.NetworkAcls {
		// Check if this NACL is associated with the target subnet
		for _, assoc := range nacl.Associations {
			if aws.ToString(assoc.SubnetId) == subnetId {
				// Found explicit association
				return aws.ToString(nacl.NetworkAclId), nil
			}
		}
		// Keep track of the default NACL in case no explicit one is found
		if aws.ToBool(nacl.IsDefault) {
			defaultNaclId = aws.ToString(nacl.NetworkAclId)
		}
	}

	// No explicit association found, return the default NACL ID
	if defaultNaclId == "" {
		return "", fmt.Errorf("could not find default network ACL for VPC %s and no explicit association for subnet %s", vpcId, subnetId)
	}

	return defaultNaclId, nil
}
