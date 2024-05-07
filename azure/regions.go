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

package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {
	var regions []types.Region
	client, err := armsubscriptions.NewClient(c.cred, nil)

	if err != nil {
		c.logger.Errorf("Failed to create vhubClient: %v", err)
		return nil, err
	}
	pager := client.NewListLocationsPager(params.AccountId, nil)

	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			c.logger.Errorf("Failed to get next page of Virtual Hubs: %v", err)
			return nil, err
		}
		for _, location := range result.Value {
			//fmt.Printf("* %s - %s\n", *location.Name, *location.DisplayName)
			regions = append(regions, types.Region{
				ID:       *location.ID,
				Name:     *location.DisplayName,
				Provider: providerName,
			})
		}
	}
	return regions, err

}
