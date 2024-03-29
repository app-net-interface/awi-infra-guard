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

// NOT IMPLEMENTED
//
// Just a placeholder for providing necessary instructions to
// create a Static Routing connection.
//
// Depending on the future approach with static routing it can
// either be a simple set of routes which will be routed to that
// particular Gateway or it can provide some information how to
// obtain the current list of routes (HTTP Clients, custom logic etc.)
type StaticRoutingSetting struct{}

// ProviderSettings define an entrypoint for preparing a plan
// how to establish a connection between two providers.
//
// At least one of BGPSetting or StaticRoutingSetting should
// be not nil - otherwise a connection cannot be established.
//
// The BGP Connection is preferred over Static Routing. If
// both sides support BGP Routing, the CSP Connector will attempt
// creating such connection. Otherwise, it will try to create
// Statically Routed Session (unless at least one side doesn't
// support Static Routing as well).
type GatewayConnectionSettings struct {
	// The structure representing necessary information for
	// establishing BGP session.
	//
	// If nil, it means that the Gateway doesn't support BGP
	// routing.
	BGPSetting *BGPSetting
	// The structure representing necessary information for
	// establishing Statically Routed session.
	//
	// If nil, it means that the Gateway doesn't support
	// static routing.
	StaticRoutingSetting *StaticRoutingSetting
	// The number of Interfaces that will be exposed by the
	// Gateway.
	//
	// The number of tunnels to be established between two
	// Gateways is equal to the bigger value of NumberOfInterfaces
	// from both Gateways. The greater value need to be multiple of
	// the lesser value.
	NumberOfInterfaces uint8
	// The upper limit for number of tunnels for a Gateway. If set
	// to 0 the limit is turned off.
	MaxNumberOfTunnels uint8
}
