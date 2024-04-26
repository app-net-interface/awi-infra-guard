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
	"time"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/app-net-interface/awi-infra-guard/db"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

func NewSyncer(l *logrus.Logger, dbClient db.Client, strategy provider.Strategy,
	waitTime time.Duration) *Syncer {
	return &Syncer{
		waitTime: waitTime,
		logger:   l,
		dbClient: dbClient,
		strategy: strategy,
	}
}

type Syncer struct {
	waitTime time.Duration
	logger   *logrus.Logger
	dbClient db.Client
	strategy provider.Strategy
}

func (s *Syncer) Sync() {

	s.syncVPC()
	s.syncInstances()
	s.syncSubnets()
	s.syncRouteTables()
	s.syncACLs()
	s.syncSecurityGroups()
	s.syncNATGateways()
	s.syncRouters()
	s.syncIGWs()
	s.syncVPCEndpoints()

	// Kubernetes
	s.syncClusters()
	s.syncPods()

	s.syncNamespaces()
	s.syncK8SSsNodes()
	s.syncK8SServices()
}

func (s *Syncer) SyncPeriodically(ctx context.Context) {
	s.logger.Infof("Starting periodical sync of cloud resources")
	s.Sync()
	ticker := time.NewTicker(s.waitTime)
	for {
		select {
		case <-ticker.C:
			s.Sync()
		case <-ctx.Done():
			return
		}
	}
}

func (s *Syncer) syncVPC() {
	genericCloudSync[*types.VPC](s, types.VPCType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPC, error) {
		return cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCs, s.dbClient.PutVPC, s.dbClient.DeleteVPC)
}

func (s *Syncer) syncInstances() {
	genericCloudSync[*types.Instance](s, types.InstanceType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Instance, error) {
		return cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInstances, s.dbClient.PutInstance, s.dbClient.DeleteInstance)
}

func (s *Syncer) syncSubnets() {
	genericCloudSync[*types.Subnet](s, types.SubnetType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Subnet, error) {
		return cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSubnets, s.dbClient.PutSubnet, s.dbClient.DeleteSubnet)
}

func (s *Syncer) syncACLs() {
	genericCloudSync[*types.ACL](s, types.ACLType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.ACL, error) {
		return cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListACLs, s.dbClient.PutACL, s.dbClient.DeleteACL)
}

func (s *Syncer) syncSecurityGroups() {
	genericCloudSync[*types.SecurityGroup](s, types.SecurityGroupType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.SecurityGroup, error) {
		return cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSecurityGroups, s.dbClient.PutSecurityGroup, s.dbClient.DeleteSecurityGroup)
}

func (s *Syncer) syncRouteTables() {
	genericCloudSync[*types.RouteTable](s, types.RouteTableType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.RouteTable, error) {
		return cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouteTables, s.dbClient.PutRouteTable, s.dbClient.DeleteRouteTable)
}

func (s *Syncer) syncNATGateways() {
	genericCloudSync[*types.NATGateway](s, types.NATGatewayType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.NATGateway, error) {
		return cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListNATGateways, s.dbClient.PutNATGateway, s.dbClient.DeleteNATGateway)
}

func (s *Syncer) syncRouters() {
	genericCloudSync[*types.Router](s, types.RouterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Router, error) {

		return cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouters, s.dbClient.PutRouter, s.dbClient.DeleteRouter)
}

func (s *Syncer) syncIGWs() {
	genericCloudSync[*types.IGW](s, types.IGWType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.IGW, error) {

		return cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInternetGateways, s.dbClient.PutIGW, s.dbClient.DeleteIGW)
}

func (s *Syncer) syncVPCEndpoints() {
	genericCloudSync[*types.VPCEndpoint](s, types.VPCEndpointType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPCEndpoint, error) {

		return cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCEndpoints, s.dbClient.PutVPCEndpoint, s.dbClient.DeleteVPCEndpoint)
}

func (s *Syncer) syncClusters() {
	genericCloudSync[*types.Cluster](s, types.ClusterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Cluster, error) {
		return cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListClusters, s.dbClient.PutCluster, s.dbClient.DeleteCluster)
}

func (s *Syncer) syncPods() {
	genericK8sSync[*types.Pod](s, types.PodsType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Pod, error) {
		return k8sProvider.ListPods(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListPods, s.dbClient.PutPod, s.dbClient.DeletePod)
}

func (s *Syncer) syncNamespaces() {
	genericK8sSync[*types.Namespace](s, types.NamespaceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Namespace, error) {
		return k8sProvider.ListNamespaces(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListNamespaces, s.dbClient.PutNamespace, s.dbClient.DeleteNamespace)
}

func (s *Syncer) syncK8SServices() {
	genericK8sSync[*types.K8SService](s, types.K8sServiceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.K8SService, error) {
		return k8sProvider.ListServices(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListKubernetesServices, s.dbClient.PutKubernetesService, s.dbClient.DeleteKubernetesService)
}

func (s *Syncer) syncK8SSsNodes() {
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
	deleteFunc func(string) error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	var allRemoteObj []P
	syncTime := make(map[string]types.SyncTime)
	for _, cloudProvider := range s.strategy.GetAllProviders() {
		t := time.Now().UTC().Format(time.RFC3339)
		ok := false
		for _, account := range cloudProvider.ListAccounts() {
			s.logger.Infof("Found account %s and provider %s", account.ID, cloudProvider.GetName())
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
		return
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
		s.logger.Errorf("Error in sync: failed to list clusters: %v", err)
		return
	}
	syncTime := make(map[string]types.SyncTime)
	for _, cluster := range clusters {
		//s.logger.Infof("Syncing pods in cluster %s", cluster.Name)
		t := time.Now().UTC().Format(time.RFC3339)
		remoteObjs, err := listF(ctx, k8sProvider, cluster.Name)
		if err != nil {
			s.logger.Errorf("Sync error: failed to %s in cluster %s: %v",
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
