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
	"strings"
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {
	c.logger.Debugf("List VPCs")
	if params == nil {
		params = &infrapb.ListVPCRequest{}
	} else {
		c.creds = params.Creds
		c.accountID = params.AccountId
	}
	builder := newFilterBuilder()
	for k, v := range params.Labels {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if params.Region == "" || params.Region == "all" {
		var (
			wg            sync.WaitGroup
			allVPCs       []types.VPC
			allErrors     []error
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
				var vpcs []types.VPC
				var err error
				vpcs, err = c.getVPCsForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: *region.RegionName,
					vpcs:   vpcs,
					err:    err,
				}
			}(*region.RegionName)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
		}()
		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("Error in region %s: %v", result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allVPCs = append(allVPCs, result.vpcs...)
			}
		}

		if len(allErrors) > 0 {
			return allVPCs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		c.logger.Debugf("In account %s Found %d VPCs across all regions", c.accountID, len(allVPCs)-1)
		return allVPCs, nil
	}
	return c.getVPCsForRegion(ctx, params.Region, filters)
}

func (c *Client) getVPCsForRegion(ctx context.Context, region string, filters []awstypes.Filter) ([]types.VPC, error) {
	c.logger.Infof("Retreiving vpcs for region %s", region)
	client, err := c.getEC2Client(ctx, c.accountID, region)
	if err != nil {
		return nil, err
	}
	// Call DescribeVpcs operation
	resp, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	c.logger.Debugf("In account %s Found %d VPCs in region %s", c.accountID, len(resp.Vpcs)-1, region)
	return convertVPCs(resp.Vpcs, c.defaultRegion, region), nil
}

func convertVPCs(vpcs []awstypes.Vpc, defaultRegion string, region string) []types.VPC {
	if region == "" {
		region = defaultRegion
	}

	result := make([]types.VPC, 0, len(vpcs))
	for _, vpc := range vpcs {
		fmt.Printf("Found vpc %sin region %s with account id %s \n", *vpc.VpcId, region, *vpc.OwnerId)

		var ipv6CIDR string
		if len(vpc.Ipv6CidrBlockAssociationSet) > 0 {
			for _, ipv6Association := range vpc.Ipv6CidrBlockAssociationSet {
				ipv6CIDR = fmt.Sprintf("%s,%s", *ipv6Association.Ipv6CidrBlock, ipv6CIDR)
			}
		}
		vpcLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#VpcDetails:VpcId=%s", region, region, aws.ToString(vpc.VpcId))
		project := ""
		for _, tag := range vpc.Tags {
			if strings.ToLower(*tag.Key) == "project" {
				project = *tag.Value
				break
			}
		}
		result = append(result, types.VPC{
			Name:      convertString(getTagName(vpc.Tags)),
			ID:        aws.ToString(vpc.VpcId),
			Region:    region,
			Labels:    convertTags(vpc.Tags),
			IPv4CIDR:  *vpc.CidrBlock,
			IPv6CIDR:  ipv6CIDR,
			AccountID: *vpc.OwnerId,
			Provider:  providerName,
			Project:   project,
			SelfLink:  vpcLink,
		})
	}
	return result
}
