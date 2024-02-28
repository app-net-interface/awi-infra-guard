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

// A structure defining instructions how to perform establishing
// a connection between two sides of connection.
type BGPScenario struct {
	// If set to true, the Source side of connection should start
	// the process of creating connection. Otherwise, the destination
	// should start.
	//
	// If one side generates both IP Addresses then it is necessary
	// to follow the value of this field - otherwise it's not that
	// crucial since both sides should be able to create their
	// own IP Address and then adjust to the second side regardless
	// of the order - but for consistency this value should be followed
	// anyway.
	SourceSideStarts bool

	// If set to true, the side of the connection, which will start
	// the connection creation part, is expected to generate both
	// IP Addresses - for its own and for the peer.
	//
	// Otherwise, both sides should generate its own IP Address and
	// provide it to the other side.
	StarterGeneratesBothAddresses bool
}

// BGPScenarioFromBothConfigs returns BGPScenario object with the
// information which side of connection should start establishing
// BGP connection and how should it start establishing - whether
// it should generate IP Addresses for both sides or just for itself.
//
// If the function returns nil it means that based on the received
// settings, it is either impossible or not implemented to establish
// BGP Session between these twp sides of connection.
func BGPScenarioFromBothConfigs(
	source, dest *BGPAddressing,
) *BGPScenario {
	// If any side of connection doesn't specify BGP
	// configuration then BGP session cannot be established.
	if source == nil || dest == nil {
		return nil
	}

	// Invalid BGPAddressing values for Source.
	if !source.AcceptsBothAddresses &&
		!source.GeneratesBothAddresses &&
		!source.GeneratesOwnAndAcceptsPeerAddress {
		return nil
	}

	if !dest.AcceptsBothAddresses &&
		!dest.GeneratesBothAddresses &&
		!dest.GeneratesOwnAndAcceptsPeerAddress {
		return nil
	}

	// Scenario where one side generates both IP Addresses is in
	// favor as it makes things simpler - there is no need for
	// addressing exchange - rather one side tells other what to do.
	if source.GeneratesBothAddresses && dest.AcceptsBothAddresses {
		return &BGPScenario{
			SourceSideStarts:              true,
			StarterGeneratesBothAddresses: true,
		}
	}

	if source.AcceptsBothAddresses && dest.GeneratesBothAddresses {
		return &BGPScenario{
			SourceSideStarts:              false,
			StarterGeneratesBothAddresses: true,
		}
	}

	if source.GeneratesOwnAndAcceptsPeerAddress && dest.GeneratesOwnAndAcceptsPeerAddress {
		// Currently, if both sides generate their own IP Address, we
		// set the Source side as the starting side to keep things
		// consistent.
		//
		// Adding an option to specify that a certain side should start
		// nevertheless gives additional complexity - what if both sides
		// want to start? Should we pick random side? Should we discard
		// such BGP Session? For now, such scenario is avoided.
		//
		// If a user wants a certain side to be first, it should be
		// specified as the source.
		return &BGPScenario{
			SourceSideStarts:              true,
			StarterGeneratesBothAddresses: false,
		}
	}

	// Other combinations are not supported.
	return nil
}

// The BGPAddressing struct defines supported scenarios
// for generating BGP Addresses. Each side returns such
// structure and based on that an actual Scenario can be
// determined.
//
// Each side that supports BGP should have at least one
// option enabled. BGP Sessions will be established between
// these sides which options can work with each other.
// More about it is written in BGPScenario structure.
type BGPAddressing struct {
	// The side of connection is capable of generating both
	// IP Addresses (Own and Peer). This option can work
	// only if the other side has enabled option
	// AcceptsBothAddresses.
	GeneratesBothAddresses bool

	// If enabled, it means that the side BGP IP addresses
	// can (or must) be fully customized - it will set its
	// own IP Address and Peer IP Address to these provided
	// by the second side.
	AcceptsBothAddresses bool

	// This means that the side can work in a scenario where
	// it generates its own IP Address and accepts the Peer
	// IP Address. This option can be used for the scenario
	// where the second side also specifies this option.
	GeneratesOwnAndAcceptsPeerAddress bool
}

// Represents the BGP Settings that can be used by a gateway with a help
// of its provider.
//
// Based on BGP Settings of two sides of BGP Session, a proper plan can
// be prepared to establish BGP Session.
type BGPSetting struct {
	// Information about side's ability to generate/accept BGP Addresses,
	// whether they need to be generated by that side or accepted from
	// peer side.
	Addressing BGPAddressing
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

// The BGP Attachment mode describes what should be accomplished
// during a particular phase of creating BGP Session.
//
// Currently, CSP Connector supports two scenarios:
// 1. Authoritarian
// 2. Cooperative
//
// The first scenario expects following order of provided Modes
// - 1st side starts with AttachModeGenerateBothIPs
// - 2nd side runs with AttachModeAcceptOtherIP
// Connection established.
//
// The cooperative scenario should go as follows:
// - 1st side starts with AttachModeGenerateIP
// - 2nd side starts with AttachModeGenerateIPAndAcceptOtherIP
// - 1st side ends with AttachModeAcceptOtherIP
// Connection established./
//
// So the Cooperative mode requires one more step than the
// Authoritation mode as the 1st side needs to learn what
// 2nd side accomplished.
type AttachBGPConnectionMode int

const (
	// The Mode indicating that the provider should generate BGP IP
	// Addresses for both sides of the connection.
	//
	// If the first Gateway is attached with this network, the second
	// provider needs to run only AttachModeAcceptOtherIP.
	AttachModeGenerateBothIPs AttachBGPConnectionMode = iota
	// The Mode indicating that this is the first stage of creating
	// BGP Session. The first provider needs to generate its side of
	// BGP Session. The other side IP Address is unknown yet so the
	// provider needs to wait for the other side to complete mode
	// AttachModeGenerateIPAndAcceptOtherIP.
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
	AttachModeGenerateIPAndAcceptOtherIP
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
	AttachModeAcceptOtherIP
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
	// * AttachModeGenerateIPAndAcceptOtherIP
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
