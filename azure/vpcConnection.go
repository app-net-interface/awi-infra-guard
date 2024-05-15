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

package azure

import (
	"context"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/types"
)

// VPCConnector interface implementation
func (c *Client) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {

	return types.SingleVPCConnectionOutput{}, nil
}

func (c *Client) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	input.Region1 = "westus2"
	input.Region2 = "westus2"

	vnet1, accountID1, err := c.getVPC(ctx, input.Vpc1ID, input.Region1)
	if err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc1ID, input.Region1,
		)
	}
	vnet2, accountID2, err := c.getVPC(ctx, input.Vpc2ID, input.Region2)
	if err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc2ID, input.Region2,
		)
	}

	if err = c.BlockTrafficBetweenVPCs(ctx, vnet1, vnet2, accountID1, accountID2); err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to create a blocking rule between %s:%s and %s:%s due to %w",
			input.Region1, *vnet1.ID, input.Region2, *vnet2.ID, err,
		)
	}

	if err = c.createVnetPeering(ctx, *vnet1.ID, *vnet2.ID, accountID1); err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to create a VPC Peering from %s:%s to %s:%s due to %w",
			input.Region1, *vnet1.ID, input.Region2, *vnet2.ID, err,
		)
	}

	if err = c.createVnetPeering(ctx, *vnet2.ID, *vnet1.ID, accountID2); err != nil {
		return types.VPCConnectionOutput{}, fmt.Errorf(
			"failed to create a VPC Peering from %s:%s to %s:%s due to %w",
			input.Region2, *vnet2.ID, input.Region1, *vnet1.ID, err,
		)
	}

	return types.VPCConnectionOutput{
		Region1: input.Region1,
		Region2: input.Region2,
	}, nil
}

func (c *Client) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}

// func getNSGNameForPeeredVPCs(sourceVNET, destinationVNET string) string {

// }

func (c *Client) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	vnet1, accountID1, err := c.getVPC(ctx, input.Vpc1ID, input.Region1)
	if err != nil {
		return types.VPCDisconnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc1ID, input.Region1,
		)
	}
	vnet2, accountID2, err := c.getVPC(ctx, input.Vpc2ID, input.Region2)
	if err != nil {
		return types.VPCDisconnectionOutput{}, fmt.Errorf(
			"failed to get VPC '%s' in region '%s'", input.Vpc2ID, input.Region2,
		)
	}

	peering1 := c.getVnetPeeringFromVnet(vnet1, *vnet2.ID)
	if peering1 != "" {
		c.deleteVnetPeering(
			ctx,
			accountID1,
			parseResourceGroupName(*vnet1.ID),
			*vnet1.Name,
			vnetPeeringName(*vnet1.ID, *vnet2.ID),
		)
	} else {
		c.logger.Infof(
			"VNet Peering %s not found. Skipping it",
			vnetPeeringName(*vnet1.ID, *vnet2.ID),
		)
	}

	peering2 := c.getVnetPeeringFromVnet(vnet2, *vnet1.ID)
	if peering2 != "" {
		c.deleteVnetPeering(
			ctx,
			accountID2,
			parseResourceGroupName(*vnet2.ID),
			*vnet2.Name,
			vnetPeeringName(*vnet2.ID, *vnet1.ID),
		)
	} else {
		c.logger.Infof(
			"VNet Peering %s not found. Skipping it",
			vnetPeeringName(*vnet2.ID, *vnet1.ID),
		)
	}

	if err = c.UnblockTrafficBetweenVPCs(ctx, vnet1, vnet2, accountID1, accountID2); err != nil {
		return types.VPCDisconnectionOutput{}, fmt.Errorf(
			"failed to remove a blocking rule between %s:%s and %s:%s due to %w",
			input.Region1, *vnet1.ID, input.Region2, *vnet2.ID, err,
		)
	}

	return types.VPCDisconnectionOutput{}, nil
}
