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
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {

	var vpces []types.VPCEndpoint
	// List all regions to ensure NAT Gateways from every region are considered
	regionResult, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		c.logger.Errorf("Unable to describe regions, %v", err)
		return vpces, err
	}
	for _, region := range regionResult.Regions {
		c.logger.Debugf("Listing VPC Endpoints for AWS account %s and region %s ", params.AccountId, *region.RegionName)
		regionalCfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(*region.RegionName),
		)
		if err != nil {
			c.logger.Errorf("Unable to load SDK config for region %s, %v", *region.RegionName, err)
			continue
		}

		ec2RegionalClient := ec2.NewFromConfig(regionalCfg)
		regVpces, err := c.ListVPCEndpointsInRegion(ec2RegionalClient, *region.RegionName)
		if err != nil {
			//c.logger.Warnf("Error listing VPC Endpoints in region %s: %v", *region.RegionName, err)
			continue
		}
		vpces = append(vpces, regVpces...)
	}
	return vpces, err
}

func (c *Client) ListVPCEndpointsInRegion(client *ec2.Client, region string) ([]types.VPCEndpoint, error) {
	var veps []types.VPCEndpoint
	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, vep := range page.VpcEndpoints {
			var name, state, serviceName, subnetIds, routeTableIds string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range vep.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}
			if vep.ServiceName != nil {
				serviceName = *vep.ServiceName
			}

			//var subnetIds, routeTableIds []string
			subnetIds = strings.Join(vep.SubnetIds, ",")
			routeTableIds = strings.Join(vep.RouteTableIds, ",")

			veps = append(veps, types.VPCEndpoint{
				ID:            *vep.VpcEndpointId,
				Provider:      c.GetName(),
				AccountId:     *vep.OwnerId,
				Name:          name,
				VPCId:         *vep.VpcId,
				Region:        region,
				State:         state,
				Labels:        labels,
				ServiceName:   serviceName,
				SubnetIds:     subnetIds,
				RouteTableIds: routeTableIds,
				CreatedAt:     vep.CreationTimestamp,
			})

		}

	}

	return veps, nil
}
