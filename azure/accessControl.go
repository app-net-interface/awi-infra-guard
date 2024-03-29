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

	"github.com/app-net-interface/awi-infra-guard/types"
)

// AccessControl interface implementation
func (c *Client) AddInboundAllowRuleInVPC(ctx context.Context, account, region string, destinationVpcID string, cidrsToAllow []string, ruleName string,
	tags map[string]string) error {
	// TBD
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
	// TBD
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