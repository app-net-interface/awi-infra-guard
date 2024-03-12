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

package types

import (
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
)

const (
	ConditionLabel = "condition"
	OrCondition    = "OR"
	AndCondition   = "AND"
	Separator      = ":"
)

const (
	VPCType           = "VPC"
	InstanceType      = "Instance"
	SubnetType        = "Subnet"
	ACLType           = "ACL"
	SecurityGroupType = "SecurityGroup"
	RouteTableType    = "RouteTable"
	ClusterType       = "Cluster"
	PodsType          = "Pod"
	K8sServiceType    = "K8sService"
	K8sNodeType       = "K8sNode"
	NamespaceType     = "Namespace"
)

type VPC struct {
	ID           string
	Name         string
	Region       string
	Labels       map[string]string
	IPv4CIDR     string
	IPv6CIDR     string
	Provider     string
	AccountID    string
	LastSyncTime string
}

func (v *VPC) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *VPC) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPC) GetProvider() string {
	return v.Provider
}

type Instance struct {
	ID           string
	Name         string
	PublicIP     string
	PrivateIP    string
	SubnetID     string
	VPCID        string
	Labels       map[string]string
	State        string
	Region       string
	Zone         string
	Provider     string
	AccountID    string
	LastSyncTime string
}

func (v Instance) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *Instance) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Instance) GetProvider() string {
	return v.Provider
}

type Subnet struct {
	Name         string
	SubnetId     string
	CidrBlock    string
	VpcId        string
	Zone         string
	Labels       map[string]string
	Region       string
	Provider     string
	AccountID    string
	LastSyncTime string
}

func (v *Subnet) DbId() string {
	return CloudID(v.Provider, v.SubnetId)
}

func (v *Subnet) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Subnet) GetProvider() string {
	return v.Provider
}

type Ports []string

type ProtocolsAndPorts map[string]Ports

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
	LastSyncTime      string
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

type Account struct {
	Name         string
	ID           string
	Provider     string
	LastSyncTime string
}

type RouteTable struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Routes       []Route
	LastSyncTime string
}

func (v *RouteTable) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *RouteTable) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *RouteTable) GetProvider() string {
	return v.Provider
}

type Route struct {
	Name        string
	Destination string
	Status      string
	Target      string
	NextHopType string
	NextHopIP   string
}

type SecurityGroup struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Rules        []SecurityGroupRule
	LastSyncTime string
}

type SecurityGroupRule struct {
	Protocol  string
	PortRange string
	Source    []string
	Direction string
}

func (v *SecurityGroup) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *SecurityGroup) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *SecurityGroup) GetProvider() string {
	return v.Provider
}

type ACL struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Rules        []ACLRule
	LastSyncTime string
}

type ACLRule struct {
	Number            int
	Protocol          string
	PortRange         string
	SourceRanges      []string
	DestinationRanges []string
	Action            string
	Direction         string
}

func (v *ACL) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *ACL) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *ACL) GetProvider() string {
	return v.Provider
}

type SyncTime struct {
	Provider     string
	ResourceType string
	Time         string
}

func (v *SyncTime) DbId() string {
	return SyncTimeKey(v.Provider, v.ResourceType)
}

func (v *SyncTime) SetSyncTime(time string) {

}

func (v *SyncTime) GetProvider() string {
	return v.Provider
}

func KubernetesID(cluster, namespace, name string) string {
	n := cluster + "/" + namespace
	if name != "" {
		n += "/" + name
	}
	return n
}

func CloudID(provider, id string) string {
	return provider + ":" + id
}

func SyncTimeKey(provider string, typ string) string {
	return fmt.Sprintf("%s/%s", provider, typ)
}

func SyncTimeKeyDecode(s string) (provider, type_ string, err error) {
	split := strings.Split(s, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("Failed to determine provider and type from key %s", s)
	}
	return split[0], split[1], nil
}

type DestinationDetails struct {
	Provider string
	VPC      string
	Region   string
}

type SingleVPCConnectionParams struct {
	ConnID      string
	VpcID       string
	Region      string
	Destination DestinationDetails
}

type VPCConnectionParams struct {
	ConnID  string
	Vpc1ID  string
	Vpc2ID  string
	Region1 string
	Region2 string
}

type VPCConnectionOutput struct {
	Region1 string
	Region2 string
}

type SingleVPCConnectionOutput struct {
	Region string
}

type VPCDisconnectionParams struct {
	ConnID  string
	Vpc1ID  string
	Vpc2ID  string
	Region1 string
	Region2 string
}

type SingleVPCDisconnectionParams struct {
	ConnID string
	VpcID  string
	Region string
}

type VPCDisconnectionOutput struct {
}
