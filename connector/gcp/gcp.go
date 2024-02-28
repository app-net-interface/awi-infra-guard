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
	"errors"
	"fmt"
	"strconv"

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
	state *transactionState
}

func NewConnector(ctx context.Context, logger *logrus.Entry, config string) (provider.Provider, error) {
	parsedConfig := Config{}
	if err := yaml.Unmarshal([]byte(config), &parsedConfig); err != nil {
		return nil, fmt.Errorf("could not parse Configuration for AWS Client: %w", err)
	}

	gcpClient, err := client.NewClient(ctx, logger.WithField("logger", "gcp_client"), parsedConfig.Project)
	if err != nil {
		return nil, fmt.Errorf("Could not create GCP Client: %w", err)
	}
	return &GCPConnector{
		gcpClient: gcpClient,
		config:    &parsedConfig,
		logger:    logger,
		state:     &transactionState{},
	}, nil
}

type transactionState struct {
	VPNGatewayURL   string
	GatewayName     string
	PeerGatewayName string
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
			return fmt.Errorf("Failed to Close the gcpClient: %w", err)
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
				GeneratesBothAddresses: false,
				// TODO: Add an implementation for generating
				// own IP Address and accepting the second
				// address.
				//
				// This can be used for creating a connection
				// with Azure provider.
				GeneratesOwnAndAcceptsPeerAddress: false,
			},
			AllowedIPRanges: []string{"192.254.0.0/16"},
			// TODO: Add excluded IP Ranges that are disallowed by
			// the GCP provider.
			ExcludedIPRanges: []string{},
		},
	}, nil
}

func (c *GCPConnector) InitializeGatewayInterfaces(
	ctx context.Context, gateway, peerGateway types.Gateway,
) ([]string, error) {
	firstInterface, secondInterface, vpnGateway, err := c.GetInterfaces(ctx, gateway.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize gateway interfaces: %w", err)
	}
	c.state.VPNGatewayURL = vpnGateway
	c.state.GatewayName = gateway.Name
	c.state.PeerGatewayName = peerGateway.Name
	return []string{firstInterface, secondInterface}, nil
}

func (c *GCPConnector) InitializeASN(
	ctx context.Context, gateway, peerGateway types.Gateway,
) (uint64, error) {
	asn, err := strconv.Atoi(gateway.ASN)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ASN number: %w", err)
	}
	// TODO: Fix type casting (int -> uint64 - possible loss)
	return uint64(asn), nil
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

func (c *GCPConnector) AttachToExternalGatewayWithBGP(
	ctx context.Context,
	gateway, peerGateway types.Gateway,
	attachMode types.AttachBGPConnectionMode,
	config types.CreateBGPConnectionConfig,
) (types.OutputForConnectionWithBGP, error) {
	if attachMode != types.AttachModeAcceptOtherIP {
		return types.OutputForConnectionWithBGP{}, errors.New(
			"currently provider gcp doesn't support BGP mode other than AttachModeAcceptOtherIP",
		)
	}

	c.logger.Infof(
		"Starting the preparation of GCP Resources for establishing the connection. "+
			"Selected Cloud Router '%s'.",
		gateway.Name,
	)

	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, gateway.Name)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to get Cloud Router %s: %w", gateway.Name, err,
		)
	}
	if router == nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"Cloud Router %s doesn't exist", gateway.Name,
		)
	}
	cloudRouterURL := router.URL

	extVPNGateway, err := c.createExternalVPNGatewayIfNeeded(ctx, config)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to prepare External VPN Gateway: %w", err,
		)
	}
	if extVPNGateway == nil {
		return types.OutputForConnectionWithBGP{}, errors.New(
			"Unexpected empty External VPN Gateway object",
		)
	}
	c.logger.Debugf("Using External VPN Gateway: %v", *extVPNGateway)

	vpnTunnels, err := c.createVPNTunnelsIfNeeded(
		ctx, config, extVPNGateway.URL, c.state.VPNGatewayURL, cloudRouterURL)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to prepare VPN Tunnels: %w", err,
		)
	}
	if len(vpnTunnels) == 0 {
		return types.OutputForConnectionWithBGP{}, errors.New(
			"expected VPN tunnels created. Got nothing",
		)
	}
	c.logger.Debugf("Using following VPN Tunnels: %v", vpnTunnels)

	cloudRouterInterfaces, err := c.createCloudRouterInterfacesIfNeeded(
		ctx, config, vpnTunnels,
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to prepare Cloud Router Interfaces: %w", err,
		)
	}
	if cloudRouterInterfaces == nil {
		return types.OutputForConnectionWithBGP{}, errors.New(
			"expected Cloud Router interfaces. Got nothing",
		)
	}

	_, err = c.createCloudRouterBPGPeersIfNeeded(
		ctx, config, cloudRouterInterfaces,
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to prepare Cloud Router BGP Peers: %w", err,
		)
	}

	return types.OutputForConnectionWithBGP{}, nil
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
// that serves the role of the Gateway for creating the connection. The configuration
// allows specifying the exact VPN Gateway name to decide which pair of IPs should be
// used as interfaces. Another option allows picking random VPN Gateway to avoid the
// need of manual checking for VPN Gateways available for the Cloud Router. If both
// options are not specified, the VPN Gateway can still be picked if there is only
// one available (otherwise the Connector script won't make the decision itself).
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
		return nil, fmt.Errorf("The GCP Cloud Router '%s' was not found: %w", gateway.GatewayID, err)
	}
	if gcpRouter == nil {
		return nil, fmt.Errorf("The GCP Cloud Router '%s' object is empty", gateway.GatewayID)
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

func (c *GCPConnector) createExternalVPNGatewayIfNeeded(
	ctx context.Context, connectConfig types.CreateBGPConnectionConfig,
) (*client.ExternalVPNGateway, error) {
	c.logger.Debug("Starting to prepare External VPN Gateway")

	interfaces := connectConfig.OutsideInterfaces
	c.logger.Debugf(
		"The External VPN Gateway is supposed to have following interfaces: %v",
		interfaces)

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
		Name:       c.GenerateName(""),
		Interfaces: []client.ExternalVPNGatewayInterface{},
	}
	for _, iface := range interfaces {
		gateway.Interfaces = append(gateway.Interfaces, client.ExternalVPNGatewayInterface{
			IP: iface,
		})
	}
	gateway, err = c.gcpClient.CreateExternalVPNGateway(ctx, gateway)
	if err != nil {
		return nil, fmt.Errorf("failed to create External VPN Gateway %v: %w", *gateway, err)
	}
	c.logger.Debugf(
		"External VPN Gateway '%v' created", *gateway)
	return gateway, nil
}

func (c *GCPConnector) doesVPNTunnelMatchesCriteria(requestedTunnel client.VPNTunnel, actualTunnel client.VPNTunnel) bool {
	// Removing identifiers as they will most likely differ
	requestedTunnel.Name = ""
	requestedTunnel.URL = ""
	actualTunnel.Name = ""
	actualTunnel.URL = ""

	// Removing Shared Key as the downloaded one has it encrypted
	requestedTunnel.SharedSecret = ""
	actualTunnel.SharedSecret = ""

	c.logger.Debugf(
		"Comparing Desired VPN Tunnel %v with found VPN Tunnel %v", requestedTunnel, actualTunnel,
	)

	return requestedTunnel == actualTunnel
}

// Returns the slice of 4 pointers to the matching VPN Tunnels.
//
// The order of pointers matters. It corresponds with interfaces 0, 1, 2 and 3.
// If any pointer is nil, it means that the corresponding VPN Tunnel was not
// found and needs to be created.
func (c *GCPConnector) getMatchingVPNTunnels(
	ctx context.Context, desiredTunnels []client.VPNTunnel,
) ([]*client.VPNTunnel, error) {
	matchedTunnels := []*client.VPNTunnel{nil, nil, nil, nil}

	tunnels, err := c.gcpClient.ListVPNTunnels(ctx, c.config.Region)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the list of VPN Tunnels: %w", err)
	}

	for i := range tunnels {
		if tunnels[i] == nil {
			continue
		}
		for j := range desiredTunnels {
			if !c.doesVPNTunnelMatchesCriteria(desiredTunnels[j], *tunnels[i]) {
				continue
			}
			if matchedTunnels[j] == nil {
				c.logger.Debugf("Found matching VPN Tunnel for interface %d: %s", j, *&tunnels[i].Name)
				matchedTunnels[j] = tunnels[i]
				continue
			}
			c.logger.Warnf(
				"Found another VPN Tunnel for interface %d: %s. "+
					"Ignoring it and keeping previously found '%s' VPN Tunnel.",
				j, *&tunnels[i].Name, matchedTunnels[j].Name)
		}
	}

	return matchedTunnels, nil
}

func (c *GCPConnector) getDesiredVPNTunnels(
	connectConfig types.CreateBGPConnectionConfig, externalVPNGatewayURL, vpnGatewayURL, cloudRouterURL string,
) []client.VPNTunnel {
	tunnels := []client.VPNTunnel{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			tunnels = append(tunnels, client.VPNTunnel{
				Name: c.GenerateName(
					fmt.Sprintf("tunnel-%d-", ((i*2)+j)+1)),
				ExternalVPNGateway:       externalVPNGatewayURL,
				ExternalGatewayInterface: (i * 2) + j,
				VPNGateway:               vpnGatewayURL,
				IKEVersion:               2,
				Interface:                i,
				CloudRouter:              cloudRouterURL,
				SharedSecret:             connectConfig.SharedSecrets[(i*2)+j],
			})
		}
	}

	return tunnels
}

func (c *GCPConnector) createVPNTunnelsIfNeeded(
	ctx context.Context,
	connectConfig types.CreateBGPConnectionConfig,
	externalVPNGatewayURL,
	vpnGatewayURL,
	cloudRouterURL string,
) ([]client.VPNTunnel, error) {
	c.logger.Debug("Starting to prepare VPN Tunnels")

	desiredVPNTunnels := c.getDesiredVPNTunnels(
		connectConfig, externalVPNGatewayURL, vpnGatewayURL, cloudRouterURL)
	c.logger.Debugf(
		"For the connection, the four following VPN Tunnels should exist. %v."+
			"Names are irrelevant: they will be used if vpn tunnels are missing and need "+
			"to be created", desiredVPNTunnels)

	matchedVPNTunnels, err := c.getMatchingVPNTunnels(ctx, desiredVPNTunnels)
	if err != nil {
		return nil, fmt.Errorf("failed to verify currently existing VPN Tunnels: %w", err)
	}
	c.logger.Debugf(
		"For the connection, the following VPN tunnels were actually found %v", matchedVPNTunnels)

	vpnTunnels := []client.VPNTunnel{}
	for i := range matchedVPNTunnels {
		if matchedVPNTunnels[i] != nil {
			c.logger.Debugf(
				"VPN Tunnel for interface %d already exists with name '%s'. Using it.",
				i, matchedVPNTunnels[i].Name)
			vpnTunnels = append(vpnTunnels, *matchedVPNTunnels[i])
			continue
		}
		c.logger.Debugf("VPN Tunnel for interface %d not found. Creating it.", i)

		createdTunnel, err := c.gcpClient.CreateVPNTunnel(ctx, &desiredVPNTunnels[i], c.config.Region)
		if err != nil {
			return nil, fmt.Errorf("failed to create a VPN Tunnel %v due to: %w", desiredVPNTunnels[i], err)
		}
		if createdTunnel == nil {
			return nil, fmt.Errorf("got unexpected empty VPN Tunnel after creation. Expected: %v", desiredVPNTunnels[i])
		}
		vpnTunnels = append(vpnTunnels, *createdTunnel)
		c.logger.Debugf("VPN Tunnel %s created successfully.", desiredVPNTunnels[i].Name)
	}

	return vpnTunnels, nil
}

func (c *GCPConnector) getDesiredInterfacesForCloudRouter(
	connectConfig types.CreateBGPConnectionConfig,
	vpnTunnels []client.VPNTunnel,
) ([]client.CloudRouterInterface, error) {
	if len(vpnTunnels) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of provided VPN Tunnels: %v. Expected 4 tunnels",
			vpnTunnels)
	}
	addresses := connectConfig.BGPAddresses
	if len(addresses) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of provided BGP Tunnel IP Addresses: %v. Expected 4 addresses",
			addresses)
	}

	maskSuffix := fmt.Sprintf("/%d", REQUESTED_CIDR_SIZE)

	interfaces := []client.CloudRouterInterface{}
	for i := range vpnTunnels {
		interfaces = append(interfaces, client.CloudRouterInterface{
			Name: c.GenerateName(
				fmt.Sprintf("tunnel-%d-", (i + 1)),
			),
			VPNTunnel: vpnTunnels[i].URL,
			IpRange:   addresses[i] + maskSuffix,
		})
	}
	return interfaces, nil
}

// prepareCloudRouterInterfaces checks for the presence of
// provided desiredInterfaces in the given Cloud Router and returns
// two slices of Interfaces.
//
// The first interface represents the interfaces that will be used
// for Connection creation (it may consist of the mix of already
// existing interfaces and those that need to be created yet).
// The second slice consists of missing interfaces that needs to
// be created afterwards.
func (c *GCPConnector) prepareCloudRouterInterfaces(
	connectConfig types.CreateBGPConnectionConfig,
	cloudRouter *client.CloudRouter,
	desiredInterfaces []client.CloudRouterInterface,
) ([]client.CloudRouterInterface, []client.CloudRouterInterface, error) {
	c.logger.Debugf(
		"checking the presence of desired interfaces '%v' in Cloud Router",
		desiredInterfaces,
	)
	if cloudRouter == nil {
		return nil, nil, errors.New(
			"cannot verify missing interfaces in nil Cloud Router",
		)
	}

	interfaces := []client.CloudRouterInterface{}
	missingInterfaces := []client.CloudRouterInterface{}

	for _, desired := range desiredInterfaces {
		c.logger.Debugf(
			"checking the presence the interface '%v' in Cloud Router '%s'",
			desired, cloudRouter.Name,
		)
		found := false
		for _, existing := range cloudRouter.Interfaces {
			c.logger.Tracef(
				"comparing desired interface '%v' with existing one: '%v'",
				desired, existing,
			)
			if desired.IpRange == existing.IpRange && desired.VPNTunnel == existing.VPNTunnel {
				c.logger.Debugf(
					"interface '%v' found in Cloud Router '%s'. Will use it",
					existing, cloudRouter.Name,
				)
				interfaces = append(interfaces, existing)
				found = true
				break
			}
		}
		if !found {
			c.logger.Debugf(
				"interface '%v' not found in Cloud Router '%s'. Will add one",
				desired, cloudRouter.Name,
			)
			interfaces = append(interfaces, desired)
			missingInterfaces = append(missingInterfaces, desired)
		}
	}

	c.logger.Debugf(
		"The Cloud Router '%s' will expose the following 4 interfaces for connection purposes: %v. "+
			"The following interfaces are not created yet and will be created shortly afterwards: %v",
		cloudRouter.Name, interfaces, missingInterfaces,
	)
	return interfaces, missingInterfaces, nil
}

// Adds required interfaces to the Cloud Router if not present and returns
// the slice of Cloud Router Interfaces that are meant to be used for
// further connection creation.
func (c *GCPConnector) createCloudRouterInterfacesIfNeeded(
	ctx context.Context,
	connectConfig types.CreateBGPConnectionConfig,
	vpnTunnels []client.VPNTunnel,
) ([]client.CloudRouterInterface, error) {
	cloudRouterName := c.state.GatewayName

	c.logger.Debugf("Starting to prepare Cloud Router '%s' interfaces", cloudRouterName)
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, cloudRouterName)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot validate Cloud Router Interfaces due to the error when "+
				"obtaining Cloud Router %s: %w",
			cloudRouterName, err)
	}
	if router == nil {
		return nil, fmt.Errorf(
			"unexpectedly got empty object while obtaining Cloud Router %s",
			cloudRouterName)
	}

	desiredInterfaces, err := c.getDesiredInterfacesForCloudRouter(
		connectConfig, vpnTunnels,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot configure desired Cloud Router Interfaces for Cloud Router '%s': %w",
			cloudRouterName, err,
		)
	}
	if len(desiredInterfaces) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of interfaces prepared for Cloud Router '%s': %v. "+
				"Expected 4 interfaces",
			cloudRouterName, desiredInterfaces,
		)
	}
	c.logger.Debugf(
		"The Cloud Router '%s' requires following interfaces: %v",
		cloudRouterName, desiredInterfaces,
	)

	actualInterfaces, missingInterfaces, err := c.prepareCloudRouterInterfaces(
		connectConfig, router, desiredInterfaces,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to determine missing interfaces in the Cloud Router '%s': %w",
			cloudRouterName, err)
	}
	if len(actualInterfaces) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of interfaces extracted from Cloud Router '%s': %v. "+
				"Expected 4 interfaces",
			cloudRouterName, actualInterfaces,
		)
	}

	for _, missingInterface := range missingInterfaces {
		_, err = c.gcpClient.AddRouterInterface(
			ctx, c.config.Region, cloudRouterName, missingInterface)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to create an interface '%v' for Cloud Router '%s': %w",
				missingInterface, cloudRouterName, err,
			)
		}
	}

	c.logger.Debugf("Cloud Router '%s' interfaces are covered", cloudRouterName)
	return actualInterfaces, nil
}

func (c *GCPConnector) getDesiredBGPPeersForCloudRouter(
	connectConfig types.CreateBGPConnectionConfig,
	cloudRouterInterfaces []client.CloudRouterInterface,
) ([]client.CloudRouterBGPPeer, error) {
	if len(cloudRouterInterfaces) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of provided Cloud Router Interfaces: %v. Expected 4 interfaces",
			cloudRouterInterfaces)
	}
	addresses := connectConfig.PeerBGPAddresses
	if len(addresses) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of provided Peer IP Addresses: %v. Expected 4 addresses",
			addresses)
	}

	bgpPeers := []client.CloudRouterBGPPeer{}
	for i := range cloudRouterInterfaces {
		bgpPeers = append(bgpPeers, client.CloudRouterBGPPeer{
			Name: c.GenerateName(
				fmt.Sprintf("bgppeer-%d-", (i + 1)),
			),
			Interface:     cloudRouterInterfaces[i].Name,
			PeerIPAddress: addresses[i],
			ASN:           strconv.FormatUint(connectConfig.PeerASN, 10),
		})
	}
	return bgpPeers, nil
}

func cloudRouterBGPPeersMatch(a, b client.CloudRouterBGPPeer) bool {
	return a.ASN == b.ASN && a.Interface == b.Interface && a.PeerIPAddress == b.PeerIPAddress
}

// prepareCloudRouterBGPPeers checks for the presence of
// provided desiredBGPPeers in the given Cloud Router and returns
// two slices of BGP peers.
//
// The first slice represents the peers that will be used
// for Connection creation (it may consist of the mix of already
// existing peers and those that need to be created yet).
// The second slice consists of missing peers that needs to
// be created afterwards.
func (c *GCPConnector) prepareCloudRouterBGPPeers(
	connectConfig types.CreateBGPConnectionConfig,
	cloudRouter *client.CloudRouter,
	desiredBGPPeers []client.CloudRouterBGPPeer,
) ([]client.CloudRouterBGPPeer, []client.CloudRouterBGPPeer, error) {
	c.logger.Debugf(
		"checking the presence of desired BGP Peers '%v' in Cloud Router",
		desiredBGPPeers,
	)
	if cloudRouter == nil {
		return nil, nil, errors.New(
			"cannot verify missing BGP Peers in nil Cloud Router",
		)
	}

	bgpPeers := []client.CloudRouterBGPPeer{}
	missingBGPPeers := []client.CloudRouterBGPPeer{}

	for _, desired := range desiredBGPPeers {
		c.logger.Debugf(
			"checking the presence of the BGP peer '%v' in Cloud Router '%s'",
			desired, cloudRouter.Name,
		)
		found := false
		for _, existing := range cloudRouter.BGPPeers {
			c.logger.Tracef(
				"comparing desired BGP Peer '%v' with existing one: '%v'",
				desired, existing,
			)
			if cloudRouterBGPPeersMatch(desired, existing) {
				c.logger.Debugf(
					"BGP peer '%v' found in Cloud Router '%s'. Will use it",
					existing, cloudRouter.Name,
				)
				bgpPeers = append(bgpPeers, existing)
				found = true
				break
			}
		}
		if !found {
			c.logger.Debugf(
				"BGP peer '%v' not found in Cloud Router '%s'. Will add one",
				desired, cloudRouter.Name,
			)
			bgpPeers = append(bgpPeers, desired)
			missingBGPPeers = append(missingBGPPeers, desired)
		}
	}

	c.logger.Debugf(
		"The Cloud Router '%s' will expose the following 4 BGP Peers for connection purposes: %v. "+
			"The following BGP Peers are not created yet and will be created shortly afterwards: %v",
		cloudRouter.Name, bgpPeers, missingBGPPeers,
	)
	return bgpPeers, missingBGPPeers, nil
}

// Adds required BGP Peers to Cloud Router to accomplish the
// connection between GCP and the second provider.
//
// The Cloud Router Interfaces that are provided as a parameter
// are those interfaces that will be associated with BGP Peers.
// The Cloud Router may have more interfaces than those, which
// is why we need to mark these interfaces we want to use.
func (c *GCPConnector) createCloudRouterBPGPeersIfNeeded(
	ctx context.Context,
	connectConfig types.CreateBGPConnectionConfig,
	cloudRouterInterfaces []client.CloudRouterInterface,
) ([]client.CloudRouterBGPPeer, error) {
	cloudRouterName := c.state.GatewayName

	c.logger.Debugf("Starting to prepare Cloud Router '%s' BGP Peers", cloudRouterName)
	router, err := c.gcpClient.GetRouter(ctx, c.config.Region, cloudRouterName)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot validate Cloud Router BGP Peers due to the error when "+
				"obtaining Cloud Router %s: %w",
			cloudRouterName, err)
	}
	if router == nil {
		return nil, fmt.Errorf(
			"unexpectedly got empty object while obtaining Cloud Router %s",
			cloudRouterName)
	}

	desiredBGPPeers, err := c.getDesiredBGPPeersForCloudRouter(
		connectConfig, cloudRouterInterfaces,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot configure desired Cloud Router BGP Peers for Cloud Router '%s': %w",
			cloudRouterName, err,
		)
	}
	if len(desiredBGPPeers) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of BGP Peers prepared for Cloud Router '%s': %v. "+
				"Expected 4 BGP peers",
			cloudRouterName, desiredBGPPeers,
		)
	}
	c.logger.Debugf(
		"The Cloud Router '%s' requires following BGP Peers: %v",
		cloudRouterName, router.BGPPeers,
	)

	actualBGPPeers, missingBGPPeers, err := c.prepareCloudRouterBGPPeers(
		connectConfig, router, desiredBGPPeers,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to determine missing BGP Peers in the Cloud Router '%s': %w",
			cloudRouterName, err)
	}
	if len(actualBGPPeers) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of BGP Peers extracted from Cloud Router '%s': %v. "+
				"Expected 4 BGP Peers",
			cloudRouterName, actualBGPPeers,
		)
	}

	for _, missingBGPPeer := range missingBGPPeers {
		_, err = c.gcpClient.AddRouterBGPPeer(
			ctx, c.config.Region, cloudRouterName, missingBGPPeer)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to create a BGP Peer '%v' for Cloud Router '%s': %w",
				missingBGPPeer, cloudRouterName, err,
			)
		}
	}

	c.logger.Debugf("Cloud Router '%s' BGP Peers are covered", cloudRouterName)
	return actualBGPPeers, nil
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

func (c *GCPConnector) gcpVPNGatewaysToMap(gateways []*client.VPNGateway) map[string][]string {
	result := map[string][]string{}
	for _, g := range gateways {
		if _, ok := result[g.Name]; ok {
			// Not expecting such scenario. If you found this line
			// while debugging your issue... sorry my friend.
			c.logger.Errorf(
				"Found multiple VPN Gateways with the same name. " +
					"Some Public IPs may be missing in the final result.",
			)
		}
		result[g.Name] = g.IPAddresses
	}
	return result
}

func (c *GCPConnector) GenerateName(id string) string {
	return helper.CreateNameWithRand(
		c.state.GatewayName+"/"+c.state.GatewayName, id,
	)
}

// Returns true if the name, provided as an argument, indicates that
// the resource was created while creating a connection between
// GCP Cloud Router and AWS Transit Gateway from the configuration.
func (c *GCPConnector) IsNameOwnedByConnection(name string) bool {
	return helper.NameCreatedForIdentifier(
		name,
		c.state.GatewayName+"/"+c.state.GatewayName,
	)
}
