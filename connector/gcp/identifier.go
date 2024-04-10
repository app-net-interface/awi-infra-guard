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

package gcp

import (
	"strings"

	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

// getUnifiedName ensures that the order of both gateways doesn't
// matter so that we don't get different strings depending on
// whether what is source and what is destination.
func getUnifiedName(state *transactionState) string {
	if state == nil {
		return "<NAME_GENERATION_ERROR>"
	}
	if strings.Compare(state.GatewayName, state.PeerGatewayName) == 1 {
		return state.GatewayName + "/" + state.PeerGatewayName
	}
	return state.PeerGatewayName + "/" + state.GatewayName
}

// GenerateName generates a name for the resource in a specific
// manner, which allows to connect that particular resource with
// the connection it was created for.
//
// Generate method adds a prefix with encoded string taken from
// names of both gateways that take place in the connection.
//
// TODO: Design a way to handle possible name collisions with
// hashing algorithm being used here.
func (c *GCPConnector) GenerateName(id string) string {
	return helper.CreateNameWithRand(
		getUnifiedName(c.state), id,
	)
}

// Returns true if the name, provided as an argument, indicates that
// the resource was created while creating a connection between
// GCP Cloud Router and AWS Transit Gateway from the configuration.
func (c *GCPConnector) IsNameOwnedByConnection(name string) bool {
	return helper.NameCreatedForIdentifier(
		name,
		getUnifiedName(c.state),
	)
}
