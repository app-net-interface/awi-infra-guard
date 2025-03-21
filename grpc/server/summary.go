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
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (s *Server) Summary(ctx context.Context, in *infrapb.SummaryRequest) (*infrapb.SummaryResponse, error) {
	var accounts []types.Account
	var vpcs []types.VPC

	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		fmt.Printf("Error retreiving provider %s ", err.Error())
		return nil, err
	}
	if in.AccountId == "" {
		accounts = cloudProvider.ListAccounts()
	}
	if in.VpcId == "" {
		vpcs, err = cloudProvider.ListVPC(ctx, &infrapb.ListVPCRequest{
			Provider:  in.Provider,
			AccountId: in.AccountId,
			Region:    in.Region,
		})
		if err != nil {
			fmt.Printf("Error retreiving vpcs %s ", err.Error())
			return nil, err
		}
	}
	subnets, err := cloudProvider.ListSubnets(ctx, &infrapb.ListSubnetsRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving subnets %s ", err.Error())
		return nil, err
	}
	instances, err := cloudProvider.ListInstances(ctx, &infrapb.ListInstancesRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving instances %s ", err.Error())
		return nil, err
	}
	vmStateSummary := make(map[string]int32)
	for _, vm := range instances {
		vmStateSummary[strings.ToLower(vm.State)] += 1
	}
	vmTypeSummary := make(map[string]int32)
	for _, vm := range instances {
		vmTypeSummary[strings.ToLower(vm.Type)] += 1
	}
	acls, err := cloudProvider.ListACLs(ctx, &infrapb.ListACLsRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving acls %s ", err.Error())
		return nil, err
	}
	sgs, err := cloudProvider.ListSecurityGroups(ctx, &infrapb.ListSecurityGroupsRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving security groups %s ", err.Error())
		return nil, err
	}
	routeTables, err := cloudProvider.ListRouteTables(ctx, &infrapb.ListRouteTablesRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving route tables %s ", err.Error())
		return nil, err
	}

	natGateways, err := cloudProvider.ListNATGateways(ctx, &infrapb.ListNATGatewaysRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving nat gateways %s ", err.Error())
		return nil, err
	}

	igws, err := cloudProvider.ListInternetGateways(ctx, &infrapb.ListInternetGatewaysRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving internet gateways %s ", err.Error())
		return nil, err
	}

	routers, err := cloudProvider.ListRouters(ctx, &infrapb.ListRoutersRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving routers %s ", err.Error())
		return nil, err
	}

	vpcEndpoints, err := cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving vpc endpoints %s ", err.Error())
		return nil, err
	}
	publicIPs, err := cloudProvider.ListPublicIPs(ctx, &infrapb.ListPublicIPsRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving public ips %s ", err.Error())
		return nil, err
	}

	// Kubernetes Resources

	clusters, err := cloudProvider.ListClusters(ctx, &infrapb.ListCloudClustersRequest{
		Provider:  in.Provider,
		AccountId: in.AccountId,
		Region:    in.Region,
		VpcId:     in.VpcId,
	})
	if err != nil {
		fmt.Printf("Error retreiving clusters %s ", err.Error())
		//return nil, err
	}
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		fmt.Printf("Error retreiving provider %s ", err.Error())
		//return nil, err
	}
	var podsCount int
	pods, err := k8sProvider.ListPods(ctx, "", nil)
	if err != nil {
		fmt.Printf("Error retreiving pods %s ", err.Error())
		//return nil, err
	}
	podsStateSummary := make(map[string]int32)
	for _, pod := range pods {
		podsStateSummary[strings.ToLower(pod.State)] += 1
		for _, cl := range clusters {
			if pod.Cluster == cl.Name {
				podsCount++
			}
		}
	}
	var servicesCount int
	services, err := k8sProvider.ListServices(ctx, "", nil)
	if err != nil {
		fmt.Printf("Error retreiving services %s ", err.Error())

		//return nil, err
	}
	for _, serv := range services {
		for _, cl := range clusters {
			if serv.Cluster == cl.Name {
				servicesCount++
			}
		}
	}
	var namespacesCount int
	namespaces, err := k8sProvider.ListNamespaces(ctx, "", nil)
	if err != nil {
		fmt.Printf("Error retreiving namespaces %s ", err.Error())

		//return nil, err
	}
	for _, namespace := range namespaces {
		for _, cl := range clusters {
			if namespace.Cluster == cl.Name {
				namespacesCount++
			}
		}
	}

	summary := &infrapb.SummaryResponse{
		Count: &infrapb.Counters{
			Accounts:       int32(len(accounts)),
			Vpc:            int32(len(vpcs)),
			Subnets:        int32(len(subnets)),
			RouteTables:    int32(len(routeTables)),
			Instances:      int32(len(instances)),
			Acls:           int32(len(acls)),
			SecurityGroups: int32(len(sgs)),
			NatGateways:    int32(len(natGateways)),
			Routers:        int32(len(routers)),
			Igws:           int32(len(igws)),
			VpcEndpoints:   int32(len(vpcEndpoints)),
			PublicIps:      int32(len(publicIPs)),

			//Kubernetes
			Clusters:   int32(len(clusters)),
			Pods:       int32(podsCount),
			Services:   int32(servicesCount),
			Namespaces: int32(namespacesCount),
		},
		Statuses: &infrapb.StatusSummary{
			VmStatus:  vmStateSummary,
			PodStatus: podsStateSummary,
			VmTypes:   vmTypeSummary,
		},
	}
	fmt.Printf(" ************* Summary: %v\n *************", summary)

	return summary, nil
}
