// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

// ListVPCEndpoints lists all Azure Private Endpoints based on the provided parameters.
func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {
	c.logger.Infof("Listing Private Endpoints for account %s", params.AccountId)

	client, err := armnetwork.NewPrivateEndpointsClient(params.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Private Endpoints client: %w", err)
	}

	var endpoints []types.VPCEndpoint
	pager := client.NewListBySubscriptionPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get next page of Private Endpoints: %w", err)
		}

		for _, pe := range page.Value {
			if pe.Properties == nil {
				continue
			}

			// Filter by Region (Location in Azure)
			if params.Region != "" && params.Region != "all" && (pe.Location == nil || *pe.Location != params.Region) {
				continue
			}

			// Filter by VPC ID
			// In Azure, a Private Endpoint is in a subnet, which is in a VNet.
			// We need to check if the subnet's VNet ID matches the requested VpcId.
			var vpcID string
			var subnetID string
			if pe.Properties.Subnet != nil && pe.Properties.Subnet.ID != nil {
				subnetID = *pe.Properties.Subnet.ID
				// Extract VNet ID from Subnet ID
				// Format: /.../virtualNetworks/{vnet-name}/subnets/{subnet-name}
				parts := strings.Split(subnetID, "/")
				for i, part := range parts {
					if strings.EqualFold(part, "virtualNetworks") && i+1 < len(parts) {
						vpcID = strings.Join(parts[:i+2], "/")
						break
					}
				}
			}
			if params.VpcId != "" && vpcID != params.VpcId {
				continue
			}

			labels := make(map[string]string)
			if pe.Tags != nil {
				for k, v := range pe.Tags {
					if v != nil {
						labels[k] = *v
					}
				}
			}

			var name string
			if pe.Name != nil {
				name = *pe.Name
			}

			var state string
			if pe.Properties.ProvisioningState != nil {
				state = string(*pe.Properties.ProvisioningState)
			}

			var serviceName string
			if len(pe.Properties.PrivateLinkServiceConnections) > 0 && pe.Properties.PrivateLinkServiceConnections[0].Properties != nil {
				if pe.Properties.PrivateLinkServiceConnections[0].Properties.PrivateLinkServiceID != nil {
					serviceName = *pe.Properties.PrivateLinkServiceConnections[0].Properties.PrivateLinkServiceID
				}
			}

			var securityGroupIDs []string
			if pe.Properties.NetworkInterfaces != nil {
				for _, nic := range pe.Properties.NetworkInterfaces {
					if nic.Properties != nil && nic.Properties.NetworkSecurityGroup != nil && nic.Properties.NetworkSecurityGroup.ID != nil {
						securityGroupIDs = append(securityGroupIDs, *nic.Properties.NetworkSecurityGroup.ID)
					}
				}
			}

			endpoints = append(endpoints, types.VPCEndpoint{
				ID:               *pe.ID,
				Provider:         c.GetName(),
				AccountID:        params.AccountId,
				Name:             name,
				VPCId:            vpcID,
				Region:           *pe.Location,
				State:            state,
				Labels:           labels,
				Type:             "PrivateLink", // Azure's equivalent type
				ServiceName:      serviceName,
				SubnetIds:        []string{subnetID},
				SecurityGroupIDs: securityGroupIDs,
				// CreatedAt:        pe.Properties.CreationTimestamp, // This field does not exist
				SelfLink: fmt.Sprintf("https://portal.azure.com/#@/resource%s", *pe.ID),
			})
		}
	}

	c.logger.Infof("Successfully listed %d Private Endpoints for account %s", len(endpoints), params.AccountId)
	return endpoints, nil
}
