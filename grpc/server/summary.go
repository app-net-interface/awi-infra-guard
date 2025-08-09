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
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// Cache for summaries
type cachedSummary struct {
	summary   *infrapb.SummaryResponse
	timestamp time.Time
}

const cacheDuration = 5 * time.Minute

var (
	summaryCache = make(map[string]*cachedSummary)
	cacheMutex   sync.RWMutex
)

type resourceCounts struct {
	accounts, vpcs, subnets, instances, acls, sgs, routeTables int
	natGateways, routers, igws, vpcEndpoints, publicIPs        int
	clusters, pods, services, namespaces                       int
	loadBalancers, networkInterfaces, keyPairs                 int
	vpcConnections, vpnConcentrators                           int

	vmStateSummary   map[string]int32
	vmTypeSummary    map[string]int32
	podsStateSummary map[string]int32
	regionSummary    map[string]int32
	lbTypeSummary    map[string]int32
	healthSummary    map[string]int32

	untaggedResources  int
	publiclyAccessible int
	openSecurityGroups int
	isolatedVpcs       int
	connectedVpcs      int
}

func getCacheKey(in *infrapb.SummaryRequest) string {
	return fmt.Sprintf("%s:%s:%s:%s", in.Provider, in.AccountId, in.Region, in.VpcId)
}

func (s *Server) getCachedSummary(in *infrapb.SummaryRequest) *infrapb.SummaryResponse {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	key := getCacheKey(in)
	if cached, exists := summaryCache[key]; exists {
		if time.Since(cached.timestamp) < cacheDuration {
			return cached.summary
		}
	}
	return nil
}

func (s *Server) cacheSummary(in *infrapb.SummaryRequest, summary *infrapb.SummaryResponse) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	key := getCacheKey(in)
	summaryCache[key] = &cachedSummary{
		summary:   summary,
		timestamp: time.Now(),
	}
}

func (s *Server) processProviderSummary(ctx context.Context, cloudProvider provider.CloudProvider, in *infrapb.SummaryRequest) *resourceCounts {
	counts := &resourceCounts{
		vmStateSummary:   make(map[string]int32),
		vmTypeSummary:    make(map[string]int32),
		podsStateSummary: make(map[string]int32),
		regionSummary:    make(map[string]int32),
		lbTypeSummary:    make(map[string]int32),
		healthSummary:    make(map[string]int32),
	}

	providerName := cloudProvider.GetName()
	var accountIDs []string
	if in.AccountId != "" {
		accountIDs = append(accountIDs, in.AccountId)
	} else {
		for _, acc := range cloudProvider.ListAccounts() {
			accountIDs = append(accountIDs, acc.ID)
		}
	}
	counts.accounts = len(accountIDs)

	for _, accountID := range accountIDs {
		s.processAccountResources(ctx, cloudProvider, providerName, accountID, in, counts)
	}

	return counts
}

func (s *Server) processAccountResources(ctx context.Context, cloudProvider provider.CloudProvider, providerName, accountID string, in *infrapb.SummaryRequest, counts *resourceCounts) {
	// VPCs
	vpcs, err := cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{Provider: providerName, AccountId: accountID, Region: in.Region})
	if err != nil {
		s.logger.Warnf("Failed to list VPCs for %s: %v", providerName, err)
	} else {
		counts.vpcs += len(vpcs)
		for _, vpc := range vpcs {
			counts.regionSummary[vpc.Region]++
			if len(vpc.Labels) == 0 {
				counts.untaggedResources++
			}
		}
	}

	// Subnets
	subnets, err := cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err != nil {
		s.logger.Warnf("Failed to list subnets for %s: %v", providerName, err)
	} else {
		counts.subnets += len(subnets)
		for _, subnet := range subnets {
			if len(subnet.Labels) == 0 {
				counts.untaggedResources++
			}
		}
	}

	// Instances
	instances, err := cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err != nil {
		s.logger.Warnf("Failed to list instances for %s: %v", providerName, err)
	} else {
		counts.instances += len(instances)
		for _, vm := range instances {
			state := strings.ToLower(vm.State)
			counts.vmStateSummary[state]++
			counts.vmTypeSummary[strings.ToLower(vm.Type)]++

			// Health status
			if state == "running" {
				counts.healthSummary["healthy"]++
			} else if state == "stopped" || state == "stopping" {
				counts.healthSummary["stopped"]++
			} else {
				counts.healthSummary["degraded"]++
			}

			if len(vm.Labels) == 0 {
				counts.untaggedResources++
			}
		}
	}

	// Security Groups
	sgs, err := cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err != nil {
		s.logger.Warnf("Failed to list security groups for %s: %v", providerName, err)
	} else {
		counts.sgs += len(sgs)
		for _, sg := range sgs {
			// Check for open security groups (0.0.0.0/0 access)
			for _, rule := range sg.Rules {
				for _, source := range rule.Source {
					if strings.Contains(source, "0.0.0.0/0") {
						counts.openSecurityGroups++
						break
					}
				}
			}
		}
	}

	// Load Balancers
	lbs, err := cloudProvider.ListLBs(ctx, &infrapb.ListLBsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err != nil {
		s.logger.Warnf("Failed to list load balancers for %s: %v", providerName, err)
	} else {
		counts.loadBalancers += len(lbs)
		for _, lb := range lbs {
			counts.lbTypeSummary[strings.ToLower(lb.Type)]++
			if strings.Contains(lb.Scheme, "internet-facing") {
				counts.publiclyAccessible++
			}
		}
	}

	// Network Interfaces
	nics, err := cloudProvider.ListNetworkInterfaces(ctx, &infrapb.ListNetworkInterfacesRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err != nil {
		s.logger.Warnf("Failed to list network interfaces for %s: %v", providerName, err)
	} else {
		counts.networkInterfaces += len(nics)
	}

	// Key Pairs
	keyPairs, err := cloudProvider.ListKeyPairs(ctx, &infrapb.ListKeyPairsRequest{Provider: providerName, AccountId: accountID, Region: in.Region})
	if err != nil {
		s.logger.Warnf("Failed to list key pairs for %s: %v", providerName, err)
	} else {
		counts.keyPairs += len(keyPairs)
	}

	// Other resources
	s.processOtherResources(ctx, cloudProvider, providerName, accountID, in, counts)
}

func (s *Server) processOtherResources(ctx context.Context, cloudProvider provider.CloudProvider, providerName, accountID string, in *infrapb.SummaryRequest, counts *resourceCounts) {
	// ACLs
	acls, err := cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.acls += len(acls)
	}

	// Route Tables
	routeTables, err := cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.routeTables += len(routeTables)
	}

	// NAT Gateways
	natGateways, err := cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.natGateways += len(natGateways)
	}

	// Internet Gateways
	igws, err := cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.igws += len(igws)
	}

	// Routers
	routers, err := cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.routers += len(routers)
	}

	// VPC Endpoints
	vpcEndpoints, err := cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.vpcEndpoints += len(vpcEndpoints)
	}

	// Public IPs
	publicIPs, err := cloudProvider.ListPublicIPs(ctx, &infrapb.ListPublicIPsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.publicIPs += len(publicIPs)
	}

	// VPN Concentrators
	vpnConcentrators, err := cloudProvider.ListVPNConcentrators(ctx, &infrapb.ListVPNConcentratorsRequest{Provider: providerName, AccountId: accountID, Region: in.Region})
	if err == nil {
		counts.vpnConcentrators += len(vpnConcentrators)
	}

	// VPC Connections
	vpcConnections, err := cloudProvider.ListVpcConnections(ctx, &infrapb.ListVpcConnectionsRequest{Provider: providerName, AccountId: accountID, Region: in.Region})
	if err == nil {
		counts.vpcConnections += len(vpcConnections)
		// Count isolated vs connected VPCs
		if len(vpcConnections) == 0 {
			counts.isolatedVpcs++
		} else {
			counts.connectedVpcs++
		}
	}

	// Kubernetes clusters
	clusters, err := cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
	if err == nil {
		counts.clusters += len(clusters)
		s.processKubernetesResources(ctx, clusters, counts)
	}
}

func (s *Server) processKubernetesResources(ctx context.Context, clusters []types.Cluster, counts *resourceCounts) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil || k8sProvider == nil {
		return
	}

	relevantClusters := make(map[string]struct{})
	for _, c := range clusters {
		relevantClusters[c.Name] = struct{}{}
	}

	if len(relevantClusters) > 0 {
		// Pods
		pods, err := k8sProvider.ListPods(ctx, "", nil)
		if err == nil {
			for _, pod := range pods {
				if _, ok := relevantClusters[pod.Cluster]; ok {
					state := strings.ToLower(pod.State)
					counts.podsStateSummary[state]++
					counts.pods++
				}
			}
		}

		// Services
		services, err := k8sProvider.ListServices(ctx, "", nil)
		if err == nil {
			for _, serv := range services {
				if _, ok := relevantClusters[serv.Cluster]; ok {
					counts.services++
				}
			}
		}

		// Namespaces
		namespaces, err := k8sProvider.ListNamespaces(ctx, "", nil)
		if err == nil {
			for _, namespace := range namespaces {
				if _, ok := relevantClusters[namespace.Cluster]; ok {
					counts.namespaces++
				}
			}
		}
	}
}

func (s *Server) aggregateCounts(aggregated, counts *resourceCounts) {
	aggregated.accounts += counts.accounts
	aggregated.vpcs += counts.vpcs
	aggregated.subnets += counts.subnets
	aggregated.instances += counts.instances
	aggregated.acls += counts.acls
	aggregated.sgs += counts.sgs
	aggregated.routeTables += counts.routeTables
	aggregated.natGateways += counts.natGateways
	aggregated.routers += counts.routers
	aggregated.igws += counts.igws
	aggregated.vpcEndpoints += counts.vpcEndpoints
	aggregated.publicIPs += counts.publicIPs
	aggregated.clusters += counts.clusters
	aggregated.pods += counts.pods
	aggregated.services += counts.services
	aggregated.namespaces += counts.namespaces
	aggregated.loadBalancers += counts.loadBalancers
	aggregated.networkInterfaces += counts.networkInterfaces
	aggregated.keyPairs += counts.keyPairs
	aggregated.vpcConnections += counts.vpcConnections
	aggregated.vpnConcentrators += counts.vpnConcentrators
	aggregated.untaggedResources += counts.untaggedResources
	aggregated.publiclyAccessible += counts.publiclyAccessible
	aggregated.openSecurityGroups += counts.openSecurityGroups
	aggregated.isolatedVpcs += counts.isolatedVpcs
	aggregated.connectedVpcs += counts.connectedVpcs

	// Merge maps
	for k, v := range counts.vmStateSummary {
		aggregated.vmStateSummary[k] += v
	}
	for k, v := range counts.vmTypeSummary {
		aggregated.vmTypeSummary[k] += v
	}
	for k, v := range counts.podsStateSummary {
		aggregated.podsStateSummary[k] += v
	}
	for k, v := range counts.regionSummary {
		aggregated.regionSummary[k] += v
	}
	for k, v := range counts.lbTypeSummary {
		aggregated.lbTypeSummary[k] += v
	}
	for k, v := range counts.healthSummary {
		aggregated.healthSummary[k] += v
	}
}

func (s *Server) Summary(ctx context.Context, in *infrapb.SummaryRequest) (*infrapb.SummaryResponse, error) {
	// Check cache first
	if cached := s.getCachedSummary(in); cached != nil {
		return cached, nil
	}

	// Use context with timeout
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var providers []provider.CloudProvider
	if in.Provider != "" && in.Provider != "all" {
		provider, err := s.strategy.GetProvider(ctx, in.Provider)
		if err != nil {
			return nil, fmt.Errorf("error retrieving provider %s: %w", in.Provider, err)
		}
		providers = append(providers, provider)
	} else {
		providers = s.strategy.GetAllProviders()
	}

	// Process providers in parallel
	var wg sync.WaitGroup
	resultsChan := make(chan *resourceCounts, len(providers))

	for _, cloudProvider := range providers {
		wg.Add(1)
		go func(cp provider.CloudProvider) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					s.logger.Errorf("Panic in provider summary processing: %v", r)
				}
			}()
			counts := s.processProviderSummary(ctx, cp, in)
			resultsChan <- counts
		}(cloudProvider)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Aggregate results
	aggregated := &resourceCounts{
		vmStateSummary:   make(map[string]int32),
		vmTypeSummary:    make(map[string]int32),
		podsStateSummary: make(map[string]int32),
		regionSummary:    make(map[string]int32),
		lbTypeSummary:    make(map[string]int32),
		healthSummary:    make(map[string]int32),
	}

	for counts := range resultsChan {
		s.aggregateCounts(aggregated, counts)
	}

	// Build response
	summary := &infrapb.SummaryResponse{
		Count: &infrapb.Counters{
			Accounts:          int32(aggregated.accounts),
			Vpc:               int32(aggregated.vpcs),
			Subnets:           int32(aggregated.subnets),
			RouteTables:       int32(aggregated.routeTables),
			Instances:         int32(aggregated.instances),
			Acls:              int32(aggregated.acls),
			SecurityGroups:    int32(aggregated.sgs),
			NatGateways:       int32(aggregated.natGateways),
			Routers:           int32(aggregated.routers),
			Igws:              int32(aggregated.igws),
			VpcEndpoints:      int32(aggregated.vpcEndpoints),
			PublicIps:         int32(aggregated.publicIPs),
			InternetGateways:  int32(aggregated.igws),
			Clusters:          int32(aggregated.clusters),
			Pods:              int32(aggregated.pods),
			Services:          int32(aggregated.services),
			Namespaces:        int32(aggregated.namespaces),
			LoadBalancers:     int32(aggregated.loadBalancers),
			NetworkInterfaces: int32(aggregated.networkInterfaces),
			KeyPairs:          int32(aggregated.keyPairs),
			VpcConnections:    int32(aggregated.vpcConnections),
			VpnConcentrators:  int32(aggregated.vpnConcentrators),
		},
		Statuses: &infrapb.StatusSummary{
			VmStatus:           aggregated.vmStateSummary,
			PodStatus:          aggregated.podsStateSummary,
			VmTypes:            aggregated.vmTypeSummary,
			LbTypes:            aggregated.lbTypeSummary,
			RegionDistribution: aggregated.regionSummary,
			HealthStatus:       aggregated.healthSummary,
		},
		Security: &infrapb.SecuritySummary{
			OpenSecurityGroups:          int32(aggregated.openSecurityGroups),
			UntaggedResources:           int32(aggregated.untaggedResources),
			PubliclyAccessibleResources: int32(aggregated.publiclyAccessible),
		},
		NetworkTopology: &infrapb.NetworkTopologySummary{
			IsolatedVpcs:       int32(aggregated.isolatedVpcs),
			ConnectedVpcs:      int32(aggregated.connectedVpcs),
			PeeringConnections: int32(aggregated.vpcConnections),
		},
		Health: &infrapb.ResourceHealthSummary{
			HealthyResources:   aggregated.healthSummary["healthy"],
			DegradedResources:  aggregated.healthSummary["degraded"],
			UnhealthyResources: aggregated.healthSummary["unhealthy"],
			StoppedResources:   aggregated.healthSummary["stopped"],
		},
	}

	// Cache the result
	s.cacheSummary(in, summary)

	return summary, nil
}
