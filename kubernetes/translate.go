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

package kubernetes

import (
	"strconv"

	"github.com/app-net-interface/awi-infra-guard/types"
	v1 "k8s.io/api/core/v1"
)

func k8sServicesToTypes(cluster string, services []v1.Service) []types.K8SService {
	k8sServices := make([]types.K8SService, 0, len(services))
	for _, svc := range services {
		var ingresses []types.K8sServiceIngress
		for _, ing := range svc.Status.LoadBalancer.Ingress {
			var ports []string
			for _, portStatus := range ing.Ports {
				ports = append(ports, string(portStatus.Port))
			}
			ingresses = append(ingresses, types.K8sServiceIngress{
				Hostname: ing.Hostname,
				IP:       ing.IP,
				Ports:    ports,
			})
		}
		protocolsAndPorts := make(types.ProtocolsAndPorts, 0)
		for _, port := range svc.Spec.Ports {
			var portNr string
			if svc.Spec.Type == v1.ServiceTypeNodePort {
				portNr = strconv.Itoa(int(port.NodePort))
			} else {
				portNr = strconv.Itoa(int(port.Port))
			}
			portsForProto, ok := protocolsAndPorts[string(port.Protocol)]
			if ok && len(portsForProto) != 0 {
				portsForProto = append(portsForProto, portNr)
			} else {
				protocolsAndPorts[string(port.Protocol)] = []string{portNr}
			}
		}
		k8sServices = append(k8sServices, types.K8SService{
			Type:              string(svc.Spec.Type),
			Cluster:           cluster,
			Namespace:         svc.Namespace,
			Name:              svc.Name,
			Ingresses:         ingresses,
			Labels:            svc.Labels,
			ProtocolsAndPorts: protocolsAndPorts,
		})
	}
	return k8sServices
}

func k8sNodesToTypes(cluster string, nodes []v1.Node) []types.K8sNode {
	k8sNodes := make([]types.K8sNode, 0, len(nodes))
	for _, node := range nodes {
		k8sNodes = append(k8sNodes, types.K8sNode{
			Cluster:   cluster,
			Namespace: node.Namespace,
			Name:      node.Name,
			Addresses: node.Status.Addresses,
			Labels:    node.Labels,
		})
	}
	return k8sNodes
}

func k8sNamespacesToTypes(cluster string, namespaces []v1.Namespace) []types.Namespace {
	k8sNamespaces := make([]types.Namespace, 0, len(namespaces))
	for _, namespace := range namespaces {
		k8sNamespaces = append(k8sNamespaces, types.Namespace{
			Cluster: cluster,
			Name:    namespace.Name,
			Labels:  namespace.Labels,
		})
	}
	return k8sNamespaces
}

func k8sPodsToTypes(clusterName string, k8sPods []v1.Pod) []types.Pod {
	pods := make([]types.Pod, 0, len(k8sPods))
	for _, pod := range k8sPods {
		pods = append(pods, types.Pod{
			Cluster:   clusterName,
			Namespace: pod.GetNamespace(),
			Name:      pod.GetName(),
			Ip:        pod.Status.PodIP,
			Labels:    pod.GetLabels(),
			State:     string(pod.Status.Phase),
		})
	}
	return pods
}
