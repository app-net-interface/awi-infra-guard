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
	"strconv"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/iterator"
)

func (c *Client) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {

	var routers []types.Router

	client, err := compute.NewRoutersRESTClient(ctx)
	if err != nil {
		c.logger.Errorf("compute.NewRoutersRESTClient: %v", err)
		return routers, err
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
			return routers, err
		}

		for i, router := range resp.Value.GetRouters() {
			var advertisedRange, advertisedGroup string
			var routerName, subnetName string

			var asn uint32

			c.logger.Infof("Router [%d] = %+v %+v ", i, *router.Name, *&router.Id)
			if router.Name != nil {
				routerName = *router.Name
			}

			if router.Bgp != nil {
				for _, ag := range router.Bgp.AdvertisedGroups {
					advertisedGroup = fmt.Sprintf("%s,%s", ag, advertisedGroup)
				}
				asn = *router.Bgp.Asn
				if router.Bgp.AdvertisedIpRanges != nil {
					for _, ar := range router.Bgp.AdvertisedIpRanges {
						advertisedRange = fmt.Sprintf("%s,%s", *ar.Range, advertisedRange)
					}
				}
			}
			//Check if it's a NAT router.
			for _, nat := range router.GetNats() {
				if nat != nil {
					if len(nat.Subnetworks) > 0 {
						subnetResourceID := nat.Subnetworks[0]
						subnetName = extractResourceID(*subnetResourceID.Name)
					}
				}
			}

			router := types.Router{
				ID:              strconv.FormatUint(*router.Id, 10),
				Provider:        c.GetName(),
				Name:            routerName,
				AccountId:       params.AccountId,
				VPCId:           extractResourceID(*router.Network),
				ASN:             asn,
				AdvertisedRange: advertisedRange,
				AdvertisedGroup: advertisedGroup,
				//CIDRBlock: ,
				Region: extractResourceID(*router.Region),
				State:  "Available", // Assuming ACTIVE
				//CreatedAt:    timestamppb.New(router.GetCreationTimestamp()),
				LastSyncTime: time.Now().Format(time.RFC3339),
				SubnetId:     subnetName,
			}
			routers = append(routers, router)
		}

	}

	for i, router := range routers {
		c.logger.Infof("GCP Router [%d]  = %+v ", i, router)
	}
	return routers, err
}
