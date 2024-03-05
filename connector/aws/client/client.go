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
	"errors"
	"fmt"
	"strconv"

	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sirupsen/logrus"
)

// TODO: Add multi-region support to AWS Client.
type Client struct {
	region           string
	defaultAccountID string
	awsClient        *ec2.Client
	logger           *logrus.Entry
}

func NewClient(ctx context.Context, logger *logrus.Entry, region string) (*Client, error) {
	logger.Debugf("Creating a new AWS Client for region: %s", region)
	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get AWS config: %v", err)
	}
	cfg.Region = region
	stsclient := sts.NewFromConfig(cfg)
	var defaultAccountID string
	req, err := stsclient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		logger.Errorf("Failed to determine Account ID for default profile: %v", err)
		// TODO: Consider returning here as it may indicate that the
		// client has no required permissions or that authentication failed.
	} else {
		defaultAccountID = helper.StringPointerToString(req.Account)
	}

	client := ec2.NewFromConfig(cfg)
	if client == nil {
		return nil, fmt.Errorf("got nil AWS client")
	}

	logger.Debug("AWS Client created successfully")
	return &Client{
		region:           region,
		defaultAccountID: defaultAccountID,
		awsClient:        client,
		logger:           logger,
	}, nil
}

func (c *Client) GetTransitGateway(ctx context.Context, name string) (*TransitGateway, error) {
	gateways, err := c.ListTransitGateways(ctx)
	if err != nil {
		return nil, err
	}
	for _, g := range gateways {
		if g == nil {
			c.logger.Warnf(
				"Got Empty Transit Gateway while obtaining the list of Transit Gateways")
			continue
		}
		if g.ID == name {
			return g, nil
		}
	}
	return nil, fmt.Errorf(
		"cannot find Transit Gateway with name '%s' in AWS Region '%s'",
		name, c.region,
	)
}

// TODO: Verify if modification doesn't require recreating existing Transit
// Gateway attachments.
func (c *Client) UpdateTransitGateway(ctx context.Context, tgw TransitGateway) error {
	_, err := c.awsClient.ModifyTransitGateway(
		ctx,
		&ec2.ModifyTransitGatewayInput{
			TransitGatewayId: &tgw.ID,
			Options:          transitGatewayToModifyOptionsAWS(&tgw),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to update Transit Gateway %s: %w", tgw.ID, err)
	}
	return nil
}

func (c *Client) ListTransitGateways(ctx context.Context) ([]*TransitGateway, error) {
	gateways := map[string]*TransitGateway{}

	output, err := c.awsClient.DescribeTransitGateways(ctx, &ec2.DescribeTransitGatewaysInput{})
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, errors.New(
			"unexpected empty Transit Gateways after attempting to Describe Transit Gateways")
	}

	for i, gw := range output.TransitGateways {
		gateways[helper.StringPointerToString(gw.TransitGatewayId)] = transitGatewayFromAWS(&output.TransitGateways[i])
	}

	vpcAttachments, err := c.awsClient.DescribeTransitGatewayVpcAttachments(ctx, &ec2.DescribeTransitGatewayVpcAttachmentsInput{})
	if err != nil {
		return nil, err
	}
	if vpcAttachments == nil {
		return nil, errors.New(
			"unexpected empty VPC Attachments after attempting to Describe " +
				"Transit Gateway Vpc Attachments")
	}

	for _, attachment := range vpcAttachments.TransitGatewayVpcAttachments {
		if attachment.TransitGatewayId == nil {
			c.logger.Warnf(
				"Transit Gateway Attachment '%s' is not attached to any Transit Gateway. Ignoring it",
				helper.StringPointerToString(attachment.TransitGatewayAttachmentId),
			)
		}
		if attachment.State != types.TransitGatewayAttachmentStateAvailable {
			continue
		}
		tgw, ok := gateways[*attachment.TransitGatewayId]
		if !ok {
			return nil, fmt.Errorf(
				"found Transit Gateway Attachment '%s' to non existing Transit Gateway '%s'",
				helper.StringPointerToString(attachment.TransitGatewayAttachmentId),
				*attachment.TransitGatewayId,
			)
		}
		tgw.VPCID = helper.StringPointerToString(attachment.VpcId)
	}

	gatewaysList := []*TransitGateway{}
	for _, tgw := range gateways {
		gatewaysList = append(gatewaysList, tgw)
	}

	return gatewaysList, nil
}

func (c *Client) GetVPC(ctx context.Context, name string) (*VPC, error) {
	vpcs, err := c.awsClient.DescribeVpcs(
		ctx,
		&ec2.DescribeVpcsInput{
			VpcIds: []string{name},
		},
	)
	if err != nil {
		return nil, err
	}
	if vpcs == nil {
		return nil, errors.New(
			"unexpected empty VPC after attempting to Describe VPCs")
	}
	if len(vpcs.Vpcs) == 0 {
		return nil, fmt.Errorf(
			"VPC %s not found", name,
		)
	}
	if len(vpcs.Vpcs) > 1 {
		return nil, fmt.Errorf(
			"unexpectedly there is more than 1 VPC with name %s", name,
		)
	}
	return vpcFromAWS(&vpcs.Vpcs[0]), nil
}

func (c *Client) ListCustomerGateways(ctx context.Context, filters ...ListCustomerGatewayFilter) ([]CustomerGateway, error) {
	var awsFilters []types.Filter
	for _, f := range filters {
		awsFilters = append(awsFilters, f.GetCustomerGatewayListFilter())
	}

	output, err := c.awsClient.DescribeCustomerGateways(ctx, &ec2.DescribeCustomerGatewaysInput{
		Filters: awsFilters,
	})
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, errors.New(
			"unexpected empty Customer Gateways after attempting to Describe Customer Gateways")
	}

	gateways := []CustomerGateway{}
	for _, cgw := range output.CustomerGateways {
		customerGateway := customerGatewayFromAWS(&cgw)
		if customerGateway != nil {
			gateways = append(gateways, *customerGateway)
		} else {
			c.logger.Warn("Got empty Customer Gateway object")
		}
	}

	return gateways, nil
}

func (c *Client) CreateCustomerGateway(ctx context.Context, ip, asn, tag string) (*CustomerGateway, error) {
	c.logger.Debugf(
		"Creating Customer Gateway with IP '%s' and ASN '%s'", ip, asn)
	converted, err := strconv.Atoi(asn)
	if err != nil {
		return nil, err
	}
	asn32 := int32(converted)
	output, err := c.awsClient.CreateCustomerGateway(ctx, &ec2.CreateCustomerGatewayInput{
		Type:              types.GatewayTypeIpsec1,
		PublicIp:          &ip,
		BgpAsn:            &asn32,
		TagSpecifications: addTag(types.ResourceTypeCustomerGateway, tag),
	})
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, fmt.Errorf(
			"unexpected empty Customer Gateway object from creating Customer Gateway with "+
				"Public IP '%s' and ASN '%s'", ip, asn,
		)
	}
	if output.CustomerGateway == nil {
		return nil, nil
	}
	cgw := customerGatewayFromAWS(output.CustomerGateway)
	c.logger.Debugf(
		"Customer Gateway created successfully with ID '%s'. State: %s", cgw.ID, cgw.State)
	return cgw, nil
}

func (c *Client) DeleteCustomerGateway(ctx context.Context, id string) error {
	_, err := c.awsClient.DeleteCustomerGateway(ctx, &ec2.DeleteCustomerGatewayInput{
		CustomerGatewayId: &id,
	})
	if err != nil {
		return fmt.Errorf("failed to delete Customer Gateway: %w", err)
	}
	return nil
}

type TunnelOption struct {
	CIDR         string
	PreSharedKey string
}

func (c *Client) ListVPNConnections(ctx context.Context, filters ...ListVPNConnectionFilter) ([]*VPNConnection, error) {
	var awsFilters []types.Filter
	for _, f := range filters {
		awsFilters = append(awsFilters, f.GetVPNConnectionListFilter())
	}

	output, err := c.awsClient.DescribeVpnConnections(ctx, &ec2.DescribeVpnConnectionsInput{
		Filters: awsFilters,
	})
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, errors.New(
			"unexpected empty VPN Connections after attempting to Describe VPN Connections")
	}
	connections := []*VPNConnection{}
	for _, conn := range output.VpnConnections {
		connections = append(connections, vpnConnectionFromAWS(&conn))
	}
	return connections, nil
}

func (c *Client) CreateVPNConnection(
	ctx context.Context,
	customerGatewayID,
	transitGatewayID string,
	tunnelOptions [2]TunnelOption,
	tag string,
) (*VPNConnection, error) {
	c.logger.Debugf(
		"Creating VPN Connection for Customer Gateway '%s' and Transit Gateway '%s' with tunnel options %v",
		customerGatewayID, transitGatewayID, tunnelOptions)
	vpnType := "ipsec.1"
	vpnTunnelOptions := []types.VpnTunnelOptionsSpecification{}
	for i := range tunnelOptions {
		vpnTunnelOptions = append(vpnTunnelOptions, types.VpnTunnelOptionsSpecification{
			TunnelInsideCidr: &tunnelOptions[i].CIDR,
			PreSharedKey:     &tunnelOptions[i].PreSharedKey,
		})
	}
	output, err := c.awsClient.CreateVpnConnection(ctx, &ec2.CreateVpnConnectionInput{
		Type:              &vpnType,
		CustomerGatewayId: &customerGatewayID,
		TransitGatewayId:  &transitGatewayID,
		Options: &types.VpnConnectionOptionsSpecification{
			TunnelOptions: vpnTunnelOptions,
		},

		TagSpecifications: addTag(types.ResourceTypeVpnConnection, tag),
	})
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, fmt.Errorf(
			"unexpected empty VPN Connection object after attempting to create VPN Connection "+
				"for Customer Gateway '%s' and VPN Gateway '%s'", customerGatewayID, transitGatewayID,
		)
	}
	vpnConnection := vpnConnectionFromAWS(output.VpnConnection)
	c.logger.Debugf(
		"Successfully created VPN Connection with ID '%s'", vpnConnection.ID)
	return vpnConnection, nil
}

func (c *Client) DeleteVPNConnection(ctx context.Context, id string) error {
	_, err := c.awsClient.DeleteVpnConnection(ctx, &ec2.DeleteVpnConnectionInput{
		VpnConnectionId: &id,
	})
	if err != nil {
		return fmt.Errorf("failed to delete VPN Connection: %w", err)
	}
	return nil
}

func (c *Client) GetDeviceConfiguration(ctx context.Context, vpnConnectionId string) (*string, error) {
	vpnType := "ipsec.1"
	config, err := c.awsClient.GetVpnConnectionDeviceSampleConfiguration(ctx, &ec2.GetVpnConnectionDeviceSampleConfigurationInput{
		VpnConnectionId:           &vpnConnectionId,
		VpnConnectionDeviceTypeId: &vpnType,
	})
	if err != nil {
		return nil, err
	}
	return config.VpnConnectionDeviceSampleConfiguration, nil
}

func addTag(resType types.ResourceType, key string) []types.TagSpecification {
	val := ""
	return []types.TagSpecification{
		{
			ResourceType: resType,
			Tags: []types.Tag{
				{
					Key:   &key,
					Value: &val,
				},
			},
		},
	}
}
