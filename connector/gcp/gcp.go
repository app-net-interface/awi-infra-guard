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

package gcp

import (
	"context"
	"errors"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/cidrpool"
	"github.com/app-net-interface/awi-infra-guard/connector/gcp/client"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// GCPConnector is a structure responsible for preparing GCP
// side of the connection between providers.
type GCPConnector struct {
	gcpClient *client.Client
	config    *Config
	logger    *logrus.Entry
	// TODO: Create a map for a state per Connection ID.
	// If the CSP Connector will become a service then
	// multiple states can exist at the same time.
	state *transactionState
}

func NewConnector(ctx context.Context, logger *logrus.Entry, config string) (provider.Provider, error) {
	parsedConfig := Config{}
	if err := yaml.Unmarshal([]byte(config), &parsedConfig); err != nil {
		return nil, fmt.Errorf("could not parse Configuration for AWS Client: %w", err)
	}

	gcpClient, err := client.NewClient(ctx, logger.WithField("logger", "gcp_client"), parsedConfig.Project)
	if err != nil {
		return nil, fmt.Errorf("could not create GCP Client: %w", err)
	}
	return &GCPConnector{
		gcpClient: gcpClient,
		config:    &parsedConfig,
		logger:    logger,
		state:     &transactionState{},
	}, nil
}

func (c *GCPConnector) Name() string {
	return "gcp"
}

// Close closes all internal clients for this client.
//
// The user should call it as soon as the connector is no
// longer needed
func (c *GCPConnector) Close() error {
	if c.gcpClient != nil {
		if err := c.gcpClient.Close(); err != nil {
			return fmt.Errorf("failed to Close the gcpClient: %w", err)
		}
		c.gcpClient = nil
	}
	c.logger.Debug("The connector has been closed successfully")
	return nil
}

// ListGateways returns the slice of abstract Gateway objects that can
// be used as Edge for connection with other Cloud Providers.
//
// Currently, the GCP Provider considers Cloud Routers as Gateways.
// However, for the Cloud Router to be an actual Gateway, it needs
// to have a VPN Gateway and a Network associated to it.
func (c *GCPConnector) ListGateways(ctx context.Context) ([]types.Gateway, error) {
	routers, err := c.gcpClient.ListRouters(ctx, c.config.Region)
	if err != nil {
		return nil, fmt.Errorf("cannot obtain a list of Cloud Routers from GCP: %w", err)
	}
	gateways := []types.Gateway{}
	for _, r := range routers {
		asn := ""
		if r.BGP != nil {
			asn = r.BGP.ASN
		}
		gateway := types.Gateway{
			Name:          r.Name,
			Kind:          "CloudRouter",
			CloudProvider: c.Name(),
			VPC:           r.Network,
			ASN:           asn,
			// TODO: Handle regions properly when multiple regions are supported.
			Region: c.config.Region,
		}
		gateways = append(gateways, gateway)
	}
	return gateways, nil
}

func (c *GCPConnector) GetGatewayConnectionSettings(ctx context.Context, gateway types.Gateway) (
	types.GatewayConnectionSettings, error,
) {
	return types.GatewayConnectionSettings{
		NumberOfInterfaces: 2,
		BGPSetting: &types.BGPSetting{
			Addressing: types.BGPAddressing{
				AcceptsBothAddresses: true,
				// TODO: Add an implementation for generating
				// both IP Addresses.
				GeneratesBothAddresses:            false,
				GeneratesOwnAndAcceptsPeerAddress: true,
			},
			AllowedIPRanges: []string{"169.254.0.0/16"},
			// TODO: Add excluded IP Ranges that are disallowed by
			// the GCP provider.
			ExcludedIPRanges: []string{},
		},
	}, nil
}

func (c *GCPConnector) InitializeCreation(
	ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
) error {
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gateway.Name)
	if err != nil {
		return fmt.Errorf(
			"failed to get Cloud Router %s: %w", gateway.Name, err,
		)
	}
	if router == nil {
		return fmt.Errorf(
			"the Cloud Router %s doesn't exist", gateway.Name,
		)
	}

	// The GCP Gateway is expected to already have Interfaces created, which is why
	// obtaining them occurs during InitializeCreation step.
	firstInterface, secondInterface, vpnGateway, err := c.GetInterfaces(ctx, gateway.Name)
	if err != nil {
		return fmt.Errorf("cannot initialize gateway interfaces: %w", err)
	}

	c.state = NewTransactionState(
		router.URL,
		vpnGateway,
		gateway.Name,
		peerGateway.Name,
		[]string{firstInterface, secondInterface},
	)
	return nil
}

func (c *GCPConnector) InitializeGatewayInterfaces(
	ctx context.Context, gateway, peerGateway types.Gateway,
) ([]string, error) {
	return c.state.OwnInterfaces, nil
}

func (c *GCPConnector) InitializeASN(
	ctx context.Context, gateway, peerGateway types.Gateway,
) (uint64, error) {
	asn := helper.StringToUInt64Pointer(gateway.ASN)
	if asn == nil {
		return 0, fmt.Errorf("failed to parse ASN number '%v'", asn)
	}
	return *asn, nil
}

func (c *GCPConnector) GetVPCForGateway(ctx context.Context, gateway types.Gateway) (string, error) {
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gateway.Name)
	if err != nil {
		return "", fmt.Errorf("failed to get GCP Cloud Router %s: %w", gateway.Name, err)
	}
	return router.Network, nil
}

func (c *GCPConnector) AttachToExternalGatewayWithStaticRouting() error {
	return errors.New("AttachToExternalGatewayWithStaticRouting not implemented in gcp")
}

func (c *GCPConnector) getOperations(
	attachMode types.AttachBGPConnectionMode,
) ([]StateOperation, error) {
	// The difference in operations for attach mode is mostly
	// for the Cloud Router BGP Peers.
	//
	// If the method was called for method AttachModeGenerateIP
	// it means that this is the beginning of creating connection
	// with cooperative mode which means that at this stage peering
	// BGP Addresseses won't be available.
	switch attachMode {
	case types.AttachModeGenerateIPAndAcceptOtherIP, types.AttachModeAcceptOtherIP:
		return []StateOperation{
			c.CreateExternalVPNGateway,
			c.CreateVPNTunnels,
			c.CreateCloudRouterInterfaces,
			c.CreateCloudRouterBGPPeers,
		}, nil
	case types.AttachModeGenerateIP:
		return []StateOperation{
			c.CreateExternalVPNGateway,
			c.CreateVPNTunnels,
			c.CreateCloudRouterInterfaces,
		}, nil
	}
	return nil, fmt.Errorf(
		"failed to get list of operations. Unhandled attachMode: %v", attachMode,
	)
}

func (c *GCPConnector) generateBGPIPs(pools []*cidrpool.CIDRV4Pool) ([]string, error) {
	numberOfIPs := c.state.BGPConnectionConfig.NumberOfTunnels

	ips := make([]string, numberOfIPs)

	for i := range ips {
		if pools[i] == nil {
			return nil, errors.New(
				"unexpected nil pool",
			)
		}
		ip, err := pools[i].GetIP()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to generate IP from CIDR Pool for BGP: %w", err,
			)
		}
		if ip == "" {
			return nil, errors.New(
				"failed to generate IP from CIDR Pool for BGP. The CIDR Pool is full",
			)
		}
		ips[i] = ip
	}

	return ips, nil
}

func (c *GCPConnector) AttachToExternalGatewayWithBGP(
	ctx context.Context,
	gateway, peerGateway types.Gateway,
	attachMode types.AttachBGPConnectionMode,
	config types.CreateBGPConnectionConfig,
) (types.OutputForConnectionWithBGP, error) {
	c.logger.Infof(
		"Starting the preparation of GCP Resources for establishing the connection. "+
			"Selected Cloud Router '%s'.",
		gateway.Name,
	)

	if attachMode == types.AttachModeGenerateBothIPs {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"the GCP Provider does not support Attach Mode Generate Both IPs. " +
				"It seems that there is an implementation bug - such situation " +
				"should not happen",
		)
	}

	c.state.BGPConnectionConfig = config

	output := types.OutputForConnectionWithBGP{
		Interfaces:       c.state.OwnInterfaces,
		PeerBGPAddresses: config.BGPAddresses,
		BGPAddresses:     config.PeerBGPAddresses,
		SharedSecrets:    config.SharedSecrets,
	}

	if attachMode == types.AttachModeGenerateIP || attachMode == types.AttachModeGenerateIPAndAcceptOtherIP {
		ips, err := c.generateBGPIPs(c.state.BGPConnectionConfig.BGPCIDRPools)
		if err != nil {
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to generate IP Addresses for BGP Session: %w", err,
			)
		}
		c.state.OwnBGPAddresses = ips
		output.BGPAddresses = ips
	} else if attachMode == types.AttachModeAcceptOtherIP {
		if len(c.state.OwnBGPAddresses) == 0 {
			c.state.OwnBGPAddresses = config.BGPAddresses
		}
	}

	operations, err := c.getOperations(attachMode)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to obtain list of operations to do: %w", err,
		)
	}

	for _, doOperation := range operations {
		if err = doOperation(ctx); err != nil {
			// TODO: Add a name of the operation which failed.
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to create a connection: %v",
				err,
			)
		}
		c.state.state, err = getNextState(c.state.state)
		if err != nil {
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to update creating connection state: %v",
				err,
			)
		}
	}

	return output, nil
}

func (c *GCPConnector) DeleteConnectionResources(
	ctx context.Context,
	gateway types.Gateway,
	peerGateway types.Gateway,
) error {
	c.state.GatewayName = gateway.Name
	c.state.PeerGatewayName = peerGateway.Name
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gateway.Name)
	if err != nil {
		return fmt.Errorf(
			"could not get GCP CloudRouter with ID '%s': %w",
			gateway.Name, err)
	}
	if router == nil {
		return fmt.Errorf(
			"got empty Cloud Router object for ID '%s'", gateway.Name)
	}

	if err = c.deleteCloudRouterBGPPeersForConnection(ctx, *router); err != nil {
		return fmt.Errorf(
			"failed to clean up BGP Peers for Cloud Router '%s': %w",
			gateway.Name, err,
		)
	}

	if err = c.deleteCloudRouterInterfacesForConnection(ctx, *router); err != nil {
		return fmt.Errorf(
			"failed to clean up Interfaces for Cloud Router '%s': %w",
			gateway.Name, err,
		)
	}

	if err = c.deleteVPNTunnelsForConnection(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up VPN Tunnels for Cloud Router '%s': %w",
			gateway.Name, err,
		)
	}

	if err = c.deleteExternalVPNGatewaysForConnection(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up External VPN Gateways for Cloud Router '%s': %w",
			gateway.Name, err,
		)
	}

	c.logger.Infof("all resources removed successully")
	return nil
}

func (c *GCPConnector) deleteCloudRouterBGPPeersForConnection(
	ctx context.Context, cloudRouter client.CloudRouter,
) error {
	peersToDelete := []string{}
	for _, peer := range cloudRouter.BGPPeers {
		if c.IsNameOwnedByConnection(peer.Name) {
			c.logger.Debugf("found BGP Peer to delete: %v", peer)
			peersToDelete = append(peersToDelete, peer.Name)
		}
	}
	for _, peerName := range peersToDelete {
		err := c.gcpClient.DeleteRouterBGPPeer(
			ctx, c.config.Region, cloudRouter.Name, peerName)
		if err != nil {
			return fmt.Errorf(
				"failed to remove BGP Peer '%s' from Cloud Router '%s'",
				peerName, cloudRouter.Name,
			)
		}
	}
	c.logger.Debugf(
		"removed following BGP Peers '%v' from Cloud Router '%s'",
		peersToDelete, cloudRouter.Name)
	return nil
}

func (c *GCPConnector) deleteCloudRouterInterfacesForConnection(
	ctx context.Context, cloudRouter client.CloudRouter,
) error {
	interfacesToDelete := []string{}
	for _, iface := range cloudRouter.Interfaces {
		if c.IsNameOwnedByConnection(iface.Name) {
			c.logger.Debugf("found Interface to delete: %v", iface)
			interfacesToDelete = append(interfacesToDelete, iface.Name)
		}
	}
	for _, interfaceName := range interfacesToDelete {
		err := c.gcpClient.DeleteRouterInterface(
			ctx, c.config.Region, cloudRouter.Name, interfaceName)
		if err != nil {
			return fmt.Errorf(
				"failed to remove Interface '%s' from Cloud Router '%s'",
				interfaceName, cloudRouter.Name,
			)
		}
	}
	c.logger.Debugf(
		"removed following Interfaces '%v' from Cloud Router '%s'",
		interfacesToDelete, cloudRouter.Name)
	return nil
}

func (c *GCPConnector) deleteVPNTunnelsForConnection(ctx context.Context) error {
	tunnels, err := c.gcpClient.ListVPNTunnels(ctx, c.config.Region)
	if err != nil {
		return fmt.Errorf(
			"cannot obtain the list of existing VPN Tunnels: %w", err)
	}
	for i := range tunnels {
		if tunnels[i] == nil {
			continue
		}
		if !c.IsNameOwnedByConnection(tunnels[i].Name) {
			continue
		}
		c.logger.Debugf("found VPN Tunnel to delete: %v", tunnels[i])
		if err = c.gcpClient.DeleteVPNTunnel(ctx, c.config.Region, tunnels[i].Name); err != nil {
			return fmt.Errorf(
				"failed to remove VPN Tunnel %v due to: %w", tunnels[i], err,
			)
		}
		c.logger.Debugf("VPN Tunnel %v deleted successfully", tunnels[i])
	}
	return nil
}

func (c *GCPConnector) deleteExternalVPNGatewaysForConnection(ctx context.Context) error {
	gateways, err := c.gcpClient.ListExternalVPNGateway(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot obtain the list of existing External VPN Gateways: %w", err)
	}
	for i := range gateways {
		if gateways[i] == nil {
			continue
		}
		if !c.IsNameOwnedByConnection(gateways[i].Name) {
			continue
		}
		c.logger.Debugf("found External VPN Gateway to delete: %v", gateways[i])
		if err = c.gcpClient.DeleteExternalVPNGateway(ctx, gateways[i].Name); err != nil {
			return fmt.Errorf(
				"failed to remove External VPN Gateway %v due to: %w", gateways[i], err,
			)
		}
		c.logger.Debugf("External VPN Gateway %v deleted successfully", gateways[i])
	}
	return nil
}

// GetInterfaces returns two interfaces that can be used for creating the connection
// along with the name of VPN Gateway they come from.
//
// The interfaces are picked from the VPN Gateway associated with the GCP Cloud Router
// that serves the role of the Gateway for creating the connection.
func (c *GCPConnector) GetInterfaces(ctx context.Context, routerID string) (
	firstInterface, secondInterface, vpnGateway string, err error,
) {
	_, vpnGateways, err := c.getGCPRouterWithAttachedVPNGateways(ctx, routerID)
	if err != nil {
		return "", "", "", fmt.Errorf(
			"cannot get GCP Cloud Router and its VPN Gateways to find out available "+
				"public IPs for GCP Gateway '%s': %w", routerID, err)
	}

	if len(vpnGateways) == 0 {
		return "", "", "", fmt.Errorf(
			"cannot use GCP Gateway '%s' to connect with other CSP as it does not expose "+
				"any public IP Addresses that could be used as interface. The corresponding "+
				"Cloud Router needs to be within the Network where VPN Gateways are created",
			routerID)
	}

	// TODO: Implement a mechanism to keep track of VPN Gateways that are already
	// a part of some connection
	gw := vpnGateways[0]

	if len(gw.IPAddresses) < 2 {
		return "", "", "", fmt.Errorf(
			"expected at least 2 public IP Addresses but the found VPNGateway '%s' has only '%d': %v",
			gw.URL, len(gw.IPAddresses), gw.IPAddresses,
		)
	}

	return gw.IPAddresses[0], gw.IPAddresses[1], gw.URL, nil
}

// GetGateway returns the Cloud Router object, considered as currently
// selected Gateway, based on the Connection options.
func (c *GCPConnector) GetGateway(ctx context.Context, gateway types.GatewayIdentifier) (*types.Gateway, error) {
	gcpRouter, err := c.gcpClient.GetRouter(
		ctx,
		gateway.Region,
		gateway.GatewayID)
	if err != nil {
		return nil, fmt.Errorf("the GCP Cloud Router '%s' was not found: %w", gateway.GatewayID, err)
	}
	if gcpRouter == nil {
		return nil, fmt.Errorf("the GCP Cloud Router '%s' object is empty", gateway.GatewayID)
	}
	asn := ""
	if gcpRouter.BGP != nil {
		asn = gcpRouter.BGP.ASN
	}
	gw := types.Gateway{
		Name:          gcpRouter.Name,
		Kind:          "CloudRouter",
		CloudProvider: gateway.Provider,
		VPC:           gcpRouter.Network,
		ASN:           asn,
		Region:        gateway.Region,
	}
	return &gw, nil
}

func (c *GCPConnector) getGCPRouterWithAttachedVPNGateways(
	ctx context.Context,
	gatewayName string,
) (*client.CloudRouter, []*client.VPNGateway, error) {
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gatewayName)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"could not get GCP CloudRouter with ID '%s': %w",
			gatewayName, err)
	}
	if router == nil {
		return nil, nil, fmt.Errorf(
			"got empty Cloud Router object for ID '%s'", gatewayName)
	}
	if router.Network == "" {
		return router, nil, nil
	}
	vpnGateways, err := c.gcpClient.FindVPNGatewaysForNetwork(
		ctx, c.config.Region, router.Network)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"could not get the list of VPN Gateways for Cloud Router '%s': %w",
			gatewayName, err)
	}
	return router, vpnGateways, nil
}

func (c *GCPConnector) GetCIDRs(ctx context.Context, gateway types.Gateway) ([]string, error) {
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gateway.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get GCP Cloud Router %s: %w", gateway.Name, err)
	}
	if router == nil {
		return nil, fmt.Errorf(
			"failed to get GCP Cloud Router %s. Doesn't exist: %w",
			gateway.Name, err)
	}
	if router.Network == "" {
		c.logger.Debugf(
			"Cloud Router %s has no VPC associated: no CIDRs found",
			gateway.Name)
		return nil, nil
	}
	network, err := c.gcpClient.GetNetwork(ctx, router.Network)
	if err != nil {
		return nil, fmt.Errorf("failed to get Network %s: %w", router.Network, err)
	}
	if network == nil {
		return nil, fmt.Errorf("failed to get Network %s: got empty Network object", router.Network)
	}

	cidrs := []string{}
	for _, subnetworkID := range network.Subnets {
		subnetwork, err := c.gcpClient.GetSubnetwork(ctx, c.config.Region, subnetworkID)
		if err != nil {
			return nil, fmt.Errorf("failed to get Subnetwork %s: %w", subnetwork, err)
		}
		if subnetwork == nil {
			return nil, fmt.Errorf("failed to get Subnetwork %s: got empty Subnetwork object", subnetwork)
		}
		cidrs = append(cidrs, subnetwork.CIDR)
	}

	return cidrs, nil
}
