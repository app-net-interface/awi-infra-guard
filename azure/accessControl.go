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

	"github.com/app-net-interface/awi-infra-guard/types"
)

type vpcPolicy string

const (
	vpcPolicyAllow = "allow"
	vpcPolicyDeny  = "deny"
)

// func (c *Client) refreshVnetSubnetsWithVPCPolicy(
// 	ctx context.Context,
// 	vnet armnetwork.VirtualNetwork,
// 	inboundVnet string,
// 	policy vpcPolicy,
// ) error {
// 	c.logger.Trace(
// 		"updating virtual network '%s' subnets with VPC Policy %s",
// 		vnet,
// 	)
// 	if vnet.Properties == nil {
// 		c.logger.Warnf(
// 			"virtual network '%s' has no properties - skipping policy update",
// 			helper.StringPointerToString(vnet.ID),
// 		)
// 		return nil
// 	}

// 	for i := range vnet.Properties.Subnets {
// 		if vnet.Properties.Subnets[i] == nil {
// 			c.logger.Warnf(
// 				"virtual network '%s' has a nil subnet pointer - skipping subnet entry",
// 				helper.StringPointerToString(vnet.ID),
// 			)
// 			continue
// 		}
// 		if vnet.Properties.Subnets[i].Properties == nil {
// 			c.logger.Warnf(
// 				"virtual network '%s' has a subnet %s with no properties - skipping subnet entry",
// 				helper.StringPointerToString(vnet.ID),
// 				helper.StringPointerToString(vnet.Properties.Subnets[i].ID),
// 			)
// 			continue
// 		}
// 	}

// 	return nil
// }

func getVnetSourceIDFromAWITags(tags map[string]string) (string, error) {
	tagValue, ok := tags["awi"]
	if !ok {
		return "", fmt.Errorf(
			"expected request key tag 'awi' with source ID but found none. Got tags: %v",
			tags,
		)
	}
	if !strings.HasPrefix(tagValue, "default-") {
		return "", fmt.Errorf(
			"the value of 'awi' tag from request has invalid prefix. Expected 'default-' but got: %s",
			tagValue,
		)
	}
	return strings.TrimPrefix(tagValue, "default-"), nil
}

// AccessControl interface implementation
func (c *Client) AddInboundAllowRuleInVPC(
	ctx context.Context,
	account string,
	region string,
	destinationVpcID string,
	cidrsToAllow []string,
	ruleName string,
	tags map[string]string,
) error {

	vnet, vnetAccount, err := c.getVPC(
		ctx, destinationVpcID, region,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to get a VPC '%s' in order to create an allow rule for it: %v",
			destinationVpcID, err,
		)
	}

	if account == "" {
		account = vnetAccount
	}

	sourceID, err := getVnetSourceIDFromAWITags(tags)
	if err != nil {
		return fmt.Errorf(
			"failed to obtain the ID of Source VPC: %w", err,
		)
	}

	err = c.refreshSubnetSecurityGroupWithVPCInbound(
		ctx,
		account,
		region,
		cidrsToAllow,
		vpcPolicyAllow,
		vnet,
		sourceID,
		ruleName,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to refresh Security Groups for subnets from VNet %s: %w",
			destinationVpcID, err,
		)
	}

	return nil
}

func (c *Client) AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, labels map[string]string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	// TBD
	return "", nil, nil
}

func (c *Client) AddInboundAllowRuleBySubnetMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, subnetCidrs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, subnets []types.Subnet, err error) {
	// TBD
	return "", nil, nil, nil
}

func (c *Client) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, instancesIPs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	// TBD
	return "", nil, nil
}

func (c *Client) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, account, region string, loadBalancerDNS string, vpcID string,
	ruleName string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (loadBalancerId, ruleId string, err error) {
	// TBD
	return "", "", nil
}

func (c *Client) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, account, region string, vpcID string, ruleName string) error {
	vnet, vnetAccount, err := c.getVPC(
		ctx, vpcID, region,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to get a VPC '%s' in order to create an allow rule for it: %v",
			vpcID, err,
		)
	}
	if account == "" {
		account = vnetAccount
	}

	err = c.deleteVPCInboundFromSubnets(
		ctx,
		account,
		region,
		vnet,
		ruleName,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to remove VPC Inbound rules from VNET '%s' subnets: %w",
			vpcID, err,
		)
	}

	return nil
}

func (c *Client) RemoveInboundAllowRulesFromVPCById(ctx context.Context, account, region string, vpcID string, instanceIDs []string,
	loadBalancersIDs []string, ruleId string) error {
	// TBD
	return nil
}

func (c *Client) RemoveInboundAllowRuleRulesByTags(ctx context.Context, account, region string, vpcID string, ruleName string, tags map[string]string) error {
	// TBD
	return nil
}

func (c *Client) RefreshInboundAllowRule(ctx context.Context, account, region string, ruleId string, cidrsToAdd []string, cidrsToRemove []string,
	destinationLabels map[string]string, destinationPrefixes []string, destinationVPCId string,
	protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error) {
	// TBD
	return nil, nil, nil
}
