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

package db

import (
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/types"
)

func (client *boltClient) PutVpcConnection(v *types.VPCConnection) error {
	return update(client, v, v.DbId(), vpcConnectionTable)
}

func (client *boltClient) GetVpcConnection(id string) (*types.VPCConnection, error) {
	vpcConnections, err := list[types.VPCConnection](client, vpcConnectionTable)
	if err != nil {
		return nil, fmt.Errorf("failed to list vpc connections: %w", err)
	}

	for _, conn := range vpcConnections {
		fmt.Printf("DEBUG: Checking VPC connection with FromVpcId: %s or ToVpcId: %s against %s\n",
			conn.FromVpcId, conn.ToVpcId, id)
		// Check both FromVpcId and ToVpcId since the connection could be in either direction
		if conn.FromVpcId == id || conn.ToVpcId == id {
			return conn, nil
		}
	}
	return nil, fmt.Errorf("vpc connection with vpc id %s not found", id)
}

func (client *boltClient) ListVpcConnections() ([]*types.VPCConnection, error) {
	return list[types.VPCConnection](client, vpcConnectionTable)
}

func (client *boltClient) DeleteVpcConnection(id string) error {
	return delete_(client, id, vpcConnectionTable)
}

// hasTransitGatewayRoute checks if a route table has a route to the specified transit gateway
func hasTransitGatewayRoute(rt *types.RouteTable, tgwID string) bool {
	if rt == nil {
		return false
	}
	for _, route := range rt.Routes {
		// Check if the route target contains the TGW ID
		if route.Target == tgwID {
			return true
		}
	}
	return false
}

// SyncVpcConnections is now handled in the sync package instead of here
func (client *boltClient) SyncVpcConnections() error {
	// No-op: Sync logic moved to sync package
	return nil
}

// UpdateVpcConnectionStatus analyzes cached route tables to determine actual connectivity status
// of Transit Gateway connections based on route table configurations
func (client *boltClient) UpdateVpcConnectionStatus() error {
	// Get all VPC connections
	connections, err := client.ListVpcConnections()
	if err != nil {
		return fmt.Errorf("failed to list VPC connections: %w", err)
	}

	// Get all route tables for lookup
	routeTables, err := client.ListRouteTables()
	if err != nil {
		return fmt.Errorf("failed to list route tables: %w", err)
	}

	// Create a map of VPC ID to its route tables for quick lookup
	vpcRouteTablesMap := make(map[string][]*types.RouteTable)
	for _, rt := range routeTables {
		vpcRouteTablesMap[rt.VpcID] = append(vpcRouteTablesMap[rt.VpcID], rt)
	}

	// Analyze each connection
	for _, conn := range connections {
		// Only process Transit Gateway connections
		if conn.ConnectionType != "transit_gateway" {
			continue
		}

		// Parse connection ID to get TGW ID (format: tgwID:attachment1:attachment2)
		parts := strings.Split(conn.ID, ":")
		if len(parts) != 3 {
			fmt.Printf("WARNING: Invalid transit gateway connection ID format: %s\n", conn.ID)
			continue
		}
		tgwID := parts[0]

		// Get route tables for both VPCs
		fromVpcRouteTables := vpcRouteTablesMap[conn.FromVpcId]
		toVpcRouteTables := vpcRouteTablesMap[conn.ToVpcId]

		// Check if either VPC has a route to the TGW
		fromVpcHasRoute := false
		toVpcHasRoute := false

		for _, rt := range fromVpcRouteTables {
			if hasTransitGatewayRoute(rt, tgwID) {
				fromVpcHasRoute = true
				break
			}
		}

		for _, rt := range toVpcRouteTables {
			if hasTransitGatewayRoute(rt, tgwID) {
				toVpcHasRoute = true
				break
			}
		}

		// Update connection status based on routing
		oldStatus := conn.Status
		if fromVpcHasRoute && toVpcHasRoute {
			conn.Status = "active"
		} else if fromVpcHasRoute || toVpcHasRoute {
			conn.Status = "partial"
		} else {
			conn.Status = "inactive"
		}

		// Only update if status changed
		if oldStatus != conn.Status {
			fmt.Printf("Updating TGW connection %s status from %s to %s\n", conn.ID, oldStatus, conn.Status)
			if err := client.PutVpcConnection(conn); err != nil {
				fmt.Printf("ERROR: Failed to update connection %s: %v\n", conn.ID, err)
				continue
			}
		}
	}

	return nil
}
