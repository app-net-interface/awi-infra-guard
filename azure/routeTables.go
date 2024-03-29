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
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListRouteTables(ctx context.Context, params *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error) {
	var routeTables []types.RouteTable

	routeTablesClient, err := armnetwork.NewRouteTablesClient(params.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create route tables client: %w", err)
	}

	// List all route tables in the subscription
	pager := routeTablesClient.NewListAllPager(nil)
	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of route tables: %w", err)
		}

		for _, rt := range result.Value {
			labels := make(map[string]string)
			if rt.Tags != nil {
				for k, v := range rt.Tags {
					labels[k] = *v
				}
			}

			var routes []types.Route

			if rt.Properties.Routes != nil {
				for _, route := range rt.Properties.Routes {
					if route.Properties != nil {
						var nextHopIP string
						// Assuming NextHopIPAddress is where you might find the IP, but it depends on your model.
						if route.Properties.NextHopIPAddress != nil {
							nextHopIP = *route.Properties.NextHopIPAddress
						}
						routes = append(routes, types.Route{
							Name:        *route.Name,
							Destination: *route.Properties.AddressPrefix,
							NextHopType: string(*route.Properties.NextHopType),
							NextHopIP:   nextHopIP,
						})
					}
				}
			}

			routeTable := types.RouteTable{
				Name:         *rt.Name,
				ID:           *rt.ID,
				VpcID:        "Not Attached", // This might need to be derived from associated subnets.
				Region:       *rt.Location,
				Labels:       labels,
				AccountID:    params.AccountId,
				Provider:     c.GetName(),
				Routes:       routes,
				LastSyncTime: "", // Populate this field as needed
			}

			routeTables = append(routeTables, routeTable)
			c.logger.Debugf("Added route table are %+v", routeTable)
		}
	}

	// Step 2: List all VNets and their subnets, noting any route table associations.
	va, err := ListVNetSubnetAssociations(ctx, params.AccountId, c.cred)
	if err != nil {
		return nil, err
	}

	// Step 3: Compare both lists and update the RouteTables list with VPCId and subnet.
	for i, rt := range routeTables {
		if association, ok := va.RtAssociations[rt.ID]; ok {
			routeTables[i].VpcID = association.VNetID // Update with VNet ID
			//routeTables[i].Subnets = association.SubnetIDs // Update with associated subnet IDs
		}
		// Note: Route tables without no subnet (VPC) association will simply not be updated.
	}

	return routeTables, nil
}
