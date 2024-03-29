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

package provider

import (
	"context"

	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

// Provider is an interface for a single Cloud Service Provider.
//
// The CSP Connector establishes a connection between two Gateways.
// Each Gateway is associated with a certain Provider that will be
// used to create/delete all necessary resources on the actual
// Cloud Service Provider side.
//
// The Provider structures need to implement all methods below, which
// will give the following abilities:
//   - List all resources that can be considered Gateway from the CSP
//     Connector point of view,
//   - Gather information required for establishing a connection (for
//     example VPC CIDRs to prepare Route Tables)
//   - Create resources needed for establishing the Connection with other
//     provider
//   - Remove resources, that were previously created, if the connection
//     is no longer needed.
//
// Providers can and should implement their own optimization ways such
// as developing internal Cache Storages for avoiding constant loading
// of the same resources.
//
// For example, an AWS Connector, which implements Provider interface,
// uses AWS SDK underneath. When the CSP Connector tries to establish
// a connection between AWS Gateway and some other Gateway, the AWS
// Provider will use that particular SDK to create Customer Gateways,
// VPN Gateways, VPN Connections and generate all necessary information
// for the second Provider.
type Provider interface {
	// Name returns the name of the Provider that will be treated as an identifier
	// for that specific provider.
	//
	// The Name() should return a unique provider name in terms of your project.
	//
	// The name is case-insensitive. Each part of the code, that uses this method,
	// will run strings.ToLower() on the result of this method to make sure that
	// the casing won't make any troubles.
	//
	// TODO: Confirm case-insensitivity.
	Name() string

	// Looks for Gateway matching the provided identifier.
	//
	// Returns the pointer to the Gateway structure if it found one or returns nil if
	// there was no such Gateway. If there is an error such as Connection Problem,
	// lack of authorization etc. it should be reported as an error.
	//
	// This method should be deterministic. A single Identifier should allow the Provider
	// to match either zero or one Gateway. If some details are missing some valid information
	// and causes a possibility for multiple matches, the Provider should not look for such
	// Gateways but immediately report an error that the provided information makes it impossible
	// to determine a single matching gateway. For example, if the AWS Provider receives GetGateway
	// with no region provided, it should report an error - even if there was only one Transit Gateway
	// in the entire AWS, the provider should report an error due to the POSSIBILITY that in certain
	// circumstances multiple Gateways could match the same criteria (two Transit Gateways with the
	// same name in different regions).
	//
	// The GetGateway should return an error if there was an actual error while trying to get the
	// matching Gateway. Error NotFound SHOULD NOT be treated as an error. If the resource was
	// not found then a nil response should be returned.
	GetGateway(ctx context.Context, gateway types.GatewayIdentifier) (*types.Gateway, error)

	// Lists all resources that can be treated as a Gateway by a given provider.
	//
	// For example, AWS Provider will return all existing Transit Gateways and GCP Provider
	// will return all Cloud Routers.
	ListGateways(ctx context.Context) ([]types.Gateway, error)

	// GetGatewayConnectionSettings returns necessary information for
	// establishing a connection with other Gateway.
	GetGatewayConnectionSettings(
		ctx context.Context, gateway types.Gateway,
	) (types.GatewayConnectionSettings, error)

	// InitializeCreation is a method for each provider to learn about the
	// planned Create operation so the provider can initialize potential
	// transaction objects, caches etc.
	InitializeCreation(
		ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
	) error

	// InitializeGatewayInterfaces creates or looks up resources to obtain
	// Public IP Addresses that should be used as Interfaces for creating
	// Tunnels toward that Gateway.
	//
	// If the returned slice is empty, it means that the provider was not
	// able to verify Interfaces on the initial phase of creating a connection
	// and that most likely this information will be available after calling
	// AttachToExternalGatewayWithBGP or AttachToExternalGatewayWithStaticRouting
	// for that provider.
	InitializeGatewayInterfaces(
		ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
	) ([]string, error)

	// InitializeASN creates or looks up resources to check what is the
	// ASN for that particular gateway.
	InitializeASN(
		ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
	) (uint64, error)

	// AttachToExternalGatewayWithBGP performs necessary operations for
	// the Gateway with the help of its Provider to establish its side
	// of the connection with other Gateway (peerGateway).
	//
	// AttachMode specifies instructions for establishing BGP session,
	// whether the provider should generate IP Addresses for this Gateway
	// and the peer Gateway, generate only its own BGP IP Address or just
	// attach the existing one, provided as an argument in config parameter.
	AttachToExternalGatewayWithBGP(
		ctx context.Context,
		gateway types.Gateway,
		peerGateway types.Gateway,
		attachMode types.AttachBGPConnectionMode,
		config types.CreateBGPConnectionConfig,
	) (types.OutputForConnectionWithBGP, error)

	// NOT IMPLEMENTED
	//
	// Just a placeholder for future CSP improvements.
	AttachToExternalGatewayWithStaticRouting() error

	DeleteConnectionResources(
		ctx context.Context,
		gateway types.Gateway,
		peerGateway types.Gateway,
	) error

	// Returns a slice of CIDRs that can be reached using a connection with
	// that particular Gateway.
	GetCIDRs(ctx context.Context, gateway types.Gateway) ([]string, error)

	// Returns the ID of a VPC where associated with Gateway resource.
	GetVPCForGateway(ctx context.Context, gateway types.Gateway) (string, error)

	// Closes all underlying clients/sockets/files etc. used by that particular
	// Provider.
	Close() error
}
