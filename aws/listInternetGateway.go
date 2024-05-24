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

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {

	var igws []types.IGW
	// List all regions to ensure NAT Gateways from every region are considered
	regionResult, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		c.logger.Errorf("Unable to describe regions, %v", err)
		return igws, err
	}
	for _, region := range regionResult.Regions {
		c.logger.Debugf("Listing Internet Gateways for AWS account %s and region %s ", params.AccountId, *region.RegionName)
		regionalCfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(*region.RegionName),
		)
		if err != nil {
			c.logger.Errorf("Unable to load SDK config for region %s, %v", *region.RegionName, err)
			continue
		}

		ec2RegionalClient := ec2.NewFromConfig(regionalCfg)
		regIgws, err := c.ListInternetGatewaysInRegion(ec2RegionalClient, *region.RegionName)
		if err != nil {
			//c.logger.Warnf("Failed to list Internet Gateways in region %s: %v", *region.RegionName, err)
			continue
		}

		//for i, natGateway := range natGateways {
		//	c.logger.Infof("NAT GW [%d] %+v\n", i, natGateway)
		//}
		igws = append(igws, regIgws...)
	}
	return igws, err
}

func (c *Client) ListInternetGatewaysInRegion(client *ec2.Client, region string) ([]types.IGW, error) {
	var igws []types.IGW
	paginator := ec2.NewDescribeInternetGatewaysPaginator(client, &ec2.DescribeInternetGatewaysInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, igw := range page.InternetGateways {
			var vpcId, name string
			var state string = "unattached"
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range igw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			if len(igw.Attachments) > 0 {
				vpcId = *igw.Attachments[0].VpcId
				state = string(igw.Attachments[0].State)
			}

			igws = append(igws, types.IGW{
				ID:            *igw.InternetGatewayId,
				Provider:      c.GetName(),
				AccountID:     *igw.OwnerId,
				Name:          name,
				AttachedVpcId: vpcId,
				Region:        region,
				State:         state,
				Labels:        labels,
			})
		}
	}
	return igws, nil
}
