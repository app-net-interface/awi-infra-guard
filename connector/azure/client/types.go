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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

type VNet struct {
	ID              string
	Name            string
	AddressPrefixes []string
}

func vNetFromAzure(vnet armnetwork.VirtualNetwork) VNet {
	var addressPrefixes []string
	if vnet.Properties != nil && vnet.Properties.AddressSpace != nil {
		addressPrefixes = helper.SliceStringPointersToSliceStrings(
			vnet.Properties.AddressSpace.AddressPrefixes,
		)
	}
	return VNet{
		ID:              helper.StringPointerToString(vnet.ID),
		Name:            helper.StringPointerToString(vnet.Name),
		AddressPrefixes: addressPrefixes,
	}
}

type ResourceGroup struct {
	Location string
}

type BGPAddress struct {
	DefaultAddresses []string
	CustomAddresses  []string
	ConfigurationID  string
}

func bgpAddressFromAzure(addr armnetwork.IPConfigurationBgpPeeringAddress) BGPAddress {
	output := BGPAddress{
		ConfigurationID: helper.StringPointerToString(addr.IPConfigurationID),
	}
	for i := range addr.DefaultBgpIPAddresses {
		if addr.DefaultBgpIPAddresses[i] == nil {
			continue
		}
		output.DefaultAddresses = append(output.DefaultAddresses, *addr.DefaultBgpIPAddresses[i])
	}
	for i := range addr.CustomBgpIPAddresses {
		if addr.CustomBgpIPAddresses[i] == nil {
			continue
		}
		output.CustomAddresses = append(output.CustomAddresses, *addr.CustomBgpIPAddresses[i])
	}
	return output
}

type VNetGateway struct {
	ID   string
	Name string
	ASN  string
	// If set to false it means that the VPN Gateway provisioning
	// failed, vpn Gateway is being deleted or it is in the update
	// process and it should not be modified at that particular
	// moment.
	ReadyToUse bool
	// TODO: Handle Public IPs that may be in use by other
	// CSP Connections. Find out if this is allowed/accepted.
	PublicIPsIDs []string

	Addresses []string

	BGPAddresses []BGPAddress

	VNet string
}

func vnetNameFromVNetGwIPConfiguration(
	config *armnetwork.VirtualNetworkGatewayIPConfiguration,
) string {
	if config == nil {
		return ""
	}
	if config.Properties == nil {
		return ""
	}
	if config.Properties.Subnet == nil {
		return ""
	}
	if config.Properties.Subnet.ID == nil {
		return ""
	}
	chunks := strings.Split(*config.Properties.Subnet.ID, "/")
	if len(chunks) < 3 {
		return ""
	}
	return chunks[len(chunks)-3]
}

func vnetGatewayFromAzure(gw armnetwork.VirtualNetworkGateway) VNetGateway {
	output := VNetGateway{
		ID:   helper.StringPointerToString(gw.ID),
		Name: helper.StringPointerToString(gw.Name),
	}
	properties := gw.Properties
	if properties == nil {
		return output
	}
	for i := range properties.IPConfigurations {
		if properties.IPConfigurations[i] == nil {
			continue
		}
		if properties.IPConfigurations[i].Properties == nil {
			continue
		}
		if properties.IPConfigurations[i].Properties.PublicIPAddress != nil {
			output.PublicIPsIDs = append(output.PublicIPsIDs,
				*properties.IPConfigurations[i].Properties.PublicIPAddress.ID,
			)
		}
	}

	if properties.CustomRoutes != nil {
		for i := range properties.CustomRoutes.AddressPrefixes {
			output.Addresses =
				append(
					output.Addresses,
					helper.StringPointerToString(properties.CustomRoutes.AddressPrefixes[i]))
		}
	}

	if properties.ProvisioningState != nil {
		// Consider checking Active flag.
		output.ReadyToUse =
			*properties.ProvisioningState == armnetwork.ProvisioningStateSucceeded
	}

	if len(properties.IPConfigurations) > 0 {
		output.VNet = vnetNameFromVNetGwIPConfiguration(properties.IPConfigurations[0])
	}

	if properties.BgpSettings == nil {
		return output
	}
	bgpSettings := properties.BgpSettings

	output.ASN = helper.Int64PointerToString(bgpSettings.Asn)
	for i := range bgpSettings.BgpPeeringAddresses {
		if bgpSettings.BgpPeeringAddresses[i] == nil {
			continue
		}
		output.BGPAddresses = append(
			output.BGPAddresses,
			bgpAddressFromAzure(*bgpSettings.BgpPeeringAddresses[i]))
	}

	return output
}

func vnetGatewayToAzure(gw VNetGateway) armnetwork.VirtualNetworkGateway {
	output := armnetwork.VirtualNetworkGateway{
		ID:   helper.StringToStringPointer(gw.ID),
		Name: helper.StringToStringPointer(gw.Name),
		Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
			BgpSettings: &armnetwork.BgpSettings{
				Asn: helper.StringToInt64Pointer(gw.ASN),
			},
		},
	}

	if len(gw.Addresses) > 0 {
		output.Properties.CustomRoutes = &armnetwork.AddressSpace{}
	}
	for i := range gw.Addresses {
		output.Properties.CustomRoutes.AddressPrefixes = append(
			output.Properties.CustomRoutes.AddressPrefixes,
			helper.StringToStringPointer(gw.Addresses[i]),
		)
	}

	// TODO: Add Public IPs

	for i := range gw.BGPAddresses {
		addresses := make([]*string, len(gw.BGPAddresses[i].CustomAddresses))
		for i, customAddr := range gw.BGPAddresses[i].CustomAddresses {
			addresses[i] = helper.StringToStringPointer(customAddr)
		}

		output.Properties.BgpSettings.BgpPeeringAddresses = append(
			output.Properties.BgpSettings.BgpPeeringAddresses,
			&armnetwork.IPConfigurationBgpPeeringAddress{
				CustomBgpIPAddresses: addresses,
			},
		)
	}
	return output
}

// Local Network Gateway serves the role of marking
// a single point of the connection with the other side.
//
// It provides information about IP Address of the public
// interface from the second side and its BGP Information
// such as ASN number and BGP IP Address.
type LocalVNetGateway struct {
	ID         string
	Name       string
	GatewayIP  string
	FQDN       string
	ReadyToUse bool
	// A list of Network Address Prefixes that can be reached
	// through that Gateway.
	//
	// This field should be set if there is no Azure Route
	// Server which would automatically collect all Routes
	// from the underlying network and all peered networks.
	NetworkAddresses []string
	// If set to false it means that the Local VNet Gateway
	// provisioning failed, is being deleted or it is in the update
	// process and it should not be modified at that particular
	// moment.
	ASN              string
	PeerBGPAddresses []BGPAddress
}

func localVnetGatewayFromAzure(gw armnetwork.LocalNetworkGateway) LocalVNetGateway {
	output := LocalVNetGateway{
		ID:   helper.StringPointerToString(gw.ID),
		Name: helper.StringPointerToString(gw.Name),
	}

	if gw.Properties == nil {
		return output
	}

	properties := gw.Properties

	output.GatewayIP = helper.StringPointerToString(properties.GatewayIPAddress)
	output.FQDN = helper.StringPointerToString(properties.Fqdn)

	if properties.ProvisioningState != nil {
		// Consider checking Active flag.
		output.ReadyToUse =
			*properties.ProvisioningState == armnetwork.ProvisioningStateSucceeded
	}

	if properties.LocalNetworkAddressSpace != nil {
		addressPrefixes := properties.LocalNetworkAddressSpace.AddressPrefixes
		for i := range addressPrefixes {
			output.NetworkAddresses = append(
				output.NetworkAddresses,
				helper.StringPointerToString(addressPrefixes[i]))
		}
	}

	if properties.BgpSettings == nil {
		return output
	}
	bgpSetting := properties.BgpSettings

	output.ASN = helper.Int64PointerToString(bgpSetting.Asn)

	for i := range bgpSetting.BgpPeeringAddresses {
		if bgpSetting.BgpPeeringAddresses[i] == nil {
			continue
		}
		output.PeerBGPAddresses = append(
			output.PeerBGPAddresses,
			bgpAddressFromAzure(*bgpSetting.BgpPeeringAddresses[i]),
		)
	}

	return output
}

func localVnetGatewayToAzure(gw LocalVNetGateway) armnetwork.LocalNetworkGateway {
	output := armnetwork.LocalNetworkGateway{
		ID:   helper.StringToStringPointer(gw.ID),
		Name: helper.StringToStringPointer(gw.Name),
		Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
			GatewayIPAddress: helper.StringToStringPointer(gw.GatewayIP),
			Fqdn:             helper.StringToStringPointer(gw.FQDN),
			BgpSettings: &armnetwork.BgpSettings{
				Asn: helper.StringToInt64Pointer(gw.ASN),
			},
		},
	}
	if len(gw.NetworkAddresses) > 0 {
		output.Properties.LocalNetworkAddressSpace = &armnetwork.AddressSpace{}
	}
	for _, addr := range gw.NetworkAddresses {
		output.Properties.LocalNetworkAddressSpace.AddressPrefixes = append(
			output.Properties.LocalNetworkAddressSpace.AddressPrefixes,
			helper.StringToStringPointer(addr),
		)
	}

	for _, addr := range gw.PeerBGPAddresses {
		addresses := make([]*string, len(addr.CustomAddresses))
		for i, customAddr := range addr.CustomAddresses {
			addresses[i] = helper.StringToStringPointer(customAddr)
		}

		output.Properties.BgpSettings.BgpPeeringAddress = helper.StringToStringPointer(strings.Join(
			addr.CustomAddresses, ",",
		))
	}
	return output
}

type NetworkGatewayConnection struct {
	ID                    string
	Name                  string
	SharedKey             string
	BGPEnabled            *bool
	NetworkGatewayID      string
	LocalNetworkGatewayID string
	ConnectionType        string
	ConnectionProtocol    string
}

func networkGatewayFromAzure(conn armnetwork.VirtualNetworkGatewayConnection) NetworkGatewayConnection {
	output := NetworkGatewayConnection{
		ID:   helper.StringPointerToString(conn.ID),
		Name: helper.StringPointerToString(conn.Name),
	}

	if conn.Properties == nil {
		return output
	}
	properties := conn.Properties

	output.BGPEnabled = properties.EnableBgp
	output.SharedKey = helper.StringPointerToString(properties.SharedKey)

	if properties.ConnectionType != nil {
		output.ConnectionType = string(*properties.ConnectionType)
	}
	if properties.ConnectionProtocol != nil {
		output.ConnectionProtocol = string(*properties.ConnectionProtocol)
	}

	if properties.VirtualNetworkGateway1 != nil {
		output.NetworkGatewayID = helper.StringPointerToString(
			properties.VirtualNetworkGateway1.ID,
		)
	}

	if properties.LocalNetworkGateway2 != nil {
		output.LocalNetworkGatewayID = helper.StringPointerToString(
			properties.LocalNetworkGateway2.ID)
	}

	return output
}

func networkGatewayToAzure(conn NetworkGatewayConnection) armnetwork.VirtualNetworkGatewayConnection {
	output := armnetwork.VirtualNetworkGatewayConnection{
		ID:   helper.StringToStringPointer(conn.ID),
		Name: helper.StringToStringPointer(conn.Name),
		Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
			SharedKey: helper.StringToStringPointer(conn.SharedKey),
			EnableBgp: conn.BGPEnabled,
			VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
				ID: helper.StringToStringPointer(conn.NetworkGatewayID),
			},
			LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
				ID: helper.StringToStringPointer(conn.LocalNetworkGatewayID),
			},
		},
	}

	if conn.ConnectionType != "" {
		connType := armnetwork.VirtualNetworkGatewayConnectionType(conn.ConnectionType)
		output.Properties.ConnectionType = &connType
	}
	if conn.ConnectionProtocol != "" {
		connProtocol := armnetwork.VirtualNetworkGatewayConnectionProtocol(conn.ConnectionProtocol)
		output.Properties.ConnectionProtocol = &connProtocol
	}

	return output
}

type PublicIP struct {
	ID      string
	Name    string
	Address string
}

func publicIPFromAzure(ip armnetwork.PublicIPAddress) PublicIP {
	output := PublicIP{
		ID:   helper.StringPointerToString(ip.ID),
		Name: helper.StringPointerToString(ip.Name),
	}
	if ip.Properties != nil {
		output.Address = helper.StringPointerToString(ip.Properties.IPAddress)
	}
	return output
}

func IDToName(id string) string {
	chunks := strings.Split(id, "/")
	// TODO: Check if that's possible.
	if len(chunks) == 0 {
		return ""
	}
	return chunks[len(chunks)-1]
}
