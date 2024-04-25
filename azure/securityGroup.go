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
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
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

func (c *Client) deleteNetworkSecurityGroup(
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

func (c *Client) deleteVPCInboundFromSubnets(
	ctx context.Context,
	account string,
	region string,
	vnet armnetwork.VirtualNetwork,
	connectionTag string,
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
			continue
		}
		err := c.deleteVPCInboundFromNSG(
			ctx,
			account,
			region,
			helper.StringPointerToString(subnetProps.NetworkSecurityGroup.ID),
			connectionTag,
		)
		if err != nil {
			return fmt.Errorf(
				"failed to remove Inbound VPC Policy from NSG '%s' from subnet '%s': %w",
				helper.StringPointerToString(subnetProps.NetworkSecurityGroup.ID),
				helper.StringPointerToString(vnet.Properties.Subnets[i].ID),
				err,
			)
		}
	}
	return nil
}

func (c *Client) removeVPCPolicyRulesFromNSG(
	nsg *armnetwork.SecurityGroup,
	connectionTag string,
) error {
	if nsg == nil || nsg.Properties == nil {
		return errors.New(
			"cannot remove VPC Policies as nsg is nil or has no properties",
		)
	}

	vpcRuleName, ok := nsg.Tags[connectionTag]
	if !ok {
		return fmt.Errorf(
			"cannot remove VPC Policy Rules from NSG '%s' as there is no "+
				"tag '%s' which would tell what NSG Rules belonged to that policy",
			helper.StringPointerToString(nsg.ID), connectionTag,
		)
	}
	if vpcRuleName == nil {
		return fmt.Errorf(
			"cannot remove VPC Policy Rules from NSG '%s' as the value for "+
				"tag '%s' is nil",
			helper.StringPointerToString(nsg.ID), connectionTag,
		)
	}

	securityRulesWithoutVPCRules := make([]*armnetwork.SecurityRule, 0, len(nsg.Properties.SecurityRules))

	for i := range nsg.Properties.SecurityRules {
		if nsg.Properties.SecurityRules[i] == nil {
			continue
		}
		ruleName, err := extractAwiNSGRuleName(
			helper.StringPointerToString(nsg.Properties.SecurityRules[i].Name),
		)
		if err != nil || ruleName != *vpcRuleName {
			// Regular rules will not match our expected form so an error simply
			// indicates it is a different rule.
			//
			// We want to preserve only rules that are not matching our expected name.
			securityRulesWithoutVPCRules = append(securityRulesWithoutVPCRules, nsg.Properties.SecurityRules[i])
		}
	}

	nsg.Properties.SecurityRules = securityRulesWithoutVPCRules

	return nil
}

func (c *Client) deleteVPCInboundFromNSG(
	ctx context.Context,
	account string,
	region string,
	nsgID string,
	connectionTag string,
) error {
	nsg, account, err := c.getNSG(ctx, nsgID, region)
	if err != nil {
		return fmt.Errorf(
			"failed to get NSG for update %s: %w",
			nsgID, err,
		)
	}
	err = c.removeVPCPolicyRulesFromNSG(
		&nsg,
		connectionTag,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update VPC Policy Rules for NSG '%s': %w",
			nsgID, err,
		)
	}

	if nsg.Tags != nil {
		delete(nsg.Tags, connectionTag)
	}

	// TODO: Delete Network Security Group if it was created by AWI and
	// there are no other policies.
	return c.createNetworkSecurityGroup(
		ctx,
		helper.StringPointerToString(nsg.Name),
		region,
		account,
		parseResourceGroupName(helper.StringPointerToString(nsg.ID)),
		nsg,
	)
}

// refreshSubnetSecurityGroupWithVPCInbound checks if there is a Security
// Group created for Subnet - if there is no NSG, it will create one.
//
// After ensuring a Network Security Group exists, it will either create
// or update its rules to allow/block the inbound traffic from other VNet
// depending on the policy. If the policy is already properly configured,
// nothing will happen.
func (c *Client) refreshSubnetSecurityGroupWithVPCInbound(
	ctx context.Context,
	account string,
	region string,
	inboundCIDRs []string,
	policy vpcPolicy,
	vnet armnetwork.VirtualNetwork,
	sourceVnetID string,
	connectionTag string,
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
			err := c.createNewNetworkSecurityGroup(
				ctx,
				account,
				region,
				helper.StringPointerToString(
					vnet.Properties.Subnets[i].ID,
				),
				sourceVnetID,
				inboundCIDRs,
				policy,
				connectionTag,
			)
			if err != nil {
				return fmt.Errorf(
					"failed to create a new Network Security Group for a Subnet '%s' in VNet '%s'",
					helper.StringPointerToString(vnet.Properties.Subnets[i].ID),
					helper.StringPointerToString(vnet.ID),
				)
			}
			continue
		}
		err := c.updateNetworkSecurityGroup(
			ctx,
			region,
			helper.StringPointerToString(subnetProps.NetworkSecurityGroup.ID),
			sourceVnetID,
			inboundCIDRs,
			policy,
			connectionTag,
		)
		if err != nil {
			return fmt.Errorf(
				"failed to update Network Security Group for a Subnet '%s' in VNet '%s'",
				helper.StringPointerToString(vnet.Properties.Subnets[i].ID),
				helper.StringPointerToString(vnet.ID),
			)
		}
	}
	return nil
}

const (
	awiNSGRulesTagPrefix       string = "awi-nsg-"
	awiNSGNamePrefix           string = "awi-nsg-"
	awiNSGRulesNamePrefix      string = "awi-nsg-vpc-rules-"
	awiVPCRuleStartingPriority int32  = 2000
)

// func getAwiNSGNameTag(inboundVPC string) string {
// 	return awiNSGRulesTagPrefix + inboundVPC
// }

func getAwiNSGName(subnetName, inboundVPC string) string {
	return fmt.Sprintf(
		"%s%s-%s",
		awiNSGNamePrefix,
		parseResourceName(subnetName),
		parseResourceName(inboundVPC),
	)
}

func getAwiNSGRuleName(inboundVPC string) string {
	return awiNSGRulesNamePrefix + parseResourceName(inboundVPC)
}

func extractAwiNSGRuleName(nameWithSuffix string) (string, error) {
	if len(nameWithSuffix) < 5 {
		return "", fmt.Errorf(
			"cannot extract AwiNSGRuleName from '%s' as its too short. "+
				"It is expected to have identifier suffix such as ':0003'",
			nameWithSuffix,
		)
	}
	if nameWithSuffix[len(nameWithSuffix)-5] != ':' {
		return "", fmt.Errorf(
			"invalid AwiNSGRuleNameWithSuffix'%s'. "+
				"It is expected to have identifier suffix such as ':0003'",
			nameWithSuffix,
		)
	}
	return nameWithSuffix[:len(nameWithSuffix)-5], nil
}

func getAwiNSGRuleNameWithIDSuffix(inboundVPC string, ruleID int) (string, error) {
	if ruleID >= 10000 {
		return "", fmt.Errorf(
			"cannot generate NSG Rule with a ruleID with more than 4 digits: %d", ruleID,
		)
	}
	return awiNSGRulesNamePrefix + parseResourceName(inboundVPC) + fmt.Sprintf(":%04d", ruleID), nil
}

func takenPriorities(sg armnetwork.SecurityGroup) helper.Set[int32] {
	priorities := helper.Set[int32]{}

	if sg.Properties == nil {
		return priorities
	}

	for i := range sg.Properties.SecurityRules {
		if sg.Properties.SecurityRules[i] == nil || sg.Properties.SecurityRules[i].Properties == nil {
			continue
		}
		if sg.Properties.SecurityRules[i].Properties.Priority != nil {
			priorities.Set(*sg.Properties.SecurityRules[i].Properties.Priority)
		}
	}

	return priorities
}

func (c *Client) deletePreviousVPCPolicyRules(
	nsg *armnetwork.SecurityGroup,
	sourceVnetID string,
) {
	if nsg == nil || nsg.Properties == nil {
		return
	}

	securityRulesWithoutVPCRules := make([]*armnetwork.SecurityRule, 0, len(nsg.Properties.SecurityRules))
	vpcRuleName := getAwiNSGRuleName(sourceVnetID)

	for i := range nsg.Properties.SecurityRules {
		if nsg.Properties.SecurityRules[i] == nil {
			continue
		}
		ruleName, err := extractAwiNSGRuleName(
			helper.StringPointerToString(nsg.Properties.SecurityRules[i].Name),
		)
		if err != nil || ruleName != vpcRuleName {
			// Regular rules will not match our expected form so an error simply
			// indicates it is a different rule.
			//
			// We want to preserve only rules that are not matching our expected name.
			securityRulesWithoutVPCRules = append(securityRulesWithoutVPCRules, nsg.Properties.SecurityRules[i])
		}
	}

	nsg.Properties.SecurityRules = securityRulesWithoutVPCRules
}

func (c *Client) addVPCPolicyRulesToNSG(
	nsg *armnetwork.SecurityGroup,
	inboundCIDRs []string,
	sourceVnetID string,
	policy vpcPolicy,
	connectionTag string,
) error {
	if nsg == nil {
		return errors.New("cannot add VNet Policy rules to nil NSG")
	}

	c.deletePreviousVPCPolicyRules(nsg, sourceVnetID)

	prioritiesInUse := takenPriorities(*nsg)

	securityRules := make([]*armnetwork.SecurityRule, 0, len(inboundCIDRs))
	currentPriority := awiVPCRuleStartingPriority

	access := armnetwork.SecurityRuleAccessDeny
	if policy == vpcPolicyAllow {
		access = armnetwork.SecurityRuleAccessAllow
	}

	ruleID := 1

	for _, cidr := range inboundCIDRs {
		// Priorities cannot overlap with already existing
		// rules so find first non used priority value.
		for prioritiesInUse.Has(currentPriority) {
			currentPriority++
			if currentPriority > 4096 {
				return fmt.Errorf(
					"cannot attach VNet Policy rules to NSG '%s' as there are no more "+
						"available rule slots",
					helper.StringPointerToString(nsg.ID),
				)
			}
		}

		ruleName, err := getAwiNSGRuleNameWithIDSuffix(sourceVnetID, ruleID)
		if err != nil {
			return fmt.Errorf(
				"failed to prepare a rule name for NSG '%s': %v",
				helper.StringPointerToString(nsg.ID),
				err,
			)
		}

		securityRules = append(
			securityRules,
			&armnetwork.SecurityRule{
				Name: to.Ptr(ruleName),
				Properties: &armnetwork.SecurityRulePropertiesFormat{
					Priority:                 to.Ptr(currentPriority),
					Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
					SourceAddressPrefix:      to.Ptr(cidr),
					DestinationAddressPrefix: to.Ptr("*"),
					Access:                   to.Ptr(access),
					Direction:                to.Ptr(armnetwork.SecurityRuleDirectionInbound),
					SourcePortRange:          to.Ptr("*"),
					DestinationPortRange:     to.Ptr("*"),
				},
			},
		)
		currentPriority++
	}

	nsg.Properties.SecurityRules = securityRules

	return nil
}

func (c *Client) updateNetworkSecurityGroup(
	ctx context.Context,
	region string,
	nsgID string,
	sourceVnetID string,
	inboundCIDRs []string,
	policy vpcPolicy,
	connectionTag string,
) error {
	nsg, account, err := c.getNSG(ctx, nsgID, region)
	if err != nil {
		return fmt.Errorf(
			"failed to get NSG for update %s: %w",
			nsgID, err,
		)
	}
	err = c.addVPCPolicyRulesToNSG(
		&nsg,
		inboundCIDRs,
		sourceVnetID,
		policy,
		connectionTag,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update VPC Policy Rules for NSG '%s': %w",
			nsgID, err,
		)
	}

	if nsg.Tags == nil {
		nsg.Tags = map[string]*string{
			connectionTag: to.Ptr(getAwiNSGRuleName(sourceVnetID)),
		}
	} else {
		nsg.Tags[connectionTag] = to.Ptr(getAwiNSGRuleName(sourceVnetID))
	}

	return c.createNetworkSecurityGroup(
		ctx,
		helper.StringPointerToString(nsg.Name),
		region,
		account,
		parseResourceGroupName(helper.StringPointerToString(nsg.ID)),
		nsg,
	)
}

func (c *Client) createNewNetworkSecurityGroup(
	ctx context.Context,
	account string,
	region string,
	subnetID string,
	sourceVnetID string,
	inboundCIDRs []string,
	policy vpcPolicy,
	connectionTag string,
) error {
	ngsName := getAwiNSGName(subnetID, sourceVnetID)

	nsg := armnetwork.SecurityGroup{
		Name:     to.Ptr(ngsName),
		Location: &region,
		Tags: map[string]*string{
			connectionTag: to.Ptr(getAwiNSGRuleName(sourceVnetID)),
		},
		Properties: &armnetwork.SecurityGroupPropertiesFormat{
			Subnets: []*armnetwork.Subnet{
				{
					ID: &subnetID,
				},
			},
		},
	}

	c.addVPCPolicyRulesToNSG(
		&nsg, inboundCIDRs, sourceVnetID, policy, connectionTag,
	)

	resGroup := parseResourceGroupName(sourceVnetID)
	if resGroup == "" {
		return fmt.Errorf(
			"failed to process Resource Group from Resource ID '%s'", sourceVnetID,
		)
	}

	return c.createNetworkSecurityGroup(
		ctx,
		ngsName,
		region,
		account,
		resGroup,
		nsg,
	)
}

func (c *Client) createNetworkSecurityGroup(
	ctx context.Context,
	name string,
	location string,
	accountID string,
	resourceGroup string,
	sg armnetwork.SecurityGroup,
) error {
	client, ok := c.accountClients[accountID]
	if !ok {
		return fmt.Errorf(
			"account ID '%s' is not associated with any clients", accountID,
		)
	}

	future, err := client.NSG.BeginCreateOrUpdate(
		ctx,
		resourceGroup,
		name,
		sg,
		nil,
	)
	if err != nil {
		return fmt.Errorf("cannot create Network Security Group: %w", err)
	}

	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the nsg create or update future response: %w",
			err,
		)
	}

	return nil
}
