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

package azure

import (
	"context"
	"errors"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/azure/client"
	"github.com/app-net-interface/awi-infra-guard/connector/cidrpool"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

func getIPsFromPools(pools []*cidrpool.CIDRV4Pool, numberOfTunnels uint8) ([]string, error) {
	ips := make([]string, numberOfTunnels)

	if len(pools) != int(numberOfTunnels) {
		return nil, fmt.Errorf(
			"the number of CIDR Pools %d doesn't match the number of tunnels %d "+
				"to be created. Each IP is supposed to be generated from a different "+
				"pool reserved for that particular IPSec Tunnel",
			len(pools), numberOfTunnels,
		)
	}

	for i := uint8(0); i < numberOfTunnels; i++ {
		if pools[i] == nil {
			return nil, errors.New(
				"unexpected nil pool",
			)
		}
		ip, err := pools[i].GetIP()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to generate IP for BGP: %w", err,
			)
		}
		if ip == "" {
			return nil, fmt.Errorf(
				"failed to generate IP for BGP. The pool is full",
			)
		}
		ips[i] = ip
	}

	return ips, nil
}

// TODO: getCustomBGPAddresses returns pair of first custom BGP Addresses.
// It won't work if Vnet Gateway acts as a gateway for more than one
// connection.
func (c *AzureConnector) getCustomBGPAddresses(gw client.VNetGateway) ([2]string, error) {
	// TODO: Attach check that VNet Gateway was already updated
	// TODO: Cover checking if there are two BGP Addresses created.
	if len(gw.BGPAddresses) != 2 {
		return [2]string{"", ""}, fmt.Errorf(
			"unexpected number of Vnet Gateway BGP Interfaces. Expected 2 Interfaces but got %v",
			gw.BGPAddresses,
		)
	}

	if len(gw.BGPAddresses[0].CustomAddresses) > 0 &&
		len(gw.BGPAddresses[1].CustomAddresses) > 0 {

		return [2]string{
			gw.BGPAddresses[0].CustomAddresses[0],
			gw.BGPAddresses[1].CustomAddresses[0],
		}, nil
	}
	return [2]string{"", ""}, nil
}

func (c *AzureConnector) AddBGPAddressesToVnetGateway(
	ctx context.Context,
	cidrPools []*cidrpool.CIDRV4Pool,
	gatewayName string,
	numberOfTunnels uint8,
) error {
	c.logger.Debug("Adding BGP Addresses to Vnet Gateway")

	gw, err := c.azClient.GetVNetGateway(ctx, gatewayName)
	if err != nil {
		return fmt.Errorf(
			"failed to get VPN Gateway %s: %w",
			gatewayName, err,
		)
	}

	// TODO: Fix checking if Custom BGP Addresses are already
	// assigned.
	//
	// bgpAddresses, err := c.getCustomBGPAddresses(gw)
	// if err != nil {
	// 	return fmt.Errorf(
	// 		"failed to get currently assigned Custom GBP Addresses to VNet Gateway %s: %w",
	// 		gatewayName, err,
	// 	)
	// }

	// if len(gw.BGPAddresses) != 0 {
	// 	c.logger.Infof(
	// 		"BGP IP Addresses already assigned. "+
	// 			"First custom BGP IP Address: %s."+
	// 			"Second custom BGP IP Address: %s",
	// 		bgpAddresses[0], bgpAddresses[1],
	// 	)
	// 	return nil
	// }

	if len(c.transaction.BGPAddresses) == 0 {
		ips, err := getIPsFromPools(cidrPools, numberOfTunnels)
		if err != nil {
			return fmt.Errorf(
				"failed to generate IP Addresses for BGP: %w",
				err,
			)
		}
		c.transaction.BGPAddresses = ips
	}

	for i := 0; i < 2; i++ {
		// TODO: Decide how to handle this.
		// gw.BGPAddresses[i].CustomAddresses = append(gw.BGPAddresses[i].CustomAddresses, addr)
		if len(c.transaction.BGPAddresses) == 2 {
			gw.BGPAddresses[i].CustomAddresses = []string{c.transaction.BGPAddresses[i]}
		} else if len(c.transaction.BGPAddresses) == 4 {
			gw.BGPAddresses[i].CustomAddresses = []string{
				c.transaction.BGPAddresses[i*2],
				c.transaction.BGPAddresses[i*2+1],
			}
		} else {
			return fmt.Errorf("NOT IMPLEMENTED - unhandled situation")
		}
	}

	if err = c.azClient.UpdateVNetGateway(ctx, gw); err != nil {
		return fmt.Errorf(
			"failed to update VPN Gateway %v with BGP Addresses: %w",
			gatewayName, err,
		)
	}

	return nil
}

func (c *AzureConnector) createLocalVNetGateways(
	ctx context.Context,
	config types.CreateBGPConnectionConfig,
	gatewayName string,
) ([]string, error) {
	c.logger.Debug("Creating Local Network Gateways")

	gw, err := c.azClient.GetVNetGateway(ctx, gatewayName)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get VPN Gateway %s: %w",
			gatewayName, err,
		)
	}

	gatewayNames := make([]string, len(config.OutsideInterfaces))
	for i := uint8(0); i < config.NumberOfTunnels; i++ {
		localGW := client.LocalVNetGateway{
			Name:             c.GenerateName(fmt.Sprintf("localvnetgw-%d", i)),
			GatewayIP:        config.OutsideInterfaces[i],
			ASN:              fmt.Sprintf("%d", config.PeerASN),
			NetworkAddresses: gw.Addresses,
		}
		localGW.PeerBGPAddresses = append(localGW.PeerBGPAddresses,
			client.BGPAddress{
				ConfigurationID: c.GenerateName(fmt.Sprintf("locvnetgw-ipconf-%d", i)),
				CustomAddresses: []string{config.PeerBGPAddresses[i]},
			})
		if err := c.azClient.CreateLocalVNetGateway(ctx, localGW); err != nil {
			return nil, fmt.Errorf(
				"failed to create Local Network Gateway: %w", err,
			)
		}
		gatewayNames[i] = localGW.Name
	}

	return gatewayNames, nil
}

func (c *AzureConnector) deleteLocalVNetGateways(ctx context.Context) error {
	c.logger.Debug("Deleting Local Network Gateways")

	gws, err := c.azClient.ListLocalVNetGateways(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot obtain the list of existing Local Network Gateways: %w", err)
	}
	for _, gw := range gws {
		if !c.IsNameOwnedByConnection(gw.Name) {
			continue
		}
		c.logger.Debugf("found Local Network Gateways to delete: %v", gw)
		if err = c.azClient.DeleteLocalVNetGateway(ctx, gw.Name); err != nil {
			return fmt.Errorf(
				"failed to remove Local Network Gateways %v due to: %w", gw, err,
			)
		}
		c.logger.Debugf("Local Network Gateways %v deleted successfully", gw)
	}
	return nil
}

func (c *AzureConnector) createNetworkGatewayConnection(
	ctx context.Context,
	config types.CreateBGPConnectionConfig,
	vpnGateway string,
	localGateways []string,
) error {
	localGatewaysIDs := make([]string, len(localGateways))

	vpnGW, err := c.azClient.GetVNetGateway(ctx, vpnGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to retrieve Virtual Network Gateway ID: %s", err,
		)
	}

	for i, localGW := range localGateways {
		localNetworkGW, err := c.azClient.GetLocalVNetGateway(ctx, localGW)
		if err != nil {
			return fmt.Errorf(
				"failed to retrieve Local Network Gateway ID: %s", err,
			)
		}
		localGatewaysIDs[i] = localNetworkGW.ID
	}

	for i, gwID := range localGatewaysIDs {
		enabled := true
		err := c.azClient.CreateNetworkGatewayConnection(
			ctx,
			client.NetworkGatewayConnection{
				Name:                  c.GenerateName(fmt.Sprintf("netgwconn-%d", i)),
				SharedKey:             config.SharedSecrets[i],
				BGPEnabled:            &enabled,
				NetworkGatewayID:      vpnGW.ID,
				LocalNetworkGatewayID: gwID,
				ConnectionType:        "IPsec",
				ConnectionProtocol:    "IKEv2",
			},
		)
		if err != nil {
			return fmt.Errorf(
				"failed to create Network Gateway Connection: %w", err,
			)
		}
	}

	return nil
}

func (c *AzureConnector) deleteNetworkGatewayConnections(ctx context.Context) error {
	c.logger.Debug("Deleting Network Gateway Connections")
	connections, err := c.azClient.ListNetworkGatewayConnections(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot obtain the list of existing Network Gateway Connections: %w", err)
	}
	for _, conn := range connections {
		c.logger.Tracef("Found existing Network Gateway Connection: %v", conn.Name)
		if !c.IsNameOwnedByConnection(conn.Name) {
			c.logger.Trace("Not owned")
			continue
		}
		c.logger.Debugf("found Network Gateway Connections to delete: %v", conn)
		if err = c.azClient.DeleteNetworkGatewayConnection(ctx, conn.Name); err != nil {
			return fmt.Errorf(
				"failed to remove Network Gateway Connections %v due to: %w", conn, err,
			)
		}
		c.logger.Debugf("Network Gateway Connections %v deleted successfully", conn)
	}
	return nil
}
