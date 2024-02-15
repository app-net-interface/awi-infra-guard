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
	"time"

	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

func (c *Client) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	net1, err := c.findNetwork(ctx, "", input.Vpc1ID)
	if err != nil {
		return types.VPCConnectionOutput{}, err
	}
	if len(net1) == 0 {
		return types.VPCConnectionOutput{}, fmt.Errorf("VPC %s not found", input.Vpc1ID)
	}
	net2, err := c.findNetwork(ctx, "", input.Vpc2ID)
	if err != nil {
		return types.VPCConnectionOutput{}, err
	}
	if len(net2) == 0 {
		return types.VPCConnectionOutput{}, fmt.Errorf("VPC %s not found", input.Vpc2ID)
	}
	if net1[0].project != net2[0].project {
		return types.VPCConnectionOutput{}, fmt.Errorf("VPC connection must be done within same project")
	}
	err = c.createVPCPeering(ctx, net1[0].project, net1[0], net2[0])
	if err != nil {
		return types.VPCConnectionOutput{}, err
	}
	err = c.createVPCPeering(ctx, net1[0].project, net2[0], net1[0])
	if err != nil {
		return types.VPCConnectionOutput{}, err
	}
	return types.VPCConnectionOutput{
		Region1: "global",
		Region2: "global",
	}, nil
}

func (c *Client) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {
	return types.SingleVPCConnectionOutput{
		Region: "global",
	}, nil
}

func (c *Client) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	net1, err := c.findNetwork(ctx, "", input.Vpc1ID)
	if err != nil {
		return types.VPCDisconnectionOutput{}, err
	}
	if len(net1) == 0 {
		return types.VPCDisconnectionOutput{}, fmt.Errorf("VPC %s not found", input.Vpc1ID)
	}
	net2, err := c.findNetwork(ctx, "", input.Vpc2ID)
	if err != nil {
		return types.VPCDisconnectionOutput{}, err
	}
	if len(net2) == 0 {
		return types.VPCDisconnectionOutput{}, fmt.Errorf("VPC %s not found", input.Vpc2ID)
	}
	if net1[0].project != net2[0].project {
		return types.VPCDisconnectionOutput{}, fmt.Errorf("VPC connection must be done within same project")
	}
	err = c.deleteVPCPeering(ctx, net1[0].project, net1[0], net2[0])
	if err != nil {
		return types.VPCDisconnectionOutput{}, err
	}
	err = c.deleteVPCPeering(ctx, net1[0].project, net2[0], net1[0])
	if err != nil {
		return types.VPCDisconnectionOutput{}, err
	}
	return types.VPCDisconnectionOutput{}, nil
}

func (c *Client) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	return types.VPCDisconnectionOutput{}, nil
}

func peeringName(network1, network2 network) string {
	return network1.name + "-to-" + network2.name
}

func (c *Client) createVPCPeering(ctx context.Context, projectID string, network1, network2 network) error {

	// Define the peering configuration
	peering := &compute.NetworksAddPeeringRequest{
		NetworkPeering: &compute.NetworkPeering{
			Name:                           peeringName(network1, network2),
			Network:                        network2.url,
			ExportCustomRoutes:             true,
			ExchangeSubnetRoutes:           true,
			ExportSubnetRoutesWithPublicIp: true,
		},
	}

	// Create the VPC peering
	operation, err := c.computeService.Networks.AddPeering(projectID, network1.name, peering).Context(ctx).Do()
	if err != nil {
		if strings.Contains(err.Error(), "There is already a peering") {
			return nil
		}
		return err
	}

	// Wait for the operation to complete (optional)
	if err := waitForOperation(ctx, c.computeService, projectID, operation.Name); err != nil {
		return err
	}

	c.logger.Infof("VPC peering created successfully between %s and %s\n", network1.name, network2.name)
	return nil
}

// Helper function to wait for the operation to complete
func waitForOperation(ctx context.Context, service *compute.Service, projectID, operationName string) error {
	for {
		operation, err := service.GlobalOperations.Get(projectID, operationName).Context(ctx).Do()
		if err != nil {
			return err
		}

		if operation.Status == "DONE" {
			if operation.Error != nil {
				return fmt.Errorf("operation failed: %v", operation.Error.Errors)
			}
			return nil
		}

		// Sleep for a while before checking the operation status again
		time.Sleep(time.Second * 5)
	}
}

func (c *Client) deleteVPCPeering(ctx context.Context, projectID string, net1, net2 network) error {
	peering := peeringName(net1, net2) // Delete the VPC peering
	operation, err := c.computeService.Networks.RemovePeering(projectID, net1.name, &compute.NetworksRemovePeeringRequest{
		Name: peering,
	}).Context(ctx).Do()

	if err != nil {
		if strings.Contains(err.Error(), "Error 400: There is no peering") {
			c.logger.Infof("VPC peering '%s' doesn't exist in network '%s'\n", peering, net1.name)
			return nil
		}
		return err
	}

	// Wait for the operation to complete (optional)
	if err := waitForOperation(ctx, c.computeService, projectID, operation.Name); err != nil {
		return err
	}

	c.logger.Infof("VPC peering '%s' deleted successfully from network '%s'\n", peering, net1.name)
	return nil
}
