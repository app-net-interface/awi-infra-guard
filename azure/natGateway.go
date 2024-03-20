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

package azure

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {

	var natGateways []types.NATGateway
	natClient, err := armnetwork.NewNatGatewaysClient(params.AccountId, c.cred, nil)
	if err != nil {
		c.logger.Errorf("Failed to create natClient: %v", err)
		return natGateways, err
	}
	publicIPClient, err := armnetwork.NewPublicIPAddressesClient(params.AccountId, c.cred, nil)
	if err != nil {
		c.logger.Errorf("Failed to create publicIPClient: %v", err)
	}

	pager := natClient.NewListAllPager(nil)

	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			c.logger.Errorf("Failed to get next page of NAT Gateways: %v", err)
			return natGateways, err
		}
		var publicIp, vnetName, vnetId, rgName, subnetID string
		var labels map[string]string = make(map[string]string)

		for _, natGateway := range result.Value {

			// Extracting Tags
			if natGateway.Tags != nil {
				for k, v := range natGateway.Tags {
					if v != nil {
						labels[k] = *v
					}
				}
			}

			// Extract VNet ID from the first associated subnet
			if natGateway.Properties.Subnets != nil && len(natGateway.Properties.Subnets) > 0 {
				subnetID = *(natGateway.Properties.Subnets)[0].ID
				//subnetId = extractLastSegment(subnetID)

				// Split the subnet ID to get the resource group name
				parts := strings.Split(subnetID, "/")
				rgName = parts[4]
				vnetName = parts[8]

				vnetId = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s", params.AccountId, rgName, vnetName)
			}
			// Assuming the first public IP association is the primary one
			if natGateway.Properties.PublicIPAddresses != nil && len(natGateway.Properties.PublicIPAddresses) > 0 {
				publicIPID := *(natGateway.Properties.PublicIPAddresses)[0].ID
				publicIpName := extractLastSegment(publicIPID)
				// Optionally, make a call to get the actual IP address from the Public IP resource
				publicIPResp, err := publicIPClient.Get(ctx, rgName, publicIpName, nil)
				if err != nil {
					c.logger.Errorf("Failed to get public IP address: %v", err)
					continue
				}

				if publicIPResp.Properties.IPAddress != nil {
					publicIp = *publicIPResp.Properties.IPAddress
				}
			}
			natGateways = append(natGateways, types.NATGateway{
				ID:        *natGateway.ID,
				Name:      *natGateway.Name,
				AccountId: params.AccountId,
				Provider:  c.GetName(),
				PublicIp:  publicIp,
				VpcId:     vnetId,
				PrivateIp: "N/A",
				SubnetId:  subnetID,
				Region:    *natGateway.Location,
				State:     string(*natGateway.Properties.ProvisioningState),
				Labels:    labels,
			},
			)
		}
	}
	return natGateways, err
}

// Helper function to extract the last segment of a resource ID (commonly used to get the name or ID of the resource)
func extractLastSegment(resourceID string) string {
	segments := strings.Split(resourceID, "/")
	return segments[len(segments)-1]
}
