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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
)

func (s *Server) Summary(ctx context.Context, in *infrapb.SummaryRequest) (*infrapb.SummaryResponse, error) {
	var providers []provider.CloudProvider
	if in.Provider != "" && in.Provider != "all" {
		provider, err := s.strategy.GetProvider(ctx, in.Provider)
		if err != nil {
			return nil, fmt.Errorf("error retrieving provider %s: %w", in.Provider, err)
		}
		providers = append(providers, provider)
	} else {
		// If no specific provider, get all available providers
		providers = s.strategy.GetAllProviders()
	}

	// Initialize aggregate counters and summaries
	totalAccounts := 0
	totalVPCs := 0
	totalSubnets := 0
	totalInstances := 0
	totalACLs := 0
	totalSGs := 0
	totalRouteTables := 0
	totalNatGateways := 0
	totalRouters := 0
	totalIGWs := 0
	totalVpcEndpoints := 0
	totalPublicIPs := 0
	totalClusters := 0
	totalPods := 0
	totalServices := 0
	totalNamespaces := 0

	vmStateSummary := make(map[string]int32)
	vmTypeSummary := make(map[string]int32)
	podsStateSummary := make(map[string]int32)

	k8sProvider, _ := s.strategy.GetKubernetesProvider()

	for _, cloudProvider := range providers {
		providerName := cloudProvider.GetName()
		var accountIDs []string
		if in.AccountId != "" {
			accountIDs = append(accountIDs, in.AccountId)
		} else {
			for _, acc := range cloudProvider.ListAccounts() {
				accountIDs = append(accountIDs, acc.ID)
			}
		}
		totalAccounts += len(accountIDs)

		for _, accountID := range accountIDs {
			vpcs, _ := cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{Provider: providerName, AccountId: accountID, Region: in.Region})
			totalVPCs += len(vpcs)

			subnets, _ := cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalSubnets += len(subnets)

			instances, _ := cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalInstances += len(instances)
			for _, vm := range instances {
				vmStateSummary[strings.ToLower(vm.State)]++
				vmTypeSummary[strings.ToLower(vm.Type)]++
			}

			acls, _ := cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalACLs += len(acls)

			sgs, _ := cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalSGs += len(sgs)

			routeTables, _ := cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalRouteTables += len(routeTables)

			natGateways, _ := cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalNatGateways += len(natGateways)

			igws, _ := cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalIGWs += len(igws)

			routers, _ := cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalRouters += len(routers)

			vpcEndpoints, _ := cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalVpcEndpoints += len(vpcEndpoints)

			publicIPs, _ := cloudProvider.ListPublicIPs(ctx, &infrapb.ListPublicIPsRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalPublicIPs += len(publicIPs)

			// Kubernetes Resources
			clusters, _ := cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{Provider: providerName, AccountId: accountID, Region: in.Region, VpcId: in.VpcId})
			totalClusters += len(clusters)

			if k8sProvider != nil {
				relevantClusters := make(map[string]struct{})
				for _, c := range clusters {
					relevantClusters[c.Name] = struct{}{}
				}
				if len(relevantClusters) > 0 {
					pods, _ := k8sProvider.ListPods(ctx, "", nil)
					for _, pod := range pods {
						if _, ok := relevantClusters[pod.Cluster]; ok {
							podsStateSummary[strings.ToLower(pod.State)]++
							totalPods++
						}
					}
					services, _ := k8sProvider.ListServices(ctx, "", nil)
					for _, serv := range services {
						if _, ok := relevantClusters[serv.Cluster]; ok {
							totalServices++
						}
					}
					namespaces, _ := k8sProvider.ListNamespaces(ctx, "", nil)
					for _, namespace := range namespaces {
						if _, ok := relevantClusters[namespace.Cluster]; ok {
							totalNamespaces++
						}
					}
				}
			}
		}
	}

	summary := &infrapb.SummaryResponse{
		Count: &infrapb.Counters{
			Accounts:       int32(totalAccounts),
			Vpc:            int32(totalVPCs),
			Subnets:        int32(totalSubnets),
			RouteTables:    int32(totalRouteTables),
			Instances:      int32(totalInstances),
			Acls:           int32(totalACLs),
			SecurityGroups: int32(totalSGs),
			NatGateways:    int32(totalNatGateways),
			Routers:        int32(totalRouters),
			Igws:           int32(totalIGWs),
			VpcEndpoints:   int32(totalVpcEndpoints),
			PublicIps:      int32(totalPublicIPs),
			Clusters:       int32(totalClusters),
			Pods:           int32(totalPods),
			Services:       int32(totalServices),
			Namespaces:     int32(totalNamespaces),
		},
		Statuses: &infrapb.StatusSummary{
			VmStatus:  vmStateSummary,
			PodStatus: podsStateSummary,
			VmTypes:   vmTypeSummary,
		},
	}
	return summary, nil
}
