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

package db

import "github.com/app-net-interface/awi-infra-guard/types"

const (
	vpcTable               = "vpcs"
	regionTable            = "regions"
	instanceTable          = "instances"
	subnetTable            = "subnets"
	clusterTable           = "clusters"
	podTable               = "pods"
	kubernetesServiceTable = "kubernetes_services"
	kubernetesNodeTable    = "kubernetes_nodes"
	namespaceTable         = "namespaces"
	accountTable           = "accounts"
	routeTableTable        = "route_tables"
	aclTable               = "acls"
	securityGroupTable     = "security_groups"
	ngTable                = "nat_gateways"
	routerTable            = "routers"
	igwTable               = "igws"
	vpcEndpointTable       = "vpcEndpoints"
	syncTimeTable          = "sync_time"
)

// Add bolt db table to this list; or it will cause a panic
var tableNames = []string{
	vpcTable,
	regionTable,
	instanceTable,
	subnetTable,
	clusterTable,
	podTable,
	kubernetesServiceTable,
	kubernetesNodeTable,
	namespaceTable,
	accountTable,
	routeTableTable,
	ngTable,
	routerTable,
	igwTable,
	vpcEndpointTable,
	aclTable,
	securityGroupTable,
	syncTimeTable,
}

type DbObject interface {
	DbId() string
	GetProvider() string
	SetSyncTime(string)
}

type Client interface {
	Open(filename string) error
	Close() error
	DropDB() error

	// VPC
	PutRegion(region *types.Region) error
	GetRegion(id string) (*types.Region, error)
	ListRegions() ([]*types.Region, error)
	DeleteRegion(id string) error

	// VPC
	PutVPC(vpc *types.VPC) error
	GetVPC(id string) (*types.VPC, error)
	ListVPCs() ([]*types.VPC, error)
	DeleteVPC(id string) error

	// Instance
	PutInstance(instance *types.Instance) error
	GetInstance(id string) (*types.Instance, error)
	ListInstances() ([]*types.Instance, error)
	DeleteInstance(id string) error

	//Subnet
	PutSubnet(subnet *types.Subnet) error
	GetSubnet(id string) (*types.Subnet, error)
	ListSubnets() ([]*types.Subnet, error)
	DeleteSubnet(id string) error

	//ACL
	PutACL(acl *types.ACL) error
	GetACL(id string) (*types.ACL, error)
	ListACLs() ([]*types.ACL, error)
	DeleteACL(id string) error

	// RouteTable
	PutRouteTable(routeTable *types.RouteTable) error
	GetRouteTable(id string) (*types.RouteTable, error)
	ListRouteTables() ([]*types.RouteTable, error)
	DeleteRouteTable(id string) error

	// NAT Gatweay
	ListNATGateways() ([]*types.NATGateway, error)
	PutNATGateway(ng *types.NATGateway) error
	GetNATGateway(id string) (*types.NATGateway, error)
	DeleteNATGateway(id string) error

	// Router
	ListRouters() ([]*types.Router, error)
	PutRouter(ng *types.Router) error
	GetRouter(id string) (*types.Router, error)
	DeleteRouter(id string) error

	// IGW
	ListInternetGateways() ([]*types.IGW, error)
	PutIGW(ng *types.IGW) error
	GetIGW(id string) (*types.IGW, error)
	DeleteIGW(id string) error

	// VPCEndpoint
	ListVPCEndpoints() ([]*types.VPCEndpoint, error)
	PutVPCEndpoint(ng *types.VPCEndpoint) error
	GetVPCEndpoint(id string) (*types.VPCEndpoint, error)
	DeleteVPCEndpoint(id string) error

	//Security Group
	PutSecurityGroup(securityGroup *types.SecurityGroup) error
	GetSecurityGroup(id string) (*types.SecurityGroup, error)
	ListSecurityGroups() ([]*types.SecurityGroup, error)
	DeleteSecurityGroup(id string) error

	// Kubernets

	//  Cluster
	PutCluster(cluster *types.Cluster) error
	GetCluster(id string) (*types.Cluster, error)
	ListClusters() ([]*types.Cluster, error)
	DeleteCluster(id string) error
	// Pods
	PutPod(pod *types.Pod) error
	GetPod(id string) (*types.Pod, error)
	ListPods() ([]*types.Pod, error)
	DeletePod(id string) error
	//Services
	PutKubernetesService(service *types.K8SService) error
	GetKubernetesService(id string) (*types.K8SService, error)
	ListKubernetesServices() ([]*types.K8SService, error)
	DeleteKubernetesService(id string) error
	// Nodes
	PutKubernetesNode(node *types.K8sNode) error
	GetKubernetesNode(id string) (*types.K8sNode, error)
	ListKubernetesNodes() ([]*types.K8sNode, error)
	DeleteKubernetesNode(id string) error

	//Namespaces
	PutNamespace(namespace *types.Namespace) error
	GetNamespace(id string) (*types.Namespace, error)
	ListNamespaces() ([]*types.Namespace, error)
	DeleteNamespace(id string) error

	// Time
	PutSyncTime(time *types.SyncTime) error
	GetSyncTime(id string) (*types.SyncTime, error)
	ListSyncTimes() ([]*types.SyncTime, error)
	DeleteSyncTime(id string) error
}
