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

package sqlite

import (
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/types"
)

func (client *sqliteClient) PutVpcConnection(v *types.VPCConnection) error {
	return client.putObject(v, v.DbId(), vpcConnectionTable)
}

func (client *sqliteClient) GetVpcConnection(id string) (*types.VPCConnection, error) {
	vpcConnection := &types.VPCConnection{}
	err := client.getObject(id, vpcConnectionTable, vpcConnection)
	if err != nil {
		return nil, err
	}
	return vpcConnection, nil
}

func (client *sqliteClient) ListVpcConnections() ([]*types.VPCConnection, error) {
	objects, err := client.listObjects(vpcConnectionTable, func() interface{} { return &types.VPCConnection{} })
	if err != nil {
		return nil, err
	}

	vpcConnections := make([]*types.VPCConnection, len(objects))
	for i, obj := range objects {
		vpcConnections[i] = obj.(*types.VPCConnection)
	}
	return vpcConnections, nil
}

func (client *sqliteClient) DeleteVpcConnection(id string) error {
	return client.deleteObject(id, vpcConnectionTable)
}

// SyncVpcConnections is now handled in the sync package instead of here
func (client *sqliteClient) SyncVpcConnections() error {
	// No-op: Sync logic moved to sync package
	return nil
}

// UpdateVpcConnectionStatus analyzes cached route tables to determine actual connectivity status
// of Transit Gateway connections based on route table configurations
func (client *sqliteClient) UpdateVpcConnectionStatus() error {
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
