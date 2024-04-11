// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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

package gcp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

func (c *Client) ListACLs(ctx context.Context, params *infrapb.ListACLsRequest) ([]types.ACL, error) {
	if params == nil {
		params = &infrapb.ListACLsRequest{}
	}
	var net network
	var err error
	if params.GetVpcId() != "" {
		net, err = c.vpcIdToSingleNetwork(ctx, params.GetAccountId(), params.GetVpcId())
		if err != nil {
			return nil, err
		}
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	acls := make([]types.ACL, 0)
	f := func(projectID string) error {
		iter, err := c.computeService.Firewalls.List(projectID).Context(ctx).Do()
		if err != nil {
			return err
		}

		for _, item := range iter.Items {
			acl := convertFirewall(projectID, networks, item)
			if !(params.GetVpcId() == "" || net.id == acl.VpcID || net.name == acl.VpcID) {
				continue
			}

			acls = append(acls, acl)
		}
		return nil
	}
	if params.GetAccountId() == "" {
		for projectID := range c.projectIDs {
			err := f(projectID)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}

	return acls, nil
}

// AddInboundAllowRuleInVPC allows specified CIDRs in VPC (network)
// supported VPC IDs format:
// 1. URL of the network resource for this firewall rule with project name information. For example:
// - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network
// - projects/myproject/global/networks/my-network -
// 2. Name or ID, for example:
// - my-network
// - 235083625034176684.
// In this case given network will be looked up in all projects specified in client.projectIDs and firewall rules will
// be applied in all found networks.
// Tagging firewall rules is not allowed in GCP.
func (c *Client) AddInboundAllowRuleInVPC(ctx context.Context, project, _ string, destinationVpcID string, cidrsToAllow []string,
	ruleName string, _ map[string]string) error {
	if destinationVpcID == "" || ruleName == "" || len(cidrsToAllow) == 0 {
		return fmt.Errorf("wrong paramters in call of function AddInboundAllowRuleInVPC")
	}

	networks, err := c.findNetwork(ctx, project, destinationVpcID)
	if err != nil {
		return err
	}

	for _, net := range networks {
		rb := &compute.Firewall{
			Allowed: []*compute.FirewallAllowed{
				{
					IPProtocol: "all",
				},
			},
			Name:         ruleName,
			Network:      net.fullUrl,
			SourceRanges: cidrsToAllow,
		}
		_, err := c.computeService.Firewalls.Insert(net.project, rb).Context(ctx).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

// AddInboundAllowRuleByLabelsMatch allows cidrsToAllow with protocolsAndPorts to all instances which match to labels
func (c *Client) AddInboundAllowRuleByLabelsMatch(ctx context.Context,
	project, _ string,
	vpcID string,
	ruleName string,
	labels map[string]string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts,
) (ruleId string, instances []types.Instance, err error) {
	net, err := c.vpcIdToSingleNetwork(ctx, project, vpcID)
	if err != nil {
		return "", nil, err
	}
	instances, err = c.addNetworkTagToInstancesByNetworkAndLabels(ctx, ruleName, net, labels)
	if err != nil {
		return "", nil, err
	}

	ruleId, err = c.addFirewallRule(ctx, net, ruleName, cidrsToAllow, protocolsAndPorts)
	if err != nil {
		return "", nil, err
	}

	return ruleId, instances, nil
}

// AddInboundAllowRuleBySubnetMatch allows cidrsToAllow with protocolsAndPorts to all instances which are within subnets which have subnetCidrs
func (c *Client) AddInboundAllowRuleBySubnetMatch(ctx context.Context,
	project, _ string,
	vpcID string,
	ruleName string,
	subnetCidrs []string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (
	ruleId string, instances []types.Instance, subnets []types.Subnet, err error) {
	net, err := c.vpcIdToSingleNetwork(ctx, project, vpcID)
	if err != nil {
		return "", nil, nil, err
	}
	instances, subnets, err = c.addNetworkTagToInstancesByNetworkAndSubnets(ctx, ruleName, net, subnetCidrs)
	if err != nil {
		return "", nil, nil, err
	}

	ruleId, err = c.addFirewallRule(ctx, net, ruleName, cidrsToAllow, protocolsAndPorts)
	if err != nil {
		return "", nil, nil, err
	}

	return ruleId, instances, subnets, nil
}

// AddInboundAllowRuleByInstanceIPMatch allows cidrsToAllow with protocolsAndPorts to all instances which have IP within instancesIPs
func (c *Client) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context,
	project, _ string,
	vpcID string, ruleName string,
	instancesIPs []string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts,
) (ruleId string, instances []types.Instance, err error) {
	net, err := c.vpcIdToSingleNetwork(ctx, project, vpcID)
	if err != nil {
		return "", nil, err
	}
	instances, err = c.addNetworkTagToInstancesByNetworkAndIps(ctx, ruleName, net, instancesIPs)
	if err != nil {
		return "", nil, err
	}

	ruleId, err = c.addFirewallRule(ctx, net, ruleName, cidrsToAllow, protocolsAndPorts)
	if err != nil {
		return "", nil, err
	}

	return ruleId, instances, nil
}

func convertFirewall(projectID string, networks []types.VPC, firewall *compute.Firewall) types.ACL {
	if firewall == nil {
		return types.ACL{}
	}

	rules := make([]types.ACLRule, 0, len(firewall.Allowed))
	for _, rule := range firewall.Allowed {
		rules = append(rules, types.ACLRule{
			Number:            int(firewall.Priority),
			Protocol:          rule.IPProtocol,
			PortRange:         strings.Join(rule.Ports, ","),
			SourceRanges:      firewall.SourceRanges,
			DestinationRanges: firewall.DestinationRanges,
			Action:            "Allow",
			Direction:         firewall.Direction,
		})
	}
	for _, rule := range firewall.Denied {
		rules = append(rules, types.ACLRule{
			Number:            int(firewall.Priority),
			Protocol:          rule.IPProtocol,
			PortRange:         strings.Join(rule.Ports, ","),
			SourceRanges:      firewall.SourceRanges,
			DestinationRanges: firewall.DestinationRanges,
			Action:            "Deny",
			Direction:         firewall.Direction,
		})
	}
	acl := types.ACL{
		ID:        strconv.FormatUint(firewall.Id, 10),
		Name:      firewall.Name,
		AccountID: projectID,
		Provider:  providerName,
		Rules:     rules,
	}
	network := strings.Split(firewall.Network, "/")
	if len(network) != 0 {
		name := network[len(network)-1]
		for _, v := range networks {
			if v.Name == name || v.ID == name {
				acl.VpcID = v.ID
				break
			}
		}
	}

	return acl
}

func (c *Client) addFirewallRule(ctx context.Context, net network, firewallRuleName string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts,
) (id string, err error) {
	rules := make([]*compute.FirewallAllowed, 0, len(protocolsAndPorts))
	for protocol, ports := range protocolsAndPorts {
		rules = append(rules, &compute.FirewallAllowed{
			IPProtocol: protocol,
			Ports:      ports,
		})
	}

	rb := &compute.Firewall{
		Allowed:      rules,
		Name:         firewallRuleName,
		Network:      net.fullUrl,
		SourceRanges: cidrsToAllow,
		TargetTags:   []string{firewallRuleName},
	}
	_, err = c.computeService.Firewalls.Insert(net.project, rb).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	firewall, err := c.computeService.Firewalls.Get(net.project, firewallRuleName).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(firewall.Id, 10), nil
}

func (c *Client) updateFirewallRule(ctx context.Context, net network, firewallRuleId string,
	cidrsToAdd, cidrsToRemove []string, protocolsAndPorts types.ProtocolsAndPorts) (ruleName string, err error) {
	firewall, err := c.computeService.Firewalls.Get(net.project, firewallRuleId).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	ruleName = firewall.Name
	cidrsToRemoveMap := make(map[string]struct{}, len(cidrsToRemove))
	for _, cidr := range cidrsToRemove {
		cidrsToRemoveMap[cidr] = struct{}{}
	}
	newCidrsToAllow := cidrsToAdd
	for _, cidr := range firewall.SourceRanges {
		if _, ok := cidrsToRemoveMap[cidr]; !ok {
			newCidrsToAllow = append(newCidrsToAllow, cidr)
		}
	}

	rules := make([]*compute.FirewallAllowed, 0, len(protocolsAndPorts))
	for protocol, ports := range protocolsAndPorts {
		rules = append(rules, &compute.FirewallAllowed{
			IPProtocol: protocol,
			Ports:      ports,
		})
	}

	rb := &compute.Firewall{
		Allowed:      rules,
		Name:         ruleName,
		Network:      net.fullUrl,
		SourceRanges: newCidrsToAllow,
		TargetTags:   []string{ruleName},
	}
	_, err = c.computeService.Firewalls.Update(net.project, ruleName, rb).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return ruleName, nil
}

func (c *Client) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context,
	project, region string,
	loadBalancerDNS string,
	vpcID string,
	ruleName string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (loadBalancerId, ruleId string, err error) {
	return "", "", fmt.Errorf("method not supported in GCP provider yet")
}

func (c *Client) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, project, _ string, vpcID string, ruleName string) error {
	net, err := c.vpcIdToSingleNetwork(ctx, project, vpcID)
	if err != nil {
		return err
	}
	_, err = c.computeService.Firewalls.Delete(net.project, ruleName).Context(ctx).Do()
	if err != nil {
		if strings.Contains(err.Error(), "notFound") {
			c.logger.Infof("firewall rule %s is already removed", ruleName)
			return nil
		}
		return err
	}
	c.logger.Infof("Succesfully removed firewall rule %s", ruleName)
	return err
}

func (c *Client) RemoveInboundAllowRulesFromVPCById(ctx context.Context,
	project, _ string,
	vpcID string,
	instanceIDs []string,
	_ []string,
	ruleId string) error {

	net, err := c.vpcIdToSingleNetwork(ctx, project, vpcID)
	if err != nil {
		return err
	}

	list, err := c.computeService.Firewalls.List(net.project).Context(ctx).Do()
	if err != nil {
		return err
	}
	var ruleName string
	for _, firewall := range list.Items {
		if strconv.FormatUint(firewall.Id, 10) == ruleId {
			_, err = c.computeService.Firewalls.Delete(net.project, ruleId).Context(ctx).Do()
			if err != nil {
				return err
			}
			c.logger.Infof("Succesfully removed firewall rule %s", firewall.Name)
			ruleName = firewall.Name
			break
		}
	}
	if ruleName == "" {
		c.logger.Infof("firewall rule with ID %s is already removed", ruleId)
		return nil
	}

	err = c.removeNetworkTagFromInstancesByIDs(ctx, net.project, ruleName, instanceIDs)
	if err != nil {
		return err
	}

	return nil
}

// RemoveInboundAllowRuleRulesByTags removes firewall rule with name ruleName, tagging firewall rules is not supported in GCP, so tags are ignored
func (c *Client) RemoveInboundAllowRuleRulesByTags(ctx context.Context, project, region string, vpcID string, ruleName string, tags map[string]string) error {
	return c.RemoveInboundAllowRuleFromVPCByName(ctx, project, region, vpcID, ruleName)
}

func (c *Client) RefreshInboundAllowRule(ctx context.Context,
	project, _ string,
	ruleId string,
	cidrsToAdd []string,
	cidrsToRemove []string,
	destinationLabels map[string]string,
	destinationPrefixes []string,
	destinationVPCId string,
	protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error) {
	net, err := c.vpcIdToSingleNetwork(ctx, project, destinationVPCId)
	if err != nil {
		return nil, nil, err
	}
	ruleName, err := c.updateFirewallRule(ctx, net, ruleId, cidrsToAdd, cidrsToRemove, protocolsAndPorts)
	if err != nil {
		return nil, nil, err
	}
	if len(destinationLabels) > 0 {
		instances, err = c.addNetworkTagToInstancesByNetworkAndLabels(ctx, ruleName, net, destinationLabels)
	} else if len(destinationPrefixes) > 0 {
		instances, subnets, err = c.addNetworkTagToInstancesByNetworkAndSubnets(ctx, ruleName, net, destinationPrefixes)
	}
	return
}
