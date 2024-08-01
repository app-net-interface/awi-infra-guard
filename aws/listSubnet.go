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
	subnets := convertSubnets(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), params.GetRegion(), out.Subnets)
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
	// Call DescribeVpcs operation
	resp, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertSubnets(c.defaultAccountID, c.defaultRegion, c.accountID, regionName, resp.Subnets), nil
}

func convertSubnets(defaultAccount, defaultRegion, account, region string, subnets []awstypes.Subnet) []types.Subnet {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}

	result := make([]types.Subnet, 0, len(subnets))
	for _, subnet := range subnets {

		subnetLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#SubnetDetails:subnetId=%s", region, region, aws.ToString(subnet.SubnetId))
		result = append(result, types.Subnet{
			Zone:      convertString(subnet.AvailabilityZone),
			SubnetId:  convertString(subnet.SubnetId),
			Name:      convertString(getTagName(subnet.Tags)),
			VpcId:     convertString(subnet.VpcId),
			CidrBlock: convertString(subnet.CidrBlock),
			Labels:    convertTags(subnet.Tags),
			Region:    region,
			AccountID: *subnet.OwnerId,
			Provider:  providerName,
			SelfLink:  subnetLink,
		})
	}

	return result
}
