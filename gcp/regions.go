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

package gcp

import (
	"context"
	"strconv"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

func (c *Client) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {
	var regions []types.Region

	// Create a new client
	computeService, err := compute.NewService(ctx)
	if err != nil {
		c.logger.Errorf("Failed initialize compute service", err)
		return regions, err
	}

	// List all regions in the project
	req := computeService.Regions.List(params.AccountId)
	if err := req.Pages(ctx, func(page *compute.RegionList) error {
		for _, region := range page.Items {
			//c.logger.Debugf("* %s - %s\n", region.Name, region.Description)
			regions = append(regions, types.Region{
				ID:       strconv.FormatUint(region.Id, 10),
				Name:     region.Name,
				Provider: providerName,
			})
		}
		return err
	}); err != nil {
		c.logger.Warnf("Failed to list regions: %s", err)
	}
	return regions, err
}
