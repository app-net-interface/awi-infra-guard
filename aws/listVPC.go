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

	c.logger.Infof("Syncing VPCs")
	if params == nil {
		params = &infrapb.ListVPCRequest{}
	}
	builder := newFilterBuilder()
	for k, v := range params.Labels {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if params.Region == "" || params.Region == "all" {
		var (
			wg            sync.WaitGroup
			allvpcs       []types.VPC
			resultChannel = make(chan []types.VPC)
			errorChannel  = make(chan error)
		)

		regionalClients, err := c.getAllClientsForProfile(params.AccountId)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}
		for regionName, awsRegionClient := range regionalClients {
			wg.Add(1)
			go func(regionName string, awsRegionClient awsClient) {
				defer wg.Done()

				vpcs, err := c.getVPCsForRegion(ctx, params.AccountId, regionName, filters)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS VPCs: %v", err)
					return
				}
				resultChannel <- vpcs
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for vpcs := range resultChannel {
			allvpcs = append(allvpcs, vpcs...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return allvpcs, nil
	}

	return c.getVPCsForRegion(ctx, params.AccountId, params.Region, filters)
}

func (c *Client) getVPCsForRegion(ctx context.Context, account, region string, filters []awstypes.Filter) ([]types.VPC, error) {
	client, err := c.getEC2Client(ctx, account, region)
	// Call DescribeVpcs operation
	resp, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertVPCs(resp.Vpcs, c.defaultAccountID, c.defaultRegion, account, region), nil
}

func convertVPCs(vpcs []awstypes.Vpc, defaultAccount, defaultRegion, account, region string) []types.VPC {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	result := make([]types.VPC, 0, len(vpcs))
	for _, vpc := range vpcs {
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
			AccountID: account,
			Provider:  providerName,
			Project:   project,
			SelfLink:  vpcLink,
		})
	}
	return result
}
