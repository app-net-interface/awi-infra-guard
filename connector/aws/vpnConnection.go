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

package aws

import (
	"context"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/aws/client"
	"github.com/app-net-interface/awi-infra-guard/connector/cidrpool"
)

func (c *AWSConnector) createTunnelOption(
	pool *cidrpool.CIDRV4Pool,
	secret string,
) (*client.TunnelOption, error) {
	cidr, err := pool.Get(REQUESTED_CIDR_SIZE)
	if err != nil {
		return nil, fmt.Errorf("error while trying to obtain a new CIDR: %w", err)
	}
	if cidr == nil {
		return nil, fmt.Errorf(
			"could not obtain CIDR with size %d from the pool %v. "+
				"Cannot fit requested CIDR size in the pool",
			REQUESTED_CIDR_SIZE, pool)
	}

	return &client.TunnelOption{
		CIDR:         cidr.String(),
		PreSharedKey: secret,
	}, nil
}

func (c *AWSConnector) createTunnelOptions(
	bgpPool [2]*cidrpool.CIDRV4Pool,
	secrets [2]string,
) ([2]client.TunnelOption, error) {
	tunnelOptions := [2]client.TunnelOption{}

	for i := 0; i < 2; i++ {
		tunnelOption, err := c.createTunnelOption(bgpPool[i], secrets[i])
		if err != nil {
			return [2]client.TunnelOption{}, fmt.Errorf(
				"failed to prepare tunnel options: %w", err,
			)
		}
		tunnelOptions[i] = *tunnelOption
	}

	return tunnelOptions, nil
}

func (c *AWSConnector) getVPNConnectionForCustomerGateway(
	ctx context.Context,
	transitGatewayID string,
	customerGatewayID string,
) (*client.VPNConnection, error) {
	connections, err := c.awsClient.ListVPNConnections(
		ctx,
		client.ListVPNConnectionFilterCustomerGatewayID{
			Value: customerGatewayID,
		},
		client.ListVPNConnectionFilterTransitGatewayID{
			Value: transitGatewayID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain list of VPN Connections between Transit Gateway "+
				"%s and Customer Gateway %s: %w",
			transitGatewayID, customerGatewayID, err,
		)
	}
	if len(connections) > 1 {
		return nil, fmt.Errorf(
			"found more than 1 existing VPN Connection between Transit Gateway "+
				"%s and Customer Gateway %s: %v. Unable to determine what should be done next",
			transitGatewayID, customerGatewayID, connections,
		)
	}
	if len(connections) == 0 {
		return nil, nil
	}
	if connections[0] == nil {
		return nil, fmt.Errorf(
			"got nil VPN Connection between Transit Gateway "+
				"%s and Customer Gateway %s. Unable to determine what should be done next",
			transitGatewayID, customerGatewayID,
		)
	}
	return connections[0], nil
}

func (c *AWSConnector) createVPNConnectionForCustomerGateway(
	ctx context.Context,
	transitGatewayID string,
	customerGatewayID string,
	bgpPools [2]*cidrpool.CIDRV4Pool,
	secrets [2]string,
) (*client.VPNConnection, error) {
	connection, err := c.getVPNConnectionForCustomerGateway(
		ctx, transitGatewayID, customerGatewayID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to check if a VPN Connection between Transit Gateway %s "+
				"and Customer Gateway %s already exist: %w",
			transitGatewayID, customerGatewayID, err,
		)
	}
	if connection != nil {
		c.logger.Warnf(
			"VPN Connection between Transit Gateway %s and Customer Gateway %s "+
				"already exists: %s. The CSP Connector will keep it instead of creating "+
				"a new connection - please note, if the VPN connection on the other side "+
				"was created from scratch the secrets will most likely not match! In that case "+
				"delete stale VPN Connection and let CSP Connector recreate connection on "+
				"both sides to use the same Shared Secret",
			transitGatewayID, customerGatewayID, connection.ID,
		)
		return connection, nil
	}
	c.logger.Debugf(
		"VPN Connection between Transit Gateway %s and Customer Gateway %s "+
			"not found. Creating it.",
		transitGatewayID, customerGatewayID,
	)

	tunnelOpts, err := c.createTunnelOptions(bgpPools, secrets)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to prepare tunnel options for VPN Connection between "+
				"Transit Gateway %s and Customer Gateway %s: %w",
			transitGatewayID, customerGatewayID, err,
		)
	}

	vpnConnection, err := c.awsClient.CreateVPNConnection(
		ctx,
		customerGatewayID,
		transitGatewayID,
		tunnelOpts,
		c.generateTagName(),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VPN Connection between "+
				"Transit Gateway %s and Customer Gateway %s: %w",
			transitGatewayID, customerGatewayID, err,
		)
	}
	if vpnConnection == nil {
		return nil, fmt.Errorf(
			"failed to create VPN Connection between "+
				"Transit Gateway %s and Customer Gateway %s. Got nil object",
			transitGatewayID, customerGatewayID,
		)
	}
	return vpnConnection, nil
}

func (c *AWSConnector) createVPNConnectionsIfNeeded(
	ctx context.Context,
	transitGatewayID string,
	customerGatewayIDs []string,
	bgpPools []*cidrpool.CIDRV4Pool,
	secrets []string,
) ([]*client.VPNConnection, error) {
	c.logger.Debug("Checking AWS VPN Connections")
	if len(secrets) != len(customerGatewayIDs)*2 {
		return nil, fmt.Errorf(
			"number of secrets doesn't match with the number of requested "+
				"connections between Transit Gateway %s and Customer Gateways %v ."+
				"Every VPN Connection requires 2 secrets so the number of secrets should "+
				"be equal to %d but got %d",
			transitGatewayID, customerGatewayIDs, len(customerGatewayIDs)*2, len(secrets),
		)
	}
	vpnConnections := make([]*client.VPNConnection, len(customerGatewayIDs))
	for i, cgw := range customerGatewayIDs {
		connection, err := c.createVPNConnectionForCustomerGateway(
			ctx,
			transitGatewayID,
			cgw,
			[2]*cidrpool.CIDRV4Pool{bgpPools[i*2], bgpPools[i*2+1]},
			[2]string{secrets[i*2], secrets[i*2+1]},
		)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to handle connection between Transit Gateway %s and "+
					"Customer Gateway %s due to: %w",
				transitGatewayID, cgw, err,
			)
		}
		vpnConnections[i] = connection
	}

	c.logger.Debug("VPN Connections Covered")
	return vpnConnections, nil
}

func (c *AWSConnector) deleteVPNConnectionsForConnection(ctx context.Context) error {
	c.logger.Debug("Looking for VPN Connections that should be deleted")
	connections, err := c.awsClient.ListVPNConnections(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot get existing VPN Connections due to: %w", err,
		)
	}

	for _, conn := range connections {
		if !c.shouldResourceBeDeleted(conn.Tags) {
			continue
		}
		c.logger.Debugf("Found VPN Connection to delete '%s'", conn.ID)
		if err = c.awsClient.DeleteVPNConnection(ctx, conn.ID); err != nil {
			return fmt.Errorf("failed to clean up VPN Connection: %w", err)
		}
		c.logger.Debugf("Successfully deleted VPN Connection '%s'", conn.ID)
	}

	c.logger.Debug("VPN Connections are cleared")
	return nil
}
