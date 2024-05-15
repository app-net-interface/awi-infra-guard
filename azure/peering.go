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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func (c *Client) getVnetPeeringFromVnet(vnet armnetwork.VirtualNetwork, destVNetID string) string {
	if vnet.Properties == nil {
		return ""
	}
	peeringName := vnetPeeringName(*vnet.ID, destVNetID)
	for i := range vnet.Properties.VirtualNetworkPeerings {
		if vnet.Properties.VirtualNetworkPeerings[i] == nil {
			continue
		}
		peering := vnet.Properties.VirtualNetworkPeerings[i]
		if *peering.Name == peeringName {
			return *peering.Name
		}
	}
	return ""
}

func (c *Client) deleteVnetPeering(
	ctx context.Context,
	accountID string,
	resourceGroup string,
	vnetName string,
	vnetPeeringName string,
) error {
	client, ok := c.accountClients[accountID]
	if !ok {
		return fmt.Errorf(
			"account ID '%s' is not associated with any clients", accountID,
		)
	}
	future, err := client.VNETPeering.BeginDelete(ctx, resourceGroup, vnetName, vnetPeeringName, nil)
	if err != nil {
		return fmt.Errorf("cannot delete VNet peering '%s': %w", vnetPeeringName, err)
	}
	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the VNet Peering delete future response: %w",
			err)
	}
	return nil
}

func (c *Client) createVnetPeering(
	ctx context.Context,
	sourceVnetID string,
	destinationVnetID string,
	accountID string,
) error {
	peering := armnetwork.VirtualNetworkPeering{
		Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
			AllowVirtualNetworkAccess: to.Ptr(true),
			AllowForwardedTraffic:     to.Ptr(false),
			AllowGatewayTransit:       to.Ptr(false),
			UseRemoteGateways:         to.Ptr(false),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr(destinationVnetID),
			},
		},
	}
	client, ok := c.accountClients[accountID]
	if !ok {
		return fmt.Errorf(
			"account ID '%s' is not associated with any clients", accountID,
		)
	}
	resGroup := parseResourceGroupName(sourceVnetID)
	if resGroup == "" {
		return fmt.Errorf(
			"failed to process Resource Group from Resource ID '%s'", sourceVnetID,
		)
	}
	peeringName := vnetPeeringName(sourceVnetID, destinationVnetID)

	future, err := client.VNETPeering.BeginCreateOrUpdate(
		ctx,
		resGroup,
		parseResourceName(sourceVnetID),
		peeringName,
		peering,
		nil,
	)
	if err != nil {
		return fmt.Errorf("cannot create vnet peering: %w", err)
	}

	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the vnet peering create or update future response: %w",
			err,
		)
	}

	return nil
}

func vnetPeeringName(
	sourceVnetID, destinationVnetID string,
) string {
	return fmt.Sprintf(
		"%s-%s-peering",
		parseResourceName(sourceVnetID),
		parseResourceName(destinationVnetID),
	)
}
