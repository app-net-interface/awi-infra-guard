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

// ListVPC returns a slice of VPC objects for a given subscription
func (c *Client) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {

	var vpclist []types.VPC
	var ipv4Cidr, ipv6Cidr string
	accounts := c.ListAccounts()

	for _, account := range accounts {
		//subscriptionID := strings.Split(account.ID, "/")[2]
		//fmt.Printf("Subscription ID : %s\n", subscriptionID)
		vnetClient, err := armnetwork.NewVirtualNetworksClient(account.ID, c.cred, nil)
		if err != nil {
			fmt.Printf("failed to create VNet client: %w", err)
			return nil, err
		}
		pager := vnetClient.NewListAllPager(nil)

		for pager.More() {
			resp, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get the next page of VNets: %w", err)
			}
			for _, vnet := range resp.VirtualNetworkListResult.Value {
				// Convert tags to labels
				labels := make(map[string]string)
				if vnet.Tags != nil {
					for k, v := range vnet.Tags {
						labels[k] = *v
					}
				}

				for _, addressPrefix := range vnet.Properties.AddressSpace.AddressPrefixes {
					if isIPv4CIDR(*addressPrefix) {
						ipv4Cidr = fmt.Sprintf("%s,%s", *addressPrefix, ipv4Cidr)
					} else {
						ipv6Cidr = fmt.Sprintf("%s,%s", *addressPrefix, ipv6Cidr)
					}
				}

				vpc := types.VPC{
					ID:        *vnet.ID,
					Name:      *vnet.Name,
					Region:    *vnet.Location,
					Labels:    labels,
					IPv4CIDR:  ipv4Cidr,
					IPv6CIDR:  ipv6Cidr,
					Provider:  "Azure",
					AccountID: account.ID,
				}

				vpclist = append(vpclist, vpc)
				// Reset ip prefix string for the next VPC
				ipv4Cidr = ""
				ipv6Cidr = ""
			}
		}
	}

	return vpclist, nil
}

func (c *Client) getVPC(ctx context.Context, id, region string) (
	armnetwork.VirtualNetwork, string, error,
) {
	for account, client := range c.accountClients {
		pager := client.VNET.NewListAllPager(nil)

		for pager.More() {
			resp, err := pager.NextPage(ctx)
			if err != nil {
				return armnetwork.VirtualNetwork{}, "", fmt.Errorf("failed to get the next page of VNets: %w", err)
			}
			for _, vnet := range resp.VirtualNetworkListResult.Value {
				if vnet.Location == nil {
					continue
				}
				if vnet.ID == nil {
					continue
				}
				if *vnet.ID == id && *vnet.Location == region {
					return *vnet, account, nil
				}
			}
		}
	}
	return armnetwork.VirtualNetwork{}, "", fmt.Errorf(
		"vnet '%s' not found in region '%s'", id, region,
	)
}

// VPCConnector interface implementation
func (c *Client) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {

	return types.SingleVPCConnectionOutput{}, nil
}

func (c *Client) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	vnet1, accountID1, err := c.getVPC(ctx, input.Vpc1ID, input.Region1)
	if err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc1ID, input.Region1,
		)
	}
	vnet2, accountID2, err := c.getVPC(ctx, input.Vpc2ID, input.Region2)
	if err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc2ID, input.Region2,
		)
	}

	if err = c.createVnetPeering(ctx, *vnet1.ID, *vnet2.ID, accountID1); err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to create a VPC Peering from %s:%s to %s:%s due to %w",
			input.Region1, *vnet1.ID, input.Region2, *vnet2.ID, err,
		)
	}

	if err = c.createVnetPeering(ctx, *vnet2.ID, *vnet1.ID, accountID2); err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to create a VPC Peering from %s:%s to %s:%s due to %w",
			input.Region2, *vnet2.ID, input.Region1, *vnet1.ID, err,
		)
	}

	return types.VPCConnectionOutput{
		Region1: input.Region1,
		Region2: input.Region2,
	}, nil
}

func (c *Client) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}

func (c *Client) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	vnet1, accountID1, err := c.getVPC(ctx, input.Vpc1ID, input.Region1)
	if err != nil {
		return types.VPCDisconnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc1ID, input.Region1,
		)
	}
	vnet2, accountID2, err := c.getVPC(ctx, input.Vpc2ID, input.Region2)
	if err != nil {
		return types.VPCDisconnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc2ID, input.Region2,
		)
	}

	peering1 := c.getVnetPeeringFromVnet(vnet1, *vnet2.ID)
	if peering1 != "" {
		c.deleteVnetPeering(
			ctx,
			accountID1,
			parseResourceGroupName(*vnet1.ID),
			*vnet1.Name,
			vnetPeeringName(*vnet1.ID, *vnet2.ID),
		)
	} else {
		c.logger.Infof(
			"VNet Peering %s not found. Skipping it",
			vnetPeeringName(*vnet1.ID, *vnet2.ID),
		)
	}

	peering2 := c.getVnetPeeringFromVnet(vnet2, *vnet1.ID)
	if peering2 != "" {
		c.deleteVnetPeering(
			ctx,
			accountID2,
			parseResourceGroupName(*vnet2.ID),
			*vnet2.Name,
			vnetPeeringName(*vnet2.ID, *vnet1.ID),
		)
	} else {
		c.logger.Infof(
			"VNet Peering %s not found. Skipping it",
			vnetPeeringName(*vnet2.ID, *vnet1.ID),
		)
	}

	return types.VPCDisconnectionOutput{}, nil
}
