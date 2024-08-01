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
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/iterator"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {

	var natGateways []types.NATGateway
	var vpcId string
	var region string

	client, err := compute.NewRoutersRESTClient(ctx)
	if err != nil {
		c.logger.Errorf("compute.NewRoutersRESTClient: %v", err)
		return natGateways, err
	}
	defer client.Close()

	// List all routers in the project
	req := &computepb.AggregatedListRoutersRequest{
		Project: params.AccountId,
	}

	it := client.AggregatedList(ctx, req)

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.logger.Errorf("Failed to list routers: %v", err)
			return natGateways, err
		}

		for _, router := range resp.Value.GetRouters() {
			//c.logger.Infof("Router [%d] = %+v %+v ", *router.Name, *&router.Id)
			for i, nat := range router.GetNats() {
				if nat != nil {
					var routerName, subnetName string
					if router.Name != nil {
						routerName = *router.Name
					}
					if len(nat.Subnetworks) > 0 {
						subnetResourceID := nat.Subnetworks[0]
						subnetName = extractResourceID(*subnetResourceID.Name)
					}
					if router.Network != nil {
						vpcId = extractResourceID(*router.Network)
					}
					if router.Region != nil {
						region = extractResourceID(*router.Region)
					}

					natGateway := types.NATGateway{
						ID:        strconv.FormatUint(*router.Id, 10),
						Provider:  c.GetName(),
						Name:      routerName,
						AccountID: params.AccountId,
						VpcId:     vpcId,
						Region:    region,
						State:     "Available", // Assuming ACTIVE
						//CreatedAt:    timestamppb.New(router.GetCreationTimestamp()),
						LastSyncTime: time.Now().Format(time.RFC3339),
						SubnetId:     subnetName,
					}
					c.logger.Debugf("GCP NAT GW [%d]  = %+v ", i, natGateway)
					natGateways = append(natGateways, natGateway)
				}
			}

		}
	}
	return natGateways, err
}
