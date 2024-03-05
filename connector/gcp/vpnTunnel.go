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
	"context"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/gcp/client"
)

// PlanVPNTunnels creates objects representing desired VPN Tunnels.
//
// The GCP Provider prepares the number of tunnels equal to the
// length of sharedSecrets provided as an argument.
//
// Each VPN Tunnel consists of:
// * Cloud Router URL
// * VPN Gateway URL
// * External VPN Gateway
//
// The first two elements come from the GCP Gateway that was
// selected for connection creation. The latter one is a resource
// created by the CSP Connector to represent the second side of
// the connection.
//
// Additionally, each VPN Tunnel has a shared secret which will
// be picked from the provided slice. The order of secrets is
// important - to demonstrate, let's visualise situation where
// VPN Gateway defines 2 interfaces (IF_0 and IF_1) and External
// VPN gateway defines 4 interfaces (EXT_IF_0, EXT_IF_1, EXT_IF_2,
// EXT_IF_3). Such scenario would produce following Tunnels:
//
// Tunnel name		VPN GW IF		EXT IF		Shared Secret
// tunnel-1			IF_0			EXT_IF_0	SECRET_0
// tunnel-2			IF_0			EXT_IF_1	SECRET_1
// tunnel-3			IF_1			EXT_IF_2	SECRET_2
// tunnel-4			IF_1			EXT_IF_3	SECRET_3
func (c *GCPConnector) PlanVPNTunnels(
	cloudRouterURL string,
	vpnGatewayURL string,
	externalVPNGatewayURL string,
	sharedSecrets []string,
) []client.VPNTunnel {
	var numberOfTunnels uint8 = uint8(len(sharedSecrets))
	tunnels := make([]client.VPNTunnel, 0, numberOfTunnels)

	for i := uint8(0); i < NUMBER_OF_GATEWAY_INTERFACES; i++ {
		for j := uint8(0); j < (numberOfTunnels / NUMBER_OF_GATEWAY_INTERFACES); j++ {
			tunnels = append(tunnels, client.VPNTunnel{
				Name: c.GenerateName(
					fmt.Sprintf("tunnel-%d-", ((i*(numberOfTunnels/NUMBER_OF_GATEWAY_INTERFACES))+j)+1)),
				ExternalVPNGateway:       externalVPNGatewayURL,
				ExternalGatewayInterface: int((i * (numberOfTunnels / NUMBER_OF_GATEWAY_INTERFACES)) + j),
				VPNGateway:               vpnGatewayURL,
				IKEVersion:               2,
				Interface:                int(i),
				CloudRouter:              cloudRouterURL,
				SharedSecret:             sharedSecrets[(i*(numberOfTunnels/NUMBER_OF_GATEWAY_INTERFACES))+j],
			})
		}
	}

	return tunnels
}

// tunnelMatchesSourceAndDestination returns true if the
// tunnel connects Cloud Router and VPN Gateway from the GCP
// Gateway and the External VPN Gateway representing the
// second side of the connection.
func (c *GCPConnector) tunnelMatchesSourceAndDestination(
	tunnel client.VPNTunnel,
	cloudRouterURL string,
	vpnGatewayURL string,
	externalVPNGatewayURL string,
) bool {
	return tunnel.CloudRouter == cloudRouterURL &&
		tunnel.VPNGateway == vpnGatewayURL &&
		tunnel.ExternalVPNGateway == externalVPNGatewayURL
}

// findExistingVPNTunnels looks for already created VPN Tunnels,
// which are established between two sides of connection.
//
// Returns the slice of pointers to VPN Tunnels. A nil value at
// a position X indicates that the VPN Tunnel representing a
// VPN Tunnel attached to External VPN Gateway at interface X
// was not found. Otherwise, the given interface should not
// be created as it already exists.
func (c *GCPConnector) findExistingVPNTunnels(
	ctx context.Context,
	cloudRouterURL string,
	vpnGatewayURL string,
	externalVPNGatewayURL string,
	numberOfInterfaces uint8,
) ([]*client.VPNTunnel, error) {
	matchedTunnels := make([]*client.VPNTunnel, numberOfInterfaces)

	tunnels, err := c.gcpClient.ListVPNTunnels(ctx, c.config.Region)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the list of VPN Tunnels: %w", err)
	}

	for i := range tunnels {
		if tunnels[i] == nil {
			continue
		}

		if !c.tunnelMatchesSourceAndDestination(
			*tunnels[i], cloudRouterURL, vpnGatewayURL, externalVPNGatewayURL,
		) {
			continue
		}

		tunnelID := tunnels[i].ExternalGatewayInterface
		if tunnelID >= len(matchedTunnels) {
			c.logger.Warnf(
				"Found VPN Tunnel between Cloud Router %s with VPN Gateway %s and "+
					"External VPN Gateway %s with the External VPN Gateway Interface ID %d "+
					"(starting from the ID 0) "+
					"but expected %d interfaces at max (the expectation is that the "+
					"External VPN Gateway has the number of interfaces equal to the "+
					"total number of VPN Tunnels to be created between gateways)",
				cloudRouterURL, vpnGatewayURL, externalVPNGatewayURL, tunnelID,
				numberOfInterfaces,
			)
		}

		if matchedTunnels[tunnelID] == nil {
			c.logger.Debugf("Found matching VPN Tunnel for interface %d: %s", tunnelID, tunnels[i].Name)
			matchedTunnels[tunnelID] = tunnels[i]
			continue
		}

		c.logger.Warnf(
			"Found another VPN Tunnel for interface %d: %s. "+
				"Ignoring it and keeping previously found '%s' VPN Tunnel.",
			tunnelID, tunnels[i].Name, matchedTunnels[tunnelID].Name)
	}

	return matchedTunnels, nil
}

// createVPNTunnels creates VPN Tunnels connecting both sides
// of the connection. The GCP side is represented by the Cloud Router
// and the VPN Gateway and the other side is represented by the
// External VPN Gateway.
//
// CreateVPNTunnels verifies if there are any already created VPN
// Tunnels and reuses them.
func (c *GCPConnector) createVPNTunnels(
	ctx context.Context,
	cloudRouterURL string,
	vpnGatewayURL string,
	externalVPNGatewayURL string,
	sharedSecrets []string,
) ([]client.VPNTunnel, error) {
	c.logger.Debug("Starting to prepare VPN Tunnels")

	numberOfTunnels := uint8(len(sharedSecrets))

	tunnelsToCreate := c.PlanVPNTunnels(
		cloudRouterURL,
		vpnGatewayURL,
		externalVPNGatewayURL,
		sharedSecrets,
	)
	c.logger.Debugf(
		"For the connection, the following VPN Tunnels should exist. %v."+
			"Names are irrelevant: they will be used if vpn tunnels are missing and need "+
			"to be created", tunnelsToCreate)

	// TODO: Verify if the Shared Secret should not be overwritten for existing
	// VPN Tunnels to ensure that both sides use proper Shared Keys.
	currentTunnels, err := c.findExistingVPNTunnels(
		ctx,
		cloudRouterURL,
		vpnGatewayURL,
		externalVPNGatewayURL,
		numberOfTunnels,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to verify currently existing VPN Tunnels: %w", err)
	}
	for i := range currentTunnels {
		if currentTunnels[i] != nil {
			c.logger.Debugf(
				"the vpn tunnel with ID %d is already satisfied: %v", i, currentTunnels[i])
		}
	}

	finalTunnels := []client.VPNTunnel{}
	for i := range currentTunnels {
		if currentTunnels[i] != nil {
			c.logger.Debugf(
				"VPN Tunnel for interface %d already exists with name '%s'. Using it.",
				i, currentTunnels[i].Name)
			finalTunnels = append(finalTunnels, *currentTunnels[i])
			continue
		}
		c.logger.Debugf("VPN Tunnel for interface %d not found. Creating it.", i)

		createdTunnel, err := c.gcpClient.CreateVPNTunnel(ctx, &tunnelsToCreate[i], c.config.Region)
		if err != nil {
			return nil, fmt.Errorf("failed to create a VPN Tunnel %v due to: %w", tunnelsToCreate[i], err)
		}
		if createdTunnel == nil {
			return nil, fmt.Errorf("got unexpected empty VPN Tunnel after creation. Expected: %v", tunnelsToCreate[i])
		}
		finalTunnels = append(finalTunnels, *createdTunnel)
		c.logger.Debugf("VPN Tunnel %s created successfully.", tunnelsToCreate[i].Name)
	}

	return finalTunnels, nil
}
