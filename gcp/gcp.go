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

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
	computeServ "google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/iterator"
)

const providerName = "GCP"

type Client struct {
	instancesClient  *compute.InstancesClient
	subnetsClient    *compute.SubnetworksClient
	networksClient   *compute.NetworksClient
	computeService   *computeServ.Service
	containerService *container.Service
	logger           *logrus.Logger
	projectIDs       map[string]struct{}
}

func NewClient(ctx context.Context, logger *logrus.Logger) (*Client, error) {
	projectIDs, err := findProjects(ctx, logger)
	if err != nil {
		return nil, err
	}
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	subnetsClient, err := compute.NewSubnetworksRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	networksClient, err := compute.NewNetworksRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	computeService, err := computeServ.NewService(ctx)
	if err != nil {
		return nil, err
	}
	containerService, err := container.NewService(ctx)
	if err != nil {
		return nil, err
	}
	projectsMap := make(map[string]struct{}, len(projectIDs))
	for _, v := range projectIDs {
		projectsMap[v] = struct{}{}
	}
	return &Client{
		instancesClient:  instancesClient,
		subnetsClient:    subnetsClient,
		networksClient:   networksClient,
		computeService:   computeService,
		containerService: containerService,
		projectIDs:       projectsMap,
		logger:           logger,
	}, nil
}

func findProjects(ctx context.Context, logger *logrus.Logger) ([]string, error) {
	projectsClient, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return nil, err
	}

	var projectIDs []string
	iter := projectsClient.SearchProjects(ctx, &resourcemanagerpb.SearchProjectsRequest{})
	for {
		project, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		logger.Infof("Found GCP project %s (ID %s)", project.GetDisplayName(), project.GetProjectId())
		networksClient, err := compute.NewNetworksRESTClient(ctx)
		if err != nil {
			return nil, err
		}
		iter := networksClient.List(ctx, &computepb.ListNetworksRequest{
			Project: project.GetName(),
		})

		addProject := true
		_, err = iter.Next()
		if err != nil && err != iterator.Done {
			logger.Errorf("GCP Project %s will be ignored because compute API operation failed with error: %v",
				project.GetDisplayName(), err)
			addProject = false
		}
		if addProject {
			projectIDs = append(projectIDs, project.GetProjectId())
		}
	}

	return projectIDs, nil
}

func (c *Client) GetName() string {
	return providerName
}

func (c *Client) ListAccounts() []types.Account {
	accounts := make([]types.Account, 0, len(c.projectIDs))
	for k := range c.projectIDs {
		accounts = append(accounts, types.Account{
			ID: k,
		})
	}
	return accounts
}

func (c *Client) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}

func (c *Client) checkProject(project string) error {
	_, ok := c.projectIDs[project]
	if !ok {
		return fmt.Errorf("provided project: %s not configured", project)
	}
	return nil
}

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {

	return nil, nil
}

func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {

	return nil, nil
}

func (c *Client) GetVPCIndex(ctx context.Context, vpcIndex *infrapb.GetVPCIndexRequest) (*types.VPCIndex, error) {
	// This logic is handled by the DB strategy.
	return nil, fmt.Errorf("GetVPCIndex not implemented directly in GCP client; use DB strategy")
}

// Add placeholder implementation for ListVpcGraphNodes
func (c *Client) ListVpcGraphNodes(ctx context.Context, params *infrapb.ListVpcGraphNodesRequest) ([]types.VpcGraphNode, error) {
	// This logic is handled by the DB strategy, which builds nodes from existing DB data.
	// The real provider doesn't need a direct implementation unless bypassing the DB.
	return nil, fmt.Errorf("ListVpcGraphNodes not implemented directly in GCP client; use DB strategy")
}

// Add placeholder implementation for ListVpcGraphEdges
func (c *Client) ListVpcGraphEdges(ctx context.Context, params *infrapb.ListVpcGraphEdgesRequest) ([]types.VpcGraphEdge, error) {
	// This logic is handled by the DB strategy, which builds edges from existing DB data.
	// The real provider doesn't need a direct implementation unless bypassing the DB.
	return nil, fmt.Errorf("ListVpcGraphEdges not implemented directly in GCP client; use DB strategy")
}

// Update placeholder implementation for GetVpcConnectivityGraph
func (c *Client) GetVpcConnectivityGraph(ctx context.Context, params *infrapb.GetVpcConnectivityGraphRequest) ([]types.VpcGraphNode, []types.VpcGraphEdge, error) {
	// This logic is handled by the DB strategy.
	return nil, nil, fmt.Errorf("GetVpcConnectivityGraph not implemented directly in GCP client; use DB strategy")
}
