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

package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

const providerName = "Azure"

type Client struct {
	cred       *azidentity.DefaultAzureCredential
	logger     *logrus.Logger
	vnetClient **armnetwork.VirtualNetworksClient
}

// NewClient initializes a new Azure client with all necessary clients for compute, network, and subscriptions.
func NewClient(ctx context.Context, logger *logrus.Logger) (*Client, error) {
	// Subscription ID from environment variable
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to obtain a credential:", err)
		return nil, err
	}
	client := &Client{
		cred:   cred,
		logger: logger,
	}

	return client, nil
}

func (c *Client) GetName() string {
	return providerName
}

func (c *Client) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}

func (c *Client) GetSubnet(ctx context.Context, input *infrapb.GetSubnetRequest) (types.Subnet, error) {
	// TBD
	return types.Subnet{}, nil
}

func (c *Client) GetVPCIDForCIDR(ctx context.Context, input *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) GetCIDRsForLabels(ctx context.Context, input *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetIPsForLabels(ctx context.Context, input *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetInstancesForLabels(ctx context.Context, input *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetVPCIDWithTag(ctx context.Context, input *infrapb.GetVPCIDWithTagRequest) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {

	return nil, nil
}

func (c *Client) GetVPCIndex(ctx context.Context, vpcIndex *infrapb.GetVPCIndexRequest) (*types.VPCIndex, error) {
	// This logic is handled by the DB strategy.
	return nil, fmt.Errorf("GetVPCIndex not implemented directly in Azure client; use DB strategy")
}

// Add placeholder implementation for ListVpcGraphNodes
func (c *Client) ListVpcGraphNodes(ctx context.Context, params *infrapb.ListVpcGraphNodesRequest) ([]types.VpcGraphNode, error) {
	// This logic is handled by the DB strategy, which builds nodes from existing DB data.
	// The real provider doesn't need a direct implementation unless bypassing the DB.
	return nil, fmt.Errorf("ListVpcGraphNodes not implemented directly in Azure client; use DB strategy")
}

// Add placeholder implementation for ListVpcGraphEdges
func (c *Client) ListVpcGraphEdges(ctx context.Context, params *infrapb.ListVpcGraphEdgesRequest) ([]types.VpcGraphEdge, error) {
	// This logic is handled by the DB strategy, which builds edges from existing DB data.
	// The real provider doesn't need a direct implementation unless bypassing the DB.
	return nil, fmt.Errorf("ListVpcGraphEdges not implemented directly in Azure client; use DB strategy")
}

// Update placeholder implementation for GetVpcConnectivityGraph
func (c *Client) GetVpcConnectivityGraph(ctx context.Context, params *infrapb.GetVpcConnectivityGraphRequest) ([]types.VpcGraphNode, []types.VpcGraphEdge, error) {
	// This logic is handled by the DB strategy.
	return nil, nil, fmt.Errorf("GetVpcConnectivityGraph not implemented directly in Azure client; use DB strategy")
}

// Update placeholder implementation for GetVpcConnectionGraph
func (c *Client) GetVpcConnectionGraph(ctx context.Context, params *infrapb.GetVpcConnectionGraphRequest) (*types.VpcConnectionGraph, error) {
	// This logic is handled by the DB strategy.
	return nil, fmt.Errorf("GetVpcConnectionGraph not implemented directly in AWS client; use DB strategy")
}

// GetInstanceConnectivityGraph is a placeholder implementation.
func (c *Client) GetInstanceConnectivityGraph(ctx context.Context, params *infrapb.GetInstanceConnectivityGraphRequest) ([]types.InstanceGraphNode, []types.InstanceGraphEdge, error) {
	c.logger.Infof("GetInstanceConnectivityGraph called for Azure VM %s (Not Implemented)", params.InstanceId)
	// TODO: Implement logic to fetch VM, NIC, Subnet, Route Table, NSGs
	// and build the nodes and edges specific to Azure resources.
	// Note: Azure instance ID might need parsing to get resource group/name.
	return nil, nil, fmt.Errorf("GetInstanceConnectivityGraph not implemented for Azure provider")
}

// ListVpcConnections implements provider.CloudProvider interface
func (c *Client) ListVpcConnections(ctx context.Context, input *infrapb.ListVpcConnectionsRequest) ([]types.VPCConnection, error) {
	vpcConns := []types.VPCConnection{}

	if c.vnetClient == nil {
		return nil, fmt.Errorf("vnet client not initialized")
	}

	// List all VNets
	pager := (*c.vnetClient).NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list VNets: %v", err)
		}

		for _, vnet := range nextResult.Value {
			if vnet.Properties == nil || vnet.Properties.VirtualNetworkPeerings == nil {
				continue
			}

			// Each VNet can have multiple peering connections
			for _, peering := range vnet.Properties.VirtualNetworkPeerings {
				if peering.Properties == nil || peering.Properties.RemoteVirtualNetwork == nil {
					continue
				}

				vpcConn := types.VPCConnection{
					Provider: providerName,
					ID:       *peering.ID,
					Name:     *peering.Name,
					// We'll get subscription ID from the resource ID
					Account:   "TODO", // Need to extract subscription from resource ID
					Region:    *vnet.Location,
					FromVpcId: *vnet.ID,
					ToVpcId:   *peering.Properties.RemoteVirtualNetwork.ID,
					// Need to parse subscription IDs from resource IDs
					FromVpcAccountId: "TODO", // Need to extract from vnet.ID
					ToVpcAccountId:   "TODO", // Need to extract from remoteVnet ID
					FromVpcRegion:    *vnet.Location,
					ToVpcRegion:      "TODO", // Need to get from remote VNet
					Status:           string(*peering.Properties.PeeringState),
				}

				vpcConns = append(vpcConns, vpcConn)
			}
		}
	}

	return vpcConns, nil
}
