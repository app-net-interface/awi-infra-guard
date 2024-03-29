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

import (
	"time"

	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/boltdb/bolt"
)

type boltClient struct {
	db *bolt.DB
}

func NewBoltClient() Client {
	return &boltClient{}
}

func (client *boltClient) Open(filename string) error {
	options := &bolt.Options{Timeout: time.Second}
	var err error
	client.db, err = bolt.Open(filename, 0600, options)
	if err != nil {
		return err
	}

	return client.db.Update(func(tx *bolt.Tx) error {
		for _, tableName := range tableNames {
			_, err := tx.CreateBucketIfNotExists([]byte(tableName))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (client *boltClient) Close() error {
	return client.db.Close()
}

func (client *boltClient) DropDB() error {
	return client.db.Update(func(tx *bolt.Tx) error {
		for _, tableName := range tableNames {
			if err := tx.DeleteBucket([]byte(tableName)); err != nil {
				return err
			}
			_, err := tx.CreateBucket([]byte(tableName))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (client *boltClient) PutVPC(vpc *types.VPC) error {
	return update(client, vpc, vpc.DbId(), vpcTable)
}

func (client *boltClient) GetVPC(id string) (*types.VPC, error) {
	return get[types.VPC](client, id, vpcTable)
}

func (client *boltClient) ListVPCs() ([]*types.VPC, error) {
	return list[types.VPC](client, vpcTable)
}

func (client *boltClient) DeleteVPC(id string) error {
	return delete_(client, id, vpcTable)
}

// Instance
func (client *boltClient) PutInstance(instance *types.Instance) error {
	return update(client, instance, instance.DbId(), instanceTable)
}

func (client *boltClient) GetInstance(id string) (*types.Instance, error) {
	return get[types.Instance](client, id, instanceTable)
}

func (client *boltClient) ListInstances() ([]*types.Instance, error) {
	return list[types.Instance](client, instanceTable)
}

func (client *boltClient) DeleteInstance(id string) error {
	return delete_(client, id, instanceTable)
}

// Subnet
func (client *boltClient) PutSubnet(subnet *types.Subnet) error {
	return update(client, subnet, subnet.DbId(), subnetTable)
}

func (client *boltClient) GetSubnet(id string) (*types.Subnet, error) {
	return get[types.Subnet](client, id, subnetTable)
}

func (client *boltClient) ListSubnets() ([]*types.Subnet, error) {
	return list[types.Subnet](client, subnetTable)
}

func (client *boltClient) DeleteSubnet(id string) error {
	return delete_(client, id, subnetTable)
}

// RouteTable
func (client *boltClient) PutRouteTable(routeTable *types.RouteTable) error {
	return update(client, routeTable, routeTable.DbId(), routeTableTable)
}

func (client *boltClient) GetRouteTable(id string) (*types.RouteTable, error) {
	return get[types.RouteTable](client, id, routeTableTable)
}

func (client *boltClient) ListRouteTables() ([]*types.RouteTable, error) {
	return list[types.RouteTable](client, routeTableTable)
}

func (client *boltClient) DeleteRouteTable(id string) error {
	return delete_(client, id, routeTableTable)
}

// NAT Gateway
func (client *boltClient) PutNATGateway(ng *types.NATGateway) error {
	return update(client, ng, ng.DbId(), ngTable)
}

func (client *boltClient) GetNATGateway(id string) (*types.NATGateway, error) {
	return get[types.NATGateway](client, id, ngTable)
}

func (client *boltClient) ListNATGateways() ([]*types.NATGateway, error) {
	return list[types.NATGateway](client, ngTable)
}

func (client *boltClient) DeleteNATGateway(id string) error {
	return delete_(client, id, ngTable)
}

// Router
func (client *boltClient) PutRouter(router *types.Router) error {
	return update(client, router, router.DbId(), routerTable)
}

func (client *boltClient) GetRouter(id string) (*types.Router, error) {
	return get[types.Router](client, id, routerTable)
}

func (client *boltClient) ListRouters() ([]*types.Router, error) {
	return list[types.Router](client, routerTable)
}

func (client *boltClient) DeleteRouter(id string) error {
	return delete_(client, id, routerTable)
}

// SecurityGroup
func (client *boltClient) PutSecurityGroup(securityGroup *types.SecurityGroup) error {
	return update(client, securityGroup, securityGroup.DbId(), securityGroupTable)
}

func (client *boltClient) GetSecurityGroup(id string) (*types.SecurityGroup, error) {
	return get[types.SecurityGroup](client, id, securityGroupTable)
}

func (client *boltClient) ListSecurityGroups() ([]*types.SecurityGroup, error) {
	return list[types.SecurityGroup](client, securityGroupTable)
}

func (client *boltClient) DeleteSecurityGroup(id string) error {
	return delete_(client, id, securityGroupTable)
}

// ACL
func (client *boltClient) PutACL(acl *types.ACL) error {
	return update(client, acl, acl.DbId(), aclTable)
}

func (client *boltClient) GetACL(id string) (*types.ACL, error) {
	return get[types.ACL](client, id, aclTable)
}

func (client *boltClient) ListACLs() ([]*types.ACL, error) {
	return list[types.ACL](client, aclTable)
}

func (client *boltClient) DeleteACL(id string) error {
	return delete_(client, id, aclTable)
}

// Cluster
func (client *boltClient) PutCluster(cluster *types.Cluster) error {
	return update(client, cluster, cluster.DbId(), clusterTable)
}

func (client *boltClient) GetCluster(id string) (*types.Cluster, error) {
	return get[types.Cluster](client, id, clusterTable)
}

func (client *boltClient) ListClusters() ([]*types.Cluster, error) {
	return list[types.Cluster](client, clusterTable)
}

func (client *boltClient) DeleteCluster(id string) error {
	return delete_(client, id, clusterTable)
}

// Pod
func (client *boltClient) PutPod(pod *types.Pod) error {
	return update(client, pod, pod.DbId(), podTable)
}

func (client *boltClient) GetPod(id string) (*types.Pod, error) {
	return get[types.Pod](client, id, podTable)
}

func (client *boltClient) ListPods() ([]*types.Pod, error) {
	return list[types.Pod](client, podTable)
}

func (client *boltClient) DeletePod(id string) error {
	return delete_(client, id, podTable)
}

// KubernetesService
func (client *boltClient) PutKubernetesService(service *types.K8SService) error {
	return update(client, service, service.DbId(), kubernetesServiceTable)
}

func (client *boltClient) GetKubernetesService(id string) (*types.K8SService, error) {
	return get[types.K8SService](client, id, kubernetesServiceTable)
}

func (client *boltClient) ListKubernetesServices() ([]*types.K8SService, error) {
	return list[types.K8SService](client, kubernetesServiceTable)
}

func (client *boltClient) DeleteKubernetesService(id string) error {
	return delete_(client, id, kubernetesServiceTable)
}

// KubernetesNode
func (client *boltClient) PutKubernetesNode(node *types.K8sNode) error {
	return update(client, node, node.DbId(), kubernetesNodeTable)
}

func (client *boltClient) GetKubernetesNode(id string) (*types.K8sNode, error) {
	return get[types.K8sNode](client, id, kubernetesNodeTable)
}

func (client *boltClient) ListKubernetesNodes() ([]*types.K8sNode, error) {
	return list[types.K8sNode](client, kubernetesNodeTable)
}

func (client *boltClient) DeleteKubernetesNode(id string) error {
	return delete_(client, id, kubernetesNodeTable)
}

// Namespace
func (client *boltClient) PutNamespace(namespace *types.Namespace) error {
	return update(client, namespace, namespace.DbId(), namespaceTable)
}

func (client *boltClient) GetNamespace(id string) (*types.Namespace, error) {
	return get[types.Namespace](client, id, namespaceTable)
}

func (client *boltClient) ListNamespaces() ([]*types.Namespace, error) {
	return list[types.Namespace](client, namespaceTable)
}

func (client *boltClient) DeleteNamespace(id string) error {
	return delete_(client, id, namespaceTable)
}

// SyncTime
func (client *boltClient) PutSyncTime(syncTime *types.SyncTime) error {
	return update(client, syncTime, syncTime.DbId(), syncTimeTable)
}
func (client *boltClient) GetSyncTime(id string) (*types.SyncTime, error) {
	return get[types.SyncTime](client, id, syncTimeTable)
}
func (client *boltClient) ListSyncTimes() ([]*types.SyncTime, error) {
	return list[types.SyncTime](client, syncTimeTable)
}
func (client *boltClient) DeleteSyncTime(id string) error {
	return delete_(client, id, syncTimeTable)
}
