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
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// This file contains structures required for obtaining the details
// of VPN Connections from AWS Provider.

type XMLVpnConnection struct {
	XMLName           xml.Name         `xml:"vpn_connection"`
	ID                string           `xml:"id,attr"`
	CustomerGatewayID string           `xml:"customer_gateway_id"`
	VpnGatewayID      string           `xml:"vpn_gateway_id"`
	VpnConnectionType string           `xml:"vpn_connection_type"`
	IpsecTunnels      []XMLIpsecTunnel `xml:"ipsec_tunnel"`
}

func (xml XMLVpnConnection) String() string {
	b, err := json.Marshal(xml)
	if err != nil {
		return "XMLVpnConnection{<cannot parse configuration>}"
	}
	return fmt.Sprintf("XMLVpnConnection {%s}", string(b))
}

type XMLIpsecTunnel struct {
	CustomerGateway XMLCustomerGateway `xml:"customer_gateway"`
	VpnGateway      XMLVpnGateway      `xml:"vpn_gateway"`
	Ike             XMLIke             `xml:"ike"`
	Ipsec           XMLIpsec           `xml:"ipsec"`
}

func (xml XMLIpsecTunnel) String() string {
	b, err := json.Marshal(xml)
	if err != nil {
		return "XMLVpnConnection{<cannot parse configuration>}"
	}
	return fmt.Sprintf("XMLVpnConnection {%s}", string(b))
}

type XMLCustomerGateway struct {
	TunnelOutsideAddress XMLAddress `xml:"tunnel_outside_address"`
	TunnelInsideAddress  XMLAddress `xml:"tunnel_inside_address"`
	Bgp                  XMLBgp     `xml:"bgp"`
}

type XMLVpnGateway struct {
	TunnelOutsideAddress XMLAddress `xml:"tunnel_outside_address"`
	TunnelInsideAddress  XMLAddress `xml:"tunnel_inside_address"`
	Bgp                  XMLBgp     `xml:"bgp"`
}

type XMLAddress struct {
	IPAddress   string `xml:"ip_address"`
	NetworkMask string `xml:"network_mask,omitempty"`
	NetworkCIDR string `xml:"network_cidr,omitempty"`
}

type XMLBgp struct {
	Asn      string `xml:"asn"`
	HoldTime string `xml:"hold_time"`
}

type XMLIke struct {
	AuthenticationProtocol string `xml:"authentication_protocol"`
	EncryptionProtocol     string `xml:"encryption_protocol"`
	Lifetime               string `xml:"lifetime"`
	PerfectForwardSecrecy  string `xml:"perfect_forward_secrecy"`
	Mode                   string `xml:"mode"`
	PreSharedKey           string `xml:"pre_shared_key"`
}

type XMLIpsec struct {
	Protocol                      string               `xml:"protocol"`
	AuthenticationProtocol        string               `xml:"authentication_protocol"`
	EncryptionProtocol            string               `xml:"encryption_protocol"`
	Lifetime                      string               `xml:"lifetime"`
	PerfectForwardSecrecy         string               `xml:"perfect_forward_secrecy"`
	Mode                          string               `xml:"mode"`
	ClearDfBit                    string               `xml:"clear_df_bit"`
	FragmentationBeforeEncryption string               `xml:"fragmentation_before_encryption"`
	TcpMssAdjustment              string               `xml:"tcp_mss_adjustment"`
	DeadPeerDetection             XMLDeadPeerDetection `xml:"dead_peer_detection"`
}

type XMLDeadPeerDetection struct {
	Interval string `xml:"interval"`
	Retries  string `xml:"retries"`
}

func VPNConnectionConfigToObject(config string) (*XMLVpnConnection, error) {
	vpnConnection := XMLVpnConnection{}
	if err := xml.Unmarshal([]byte(config), &vpnConnection); err != nil {
		return nil, fmt.Errorf("failed to process VPN Connection Configuration: %w", err)
	}
	return &vpnConnection, nil
}
