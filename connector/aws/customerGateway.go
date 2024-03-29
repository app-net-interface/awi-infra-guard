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
	"strings"

	"github.com/app-net-interface/awi-infra-guard/connector/aws/client"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

func (c *AWSConnector) excludeDeletingOrDeletedCustomerGateways(
	gateways []client.CustomerGateway,
) []client.CustomerGateway {
	filtered := make([]client.CustomerGateway, 0, len(gateways))
	for _, gw := range gateways {
		state := strings.ToUpper(gw.State)
		if state == "DELETING" || state == "DELETED" {
			continue
		}
		filtered = append(filtered, gw)
	}
	return filtered
}

func (c *AWSConnector) getCustomerGatewayForIP(
	ctx context.Context, ip string,
) (*client.CustomerGateway, error) {
	gateways, err := c.awsClient.ListCustomerGateways(ctx, client.ListCustomerGatewayFilterIP{Value: ip})
	if err != nil {
		return nil, fmt.Errorf(
			"cannot verify current state of AWS Customer Gateways due to: %w", err,
		)
	}
	gateways = c.excludeDeletingOrDeletedCustomerGateways(gateways)
	if len(gateways) == 0 {
		return nil, nil
	}
	if len(gateways) > 1 {
		return nil, fmt.Errorf(
			"got multiple customer gateways pointing to the same external IP: %v",
			gateways,
		)
	}
	return &gateways[0], nil
}

func (c *AWSConnector) listCustomerGatewaysForTheConnection(
	ctx context.Context,
) ([]client.CustomerGateway, error) {
	gateways, err := c.awsClient.ListCustomerGateways(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot get existing Customer Gateways due to: %w", err,
		)
	}
	gateways = c.excludeDeletingOrDeletedCustomerGateways(gateways)

	output := make([]client.CustomerGateway, 0, len(gateways))
	for _, gw := range gateways {
		if !c.isResourceOwnedByConnection(gw.Tags) {
			continue
		}
		output = append(output, gw)
	}

	return output, nil
}

// Logs Customer Gateways which are assigned to this connection
// but are not reflecting any of currentInterfaces.
//
// Such situation can happen when the outside IP Address of
// second side of connection has changed - such Customer Gateway
// most likely should be deleted but this is a decision for the
// user.
func (c *AWSConnector) logRedundantCustomerGateways(
	ctx context.Context, currentInterfaces []string,
) error {
	interfaceMap := helper.SetFromSlice[string](currentInterfaces)

	gateways, err := c.listCustomerGatewaysForTheConnection(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot verify current state of AWS Customer Gateways due to: %w", err,
		)
	}

	for _, gw := range gateways {
		if !interfaceMap.Has(gw.IP) {
			c.logger.Errorf(
				"Found Customer Gateway %v assigned to this connection but it "+
					"doesn't match requested interfaces: %v. Please make sure "+
					"that Customer Gateway is necessary for the Gateway: %s",
				gw, currentInterfaces, c.transactionState.PeerGatewayName,
			)
			if c.isResourceOwnedByDifferentConnection(gw.Tags) {
				c.logger.Errorf(
					"Customer Gateway %s is also a part of different connections.",
					gw.ID,
				)
			} else {
				c.logger.Errorf(
					"Customer Gateway %s is not assigned to any other connection "+
						"so most likely it is a stale entry and can be removed (requires "+
						"confirmation)",
					gw.ID,
				)
			}
		}
	}
	return nil
}

func (c *AWSConnector) createCustomerGatewayIfNeeded(
	ctx context.Context, ip, asn string,
) (string, error) {
	gw, err := c.getCustomerGatewayForIP(ctx, ip)
	if err != nil {
		return "", fmt.Errorf(
			"found a problem while inspecting the presence for Customer Gateway for IP %s: %w",
			ip, err,
		)
	}
	if gw != nil {
		c.logger.Infof(
			"Customer Gateway for IP '%s' already exists under ID '%s'. Using it.",
			ip, gw.ID)

		if gw.ASN != asn {
			return "", fmt.Errorf(
				"customer Gateway '%s' for IP '%s', which already existed, has non matching ASN."+
					"Expected: %s, Got: %s",
				gw.ID, ip, asn, gw.ASN,
			)
		}
		return gw.ID, nil
	}

	c.logger.Debugf(
		"Customer Gateway for IP '%s' does not exist. Creating it.",
		ip)

	cgw, err := c.awsClient.CreateCustomerGateway(
		ctx,
		ip,
		asn,
		c.generateTagName())
	if err != nil {
		return "", fmt.Errorf(
			"could not create AWS Customer Gateway for IP '%s': %w", ip, err,
		)
	}
	if cgw == nil {
		return "", errors.New(
			"the Created Customer Gateway object is empty. Try restarting the script",
		)
	}
	return cgw.ID, nil
}

func (c *AWSConnector) createCustomerGatewaysIfNeeded(
	ctx context.Context, interfaces []string, asn string,
) ([]string, error) {
	c.logger.Debug("Checking existing AWS Customer Gateways")

	c.logRedundantCustomerGateways(ctx, interfaces)

	c.logger.Debugf(
		"The Gateway %s requires Customer Gateways for Gateway %s interfaces %v "+
			"to establish a connection. Checking if any already exist.",
		c.transactionState.GatewayName, c.transactionState.PeerGatewayName,
		interfaces,
	)

	customerGateways := []string{}
	for _, iface := range interfaces {
		gwID, err := c.createCustomerGatewayIfNeeded(
			ctx, iface, asn,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"failed handling Customer Gateway for IP %s: %w",
				iface, err,
			)
		}
		customerGateways = append(customerGateways, gwID)
	}

	c.logger.Debug("AWS Customer Gateways covered")
	return customerGateways, nil
}

func (c *AWSConnector) deleteCustomerGatewaysForConnection(ctx context.Context) error {
	c.logger.Debug("Looking for Customer Gateways that should be deleted")
	gateways, err := c.listCustomerGatewaysForTheConnection(ctx)
	if err != nil {
		return fmt.Errorf(
			"cannot get existing Customer Gateways due to: %w", err,
		)
	}

	for _, gw := range gateways {
		if c.isResourceOwnedByDifferentConnection(gw.Tags) {
			c.logger.Debugf(
				"Customer Gateway '%s' is used by different connections. It won't be deleted.",
				gw.ID)
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
