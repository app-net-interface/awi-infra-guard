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

package sync

import (
	"context"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/app-net-interface/awi-infra-guard/grpc/config"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/app-net-interface/awi-infra-guard/db"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

func NewSyncer(l *logrus.Logger, dbClient db.Client, strategy provider.Strategy,
	sc *config.SyncConfig) *Syncer {
	return &Syncer{
		sc:       sc,
		logger:   l,
		dbClient: dbClient,
		strategy: strategy,
	}
}

type Syncer struct {
	sc       *config.SyncConfig
	logger   *logrus.Logger
	dbClient db.Client
	strategy provider.Strategy
}

func (s *Syncer) ParallelSync(ctx context.Context, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	s.logger.Infof("*****************Sync Start*****************")
	allResource := s.sc.HasCloudResource("all")
	var wg sync.WaitGroup

	if allResource || s.sc.HasCloudResource("region") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncRegions(ctx)
		}()
	}

	if allResource || s.sc.HasCloudResource("vpc") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncVPC(ctx)
		}()
	}

	if allResource || s.sc.HasCloudResource("instance") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncInstances(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("publicip") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncPublicIPs(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("subnet") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncSubnets(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("acl") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncACLs(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("routetable") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncRouteTables(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("securitygroup") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncSecurityGroups(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("natgateway") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncNATGateways(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("router") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncRouters(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("internetgateway") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncIGWs(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("vpcendpoint") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncVPCEndpoints(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("lb") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncLBs(ctx)
		}()
	}
	if allResource || s.sc.HasCloudResource("networkinterface") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			s.syncNetworkInterfaces(ctx)
		}()
	}

	// Kubernetes

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		s.syncClusters(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		s.syncPods(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		s.syncNamespaces(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		s.syncK8SServices(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		s.syncK8SSsNodes(ctx)
	}()

	wg.Wait()

	s.logger.Infof("*****************Cloud Sync End*****************")
	s.logger.Infof("*****************K8S Sync Start*****************")
	s.logger.Infof("*****************K8S Sync End*****************")
}

func (s *Syncer) Sync(ctx context.Context, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	s.logger.Infof("*****************Sync Start*****************")
	allResource := s.sc.HasCloudResource("all")

	if allResource || s.sc.HasCloudResource("region") {
		s.syncRegions(ctx)
	}

	if allResource || s.sc.HasCloudResource("vpc") {
		s.syncVPC(ctx)
	}

	if allResource || s.sc.HasCloudResource("instance") {

		s.syncInstances(ctx)
	}
	if allResource || s.sc.HasCloudResource("publicip") {

		s.syncPublicIPs(ctx)
	}
	if allResource || s.sc.HasCloudResource("subnet") {

		s.syncSubnets(ctx)
	}
	if allResource || s.sc.HasCloudResource("acl") {

		s.syncACLs(ctx)
	}
	if allResource || s.sc.HasCloudResource("routetable") {

		s.syncRouteTables(ctx)
	}
	if allResource || s.sc.HasCloudResource("securitygroup") {

		s.syncSecurityGroups(ctx)
	}
	if allResource || s.sc.HasCloudResource("natgateway") {

		s.syncNATGateways(ctx)
	}
	if allResource || s.sc.HasCloudResource("router") {

		s.syncRouters(ctx)
	}
	if allResource || s.sc.HasCloudResource("internetgateway") {

		s.syncIGWs(ctx)
	}
	if allResource || s.sc.HasCloudResource("vpcendpoint") {

		s.syncVPCEndpoints(ctx)
	}
	if allResource || s.sc.HasCloudResource("lb") {
		s.syncLBs(ctx)
	}
	if allResource || s.sc.HasCloudResource("networkinterface") {
		s.syncNetworkInterfaces(ctx)
	}

	// Kubernetes
	s.logger.Errorf("*****************Cloud Sync End*****************")

	s.logger.Errorf("*****************K8S Sync Start*****************")

	s.syncClusters(ctx)

	s.syncPods(ctx)

	s.syncNamespaces(ctx)

	s.syncK8SServices(ctx)

	s.syncK8SSsNodes(ctx)

	s.logger.Errorf("*****************K8S Sync End*****************")
}

func (s *Syncer) SyncPeriodically(ctx context.Context) {
	s.logger.Infof("Starting periodical sync of cloud resources every %s seconds", s.sc.SyncWaitTime.String())
	done := make(chan struct{}, 1)
	s.Sync(ctx, done)
	ticker := time.NewTicker(s.sc.SyncWaitTime)
	for {
		select {
		case <-ticker.C:
			go func() {
				select {
				case <-done:
					s.Sync(ctx, done)
				default:
					s.logger.Errorf("Previous sync operation is still running, skipping this sync")
				}
			}()
		case <-ctx.Done():
			s.logger.Infof("Context done - time out")
		}
	}
}
func (s *Syncer) syncRegions(ctx context.Context) (e error) {
	e = genericCloudSync[*types.Region](s, types.RegionType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Region, error) {
		return cloudProvider.ListRegions(ctx, &infrapb.ListRegionsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRegions, s.dbClient.PutRegion, s.dbClient.DeleteRegion)
	if e != nil {
		s.logger.Errorf("Sync error: failed to sync regions: %v", e)
	}
	return
}

func (s *Syncer) syncVPC(ctx context.Context) {
	genericCloudSync[*types.VPC](s, types.VPCType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPC, error) {
		return cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCs, s.dbClient.PutVPC, s.dbClient.DeleteVPC)
}

func (s *Syncer) syncInstances(ctx context.Context) {
	genericCloudSync[*types.Instance](s, types.InstanceType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Instance, error) {
		return cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInstances, s.dbClient.PutInstance, s.dbClient.DeleteInstance)
}

func (s *Syncer) syncPublicIPs(ctx context.Context) {
	genericCloudSync[*types.PublicIP](s, types.PublicIPType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.PublicIP, error) {
		return cloudProvider.ListPublicIPs(ctx, &infrapb.ListPublicIPsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListPublicIPs, s.dbClient.PutPublicIP, s.dbClient.DeletePublicIP)
}

func (s *Syncer) syncSubnets(ctx context.Context) {
	genericCloudSync[*types.Subnet](s, types.SubnetType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Subnet, error) {
		return cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSubnets, s.dbClient.PutSubnet, s.dbClient.DeleteSubnet)
}

func (s *Syncer) syncACLs(ctx context.Context) {
	genericCloudSync[*types.ACL](s, types.ACLType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.ACL, error) {
		return cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListACLs, s.dbClient.PutACL, s.dbClient.DeleteACL)
}

func (s *Syncer) syncSecurityGroups(ctx context.Context) {
	genericCloudSync[*types.SecurityGroup](s, types.SecurityGroupType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.SecurityGroup, error) {
		return cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSecurityGroups, s.dbClient.PutSecurityGroup, s.dbClient.DeleteSecurityGroup)
}

func (s *Syncer) syncRouteTables(ctx context.Context) {
	genericCloudSync[*types.RouteTable](s, types.RouteTableType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.RouteTable, error) {
		return cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouteTables, s.dbClient.PutRouteTable, s.dbClient.DeleteRouteTable)
}

func (s *Syncer) syncNATGateways(ctx context.Context) {
	genericCloudSync[*types.NATGateway](s, types.NATGatewayType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.NATGateway, error) {
		return cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListNATGateways, s.dbClient.PutNATGateway, s.dbClient.DeleteNATGateway)
}

func (s *Syncer) syncRouters(ctx context.Context) {
	genericCloudSync[*types.Router](s, types.RouterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Router, error) {
		return cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouters, s.dbClient.PutRouter, s.dbClient.DeleteRouter)
}

func (s *Syncer) syncIGWs(ctx context.Context) {
	genericCloudSync[*types.IGW](s, types.IGWType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.IGW, error) {
		return cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInternetGateways, s.dbClient.PutIGW, s.dbClient.DeleteIGW)
}

func (s *Syncer) syncVPCEndpoints(ctx context.Context) {
	genericCloudSync[*types.VPCEndpoint](s, types.VPCEndpointType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPCEndpoint, error) {
		return cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCEndpoints, s.dbClient.PutVPCEndpoint, s.dbClient.DeleteVPCEndpoint)
}

func (s *Syncer) syncLBs(ctx context.Context) {
	genericCloudSync[*types.LB](s, types.LBType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.LB, error) {
		return cloudProvider.ListLBs(ctx, &infrapb.ListLBsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListLBs, s.dbClient.PutLB, s.dbClient.DeleteLB)
}

func (s *Syncer) syncNetworkInterfaces(ctx context.Context) {
	genericCloudSync[*types.NetworkInterface](s, types.NetworkInterfaceType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.NetworkInterface, error) {
		return cloudProvider.ListNetworkInterfaces(ctx, &infrapb.ListNetworkInterfacesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListNetworkInterfaces, s.dbClient.PutNetworkInterface, s.dbClient.DeleteNetworkInterface)
}

/* End sync cloud resources */
/* Start sync kubernetes resources */

func (s *Syncer) syncClusters(ctx context.Context) {
	genericCloudSync[*types.Cluster](s, types.ClusterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Cluster, error) {
		return cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListClusters, s.dbClient.PutCluster, s.dbClient.DeleteCluster)
}

func (s *Syncer) syncPods(ctx context.Context) {
	genericK8sSync[*types.Pod](s, types.PodsType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Pod, error) {
		return k8sProvider.ListPods(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListPods, s.dbClient.PutPod, s.dbClient.DeletePod)
}

func (s *Syncer) syncNamespaces(ctx context.Context) {
	genericK8sSync[*types.Namespace](s, types.NamespaceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Namespace, error) {
		return k8sProvider.ListNamespaces(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListNamespaces, s.dbClient.PutNamespace, s.dbClient.DeleteNamespace)
}

func (s *Syncer) syncK8SServices(ctx context.Context) {
	genericK8sSync[*types.K8SService](s, types.K8sServiceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.K8SService, error) {
		return k8sProvider.ListServices(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListKubernetesServices, s.dbClient.PutKubernetesService, s.dbClient.DeleteKubernetesService)
}

func (s *Syncer) syncK8SSsNodes(ctx context.Context) {
	genericK8sSync[*types.K8sNode](s, types.K8sNodeType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.K8sNode, error) {
		return k8sProvider.ListNodes(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListKubernetesNodes, s.dbClient.PutKubernetesNode, s.dbClient.DeleteKubernetesNode)
}

func genericCloudSync[P interface {
	*T
	db.DbObject
}, T any](s *Syncer,
	typeName string,
	listF func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]T, error),
	logger *logrus.Logger,
	listFunc func() ([]P, error),
	putFunc func(P) error,
	deleteFunc func(string) error) error {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*900)
	defer cancel()
	var allRemoteObj []P
	syncTime := make(map[string]types.SyncTime)
	//var tType T
	for _, cloudProvider := range s.strategy.GetAllProviders() {
		t := time.Now().UTC().Format(time.RFC3339)
		ok := false
		for _, account := range cloudProvider.ListAccounts() {
			s.logger.Debugf("Calling function %s for account %s and provider %s ", runtime.FuncForPC(reflect.ValueOf(listF).Pointer()).Name(), account.ID, cloudProvider.GetName())
			remoteObjs, err := listF(ctx, cloudProvider, account.ID)
			if err != nil {
				s.logger.Errorf("Sync error: failed to List %s in provider %s: %v",
					typeName, cloudProvider.GetName(), err)
				continue
			} else {
				ok = true
			}
			ptrs := make([]P, len(remoteObjs))
			for k, v := range remoteObjs {
				v := v
				p := P(new(T))
				p = &v
				p.SetSyncTime(t)
				ptrs[k] = p
			}
			allRemoteObj = append(allRemoteObj, ptrs...)
		}
		if ok {
			syncTime[types.SyncTimeKey(cloudProvider.GetName(), typeName)] = types.SyncTime{
				Provider:     cloudProvider.GetName(),
				ResourceType: typeName,
				Time:         t,
			}
		}
	}

	local, err := listFunc()
	if err != nil {
		logger.Errorf("Sync error: failed to List %T in database: %v", new(*T), err)
		return err
	}
	remoteSet := make(map[string]struct{})
	for _, obj := range allRemoteObj {
		if obj == nil {
			continue
		}
		remoteSet[obj.DbId()] = struct{}{}
		err := putFunc(obj)
		if err != nil {
			logger.Errorf("Sync error: failed to put %T in database: %v", obj, err)
		}
	}
	// delete those which don't exist anymore
	for _, localObj := range local {
		if localObj == nil {
			continue
		}
		_, ok := remoteSet[localObj.DbId()]
		if ok {
			continue
		}

		// assuming lost connection to this provider, not removing existing objects
		_, ok = syncTime[types.SyncTimeKey(localObj.GetProvider(), typeName)]
		if !ok {
			continue
		}
		err := deleteFunc(localObj.DbId())
		if err != nil {
			logger.Errorf("Sync error: failed to delete %T from database: %v", localObj, err)
		}
	}
	for _, v := range syncTime {
		err := s.dbClient.PutSyncTime(&v)
		if err != nil {
			logger.Errorf("Sync error: to put sync time of %s", v.DbId())
		}
	}
	return err
}

func genericK8sSync[P interface {
	*T
	db.DbObject
}, T any](s *Syncer,
	typeName string,
	listF func(ctx context.Context, cloudProvider provider.Kubernetes, clusterName string) ([]T, error),
	logger *logrus.Logger,
	listFunc func() ([]P, error),
	putFunc func(P) error,
	deleteFunc func(string) error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	var allRemoteObjs []P
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		s.logger.Errorf("Error in sync: failed to get kuberntes provider: %v", err)
		return
	}
	clusters, err := k8sProvider.ListClusters(ctx)
	if err != nil {
		s.logger.Warnf("Error in sync: failed to list clusters: %v", err)
		return
	}
	syncTime := make(map[string]types.SyncTime)
	for _, cluster := range clusters {
		//s.logger.Infof("Syncing pods in cluster %s", cluster.Name)
		t := time.Now().UTC().Format(time.RFC3339)
		remoteObjs, err := listF(ctx, k8sProvider, cluster.Name)
		if err != nil {
			s.logger.Warnf("Sync error: failed to access %s in cluster %s: %v",
				typeName, cluster.Name, err)
			continue
		}
		syncTime[types.SyncTimeKey(cluster.Name, typeName)] = types.SyncTime{
			Provider:     cluster.Name,
			ResourceType: typeName,
			Time:         t,
		}
		ptrs := make([]P, len(remoteObjs))
		for k, v := range remoteObjs {
			v := v
			p := P(new(T))
			p = &v
			p.SetSyncTime(t)
			ptrs[k] = p
		}

		allRemoteObjs = append(allRemoteObjs, ptrs...)
	}

	local, err := listFunc()
	if err != nil {
		logger.Errorf("Sync error: failed to List %T in database: %v", new(*T), err)
		return
	}
	remoteSet := make(map[string]struct{})
	for _, obj := range allRemoteObjs {
		if obj == nil {
			continue
		}
		remoteSet[obj.DbId()] = struct{}{}
		err := putFunc(obj)
		if err != nil {
			logger.Errorf("Sync error: failed to put %T in database: %v", obj, err)
		}
	}
	// delete those which don't exist anymore
	for _, localObj := range local {
		if localObj == nil {
			continue
		}
		_, ok := remoteSet[localObj.DbId()]
		if ok {
			continue
		}
		// assuming lost connection to this cluster, not removing existing objects
		_, ok = syncTime[types.SyncTimeKey(localObj.GetProvider(), typeName)]
		if !ok {
			continue
		}
		err := deleteFunc(localObj.DbId())
		if err != nil {
			logger.Errorf("Sync error: failed to delete %T from database: %v", localObj, err)
		}
	}
	for _, v := range syncTime {
		err := s.dbClient.PutSyncTime(&v)
		if err != nil {
			logger.Errorf("Sync error: to put sync time of %s", v.DbId())
		}
	}
}
