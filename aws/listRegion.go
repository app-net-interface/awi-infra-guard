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
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {

	var awsRegions []awsTypes.Region
	var regions []types.Region

	c.accountID = params.AccountId
	c.creds = params.Creds

	awsRegions, err := c.getAllRegions(ctx)

	for _, region := range awsRegions {
		//	c.logger.Debugf("%d.%s", count, *region.RegionName)
		regions = append(regions, types.Region{
			ID:       *region.RegionName,
			Name:     *region.RegionName,
			Provider: providerName,
		})
	}
	c.logger.Infof("Found %d regions enabled in this account", len(regions)-1)
	return regions, err
}
