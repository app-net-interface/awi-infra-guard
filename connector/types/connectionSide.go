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

// Represents the BGP Settings that can be used by a gateway with a help
// of its provider.
//
// Based on BGP Settings of two sides of BGP Session, a proper plan can
// be prepared to establish BGP Session.
type BGPSetting struct {
	// If set to true, it means that the side can pick its own IP Address.
	//
	// It will use that if the second side does not define PickOtherIPAddress
	// as true.
	PickOwnIPAddress bool
	// If set to true, it means that the side is capable of preparing IP
	// Address for the second side of connection.
	//
	// Providers such as AWS generate configuration for both sides of BGP
	// session so the other side does not need to generate it.
	PickOtherIPAddress bool
	// The CIDRs for Network from which IP Addresses can be picked to represent
	// this side of BGP Connection.
	//
	// In most situations it will be a single 169.254.0.0/16 block.
	//
	// If BGPIPRanges do not overlap, the BGP Session cannot be established.
	AllowedIPRanges []string
	// CIDRs within BGPIPRange which are reserved and cannot be used.
	ExcludedIPRanges []string
}

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
}

// The BGP Attachment mode describes what should be accomplished
// during a particular phase of creating BGP Session.
//
// Currently, CSP Connector supports two scenarios.
//
// First scenario with master mind:
//   - 1st provider generates BGP IP Addresses for both
//     sides of the session and creates all necessary
//     resources for its side
//   - 2nd provider ackknowledges IP Addresses that were
//     generated for its gateway and creates IP resources
//     which will use those IP Addresses and point at the
//     IP Addresses from the second side
type AttachBGPConnectionMode int

const (
	// The Mode indicating that the provider should generate BGP IP
	// Addresses for both sides of the connection.
	//
	// If the first Gateway is attached with this network, the second
	// provider needs to run only AttachModeAttachOtherIP.
	AttachModeGenerateBothIPs AttachBGPConnectionMode = iota
	// The Mode indicating that this is the first stage of creating
	// BGP Session. The first provider needs to generate its side of
	// BGP Session. The other side IP Address is unknown yet so the
	// provider needs to wait for the other side to complete mode
	// AttachModeGenerateIPAndAttachOtherIP.
	AttachModeGenerateIP
	// This Mode is used at the second stage of creating BGP session
	// if the first provider used AttachModeGenerateIP and returned
	// its BGP IP Address. The second provider will generate its
	// own IP Address for BGP Session and create proper resources
	// with the knowledge of the BGP IP Address from the other side.
	//
	// The second side of connection should be completed and now the
	// first side needs to learn the generated IP Address to complete
	// its side.
	AttachModeGenerateIPAndAttachOtherIP
	// The final stage of establishing BGP Session. After second provider
	// attached itself to this gateway's BGP IP Address and generated its
	// own BGP IP Address, this provider can use it to complete this side
	// of connection.
	//
	// This mode is expected to be called for second provider if the first
	// provider used AttachModeGenerateBothIPs or for the first provider
	// if it used AttachModeGenerateIP mode.
	//
	// After completing the Attachment process with this mode, the connection
	// should be considered to be established on both sides.
	AttachModeAttachOtherIP
)

type CreateBGPConnectionConfig struct {
	// The public IP Addresses of the second gateway.
	OutsideInterfaces []string
	// The ASN associated with this Gateway.
	ASN uint64
	// The identifying ASN of the other side.
	PeerASN uint64
	// Number of tunnels that should be created.
	NumberOfTunnels uint8
	// BGP Addresses that should be used by this gateway.
	// If this is not nil it means that the second provider
	// generated those for this gateway and expects this gateway
	// to use them.
	BGPAddresses []string
	// BGP Addresses of the other side.
	PeerBGPAddresses []string
	// Shared Secrets.
	SharedSecrets []string
}

// The data obtained during creating provider specific resources.
//
// If during Gateway Attachment, the provider generated BGP IP Addresses,
// picked IPSec Tunnel PreShared secrets, picked public IP Addresses as
// Gateway interfaces, those informations will be returned with this
// structure.
type OutputForConnectionWithBGP struct {
	// The IP Addresses generated for that particular gateway. This slice
	// is expected to be not nil for modes
	// * AttachModeGenerateBothIPs
	// * AttachModeGenerateIP
	// * AttachModeGenerateIPAndAttachOtherIP
	BGPAddresses []string
	// The IP Addresses generated for the other gateway. This slice is
	// expected to be not nil only for AttachModeGenerateBothIPs mode.
	PeerBGPAddresses []string
	// Shared Secrets.
	SharedSecrets []string
	// Interfaces returns list of interfaces for providers that
	// cannot return interfaces during InitializeGatewayInterface
	// phase.
	Interfaces []string
}
