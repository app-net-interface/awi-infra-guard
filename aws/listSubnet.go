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
	input := &ec2.DescribeSubnetsInput{
		Filters: builder.build(),
	}
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allSubnets    []types.Subnet
			resultChannel = make(chan []types.Subnet)
			errorChannel  = make(chan error)
		)

		regionalClients, err := c.getAllClientsForProfile(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		for regionName, awsRegionClient := range regionalClients {
			wg.Add(1)
			go func(regionName string, awsRegionClient awsClient) {
				defer wg.Done()

				out, err := awsRegionClient.ec2Client.DescribeSubnets(ctx, input)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS subnets: %v", err)
					return
				}

				subnets := convertSubnets(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), regionName, out.Subnets)
				resultChannel <- subnets
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for subnets := range resultChannel {
			allSubnets = append(allSubnets, subnets...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return allSubnets, nil
	}

	client, err := c.getEC2Client(ctx, params.GetAccountId(), params.GetRegion())
	if err != nil {
		return nil, err
	}
	out, err := client.DescribeSubnets(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("could not get AWS subnets: %v", err)
	}
	return convertSubnets(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), params.GetRegion(), out.Subnets), nil
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
		result = append(result, types.Subnet{
			Zone:      convertString(subnet.AvailabilityZone),
			SubnetId:  convertString(subnet.SubnetId),
			Name:      convertString(getTagName(subnet.Tags)),
			VpcId:     convertString(subnet.VpcId),
			CidrBlock: convertString(subnet.CidrBlock),
			Labels:    convertTags(subnet.Tags),
			Region:    region,
			AccountID: account,
			Provider:  providerName,
		})
	}

	return result
}
