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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/sirupsen/logrus"
)

type Client struct {
	logger                         *logrus.Entry
	subscriptionID                 string
	location                       string
	resourceGroup                  string
	vnetClient                     *armnetwork.VirtualNetworksClient
	vnetGatewayClient              *armnetwork.VirtualNetworkGatewaysClient
	localVnetGatewayClient         *armnetwork.LocalNetworkGatewaysClient
	networkGatewayConnectionClient *armnetwork.VirtualNetworkGatewayConnectionsClient
	publicIPClient                 *armnetwork.PublicIPAddressesClient
}

func NewClient(
	ctx context.Context, logger *logrus.Entry, location, resourceGroup string,
) (*Client, error) {
	subscriptionID := os.Getenv("CSP_AZ_SUB_ID")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain credentials for Azure Client: %v", err,
		)
	}
	vnetClient, err := armnetwork.NewVirtualNetworksClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VNet Client: %v", err,
		)
	}
	vnetGatewayClient, err := armnetwork.NewVirtualNetworkGatewaysClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VNet Gateway Client: %v", err,
		)
	}

	localVnetGatewayClient, err := armnetwork.NewLocalNetworkGatewaysClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Local VNet Gateway Client: %v", err,
		)
	}

	networkGatewayConnectionClient, err := armnetwork.NewVirtualNetworkGatewayConnectionsClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VPN Connection Client: %v", err,
		)
	}

	publicIPClient, err := armnetwork.NewPublicIPAddressesClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Public IP Address Client: %v", err,
		)
	}

	return &Client{
		logger:                         logger,
		subscriptionID:                 subscriptionID,
		resourceGroup:                  resourceGroup,
		location:                       location,
		vnetClient:                     vnetClient,
		vnetGatewayClient:              vnetGatewayClient,
		localVnetGatewayClient:         localVnetGatewayClient,
		networkGatewayConnectionClient: networkGatewayConnectionClient,
		publicIPClient:                 publicIPClient,
	}, nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) GetVNet(ctx context.Context, id string) (VNet, error) {
	resp, err := c.vnetClient.Get(
		ctx, c.resourceGroup, id, nil,
	)
	if err != nil {
		return VNet{}, fmt.Errorf(
			"failed to get a VNet '%s': %v", id, err,
		)
	}
	return vNetFromAzure(resp.VirtualNetwork), nil
}

func (c *Client) GetVNetGateway(ctx context.Context, id string) (VNetGateway, error) {
	gw, err := c.vnetGatewayClient.Get(
		ctx, c.resourceGroup, id, nil,
	)
	if err != nil {
		return VNetGateway{}, fmt.Errorf(
			"failed to get VNet Gateway '%s': %w", id, err,
		)
	}
	return vnetGatewayFromAzure(gw.VirtualNetworkGateway), nil
}

func (c *Client) UpdateVNetGateway(ctx context.Context, gw VNetGateway) error {
	azGW := vnetGatewayToAzure(gw)
	marshalled, err := json.Marshal(azGW)
	if err != nil {
		return fmt.Errorf(
			"failed to marshall Virtual Network Gateway %v: %w", azGW, err,
		)
	}

	currentGW, err := c.vnetGatewayClient.Get(
		ctx, c.resourceGroup, gw.Name, nil,
	)
	if err != nil {
		return fmt.Errorf(
			"cannot update VNet Gateway '%s' as it's not found: %w", gw.Name, err,
		)
	}

	if err = json.Unmarshal(marshalled, &currentGW.VirtualNetworkGateway); err != nil {
		return fmt.Errorf(
			"failed to prepare updated structure by marshalling back to original: %w",
			err,
		)
	}
	future, err := c.vnetGatewayClient.BeginCreateOrUpdate(
		ctx,
		c.resourceGroup,
		gw.Name,
		currentGW.VirtualNetworkGateway,
		nil)
	if err != nil {
		return fmt.Errorf("cannot create virtual network gateway: %w", err)
	}

	_, err = future.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot get the virtual network gateway create or update future response: %w", err)
	}

	return nil
}

func (c *Client) ListVNetGateways(ctx context.Context) ([]VNetGateway, error) {
	pager := c.vnetGatewayClient.NewListPager(
		c.resourceGroup, nil,
	)
	if pager == nil {
		return nil, fmt.Errorf(
			"failed to get List Pager for VNet Gateways. Got nil object",
		)
	}
	output := []VNetGateway{}

	for pager.More() {
		gws, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain a page of VNet Gateways: %w", err)
		}
		for i := range gws.Value {
			if gws.Value[i] == nil {
				c.logger.Warnf("Got nil VNet Gateway on the Page: %v", gws)
				continue
			}
			output = append(output, vnetGatewayFromAzure(*gws.Value[i]))
		}

	}
	return output, nil
}

func (c *Client) GetLocalVNetGateway(ctx context.Context, id string) (LocalVNetGateway, error) {
	gw, err := c.localVnetGatewayClient.Get(
		ctx, c.resourceGroup, id, nil,
	)
	if err != nil {
		return LocalVNetGateway{}, fmt.Errorf(
			"failed to get Local Network Gateway '%s': %v", id, err,
		)
	}
	return localVnetGatewayFromAzure(gw.LocalNetworkGateway), nil
}

func (c *Client) ListLocalVNetGateways(ctx context.Context) ([]LocalVNetGateway, error) {
	pager := c.localVnetGatewayClient.NewListPager(
		c.resourceGroup, nil,
	)
	if pager == nil {
		return nil, fmt.Errorf(
			"failed to get List Pager for Local Network Gateways. Got nil object",
		)
	}
	output := []LocalVNetGateway{}

	for pager.More() {
		gws, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain a page of Local Network Gateways: %w", err)
		}
		for i := range gws.Value {
			if gws.Value[i] == nil {
				c.logger.Warnf("Got nil Local Network Gateway on the Page: %v", gws)
				continue
			}
			output = append(output, localVnetGatewayFromAzure(*gws.Value[i]))
		}

	}
	return output, nil
}

func (c *Client) CreateLocalVNetGateway(ctx context.Context, gw LocalVNetGateway) error {
	// TODO: Ensure that Create method checks if there is no resource existing already.
	// Otherwise, it may override some data.
	//
	// Eventually, make sure create method lacks some information necessary for update,
	// to avoid risk of sending a create to a resource that was created shortly after
	// ensuring it does not exist (race).

	azGW := localVnetGatewayToAzure(gw)
	azGW.Location = helper.StringToStringPointer(c.location)

	future, err := c.localVnetGatewayClient.BeginCreateOrUpdate(
		ctx,
		c.resourceGroup,
		gw.Name,
		azGW,
		nil)
	if err != nil {
		return fmt.Errorf("cannot create local network gateway: %w", err)
	}

	_, err = future.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot get the local network gateway create or update future response: %w", err)
	}

	return nil
}

func (c *Client) DeleteLocalVNetGateway(
	ctx context.Context, name string,
) error {
	future, err := c.localVnetGatewayClient.BeginDelete(ctx, c.resourceGroup, name, nil)
	if err != nil {
		return fmt.Errorf("cannot delete Local Network Gateway: %w", err)
	}
	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the Local Network Gateway delete future response: %w",
			err)
	}
	return nil
}

func (c *Client) ListNetworkGatewayConnections(ctx context.Context) ([]NetworkGatewayConnection, error) {
	pager := c.networkGatewayConnectionClient.NewListPager(
		c.resourceGroup, nil,
	)
	if pager == nil {
		return nil, fmt.Errorf(
			"failed to get List Pager for Network Gateway Connections. Got nil object",
		)
	}
	output := []NetworkGatewayConnection{}

	for pager.More() {
		gws, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain a page of Network Gateway Connections: %w", err)
		}
		for i := range gws.Value {
			if gws.Value[i] == nil {
				c.logger.Warnf("Got nil Network Gateway Connection on the Page: %v", gws)
				continue
			}
			output = append(output, networkGatewayFromAzure(*gws.Value[i]))
		}

	}
	return output, nil
}

func (c *Client) CreateNetworkGatewayConnection(
	ctx context.Context, conn NetworkGatewayConnection,
) error {
	azConn := networkGatewayToAzure(conn)
	azConn.Location = helper.StringToStringPointer(c.location)

	future, err := c.networkGatewayConnectionClient.BeginCreateOrUpdate(
		ctx,
		c.resourceGroup,
		conn.Name,
		azConn,
		nil)
	if err != nil {
		return fmt.Errorf("cannot create Network Gateway Connection: %w", err)
	}

	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the Network Gateway Connection create or update future response: %w",
			err)
	}

	return nil
}

func (c *Client) DeleteNetworkGatewayConnection(
	ctx context.Context, name string,
) error {
	future, err := c.networkGatewayConnectionClient.BeginDelete(ctx, c.resourceGroup, name, nil)
	if err != nil {
		return fmt.Errorf("cannot delete Network Gateway Connection: %w", err)
	}
	if _, err = future.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf(
			"cannot get the Network Gateway Connection delete future response: %w",
			err)
	}
	return nil
}

func (c *Client) GetPublicIP(ctx context.Context, id string) (PublicIP, error) {
	resp, err := c.publicIPClient.Get(
		ctx, c.resourceGroup, id, nil,
	)
	if err != nil {
		return PublicIP{}, fmt.Errorf(
			"failed to get a VNet '%s': %v", id, err,
		)
	}
	return publicIPFromAzure(resp.PublicIPAddress), nil
}
