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

func (c *Client) ApplyAccessRulesToSubnet(
	ctx context.Context,
	account string,
	region string,
	subnet armnetwork.Subnet,
	rules accesscontrol.AccessControlRuleSet,
) error {
	if subnet.Properties == nil {
		c.logger.Warnf(
			"cannot update subnet '%s' as it has no properties",
			helper.StringPointerToString(subnet.ID),
		)
		return nil
	}

	// TODO: This situation may happen if a new Subnet
	// was created during processing this method. AWI ensures
	// first that all existing subnets are associated with any
	// NSG but this is something that can happen anyway.
	//
	// Solving this is a part of greater issue - updating subnets
	// that were created after creating a connection.
	if subnet.Properties.NetworkSecurityGroup == nil {
		c.logger.Warnf(
			"the subnet '%s' has no NSG network attached. "+
				"Cannot attach following rules there: %v",
			helper.StringPointerToString(subnet.ID), rules,
		)
		return nil
	}

	nsg, nsgAcc, err := c.getNSG(
		ctx,
		helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
		helper.StringPointerToString(&region),
	)
	if err != nil {
		return fmt.Errorf(
			"cannot get Network Security Group %s associated with subnet %s: %w",
			helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
			helper.StringPointerToString(subnet.ID),
			err,
		)
	}

	err = c.ApplyAccessRulesToNSG(
		ctx,
		nsgAcc,
		region,
		nsg,
		rules,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to apply rules %v to associated Network Security Group %s: %w",
			rules,
			helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
			err,
		)
	}

	return nil
}

func (c *Client) DeleteAccessRulesFromSubnet(
	ctx context.Context,
	region string,
	subnet armnetwork.Subnet,
	rules accesscontrol.RuleNames,
) error {
	if subnet.Properties == nil {
		c.logger.Warnf(
			"cannot update subnet '%s' as it has no properties",
			helper.StringPointerToString(subnet.ID),
		)
		return nil
	}

	// TODO: This situation may happen if a new Subnet
	// was created during processing this method. AWI ensures
	// first that all existing subnets are associated with any
	// NSG but this is something that can happen anyway.
	//
	// Solving this is a part of greater issue - updating subnets
	// that were created after creating a connection.
	if subnet.Properties.NetworkSecurityGroup == nil {
		c.logger.Warnf(
			"the subnet '%s' has no NSG network attached. "+
				"Cannot attach following rules there: %v",
			helper.StringPointerToString(subnet.ID), rules,
		)
		return nil
	}

	nsg, nsgAcc, err := c.getNSG(
		ctx,
		helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
		helper.StringPointerToString(&region),
	)
	if err != nil {
		return fmt.Errorf(
			"cannot get Network Security Group %s associated with subnet %s: %w",
			helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
			helper.StringPointerToString(subnet.ID),
			err,
		)
	}
	err = c.DeleteAccessRulesFromNSG(
		ctx,
		nsgAcc,
		region,
		nsg,
		rules,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to delete rules %v from associated Network Security Group %s: %w",
			rules,
			helper.StringPointerToString(subnet.Properties.NetworkSecurityGroup.ID),
			err,
		)
	}

	return nil
}
