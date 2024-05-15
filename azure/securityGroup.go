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

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListSecurityGroups(ctx context.Context, input *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error) {

	var secGroups []types.SecurityGroup
	var vNetName string
	if input.VpcId != "" {
		parts := strings.Split(input.VpcId, "/")
		if len(parts) > 0 {
			vNetName = parts[len(parts)-1]
		}
	}
	c.logger.Debugf("Retrieving security groups for account %s and VPC %s", input.AccountId, vNetName)
	vmClient, err := armcompute.NewVirtualMachinesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM client: %w", err)
	}
	nicClient, err := armnetwork.NewInterfacesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network interfaces client: %w", err)
	}

	// Example logic for listing all VMs (simplified for demonstration)
	vmPager := vmClient.NewListAllPager(nil)
	for vmPager.More() {
		vmResult, err := vmPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of VMs: %w", err)
		}
		for _, vm := range vmResult.Value {
			if vm.Properties.NetworkProfile == nil {
				continue
			}
			for _, nicRef := range vm.Properties.NetworkProfile.NetworkInterfaces {
				nic, err := nicClient.Get(ctx, parseResourceGroupName(*nicRef.ID), parseResourceName(*nicRef.ID), nil)
				if err != nil {
					fmt.Printf("Failed to get NIC: %v\n", err)
					continue
				}
				if nic.Properties == nil || nic.Properties.IPConfigurations == nil || nic.Properties.NetworkSecurityGroup == nil {
					continue
				}

				for _, ipConf := range nic.Properties.IPConfigurations {
					if ipConf.Properties == nil || ipConf.Properties.Subnet == nil || !strings.Contains(*ipConf.Properties.Subnet.ID, vNetName) {
						continue
					}
					if nic.Properties == nil || nic.Properties.NetworkSecurityGroup == nil {
						continue
					}
					// Extract VNet ID from the subnet ID
					subnetID := *ipConf.Properties.Subnet.ID
					vNetID := extractVNetIDFromSubnetID(subnetID)

					var region string
					if nic.Interface.Properties.NetworkSecurityGroup.Location != nil {
						region = *nic.Interface.Properties.NetworkSecurityGroup.Location
					}

					//c.logger.Debugf("Azure security group %+v", nic.Interface.Properties.NetworkSecurityGroup)
					secGroup := types.SecurityGroup{

						ID: *nic.Interface.Properties.NetworkSecurityGroup.ID,
						// Azure bug: NSG has a name in JSON but , not in the structure.
						Name:  parseResourceName(*nic.Interface.Properties.NetworkSecurityGroup.ID),
						VpcID: vNetID,
						//Labels:    convertToStringMap(nic.Interface.Properties.NetworkSecurityGroup.Tags),
						Region:    region,
						Provider:  c.GetName(),
						AccountID: input.AccountId,
					}

					//secGroup.Rules = convertToSecurityGroupRule(nic.Interface.Properties.NetworkSecurityGroup.Properties.SecurityRules)
					//c.logger.Debugf("Azure security group = %v", secGroup)
					secGroups = append(secGroups, secGroup)
					break // Assuming a single NIC per VM for simplicity
				}
			}
		}
	}
	return secGroups, nil
}

func (c *Client) getNSG(ctx context.Context, id, region string) (
	armnetwork.SecurityGroup, string, error,
) {
	for account, client := range c.accountClients {
		pager := client.NSG.NewListAllPager(nil)

		for pager.More() {
			resp, err := pager.NextPage(ctx)
			if err != nil {
				return armnetwork.SecurityGroup{}, "", fmt.Errorf(
					"failed to get the next page of Network Security Groups: %w", err)
			}
			for _, nsg := range resp.SecurityGroupListResult.Value {
				if nsg.Location == nil {
					continue
				}
				if nsg.ID == nil {
					continue
				}
				if *nsg.ID == id && *nsg.Location == region {
					return *nsg, account, nil
				}
			}
		}
	}
	return armnetwork.SecurityGroup{}, "", fmt.Errorf(
		"network security group '%s' not found in region '%s'", id, region,
	)
}

func (c *Client) updateNetworkSecurityGroup(
	ctx context.Context,
	account string,
	nsg armnetwork.SecurityGroup,
) error {
	_, err := c.createNetworkSecurityGroup(
		ctx,
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

func (c *Client) createNetworkSecurityGroup(
	ctx context.Context,
	name string,
	location string,
	accountID string,
	resourceGroup string,
	sg armnetwork.SecurityGroup,
) (string, error) {
	client, ok := c.accountClients[accountID]
	if !ok {
		return "", fmt.Errorf(
			"account ID '%s' is not associated with any clients", accountID,
		)
	}

	sg.Location = &location

	future, err := client.NSG.BeginCreateOrUpdate(
		ctx,
		resourceGroup,
		name,
		sg,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("cannot create Network Security Group: %w", err)
	}

	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return "", fmt.Errorf(
			"cannot get the nsg create or update future response: %w",
			err,
		)
	}

	res, err := future.Result(ctx)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create/update network security group: %w", err,
		)
	}

	return helper.StringPointerToString(res.ID), nil
}
