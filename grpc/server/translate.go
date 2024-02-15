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
)

func grpcProtocolsAndPortToTypes(in map[string]*infrapb.Ports) types.ProtocolsAndPorts {
	out := make(types.ProtocolsAndPorts, len(in))
	for k, v := range in {
		ports := make([]string, 0, len(v.GetPorts()))
		for _, port := range v.GetPorts() {
			ports = append(ports, port)
		}
		out[k] = ports
	}
	return out
}

func typesInstanceToGrpc(in []types.Instance) []*infrapb.Instance {
	out := make([]*infrapb.Instance, 0, len(in))
	for _, instance := range in {
		out = append(out, &infrapb.Instance{
			Id:           instance.ID,
			Name:         instance.Name,
			PublicIP:     instance.PublicIP,
			PrivateIP:    instance.PrivateIP,
			SubnetID:     instance.SubnetID,
			VpcId:        instance.VPCID,
			Zone:         instance.Zone,
			Region:       instance.Region,
			Labels:       instance.Labels,
			Provider:     instance.Provider,
			AccountId:    instance.AccountID,
			State:        instance.State,
			LastSyncTime: instance.LastSyncTime,
		})
	}
	return out
}

func typesSubnetsToGrpc(in []types.Subnet) []*infrapb.Subnet {
	out := make([]*infrapb.Subnet, 0, len(in))
	for _, subnet := range in {
		out = append(out, &infrapb.Subnet{
			SubnetId:     subnet.SubnetId,
			Name:         subnet.Name,
			CidrBlock:    subnet.CidrBlock,
			VpcId:        subnet.VpcId,
			Zone:         subnet.Zone,
			Region:       subnet.Region,
			Labels:       subnet.Labels,
			Provider:     subnet.Provider,
			AccountId:    subnet.AccountID,
			LastSyncTime: subnet.LastSyncTime,
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
			Provider:     vpc.Provider,
			AccountId:    vpc.AccountID,
			LastSyncTime: vpc.LastSyncTime,
		})
	}
	return out
}

func typesPodsToGrpc(in []types.Pod) []*infrapb.Pod {
	out := make([]*infrapb.Pod, 0, len(in))
	for _, pod := range in {
		out = append(out, &infrapb.Pod{
			Cluster:      pod.Cluster,
			Namespace:    pod.Namespace,
			Name:         pod.Name,
			Ip:           pod.Ip,
			Labels:       pod.Labels,
			State:        pod.State,
			LastSyncTime: pod.LastSyncTime,
		})
	}
	return out
}

func typesServicesToGrpc(in []types.K8SService) []*infrapb.K8SService {
	out := make([]*infrapb.K8SService, 0, len(in))
	for _, svc := range in {
		out = append(out, &infrapb.K8SService{
			Cluster:      svc.Cluster,
			Namespace:    svc.Namespace,
			Name:         svc.Name,
			Ingresses:    typesIngressesToGrpc(svc.Ingresses),
			Labels:       svc.Labels,
			Type:         svc.Type,
			LastSyncTime: svc.LastSyncTime,
		})
	}
	return out
}

func typesIngressesToGrpc(in []types.K8sServiceIngress) []*infrapb.K8SService_Ingress {
	out := make([]*infrapb.K8SService_Ingress, 0, len(in))
	for _, ing := range in {
		out = append(out, &infrapb.K8SService_Ingress{
			Hostname: ing.Hostname,
			IP:       ing.IP,
			Ports:    ing.Ports,
		})
	}
	return out
}

func typesClustersToGrpc(in []types.Cluster) []*infrapb.Cluster {
	out := make([]*infrapb.Cluster, 0, len(in))
	for _, cluster := range in {
		out = append(out, &infrapb.Cluster{
			Name:         cluster.Name,
			FullName:     cluster.FullName,
			Arn:          cluster.Arn,
			VpcId:        cluster.VpcID,
			Region:       cluster.Region,
			Project:      cluster.Project,
			Labels:       cluster.Labels,
			Provider:     cluster.Provider,
			AccountId:    cluster.AccountID,
			Id:           cluster.Id,
			LastSyncTime: cluster.LastSyncTime,
		})
	}
	return out
}

func typesNamespacesToGrpc(in []types.Namespace) []*infrapb.Namespace {
	out := make([]*infrapb.Namespace, 0, len(in))
	for _, namespace := range in {
		out = append(out, &infrapb.Namespace{
			Cluster:      namespace.Cluster,
			Name:         namespace.Name,
			Labels:       namespace.Labels,
			LastSyncTime: namespace.LastSyncTime,
		})
	}
	return out
}

func typesNodesToGrpc(in []types.K8sNode) []*infrapb.Node {
	out := make([]*infrapb.Node, 0, len(in))
	for _, node := range in {
		addresses := make([]string, 0, len(node.Addresses))
		for _, address := range node.Addresses {
			addresses = append(addresses, fmt.Sprintf("%s:%s", address.Type, address.Address))
		}
		out = append(out, &infrapb.Node{
			Cluster:      node.Cluster,
			Name:         node.Name,
			Namespace:    node.Namespace,
			Addresses:    addresses,
			LastSyncTime: node.LastSyncTime,
		})
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
		out = append(out, &infrapb.RouteTable{
			Provider:     rt.Provider,
			Id:           rt.ID,
			Name:         rt.Name,
			VpcId:        rt.VpcID,
			Region:       rt.Region,
			AccountId:    rt.AccountID,
			Labels:       rt.Labels,
			Routes:       routes,
			LastSyncTime: rt.LastSyncTime,
		})
	}
	return out
}
