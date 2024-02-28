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
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/connector/aws/client"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type AWSConnector struct {
	awsClient        *client.Client
	config           *Config
	logger           *logrus.Entry
	transactionState *transactionState
}

type transactionState struct {
	GatewayName     string
	PeerGatewayName string
}

func NewConnector(ctx context.Context, logger *logrus.Entry, config string) (provider.Provider, error) {
	parsedConfig := Config{}
	if err := yaml.Unmarshal([]byte(config), &parsedConfig); err != nil {
		return nil, fmt.Errorf("could not parse Configuration for AWS Client: %w", err)
	}

	awsClient, err := client.NewClient(ctx, logger.WithField("logger", "aws_client"), parsedConfig.Region)
	if err != nil {
		return nil, fmt.Errorf("Could not create AWS Client: %w", err)
	}
	return &AWSConnector{
		awsClient:        awsClient,
		config:           &parsedConfig,
		logger:           logger,
		transactionState: &transactionState{},
	}, nil
}

func (c *AWSConnector) Name() string {
	return "aws"
}

func (c *AWSConnector) Close() error {
	return nil
}

func (c *AWSConnector) GetGatewayConnectionSettings(
	ctx context.Context, gateway types.Gateway,
) (types.GatewayConnectionSettings, error) {
	// TODO: At this stage, the provider needs to
	// go through other Gateways and see IPs that
	// were already associated.
	return types.GatewayConnectionSettings{
		NumberOfInterfaces: 4,
		BGPSetting: &types.BGPSetting{
			// Currently, AWS SDK doesn't support providing
			// custom IP Addresses for BGP for either Own
			// IP Address or Peer.
			Addressing: types.BGPAddressing{
				GeneratesBothAddresses:            true,
				AcceptsBothAddresses:              false,
				GeneratesOwnAndAcceptsPeerAddress: false,
			},
			AllowedIPRanges: []string{"192.254.0.0/16"},
			// TODO: Specify Excluded IP Ranges by checking
			// already used IP Ranges and including provider
			// reserved IP Ranges.
			ExcludedIPRanges: []string{},
		},
	}, nil
}

func (c *AWSConnector) ListGateways(ctx context.Context) ([]types.Gateway, error) {
	transitGateways, err := c.awsClient.ListTransitGateways(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list Transit Gateways from AWS: %w", err)
	}
	gateways := []types.Gateway{}
	for i := range transitGateways {
		if transitGateways[i] == nil {
			return nil, fmt.Errorf(
				"got empty Transit Gateway object in the list of transit gateways: %v",
				transitGateways)
		}
		// TODO: Handle multi-region configuration when it is implemented.
		//
		// Currently, the CSP Connector works in a single region context.
		gateways = append(gateways, c.transitGatewayToGateway(*transitGateways[i], c.config.Region))
	}
	return gateways, nil
}

func (c *AWSConnector) InitializeASN(
	ctx context.Context, gateway, peerGateway types.Gateway,
) (uint64, error) {
	asn, err := strconv.Atoi(gateway.ASN)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ASN number: %w", err)
	}
	// TODO: Fix type casting (int -> uint64 - possible loss)
	return uint64(asn), nil
}

func (c *AWSConnector) InitializeGatewayInterfaces(
	ctx context.Context, gateway, peerGateway types.Gateway,
) ([]string, error) {
	c.transactionState.GatewayName = gateway.Name
	c.transactionState.PeerGatewayName = peerGateway.Name
	// TODO: Check if AWS Interfaces can be initialized in
	// this step. Can we create resources which will give
	// us interface IP Addresses without knowing interfaces
	// of the other side yet and their BGP Configuration.
	return []string{}, nil
}

func (c *AWSConnector) GetGateway(ctx context.Context, gateway types.GatewayIdentifier) (*types.Gateway, error) {
	transitGateway, err := c.awsClient.GetTransitGateway(ctx, gateway.GatewayID)
	if err != nil {
		return nil, fmt.Errorf(
			"The AWS Transit Gateway '%s' was not found: %w", gateway.GatewayID, err)
	}
	if transitGateway == nil {
		return nil, fmt.Errorf(
			"The AWS Transit Gateway '%s' object is empty", gateway.GatewayID)
	}
	gw := c.transitGatewayToGateway(*transitGateway, gateway.Region)
	return &gw, nil
}

func (c *AWSConnector) GetVPCForGateway(ctx context.Context, gateway types.Gateway) (string, error) {
	tgw, err := c.awsClient.GetTransitGateway(ctx, gateway.Name)
	if err != nil {
		return "", fmt.Errorf(
			"failed to get Transit Gateway '%s': %w",
			gateway.Name, err,
		)
	}
	if tgw == nil {
		return "", fmt.Errorf(
			"failed to get Transit Gateway '%s': doesn't exist",
			gateway.Name,
		)
	}
	return tgw.VPCID, nil
}

// NOT IMPLEMENTED
func (c *AWSConnector) AttachToExternalGatewayWithStaticRouting() error {
	return errors.New("AttachToExternalGatewayWithStaticRouting not implemented in aws")
}

func xmlVpnConnectionToGeneratedBGPAddresses(connections []client.XMLVpnConnection) types.OutputForConnectionWithBGP {
	bgpAddresses := types.OutputForConnectionWithBGP{
		BGPAddresses:     []string{},
		PeerBGPAddresses: []string{},
		SharedSecrets:    []string{},
		Interfaces:       []string{},
	}

	for _, conn := range connections {
		for _, tunnel := range conn.IpsecTunnels {
			bgpAddresses.Interfaces = append(
				bgpAddresses.Interfaces, tunnel.VpnGateway.TunnelOutsideAddress.IPAddress)
			bgpAddresses.BGPAddresses = append(
				bgpAddresses.BGPAddresses, tunnel.VpnGateway.TunnelInsideAddress.IPAddress)
			bgpAddresses.PeerBGPAddresses = append(
				bgpAddresses.PeerBGPAddresses, tunnel.CustomerGateway.TunnelInsideAddress.IPAddress)
			bgpAddresses.SharedSecrets = append(
				bgpAddresses.SharedSecrets, tunnel.Ike.PreSharedKey)
		}
	}

	return bgpAddresses
}

func (c *AWSConnector) ensureTransitGatewayPropagatesRoutes(
	ctx context.Context, transitGatewayName string,
) error {
	transitGateway, err := c.awsClient.GetTransitGateway(ctx, transitGatewayName)
	if err != nil {
		return fmt.Errorf(
			"The AWS Transit Gateway '%s' was not found: %w",
			transitGatewayName, err)
	}
	if transitGateway == nil {
		return fmt.Errorf(
			"The AWS Transit Gateway '%s' returned nil object", transitGatewayName)
	}
	if transitGateway.AutoAcceptSharedAttachments &&
		transitGateway.DefaultRouteTableAssociation &&
		transitGateway.DefaultRouteTablePropagation &&
		transitGateway.DnsSupport &&
		transitGateway.VpnEcmpSupport {
		c.logger.Debugf(
			"The Transit Gateway '%s' propagating options are as expected. "+
				"Nothing to do", transitGatewayName,
		)
		return nil
	}

	c.logger.Warnf(
		"The Transit Gateway '%s' options need to be readjusted. Got: "+
			"AutoAcceptSharedAttachments: '%v', "+
			"DefaultRouteTableAssociation: '%v', "+
			"DefaultRouteTablePropagation: '%v', "+
			"DnsSupport: '%v', "+
			"VpnEcmpSupport: '%v', "+
			"but AWS Provider expects them to be all set to true for "+
			"establishing the connection with BGP Session. Updating Transit Gateway",
		transitGatewayName, transitGateway.AutoAcceptSharedAttachments,
		transitGateway.DefaultRouteTableAssociation,
		transitGateway.DefaultRouteTablePropagation,
		transitGateway.DnsSupport, transitGateway.VpnEcmpSupport,
	)
	transitGateway.AutoAcceptSharedAttachments = true
	transitGateway.DefaultRouteTableAssociation = true
	transitGateway.DefaultRouteTablePropagation = true
	transitGateway.DnsSupport = true
	transitGateway.VpnEcmpSupport = true

	if err := c.awsClient.UpdateTransitGateway(ctx, *transitGateway); err != nil {
		return fmt.Errorf(
			"the update of Transit Gateway '%s' propagation options failed: %w",
			transitGatewayName, err)
	}
	return nil
}

func (c *AWSConnector) AttachToExternalGatewayWithBGP(
	ctx context.Context,
	gateway, peerGateway types.Gateway,
	attachMode types.AttachBGPConnectionMode,
	config types.CreateBGPConnectionConfig,
) (types.OutputForConnectionWithBGP, error) {
	if attachMode != types.AttachModeGenerateBothIPs {
		return types.OutputForConnectionWithBGP{}, errors.New(
			"AWS SDK doesn't support accepting custom IP Addresses for BGP. " +
				"The only supported Attach Mode is AttachModeGenerateBothIPs. " +
				"This error indicates that CSP Connector, due to unknown reason, " +
				"has specified a scenario where AWS Provider serves different role" +
				"than Authoritarian Address provider - it should NOT happen.",
		)
	}
	if err := c.ensureTransitGatewayPropagatesRoutes(ctx, gateway.Name); err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to perform attachment to External Gateway since Transit Gateway "+
				"'%s' could not be prepared for propagating routes: %v",
			gateway.Name, err,
		)
	}

	transitGateway, err := c.awsClient.GetTransitGateway(ctx, gateway.Name)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"The AWS Transit Gateway '%s' was not found: %w",
			gateway.Name, err)
	}

	customerGateways, err := c.createCustomerGatewaysIfNeeded(
		ctx, config.OutsideInterfaces, strconv.FormatUint(config.PeerASN, 10),
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"got problem while trying to adjust AWS Customer Gateways: %w", err,
		)
	}

	vpnConnections, err := c.createVPNConnectionsIfNeeded(
		ctx, transitGateway.ID, transitGateway.VPCID, customerGateways,
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"got problem while trying to adjust AWS VPN Connections: %w", err,
		)
	}

	configurations := []client.XMLVpnConnection{}
	for i := range vpnConnections {
		if vpnConnections[i] == nil {
			return types.OutputForConnectionWithBGP{}, errors.New("got nil VPN Connection")
		}
		config, err := client.VPNConnectionConfigToObject(vpnConnections[i].Configuration)
		if err != nil || config == nil {
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to parse VPN Connection configuration %s or got it empty. Err: %w",
				vpnConnections[i].Configuration, err,
			)
		}
		configurations = append(configurations, *config)
	}

	return xmlVpnConnectionToGeneratedBGPAddresses(configurations), nil
}

func (c *AWSConnector) DeleteConnectionResources(
	ctx context.Context,
	gateway types.Gateway,
	peerGateway types.Gateway,
) error {
	c.transactionState.GatewayName = gateway.Name
	c.transactionState.PeerGatewayName = peerGateway.Name
	if err := c.deleteVPNConnectionsForConnection(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up VPN Connection: %w", err,
		)
	}
	if err := c.deleteCustomerGatewaysForConnection(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up Customer Gateways: %w", err,
		)
	}
	return nil
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

func (c *AWSConnector) deleteCustomerGatewaysForConnection(ctx context.Context) error {
	c.logger.Debug("Looking for Customer Gateways that should be deleted")
	gateways, err := c.awsClient.ListCustomerGateways(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot get existing Customer Gateways due to: %w", err,
		)
	}

	for _, gw := range gateways {
		if !c.shouldResourceBeDeleted(gw.Tags) {
			continue
		}
		c.logger.Debugf("Found Customer Gateway to delete '%s'", gw.ID)
		if err = c.awsClient.DeleteCustomerGateway(ctx, gw.ID); err != nil {
			// TODO: Consider situation where Customer Gateway, created by the
			// CSP script was updated with VPN Connections created by someone
			// else. How should we treat such situations?
			return fmt.Errorf("failed to clean up Customer Gateway: %w", err)
		}
		c.logger.Debugf("Successfully deleted Customer Gateway '%s'", gw.ID)
	}

	c.logger.Debug("Customer Gateways are cleared")
	return nil
}

func (c *AWSConnector) transitGatewayToGateway(tgw client.TransitGateway, region string) types.Gateway {
	return types.Gateway{
		Name:          tgw.ID,
		Kind:          "TransitGateway",
		CloudProvider: c.Name(),
		Region:        region,
		VPC:           tgw.VPCID,
		ASN:           tgw.ASN,
	}
}

func (c *AWSConnector) filterOutNotMatchingCustomerGateways(gateways []client.CustomerGateway, ip string) []client.CustomerGateway {
	filtered := []client.CustomerGateway{}
	for _, gw := range gateways {
		if gw.IP != ip {
			continue
		}
		state := strings.ToUpper(gw.State)
		if state == "DELETING" || state == "DELETED" {
			c.logger.Debugf(
				"Found Customer Gateway '%s' with IP '%s' but in '%s' state. Ignoring it.",
				gw.ID, gw.IP, gw.State,
			)
			continue
		}
		filtered = append(filtered, gw)
	}
	return filtered
}

func (c *AWSConnector) pickCustomerGatewayWithIPAndASN(gateways []client.CustomerGateway, ip, asn string) (*client.CustomerGateway, error) {
	filtered := c.filterOutNotMatchingCustomerGateways(gateways, ip)
	if len(filtered) == 0 {
		return nil, nil
	}
	if len(filtered) > 1 {
		c.logger.Warnf(
			"found more than 1 customer gateways with IP '%s': %v. Checking if any meets additional criteria",
			ip, filtered,
		)
	}
	var pendingGw *client.CustomerGateway
	var availableGw *client.CustomerGateway

	for _, gw := range filtered {
		state := strings.ToUpper(gw.State)
		if gw.ASN != asn {
			c.logger.Warnf(
				"found customer gateway '%s' with proper IP '%s' but different ASN than requested: '%s'. "+
					"Checking other customer gateways",
				gw.ID, gw.IP, gw.ASN,
			)
			continue
		}
		if state == "PENDING" {
			pendingGw = &client.CustomerGateway{}
			*pendingGw = gw
			continue
		}
		if state == "AVAILABLE" {
			availableGw = &client.CustomerGateway{}
			*availableGw = gw
			continue
		}
		return nil, fmt.Errorf(
			"found customer gateway '%s' with invalid state: %s",
			gw.ID, gw.State,
		)
	}

	if pendingGw != nil && availableGw != nil {
		c.logger.Warnf(
			"found two customer gateways with IP '%s' and ASN '%s'. "+
				"One in pending state: '%s' and second one currently available: '%s'. "+
				"THE ONE AVAILABLE WILL BE PICKED. If something changes and the available "+
				"Customer Gateway is removed and the pending one becomes an available one, the"+
				"script may need to be run again to update customer gateways, "+
				"to be run again so that the correct Customer Gateway is picked.",
			ip, asn, pendingGw.ID, availableGw.ID,
		)
		return availableGw, nil
	}

	if availableGw != nil {
		return availableGw, nil
	}

	if pendingGw != nil {
		c.logger.Warnf(
			"found AWS customer gateway with proper IP Address '%s' and ASN '%s'"+
				"but the Gateway is yet in a pending state. Make sure the Customer "+
				"Gateway will become available as soon as possible, otherwise Cloud "+
				"Provider Connection may not work properly.",
			ip, asn,
		)
	}

	return nil, fmt.Errorf(
		"found Customer Gateway(s) with requested IP Address '%s' but with no proper ASN. "+
			"Expected ASN: '%s'. Got Gateways: %v", ip, asn, filtered,
	)
}

// createCustomerGatewaysIfNeeded verifies if proper AWS Customer Gateways
// exist and either creates missing one or reports an error if Customer
// Gateways represent the state that is hard to overcome without additional
// input from the user (like the existing Customer Gateway with requested
// IP Address but different ASN number).
//
// The method returns list of Customer Gateway IDs representing:
// * ID of Customer Gateway representing firstGCPInterface
// * ID of Customer Gateway representing secondGCPInterface
func (c *AWSConnector) createCustomerGatewaysIfNeeded(
	ctx context.Context, interfaces []string, ASN string,
) ([]string, error) {
	c.logger.Debug("Checking existing AWS Customer Gateways")
	gateways, err := c.awsClient.ListCustomerGateways(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot verify current state of AWS Customer Gateways due to: %w", err,
		)
	}

	customerGateways := []string{}

	for _, iface := range interfaces {
		gw, err := c.pickCustomerGatewayWithIPAndASN(gateways, iface, ASN)
		if err != nil {
			return nil, fmt.Errorf(
				"found a problem while inspecting the presence for Customer Gateway for IP %s: %w",
				iface, err,
			)
		}

		if gw == nil {
			c.logger.Debugf("Customer Gateway for IP '%s' does not exist. Creating it.", iface)
			cgw, err := c.awsClient.CreateCustomerGateway(
				ctx,
				iface,
				ASN,
				c.generateTagName())
			if err != nil {
				return nil, fmt.Errorf("Could not create AWS Customer Gateway for IP '%s': %w", iface, err)
			}
			if cgw == nil {
				return nil, errors.New("The Created Customer Gateway object is empty. Try restarting the script.")
			}
			customerGateways = append(customerGateways, cgw.ID)
		} else {
			c.logger.Debugf("Customer Gateway for IP '%s' already exists under ID '%s'. Using it.", iface, gw.ID)
			customerGateways = append(customerGateways, gw.ID)
		}
	}

	c.logger.Debug("AWS Customer Gateways covered")
	return customerGateways, nil
}

func (c *AWSConnector) pickVPNConnectionWithCustomerGatewayTransitGateway(
	vpnConnections []*client.VPNConnection, customerGatewayID, transitGatewayID string,
) (*client.VPNConnection, error) {
	matching := []*client.VPNConnection{}

	for _, conn := range vpnConnections {
		if customerGatewayID == conn.CustomerGatewayID && transitGatewayID == conn.TransitGatewayID {
			connection := *conn
			matching = append(matching, &connection)
		}
	}

	if len(matching) == 0 {
		return nil, nil
	}

	if len(matching) > 1 {
		c.logger.Warnf(
			"Found more VPN Connections than 1 for Customer Gateway ID '%s' and "+
				"Transit Gateway ID '%s'", customerGatewayID, transitGatewayID)
	}

	return matching[0], nil
}

func getCIDRPool(cidrsInUse []string) (*helper.CIDRV4Pool, error) {
	pool, err := helper.NewCIDRV4Pool("169.254.0.0/16")
	if err != nil {
		return nil, fmt.Errorf("cannot create a CIDR Pool: %w", err)
	}
	if pool == nil {
		return nil, fmt.Errorf("got empty CIDR Pool object")
	}
	reservedBlocks := []string{
		"169.254.0.0/30",
		"169.254.1.0/30",
		"169.254.2.0/30",
		"169.254.3.0/30",
		"169.254.4.0/30",
		"169.254.5.0/30",
		"169.254.169.252/30",
	}
	forbiddenCIDRs := append(reservedBlocks, cidrsInUse...)
	for _, forbiddenCIDR := range forbiddenCIDRs {
		_, cidr, err := net.ParseCIDR(forbiddenCIDR)
		if err != nil {
			return nil, fmt.Errorf(
				"got problem while parsing one of AWS Reserved Blocks '%s' for Tunneling: %w",
				forbiddenCIDR, err,
			)
		}
		if err = pool.ExcludeCIDRFromPool(cidr); err != nil {
			return nil, fmt.Errorf(
				"got problem while preventing CIDR pool from using"+
					"AWS Reserved Blocks '%s' for Tunneling: %w",
				forbiddenCIDR, err,
			)
		}
	}
	return pool, nil
}

func (c *AWSConnector) getIDsOfTransitGatewaysBelongingToVPC(
	transitGateways []*client.TransitGateway, vpc string,
) map[string]struct{} {
	selected := map[string]struct{}{}
	for i := range transitGateways {
		if transitGateways[i] == nil {
			c.logger.Infof(
				"Registered empty Transit Gateway object within the list of Transit Gateways: %v",
				transitGateways,
			)
			continue
		}
		if transitGateways[i].VPCID == vpc {
			selected[transitGateways[i].ID] = struct{}{}
		}
	}
	return selected
}

func (c *AWSConnector) getTunnelCIDRsInUseForVPCTransitGateway(
	ctx context.Context, vpc string, vpnConnections []*client.VPNConnection,
) ([]string, error) {
	transitGateways, err := c.awsClient.ListTransitGateways(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot obtain list of Transit Gateways: %w", err)
	}
	tgwIDs := c.getIDsOfTransitGatewaysBelongingToVPC(transitGateways, vpc)

	cidrs := []string{}
	for i := range vpnConnections {
		if vpnConnections[i] == nil {
			c.logger.Infof(
				"Registered empty VPN Connection object within the list of VPN Connections: %v",
				vpnConnections,
			)
			continue
		}
		if _, ok := tgwIDs[vpnConnections[i].TransitGatewayID]; !ok {
			continue
		}
		for _, option := range vpnConnections[i].TunnelOptions {
			cidrs = append(cidrs, option.CIDR)
		}
	}
	return cidrs, nil
}

func (c *AWSConnector) createTunnelOption(
	pool *helper.CIDRV4Pool,
) (*client.TunnelOption, error) {
	cidr, err := pool.Get(REQUESTED_CIDR_SIZE)
	if err != nil {
		return nil, fmt.Errorf("error while trying to obtain a new CIDR: %w", err)
	}
	if cidr == nil {
		return nil, fmt.Errorf(
			"could not obtain CIDR with size %d from the pool %v. "+
				"Cannot fit requested CIDR size in the pool.", REQUESTED_CIDR_SIZE, pool)
	}

	return &client.TunnelOption{
		CIDR: cidr.String(),
	}, nil
}

// Prepares two pairs of new Tunnel Options.
//
// This method scans through all Transit Gateways which belong to the
// same VPC as the requested Transit Gateway and then find all VPN
// Connections associated with these Transit Gateways in order to
// collect all CIDRs that are being in use by these VPN Connections
// and find out which CIDR can be used to keep it unique within that
// particular VPC.
//
// It also generates PreSharedKeys.
func (c *AWSConnector) prepareNewTunnelOptions(
	ctx context.Context, vpc string, vpnConnections []*client.VPNConnection,
) ([][]client.TunnelOption, error) {
	takenTunnelCIDRs, err := c.getTunnelCIDRsInUseForVPCTransitGateway(ctx, vpc, vpnConnections)
	if err != nil {
		return nil, fmt.Errorf(
			"could not verify CIDRs from already existing VPN Connections: %w", err)
	}
	pool, err := getCIDRPool(takenTunnelCIDRs)
	if err != nil {
		return nil, fmt.Errorf(
			"could not create CIDR Pool for preparing Tunnel Options: %w", err)
	}

	tunnelOptions := [][]client.TunnelOption{}

	for i := 0; i < 2; i++ {
		options := []client.TunnelOption{}
		for j := 0; j < 2; j++ {
			tunnelOption, err := c.createTunnelOption(pool)
			if err != nil {
				return nil, fmt.Errorf(
					"failed to prepare tunnel options: %w", err,
				)
			}
			options = append(options, *tunnelOption)
		}
		tunnelOptions = append(tunnelOptions, options)
	}

	return tunnelOptions, nil
}

func (c *AWSConnector) createVPNConnectionsIfNeeded(
	ctx context.Context, transitGatewayID, vpcID string, customerGatewayIDs []string,
) ([]*client.VPNConnection, error) {
	c.logger.Debug("Checking AWS VPN Connections")
	if len(customerGatewayIDs) != 2 {
		return nil, fmt.Errorf(
			"Unexpected number of Customer Gateway IDs. Expected 2. Got the following: %v",
			customerGatewayIDs,
		)
	}
	connections, err := c.awsClient.ListVPNConnections(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot verify current state of AWS VPN Connections due to: %w", err,
		)
	}
	firstConnection, err := c.pickVPNConnectionWithCustomerGatewayTransitGateway(
		connections, customerGatewayIDs[0], transitGatewayID)
	if err != nil {
		return nil, fmt.Errorf(
			"found a problem while inspecting the presence for VPN Connection for "+
				" Customer Gateway '%s' and Transit Gateway '%s': %w",
			customerGatewayIDs[0], transitGatewayID, err,
		)
	}
	secondConnection, err := c.pickVPNConnectionWithCustomerGatewayTransitGateway(
		connections, customerGatewayIDs[1], transitGatewayID)
	if err != nil {
		return nil, fmt.Errorf(
			"found a problem while inspecting the presence for VPN Connection for "+
				" Customer Gateway '%s' and Transit Gateway '%s': %w",
			customerGatewayIDs[1], transitGatewayID, err,
		)
	}

	vpnConnections := []*client.VPNConnection{}

	// TODO: This step can be optimized by prior check for the existence
	// of connections. We assume that existing VPN Connections may have
	// already tunnel options and if not it means that the user should
	// solve it on his own.
	tunnelOptions, err := c.prepareNewTunnelOptions(ctx, vpcID, vpnConnections)
	if err != nil {
		return nil, fmt.Errorf(
			"found a problem while trying to obtain new possible tunnel options: %w",
			err,
		)
	}
	if len(tunnelOptions) != 2 {
		return nil, fmt.Errorf(
			"internal error. Expected exactly 2 tunnel options but got %d",
			len(tunnelOptions),
		)
	}

	for _, conn := range []struct {
		customerGatewayID string
		conn              *client.VPNConnection
		tunnelOptions     []client.TunnelOption
	}{
		{customerGatewayIDs[0], firstConnection, tunnelOptions[0]},
		{customerGatewayIDs[1], secondConnection, tunnelOptions[1]},
	} {
		if conn.conn == nil {
			c.logger.Debugf(
				"VPN Connection for Customer Gateway '%s' and Transit Gateway '%s' "+
					"does not exist. Creating it.", conn.customerGatewayID, transitGatewayID)
			vpnConn, err := c.awsClient.CreateVPNConnection(
				ctx,
				conn.customerGatewayID,
				transitGatewayID,
				conn.tunnelOptions,
				c.generateTagName())
			if err != nil {
				return nil, fmt.Errorf(
					"Could not create AWS VPN Connection for Customer Gateway '%s' "+
						"and Transit Gateway '%s': %w", conn.customerGatewayID, transitGatewayID, err)
			}
			if vpnConn == nil {
				return nil, errors.New("The Created VPN Connection object is empty. Try restarting the script.")
			}
			vpnConnections = append(vpnConnections, vpnConn)
		} else {
			c.logger.Debugf(
				"VPN Connection for Customer Gateway '%s' and Transit Gateway '%s' "+
					"already exists under ID '%s'. Using it.", conn.customerGatewayID, transitGatewayID, conn.conn.ID)
			vpnConnections = append(vpnConnections, conn.conn)
		}
	}

	c.logger.Debug("VPN Connections Covered")
	return vpnConnections, nil
}

func (c *AWSConnector) generateTagName() string {
	return helper.CreateName(
		// TODO: Enforce one order of connection sides.
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName, "",
	)
}

func (c *AWSConnector) shouldResourceBeDeleted(tags []client.Tag) bool {
	shouldDelete := false
	for _, tag := range tags {
		if c.isTagOwnedByAnotherConnection(tag.Key) {
			return false
		}
		if !shouldDelete && c.isTagOwnedByConnection(tag.Key) {
			shouldDelete = true
		}
	}
	return shouldDelete
}

func (c *AWSConnector) isTagOwnedByConnection(tagName string) bool {
	return helper.NameCreatedForIdentifier(
		tagName,
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName,
	)
}

func (c *AWSConnector) isTagOwnedByAnotherConnection(tagName string) bool {
	return helper.NameCreatedByScript(tagName) && !helper.NameCreatedForIdentifier(
		tagName,
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName,
	)
}
