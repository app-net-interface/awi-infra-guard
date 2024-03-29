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

package types

// Request provides information about the Connection to
// be either created or removed.
//
// It specifies the information about both Gateways that
// take a part in the connection either by directly specifying
// both gateways or naming the Connection ID.
type Request struct {
	// The Identifier of the Connection.
	ConnectionID *string `yaml:"connectionID,omitempty"`
	// The details of the connection to be created.
	ConnectionDetails *ConnectionDetails `yaml:"connectionDetails,omitempty"`
}

type GatewayIdentifier struct {
	// The Gateway ID which will be used as an edge of a single
	// side of the connection.
	//
	// The Gateway ID is either the ID or the Name of the
	// resource which acts as a Gateway for a given Provider.
	// For example, in the AWS, the Transit Gateway is considered
	// as a Gateway in terms of CSP Connection while the GCP Gateway
	// refers to Cloud Router.
	GatewayID string `yaml:"id,omitempty"`
	// The region where the Provider should store most of the resources
	// needed for establishing the connection.
	//
	// Be aware that some providers may create resources which are not
	// associated with any region and they are global.
	Region string `yaml:"region,omitempty"`
	// The Provider that should be used for preparing a single side of
	// the connection.
	//
	// The Provider name must be available in Providers subsection of the
	// given configuration. CSP Connector will validate if Source Provider
	// and Destination Provider were configured and that proper clients can
	// be created without errors.
	Provider string `yaml:"provider,omitempty"`
	// Details is a map of additional keys and values that may
	// be required for specific providers to establish connection on
	// their side.
	//
	// CURRENTLY UNUSED.
	Details map[string]interface{} `yaml:"details,omitempty"`
}

type ConnectionDetails struct {
	// The Source Gateway.
	Source GatewayIdentifier `yaml:"source,omitempty"`
	// The Destination Gateway
	Destination GatewayIdentifier `yaml:"destination,omitempty"`
}
