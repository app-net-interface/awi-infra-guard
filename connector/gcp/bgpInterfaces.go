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

func (c *GCPConnector) getDesiredInterfacesForCloudRouter(
	connectConfig types.CreateBGPConnectionConfig,
	vpnTunnels []client.VPNTunnel,
) ([]client.CloudRouterInterface, error) {
	if len(vpnTunnels) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of provided VPN Tunnels: %v. Expected %d tunnels",
			vpnTunnels, connectConfig.NumberOfTunnels)
	}
	addresses := c.state.OwnBGPAddresses
	if len(addresses) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of provided BGP Tunnel IP Addresses: %v. Expected %d addresses",
			addresses, connectConfig.NumberOfTunnels)
	}

	maskSuffix := fmt.Sprintf("/%d", 30)

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
		"The Cloud Router '%s' will expose the following %d interfaces for connection purposes: %v. "+
			"The following interfaces are not created yet and will be created shortly afterwards: %v",
		cloudRouter.Name, connectConfig.NumberOfTunnels, interfaces, missingInterfaces,
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
	if len(desiredInterfaces) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of interfaces prepared for Cloud Router '%s': %v. "+
				"Expected %d interfaces",
			cloudRouterName, desiredInterfaces, connectConfig.NumberOfTunnels,
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
	if len(actualInterfaces) != int(connectConfig.NumberOfTunnels) {
		return nil, fmt.Errorf(
			"unexpected number of interfaces extracted from Cloud Router '%s': %v. "+
				"Expected %d interfaces",
			cloudRouterName, actualInterfaces, connectConfig.NumberOfTunnels,
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
