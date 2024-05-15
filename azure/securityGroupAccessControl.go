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
	"crypto/sha256"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	accesscontrol "github.com/app-net-interface/awi-infra-guard/azure/accessControl"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

// ApplyAccessRulesToNSG creates access rules from the provided
// AccessControlRuleSet. This operation is preceeded with the
// removal of older entries associated with the same name to
// avoid rule duplication and to refresh possibly stale rules.
func (c *Client) ApplyAccessRulesToNSG(
	ctx context.Context,
	account string,
	region string,
	nsg armnetwork.SecurityGroup,
	rules accesscontrol.AccessControlRuleSet,
) error {
	if nsg.Properties == nil {
		return fmt.Errorf(
			"cannot add Access Rules to Network Security Group %s as it has no properties",
			helper.StringPointerToString(nsg.ID),
		)
	}

	_, err := accesscontrol.RemoveRulesFromSecurityGroup(&nsg, rules.RuleNamesForVPC())
	if err != nil {
		return fmt.Errorf(
			"failed to remove rules '%v' from Network Security Group '%s': %w",
			rules, helper.StringPointerToString(nsg.ID), err,
		)
	}

	prioritiesInUse := takenPriorities(nsg)

	securityRules, err := rules.GenerateSecurityGroupRulesForVPC(
		prioritiesInUse,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to generate Security Rules: %w", err,
		)
	}

	if nsg.Properties.SecurityRules == nil {
		nsg.Properties.SecurityRules = make([]*armnetwork.SecurityRule, 0, len(securityRules))
	}

	for i := range securityRules {
		nsg.Properties.SecurityRules = append(
			nsg.Properties.SecurityRules,
			&securityRules[i],
		)
	}

	err = c.updateNetworkSecurityGroup(
		ctx,
		account,
		nsg,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update Network Security Group '%s' with new Security Group Rules: %w",
			helper.StringPointerToString(nsg.ID), err,
		)
	}

	return nil
}

func (c *Client) DeleteAccessRulesFromNSG(
	ctx context.Context,
	account string,
	region string,
	nsg armnetwork.SecurityGroup,
	rules accesscontrol.RuleNames,
) error {
	if nsg.Properties == nil {
		return fmt.Errorf(
			"cannot add Access Rules to Network Security Group %s as it has no properties",
			helper.StringPointerToString(nsg.ID),
		)
	}

	changed, err := accesscontrol.RemoveRulesFromSecurityGroup(&nsg, rules)
	if err != nil {
		return fmt.Errorf(
			"failed to remove rules '%v' from Network Security Group '%s': %w",
			rules, helper.StringPointerToString(nsg.ID), err,
		)
	}
	if !changed {
		c.logger.Debugf(
			"Network Security Group %s had no rules associated with names "+
				"%v. Update is redundant",
			helper.StringPointerToString(nsg.ID), rules,
		)
		return nil
	}

	err = c.updateNetworkSecurityGroup(
		ctx,
		account,
		nsg,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update Network Security Group '%s' with removed Security Group Rules: %w",
			helper.StringPointerToString(nsg.ID), err,
		)
	}

	return nil
}

func awiNSGName(subnetID string) string {
	hasher := sha256.New()
	hasher.Write([]byte(subnetID))
	hashBytes := hasher.Sum(nil)

	return fmt.Sprintf("awi-nsg-%x", hashBytes)
}

func (c *Client) createAWINetworkSecurityGroup(
	ctx context.Context,
	account string,
	region string,
	subnetID string,
) (string, error) {
	ngsName := awiNSGName(subnetID)

	nsg := armnetwork.SecurityGroup{
		Name:     to.Ptr(ngsName),
		Location: &region,
		Properties: &armnetwork.SecurityGroupPropertiesFormat{

			Subnets: []*armnetwork.Subnet{
				{
					ID: &subnetID,
				},
			},
		},
	}

	resGroup := parseResourceGroupName(subnetID)
	if resGroup == "" {
		return "", fmt.Errorf(
			"failed to process Resource Group from Resource ID '%s'", subnetID,
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

func takenPriorities(sg armnetwork.SecurityGroup) helper.Set[uint] {
	priorities := helper.Set[uint]{}

	if sg.Properties == nil {
		return priorities
	}

	for i := range sg.Properties.SecurityRules {
		if sg.Properties.SecurityRules[i] == nil || sg.Properties.SecurityRules[i].Properties == nil {
			continue
		}
		if sg.Properties.SecurityRules[i].Properties.Priority != nil {
			priorities.Set(uint(*sg.Properties.SecurityRules[i].Properties.Priority))
		}
	}

	return priorities
}
