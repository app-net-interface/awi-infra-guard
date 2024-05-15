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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListSubnets(ctx context.Context, params *infrapb.ListSubnetsRequest) ([]types.Subnet, error) {
	c.logger.Infof("Listing subnet for account %s  and VPC %s ", params.AccountId, params.VpcId)
	var subnets []types.Subnet

	// Create a virtual networks client
	vnetsClient, err := armnetwork.NewVirtualNetworksClient(params.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create virtual networks client: %w", err)
	}

	// List all VNets in the subscription
	// List all subnets in a vNet and keep adding
	vnetsPager := vnetsClient.NewListAllPager(nil)
	for vnetsPager.More() {
		vnetsResult, err := vnetsPager.NextPage(ctx)
		if err != nil {
			return subnets, fmt.Errorf("failed to get the next page of virtual networks: %w", err)
		}
		for _, vnet := range vnetsResult.Value {
			// Extract the resource group name from the VNet ID
			rgName := ExtractResourceGroupName(*vnet.ID)

			// Now, for each VNet, list its subnets
			subnetsClient, err := armnetwork.NewSubnetsClient(params.AccountId, c.cred, nil)
			if err != nil {
				return subnets, fmt.Errorf("failed to create subnets client: %w", err)
			}
			subnetsPager := subnetsClient.NewListPager(rgName, *vnet.Name, nil)
			for subnetsPager.More() {
				subnetResult, err := subnetsPager.NextPage(ctx)
				if err != nil {
					return subnets, fmt.Errorf("failed to get the next page of subnets: %w", err)
				}
				for i, subnet := range subnetResult.Value {
					cidrBlock := ""
					if subnet.Properties != nil && subnet.Properties.AddressPrefix != nil {
						cidrBlock = *subnet.Properties.AddressPrefix
					}
					// Inherit tag from vnet
					labels := make(map[string]string)
					// Assuming you want to inherit VNet's tags for its subnets
					if vnet.Tags != nil {
						for k, v := range vnet.Tags {
							labels[k] = *v
						}
					}
					subnets = append(subnets, types.Subnet{
						SubnetId:  *subnet.ID,
						Name:      *subnet.Name,
						CidrBlock: cidrBlock,
						VpcId:     *vnet.ID,
						Zone:      "unknown",
						Region:    *vnet.Location,
						Provider:  c.GetName(),
						AccountID: params.AccountId,
						Labels:    labels, // Update this as needed
					})
					c.logger.Debugf("Added subnet %+v", subnets[i])
				}
			}
		}
	}
	return subnets, nil
}

func (c *Client) getSubnet(ctx context.Context, resGroupName, vnetName, subnetName string) (
	armnetwork.Subnet, string, error,
) {
	for account, client := range c.accountClients {
		subnet, err := client.Subnet.Get(
			ctx, resGroupName, vnetName, subnetName, nil,
		)
		if err != nil {
			return armnetwork.Subnet{}, account, fmt.Errorf(
				"failed to get Subnet '%s:%s': %w", vnetName, subnetName, err)
		}
		return subnet.Subnet, account, nil
	}

	return armnetwork.Subnet{}, "", fmt.Errorf(
		"subnet '%s:%s' not found", vnetName, subnetName,
	)
}

func (c *Client) updateSubnet(
	ctx context.Context,
	account string,
	subnet armnetwork.Subnet,
) error {
	_, err := c.createSubnet(
		ctx,
		account,

		helper.StringPointerToString(nsg.Name),
		helper.StringPointerToString(nsg.Location),
		account,
		parseResourceGroupName(helper.StringPointerToString(nsg.ID)),
		nsg,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update Network Security Group %v: %w", nsg, err,
		)
	}
	return nil
}

func (c *Client) createSubnet(
	ctx context.Context,
	accountID string,
	resourceGroup string,
	name string,
	vnetName string,
	subnet armnetwork.Subnet,
) (string, error) {
	client, ok := c.accountClients[accountID]
	if !ok {
		return "", fmt.Errorf(
			"account ID '%s' is not associated with any clients", accountID,
		)
	}

	future, err := client.Subnet.BeginCreateOrUpdate(
		ctx,
		resourceGroup,
		vnetName,
		name,
		subnet,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("cannot create Subnet: %w", err)
	}

	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return "", fmt.Errorf(
			"cannot get the subnet create or update future response: %w",
			err,
		)
	}

	res, err := future.Result(ctx)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create/update subnet: %w", err,
		)
	}

	return helper.StringPointerToString(res.ID), nil
}

// Helper function to extract the resource group name from a resource ID
func ExtractResourceGroupName(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if part == "resourceGroups" && i < len(parts)-1 {
			return parts[i+1]
		}
	}
	return ""
}
