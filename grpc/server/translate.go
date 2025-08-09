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

package server

import (
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func grpcProtocolsAndPortToTypes(in map[string]*infrapb.Ports) types.ProtocolsAndPorts {
	out := make(types.ProtocolsAndPorts, len(in))
	for k, v := range in {
		ports := make([]string, 0, len(v.GetPorts()))
		ports = append(ports, v.GetPorts()...)
		out[k] = ports
	}
	return out
}

func typesAccountsToGrpc(in []types.Account) []*infrapb.Account {
	out := make([]*infrapb.Account, 0, len(in))
	for _, account := range in {
		out = append(out, &infrapb.Account{
			Provider: account.Provider,
			Id:       account.ID,
			Name:     account.Name,
		})
	}
	return out
}

func typesRegionsToGrpc(in []types.Region) []*infrapb.Region {
	out := make([]*infrapb.Region, 0, len(in))
	for _, region := range in {
		out = append(out, &infrapb.Region{
			Provider: region.Provider,
			Id:       region.ID,
			Name:     region.Name,
		})
	}
	return out
}

func typesInstanceToGrpc(in []types.Instance) []*infrapb.Instance {
	out := make([]*infrapb.Instance, 0, len(in))
	for _, instance := range in {
		var createdAtPb *timestamppb.Timestamp
		var updatedAtPb *timestamppb.Timestamp
		if instance.CreatedAt != nil {
			createdAtPb = timestamppb.New(*instance.CreatedAt)
		}
		if instance.UpdatedAt != nil {
			// If UpdatedAt is set, we can use it as the last sync time
			updatedAtPb = timestamppb.New(*instance.UpdatedAt)
		}
		out = append(out, &infrapb.Instance{
			Id:               instance.ID,
			Name:             instance.Name,
			PublicIP:         instance.PublicIP,
			PrivateIP:        instance.PrivateIP,
			SubnetID:         instance.SubnetID,
			VpcId:            instance.VPCID,
			Zone:             instance.Zone,
			Project:          instance.Project,
			Region:           instance.Region,
			Labels:           instance.Labels,
			Provider:         instance.Provider,
			AccountId:        instance.AccountID,
			State:            instance.State,
			Type:             instance.Type,
			SecurityGroupIds: instance.SecurityGroupIDs,
			InterfaceIds:     instance.InterfaceIDs,
			LastSyncTime:     instance.LastSyncTime,
			SelfLink:         instance.SelfLink,
			CreatedAt:        createdAtPb,
			UpdatedAt:        updatedAtPb,
			SecurityStatus:   typesInstanceSecurityStatusToGrpc(instance.SecurityStatus),
		})
	}
	return out
}

func typesSubnetsToGrpc(in []types.Subnet) []*infrapb.Subnet {
	out := make([]*infrapb.Subnet, 0, len(in))
	for _, subnet := range in {
		var createdAtPb *timestamppb.Timestamp
		if subnet.CreatedAt != nil {
			createdAtPb = timestamppb.New(*subnet.CreatedAt)
		}
		var updatedAtPb *timestamppb.Timestamp
		if subnet.UpdatedAt != nil {
			updatedAtPb = timestamppb.New(*subnet.UpdatedAt)
		}

		out = append(out, &infrapb.Subnet{
			Id:            subnet.SubnetId,
			Name:          subnet.Name,
			CidrBlock:     subnet.CidrBlock,
			VpcId:         subnet.VpcId,
			Zone:          subnet.Zone,
			Region:        subnet.Region,
			Labels:        subnet.Labels,
			Provider:      subnet.Provider,
			AccountId:     subnet.AccountID,
			LastSyncTime:  subnet.LastSyncTime,
			SelfLink:      subnet.SelfLink,
			RouteTableIds: subnet.RouteTableIds,
			AclIds:        subnet.NetworkAclIds,
			CreatedAt:     createdAtPb,
			UpdatedAt:     updatedAtPb,
		})
	}
	return out
}

func typesVpcsToGrpc(in []types.VPC) []*infrapb.VPC {
	out := make([]*infrapb.VPC, 0, len(in))
	for _, vpc := range in {
		out = append(out, &infrapb.VPC{
			Id:           vpc.ID,
			Name:         vpc.Name,
			Region:       vpc.Region,
			Labels:       vpc.Labels,
			Ipv4Cidr:     vpc.IPv4CIDR,
			Ipv6Cidr:     vpc.IPv6CIDR,
			Provider:     vpc.Provider,
			AccountId:    vpc.AccountID,
			LastSyncTime: vpc.LastSyncTime,
			SelfLink:     vpc.SelfLink,
			Project:      vpc.Project,
		})
	}
	return out
}

func typesRoutersToGrpc(in []types.Router) []*infrapb.Router {
	out := make([]*infrapb.Router, 0, len(in))
	for _, router := range in {
		out = append(out, &infrapb.Router{
			Id:              router.ID,
			Name:            router.Name,
			VpcId:           router.VPCId,
			Asn:             router.ASN,
			AdvertisedRange: router.AdvertisedRange,
			AdvertisedGroup: router.AdvertisedGroup,
			SubnetId:        router.SubnetId,
			Provider:        router.Provider,
			Region:          router.Region,
			State:           router.State,
			Labels:          router.Labels,
			AccountId:       router.AccountID,
			CreatedAt:       timestamppb.New(router.CreatedAt),
			LastSyncTime:    router.LastSyncTime,
			SelfLink:        router.SelfLink,
		})
	}
	return out
}

func typesIGWsToGrpc(in []types.IGW) []*infrapb.IGW {
	out := make([]*infrapb.IGW, 0, len(in))
	for _, igw := range in {
		out = append(out, &infrapb.IGW{
			Id:            igw.ID,
			Name:          igw.Name,
			AttachedVpcId: igw.AttachedVpcId,
			Provider:      igw.Provider,
			Region:        igw.Region,
			State:         igw.State,
			Labels:        igw.Labels,
			AccountId:     igw.AccountID,
			CreatedAt:     igw.CreatedAt,
			LastSyncTime:  igw.LastSyncTime,
			SelfLink:      igw.SelfLink,
		})
	}
	return out
}

func typesVPCEndpointsToGrpc(in []types.VPCEndpoint) []*infrapb.VPCEndpoint {
	out := make([]*infrapb.VPCEndpoint, 0, len(in))
	for _, vpce := range in {
		var createdAt *timestamppb.Timestamp
		if vpce.CreatedAt != nil {
			createdAt = timestamppb.New(*vpce.CreatedAt)
		}
		out = append(out, &infrapb.VPCEndpoint{
			Id:               vpce.ID,
			Name:             vpce.Name,
			VpcId:            vpce.VPCId,
			Provider:         vpce.Provider,
			Region:           vpce.Region,
			State:            vpce.State,
			Labels:           vpce.Labels,
			AccountId:        vpce.AccountID,
			RouteTableIds:    vpce.RouteTableIds,
			SubnetIds:        vpce.SubnetIds,
			SecurityGroupIds: vpce.SecurityGroupIDs,
			ServiceName:      vpce.ServiceName,
			Type:             vpce.Type,
			CreatedAt:        createdAt,
			LastSyncTime:     vpce.LastSyncTime,
			SelfLink:         vpce.SelfLink,
		})
	}
	return out
}

func typesNATGatewaysToGrpc(in []types.NATGateway) []*infrapb.NATGateway {
	out := make([]*infrapb.NATGateway, 0, len(in))
	for _, gateway := range in {
		out = append(out, &infrapb.NATGateway{
			Id:           gateway.ID,
			Name:         gateway.Name,
			Provider:     gateway.Provider,
			VpcId:        gateway.VpcId,
			Region:       gateway.Region,
			State:        gateway.State,
			Labels:       gateway.Labels,
			AccountId:    gateway.AccountID,
			PublicIp:     gateway.PublicIp,
			PrivateIp:    gateway.PrivateIp,
			SubnetId:     gateway.SubnetId,
			CreatedAt:    timestamppb.New(gateway.CreatedAt),
			LastSyncTime: gateway.LastSyncTime,
			SelfLink:     gateway.SelfLink,
		})
	}
	return out
}

func typesACLsToGrpc(in []types.ACL) []*infrapb.ACL {
	out := make([]*infrapb.ACL, 0, len(in))
	for _, acl := range in {
		rules := make([]*infrapb.ACL_ACLRule, 0, len(acl.Rules))
		for _, r := range acl.Rules {
			rules = append(rules, &infrapb.ACL_ACLRule{
				Number:            int32(r.Number),
				Protocol:          r.Protocol,
				PortRange:         r.PortRange,
				SourceRanges:      r.SourceRanges,
				DestinationRanges: r.DestinationRanges,
				Action:            r.Action,
				Direction:         r.Direction,
			})
		}
		out = append(out, &infrapb.ACL{
			Provider:     acl.Provider,
			Id:           acl.ID,
			Name:         acl.Name,
			VpcId:        acl.VpcID,
			Region:       acl.Region,
			AccountId:    acl.AccountID,
			Labels:       acl.Labels,
			Rules:        rules,
			LastSyncTime: acl.LastSyncTime,
			SelfLink:     acl.SelfLink,
		})
	}
	return out
}

func typesSgsToGrpc(in []types.SecurityGroup) []*infrapb.SecurityGroup {
	out := make([]*infrapb.SecurityGroup, 0, len(in))
	for _, acl := range in {
		rules := make([]*infrapb.SecurityGroup_SecurityGroupRule, 0, len(acl.Rules))
		for _, r := range acl.Rules {
			rules = append(rules, &infrapb.SecurityGroup_SecurityGroupRule{
				Protocol:  r.Protocol,
				PortRange: r.PortRange,
				Source:    r.Source,
				Direction: r.Direction,
			})
		}
		out = append(out, &infrapb.SecurityGroup{
			Provider:     acl.Provider,
			Id:           acl.ID,
			Name:         acl.Name,
			VpcId:        acl.VpcID,
			Region:       acl.Region,
			AccountId:    acl.AccountID,
			Labels:       acl.Labels,
			Rules:        rules,
			LastSyncTime: acl.LastSyncTime,
			SelfLink:     acl.SelfLink,
			RiskStatus:   typesSecurityGroupRiskStatusToGrpc(acl.SecurityStatus),
		})
	}
	return out
}

func typesRouteTableToGrpc(in []types.RouteTable) []*infrapb.RouteTable {
	out := make([]*infrapb.RouteTable, 0, len(in))
	for _, rt := range in {
		routes := make([]*infrapb.RouteTable_Route, 0, len(rt.Routes))
		for _, r := range rt.Routes {
			routes = append(routes, &infrapb.RouteTable_Route{
				Destination: r.Destination,
				Target:      r.Target,
				Status:      r.Status,
			})
		}
		fmt.Printf("DEBUG: RouteTable Subnets %v\n", rt.SubnetIds)

		out = append(out, &infrapb.RouteTable{
			Provider:     rt.Provider,
			Id:           rt.ID,
			Name:         rt.Name,
			VpcId:        rt.VpcID,
			Region:       rt.Region,
			AccountId:    rt.AccountID,
			Labels:       rt.Labels,
			Routes:       routes,
			SubnetIds:    rt.SubnetIds,
			IgwIds:       rt.IGWIds,
			NgwIds:       rt.NGWIds,
			TgwIds:       rt.TGWIds,
			LastSyncTime: rt.LastSyncTime,
			SelfLink:     rt.SelfLink,
		})
	}
	fmt.Printf("DEBUG: RouteTable %v\n", out)
	return out
}

func typesPublicIPsToGrpc(in []types.PublicIP) []*infrapb.PublicIP {
	out := make([]*infrapb.PublicIP, 0, len(in))
	for _, publicIP := range in {
		var createdAt *timestamppb.Timestamp
		if publicIP.CreatedAt != nil {
			createdAt = timestamppb.New(*publicIP.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if publicIP.UpdatedAt != nil {
			updatedAt = timestamppb.New(*publicIP.UpdatedAt)
		}
		out = append(out, &infrapb.PublicIP{
			Provider:           publicIP.Provider,
			Id:                 publicIP.ID,
			VpcId:              publicIP.VPCId,
			Region:             publicIP.Region,
			PublicIp:           publicIP.PublicIP,
			InstanceId:         publicIP.InstanceId,
			PrivateIp:          publicIP.PrivateIP,
			AccountId:          publicIP.AccountID,
			Type:               publicIP.Type,
			Labels:             publicIP.Labels,
			SelfLink:           publicIP.SelfLink,
			Byoip:              publicIP.Byoip,
			Project:            publicIP.Project,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
			LastSyncTime:       publicIP.LastSyncTime,
			NetworkInterfaceId: publicIP.NetworkInterfaceId,
		})
	}
	return out
}

func typesLBToGrpc(in []types.LB) []*infrapb.LB {
	if len(in) == 0 {
		fmt.Println("No LBs found")
		return nil
	}
	out := make([]*infrapb.LB, 0, len(in))

	for _, lb := range in {
		out = append(out, &infrapb.LB{
			Id:                     lb.ID,
			Name:                   lb.Name,
			VpcId:                  lb.VPCID,
			DnsName:                lb.DNSName,
			Provider:               lb.Provider,
			AccountId:              lb.AccountID,
			Listeners:              typesLBListenersToGrpc(lb.Listeners),
			LoadBalancerType:       lb.Type,
			Scheme:                 lb.Scheme,
			Region:                 lb.Region,
			InstanceIds:            lb.InstanceIDs,
			SubnetIds:              lb.SubnetIDs,
			SecurityGroupIds:       lb.SecurityGroupIDs,
			TargetGroupIds:         lb.TargetGroupIDs,
			CrossZoneLoadBalancing: lb.CrossZoneLoadBalancing,
			AccessLogsEnabled:      lb.AccessLogsEnabled,
			LoggingBucket:          lb.LoggingBucket,
			PublicIpAddresses:      lb.PublicIPs,
			PrivateIpAddresses:     lb.PrivateIPs,
			State:                  lb.State,
			IpAddressType:          lb.IPAddressType,
			Zone:                   lb.Zone,
			Labels:                 lb.Labels,
			Project:                lb.Project,
			CreatedAt:              timestamppb.New(lb.CreatedAt),
			SecurityStatus:         typesLoadBalancerSecurityStatusToGrpc(lb.SecurityStatus),
		})
	}
	return out
}

func typesLBListenersToGrpc(in []types.LBListener) []*infrapb.LB_Listener {
	out := make([]*infrapb.LB_Listener, 0, len(in))
	for _, listener := range in {
		out = append(out, &infrapb.LB_Listener{
			ListenerId:    listener.ListenerID,
			Protocol:      listener.Protocol,
			Port:          int32(listener.Port),
			TargetGroupId: listener.TargetGroupID,
		})
	}
	return out
}

func typesNetworkInterfacesToGrpc(in []types.NetworkInterface) []*infrapb.NetworkInterface {
	out := make([]*infrapb.NetworkInterface, len(in))
	for i, ni := range in {
		out[i] = &infrapb.NetworkInterface{
			Id:         ni.ID,
			Name:       ni.Name,
			Provider:   ni.Provider,
			AccountId:  ni.AccountID,
			VpcId:      ni.VPCID,
			InstanceId: ni.InstanceID,
			MacAddress: ni.MacAddress,
			DnsName:    ni.PublicDNSName,
			//Pr: ni.PrivateDNSName,
			Status:        ni.Status,
			LastSyncTime:  ni.LastSyncTime,
			SubnetId:      ni.SubnetID,
			InterfaceType: ni.InterfaceType,
			//AvailabilityZone:   ni.AvailabilityZone,
			Region:           ni.Region,
			PrivateIps:       ni.PrivateIPs,
			PublicIp:         ni.PublicIP,
			SecurityGroupIds: ni.SecurityGroupIDs,
		}
	}
	return out
}

func typesKeyPairsToGrpc(in []types.KeyPair) []*infrapb.KeyPair {
	out := make([]*infrapb.KeyPair, len(in))
	for i, kp := range in {
		out[i] = &infrapb.KeyPair{
			Id:                    kp.ID,
			Provider:              kp.Provider,
			AccountId:             kp.AccountID,
			Name:                  kp.Name,
			PrivateKeyFingerprint: kp.Fingerprint,
			PublicKey:             kp.PublicKey,
			CreatedAt:             timestamppb.New(kp.CreatedAt),
			Labels:                kp.Labels,
			InstanceIds:           kp.InstanceIds,
			Region:                kp.Region,
			KeyPairType:           kp.KeyPairType,
		}
	}
	return out
}

func typesVPNConcentratorsToGrpc(in []types.VPNConcentrator) []*infrapb.VPNConcentrator {
	out := make([]*infrapb.VPNConcentrator, 0, len(in))
	for _, vpnc := range in {
		out = append(out, &infrapb.VPNConcentrator{
			Id:           vpnc.ID,
			Name:         vpnc.Name,
			Provider:     vpnc.Provider,
			AccountId:    vpnc.AccountID,
			VpcId:        vpnc.VpcID,
			Region:       vpnc.Region,
			State:        vpnc.State,
			Type:         vpnc.Type,
			Asn:          vpnc.ASN,
			Labels:       vpnc.Labels,
			CreatedAt:    timestamppb.New(vpnc.CreatedAt),
			LastSyncTime: vpnc.LastSyncTime,
			SelfLink:     vpnc.SelfLink,
		})
	}
	return out
}

func typesVPCIndexToGrpc(in types.VPCIndex) *infrapb.VPCIndex {
	return &infrapb.VPCIndex{
		VpcId:               in.VpcId,
		InstanceIds:         in.InstanceIds,
		SubnetIds:           in.SubnetIds,
		AclIds:              in.AclIds,
		SecurityGroupIds:    in.SecurityGroupIds,
		RouteTableIds:       in.RouteTableIds,
		LbIds:               in.LbIds,
		IgwIds:              in.IgwIds,
		NatGatewayIds:       in.NatGatewayIds,
		NetworkInterfaceIds: in.NetworkInterfaceIds,
		KeyPairIds:          in.KeyPairIds,
		PublicIpIds:         in.PublicIpIds,
		VpcEndpointIds:      in.VpcEndpointIds,
		VpnConcentratorIds:  in.VpnConcentratorIds,
		RouterIds:           in.RouterIds,
		SecurityRisks:       typesVPCSecurityRisksToGrpc(in.SecurityRisks),
	}
}

func typesVpcGraphNodesToGrpc(in []types.VpcGraphNode) []*infrapb.VpcGraphNode {
	if in == nil {
		return nil
	}
	out := make([]*infrapb.VpcGraphNode, len(in))
	for i, node := range in {
		out[i] = &infrapb.VpcGraphNode{
			Id:           node.ID,
			ResourceType: node.ResourceType,
			Name:         node.Name,
			Properties:   node.Properties,
			Provider:     node.Provider,
			AccountId:    node.AccountId,
			Region:       node.Region,
		}
	}
	return out
}

func typesVpcConnectionGraphNodeToGrpc(in *types.VpcConnectionGraphNode) *infrapb.VpcConnectionGraphNode {
	if in == nil {
		return nil
	}

	nodeType := infrapb.VpcConnectionGraphNode_NODE_TYPE_UNSPECIFIED
	switch in.NodeType {
	case "vpc":
		nodeType = infrapb.VpcConnectionGraphNode_VPC
	case "transit_gateway":
		nodeType = infrapb.VpcConnectionGraphNode_TRANSIT_GATEWAY
	case "vpc_endpoint":
		nodeType = infrapb.VpcConnectionGraphNode_VPC_ENDPOINT
	case "transit_vpc":
		nodeType = infrapb.VpcConnectionGraphNode_TRANSIT_VPC
	case "internet_gateway":
		nodeType = infrapb.VpcConnectionGraphNode_INTERNET_GATEWAY
	case "vpn_gateway":
		nodeType = infrapb.VpcConnectionGraphNode_VPN_GATEWAY
	case "nat_gateway":
		nodeType = infrapb.VpcConnectionGraphNode_NAT_GATEWAY
	}

	return &infrapb.VpcConnectionGraphNode{
		Id:         in.Id,
		Name:       in.Name,
		NodeType:   nodeType,
		Provider:   in.Provider,
		AccountId:  in.AccountId,
		Region:     in.Region,
		Properties: in.Properties,
		Labels:     in.Labels,
	}
}

func typesVpcConnectionGraphNodesToGrpc(in []*types.VpcConnectionGraphNode) []*infrapb.VpcConnectionGraphNode {
	if in == nil {
		return nil
	}
	out := make([]*infrapb.VpcConnectionGraphNode, len(in))
	for i, node := range in {
		out[i] = typesVpcConnectionGraphNodeToGrpc(node)
	}
	return out
}

func typesVpcConnectionGraphEdgeToGrpc(in *types.VpcConnectionGraphEdge) *infrapb.VpcConnectionGraphEdge {
	if in == nil {
		return nil
	}

	connType := infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_UNSPECIFIED
	switch in.ConnectionType {
	case "peering":
		connType = infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_PEERING
	case "transit_gateway":
		connType = infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_TRANSIT_GATEWAY
	case "vpc_endpoint":
		connType = infrapb.VpcConnectionType_VPC_ENDPOINT
	case "transit_vpc":
		connType = infrapb.VpcConnectionType_TRANSIT_VPC
	}

	return &infrapb.VpcConnectionGraphEdge{
		Id:             in.Id,
		SourceNodeId:   in.SourceNodeId,
		TargetNodeId:   in.TargetNodeId,
		ConnectionType: connType,
		Status:         in.Status,
		Bidirectional:  in.Bidirectional,
		AccountId:      in.AccountId,
		Region:         in.Region,
		Provider:       in.Provider,
		RouteTableIds:  in.RouteTableIds,
		Properties:     in.Properties,
	}
}

func typesVpcConnectionGraphEdgesToGrpc(in []*types.VpcConnectionGraphEdge) []*infrapb.VpcConnectionGraphEdge {
	if in == nil {
		return nil
	}
	out := make([]*infrapb.VpcConnectionGraphEdge, len(in))
	for i, edge := range in {
		out[i] = typesVpcConnectionGraphEdgeToGrpc(edge)
	}
	return out
}

func typesVpcGraphEdgesToGrpc(in []types.VpcGraphEdge) []*infrapb.VpcGraphEdge {
	if in == nil {
		return nil
	}
	out := make([]*infrapb.VpcGraphEdge, len(in))
	for i, edge := range in {
		out[i] = &infrapb.VpcGraphEdge{
			SourceNodeId:     edge.SourceNodeID,
			TargetNodeId:     edge.TargetNodeID,
			RelationshipType: edge.RelationshipType,
			Provider:         edge.Provider,
			AccountId:        edge.AccountId,
			Region:           edge.Region,
		}
	}
	return out
}

func typesVpcInternalGraphToGrpc(in *types.VpcInternalGraph) *infrapb.VpcInternalGraph {
	if in == nil {
		return nil
	}

	out := &infrapb.VpcInternalGraph{
		Nodes:        typesVpcGraphNodesToGrpc(in.Nodes),
		Edges:        typesVpcGraphEdgesToGrpc(in.Edges),
		AccountId:    in.AccountId,
		Region:       in.Region,
		Provider:     in.Provider,
		Labels:       in.Labels,
		LastSyncTime: in.LastSyncTime,
	}
	return out
}

func typesVpcConnectionGraphToGrpc(in *types.VpcConnectionGraph) *infrapb.VpcConnectionGraph {
	if in == nil {
		return nil
	}

	out := &infrapb.VpcConnectionGraph{
		Nodes:        typesVpcConnectionGraphNodesToGrpc(in.Nodes),
		Edges:        typesVpcConnectionGraphEdgesToGrpc(in.Edges),
		Provider:     in.Provider,
		AccountId:    in.Provider,   // Switched from Accounts list to single account
		Region:       in.Regions[0], // Taking first region since proto only supports single region
		Labels:       in.Labels,
		LastSyncTime: in.LastSyncTime,
	}

	if in.SrcVpcGraph != nil {
		out.SrcVpcGraph = typesVpcInternalGraphToGrpc(in.SrcVpcGraph)
	}

	if in.DestVpcGraph != nil {
		out.DestVpcGraph = typesVpcInternalGraphToGrpc(in.DestVpcGraph)
	}

	return out
}

func typesVpcConnectionsToGrpc(in []*types.VPCConnection) []*infrapb.VpcConnection {
	if in == nil {
		return nil
	}
	out := make([]*infrapb.VpcConnection, len(in))
	for i, conn := range in {
		out[i] = &infrapb.VpcConnection{
			Id:             conn.ID,
			Name:           conn.Name,
			Provider:       conn.Provider,
			AccountId:      conn.Account,
			Region:         conn.Region,
			VpcId_1:        conn.FromVpcId,
			Vpc_1AccountId: conn.FromVpcAccountId,
			Vpc_1Region:    conn.FromVpcRegion,
			VpcId_2:        conn.ToVpcId,
			Vpc_2AccountId: conn.ToVpcAccountId,
			Vpc_2Region:    conn.ToVpcRegion,
			Status:         conn.Status,
		}

		// Set connection type
		switch conn.ConnectionType {
		case "transit_gateway":
			out[i].ConnectionType = infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_TRANSIT_GATEWAY
		case "peering":
			out[i].ConnectionType = infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_PEERING
		case "vpc_endpoint":
			out[i].ConnectionType = infrapb.VpcConnectionType_VPC_ENDPOINT
		case "transit_vpc":
			out[i].ConnectionType = infrapb.VpcConnectionType_TRANSIT_VPC
		default:
			out[i].ConnectionType = infrapb.VpcConnectionType_VPC_CONNECTION_TYPE_UNSPECIFIED
		}
	}
	return out
}

// Security status translation functions

func typesVPCSecurityRisksToGrpc(in *types.VPCSecurityRisks) *infrapb.VPCSecurityRisks {
	if in == nil {
		return nil
	}

	var lastRiskAnalysisPb *timestamppb.Timestamp
	if in.LastRiskAnalysis != nil {
		lastRiskAnalysisPb = timestamppb.New(*in.LastRiskAnalysis)
	}

	return &infrapb.VPCSecurityRisks{
		CriticalInstanceIds:        in.CriticalInstanceIds,
		CriticalSecurityGroupIds:   in.CriticalSecurityGroupIds,
		CriticalLbIds:              in.CriticalLbIds,
		HighRiskInstanceIds:        in.HighRiskInstanceIds,
		HighRiskSecurityGroupIds:   in.HighRiskSecurityGroupIds,
		HighRiskSubnetIds:          in.HighRiskSubnetIds,
		MediumRiskInstanceIds:      in.MediumRiskInstanceIds,
		MediumRiskSecurityGroupIds: in.MediumRiskSecurityGroupIds,
		MediumRiskAclIds:           in.MediumRiskAclIds,
		UntaggedInstanceIds:        in.UntaggedInstanceIds,
		UntaggedSecurityGroupIds:   in.UntaggedSecurityGroupIds,
		UntaggedSubnetIds:          in.UntaggedSubnetIds,
		UntaggedLbIds:              in.UntaggedLbIds,
		IsolatedSubnetIds:          in.IsolatedSubnetIds,
		OverExposedSubnetIds:       in.OverExposedSubnetIds,
		LastRiskAnalysis:           lastRiskAnalysisPb,
	}
}

func typesInstanceSecurityStatusToGrpc(in *types.InstanceSecurityStatus) *infrapb.InstanceSecurityStatus {
	if in == nil {
		return nil
	}

	riskLevel := infrapb.InstanceSecurityStatus_RISK_LEVEL_UNSPECIFIED
	switch in.RiskLevel {
	case types.RiskLevelSecure:
		riskLevel = infrapb.InstanceSecurityStatus_SECURE
	case types.RiskLevelLow:
		riskLevel = infrapb.InstanceSecurityStatus_LOW_RISK
	case types.RiskLevelMedium:
		riskLevel = infrapb.InstanceSecurityStatus_MEDIUM_RISK
	case types.RiskLevelHigh:
		riskLevel = infrapb.InstanceSecurityStatus_HIGH_RISK
	case types.RiskLevelCritical:
		riskLevel = infrapb.InstanceSecurityStatus_CRITICAL_RISK
	}

	var lastAssessedPb *timestamppb.Timestamp
	// TODO: Parse the timestamp string if needed
	// if in.LastSecurityScan != "" {
	//     lastAssessedPb = ... (implement timestamp parsing if needed)
	// }

	return &infrapb.InstanceSecurityStatus{
		OverallRiskLevel:      riskLevel,
		IsPubliclyAccessible:  in.IsPubliclyAccessible,
		HasOpenSshAccess:      in.HasOpenSSHAccess,
		HasOpenRdpAccess:      in.HasOpenRDPAccess,
		HasOpenDatabasePorts:  in.HasOpenDatabasePorts,
		IsUntagged:            in.IsUntagged,
		HasOverlyPermissiveSg: in.HasOverlyPermissiveSg,
		RiskSummary:           in.RiskSummary,
		ExposedPorts:          in.ExposedPorts,
		RiskySecurityGroups:   in.RiskySecurityGroups,
		Recommendations:       in.Recommendations,
		LastAssessed:          lastAssessedPb,
	}
}

func typesSecurityGroupRiskStatusToGrpc(in *types.SecurityGroupRiskStatus) *infrapb.SecurityGroupRiskStatus {
	if in == nil {
		return nil
	}

	riskLevel := infrapb.SecurityGroupRiskStatus_RISK_LEVEL_UNSPECIFIED
	switch in.RiskLevel {
	case types.RiskLevelSecure:
		riskLevel = infrapb.SecurityGroupRiskStatus_SECURE
	case types.RiskLevelLow:
		riskLevel = infrapb.SecurityGroupRiskStatus_LOW_RISK
	case types.RiskLevelMedium:
		riskLevel = infrapb.SecurityGroupRiskStatus_MEDIUM_RISK
	case types.RiskLevelHigh:
		riskLevel = infrapb.SecurityGroupRiskStatus_HIGH_RISK
	case types.RiskLevelCritical:
		riskLevel = infrapb.SecurityGroupRiskStatus_CRITICAL_RISK
	}

	var lastAssessedPb *timestamppb.Timestamp
	// TODO: Parse the timestamp string if needed
	// if in.LastSecurityScan != "" {
	//     lastAssessedPb = ... (implement timestamp parsing if needed)
	// }

	// Convert affected instances count to string array
	affectedInstancesStr := []string{}
	if in.AffectedInstances > 0 {
		// For now, just add a placeholder - in real implementation, this would be actual instance IDs
		affectedInstancesStr = append(affectedInstancesStr, fmt.Sprintf("%d instances", in.AffectedInstances))
	}

	return &infrapb.SecurityGroupRiskStatus{
		OverallRiskLevel:           riskLevel,
		AllowsSshFromInternet:      in.AllowsSSHFromInternet,
		AllowsRdpFromInternet:      in.AllowsRDPFromInternet,
		AllowsDatabaseFromInternet: in.AllowsHTTPFromInternet, // Map HTTP to database for now
		HasWidePortRanges:          in.HasOverlyPermissiveRules,
		IsUnused:                   false, // Would need to determine from actual usage
		IsUntagged:                 in.IsUntagged,
		RiskyRules:                 in.RiskyRules,
		AffectedInstances:          affectedInstancesStr,
		AttachedInstanceCount:      in.AffectedInstances,
		Recommendations:            in.Recommendations,
		LastAssessed:               lastAssessedPb,
	}
}

func typesLoadBalancerSecurityStatusToGrpc(in *types.LoadBalancerSecurityStatus) *infrapb.LoadBalancerSecurityStatus {
	if in == nil {
		return nil
	}

	riskLevel := infrapb.LoadBalancerSecurityStatus_RISK_LEVEL_UNSPECIFIED
	switch in.RiskLevel {
	case types.RiskLevelSecure:
		riskLevel = infrapb.LoadBalancerSecurityStatus_SECURE
	case types.RiskLevelLow:
		riskLevel = infrapb.LoadBalancerSecurityStatus_LOW_RISK
	case types.RiskLevelMedium:
		riskLevel = infrapb.LoadBalancerSecurityStatus_MEDIUM_RISK
	case types.RiskLevelHigh:
		riskLevel = infrapb.LoadBalancerSecurityStatus_HIGH_RISK
	case types.RiskLevelCritical:
		riskLevel = infrapb.LoadBalancerSecurityStatus_CRITICAL_RISK
	}

	var lastAssessedPb *timestamppb.Timestamp
	// TODO: Parse the timestamp string if needed
	// if in.LastSecurityScan != "" {
	//     lastAssessedPb = ... (implement timestamp parsing if needed)
	// }

	// Calculate risky backend instance count
	riskyBackendCount := int32(len(in.BackendInstanceRisks))

	return &infrapb.LoadBalancerSecurityStatus{
		OverallRiskLevel:         riskLevel,
		IsInternetFacing:         in.IsInternetFacing,
		HasRiskyBackendInstances: len(in.BackendInstanceRisks) > 0,
		HasInsecureListeners:     in.HasInsecureListeners,
		IsUntagged:               in.IsUntagged,
		HasOverlyPermissiveSg:    in.HasOpenSecurityGroups,
		TotalBackendInstances:    0, // Would need to be calculated from actual backend data
		RiskyBackendInstances:    riskyBackendCount,
		RiskyBackendInstanceIds:  in.BackendInstanceRisks,
		InsecureListeners:        in.InsecureProtocols,
		SecurityGroupIssues:      in.SecurityGroupRisks,
		Recommendations:          in.Recommendations,
		LastAssessed:             lastAssessedPb,
	}
}
