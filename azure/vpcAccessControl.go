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
	accesscontrol "github.com/app-net-interface/awi-infra-guard/azure/accessControl"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

func (c *Client) ApplyAccessRulesToVPC(
	ctx context.Context,
	account string,
	vnet armnetwork.VirtualNetwork,
	rules accesscontrol.AccessControlRuleSet,
) error {
	if vnet.Properties == nil {
		c.logger.Warnf(
			"cannot update vnet subnets as vnet '%s' has no properties",
			helper.StringPointerToString(vnet.ID),
		)
		return nil
	}

	err := c.EnsureEverySubnetInVPCHasNetworkSecurityGroup(
		ctx, account, vnet,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to set NSG for every subnet in VNET %s: %w",
			helper.StringPointerToString(vnet.ID),
			err,
		)
	}

	// Get newer vnet after updating Security Groups.
	vnet, _, err = c.getVPC(
		ctx,
		helper.StringPointerToString(vnet.ID),
		helper.StringPointerToString(vnet.Location),
	)
	if err != nil {
		return fmt.Errorf(
			"failed to get vnet: %s: %w",
			helper.StringPointerToString(vnet.ID),
			err,
		)
	}

	for i := range vnet.Properties.Subnets {
		subnet := vnet.Properties.Subnets[i]
		if subnet == nil {
			continue
		}

		err := c.ApplyAccessRulesToSubnet(
			ctx,
			account,
			helper.StringPointerToString(vnet.Location),
			*subnet,
			rules,
		)
		if err != nil {
			return fmt.Errorf(
				"failed to apply rules %v to subnet %s: %w",
				rules, helper.StringPointerToString(subnet.ID), err,
			)
		}
	}

	return nil
}

func (c *Client) DeleteAccessRulesFromVPC(
	ctx context.Context,
	vnet armnetwork.VirtualNetwork,
	rules accesscontrol.RuleNames,
) error {
	if vnet.Properties == nil {
		c.logger.Warnf(
			"cannot update vnet subnets as vnet '%s' has no properties",
			helper.StringPointerToString(vnet.ID),
		)
		return nil
	}

	for i := range vnet.Properties.Subnets {
		subnet := vnet.Properties.Subnets[i]
		if subnet == nil {
			continue
		}

		err := c.DeleteAccessRulesFromSubnet(
			ctx,
			helper.StringPointerToString(vnet.Location),
			*subnet,
			rules,
		)
		if err != nil {
			return fmt.Errorf(
				"failed to apply rules %v to subnet %s: %w",
				rules, helper.StringPointerToString(subnet.ID), err,
			)
		}
	}

	return nil
}

func (c *Client) BlockTrafficFromVPCToVPC(
	ctx context.Context,
	destinationVNET armnetwork.VirtualNetwork,
	sourceVNET armnetwork.VirtualNetwork,
	destinationVNETAccount string,
) error {
	ruleName := accesscontrol.VPCRuleName(
		helper.StringPointerToString(destinationVNET.ID),
		helper.StringPointerToString(sourceVNET.ID),
	)

	cidrsToBlock := c.getVNetCIDRs(sourceVNET)

	ruleset := accesscontrol.AccessControlRuleSet{}
	ruleset.NewDirectedVPCRules(
		ruleName,
		accesscontrol.AccessDeny,
		cidrsToBlock,
	)

	err := c.ApplyAccessRulesToVPC(
		ctx,
		destinationVNETAccount,
		destinationVNET,
		ruleset,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to apply VPC Access rules %v to VNET %s: %w",
			ruleset,
			helper.StringPointerToString(destinationVNET.ID),
			err,
		)
	}

	return nil
}

func (c *Client) UnblockTrafficFromVPCToVPC(
	ctx context.Context,
	destinationVNET armnetwork.VirtualNetwork,
	sourceVNET armnetwork.VirtualNetwork,
	destinationVNETAccount string,
) error {
	ruleName := accesscontrol.VPCRuleName(
		helper.StringPointerToString(destinationVNET.ID),
		helper.StringPointerToString(sourceVNET.ID),
	)

	err := c.DeleteAccessRulesFromVPC(
		ctx,
		destinationVNET,
		accesscontrol.RuleNames{
			ruleName,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to delete VPC Access rule %v from VNET %s: %w",
			ruleName,
			helper.StringPointerToString(destinationVNET.ID),
			err,
		)
	}

	return nil
}

func (c *Client) BlockTrafficBetweenVPCs(
	ctx context.Context,
	vnet1, vnet2 armnetwork.VirtualNetwork,
	acc1, acc2 string,
) error {
	if err := c.BlockTrafficFromVPCToVPC(ctx, vnet1, vnet2, acc1); err != nil {
		return fmt.Errorf(
			"failed to block traffic from VNET %s to VNET %s: %w",
			helper.StringPointerToString(vnet2.ID),
			helper.StringPointerToString(vnet1.ID),
			err,
		)
	}
	if err := c.BlockTrafficFromVPCToVPC(ctx, vnet2, vnet1, acc2); err != nil {
		return fmt.Errorf(
			"failed to block traffic from VNET %s to VNET %s: %w",
			helper.StringPointerToString(vnet1.ID),
			helper.StringPointerToString(vnet2.ID),
			err,
		)
	}
	return nil
}

func (c *Client) UnblockTrafficBetweenVPCs(
	ctx context.Context,
	vnet1, vnet2 armnetwork.VirtualNetwork,
	acc1, acc2 string,
) error {
	if err := c.UnblockTrafficFromVPCToVPC(ctx, vnet1, vnet2, acc1); err != nil {
		return fmt.Errorf(
			"failed to unblock traffic from VNET %s to VNET %s: %w",
			helper.StringPointerToString(vnet2.ID),
			helper.StringPointerToString(vnet1.ID),
			err,
		)
	}
	if err := c.UnblockTrafficFromVPCToVPC(ctx, vnet2, vnet1, acc2); err != nil {
		return fmt.Errorf(
			"failed to unblock traffic from VNET %s to VNET %s: %w",
			helper.StringPointerToString(vnet1.ID),
			helper.StringPointerToString(vnet2.ID),
			err,
		)
	}
	return nil
}
