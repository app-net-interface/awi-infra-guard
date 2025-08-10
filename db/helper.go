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

// extractPorts extracts individual ports from a port range string

package db

import (
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/types"
)

func extractPorts(portRange string) []string {
	ports := []string{}

	// Handle single port
	if !strings.Contains(portRange, "-") {
		ports = append(ports, strings.TrimSpace(portRange))
		return ports
	}

	// Handle port range
	parts := strings.Split(portRange, "-")
	if len(parts) == 2 {
		start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 == nil && err2 == nil {
			// For analysis, we'll focus on commonly known risky ports
			riskyPorts := []int{
				21,    // FTP
				22,    // SSH
				23,    // Telnet
				25,    // SMTP
				53,    // DNS
				80,    // HTTP
				135,   // Windows RPC
				139,   // NetBIOS
				443,   // HTTPS
				445,   // SMB
				465,   // SMTPS
				587,   // SMTP submission
				993,   // IMAPS
				995,   // POP3S
				1433,  // SQL Server
				1521,  // Oracle
				3306,  // MySQL
				3389,  // RDP
				5432,  // PostgreSQL
				6379,  // Redis
				27017, // MongoDB
			}
			for _, port := range riskyPorts {
				if port >= start && port <= end {
					ports = append(ports, strconv.Itoa(port))
				}
			}
		}
	}

	return ports
}

// removeDuplicates removes duplicate strings from a slice
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
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
