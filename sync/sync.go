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
	"runtime/debug"
	"sync"
	"time"

	"github.com/app-net-interface/awi-infra-guard/db"
	"github.com/app-net-interface/awi-infra-guard/grpc/config"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

func NewSyncer(l *logrus.Logger, dbClient db.Client, strategy provider.Strategy,
	mainCfg *config.Config) *Syncer { // Changed sc *config.SyncConfig to mainCfg *config.Config
	return &Syncer{
		sc:       &mainCfg.SyncConfig, // Access SyncConfig from the main Config
		mainCfg:  mainCfg,             // Store the main config
		logger:   l,
		dbClient: dbClient,
		strategy: strategy,
	}
}

type Syncer struct {
	sc       *config.SyncConfig
	mainCfg  *config.Config // Added to access top-level configurations like KubernetesSupported
	logger   *logrus.Logger
	dbClient db.Client
	strategy provider.Strategy
}

func (s *Syncer) ParallelSync(ctx context.Context, done chan<- struct{}) {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Errorf("Panic recovered in ParallelSync top level: %v\n%s", r, string(debug.Stack()))
		}
		if done != nil {
			// Ensure 'done' is only signaled once, even if there's an early exit or panic
			select {
			case done <- struct{}{}:
			default: // Avoid blocking if already sent or channel closed
			}
		}
	}()
	s.logger.Infof("*****************Parallel Sync Start*****************")
	overallStartTime := time.Now()

	currentCloudProviders := s.strategy.GetAllProviders() // Get providers for this specific sync cycle
	cloudProviderCount := len(currentCloudProviders)

	if cloudProviderCount == 0 {
		s.logger.Warnf("ParallelSync: No active cloud providers available in the strategy for this sync cycle. Skipping cloud resource sync operations.")
	} else {
		s.logger.Infof("ParallelSync: Starting cloud resource sync with %d active providers.", cloudProviderCount)
	}

	var wg sync.WaitGroup
	var allResourceCloud bool

	// Cloud Resources
	if cloudProviderCount > 0 { // Only launch cloud sync goroutines if providers exist
		allResourceCloud = s.sc.HasCloudResource("all") // From SyncConfig

		if allResourceCloud || s.sc.HasCloudResource("region") {
			wg.Add(1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						s.logger.Errorf("Panic in syncRegions: %v\n%s", r, string(debug.Stack()))
					}
					wg.Done()
				}()
				s.logger.Debugf("Starting parallel sync for regions")
				s.syncRegions(ctx)
				s.logger.Debugf("Finished parallel sync for regions")
			}()
		}

		if allResourceCloud || s.sc.HasCloudResource("vpc") {
			wg.Add(1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						s.logger.Errorf("Panic in syncVPC: %v\n%s", r, string(debug.Stack()))
					}
					wg.Done()
				}()
				s.logger.Debugf("Starting parallel sync for VPCs")
				s.syncVPC(ctx)
				s.logger.Debugf("Finished parallel sync for VPCs")
			}()
		}

		if allResourceCloud || s.sc.HasCloudResource("instance") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for instances")
				s.syncInstances(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for instances")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("publicip") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for public IPs")
				s.syncPublicIPs(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for public IPs")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("subnet") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for subnets")
				s.syncSubnets(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for subnets")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("acl") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for ACLs")
				s.syncACLs(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for ACLs")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("routetable") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for route tables")
				s.syncRouteTables(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for route tables")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("securitygroup") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for security groups")
				s.syncSecurityGroups(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for security groups")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("natgateway") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for NAT gateways")
				s.syncNATGateways(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for NAT gateways")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("router") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for routers")
				s.syncRouters(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for routers")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("internetgateway") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for internet gateways")
				s.syncIGWs(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for internet gateways")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("vpcendpoint") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for VPC endpoints")
				s.syncVPCEndpoints(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for VPC endpoints")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("lb") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for LBs")
				s.syncLBs(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for LBs")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("vpcconnection") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for VPC connections")
				s.syncVpcConnections(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for VPC connections")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("networkinterface") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for network interfaces")
				s.syncNetworkInterfaces(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for network interfaces")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("keypair") { // Added keypair based on your other sync function
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for key pairs")
				s.syncKeyPairs(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for key pairs")
			}()
		}
		if allResourceCloud || s.sc.HasCloudResource("vpnconcentrator") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for VPN concentrators")
				s.syncVPNConcentrators(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for VPN concentrators")
			}()
		}
	}

	// Kubernetes Resources
	// Assuming these checks s.sc.HasCloudResource("k8scluster"), s.sc.HasCloudResource("k8spod") etc. exist or you'll adapt them.
	// For simplicity, I'll use the resource names as they appear in your sync function names.
	if s.mainCfg.KubernetesSupported { // Corrected: Use s.mainCfg.KubernetesSupported
		if allResourceCloud || s.sc.HasKubernetesResource("cluster") { // Assuming "cluster" refers to K8s clusters here
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for K8s clusters")
				s.syncClusters(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for K8s clusters")
			}()
		}

		if allResourceCloud || s.sc.HasKubernetesResource("pod") { // Changed to HasKubernetesResource and simplified name
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for K8s pods")
				s.syncPods(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for K8s pods")
			}()
		}

		if allResourceCloud || s.sc.HasKubernetesResource("namespace") { // Changed to HasKubernetesResource
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for K8s namespaces")
				s.syncNamespaces(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for K8s namespaces")
			}()
		}

		if allResourceCloud || s.sc.HasKubernetesResource("service") { // Changed to HasKubernetesResource
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for K8s services")
				s.syncK8SServices(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for K8s services")
			}()
		}

		if allResourceCloud || s.sc.HasKubernetesResource("node") { // Changed to HasKubernetesResource
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.logger.Debugf("Starting parallel sync for K8s nodes")
				s.syncK8SSsNodes(ctx) // Pass the parent context
				s.logger.Debugf("Finished parallel sync for K8s nodes")
			}()
		}
	} else {
		s.logger.Infof("Kubernetes sync is disabled via configuration (kubernetesSupported: false).")
	}

	s.logger.Info("Waiting for all parallel resource syncs to complete...")
	wg.Wait()
	s.logger.Info("All parallel resource syncs completed.")

	// Sync VPC Index - this should happen after all other resources are synced.
	if cloudProviderCount > 0 && (s.sc.HasCloudResource("all") || s.sc.HasCloudResource("vpcIndex")) {
		s.logger.Info("Starting VPC index sync...")
		vpcIndexStartTime := time.Now()
		s.syncVPCIndex(ctx) // Pass the parent context
		s.logger.Infof("Finished VPC index sync in %v.", time.Since(vpcIndexStartTime))
	} else if s.sc.HasCloudResource("all") || s.sc.HasCloudResource("vpcIndex") {
		s.logger.Warn("Skipping VPC index sync: No active cloud providers were available for this cycle or vpcIndex not configured for sync.")
	}

	s.logger.Infof("*****************Parallel Sync End (Total time: %v)*****************", time.Since(overallStartTime))
}

/*
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
	if allResource || s.sc.HasCloudResource("keypair") {
		s.syncKeyPairs(ctx)
	}
	if allResource || s.sc.HasCloudResource("vpnconcentrator") {
		s.syncVPNConcentrators(ctx)
	}

	// Kubernetes
	s.logger.Errorf("*****************Cloud Sync End*****************")

	s.logger.Errorf("*****************K8S Sync Start*****************")

	s.syncClusters(ctx)

	s.syncPods(ctx)

	s.syncNamespaces(ctx)

	s.syncK8SServices(ctx)

	s.syncK8SSsNodes(ctx)

	// Sync VPC Index
	if allResource || s.sc.HasCloudResource("vpcIndex") {
		s.syncVPCIndex(ctx)
	}

	s.logger.Errorf("*****************K8S Sync End*****************")
}
*/

func (s *Syncer) SyncPeriodically(ctx context.Context) {
	s.logger.Infof("Starting periodical sync of cloud resources every %s seconds", s.sc.SyncWaitTime.String())
	done := make(chan struct{}, 1)
	// First sync
	//s.Sync(ctx, done)
	s.ParallelSync(ctx, done)
	// After the first sync, we start the ticker for periodic syncs
	ticker := time.NewTicker(s.sc.SyncWaitTime)
	defer ticker.Stop() // Ensure ticker is stopped when SyncPeriodically exits

	for {
		select {
		case <-ticker.C:
			s.logger.Debug("Ticker event received for periodic sync.")
			// Launch the sync in a new goroutine to prevent blocking the ticker processing
			// if a sync operation takes longer than the tick interval.
			go func() {
				select {
				case <-done: // Wait for the previous ParallelSync to signal completion
					s.logger.Info("Previous sync completed. Starting new periodic sync.")
					s.ParallelSync(ctx, done) // Pass the main context
					s.logger.Info("Periodic sync attempt finished.")
				default:
					s.logger.Warnf("Previous sync operation is still running or did not signal 'done'. Skipping this sync tick.")
				}
			}()
		case <-ctx.Done():
			s.logger.Infof("SyncPeriodically: Context done. Stopping periodic sync. Error: %v", ctx.Err())
			return // Exit loop when context is done
		}
	}
}

func (s *Syncer) syncRegions(ctx context.Context) (e error) {
	e = genericCloudSync[*types.Region](ctx, s, types.RegionType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Region, error) {
		return cloudProvider.ListRegions(ctx, &infrapb.ListRegionsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRegions, s.dbClient.PutRegion, s.dbClient.DeleteRegion)
	if e != nil {
		s.logger.Errorf("Sync error: failed to sync regions: %v", e)
	}
	return
}

func (s *Syncer) syncVPC(ctx context.Context) {
	genericCloudSync[*types.VPC](ctx, s, types.VPCType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPC, error) {
		return cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCs, s.dbClient.PutVPC, s.dbClient.DeleteVPC)
}

func (s *Syncer) syncVPCIndex(ctx context.Context) {
	if err := s.dbClient.SyncVPCIndexes(); err != nil {
		s.logger.Errorf("Sync error: failed to finalize VPC indexes: %v", err)
	}
}

func (s *Syncer) syncInstances(ctx context.Context) {
	genericCloudSync[*types.Instance](ctx, s, types.InstanceType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Instance, error) {
		return cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInstances, s.dbClient.PutInstance, s.dbClient.DeleteInstance)
}

func (s *Syncer) syncPublicIPs(ctx context.Context) {
	genericCloudSync[*types.PublicIP](ctx, s, types.PublicIPType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.PublicIP, error) {
		return cloudProvider.ListPublicIPs(ctx, &infrapb.ListPublicIPsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListPublicIPs, s.dbClient.PutPublicIP, s.dbClient.DeletePublicIP)
}

func (s *Syncer) syncSubnets(ctx context.Context) {
	genericCloudSync[*types.Subnet](ctx, s, types.SubnetType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Subnet, error) {
		return cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSubnets, s.dbClient.PutSubnet, s.dbClient.DeleteSubnet)
}

func (s *Syncer) syncACLs(ctx context.Context) {
	genericCloudSync[*types.ACL](ctx, s, types.ACLType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.ACL, error) {
		return cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListACLs, s.dbClient.PutACL, s.dbClient.DeleteACL)
}

func (s *Syncer) syncSecurityGroups(ctx context.Context) {
	genericCloudSync[*types.SecurityGroup](ctx, s, types.SecurityGroupType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.SecurityGroup, error) {
		return cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListSecurityGroups, s.dbClient.PutSecurityGroup, s.dbClient.DeleteSecurityGroup)
}

func (s *Syncer) syncRouteTables(ctx context.Context) {
	genericCloudSync[*types.RouteTable](ctx, s, types.RouteTableType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.RouteTable, error) {
		return cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouteTables, s.dbClient.PutRouteTable, s.dbClient.DeleteRouteTable)
}

func (s *Syncer) syncNATGateways(ctx context.Context) {
	genericCloudSync[*types.NATGateway](ctx, s, types.NATGatewayType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.NATGateway, error) {
		return cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListNATGateways, s.dbClient.PutNATGateway, s.dbClient.DeleteNATGateway)
}

func (s *Syncer) syncRouters(ctx context.Context) {
	genericCloudSync[*types.Router](ctx, s, types.RouterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Router, error) {
		return cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListRouters, s.dbClient.PutRouter, s.dbClient.DeleteRouter)
}

func (s *Syncer) syncIGWs(ctx context.Context) {
	genericCloudSync[*types.IGW](ctx, s, types.IGWType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.IGW, error) {
		return cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListInternetGateways, s.dbClient.PutIGW, s.dbClient.DeleteIGW)
}

func (s *Syncer) syncVPCEndpoints(ctx context.Context) {
	genericCloudSync[*types.VPCEndpoint](ctx, s, types.VPCEndpointType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPCEndpoint, error) {
		return cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCEndpoints, s.dbClient.PutVPCEndpoint, s.dbClient.DeleteVPCEndpoint)
}

func (s *Syncer) syncLBs(ctx context.Context) {
	genericCloudSync[*types.LB](ctx, s, types.LBType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.LB, error) {
		return cloudProvider.ListLBs(ctx, &infrapb.ListLBsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListLBs, s.dbClient.PutLB, s.dbClient.DeleteLB)
}

func (s *Syncer) syncNetworkInterfaces(ctx context.Context) {
	genericCloudSync[*types.NetworkInterface](ctx, s, types.NetworkInterfaceType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.NetworkInterface, error) {
		return cloudProvider.ListNetworkInterfaces(ctx, &infrapb.ListNetworkInterfacesRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListNetworkInterfaces, s.dbClient.PutNetworkInterface, s.dbClient.DeleteNetworkInterface)
}

func (s *Syncer) syncKeyPairs(ctx context.Context) {
	genericCloudSync[*types.KeyPair](ctx, s, types.KeyPairType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.KeyPair, error) {
		return cloudProvider.ListKeyPairs(ctx, &infrapb.ListKeyPairsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListKeyPairs, s.dbClient.PutKeyPair, s.dbClient.DeleteKeyPair)
}

func (s *Syncer) syncVPNConcentrators(ctx context.Context) {
	genericCloudSync[*types.VPNConcentrator](ctx, s, types.VPNConcentratorType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPNConcentrator, error) {
		return cloudProvider.ListVPNConcentrators(ctx, &infrapb.ListVPNConcentratorsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPNConcentrators, s.dbClient.PutVPNConcentrator, s.dbClient.DeleteVPNConcentrator)
}

// syncVpcConnections discovers and stores VPC connections from cloud providers
func (s *Syncer) syncVpcConnections(ctx context.Context) {
	genericCloudSync[*types.VPCConnection](ctx, s, types.VPCConnectionType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPCConnection, error) {
		return cloudProvider.ListVpcConnections(ctx, &infrapb.ListVpcConnectionsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVpcConnections, s.dbClient.PutVpcConnection, s.dbClient.DeleteVpcConnection)
}

/* End sync cloud resources */
/* Start sync kubernetes resources */

func (s *Syncer) syncClusters(ctx context.Context) {
	genericCloudSync[*types.Cluster](ctx, s, types.ClusterType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.Cluster, error) {
		return cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListClusters, s.dbClient.PutCluster, s.dbClient.DeleteCluster)
}

func (s *Syncer) syncPods(ctx context.Context) {
	genericK8sSync[*types.Pod](ctx, s, types.PodsType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Pod, error) {
		return k8sProvider.ListPods(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListPods, s.dbClient.PutPod, s.dbClient.DeletePod)
}

func (s *Syncer) syncNamespaces(ctx context.Context) {
	genericK8sSync[*types.Namespace](ctx, s, types.NamespaceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.Namespace, error) {
		return k8sProvider.ListNamespaces(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListNamespaces, s.dbClient.PutNamespace, s.dbClient.DeleteNamespace)
}

func (s *Syncer) syncK8SServices(ctx context.Context) {
	genericK8sSync[*types.K8SService](ctx, s, types.K8sServiceType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.K8SService, error) {
		return k8sProvider.ListServices(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListKubernetesServices, s.dbClient.PutKubernetesService, s.dbClient.DeleteKubernetesService)
}

func (s *Syncer) syncK8SSsNodes(ctx context.Context) {
	genericK8sSync[*types.K8sNode](ctx, s, types.K8sNodeType, func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]types.K8sNode, error) {
		return k8sProvider.ListNodes(ctx, clusterName, nil)
	}, s.logger, s.dbClient.ListKubernetesNodes, s.dbClient.PutKubernetesNode, s.dbClient.DeleteKubernetesNode)
}

func genericCloudSync[P interface {
	*T
	db.DbObject
}, T any](ctx context.Context, s *Syncer, // Added ctx as first parameter
	typeName string,
	listF func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]T, error),
	logger *logrus.Logger,
	listFunc func() ([]P, error),
	putFunc func(P) error,
	deleteFunc func(string) error) error {

	var allRemoteObj []P
	syncTime := make(map[string]types.SyncTime)
	var lastError error

	logger.Debugf("genericCloudSync for %s: Number of providers from strategy: %d", typeName, len(s.strategy.GetAllProviders()))
	for _, cloudProvider := range s.strategy.GetAllProviders() {
		t := time.Now().UTC().Format(time.RFC3339)
		ok := false
		accounts := cloudProvider.ListAccounts()
		logger.Debugf("genericCloudSync for %s, Provider %s: Number of accounts: %d", typeName, cloudProvider.GetName(), len(accounts))

		for _, account := range accounts {
			logger.Debugf("genericCloudSync for %s, Provider %s, Account %s: Calling listF (%s)",
				typeName, cloudProvider.GetName(), account.ID, runtime.FuncForPC(reflect.ValueOf(listF).Pointer()).Name())

			remoteObjs, err := listF(ctx, cloudProvider, account.ID) // Use passed ctx
			if err != nil {
				logger.Errorf("Sync error in genericCloudSync for %s, Provider %s, Account %s: listF failed: %v",
					typeName, cloudProvider.GetName(), account.ID, err)
				if lastError == nil {
					lastError = err
				}
				continue
			}
			ok = true
			logger.Debugf("genericCloudSync for %s, Provider %s, Account %s: listF returned %d objects.",
				typeName, cloudProvider.GetName(), account.ID, len(remoteObjs))

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
		logger.Errorf("Sync error in genericCloudSync for %s: failed to List from database: %v", typeName, err)
		if lastError == nil {
			lastError = err
		}
		return lastError
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
			if lastError == nil { // Capture error from putFunc
				lastError = err
			}
		}
	}
	for _, localObj := range local {
		if localObj == nil {
			continue
		}
		_, ok := remoteSet[localObj.DbId()]
		if ok {
			continue
		}
		_, syncTimeProviderOk := syncTime[types.SyncTimeKey(localObj.GetProvider(), typeName)]
		if !syncTimeProviderOk {
			continue
		}
		err := deleteFunc(localObj.DbId())
		if err != nil {
			logger.Errorf("Sync error: failed to delete %T from database: %v", localObj, err)
			if lastError == nil { // Capture error from deleteFunc
				lastError = err
			}
		}
	}
	for _, v := range syncTime {
		errDb := s.dbClient.PutSyncTime(&v)
		if errDb != nil {
			logger.Errorf("Sync error in genericCloudSync for %s: failed to put sync time for %s: %v", typeName, v.DbId(), errDb)
			if lastError == nil {
				lastError = errDb
			}
		}
	}
	return lastError
}

func genericK8sSync[P interface {
	*T
	db.DbObject
}, T any](ctx context.Context, s *Syncer, // Added ctx as first parameter
	typeName string,
	listF func(ctx context.Context, k8sProvider provider.Kubernetes, clusterName string) ([]T, error),
	logger *logrus.Logger,
	listFunc func() ([]P, error),
	putFunc func(P) error,
	deleteFunc func(string) error) error { // Changed to return error

	var allRemoteObj []P
	syncTime := make(map[string]types.SyncTime) // K8s resources might also need sync time per cluster/provider
	var lastError error

	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		s.logger.Errorf("Error in genericK8sSync for %s: failed to get kubernetes provider: %v", typeName, err)
		return err
	}
	clusters, err := k8sProvider.ListClusters(ctx) // Use passed ctx
	if err != nil {
		s.logger.Warnf("Error in genericK8sSync for %s: failed to list clusters: %v", typeName, err)
		return err
	}

	logger.Debugf("genericK8sSync for %s: Number of clusters from strategy: %d", typeName, len(clusters))
	for _, cluster := range clusters {
		t := time.Now().UTC().Format(time.RFC3339)
		logger.Debugf("genericK8sSync for %s, Cluster %s: Calling listF (%s)",
			typeName, cluster.Name, runtime.FuncForPC(reflect.ValueOf(listF).Pointer()).Name())

		remoteObjs, errList := listF(ctx, k8sProvider, cluster.Name) // Use passed ctx
		if errList != nil {
			s.logger.Warnf("Sync error in genericK8sSync for %s, Cluster %s: listF failed: %v",
				typeName, cluster.Name, errList)
			if lastError == nil {
				lastError = errList
			}
			continue
		}
		logger.Debugf("genericK8sSync for %s, Cluster %s: listF returned %d objects.",
			typeName, cluster.Name, len(remoteObjs))

		ptrs := make([]P, len(remoteObjs))
		for k, v := range remoteObjs {
			v := v
			p := P(new(T))
			p = &v
			p.SetSyncTime(t) // Assuming K8s objects also implement SetSyncTime or similar
			// K8s objects might need a provider/cluster identifier set if not inherently part of their DbId
			// For simplicity, assuming DbId is unique across clusters or GetProvider handles it.
			ptrs[k] = p
		}
		allRemoteObj = append(allRemoteObj, ptrs...)
		// Assuming syncTime is per cluster for K8s resources
		syncTime[types.SyncTimeKey(cluster.Name, typeName)] = types.SyncTime{
			Provider:     cluster.Name, // Or a generic "kubernetes" provider name if cluster is part of DbId
			ResourceType: typeName,
			Time:         t,
		}
	}

	local, err := listFunc()
	if err != nil {
		logger.Errorf("Sync error in genericK8sSync for %s: failed to List from database: %v", typeName, err)
		if lastError == nil {
			lastError = err
		}
		return lastError
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
			if lastError == nil {
				lastError = err
			}
		}
	}

	for _, localObj := range local {
		if localObj == nil {
			continue
		}
		_, ok := remoteSet[localObj.DbId()]
		if ok {
			continue
		}
		// For K8s, the provider key in syncTime might be the cluster name or a general K8s provider ID
		// This depends on how localObj.GetProvider() is implemented for K8s types.
		// Assuming localObj.GetProvider() returns the cluster name for this check.
		_, syncTimeProviderOk := syncTime[types.SyncTimeKey(localObj.GetProvider(), typeName)]
		if !syncTimeProviderOk {
			continue
		}
		err := deleteFunc(localObj.DbId())
		if err != nil {
			logger.Errorf("Sync error: failed to delete %T from database: %v", localObj, err)
			if lastError == nil {
				lastError = err
			}
		}
	}

	for _, v := range syncTime {
		errDb := s.dbClient.PutSyncTime(&v)
		if errDb != nil {
			logger.Errorf("Sync error in genericK8sSync for %s: failed to put sync time for %s: %v", typeName, v.DbId(), errDb)
			if lastError == nil {
				lastError = errDb
			}
		}
	}
	return lastError
}
