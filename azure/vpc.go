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

// VPCConnector interface implementation
func (c *Client) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {
	// TBD
	return types.SingleVPCConnectionOutput{}, nil
}

func (c *Client) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	// TBD
	return types.VPCConnectionOutput{}, nil
}

func (c *Client) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}

func (c *Client) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}
