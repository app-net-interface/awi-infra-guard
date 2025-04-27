// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
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

package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// formatRoute converts a Route into a readable string.
func formatRoute(route types.Route) string {
	var parts []string
	if route.Destination != "" {
		parts = append(parts, fmt.Sprintf("Dest: %s", route.Destination))
	}
	if route.Target != "" {
		parts = append(parts, fmt.Sprintf("Target: %s", route.Target))
	}
	if route.Status != "" {
		parts = append(parts, fmt.Sprintf("Status: %s", route.Status))
	}
	if route.NextHopType != "" {
		parts = append(parts, fmt.Sprintf("NextHopType: %s", route.NextHopType))
	}
	// Add other relevant fields from types.Route if needed
	return strings.Join(parts, ", ")
}

// formatSecurityGroupRule converts a SecurityGroupRule into a readable string.
func formatSecurityGroupRule(rule types.SecurityGroupRule) string {
	var parts []string
	parts = append(parts, fmt.Sprintf("Direction: %s", rule.Direction)) // Added Direction
	if rule.Protocol != "" {
		parts = append(parts, fmt.Sprintf("Protocol: %s", rule.Protocol))
	}
	if rule.PortRange != "" { // Use PortRange string directly
		parts = append(parts, fmt.Sprintf("Ports: %s", rule.PortRange))
	}
	if len(rule.Source) > 0 { // Use Source field
		parts = append(parts, fmt.Sprintf("Source: %s", strings.Join(rule.Source, ", ")))
	}
	return strings.Join(parts, ", ")
}

// formatACLRule converts an ACLRule into a readable string.
func formatACLRule(rule types.ACLRule) string {
	var parts []string
	parts = append(parts, fmt.Sprintf("Rule#: %d", rule.Number))
	parts = append(parts, fmt.Sprintf("Action: %s", rule.Action))
	parts = append(parts, fmt.Sprintf("Direction: %s", rule.Direction))
	if rule.Protocol != "" {
		parts = append(parts, fmt.Sprintf("Protocol: %s", rule.Protocol))
	}
	if rule.PortRange != "" {
		parts = append(parts, fmt.Sprintf("Ports: %s", rule.PortRange))
	}
	if len(rule.SourceRanges) > 0 {
		parts = append(parts, fmt.Sprintf("SourceCIDR: %s", strings.Join(rule.SourceRanges, ", ")))
	}
	if len(rule.DestinationRanges) > 0 {
		parts = append(parts, fmt.Sprintf("DestCIDR: %s", strings.Join(rule.DestinationRanges, ", ")))
	}
	return strings.Join(parts, ", ")
}

func (p *providerWithDB) ListVpcGraphNodes(ctx context.Context, params *infrapb.ListVpcGraphNodesRequest) ([]types.VpcGraphNode, error) {
	vpcIndexKey := types.CloudID(p.realProvider.GetName(), params.GetVpcId())
	vpcIndex, err := p.dbClient.GetVPCIndex(vpcIndexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC index for %s: %w", vpcIndexKey, err)
	}
	if vpcIndex == nil {
		return nil, fmt.Errorf("VPC index not found for %s", vpcIndexKey)
	}
	logger := p.logger
	logger.Infof("VPC Index for %s contains RouteTableIds: %v", vpcIndexKey, vpcIndex.RouteTableIds)

	var nodes []types.VpcGraphNode
	providerName := p.realProvider.GetName()

	addProp := func(props map[string]string, key, value string) {
		if value != "" {
			props[key] = value
		}
	}
	addPropSlice := func(props map[string]string, key string, values []string) {
		if len(values) > 0 {
			props[key] = strings.Join(values, ",")
		}
	}

	createNode := func(id, resourceType, name string, properties map[string]string) types.VpcGraphNode {
		if properties == nil {
			properties = make(map[string]string)
		}
		addProp(properties, "name", name)
		addProp(properties, "id", id)
		addProp(properties, "resourceType", resourceType)

		return types.VpcGraphNode{
			ID:           id,
			ResourceType: resourceType,
			Name:         name,
			Properties:   properties,
			Provider:     providerName,
			AccountId:    vpcIndex.AccountId,
			Region:       vpcIndex.Region,
		}
	}

	for _, id := range vpcIndex.InstanceIds {
		res, err := p.dbClient.GetInstance(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "privateIP", res.PrivateIP)
			addProp(props, "publicIP", res.PublicIP)
			addProp(props, "state", res.State)
			addProp(props, "type", res.Type)
			addProp(props, "zone", res.Zone)
			addProp(props, "subnetID", res.SubnetID)
			addProp(props, "vpcID", res.VPCID)
			nodes = append(nodes, createNode(res.ID, types.InstanceType, res.Name, props))
		} else {
			logger.Warnf("Failed to get instance %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.SubnetIds {
		res, err := p.dbClient.GetSubnet(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "cidrBlock", res.CidrBlock)
			addProp(props, "zone", res.Zone)
			addProp(props, "vpcID", res.VpcId)
			addPropSlice(props, "routeTableIDs", res.RouteTableIds)
			addPropSlice(props, "networkAclIDs", res.NetworkAclIds)
			nodes = append(nodes, createNode(res.SubnetId, types.SubnetType, res.Name, props))
		} else {
			logger.Warnf("Failed to get subnet %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.RouteTableIds {
		res, err := p.dbClient.GetRouteTable(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "self_link", res.SelfLink)
			addProp(props, "vpcID", res.VpcID)
			for i, route := range res.Routes {
				key := fmt.Sprintf("route_%d", i+1)
				props[key] = formatRoute(route)
			}
			nodes = append(nodes, createNode(res.ID, types.RouteTableType, res.Name, props))
		} else {
			logger.Warnf("Failed to get route table %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.NatGatewayIds {
		res, err := p.dbClient.GetNATGateway(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "publicIP", res.PublicIp)
			addProp(props, "privateIP", res.PrivateIp)
			addProp(props, "state", res.State)
			addProp(props, "subnetID", res.SubnetId)
			addProp(props, "vpcID", res.VpcId)
			nodes = append(nodes, createNode(res.ID, types.NATGatewayType, res.Name, props))
		} else {
			logger.Warnf("Failed to get NAT gateway %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.IgwIds {
		res, err := p.dbClient.GetIGW(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "state", res.State)
			addProp(props, "attachedVpcId", res.AttachedVpcId)
			nodes = append(nodes, createNode(res.ID, types.IGWType, res.Name, props))
		} else {
			logger.Warnf("Failed to get internet gateway %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.SecurityGroupIds {
		res, err := p.dbClient.GetSecurityGroup(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "self_link", res.SelfLink)
			addProp(props, "name", res.Name)
			addProp(props, "vpcID", res.VpcID)

			for i, rule := range res.Rules {
				key := fmt.Sprintf("rule_%d", i+1)
				props[key] = formatSecurityGroupRule(rule)
			}

			nodes = append(nodes, createNode(res.ID, types.SecurityGroupType, res.Name, props))
		} else {
			logger.Warnf("Failed to get security group %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.AclIds {
		res, err := p.dbClient.GetACL(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "name", res.Name)
			addProp(props, "self_link", res.SelfLink)
			addProp(props, "vpcID", res.VpcID)

			for i, rule := range res.Rules {
				key := fmt.Sprintf("rule_%d", i+1)
				props[key] = formatACLRule(rule)
			}

			nodes = append(nodes, createNode(res.ID, types.ACLType, res.Name, props))
		} else {
			logger.Warnf("Failed to get ACL %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.LbIds {
		res, err := p.dbClient.GetLB(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "dnsName", res.DNSName)
			addProp(props, "type", res.Type)
			addProp(props, "scheme", res.Scheme)
			addProp(props, "vpcID", res.VPCID)
			addProp(props, "ipAddressType", res.IPAddressType)
			addPropSlice(props, "ipAddresses", res.IPAddresses)
			addPropSlice(props, "subnetIDs", res.SubnetIDs)
			addPropSlice(props, "securityGroupIDs", res.SecurityGroupIDs)
			nodes = append(nodes, createNode(res.ID, types.LBType, res.Name, props))
		} else {
			logger.Warnf("Failed to get LB %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.VpcEndpointIds {
		res, err := p.dbClient.GetVPCEndpoint(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "serviceName", res.ServiceName)
			addProp(props, "type", res.Type)
			addProp(props, "state", res.State)
			addProp(props, "vpcID", res.VPCId)
			addPropSlice(props, "subnetIDs", res.SubnetIds)
			addPropSlice(props, "routeTableIDs", res.RouteTableIds)
			addPropSlice(props, "securityGroupIDs", res.SecurityGroupIDs)
			nodes = append(nodes, createNode(res.ID, types.VPCEndpointType, res.Name, props))
		} else {
			logger.Warnf("Failed to get VPC Endpoint %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.RouterIds {
		res, err := p.dbClient.GetRouter(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "asn", fmt.Sprintf("%d", res.ASN))
			addProp(props, "state", res.State)
			addProp(props, "vpcID", res.VPCId)
			addProp(props, "advertisedRange", res.AdvertisedRange)
			addProp(props, "advertisedGroup", res.AdvertisedGroup)
			nodes = append(nodes, createNode(res.ID, types.RouterType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Router %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.NetworkInterfaceIds {
		res, err := p.dbClient.GetNetworkInterface(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "macAddress", res.MacAddress)
			addProp(props, "status", res.Status)
			addProp(props, "publicIP", res.PublicIP)
			addProp(props, "privateIP", res.PrivateIPs[0])
			addProp(props, "subnetID", res.SubnetID)
			addProp(props, "instanceID", res.InstanceID)
			addProp(props, "description", res.Description)
			addProp(props, "type", res.InterfaceType)
			nodes = append(nodes, createNode(res.ID, types.NetworkInterfaceType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Network Interface %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.VpnConcentratorIds {
		res, err := p.dbClient.GetVPNConcentrator(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "asn", fmt.Sprintf("%d", res.ASN))
			addProp(props, "state", res.State)
			addProp(props, "type", res.Type)
			addProp(props, "self_link", res.SelfLink)

			nodes = append(nodes, createNode(res.ID, types.VPNConcentratorType, res.Name, props))
		} else {
			logger.Warnf("Failed to get VPN Concentrator %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.PublicIpIds {
		res, err := p.dbClient.GetPublicIP(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "publicIP", res.PublicIP)
			addProp(props, "instanceId", res.InstanceId)
			addProp(props, "type", res.Type)
			addProp(props, "networkInterfaceId", res.NetworkInterfaceId)
			addProp(props, "privateIP", res.PrivateIP)
			addProp(props, "self_link", res.SelfLink)
			nodes = append(nodes, createNode(res.ID, types.PublicIPType, res.ID, props))
		} else {
			logger.Warnf("Failed to get Public IP %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.ClusterIds {
		res, err := p.dbClient.GetCluster(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := make(map[string]string)
			addProp(props, "name", res.FullName)

			nodes = append(nodes, createNode(res.Id, types.ClusterType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Cluster %s for graph node: %v", id, err)
		}
	}

	return nodes, nil
}

func (p *providerWithDB) ListVpcGraphEdges(ctx context.Context, params *infrapb.ListVpcGraphEdgesRequest) ([]types.VpcGraphEdge, error) {
	vpcIndexKey := types.CloudID(p.realProvider.GetName(), params.GetVpcId())
	vpcIndex, err := p.dbClient.GetVPCIndex(vpcIndexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC index for %s: %w", vpcIndexKey, err)
	}
	if vpcIndex == nil {
		return nil, fmt.Errorf("VPC index not found for %s", vpcIndexKey)
	}

	var edges []types.VpcGraphEdge
	providerName := p.realProvider.GetName()
	logger := p.logger

	createEdge := func(sourceID, targetID, relationshipType string) types.VpcGraphEdge {
		if sourceID == "" || targetID == "" {
			logger.Warnf("Attempted to create edge with empty source/target: %s -> %s (%s)", sourceID, targetID, relationshipType)
			return types.VpcGraphEdge{}
		}
		if sourceID == vpcIndex.VpcId || targetID == vpcIndex.VpcId {
			return types.VpcGraphEdge{}
		}
		return types.VpcGraphEdge{
			SourceNodeID:     sourceID,
			TargetNodeID:     targetID,
			RelationshipType: relationshipType,
			Provider:         providerName,
			AccountId:        vpcIndex.AccountId,
			Region:           vpcIndex.Region,
		}
	}

	// 1. Subnet relationships (RouteTable, ACL)
	for _, subnetId := range vpcIndex.SubnetIds {
		subnet, err := p.dbClient.GetSubnet(types.CloudID(providerName, subnetId))
		if err != nil || subnet == nil {
			logger.Warnf("Failed to get subnet %s for edge generation: %v", subnetId, err)
			continue
		}

		// Subnet -> RouteTable edges (Iterate through all associated Route Tables)
		for _, rtId := range subnet.RouteTableIds {
			if rtId == "" { // Skip empty IDs
				continue
			}
			if contains(vpcIndex.RouteTableIds, rtId) {
				edge := createEdge(subnetId, rtId, "USES_ROUTE_TABLE")
				if edge.SourceNodeID != "" { // Check if edge creation was successful (non-empty IDs)
					edges = append(edges, edge)
				}
			} else {
				// Log if an associated RT is not found in the index (might indicate partial data or explicit association outside VPC)
				logger.Warnf("Subnet %s associated route table %s not found in VPC index", subnetId, rtId)
			}
		}

		// Subnet -> ACL edges (Iterate through all associated ACLs)
		for _, aclId := range subnet.NetworkAclIds {
			if aclId == "" { // Skip empty IDs
				continue
			}
			if contains(vpcIndex.AclIds, aclId) {
				edge := createEdge(subnetId, aclId, "USES_ACL")
				if edge.SourceNodeID != "" { // Check if edge creation was successful (non-empty IDs)
					edges = append(edges, edge)
				}
			} else {
				// Log if an associated ACL is not found in the index
				logger.Warnf("Subnet %s associated ACL %s not found in VPC index", subnetId, aclId)
			}
		}
	}

	// 2. RouteTable relationships (Routes)
	for _, rtId := range vpcIndex.RouteTableIds {
		rt, err := p.dbClient.GetRouteTable(types.CloudID(providerName, rtId))
		if err != nil || rt == nil {
			logger.Warnf("Failed to get route table %s for edge generation: %v", rtId, err)
			continue
		}

		// RouteTable -> Target (IGW, NAT, Instance, VPCE, NI, VGW)
		for _, route := range rt.Routes {
			targetId := route.Target
			relationship := "ROUTES_TO"
			if contains(vpcIndex.IgwIds, targetId) ||
				contains(vpcIndex.NatGatewayIds, targetId) ||
				contains(vpcIndex.InstanceIds, targetId) ||
				contains(vpcIndex.VpcEndpointIds, targetId) ||
				contains(vpcIndex.NetworkInterfaceIds, targetId) ||
				contains(vpcIndex.VpnConcentratorIds, targetId) ||
				contains(vpcIndex.RouterIds, targetId) {
				edge := createEdge(rtId, targetId, relationship)
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			} else if targetId != "" && targetId != "local" {
			}
		}
	}

	// 3. Instance relationships (Subnet, SecurityGroup, NetworkInterface)
	for _, instanceId := range vpcIndex.InstanceIds {
		instance, err := p.dbClient.GetInstance(types.CloudID(providerName, instanceId))
		if err != nil || instance == nil {
			logger.Warnf("Failed to get instance %s for edge generation: %v", instanceId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, instance.SubnetID) {
			edge := createEdge(instanceId, instance.SubnetID, "LOCATED_IN")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
		for _, sgId := range instance.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edge := createEdge(instanceId, sgId, "USES_SECURITY_GROUP")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		for _, niId := range instance.InterfaceIDs {
			if contains(vpcIndex.NetworkInterfaceIds, niId) {
				edge := createEdge(instanceId, niId, "HAS_INTERFACE")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
	}

	// 4. NetworkInterface relationships (Subnet, SecurityGroup, PublicIP, Instance)
	for _, niId := range vpcIndex.NetworkInterfaceIds {
		ni, err := p.dbClient.GetNetworkInterface(types.CloudID(providerName, niId))
		if err != nil || ni == nil {
			logger.Warnf("Failed to get network interface %s for edge generation: %v", niId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, ni.SubnetID) {
			edge := createEdge(niId, ni.SubnetID, "LOCATED_IN")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
		for _, sgId := range ni.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edge := createEdge(niId, sgId, "USES_SECURITY_GROUP")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		if ni.PublicIP != "" {
			pubIPId := findPublicIPIdByIP(p.dbClient, providerName, ni.PublicIP, vpcIndex.PublicIpIds)
			if pubIPId != "" {
				edge := createEdge(niId, pubIPId, "HAS_PUBLIC_IP")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
	}

	// 5. NAT Gateway relationships (Subnet, PublicIP)
	for _, natId := range vpcIndex.NatGatewayIds {
		nat, err := p.dbClient.GetNATGateway(types.CloudID(providerName, natId))
		if err != nil || nat == nil {
			logger.Warnf("Failed to get NAT gateway %s for edge generation: %v", natId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, nat.SubnetId) {
			edge := createEdge(natId, nat.SubnetId, "LOCATED_IN")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
		pubIPId := findPublicIPIdByIP(p.dbClient, providerName, nat.PublicIp, vpcIndex.PublicIpIds)
		if pubIPId != "" {
			edge := createEdge(natId, pubIPId, "USES_PUBLIC_IP")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
	}

	// 6. VPCEndpoint relationships (Subnet, RouteTable, SecurityGroup)
	for _, vpceId := range vpcIndex.VpcEndpointIds {
		vpce, err := p.dbClient.GetVPCEndpoint(types.CloudID(providerName, vpceId))
		if err != nil || vpce == nil {
			logger.Warnf("Failed to get VPC Endpoint %s for edge generation: %v", vpceId, err)
			continue
		}
		for _, subnetId := range vpce.SubnetIds {
			if contains(vpcIndex.SubnetIds, subnetId) {
				edge := createEdge(vpceId, subnetId, "USES_SUBNET")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		for _, rtId := range vpce.RouteTableIds {
			if contains(vpcIndex.RouteTableIds, rtId) {
				edge := createEdge(vpceId, rtId, "MODIFIES_ROUTE_TABLE")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		for _, sgId := range vpce.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edge := createEdge(vpceId, sgId, "USES_SECURITY_GROUP")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
	}

	// 7. LB relationships (Subnet, SecurityGroup, Instance)
	for _, lbId := range vpcIndex.LbIds {
		lb, err := p.dbClient.GetLB(types.CloudID(providerName, lbId))
		if err != nil || lb == nil {
			logger.Warnf("Failed to get LB %s for edge generation: %v", lbId, err)
			continue
		}
		for _, subnetId := range lb.SubnetIDs {
			if contains(vpcIndex.SubnetIds, subnetId) {
				edge := createEdge(lbId, subnetId, "LOCATED_IN")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		for _, sgId := range lb.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edge := createEdge(lbId, sgId, "USES_SECURITY_GROUP")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
		for _, instanceId := range lb.InstanceIDs {
			if contains(vpcIndex.InstanceIds, instanceId) {
				edge := createEdge(lbId, instanceId, "LOAD_BALANCES_TO")
				if edge.SourceNodeID != "" {
					edges = append(edges, edge)
				}
			}
		}
	}

	// 8. Public IP relationships (Instance, NetworkInterface, NAT Gateway)
	for _, pubIPId := range vpcIndex.PublicIpIds {
		pubIP, err := p.dbClient.GetPublicIP(types.CloudID(providerName, pubIPId))
		if err != nil || pubIP == nil {
			logger.Warnf("Failed to get Public IP %s for edge generation: %v", pubIPId, err)
			continue
		}
		if pubIP.InstanceId != "" && contains(vpcIndex.InstanceIds, pubIP.InstanceId) {
			edge := createEdge(pubIPId, pubIP.InstanceId, "ASSOCIATED_WITH")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
		if pubIP.NetworkInterfaceId != "" && contains(vpcIndex.NetworkInterfaceIds, pubIP.NetworkInterfaceId) {
			edge := createEdge(pubIPId, pubIP.NetworkInterfaceId, "ASSOCIATED_WITH")
			if edge.SourceNodeID != "" {
				edges = append(edges, edge)
			}
		}
	}

	// 9. VPN Concentrator / VGW relationships
	for _, vgwId := range vpcIndex.VpnConcentratorIds {
		_, err := p.dbClient.GetVPNConcentrator(types.CloudID(providerName, vgwId))
		if err != nil || vgwId == "" {
			logger.Warnf("Failed to get VPN Concentrator %s for edge generation: %v", vgwId, err)
			continue
		}
	}

	validEdges := make([]types.VpcGraphEdge, 0, len(edges))
	for _, edge := range edges {
		if edge.SourceNodeID != "" && edge.TargetNodeID != "" {
			validEdges = append(validEdges, edge)
		}
	}

	return validEdges, nil
}

func (p *providerWithDB) GetVpcConnectivityGraph(ctx context.Context, params *infrapb.GetVpcConnectivityGraphRequest) ([]types.VpcGraphNode, []types.VpcGraphEdge, error) {
	nodesReq := &infrapb.ListVpcGraphNodesRequest{
		Provider:  params.Provider,
		AccountId: params.AccountId,
		Region:    params.Region,
		VpcId:     params.VpcId,
		Creds:     params.Creds,
	}
	nodes, err := p.ListVpcGraphNodes(ctx, nodesReq)
	if err != nil {
		p.logger.Errorf("Failed to list VPC graph nodes for graph request (VPC %s): %v", params.VpcId, err)
		return nil, nil, fmt.Errorf("failed to get nodes: %w", err)
	}

	edgesReq := &infrapb.ListVpcGraphEdgesRequest{
		Provider:  params.Provider,
		AccountId: params.AccountId,
		Region:    params.Region,
		VpcId:     params.VpcId,
		Creds:     params.Creds,
	}
	edges, err := p.ListVpcGraphEdges(ctx, edgesReq)
	if err != nil {
		p.logger.Errorf("Failed to list VPC graph edges for graph request (VPC %s): %v", params.VpcId, err)
		return nil, nil, fmt.Errorf("failed to get edges: %w", err)
	}

	return nodes, edges, nil
}

func (p *providerWithDB) GetInstanceConnectivityGraph(ctx context.Context, params *infrapb.GetInstanceConnectivityGraphRequest) ([]types.InstanceGraphNode, []types.InstanceGraphEdge, error) {
	// Delegate to the graph building logic
	return p.buildInstanceGraph(ctx, params)
}

// buildInstanceGraph constructs the nodes and edges for a specific instance's connectivity.
func (p *providerWithDB) buildInstanceGraph(ctx context.Context, params *infrapb.GetInstanceConnectivityGraphRequest) ([]types.InstanceGraphNode, []types.InstanceGraphEdge, error) {
	providerName := params.GetProvider()
	accountId := params.GetAccountId()
	region := params.GetRegion()
	instanceId := params.GetInstanceId()
	logger := p.logger.WithField("instanceId", instanceId)

	nodesMap := make(map[string]types.InstanceGraphNode)
	edgesMap := make(map[string]types.InstanceGraphEdge)

	// Helper to create node properties, simplifying and adding self link and rules
	createProps := func(resource interface{}) map[string]string {
		props := make(map[string]string)

		// Use type assertion to populate common/specific properties AND SelfLink
		switch v := resource.(type) {
		case *types.Instance:
			props["privateIP"] = v.PrivateIP
			props["publicIP"] = v.PublicIP
			props["state"] = v.State
			props["type"] = v.Type
			props["zone"] = v.Zone
			props["subnetID"] = v.SubnetID
			props["vpcID"] = v.VPCID
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.Subnet:
			props["cidrBlock"] = v.CidrBlock
			props["zone"] = v.Zone
			props["vpcID"] = v.VpcId
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.RouteTable:
			props["name"] = v.Name
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.SecurityGroup:
			props["name"] = v.Name
			props["vpcID"] = v.VpcID
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
			// Add Rules
			for i, rule := range v.Rules {
				key := fmt.Sprintf("rule_%d", i+1)
				props[key] = formatSecurityGroupRule(rule)
			}
		case *types.NetworkInterface:
			props["macAddress"] = v.MacAddress
			props["status"] = v.Status
			props["publicIP"] = v.PublicIP
			if len(v.PrivateIPs) > 0 {
				props["privateIP"] = v.PrivateIPs[0]
			}
			props["subnetID"] = v.SubnetID
			props["vpcID"] = v.VPCID
			props["instanceID"] = v.InstanceID
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.ACL:
			props["vpcID"] = v.VpcID
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
				// Add Rules
				for i, rule := range v.Rules {
					key := fmt.Sprintf("rule_%d", i+1)
					props[key] = formatACLRule(rule)
				}
			}
		case *types.IGW:
			props["state"] = v.State
			props["attachedVpcId"] = v.AttachedVpcId
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.NATGateway:
			props["publicIP"] = v.PublicIp
			props["privateIP"] = v.PrivateIp
			props["state"] = v.State
			props["subnetID"] = v.SubnetId
			props["vpcID"] = v.VpcId
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.VPCEndpoint:
			props["serviceName"] = v.ServiceName
			props["type"] = v.Type
			props["state"] = v.State
			props["vpcID"] = v.VPCId
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		case *types.VPNConcentrator:
			props["asn"] = fmt.Sprintf("%d", v.ASN)
			props["state"] = v.State
			props["type"] = v.Type
			props["vpcID"] = v.VpcID
			if v.SelfLink != "" {
				props["self_link"] = v.SelfLink
			}
		}

		return props
	}

	// Helper to add a node, preventing duplicates
	addNode := func(node types.InstanceGraphNode) {
		if _, exists := nodesMap[node.ID]; !exists {
			nodesMap[node.ID] = node
		}
	}

	// Helper to add an edge, preventing duplicates
	addEdge := func(edge types.InstanceGraphEdge) {
		edgeKey := fmt.Sprintf("%s->%s:%s", edge.SourceNodeID, edge.TargetNodeID, edge.RelationshipType)
		if _, exists := edgesMap[edgeKey]; !exists {
			edgesMap[edgeKey] = edge
		}
	}

	// --- Start Fetching Data ---
	// 1. Fetch the target Instance
	instance, err := p.dbClient.GetInstance(types.CloudID(providerName, instanceId))
	if err != nil || instance == nil {
		logger.Errorf("Failed to get target instance: %v", err)
		return nil, nil, fmt.Errorf("failed to get instance %s: %w", instanceId, err)
	}
	instanceNode := types.InstanceGraphNode{
		ID:           instance.ID,
		ResourceType: types.InstanceType,
		Name:         instance.Name,
		Properties:   createProps(instance),
		Provider:     providerName, AccountID: accountId, Region: region,
	}
	addNode(instanceNode)

	// 2. Fetch Network Interfaces associated with the instance
	var networkInterfaces []*types.NetworkInterface
	if len(instance.InterfaceIDs) > 0 {
		for _, niID := range instance.InterfaceIDs {
			ni, niErr := p.dbClient.GetNetworkInterface(types.CloudID(providerName, niID))
			if niErr == nil && ni != nil {
				networkInterfaces = append(networkInterfaces, ni)
			} else {
				logger.Warnf("Failed to get network interface %s for instance %s: %v", niID, instanceId, niErr)
			}
		}
	} else {
		logger.Warnf("Instance %s does not have NetworkInterfaceIDs field populated, cannot reliably fetch NIs.", instanceId)
	}

	primarySubnetId := instance.SubnetID
	for _, ni := range networkInterfaces {
		niNode := types.InstanceGraphNode{
			ID:           ni.ID,
			ResourceType: types.NetworkInterfaceType,
			Name:         ni.Name,
			Properties:   createProps(ni),
			Provider:     providerName, AccountID: accountId, Region: region,
		}
		addNode(niNode)
		addEdge(types.InstanceGraphEdge{SourceNodeID: instanceId, TargetNodeID: ni.ID, RelationshipType: "HAS_INTERFACE", Provider: providerName, AccountID: accountId, Region: region})

		if ni.SubnetID != "" {
			primarySubnetId = ni.SubnetID
		}

		for _, sgID := range ni.SecurityGroupIDs {
			sg, sgErr := p.dbClient.GetSecurityGroup(types.CloudID(providerName, sgID))
			if sgErr == nil && sg != nil {
				sgNode := types.InstanceGraphNode{
					ID:           sg.ID,
					ResourceType: types.SecurityGroupType,
					Name:         sg.Name,
					Properties:   createProps(sg),
					Provider:     providerName, AccountID: accountId, Region: region,
				}
				addNode(sgNode)
				addEdge(types.InstanceGraphEdge{SourceNodeID: ni.ID, TargetNodeID: sgID, RelationshipType: "USES_SECURITY_GROUP", Provider: providerName, AccountID: accountId, Region: region})
			} else {
				logger.Warnf("Failed to get security group %s for NI %s: %v", sgID, ni.ID, sgErr)
			}
		}
	}

	if primarySubnetId == "" {
		logger.Errorf("Could not determine subnet for instance %s", instanceId)
		return nil, nil, fmt.Errorf("could not determine subnet for instance %s", instanceId)
	}
	subnet, err := p.dbClient.GetSubnet(types.CloudID(providerName, primarySubnetId))
	if err != nil || subnet == nil {
		logger.Errorf("Failed to get subnet %s for instance %s: %v", primarySubnetId, instanceId, err)
		return nil, nil, fmt.Errorf("failed to get subnet %s: %w", primarySubnetId, err)
	}
	subnetNode := types.InstanceGraphNode{
		ID:           subnet.SubnetId,
		ResourceType: types.SubnetType,
		Name:         subnet.Name,
		Properties:   createProps(subnet),
		Provider:     providerName, AccountID: accountId, Region: region,
	}
	addNode(subnetNode)
	addEdge(types.InstanceGraphEdge{SourceNodeID: instanceId, TargetNodeID: subnet.SubnetId, RelationshipType: "LOCATED_IN", Provider: providerName, AccountID: accountId, Region: region})

	if len(subnet.RouteTableIds) == 0 || subnet.RouteTableIds[0] == "" {
		logger.Errorf("Subnet %s has no associated route table ID", subnet.SubnetId)
		return nil, nil, fmt.Errorf("subnet %s has no associated route table ID", subnet.SubnetId)
	}
	routeTableId := subnet.RouteTableIds[0]
	routeTable, err := p.dbClient.GetRouteTable(types.CloudID(providerName, routeTableId))
	if err != nil || routeTable == nil {
		logger.Errorf("Failed to get route table %s for subnet %s: %v", routeTableId, subnet.SubnetId, err)
		return nil, nil, fmt.Errorf("failed to get route table %s: %w", routeTableId, err)
	}
	rtNode := types.InstanceGraphNode{
		ID:           routeTable.ID,
		ResourceType: types.RouteTableType,
		Name:         routeTable.Name,
		Properties:   createProps(routeTable),
		Provider:     providerName, AccountID: accountId, Region: region,
	}
	addNode(rtNode)
	addEdge(types.InstanceGraphEdge{SourceNodeID: subnet.SubnetId, TargetNodeID: routeTable.ID, RelationshipType: "USES_ROUTE_TABLE", Provider: providerName, AccountID: accountId, Region: region})

	for _, route := range routeTable.Routes {
		targetId := route.Target
		if targetId == "" || targetId == "local" {
			continue
		}

		var targetNode *types.InstanceGraphNode
		var targetErr error

		if strings.HasPrefix(targetId, "igw-") {
			res, getErr := p.dbClient.GetIGW(types.CloudID(providerName, targetId))
			if getErr == nil && res != nil {
				targetNode = &types.InstanceGraphNode{ID: res.ID, ResourceType: types.IGWType, Name: res.Name, Properties: createProps(res), Provider: providerName, AccountID: accountId, Region: region}
			} else {
				targetErr = getErr
			}
		} else if strings.HasPrefix(targetId, "nat-") {
			res, getErr := p.dbClient.GetNATGateway(types.CloudID(providerName, targetId))
			if getErr == nil && res != nil {
				targetNode = &types.InstanceGraphNode{ID: res.ID, ResourceType: types.NATGatewayType, Name: res.Name, Properties: createProps(res), Provider: providerName, AccountID: accountId, Region: region}
			} else {
				targetErr = getErr
			}
		} else if strings.HasPrefix(targetId, "vpce-") {
			res, getErr := p.dbClient.GetVPCEndpoint(types.CloudID(providerName, targetId))
			if getErr == nil && res != nil {
				targetNode = &types.InstanceGraphNode{ID: res.ID, ResourceType: types.VPCEndpointType, Name: res.Name, Properties: createProps(res), Provider: providerName, AccountID: accountId, Region: region}
			} else {
				targetErr = getErr
			}
		} else if strings.HasPrefix(targetId, "vgw-") {
			res, getErr := p.dbClient.GetVPNConcentrator(types.CloudID(providerName, targetId))
			if getErr == nil && res != nil {
				targetNode = &types.InstanceGraphNode{ID: res.ID, ResourceType: types.VPNConcentratorType, Name: res.Name, Properties: createProps(res), Provider: providerName, AccountID: accountId, Region: region}
			} else {
				targetErr = getErr
			}
		} else if strings.HasPrefix(targetId, "eni-") {
			found := false
			for _, ni := range networkInterfaces {
				if ni.ID == targetId {
					targetNode = &types.InstanceGraphNode{ID: ni.ID}
					found = true
					break
				}
			}
			if !found {
				res, getErr := p.dbClient.GetNetworkInterface(types.CloudID(providerName, targetId))
				if getErr == nil && res != nil {
					targetNode = &types.InstanceGraphNode{ID: res.ID, ResourceType: types.NetworkInterfaceType, Name: res.Name, Properties: createProps(res), Provider: providerName, AccountID: accountId, Region: region}
				} else {
					targetErr = getErr
				}
			}
		}

		if targetNode != nil {
			if targetNode.ResourceType != "" {
				addNode(*targetNode)
			}
			addEdge(types.InstanceGraphEdge{SourceNodeID: routeTable.ID, TargetNodeID: targetId, RelationshipType: "ROUTES_TO", Provider: providerName, AccountID: accountId, Region: region})
		} else {
			logger.Warnf("Failed to get route target %s for route table %s: %v", targetId, routeTable.ID, targetErr)
		}
	}

	if len(networkInterfaces) == 0 && len(instance.SecurityGroupIDs) > 0 {
		for _, sgID := range instance.SecurityGroupIDs {
			sg, sgErr := p.dbClient.GetSecurityGroup(types.CloudID(providerName, sgID))
			if sgErr == nil && sg != nil {
				sgNode := types.InstanceGraphNode{
					ID:           sg.ID,
					ResourceType: types.SecurityGroupType,
					Name:         sg.Name,
					Properties:   createProps(sg),
					Provider:     providerName, AccountID: accountId, Region: region,
				}
				addNode(sgNode)
				addEdge(types.InstanceGraphEdge{SourceNodeID: instanceId, TargetNodeID: sgID, RelationshipType: "USES_SECURITY_GROUP", Provider: providerName, AccountID: accountId, Region: region})
			} else {
				logger.Warnf("Failed to get security group %s for instance %s: %v", sgID, instanceId, sgErr)
			}
		}
	}

	if len(subnet.NetworkAclIds) > 0 && subnet.NetworkAclIds[0] != "" {
		aclId := subnet.NetworkAclIds[0]
		acl, err := p.dbClient.GetACL(types.CloudID(providerName, aclId))
		if err == nil && acl != nil {
			aclNode := types.InstanceGraphNode{
				ID:           acl.ID,
				ResourceType: types.ACLType,
				Name:         acl.Name,
				Properties:   createProps(acl),
				Provider:     providerName, AccountID: accountId, Region: region,
			}
			addNode(aclNode)
			addEdge(types.InstanceGraphEdge{SourceNodeID: subnet.SubnetId, TargetNodeID: acl.ID, RelationshipType: "USES_ACL", Provider: providerName, AccountID: accountId, Region: region})
		} else {
			logger.Warnf("Failed to get ACL %s for subnet %s: %v", aclId, subnet.SubnetId, err)
		}
	}

	finalNodes := make([]types.InstanceGraphNode, 0, len(nodesMap))
	for _, node := range nodesMap {
		finalNodes = append(finalNodes, node)
	}
	finalEdges := make([]types.InstanceGraphEdge, 0, len(edgesMap))
	for _, edge := range edgesMap {
		finalEdges = append(finalEdges, edge)
	}

	logger.Infof("Generated instance graph for %s with %d nodes and %d edges", instanceId, len(finalNodes), len(finalEdges))
	return finalNodes, finalEdges, nil
}

func contains(slice []string, item string) bool {
	if item == "" {
		return false
	}
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func findPublicIPIdByIP(dbClient Client, providerName, ipAddress string, candidateIDs []string) string {
	if ipAddress == "" {
		return ""
	}
	for _, id := range candidateIDs {
		pubIP, err := dbClient.GetPublicIP(types.CloudID(providerName, id))
		if err == nil && pubIP != nil && pubIP.PublicIP == ipAddress && strings.EqualFold(pubIP.Provider, providerName) {
			return id
		}
	}
	return ""
}
