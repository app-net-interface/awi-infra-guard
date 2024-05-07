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
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (c *Client) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {

	var regions []types.Region
	// List all regions to ensure NAT Gateways from every region are considered
	regionResult, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		c.logger.Errorf("Unable to describe regions, %v", err)
		return regions, err
	}
	for count, region := range regionResult.Regions {
		c.logger.Debugf("%d.%s", count, *region.RegionName)
		regions = append(regions, types.Region{
			ID:       *region.RegionName,
			Name:     *region.RegionName,
			Provider: providerName,
		})
	}
	return regions, err
}
