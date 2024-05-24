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

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {

	var routers []types.Router
	c.logger.Infof("Listing routers for %s ", c.GetName())
	vhubClient, err := armnetwork.NewVirtualHubsClient(params.AccountId, c.cred, nil)
	if err != nil {
		c.logger.Errorf("Failed to create vhubClient: %v", err)
		return nil, err
	}
	pager := vhubClient.NewListPager(nil)

	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			c.logger.Errorf("Failed to get next page of Virtual Hubs: %v", err)
			return nil, err
		}

		for _, vhub := range result.Value {
			router := buildRouter(vhub, params)
			routers = append(routers, router)
		}
	}
	return routers, err
}

func buildRouter(vhub *armnetwork.VirtualHub, params *infrapb.ListRoutersRequest) types.Router {
	var name, ps, location, ap string
	var secId []string
	var asn uint32

	var labels map[string]string = make(map[string]string)

	// Extracting Tags
	if vhub.Tags != nil {
		for k, v := range vhub.Tags {
			labels[k] = *v
		}
	}

	if vhub.Name != nil {
		name = *vhub.Name
	}

	if vhub.Properties.ProvisioningState != nil {
		ps = string(*vhub.Properties.ProvisioningState)
	}

	if vhub.Properties.AzureFirewall != nil && vhub.Properties.AzureFirewall.ID != nil {
		secId = []string{*vhub.Properties.AzureFirewall.ID}
	}

	if vhub.Properties.VirtualRouterAsn != nil {
		asn = uint32(*vhub.Properties.VirtualRouterAsn)
	}

	if vhub.Location != nil {
		location = *vhub.Location
	}

	if vhub.Properties.AddressPrefix != nil {
		ap = *vhub.Properties.AddressPrefix
	}

	router := types.Router{
		ID:               *vhub.ID,
		Name:             name,
		AccountID:        params.AccountId,
		Provider:         "Azure",
		VPCId:            "N/A",
		Region:           location,
		State:            ps,
		Labels:           labels,
		CIDRBlock:        ap,
		ASN:              asn,
		SecurityGroupIDs: secId,
	}
	return router
}
