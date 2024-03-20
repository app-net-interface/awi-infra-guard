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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {

	var natGateways []types.NATGateway
	// List all regions to ensure NAT Gateways from every region are considered
	regionResult, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		c.logger.Errorf("Unable to describe regions, %v", err)
	}
	for _, region := range regionResult.Regions {
		c.logger.Debugf("Listing NAT Gateways for AWS account %s and region %s ", params.AccountId, *region.RegionName)
		regionalCfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(*region.RegionName),
		)
		if err != nil {
			c.logger.Errorf("Unable to load SDK config for region %s, %v", *region.RegionName, err)
			continue
		}

		ec2RegionalClient := ec2.NewFromConfig(regionalCfg)
		ngs, err := c.listNATGateways(ec2RegionalClient, *region.RegionName)
		if err != nil {
			c.logger.Errorf("Error listing NAT Gateways in region %s: %v", *region.RegionName, err)
			continue
		}

		//for i, natGateway := range natGateways {
		//	c.logger.Infof("NAT GW [%d] %+v\n", i, natGateway)
		//}
		natGateways = append(natGateways, ngs...)
	}
	//c.logger.Infof("NAT GW to be writtend %+v ", natGateways)
	for i, _ := range natGateways {
		c.logger.Debugf("NAT GW [%d] %+v\n", i, natGateways[i])
		natGateways[i].AccountId = params.AccountId
	}
	return natGateways, err
}

func (c *Client) listNATGateways(client *ec2.Client, region string) ([]types.NATGateway, error) {
	var natGateways []types.NATGateway
	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, ngw := range page.NatGateways {
			var name, publicIp, privateIp string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range ngw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			// Assuming the first address is the public one (if exists)
			if len(ngw.NatGatewayAddresses) > 0 {
				if ngw.NatGatewayAddresses[0].PublicIp != nil {
					publicIp = *ngw.NatGatewayAddresses[0].PublicIp
				}
				if ngw.NatGatewayAddresses[0].PrivateIp != nil {
					privateIp = *ngw.NatGatewayAddresses[0].PrivateIp
				}
			}

			natGateways = append(natGateways, types.NATGateway{
				ID:        *ngw.NatGatewayId,
				Provider:  c.GetName(),
				Name:      name,
				VpcId:     *ngw.VpcId,
				Region:    region,
				State:     string(ngw.State),
				PublicIp:  publicIp,
				PrivateIp: privateIp,
				SubnetId:  convertString(ngw.SubnetId),
				Labels:    labels,
				CreatedAt: *ngw.CreateTime,
			})
		}
	}
	return natGateways, nil
}
