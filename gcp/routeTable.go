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

package gcp

import (
	"context"
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

func (c *Client) ListRouteTables(ctx context.Context, params *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error) {
	if params == nil {
		params = &infrapb.ListRouteTablesRequest{}
	}
	var net network
	var err error
	if params.GetVpcId() != "" {
		net, err = c.vpcIdToSingleNetwork(ctx, params.GetAccountId(), params.GetVpcId())
		if err != nil {
			return nil, err
		}
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	routeTables := make(map[string]*types.RouteTable)

	f := func(projectID string) error {
		iter, err := c.computeService.Routes.List(projectID).Context(ctx).Do()
		if err != nil {
			return err
		}

		for _, item := range iter.Items {
			rt, ok := routeTables[item.Network]
			if !ok {
				rt = &types.RouteTable{
					Name:      item.Name,
					ID:        strconv.FormatUint(item.Id, 10),
					Provider:  providerName,
					AccountID: projectID,
					Routes:    nil,
				}

				network := strings.Split(item.Network, "/")
				if len(network) != 0 {
					name := network[len(network)-1]
					for _, v := range networks {
						if v.Name == name || v.ID == name {
							rt.VpcID = v.ID
							break
						}
					}
				}
				if !(params.GetVpcId() == "" || net.id == rt.VpcID || net.name == rt.VpcID) {
					continue
				}
				routeTables[item.Network] = rt
			}

			route := convertRoute(item)

			rt.Routes = append(rt.Routes, route)

		}
		return nil
	}
	if params.GetAccountId() == "" {
		for projectID := range c.projectIDs {
			err := f(projectID)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}

	list := make([]types.RouteTable, 0, len(routeTables))
	for _, v := range routeTables {
		list = append(list, *v)
	}

	return list, nil
}
func convertRoute(r *compute.Route) types.Route {
	out := types.Route{
		Destination: r.DestRange,
		Status:      r.RouteStatus,
	}
	if r.NextHopGateway != "" {
		out.Target = r.NextHopGateway
	} else if r.NextHopIlb != "" {
		out.Target = r.NextHopIlb
	} else if r.NextHopIp != "" {
		out.Target = r.NextHopIp
	} else if r.NextHopInstance != "" {
		out.Target = r.NextHopInstance
	} else if r.NextHopNetwork != "" {
		out.Target = r.NextHopNetwork
	} else if r.NextHopPeering != "" {
		out.Target = r.NextHopPeering
	} else if r.NextHopVpnTunnel != "" {
		out.Target = r.NextHopVpnTunnel
	}
	return out
}
