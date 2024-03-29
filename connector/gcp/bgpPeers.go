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
	"strconv"

	"github.com/app-net-interface/awi-infra-guard/connector/gcp/client"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

func (c *GCPConnector) getDesiredBGPPeersForCloudRouter(
	connectConfig types.CreateBGPConnectionConfig,
	cloudRouterInterfaces []client.CloudRouterInterface,
) ([]client.CloudRouterBGPPeer, error) {
	if len(cloudRouterInterfaces) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of provided Cloud Router Interfaces: %v. Expected %d interfaces",
			cloudRouterInterfaces, connectConfig.NumberOfTunnels)
	}

	addresses := connectConfig.PeerBGPAddresses
	if len(addresses) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of provided Peer IP Addresses: %v. Expected %d addresses",
			addresses, connectConfig.NumberOfTunnels)
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
// BGP Session between GCP and the second provider.
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
	if len(desiredBGPPeers) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of BGP Peers prepared for Cloud Router '%s': %v. "+
				"Expected %d BGP peers",
			cloudRouterName, desiredBGPPeers, connectConfig.NumberOfTunnels,
		)
	}
	c.logger.Debugf(
		"The Cloud Router '%s' requires following BGP Peers: %v",
		cloudRouterName, router.BGPPeers,
	)

	actualBGPPeers, missingBGPPeers, err := c.prepareCloudRouterBGPPeers(
		router, desiredBGPPeers,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to determine missing BGP Peers in the Cloud Router '%s': %w",
			cloudRouterName, err)
	}
	if len(actualBGPPeers) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of BGP Peers extracted from Cloud Router '%s': %v. "+
				"Expected %d BGP Peers",
			cloudRouterName, actualBGPPeers, connectConfig.NumberOfTunnels,
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
