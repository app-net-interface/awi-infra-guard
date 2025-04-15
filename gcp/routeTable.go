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
	"fmt"
	"net"
	"strconv"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	c "google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
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

func convertRoute(r *c.Route) types.Route {
	out := types.Route{
		Destination: r.DestRange,
		Status:      r.RouteStatus,
	}
	if r.NextHopGateway != "" {
		out.NextHopType = "Gateway"
		out.Target = r.NextHopGateway
	} else if r.NextHopIlb != "" {
		out.NextHopType = "LoadBalancer"
		out.Target = r.NextHopIlb
	} else if r.NextHopIp != "" {
		out.NextHopType = "IP"
		out.Target = r.NextHopIp
		out.NextHopIP = r.NextHopIp
	} else if r.NextHopInstance != "" {
		out.NextHopType = "Instance"
		out.Target = r.NextHopInstance
	} else if r.NextHopNetwork != "" {
		out.NextHopType = "Network"
		out.Target = r.NextHopNetwork
	} else if r.NextHopPeering != "" {
		out.NextHopType = "Peering"
		out.Target = r.NextHopPeering
	} else if r.NextHopVpnTunnel != "" {
		out.NextHopType = "VPN"
		out.Target = r.NextHopVpnTunnel
	}
	return out
}

// getSubnetRoutes returns routes in the given project that:
//  1. Belong to the same VPC network as the specified subnetwork.
//  2. Overlap in CIDR range with that subnetwork.
//  3. This overlap explains association of the route with subnet
func getSubnetRoutes(ctx context.Context, projectID, region, subnetworkName string) ([]*computepb.Route, error) {
	// 1) Create clients for subnetworks and routes.
	subnetworksClient, err := compute.NewSubnetworksRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create subnetworks client: %v", err)
	}
	defer subnetworksClient.Close()

	routesClient, err := compute.NewRoutesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create routes client: %v", err)
	}
	defer routesClient.Close()

	// 2) Retrieve the specified subnetwork.
	subnetResp, err := subnetworksClient.Get(ctx, &computepb.GetSubnetworkRequest{
		Project:    projectID,
		Region:     region,
		Subnetwork: subnetworkName,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get subnetwork %q: %v", subnetworkName, err)
	}

	subnetCIDR := subnetResp.GetIpCidrRange() // e.g. "10.10.1.0/24"
	subnetNetwork := subnetResp.GetNetwork()  // e.g. "projects/<proj>/global/networks/<network>" (or full URL)
	_, subnetNet, err := net.ParseCIDR(subnetCIDR)
	if err != nil {
		return nil, fmt.Errorf("failed to parse subnetwork CIDR (%s): %v", subnetCIDR, err)
	}

	// 3) List all routes in the project using RoutesClient.
	//    The List call returns an iterator.
	req := &computepb.ListRoutesRequest{
		Project: projectID,
	}
	it := routesClient.List(ctx, req)

	// 4) Filter routes for matching VPC network and overlapping CIDR.
	var matchingRoutes []*computepb.Route
	for {
		route, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error iterating routes: %v", err)
		}

		if route.GetNetwork() == subnetNetwork {
			destRange := route.GetDestRange()
			_, routeNet, err := net.ParseCIDR(destRange)
			if err != nil {
				// If there's any parse error, skip this route.
				continue
			}
			if netsOverlap(subnetNet, routeNet) {
				matchingRoutes = append(matchingRoutes, route)
			}
		}
	}

	return matchingRoutes, nil
}
