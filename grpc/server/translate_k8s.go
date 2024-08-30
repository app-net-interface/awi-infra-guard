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

package server

import (
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// Kubernetes Types

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
			SelfLink:     pod.SelfLink,
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
			SelfLink:     svc.SelfLink,
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
			SelfLink:     cluster.SelfLink,
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
			SelfLink:     node.SelfLink,
		})
	}
	return out
}

