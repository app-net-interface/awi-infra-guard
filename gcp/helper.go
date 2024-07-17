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
	"net"
	"slices"
	"strings"
)

// Helper function to extract the last segment of a resource URL, typically the ID or name.
func extractResourceID(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func isIPv4CIDR(cidr string) bool {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return false // Not a valid CIDR notation
	}
	return ip.To4() != nil // Returns true if CIDR is IPv4
}

// ContainsAny checks if any element of slice2 is present in slice1
func ContainsAny(slice1, slice2 []string) bool {
	for _, item := range slice2 {
		if slices.Contains(slice1, item) {
			return true
		}
	}
	return false
}
