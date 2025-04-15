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
		})
	}
	return out
}

func typesSubnetsToGrpc(in []types.Subnet) []*infrapb.Subnet {
	out := make([]*infrapb.Subnet, 0, len(in))
	for _, subnet := range in {
		out = append(out, &infrapb.Subnet{
			Id:           subnet.SubnetId,
			Name:         subnet.Name,
			CidrBlock:    subnet.CidrBlock,
			VpcId:        subnet.VpcId,
			Zone:         subnet.Zone,
			Region:       subnet.Region,
			Labels:       subnet.Labels,
			Provider:     subnet.Provider,
			AccountId:    subnet.AccountID,
			LastSyncTime: subnet.LastSyncTime,
			SelfLink:     subnet.SelfLink,
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
		out = append(out, &infrapb.VPCEndpoint{
			Id:            vpce.ID,
			Name:          vpce.Name,
			VpcId:         vpce.VPCId,
			Provider:      vpce.Provider,
			Region:        vpce.Region,
			State:         vpce.State,
			Labels:        vpce.Labels,
			AccountId:     vpce.AccountID,
			RouteTableIds: vpce.RouteTableIds,
			SubnetIds:     vpce.SubnetIds,
			ServiceName:   vpce.ServiceName,
			Type:          vpce.Type,
			CreatedAt:     timestamppb.New(*vpce.CreatedAt),
			LastSyncTime:  vpce.LastSyncTime,
			SelfLink:      vpce.SelfLink,
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
		subnetIds := make([]string, 0, len(rt.SubnetIds))
		subnetIds = append(subnetIds, rt.SubnetIds...)
		gatewayIds := make([]string, 0, len(rt.GatewayIds))
		gatewayIds = append(gatewayIds, rt.GatewayIds...)

		out = append(out, &infrapb.RouteTable{
			Provider:     rt.Provider,
			Id:           rt.ID,
			Name:         rt.Name,
			VpcId:        rt.VpcID,
			Region:       rt.Region,
			AccountId:    rt.AccountID,
			Labels:       rt.Labels,
			Routes:       routes,
			SubnetIds:    subnetIds,
			GatewayIds:   gatewayIds,
			LastSyncTime: rt.LastSyncTime,
			SelfLink:     rt.SelfLink,
		})
	}
	return out
}

func typesPublicIPsToGrpc(in []types.PublicIP) []*infrapb.PublicIP {
	out := make([]*infrapb.PublicIP, 0, len(in))
	for _, publicIP := range in {

		out = append(out, &infrapb.PublicIP{
			Provider:   publicIP.Provider,
			Id:         publicIP.ID,
			VpcId:      publicIP.VPCID,
			Region:     publicIP.Region,
			PublicIp:   publicIP.PublicIP,
			InstanceId: publicIP.InstanceId,
			PrivateIp:  publicIP.PrivateIP,
			AccountId:  publicIP.AccountID,
			Type:       publicIP.Type,
			Labels:     publicIP.Labels,
			SelfLink:   publicIP.SelfLink,
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
			TargetGroupIds:         lb.TargetGroupIDs,
			CrossZoneLoadBalancing: lb.CrossZoneLoadBalancing,
			AccessLogsEnabled:      lb.AccessLogsEnabled,
			LoggingBucket:          lb.LoggingBucket,
			IpAddresses:            lb.IPAddresses,
			IpAddressType:          lb.IPAddressType,
			Zone:                   lb.Zone,
			Labels:                 lb.Labels,
			Project:                lb.Project,
			CreatedAt:              timestamppb.New(lb.CreatedAt),
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
			Status: ni.Status,
			//InterfaceType:      ni.InterfaceType,
			LastSyncTime: ni.LastSyncTime,
			SubnetId:     ni.SubnetID,
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
	}
}

func typesVpcGraphNodesToGrpc(in []types.VpcGraphNode) []*infrapb.VpcGraphNode {
	out := make([]*infrapb.VpcGraphNode, 0, len(in))
	for _, node := range in {
		if node.Properties == nil {
			// Ensure properties map is not nil for gRPC message
			node.Properties = make(map[string]string)
		}
		out = append(out, &infrapb.VpcGraphNode{
			Id:           node.ID,
			ResourceType: node.ResourceType,
			Name:         node.Name,
			Properties:   node.Properties,
			Provider:     node.Provider,
			AccountId:    node.AccountId,
		})
	}
	return out
}

func typesVpcGraphEdgesToGrpc(in []types.VpcGraphEdge) []*infrapb.VpcGraphEdge {
	out := make([]*infrapb.VpcGraphEdge, 0, len(in))
	for _, edge := range in {
		out = append(out, &infrapb.VpcGraphEdge{
			SourceNodeId:     edge.SourceNodeID,
			TargetNodeId:     edge.TargetNodeID,
			RelationshipType: edge.RelationshipType,
			Provider:         edge.Provider,
			AccountId:        edge.AccountId,
		})
	}
	return out
}
