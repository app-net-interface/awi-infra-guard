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
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
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
			fmt.Printf("failed to create VNet client: %v", err)
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

func (c *Client) getVNetCIDRs(vnet armnetwork.VirtualNetwork) []string {
	if vnet.Properties == nil {
		c.logger.Infof(
			"cannot get CIDRs from VNet as it lacks properties: %v", vnet,
		)
		return nil
	}
	if vnet.Properties.AddressSpace == nil {
		c.logger.Infof(
			"cannot get CIDRs from VNet as it lacks Address Space: %v", vnet,
		)
		return nil
	}
	prefixes := make([]string, 0, len(vnet.Properties.AddressSpace.AddressPrefixes))
	for i := range vnet.Properties.AddressSpace.AddressPrefixes {
		if vnet.Properties.AddressSpace.AddressPrefixes[i] == nil {
			c.logger.Infof(
				"The VNet %v has nil address prefix", vnet,
			)
			continue
		}
		prefixes = append(prefixes, *vnet.Properties.AddressSpace.AddressPrefixes[i])
	}
	return prefixes
}

func (c *Client) EnsureEverySubnetInVPCHasNetworkSecurityGroup(
	ctx context.Context,
	account string,
	vnet armnetwork.VirtualNetwork,
) error {
	if vnet.Properties == nil {
		c.logger.Warnf(
			"cannot update vnet subnets as vnet '%s' has no properties",
			helper.StringPointerToString(vnet.ID),
		)
		return nil
	}
	for i := range vnet.Properties.Subnets {
		if vnet.Properties.Subnets[i] == nil || vnet.Properties.Subnets[i].Properties == nil {
			continue
		}
		subnetProps := vnet.Properties.Subnets[i].Properties
		if subnetProps.NetworkSecurityGroup == nil {
			nsgID, err := c.createAWINetworkSecurityGroup(
				ctx,
				account,
				helper.StringPointerToString(vnet.Location),
				helper.StringPointerToString(vnet.Properties.Subnets[i].ID),
			)
			if err != nil {
				return fmt.Errorf(
					"failed to create AWI Network Security Group for Subnet %s: %w",
					helper.StringPointerToString(vnet.Properties.Subnets[i].ID), err,
				)
			}

			subnetProps.NetworkSecurityGroup = &armnetwork.SecurityGroup{
				ID: &nsgID,
			}

		}
	}
	return nil
}
