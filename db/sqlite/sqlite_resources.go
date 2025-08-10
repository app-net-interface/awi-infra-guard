// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

package sqlite

import (
	"github.com/app-net-interface/awi-infra-guard/types"
)

// RESOURCE CRUD OPERATIONS
// This file contains all the Put/Get/List/Delete operations for different resource types

// VPC methods
func (client *sqliteClient) PutVPC(vpc *types.VPC) error {
	return client.putObject(vpc, vpc.DbId(), vpcTable)
}

func (client *sqliteClient) GetVPC(id string) (*types.VPC, error) {
	vpc := &types.VPC{}
	err := client.getObject(id, vpcTable, vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

func (client *sqliteClient) ListVPCs() ([]*types.VPC, error) {
	objects, err := client.listObjects(vpcTable, func() interface{} { return &types.VPC{} })
	if err != nil {
		return nil, err
	}

	vpcs := make([]*types.VPC, len(objects))
	for i, obj := range objects {
		vpcs[i] = obj.(*types.VPC)
	}
	return vpcs, nil
}

func (client *sqliteClient) DeleteVPC(id string) error {
	return client.deleteObject(id, vpcTable)
}

// Region methods
func (client *sqliteClient) PutRegion(region *types.Region) error {
	return client.putObject(region, region.DbId(), regionTable)
}

func (client *sqliteClient) GetRegion(id string) (*types.Region, error) {
	region := &types.Region{}
	err := client.getObject(id, regionTable, region)
	return region, err
}

func (client *sqliteClient) ListRegions() ([]*types.Region, error) {
	objects, err := client.listObjects(regionTable, func() interface{} { return &types.Region{} })
	if err != nil {
		return nil, err
	}

	regions := make([]*types.Region, len(objects))
	for i, obj := range objects {
		regions[i] = obj.(*types.Region)
	}
	return regions, nil
}
func (client *sqliteClient) DeleteRegion(id string) error {
	_, err := client.db.Exec("DELETE FROM "+regionTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutInstance(instance *types.Instance) error {
	return client.putObject(instance, instance.DbId(), instanceTable)
}
func (client *sqliteClient) GetInstance(id string) (*types.Instance, error) {
	instance := &types.Instance{}
	err := client.getObject(id, instanceTable, instance)
	return instance, err
}
func (client *sqliteClient) ListInstances() ([]*types.Instance, error) {
	objects, err := client.listObjects(instanceTable, func() interface{} { return &types.Instance{} })
	if err != nil {
		return nil, err
	}
	var instances []*types.Instance
	for _, obj := range objects {
		instances = append(instances, obj.(*types.Instance))
	}
	return instances, nil
}
func (client *sqliteClient) DeleteInstance(id string) error {
	_, err := client.db.Exec("DELETE FROM "+instanceTable+" WHERE id = ?", id)
	return err
}

// Subnet methods
func (client *sqliteClient) PutSubnet(subnet *types.Subnet) error {
	return client.putObject(subnet, subnet.DbId(), subnetTable)
}

func (client *sqliteClient) GetSubnet(id string) (*types.Subnet, error) {
	subnet := &types.Subnet{}
	err := client.getObject(id, subnetTable, subnet)
	return subnet, err
}

func (client *sqliteClient) ListSubnets() ([]*types.Subnet, error) {
	objects, err := client.listObjects(subnetTable, func() interface{} { return &types.Subnet{} })
	if err != nil {
		return nil, err
	}

	var subnets []*types.Subnet
	for _, obj := range objects {
		subnets = append(subnets, obj.(*types.Subnet))
	}
	return subnets, nil
}

func (client *sqliteClient) DeleteSubnet(id string) error {
	return client.deleteObject(id, subnetTable)
}

// ACL methods
func (client *sqliteClient) PutACL(acl *types.ACL) error {
	return client.putObject(acl, acl.DbId(), aclTable)
}

func (client *sqliteClient) GetACL(id string) (*types.ACL, error) {
	acl := &types.ACL{}
	err := client.getObject(id, aclTable, acl)
	return acl, err
}

func (client *sqliteClient) ListACLs() ([]*types.ACL, error) {
	objects, err := client.listObjects(aclTable, func() interface{} { return &types.ACL{} })
	if err != nil {
		return nil, err
	}

	var acls []*types.ACL
	for _, obj := range objects {
		acls = append(acls, obj.(*types.ACL))
	}
	return acls, nil
}

func (client *sqliteClient) DeleteACL(id string) error {
	return client.deleteObject(id, aclTable)
}

// Placeholder implementations for all other required methods
// These would need proper implementation based on the specific types

func (client *sqliteClient) PutRouteTable(routeTable *types.RouteTable) error {
	return client.putObject(routeTable, routeTable.DbId(), routeTableTable)
}

func (client *sqliteClient) GetRouteTable(id string) (*types.RouteTable, error) {
	rt := &types.RouteTable{}
	err := client.getObject(id, routeTableTable, rt)
	return rt, err
}

func (client *sqliteClient) ListRouteTables() ([]*types.RouteTable, error) {
	objects, err := client.listObjects(routeTableTable, func() interface{} { return &types.RouteTable{} })
	if err != nil {
		return nil, err
	}

	var routeTables []*types.RouteTable
	for _, obj := range objects {
		routeTables = append(routeTables, obj.(*types.RouteTable))
	}
	return routeTables, nil
}

func (client *sqliteClient) DeleteRouteTable(id string) error {
	return client.deleteObject(id, routeTableTable)
}

func (client *sqliteClient) PutNATGateway(ng *types.NATGateway) error {
	return client.putObject(ng, ng.DbId(), ngTable)
}
func (client *sqliteClient) GetNATGateway(id string) (*types.NATGateway, error) {
	ng := &types.NATGateway{}
	err := client.getObject(id, ngTable, ng)
	return ng, err
}

func (client *sqliteClient) ListNATGateways() ([]*types.NATGateway, error) {
	objects, err := client.listObjects(ngTable, func() interface{} { return &types.NATGateway{} })
	if err != nil {
		return nil, err
	}

	var natGateways []*types.NATGateway
	for _, obj := range objects {
		natGateways = append(natGateways, obj.(*types.NATGateway))
	}
	return natGateways, nil
}

func (client *sqliteClient) DeleteNATGateway(id string) error {
	_, err := client.db.Exec("DELETE FROM "+ngTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutRouter(router *types.Router) error {
	return client.putObject(router, router.DbId(), routerTable)
}
func (client *sqliteClient) GetRouter(id string) (*types.Router, error) {
	router := &types.Router{}
	err := client.getObject(id, routerTable, router)
	return router, err
}

func (client *sqliteClient) ListRouters() ([]*types.Router, error) {
	objects, err := client.listObjects(routerTable, func() interface{} { return &types.Router{} })
	if err != nil {
		return nil, err
	}

	var routers []*types.Router
	for _, obj := range objects {
		routers = append(routers, obj.(*types.Router))
	}
	return routers, nil
}

func (client *sqliteClient) DeleteRouter(id string) error {
	_, err := client.db.Exec("DELETE FROM "+routerTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutIGW(igw *types.IGW) error {
	return client.putObject(igw, igw.DbId(), igwTable)
}
func (client *sqliteClient) GetIGW(id string) (*types.IGW, error) {
	igw := &types.IGW{}
	err := client.getObject(id, igwTable, igw)
	return igw, err
}

func (client *sqliteClient) ListInternetGateways() ([]*types.IGW, error) {
	objects, err := client.listObjects(igwTable, func() interface{} { return &types.IGW{} })
	if err != nil {
		return nil, err
	}

	var igws []*types.IGW
	for _, obj := range objects {
		igws = append(igws, obj.(*types.IGW))
	}
	return igws, nil
}

func (client *sqliteClient) DeleteIGW(id string) error {
	_, err := client.db.Exec("DELETE FROM "+igwTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutVPCEndpoint(endpoint *types.VPCEndpoint) error {
	return client.putObject(endpoint, endpoint.DbId(), vpcEndpointTable)
}
func (client *sqliteClient) GetVPCEndpoint(id string) (*types.VPCEndpoint, error) {
	endpoint := &types.VPCEndpoint{}
	err := client.getObject(id, vpcEndpointTable, endpoint)
	return endpoint, err
}

func (client *sqliteClient) ListVPCEndpoints() ([]*types.VPCEndpoint, error) {
	objects, err := client.listObjects(vpcEndpointTable, func() interface{} { return &types.VPCEndpoint{} })
	if err != nil {
		return nil, err
	}

	var endpoints []*types.VPCEndpoint
	for _, obj := range objects {
		endpoints = append(endpoints, obj.(*types.VPCEndpoint))
	}
	return endpoints, nil
}

func (client *sqliteClient) DeleteVPCEndpoint(id string) error {
	_, err := client.db.Exec("DELETE FROM "+vpcEndpointTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutPublicIP(publicIP *types.PublicIP) error {
	return client.putObject(publicIP, publicIP.DbId(), publicIPTable)
}
func (client *sqliteClient) GetPublicIP(id string) (*types.PublicIP, error) {
	publicIP := &types.PublicIP{}
	err := client.getObject(id, publicIPTable, publicIP)
	return publicIP, err
}

func (client *sqliteClient) ListPublicIPs() ([]*types.PublicIP, error) {
	objects, err := client.listObjects(publicIPTable, func() interface{} { return &types.PublicIP{} })
	if err != nil {
		return nil, err
	}

	var publicIPs []*types.PublicIP
	for _, obj := range objects {
		publicIPs = append(publicIPs, obj.(*types.PublicIP))
	}
	return publicIPs, nil
}

func (client *sqliteClient) DeletePublicIP(id string) error {
	_, err := client.db.Exec("DELETE FROM "+publicIPTable+" WHERE id = ?", id)
	return err
}

// Security Group methods
func (client *sqliteClient) PutSecurityGroup(sg *types.SecurityGroup) error {
	return client.putObject(sg, sg.DbId(), securityGroupTable)
}

func (client *sqliteClient) GetSecurityGroup(id string) (*types.SecurityGroup, error) {
	sg := &types.SecurityGroup{}
	err := client.getObject(id, securityGroupTable, sg)
	return sg, err
}

func (client *sqliteClient) ListSecurityGroups() ([]*types.SecurityGroup, error) {
	objects, err := client.listObjects(securityGroupTable, func() interface{} { return &types.SecurityGroup{} })
	if err != nil {
		return nil, err
	}

	var securityGroups []*types.SecurityGroup
	for _, obj := range objects {
		securityGroups = append(securityGroups, obj.(*types.SecurityGroup))
	}
	return securityGroups, nil
}

func (client *sqliteClient) DeleteSecurityGroup(id string) error {
	return client.deleteObject(id, securityGroupTable)
}

func (client *sqliteClient) PutLB(lb *types.LB) error {
	return client.putObject(lb, lb.DbId(), lbTable)
}
func (client *sqliteClient) GetLB(id string) (*types.LB, error) {
	lb := &types.LB{}
	err := client.getObject(id, lbTable, lb)
	return lb, err
}

func (client *sqliteClient) ListLBs() ([]*types.LB, error) {
	objects, err := client.listObjects(lbTable, func() interface{} { return &types.LB{} })
	if err != nil {
		return nil, err
	}

	var lbs []*types.LB
	for _, obj := range objects {
		lbs = append(lbs, obj.(*types.LB))
	}
	return lbs, nil
}

func (client *sqliteClient) DeleteLB(id string) error {
	_, err := client.db.Exec("DELETE FROM "+lbTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutNetworkInterface(ni *types.NetworkInterface) error {
	return client.putObject(ni, ni.DbId(), networkInterfaceTable)
}
func (client *sqliteClient) GetNetworkInterface(id string) (*types.NetworkInterface, error) {
	ni := &types.NetworkInterface{}
	err := client.getObject(id, networkInterfaceTable, ni)
	return ni, err
}

func (client *sqliteClient) ListNetworkInterfaces() ([]*types.NetworkInterface, error) {
	objects, err := client.listObjects(networkInterfaceTable, func() interface{} { return &types.NetworkInterface{} })
	if err != nil {
		return nil, err
	}

	var networkInterfaces []*types.NetworkInterface
	for _, obj := range objects {
		networkInterfaces = append(networkInterfaces, obj.(*types.NetworkInterface))
	}
	return networkInterfaces, nil
}

func (client *sqliteClient) DeleteNetworkInterface(id string) error {
	_, err := client.db.Exec("DELETE FROM "+networkInterfaceTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutKeyPair(kp *types.KeyPair) error {
	return client.putObject(kp, kp.DbId(), keyPairTable)
}
func (client *sqliteClient) GetKeyPair(id string) (*types.KeyPair, error) {
	kp := &types.KeyPair{}
	err := client.getObject(id, keyPairTable, kp)
	return kp, err
}

func (client *sqliteClient) ListKeyPairs() ([]*types.KeyPair, error) {
	objects, err := client.listObjects(keyPairTable, func() interface{} { return &types.KeyPair{} })
	if err != nil {
		return nil, err
	}

	var keyPairs []*types.KeyPair
	for _, obj := range objects {
		keyPairs = append(keyPairs, obj.(*types.KeyPair))
	}
	return keyPairs, nil
}

func (client *sqliteClient) DeleteKeyPair(id string) error {
	_, err := client.db.Exec("DELETE FROM "+keyPairTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutVPNConcentrator(vpn *types.VPNConcentrator) error {
	return client.putObject(vpn, vpn.DbId(), vpnConcentratorTable)
}
func (client *sqliteClient) GetVPNConcentrator(id string) (*types.VPNConcentrator, error) {
	vpn := &types.VPNConcentrator{}
	err := client.getObject(id, vpnConcentratorTable, vpn)
	return vpn, err
}

func (client *sqliteClient) ListVPNConcentrators() ([]*types.VPNConcentrator, error) {
	objects, err := client.listObjects(vpnConcentratorTable, func() interface{} { return &types.VPNConcentrator{} })
	if err != nil {
		return nil, err
	}

	var vpnConcentrators []*types.VPNConcentrator
	for _, obj := range objects {
		vpnConcentrators = append(vpnConcentrators, obj.(*types.VPNConcentrator))
	}
	return vpnConcentrators, nil
}

func (client *sqliteClient) DeleteVPNConcentrator(id string) error {
	_, err := client.db.Exec("DELETE FROM "+vpnConcentratorTable+" WHERE id = ?", id)
	return err
}

// Kubernetes specific
func (client *sqliteClient) PutCluster(cluster *types.Cluster) error {
	return client.putObject(cluster, cluster.DbId(), clusterTable)
}
func (client *sqliteClient) GetCluster(id string) (*types.Cluster, error) {
	cluster := &types.Cluster{}
	err := client.getObject(id, clusterTable, cluster)
	return cluster, err
}
func (client *sqliteClient) ListClusters() ([]*types.Cluster, error) {
	objects, err := client.listObjects(clusterTable, func() interface{} { return &types.Cluster{} })
	if err != nil {
		return nil, err
	}

	var clusters []*types.Cluster
	for _, obj := range objects {
		clusters = append(clusters, obj.(*types.Cluster))
	}
	return clusters, nil
}
func (client *sqliteClient) DeleteCluster(id string) error {
	_, err := client.db.Exec("DELETE FROM "+clusterTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutPod(pod *types.Pod) error {
	return client.putObject(pod, pod.DbId(), podTable)
}
func (client *sqliteClient) GetPod(id string) (*types.Pod, error) {
	pod := &types.Pod{}
	err := client.getObject(id, podTable, pod)
	return pod, err
}
func (client *sqliteClient) ListPods() ([]*types.Pod, error) {
	objects, err := client.listObjects(podTable, func() interface{} { return &types.Pod{} })
	if err != nil {
		return nil, err
	}

	var pods []*types.Pod
	for _, obj := range objects {
		pods = append(pods, obj.(*types.Pod))
	}
	return pods, nil
}
func (client *sqliteClient) DeletePod(id string) error {
	_, err := client.db.Exec("DELETE FROM "+podTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutKubernetesService(service *types.K8SService) error {
	return client.putObject(service, service.DbId(), kubernetesServiceTable)
}
func (client *sqliteClient) GetKubernetesService(id string) (*types.K8SService, error) {
	service := &types.K8SService{}
	err := client.getObject(id, kubernetesServiceTable, service)
	return service, err
}
func (client *sqliteClient) ListKubernetesServices() ([]*types.K8SService, error) {
	objects, err := client.listObjects(kubernetesServiceTable, func() interface{} { return &types.K8SService{} })
	if err != nil {
		return nil, err
	}

	var services []*types.K8SService
	for _, obj := range objects {
		services = append(services, obj.(*types.K8SService))
	}
	return services, nil
}
func (client *sqliteClient) DeleteKubernetesService(id string) error {
	_, err := client.db.Exec("DELETE FROM "+kubernetesServiceTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutKubernetesNode(node *types.K8sNode) error {
	return client.putObject(node, node.DbId(), kubernetesNodeTable)
}
func (client *sqliteClient) GetKubernetesNode(id string) (*types.K8sNode, error) {
	node := &types.K8sNode{}
	err := client.getObject(id, kubernetesNodeTable, node)
	return node, err
}
func (client *sqliteClient) ListKubernetesNodes() ([]*types.K8sNode, error) {
	objects, err := client.listObjects(kubernetesNodeTable, func() interface{} { return &types.K8sNode{} })
	if err != nil {
		return nil, err
	}

	var nodes []*types.K8sNode
	for _, obj := range objects {
		nodes = append(nodes, obj.(*types.K8sNode))
	}
	return nodes, nil
}
func (client *sqliteClient) DeleteKubernetesNode(id string) error {
	_, err := client.db.Exec("DELETE FROM "+kubernetesNodeTable+" WHERE id = ?", id)
	return err
}

func (client *sqliteClient) PutNamespace(namespace *types.Namespace) error {
	return client.putObject(namespace, namespace.DbId(), namespaceTable)
}
func (client *sqliteClient) GetNamespace(id string) (*types.Namespace, error) {
	namespace := &types.Namespace{}
	err := client.getObject(id, namespaceTable, namespace)
	return namespace, err
}
func (client *sqliteClient) ListNamespaces() ([]*types.Namespace, error) {
	objects, err := client.listObjects(namespaceTable, func() interface{} { return &types.Namespace{} })
	if err != nil {
		return nil, err
	}

	var namespaces []*types.Namespace
	for _, obj := range objects {
		namespaces = append(namespaces, obj.(*types.Namespace))
	}
	return namespaces, nil
}
func (client *sqliteClient) DeleteNamespace(id string) error {
	_, err := client.db.Exec("DELETE FROM "+namespaceTable+" WHERE id = ?", id)
	return err
}

// Resource sync times
func (client *sqliteClient) PutSyncTime(syncTime *types.SyncTime) error {
	return client.putObject(syncTime, syncTime.DbId(), syncTimeTable)
}
func (client *sqliteClient) GetSyncTime(id string) (*types.SyncTime, error) {
	syncTime := &types.SyncTime{}
	err := client.getObject(id, syncTimeTable, syncTime)
	return syncTime, err
}
func (client *sqliteClient) ListSyncTimes() ([]*types.SyncTime, error) {
	objects, err := client.listObjects(syncTimeTable, func() interface{} { return &types.SyncTime{} })
	if err != nil {
		return nil, err
	}

	var syncTimes []*types.SyncTime
	for _, obj := range objects {
		syncTimes = append(syncTimes, obj.(*types.SyncTime))
	}
	return syncTimes, nil
}
func (client *sqliteClient) DeleteSyncTime(id string) error {
	_, err := client.db.Exec("DELETE FROM "+syncTimeTable+" WHERE id = ?", id)
	return err
}
