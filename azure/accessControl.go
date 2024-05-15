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
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

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

	ruleset := accesscontrol.AccessControlRuleSet{}
	ruleset.NewDirectedVPCRules(
		accesscontrol.CustomRuleName(ruleName),
		accesscontrol.AccessAllow,
		cidrsToAllow,
	)

	err = c.ApplyAccessRulesToVPC(
		ctx,
		account,
		vnet,
		ruleset,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to apply rules %v to VPC %s: %w",
			ruleset, destinationVpcID, err,
		)
	}

	return nil
}

func (c *Client) getSubnetsFromInstances(
	ctx context.Context, instances []types.Instance,
) ([]armnetwork.Subnet, error) {

	type subnetInfo struct {
		VNetID   string
		SubnetID string
	}

	subnetInfos := helper.Set[subnetInfo]{}

	for _, instance := range instances {
		subnetInfos.Set(subnetInfo{
			VNetID:   instance.VPCID,
			SubnetID: instance.SubnetID,
		})
	}

	infos := subnetInfos.Keys()
	subnets := make([]armnetwork.Subnet, 0, len(infos))

	for _, info := range infos {
		subnet, _, err := c.getSubnet(
			ctx,
			parseResourceGroupName(info.SubnetID),
			parseResourceName(info.VNetID),
			parseResourceName(info.SubnetID),
		)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to get subnet %s: %w",
				info.SubnetID, err,
			)
		}
		subnets = append(subnets, subnet)
	}

	return subnets, nil
}

func (c *Client) prepareCustomAccessRules(
	instances []types.Instance,
	ruleName string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts,
) (accesscontrol.AccessControlRuleSet, error) {
	ruleset := accesscontrol.AccessControlRuleSet{}

	for _, instance := range instances {
		err := ruleset.NewCustomRules(
			accesscontrol.CustomRuleName(ruleName),
			accesscontrol.AccessAllow,
			[]string{instance.SubnetID},
			cidrsToAllow,
			[]string{instance.PrivateIP},
			protocolsAndPorts,
		)
		if err != nil {
			return accesscontrol.AccessControlRuleSet{}, fmt.Errorf(
				"failed to create custom rule: %w", err,
			)
		}
	}

	return ruleset, nil
}

func (c *Client) AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, labels map[string]string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {

	instances, err = c.ListInstances(ctx, &infrapb.ListInstancesRequest{
		VpcId:     vpcID,
		Zone:      region,
		AccountId: account,
		Labels:    labels,
		Region:    region,
	})
	if err != nil {
		return "", nil, fmt.Errorf(
			"failed to list Instances: %w", err,
		)
	}

	subnets, err := c.getSubnetsFromInstances(ctx, instances)
	if err != nil {
		return "", nil, fmt.Errorf(
			"failed to extract subnets associated with matched instances: %w", err,
		)
	}

	ruleset, err := c.prepareCustomAccessRules(
		instances,
		ruleName,
		cidrsToAllow,
		protocolsAndPorts,
	)
	if err != nil {
		return "", nil, fmt.Errorf(
			"failed to prepare custom access rules: %w", err,
		)
	}

	for _, subnet := range subnets {
		err = c.ApplyAccessRulesToSubnet(
			ctx,
			account,
			region,
			subnet,
			ruleset,
		)
		if err != nil {
			return "", nil, fmt.Errorf(
				"failed to apply access rules to subnet %s: %w",
				helper.StringPointerToString(subnet.ID),
				err,
			)
		}
	}

	return ruleName, instances, nil
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
	vnet, _, err := c.getVPC(
		ctx, vpcID, region,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to get a VPC '%s' in order to create an allow rule for it: %v",
			vpcID, err,
		)
	}

	err = c.DeleteAccessRulesFromVPC(
		ctx,
		vnet,
		accesscontrol.RuleNames{
			accesscontrol.CustomRuleName(ruleName),
		},
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
