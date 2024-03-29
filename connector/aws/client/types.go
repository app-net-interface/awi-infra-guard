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

package client

import (
	"fmt"
	"strconv"

	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type TransitGateway struct {
	ID string
	// TODO: Replace VPC ID with the list of IDs.
	VPCID string
	ASN   string

	// All these options will be set to enable when creating
	// a connection with a BGP.
	//
	// This is required to automatically accept routes discovered
	// from the second side of connection.
	AutoAcceptSharedAttachments  bool
	DefaultRouteTableAssociation bool
	DefaultRouteTablePropagation bool
	VpnEcmpSupport               bool
	DnsSupport                   bool
}

func (tgw TransitGateway) String() string {
	return fmt.Sprintf(
		"TransitGateway {ID: '%s', VPCID: '%s', ASN: '%s', "+
			"AutoAcceptSharedAttachments: %v, "+
			"DefaultRouteTableAssociation: %v, "+
			"DefaultRouteTablePropagation: %v, "+
			"VpnEcmpSupport: %v, "+
			"DnsSupport: %v}",
		tgw.ID, tgw.VPCID, tgw.ASN, tgw.AutoAcceptSharedAttachments,
		tgw.DefaultRouteTableAssociation, tgw.DefaultRouteTablePropagation,
		tgw.VpnEcmpSupport, tgw.DnsSupport,
	)
}

func transitGatewayFromAWS(tgw *types.TransitGateway) *TransitGateway {
	if tgw == nil {
		return nil
	}
	options := tgw.Options
	if options == nil {
		options = &types.TransitGatewayOptions{}
	}
	return &TransitGateway{
		ID:                           helper.StringPointerToString(tgw.TransitGatewayId),
		ASN:                          helper.Int64PointerToString(options.AmazonSideAsn),
		AutoAcceptSharedAttachments:  options.AutoAcceptSharedAttachments == types.AutoAcceptSharedAttachmentsValueEnable,
		DefaultRouteTableAssociation: options.DefaultRouteTableAssociation == types.DefaultRouteTableAssociationValueEnable,
		DefaultRouteTablePropagation: options.DefaultRouteTablePropagation == types.DefaultRouteTablePropagationValueEnable,
		VpnEcmpSupport:               options.VpnEcmpSupport == types.VpnEcmpSupportValueEnable,
		DnsSupport:                   options.DnsSupport == types.DnsSupportValueEnable,
	}
}

func transitGatewayToModifyOptionsAWS(tgw *TransitGateway) *types.ModifyTransitGatewayOptions {
	if tgw == nil {
		return nil
	}
	options := &types.ModifyTransitGatewayOptions{}

	if tgw.AutoAcceptSharedAttachments {
		options.AutoAcceptSharedAttachments = types.AutoAcceptSharedAttachmentsValueEnable
	} else {
		options.AutoAcceptSharedAttachments = types.AutoAcceptSharedAttachmentsValueDisable
	}

	if tgw.DefaultRouteTableAssociation {
		options.DefaultRouteTableAssociation = types.DefaultRouteTableAssociationValueEnable
	} else {
		options.DefaultRouteTableAssociation = types.DefaultRouteTableAssociationValueDisable
	}

	if tgw.DefaultRouteTablePropagation {
		options.DefaultRouteTablePropagation = types.DefaultRouteTablePropagationValueEnable
	} else {
		options.DefaultRouteTablePropagation = types.DefaultRouteTablePropagationValueDisable
	}

	if tgw.VpnEcmpSupport {
		options.VpnEcmpSupport = types.VpnEcmpSupportValueEnable
	} else {
		options.VpnEcmpSupport = types.VpnEcmpSupportValueDisable
	}

	if tgw.DnsSupport {
		options.DnsSupport = types.DnsSupportValueEnable
	} else {
		options.DnsSupport = types.DnsSupportValueDisable
	}

	asn, err := strconv.ParseInt(tgw.ASN, 10, 64)
	if err != nil {
		asn = 0
	}
	options.AmazonSideAsn = &asn

	return options
}

type CustomerGateway struct {
	ID    string
	ASN   string
	IP    string
	State string
	Tags  []Tag
}

func (cgw CustomerGateway) String() string {
	return fmt.Sprintf(
		"CustomerGateway {ID: '%s', ASN: '%s', IP: '%s', State: '%s'}",
		cgw.ID, cgw.ASN, cgw.IP, cgw.State,
	)
}

func customerGatewayFromAWS(cgw *types.CustomerGateway) *CustomerGateway {
	if cgw == nil {
		return nil
	}
	return &CustomerGateway{
		ID:    helper.StringPointerToString(cgw.CustomerGatewayId),
		IP:    helper.StringPointerToString(cgw.IpAddress),
		ASN:   helper.StringPointerToString(cgw.BgpAsn),
		State: helper.StringPointerToString(cgw.State),
		Tags:  tagsFromAWS(cgw.Tags),
	}
}

type VPNConnectionTunnelOption struct {
	// TODO: Improve tunnel options to add additional
	// layers of security.
	CIDR string
}

func (vpn VPNConnectionTunnelOption) String() string {
	return fmt.Sprintf(
		"VPNConnectionTunnelOption {CIDR: '%s'}", vpn.CIDR,
	)
}

type VPNConnection struct {
	ID                string
	CustomerGatewayID string
	TransitGatewayID  string
	TunnelOptions     []VPNConnectionTunnelOption
	Configuration     string
	Tags              []Tag
}

func (vpn VPNConnection) String() string {
	return fmt.Sprintf(
		"VPNConnection {ID: '%s', CustomerGatewayID: '%s', TransitGatewayID: '%s', "+
			"TunnelOptions: %v, Configuration: '%s'}",
		vpn.ID, vpn.CustomerGatewayID, vpn.TransitGatewayID, vpn.TunnelOptions,
		vpn.Configuration,
	)
}

func vpnConnectionFromAWS(vc *types.VpnConnection) *VPNConnection {
	if vc == nil {
		return nil
	}
	tunnelOptions := []VPNConnectionTunnelOption{}

	if vc.Options != nil {
		for _, option := range vc.Options.TunnelOptions {
			tunnelOptions = append(tunnelOptions, VPNConnectionTunnelOption{
				// TODO: Verify if we need to handle IPv6 CIDR
				CIDR: helper.StringPointerToString(option.TunnelInsideCidr),
			})
		}
	}

	return &VPNConnection{
		ID:                helper.StringPointerToString(vc.VpnConnectionId),
		CustomerGatewayID: helper.StringPointerToString(vc.CustomerGatewayId),
		TransitGatewayID:  helper.StringPointerToString(vc.TransitGatewayId),
		TunnelOptions:     tunnelOptions,
		Configuration:     helper.StringPointerToString(vc.CustomerGatewayConfiguration),
		Tags:              tagsFromAWS(vc.Tags),
	}
}

type Tag struct {
	Key   string
	Value string
}

func tagsFromAWS(tags []types.Tag) []Tag {
	out := []Tag{}
	for _, t := range tags {
		out = append(out, Tag{
			Key:   helper.StringPointerToString(t.Key),
			Value: helper.StringPointerToString(t.Value),
		})
	}
	return out
}

type VPC struct {
	Name string
	CIDR string
}

func vpcFromAWS(vpc *types.Vpc) *VPC {
	if vpc == nil {
		return nil
	}
	return &VPC{
		Name: helper.StringPointerToString(vpc.VpcId),
		CIDR: helper.StringPointerToString(vpc.CidrBlock),
	}
}
