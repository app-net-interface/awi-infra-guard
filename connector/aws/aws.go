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

package aws

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
		return nil, fmt.Errorf("could not create AWS Client: %w", err)
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
			AllowedIPRanges: []string{"169.254.0.0/16"},
			// TODO: Specify Excluded IP Ranges by checking
			// already used IP Ranges and including provider
			// reserved IP Ranges.
			ExcludedIPRanges: []string{
				"169.254.0.0/30",
				"169.254.1.0/30",
				"169.254.2.0/30",
				"169.254.3.0/30",
				"169.254.4.0/30",
				"169.254.5.0/30",
				"169.254.169.252/30",
			},
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

func (c *AWSConnector) InitializeCreation(
	ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
) error {
	c.transactionState.GatewayName = gateway.Name
	c.transactionState.PeerGatewayName = peerGateway.Name
	return nil
}

func (c *AWSConnector) InitializeASN(
	ctx context.Context, gateway, peerGateway types.Gateway,
) (uint64, error) {
	asn := helper.StringToUInt64Pointer(gateway.ASN)
	if asn == nil {
		return 0, fmt.Errorf("failed to parse ASN number: %s", gateway.ASN)
	}
	return *asn, nil
}

func (c *AWSConnector) InitializeGatewayInterfaces(
	ctx context.Context, gateway, peerGateway types.Gateway,
) ([]string, error) {
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
			"the AWS Transit Gateway '%s' was not found: %w", gateway.GatewayID, err)
	}
	if transitGateway == nil {
		return nil, fmt.Errorf(
			"the AWS Transit Gateway '%s' object is empty", gateway.GatewayID)
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

func xmlVpnConnectionToGeneratedBGPAddresses(
	connections []client.XMLVpnConnection,
	sharedSecrets []string,
) types.OutputForConnectionWithBGP {
	bgpAddresses := types.OutputForConnectionWithBGP{
		BGPAddresses:     []string{},
		PeerBGPAddresses: []string{},
		SharedSecrets:    sharedSecrets,
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
			"the AWS Transit Gateway '%s' was not found: %w",
			transitGatewayName, err)
	}
	if transitGateway == nil {
		return fmt.Errorf(
			"the AWS Transit Gateway '%s' returned nil object", transitGatewayName)
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
				"than Authoritarian Address provider - it should NOT happen",
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
			"the AWS Transit Gateway '%s' was not found: %w",
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
		ctx,
		transitGateway.ID,
		customerGateways,
		config.BGPCIDRPools,
		config.SharedSecrets,
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

	output := xmlVpnConnectionToGeneratedBGPAddresses(configurations, config.SharedSecrets)

	return output, nil
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

func (c *AWSConnector) GetCIDRs(ctx context.Context, gateway types.Gateway) ([]string, error) {
	tgw, err := c.awsClient.GetTransitGateway(ctx, gateway.Name)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get Transit Gateway '%s': %w",
			gateway.Name, err,
		)
	}
	if tgw.VPCID == "" {
		c.logger.Debugf(
			"Transit Gateway %s has no VPC associated: no CIDRs found",
			gateway.Name,
		)
		return nil, nil
	}
	vpc, err := c.awsClient.GetVPC(ctx, tgw.VPCID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get VPC '%s' which is associated with Transit Gateway '%s': %w",
			tgw.VPCID, gateway.Name, err,
		)
	}
	if vpc == nil {
		return nil, fmt.Errorf(
			"failed to get VPC '%s' which is associated with Transit Gateway '%s': "+
				"the VPC object is nil",
			tgw.VPCID, gateway.Name,
		)
	}
	return []string{vpc.CIDR}, nil
}
