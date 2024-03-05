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

func (c *GCPConnector) assertExternalVPNGatewayInterfaces(
	ipAddresses []string, gateway *client.ExternalVPNGateway,
) bool {
	if gateway == nil {
		return false
	}
	if len(gateway.Interfaces) != len(ipAddresses) {
		return false
	}
	// TODO: Consider allowing different order of IPs.
	// While we cannot support all kinds of permutations here
	// since first two IP Addresses belong to first connection and
	// last two IP addresses belong to the second connection, we
	// should be probably still able to allow mixing orders of
	// IP addresses within the same connection.
	for i := range ipAddresses {
		if ipAddresses[i] != gateway.Interfaces[i].IP {
			return false
		}
	}
	return true
}

// Returns an External VPN Gateway which matches the list of
// IP Addresses provided for it. It's purpose is to check if
// there was already created an External VPN Gateway according
// to the Interfaces from the other provider.
func (c *GCPConnector) getMatchingExternalVPNGateway(
	ctx context.Context, ipAddresses []string,
) (*client.ExternalVPNGateway, error) {
	vpnGateways, err := c.gcpClient.ListExternalVPNGateway(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot verify the current state of External VPN Gateways: %w", err)
	}
	for i := range vpnGateways {
		if vpnGateways[i] == nil {
			continue
		}
		if c.assertExternalVPNGatewayInterfaces(ipAddresses, vpnGateways[i]) {
			return vpnGateways[i], nil
		}
	}
	return nil, nil
}

func (c *GCPConnector) createExternalVPNGateway(
	ctx context.Context,
	interfaces []string,
) (*client.ExternalVPNGateway, error) {
	c.logger.Debugf(
		"Starting to prepare External VPN Gateway"+
			"The External VPN Gateway is supposed to have following interfaces: %v",
		interfaces,
	)

	c.logger.Debug(
		"Checking if there is already existing Existing VPN Gateway " +
			"that matches given criterias")
	gateway, err := c.getMatchingExternalVPNGateway(ctx, interfaces)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot verify if there was already created External VPN Gateway: %w", err)
	}
	if gateway != nil {
		c.logger.Debugf("found External VPN Gateway '%s' matching prerequisites", gateway.Name)
		return gateway, nil
	}
	c.logger.Debugf(
		"External VPN Gateway with interfaces '%v' not found. Creating one", interfaces)

	// TODO: Define some standard rules for assigning unique and identifiable name
	// for GCP resources and handle potential name conflict with either retrying or
	// renaming policy.
	gateway = &client.ExternalVPNGateway{
		Name:       c.GenerateName("ext-vpn-gw"),
		Interfaces: []client.ExternalVPNGatewayInterface{},
	}
	for _, iface := range interfaces {
		gateway.Interfaces = append(gateway.Interfaces, client.ExternalVPNGatewayInterface{
			IP: iface,
		})
	}
	created, err := c.gcpClient.CreateExternalVPNGateway(ctx, gateway)
	if err != nil {
		return nil, fmt.Errorf("failed to create External VPN Gateway %v: %w", *gateway, err)
	}
	c.logger.Debugf(
		"External VPN Gateway '%v' created", *created)
	return created, nil
}
