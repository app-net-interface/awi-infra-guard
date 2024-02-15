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

package gcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/app-net-interface/kubernetes-discovery/cluster"
	gkediscovery "github.com/app-net-interface/kubernetes-discovery/cluster/gke"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
)

func (c *Client) ListClusters(ctx context.Context, params *infrapb.ListCloudClustersRequest) ([]types.Cluster, error) {
	var net network
	if params.GetVpcId() != "" {
		nets, err := c.findNetwork(ctx, params.GetAccountId(), params.GetVpcId())
		if err != nil {
			return nil, err
		}
		if len(nets) != 1 {
			return nil, fmt.Errorf("found %d matching networks, expected 1", len(nets))
		}
		net = nets[0]
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	clusters := make([]types.Cluster, 0)
	f := func(projectId string) error {
		resp, err := c.containerService.Projects.Zones.Clusters.List(projectId, "-").Do()
		if err != nil {
			return err
		}
	clustersLoop:
		for _, gcpCluster := range resp.Clusters {
			cluster := gcpClusterToTypes(projectId, networks, gcpCluster)
			if params.GetVpcId() != "" && (cluster.VpcID != net.name && cluster.VpcID != net.id) {
				continue clustersLoop
			}
			for k, v := range params.GetLabels() {
				if cluster.Labels[k] != v {
					continue clustersLoop
				}
			}
			clusters = append(clusters, cluster)
		}
		return nil
	}
	if params.GetAccountId() == "" {
		for project := range c.projectIDs {
			err := f(project)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}
	return clusters, nil
}

func gcpClusterToTypes(project string, networks []types.VPC, cluster *container.Cluster) types.Cluster {
	vpcID := cluster.Network
	network := strings.Split(cluster.Network, "/")
	if len(network) != 0 {
		name := network[len(network)-1]
		for _, v := range networks {
			if v.Name == name || v.ID == name {
				vpcID = v.ID
				break
			}
		}
	}
	return types.Cluster{
		Name:      cluster.Name,
		FullName:  "gke_" + project + "_" + cluster.Zone + "_" + cluster.Name,
		VpcID:     vpcID,
		Labels:    cluster.ResourceLabels,
		Region:    cluster.Zone,
		Project:   project,
		AccountID: project,
		Provider:  providerName,
		Arn:       cluster.SelfLink,
		Id:        cluster.Id,
	}
}

func (c *Client) RetrieveClustersData(ctx context.Context) ([]cluster.DiscoveredCluster, error) {
	creds, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	var clusters []cluster.DiscoveredCluster
	for project := range c.projectIDs {
		creds.ProjectID = project
		gkeDiscovery, err := gkediscovery.NewClustersRetriever(creds)
		if err != nil {
			return nil, err
		}
		cls, err := gkeDiscovery.Retrieve(ctx)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, cls.DiscoveredClusters...)
	}
	return clusters, nil
}
