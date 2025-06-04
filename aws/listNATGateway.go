// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws" // Import for aws.ToString
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListNATGateways called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListNATGateways: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListNATGateways: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListNATGateways called with VPC ID: %s, Region: %s", operationalAccountID, params.GetVpcId(), params.GetRegion())

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allNGWs       []types.NATGateway
			allErrors     []error
			resultChannel = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListNATGateways: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListNATGateways: Iterating %d regions.", operationalAccountID, len(regions))

		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListNATGateways: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getNATGatewaysForRegion
				regNgws, err := c.getNATGatewaysForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					ngws:   regNgws,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}
		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListNATGateways: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListNATGateways: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allNGWs = append(allNGWs, result.ngws...)
			}
		}
		c.logger.Infof("[AccountID: %s] ListNATGateways: Found %d NGWs across %d regions", operationalAccountID, len(allNGWs), len(regions))

		if len(allErrors) > 0 {
			return allNGWs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allNGWs, nil // Added return for success case
	}
	c.logger.Debugf("[AccountID: %s] ListNATGateways: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getNATGatewaysForRegion
	return c.getNATGatewaysForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getNATGatewaysForRegion now accepts operationalAccountID
func (c *Client) getNATGatewaysForRegion(ctx context.Context, regionName string, filters []awstypes.Filter, operationalAccountID string) ([]types.NATGateway, error) {
	c.logger.Debugf("[AccountID: %s] getNATGatewaysForRegion: Region %s", operationalAccountID, regionName)
	var natGateways []types.NATGateway
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getNATGatewaysForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{
		Filter: filters, // Note: DescribeNatGatewaysInput uses Filter, not Filters. Ensure this is correct for your SDK version.
		// If your SDK version uses Filters, it should be:
		// Filters: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx) // Use ctx from parameters
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getNATGatewaysForRegion: DescribeNatGateways NextPage failed for region %s: %v", operationalAccountID, regionName, err)
			return nil, err
		}
		for _, ngw := range page.NatGateways {
			var name, publicIp, privateIp string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range ngw.Tags {
				labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
				if aws.ToString(tag.Key) == "Name" || aws.ToString(tag.Key) == "name" {
					name = aws.ToString(tag.Value)
				}
			}

			// Assuming the first address is the public one (if exists)
			if len(ngw.NatGatewayAddresses) > 0 {
				publicIp = aws.ToString(ngw.NatGatewayAddresses[0].PublicIp)
				privateIp = aws.ToString(ngw.NatGatewayAddresses[0].PrivateIp)
				// Note: A NAT Gateway usually has an AllocationId (Elastic IP ID) instead of a direct PublicIp in NatGatewayAddresses.
				// The PublicIp is associated with that AllocationId.
				// You might need to describe the Elastic IP using ngw.NatGatewayAddresses[0].AllocationId to get the PublicIp if it's not directly here.
				// However, if the SDK populates PublicIp directly in NatGatewayAddresses, this is fine.
			}

			natGateways = append(natGateways, types.NATGateway{
				ID:        aws.ToString(ngw.NatGatewayId),
				Provider:  c.GetName(),
				Name:      name,
				VpcId:     aws.ToString(ngw.VpcId),
				Region:    regionName,
				State:     string(ngw.State),
				PublicIp:  publicIp, // This might be derived from AllocationId if not directly available
				PrivateIp: privateIp,
				SubnetId:  aws.ToString(ngw.SubnetId),
				AccountID: operationalAccountID, // Use operationalAccountID
				Labels:    labels,
				CreatedAt: aws.ToTime(ngw.CreateTime),
				SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#NatGatewayDetails:natGatewayId=%s", regionName, regionName, aws.ToString(ngw.NatGatewayId)),
			})
		}
	}
	c.logger.Debugf("[AccountID: %s] getNATGatewaysForRegion: Found %d NAT Gateways in region %s.", operationalAccountID, len(natGateways), regionName)
	return natGateways, nil
}
