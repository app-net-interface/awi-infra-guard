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

	"github.com/app-net-interface/awi-infra-guard/connector/gcp/client"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

type transactionState struct {
	state creationState

	// Input data
	CloudRouterURL  string
	VPNGatewayURL   string
	GatewayName     string
	PeerGatewayName string
	OwnInterfaces   []string

	// SetBeforeCreation
	BGPConnectionConfig types.CreateBGPConnectionConfig
	OwnBGPAddresses     []string

	// SetDuringCreation
	ExternalVPNGatewayURL string
	VPNTunnels            []client.VPNTunnel
	CloudRouterInterfaces []client.CloudRouterInterface
	CloudRouterBGPPeers   []client.CloudRouterBGPPeer
}

func NewTransactionState(
	cloudRouterURL string,
	vpnGatewayURL string,
	gatewayName string,
	peerGatewayName string,
	ownInterfaces []string,
) *transactionState {
	return &transactionState{
		state:           stateCreateExternalVPNGateway,
		CloudRouterURL:  cloudRouterURL,
		VPNGatewayURL:   vpnGatewayURL,
		GatewayName:     gatewayName,
		PeerGatewayName: peerGatewayName,
		OwnInterfaces:   ownInterfaces,
	}
}

type creationState uint8

const (
	stateCreateExternalVPNGateway creationState = iota
	stateCreateVPNTunnels
	stateCreateBGPInterfaces
	stateCreateBGPPeers
	stateCreationCompleted
)

func getNextState(current creationState) (creationState, error) {
	switch current {
	case stateCreateExternalVPNGateway:
		return stateCreateVPNTunnels, nil
	case stateCreateVPNTunnels:
		return stateCreateBGPInterfaces, nil
	case stateCreateBGPInterfaces:
		return stateCreateBGPPeers, nil
	case stateCreateBGPPeers:
		return stateCreationCompleted, nil
	case stateCreationCompleted:
		return stateCreationCompleted, errors.New(
			"the creation was already completed. Nothing more to do",
		)
	}
	return current, fmt.Errorf(
		"unhandled creation state: %v", current,
	)
}

type StateOperation func(ctx context.Context) error

func (c *GCPConnector) CreateExternalVPNGateway(ctx context.Context) error {
	if c.state.state != stateCreateExternalVPNGateway {
		c.logger.Debugf(
			"creating state is not ExternalVPNGateway so skipping this stage. Current state: %v",
			c.state.state,
		)
		return nil
	}
	if c.state.ExternalVPNGatewayURL != "" {
		c.logger.Infof("External VPN Gateway already created: %s", c.state.ExternalVPNGatewayURL)
		return nil
	}
	extVPNGateway, err := c.createExternalVPNGateway(ctx, c.state.BGPConnectionConfig.OutsideInterfaces)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare External VPN Gateway: %w", err,
		)
	}
	if extVPNGateway == nil {
		return errors.New(
			"unexpected empty External VPN Gateway object",
		)
	}
	c.state.ExternalVPNGatewayURL = extVPNGateway.URL
	c.logger.Debugf("Using External VPN Gateway: %v", *extVPNGateway)
	return nil
}

func (c *GCPConnector) CreateVPNTunnels(ctx context.Context) error {
	if c.state.state != stateCreateVPNTunnels {
		c.logger.Debugf(
			"creating state is not VPNTunnels so skipping this stage. Current state: %v",
			c.state.state,
		)
		return nil
	}
	if len(c.state.VPNTunnels) > 0 {
		c.logger.Infof("VPN Tunnels already created: %v", c.state.VPNTunnels)
		return nil
	}

	vpnTunnels, err := c.createVPNTunnels(
		ctx,
		c.state.CloudRouterURL,
		c.state.VPNGatewayURL,
		c.state.ExternalVPNGatewayURL,
		c.state.BGPConnectionConfig.SharedSecrets,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare VPN Tunnels: %w", err,
		)
	}
	if len(vpnTunnels) == 0 {
		return errors.New(
			"expected VPN tunnels created. Got nothing",
		)
	}

	c.state.VPNTunnels = vpnTunnels
	c.logger.Debugf("Using following VPN Tunnels: %v", vpnTunnels)
	return nil
}

func (c *GCPConnector) CreateCloudRouterInterfaces(ctx context.Context) error {
	if c.state.state != stateCreateBGPInterfaces {
		c.logger.Debugf(
			"creating state is not BGPInterface so skipping this stage. Current state: %v",
			c.state.state,
		)
		return nil
	}
	if len(c.state.CloudRouterInterfaces) > 0 {
		c.logger.Infof("Cloud Router Interfaces already created: %v", c.state.CloudRouterInterfaces)
		return nil
	}

	cloudRouterInterfaces, err := c.createCloudRouterInterfacesIfNeeded(
		ctx, c.state.BGPConnectionConfig, c.state.VPNTunnels,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare Cloud Router Interfaces: %w", err,
		)
	}
	if cloudRouterInterfaces == nil {
		return errors.New(
			"expected Cloud Router interfaces. Got nothing",
		)
	}

	c.state.CloudRouterInterfaces = cloudRouterInterfaces
	return nil
}

func (c *GCPConnector) CreateCloudRouterBGPPeers(ctx context.Context) error {
	if c.state.state != stateCreateBGPPeers {
		c.logger.Debugf(
			"creating state is not BGPPeers so skipping this stage. Current state: %v",
			c.state.state,
		)
		return nil
	}
	if len(c.state.CloudRouterBGPPeers) > 0 {
		c.logger.Infof("Cloud Router BGP Peers already created: %v", c.state.CloudRouterBGPPeers)
		return nil
	}

	cloudRouterPeers, err := c.createCloudRouterBPGPeersIfNeeded(
		ctx, c.state.BGPConnectionConfig, c.state.CloudRouterInterfaces,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare Cloud Router BGP Peers: %w", err,
		)
	}
	if cloudRouterPeers == nil {
		return errors.New(
			"expected Cloud Router BGP Peers. Got nothing",
		)
	}

	c.state.CloudRouterBGPPeers = cloudRouterPeers
	return nil
}
