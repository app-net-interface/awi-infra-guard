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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (c *Client) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {

	var routers []types.Router
	c.accountID = params.AccountId
	regionResult, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		c.logger.Errorf("Unable to describe regions, %v", err)
		return routers, err
	}
	for _, region := range regionResult.Regions {
		regionalCfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(*region.RegionName),
		)
		ec2RegionalClient := ec2.NewFromConfig(regionalCfg)
		regionalRouters, err := c.ListRoutersForRegion(ec2RegionalClient, *region.RegionName)
		if err != nil {
			//c.logger.Warnf("Error listing Transit Gateways in region %s: %v", *region.RegionName, err)
			continue
		}
		routers = append(routers, regionalRouters...)

	}
	return routers, nil
}

func (c *Client) ListRoutersForRegion(client *ec2.Client, region string) ([]types.Router, error) {

	var routers []types.Router

	paginator := ec2.NewDescribeTransitGatewaysPaginator(client, &ec2.DescribeTransitGatewaysInput{})

	for paginator.HasMorePages() {

		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, tgw := range page.TransitGateways {
			var name string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range tgw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			routers = append(routers, types.Router{
				ID:        *tgw.TransitGatewayId,
				Provider:  c.GetName(),
				Name:      name,
				Region:    region,
				State:     string(tgw.State),
				Labels:    labels,
				AccountId: *tgw.OwnerId,
				CreatedAt: *tgw.CreationTime,
			})
		}
	}
	return routers, nil
}
