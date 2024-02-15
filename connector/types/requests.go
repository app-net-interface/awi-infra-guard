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

package types

// ConnectionRequest provides information about the desired
// connection that should be established with the Cloud Service
// Provider Connector.
//
// It specifies both the Source and the Destination Gateways that
// should be able to pass the traffic between them so it is possible
// for Virtual Machines, which have access to them, to talk with
// themselves.
//
// The ConnectionRequest can be also used for restarting the process
// of creating Connection which was started earlier - look at
// ConnectionID field for more details.
type Request struct {
	// The Identifier of the Connection to be recreated.
	//
	// If the ConnectionID is specified, the CSP Connector will
	// try to recreate the Connection based on the corresponding
	// entry in the database.
	//
	// This can be used for retrying establishing the Connection
	// if there was an error or the Connector was interrupted and
	// the Connection is stuck in IN PROGRESS state.
	ConnectionID *string
	// The details of the connection to be created.
	ConnectionDetails *ConnectionDetails
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
	GatewayID string
	// The region where the Provider should store most of the resources
	// needed for establishing the connection.
	//
	// Be aware that some providers may create resources which are not
	// associated with any region and they are global.
	Region string
	// The Provider that should be used for preparing a single side of
	// the connection.
	//
	// The Provider name must be available in Providers subsection of the
	// given configuration. CSP Connector will validate if Source Provider
	// and Destination Provider were configured and that proper clients can
	// be created without errors.
	Provider string
	// Details is a map of additional keys and values that may
	// be required for specific providers to establish connection on
	// their side.
	Details map[string]interface{}
}

type ConnectionDetails struct {
	// The Source Gateway.
	Source GatewayIdentifier
	// The Destination Gateway
	Destination GatewayIdentifier
}

// The Provider Interface describes necessary information
// required for establishing the connection with other
// Cloud Service Providers.
//
// The Interface provides information
type ProviderInterface struct {
}
