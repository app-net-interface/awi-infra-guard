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

package types

import (
	v1 "k8s.io/api/core/v1"
)

/* Start Kubernetes types */

//Kubernetes

type Cluster struct {
	Name         string
	FullName     string
	Arn          string
	VpcID        string
	Region       string
	Project      string
	Labels       map[string]string
	Provider     string
	AccountID    string
	Id           string
	SelfLink     string
	LastSyncTime string
}

func (v *Cluster) DbId() string {
	return CloudID(v.Provider, v.Name)
}

func (v *Cluster) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Cluster) GetProvider() string {
	return v.Provider
}

type Pod struct {
	Cluster      string
	Namespace    string
	Name         string
	Ip           string
	Labels       map[string]string
	State        string
	SelfLink     string
	LastSyncTime string
}

func (v *Pod) DbId() string {
	return KubernetesID(v.Cluster, v.Namespace, v.Name)
}

func (v *Pod) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Pod) GetProvider() string {
	return v.Cluster
}

type K8SService struct {
	Cluster           string
	Namespace         string
	Name              string
	Type              string
	ProtocolsAndPorts ProtocolsAndPorts
	Ingresses         []K8sServiceIngress
	Labels            map[string]string
	SelfLink          string

	LastSyncTime string
}

func (v *K8SService) DbId() string {
	return KubernetesID(v.Cluster, v.Namespace, v.Name)
}

func (v *K8SService) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *K8SService) GetProvider() string {
	return v.Cluster
}

type K8sServiceIngress struct {
	Hostname string
	IP       string
	Ports    []string
}

type K8sNode struct {
	Cluster      string
	Name         string
	Namespace    string
	Addresses    []v1.NodeAddress
	Labels       map[string]string
	SelfLink     string
	LastSyncTime string
}

func (v *K8sNode) DbId() string {
	return KubernetesID(v.Cluster, v.Namespace, v.Name)
}

func (v *K8sNode) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *K8sNode) GetProvider() string {
	return v.Cluster
}

type Namespace struct {
	Cluster      string
	Name         string
	Labels       map[string]string
	SelfLink     string
	LastSyncTime string
}

func (v *Namespace) DbId() string {
	return KubernetesID(v.Cluster, v.Name, "")
}

func (v *Namespace) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Namespace) GetProvider() string {
	return v.Cluster
}

func KubernetesID(cluster, namespace, name string) string {
	n := cluster + "/" + namespace
	if name != "" {
		n += "/" + name
	}
	return n
}

/* End Kubernetes types */
