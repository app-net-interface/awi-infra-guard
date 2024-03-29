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
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type AzureConnector struct {
	azClient    *client.Client
	config      *Config
	logger      *logrus.Entry
	transaction *transaction
}

type transaction struct {
	BGPAddresses []string
	PublicIPs    []string

	GatewayName     string
	PeerGatewayName string
}

func NewConnector(ctx context.Context, logger *logrus.Entry, config string) (provider.Provider, error) {
	parsedConfig := Config{}
	if err := yaml.Unmarshal([]byte(config), &parsedConfig); err != nil {
		return nil, fmt.Errorf("could not parse Configuration for Azure Client: %w", err)
	}

	azClient, err := client.NewClient(
		ctx,
		logger.WithField("logger", "azure_client"),
		parsedConfig.Location,
		parsedConfig.ResourceGroup)
	if err != nil {
		return nil, fmt.Errorf("could not create Azure Client: %w", err)
	}
	return &AzureConnector{
		azClient: azClient,
		config:   &parsedConfig,
		logger:   logger,
	}, nil
}

func (c *AzureConnector) Name() string {
	return "azure"
}

func (c *AzureConnector) gatewayFromVNetGateway(gw client.VNetGateway, region string) types.Gateway {
	return types.Gateway{
		Name:          gw.Name,
		CloudProvider: c.Name(),
		Kind:          "VirtualNetworkGateway",
		Region:        region,
	}
}

func (c *AzureConnector) GetGateway(ctx context.Context, gateway types.GatewayIdentifier) (*types.Gateway, error) {
	gw, err := c.azClient.GetVNetGateway(ctx, gateway.GatewayID)
	if err != nil {
		return nil, fmt.Errorf("could not get Gateway %v for Azure: %w", gateway, err)
	}
	output := c.gatewayFromVNetGateway(gw, c.config.Location)
	return &output, nil
}

func (c *AzureConnector) ListGateways(ctx context.Context) ([]types.Gateway, error) {
	gws, err := c.azClient.ListVNetGateways(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list Gateways for Azure: %w", err)
	}
	output := []types.Gateway{}
	for _, gw := range gws {
		output = append(output, c.gatewayFromVNetGateway(gw, c.config.Location))
	}
	return output, nil
}

func (c *AzureConnector) GetGatewayConnectionSettings(
	ctx context.Context, gateway types.Gateway,
) (types.GatewayConnectionSettings, error) {
	return types.GatewayConnectionSettings{
		// We keep this value hardcoded for now accordingly
		// to GCloud documents on connecting two cloud
		// providers.
		NumberOfInterfaces: 2,
		MaxNumberOfTunnels: 4,
		BGPSetting: &types.BGPSetting{
			Addressing: types.BGPAddressing{
				// It could be implemented, however.
				GeneratesBothAddresses: false,
				// If BGP was already enabled, it means most likely, that
				// the BGP IP Addresses were already associated and
				// should not be modified or else it could break existing
				// connections.
				AcceptsBothAddresses:              true,
				GeneratesOwnAndAcceptsPeerAddress: true,
			},
			AllowedIPRanges: []string{
				"169.254.21.0/24",
				"169.254.22.0/24",
			},
		},
	}, nil
}

func (c *AzureConnector) InitializeGatewayInterfaces(
	ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
) ([]string, error) {
	gw, err := c.azClient.GetVNetGateway(ctx, gateway.Name)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get VPN Gateway %v: %w",
			gw, err,
		)
	}
	// TODO: Decide what to do if there are more than two
	// Public IPs.
	//
	// Consider the situation where there are more CSP
	// connections with the same gateway and both of them
	// use different Public IPs to reach Azure Gateway.
	if len(gw.PublicIPsIDs) != 2 {
		return nil, fmt.Errorf(
			"unexpected number of Public IPs for Gateway %v: "+
				"expecting two addresses but found: %v",
			gw, gw.PublicIPsIDs,
		)
	}

	ips := make([]string, 2)
	for i, ipID := range gw.PublicIPsIDs {
		publicIP, err := c.azClient.GetPublicIP(ctx, client.IDToName(ipID))
		if err != nil {
			return nil, fmt.Errorf(
				"failed to resolve Public IP Address %s: %w",
				ipID, err,
			)
		}
		if publicIP.Address == "" {
			return nil, fmt.Errorf(
				"failed to resolve Public IP Address %s. Got empty string",
				ipID,
			)
		}
		ips[i] = publicIP.Address
	}
	c.transaction.PublicIPs = ips
	return ips, nil
}

func (c *AzureConnector) InitializeCreation(
	ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
) error {
	c.logger.Infof(
		"Starting to create a connection of Azure Gateway %s with %s:%s",
		gateway.Name, peerGateway.CloudProvider, peerGateway.Name,
	)

	c.transaction = &transaction{
		GatewayName:     gateway.Name,
		PeerGatewayName: peerGateway.Name,
	}
	return nil
}

func (c *AzureConnector) InitializeASN(
	ctx context.Context, gateway types.Gateway, peerGateway types.Gateway,
) (uint64, error) {
	gw, err := c.azClient.GetVNetGateway(ctx, gateway.Name)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to get VPN Gateway %s: %w",
			gateway.Name, err,
		)
	}

	if gw.ASN != "" {
		asn := helper.StringToUInt64Pointer(gw.ASN)
		if asn == nil {
			return 0, fmt.Errorf(
				"the ASN number '%s' cannot be parsed from gateway: %v",
				gw.ASN, gateway,
			)
		}
		return *asn, nil
	}

	return 0, nil
}

func (c *AzureConnector) AttachToExternalGatewayWithBGP(
	ctx context.Context,
	gateway types.Gateway,
	peerGateway types.Gateway,
	attachMode types.AttachBGPConnectionMode,
	config types.CreateBGPConnectionConfig,
) (types.OutputForConnectionWithBGP, error) {
	c.logger.Infof(
		"Running attachment process with BGP Session. Mode: %d", attachMode,
	)

	if attachMode == types.AttachModeGenerateBothIPs {
		return types.OutputForConnectionWithBGP{}, errors.New("NOT IMPLEMENTED")
	}
	if attachMode == types.AttachModeGenerateIPAndAcceptOtherIP {
		return types.OutputForConnectionWithBGP{}, errors.New("NOT IMPLEMENTED")
	}

	if attachMode == types.AttachModeGenerateIP {

		err := c.AddBGPAddressesToVnetGateway(
			ctx,
			config.BGPCIDRPools,
			gateway.Name,
			config.NumberOfTunnels,
		)
		if err != nil {
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to initialize BGP Addresses in Virtual Network Gateway %s: %w",
				gateway.Name, err,
			)
		}

		return types.OutputForConnectionWithBGP{
			BGPAddresses:  c.transaction.BGPAddresses,
			SharedSecrets: config.SharedSecrets,
			Interfaces:    c.transaction.PublicIPs,
		}, nil
	}

	if len(config.BGPAddresses) != 0 {
		c.transaction.BGPAddresses = config.BGPAddresses

		err := c.AddBGPAddressesToVnetGateway(
			ctx,
			config.BGPCIDRPools,
			gateway.Name,
			config.NumberOfTunnels,
		)
		if err != nil {
			return types.OutputForConnectionWithBGP{}, fmt.Errorf(
				"failed to initialize BGP Addresses in Virtual Network Gateway %s: %w",
				gateway.Name, err,
			)
		}
	}

	localGWs, err := c.createLocalVNetGateways(
		ctx,
		config,
		gateway.Name,
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to create Local Network Gateways: %w",
			err,
		)
	}

	err = c.createNetworkGatewayConnection(
		ctx,
		config,
		gateway.Name,
		localGWs,
	)
	if err != nil {
		return types.OutputForConnectionWithBGP{}, fmt.Errorf(
			"failed to create VPN Connections: %w",
			err,
		)
	}

	return types.OutputForConnectionWithBGP{}, nil
}

// NOT IMPLEMENTED
//
// Just a placeholder for future CSP improvements.
func (c *AzureConnector) AttachToExternalGatewayWithStaticRouting() error {
	return errors.New("NOT IMPLEMENTED")
}

func (c *AzureConnector) DeleteConnectionResources(
	ctx context.Context,
	gateway types.Gateway,
	peerGateway types.Gateway,
) error {
	c.transaction = &transaction{
		GatewayName:     gateway.Name,
		PeerGatewayName: peerGateway.Name,
	}

	if err := c.deleteNetworkGatewayConnections(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up Network Gateway Connections for gateway '%s': %w",
			gateway.Name, err,
		)
	}

	if err := c.deleteLocalVNetGateways(ctx); err != nil {
		return fmt.Errorf(
			"failed to clean up Local Network Gateways for gateway '%s': %w",
			gateway.Name, err,
		)
	}

	return nil
}

func (c *AzureConnector) GetVPCForGateway(ctx context.Context, gateway types.Gateway) (string, error) {
	vnetGW, err := c.azClient.GetVNetGateway(ctx, gateway.Name)
	if err != nil {
		return "", fmt.Errorf(
			"failed to get Virtual Network Gateway %s: %w", gateway.Name, err,
		)
	}
	return vnetGW.VNet, nil
}

func (c *AzureConnector) Close() error {
	if err := c.azClient.Close(); err != nil {
		return fmt.Errorf(
			"failed to close Azure Connector: %w", err,
		)
	}
	return nil
}

func (c *AzureConnector) GetCIDRs(ctx context.Context, gateway types.Gateway) ([]string, error) {
	vnetName, err := c.GetVPCForGateway(ctx, gateway)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get Virtual Network Name for Virtual Network Gateway %s: %w",
			gateway.Name, err,
		)
	}
	vnet, err := c.azClient.GetVNet(ctx, vnetName)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get Virtual Network %s: %w", vnetName, err,
		)
	}
	return vnet.AddressPrefixes, nil
}
