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
	"fmt"
	"strings"

	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

type CloudRouterInterface struct {
	Name      string
	VPNTunnel string
	IpRange   string
}

func (c CloudRouterInterface) String() string {
	return fmt.Sprintf(
		"CloudRouterInterface {Name: '%s', VPNTunnel: '%s', IPRange: '%s'}",
		c.Name, c.VPNTunnel, c.IpRange,
	)
}

type CloudRouterBGP struct {
	ASN              string
	AdvertiseMode    string
	AdvertisedGroups []string
}

func (c CloudRouterBGP) String() string {
	return fmt.Sprintf(
		"CloudRouterBGP {ASN: '%s', AdvertiseMode: '%s', AdvertisedGroups: %v}",
		c.ASN, c.AdvertiseMode, c.AdvertisedGroups,
	)
}

func cloudRouterBGPFromGCP(bgp *computepb.RouterBgp) *CloudRouterBGP {
	if bgp == nil {
		return nil
	}
	return &CloudRouterBGP{
		AdvertiseMode:    helper.StringPointerToString(bgp.AdvertiseMode),
		AdvertisedGroups: bgp.AdvertisedGroups,
		ASN:              helper.UInt32PointerToString(bgp.Asn),
	}
}

func cloudRouterBGPToGCP(bgp *CloudRouterBGP) *computepb.RouterBgp {
	if bgp == nil {
		return nil
	}
	return &computepb.RouterBgp{
		AdvertisedGroups: bgp.AdvertisedGroups,
		AdvertiseMode:    helper.StringToStringPointer(bgp.AdvertiseMode),
		Asn:              helper.StringToUInt32Pointer(bgp.ASN),
	}
}

type CloudRouterBGPPeer struct {
	Name          string
	ASN           string
	Interface     string
	PeerIPAddress string
}

func (c CloudRouterBGPPeer) String() string {
	return fmt.Sprintf(
		"CloudRouterBGPPeer {Name: '%s', ASN: '%s', Interface: '%s', PeerIPAddress: '%s'}",
		c.Name, c.ASN, c.Interface, c.PeerIPAddress,
	)
}

func cloudRouterBGPPeerFromGCP(bgpPeer *computepb.RouterBgpPeer) *CloudRouterBGPPeer {
	if bgpPeer == nil {
		return nil
	}
	return &CloudRouterBGPPeer{
		Name:          helper.StringPointerToString(bgpPeer.Name),
		ASN:           helper.UInt32PointerToString(bgpPeer.PeerAsn),
		Interface:     helper.StringPointerToString(bgpPeer.InterfaceName),
		PeerIPAddress: helper.StringPointerToString(bgpPeer.PeerIpAddress),
	}
}

func cloudRouterBGPPeerToGCP(bgpPeer *CloudRouterBGPPeer) *computepb.RouterBgpPeer {
	if bgpPeer == nil {
		return nil
	}
	return &computepb.RouterBgpPeer{
		Name:          helper.StringToStringPointer(bgpPeer.Name),
		PeerAsn:       helper.StringToUInt32Pointer(bgpPeer.ASN),
		InterfaceName: helper.StringToStringPointer(bgpPeer.Interface),
		PeerIpAddress: helper.StringToStringPointer(bgpPeer.PeerIPAddress),
	}
}

type CloudRouter struct {
	Name       string
	Network    string
	BGP        *CloudRouterBGP
	BGPPeers   []CloudRouterBGPPeer
	URL        string
	Interfaces []CloudRouterInterface
}

func (c CloudRouter) String() string {
	return fmt.Sprintf(
		"CloudRouter {Name: '%s', Network: '%s', BGP: %v, BGPPeers: %v, URL: '%s', Interfaces: %v}",
		c.Name, c.Network, c.BGP, c.BGPPeers, c.URL, c.Interfaces,
	)
}

func cloudRouterFromGCP(router *computepb.Router) *CloudRouter {
	if router == nil {
		return nil
	}
	cloudRouter := &CloudRouter{
		Name:       helper.StringPointerToString(router.Name),
		Network:    helper.StringPointerToString(router.Network),
		BGP:        cloudRouterBGPFromGCP(router.Bgp),
		URL:        helper.StringPointerToString(router.SelfLink),
		Interfaces: []CloudRouterInterface{},
		BGPPeers:   []CloudRouterBGPPeer{},
	}

	for i := range router.Interfaces {
		if router.Interfaces[i] == nil {
			continue
		}
		cloudRouter.Interfaces = append(cloudRouter.Interfaces, CloudRouterInterface{
			Name:      helper.StringPointerToString(router.Interfaces[i].Name),
			VPNTunnel: helper.StringPointerToString(router.Interfaces[i].LinkedVpnTunnel),
			IpRange:   helper.StringPointerToString(router.Interfaces[i].IpRange),
		})
	}

	for i := range router.BgpPeers {
		if router.BgpPeers[i] == nil {
			continue
		}
		bgpPeer := cloudRouterBGPPeerFromGCP(
			router.BgpPeers[i],
		)
		if bgpPeer == nil {
			continue
		}
		cloudRouter.BGPPeers = append(cloudRouter.BGPPeers, *bgpPeer)
	}

	return cloudRouter
}

func cloudRouterToGCP(router *CloudRouter) *computepb.Router {
	if router == nil {
		return nil
	}
	cloudRouter := &computepb.Router{
		Name:     &router.Name,
		Network:  &router.Network,
		SelfLink: &router.URL,
		Bgp:      cloudRouterBGPToGCP(router.BGP),
	}
	// Considering that GCP Patch system may treat nil fields
	// as these that should NOT be patched, it is important to not
	// send an empty slice instead of nil as it may clear all existing
	// interfaces.
	// If we want to do so, a special method should be created, which
	// will explicitely call patch with allocated empty slice of interfaces.
	//
	// Same goes for BGP Peers below.
	//
	// TODO: Confirm that the comment above is true - currently it is an
	// assumption.
	if len(router.Interfaces) > 0 {
		cloudRouter.Interfaces = []*computepb.RouterInterface{}
		for i := range router.Interfaces {
			cloudRouter.Interfaces = append(cloudRouter.Interfaces, &computepb.RouterInterface{
				Name:            &router.Interfaces[i].Name,
				IpRange:         &router.Interfaces[i].IpRange,
				LinkedVpnTunnel: &router.Interfaces[i].VPNTunnel,
			})
		}
	}
	if len(router.BGPPeers) > 0 {
		cloudRouter.BgpPeers = []*computepb.RouterBgpPeer{}
		for i := range router.BGPPeers {
			bgpPeer := cloudRouterBGPPeerToGCP(&router.BGPPeers[i])
			if bgpPeer != nil {
				cloudRouter.BgpPeers = append(cloudRouter.BgpPeers, bgpPeer)
			}
		}

	}
	return cloudRouter
}

type VPNGateway struct {
	Name        string
	URL         string
	Network     string
	IPAddresses []string
}

func (v VPNGateway) String() string {
	return fmt.Sprintf(
		"VPNGateway {Name: '%s', URL: '%s', Network: '%s', IPAddresses: %v}",
		v.Name, v.URL, v.Network, v.IPAddresses,
	)
}

func vpnGatewayFromGCP(gateway *computepb.VpnGateway) *VPNGateway {
	if gateway == nil {
		return nil
	}
	addresses := []string{}
	for _, i := range gateway.VpnInterfaces {
		if i == nil {
			continue
		}
		addresses = append(addresses, *i.IpAddress)
	}
	return &VPNGateway{
		Name:        helper.StringPointerToString(gateway.Name),
		Network:     helper.StringPointerToString(gateway.Network),
		URL:         helper.StringPointerToString(gateway.SelfLink),
		IPAddresses: addresses,
	}
}

type Network struct {
	Name    string
	Subnets []string
}

func (n Network) String() string {
	return fmt.Sprintf(
		"Network {Name: '%s', Subnets: '%v'}", n.Name, n.Subnets,
	)
}

func networkFromGCP(network *computepb.Network) *Network {
	if network == nil {
		return nil
	}
	return &Network{
		Name:    helper.StringPointerToString(network.Name),
		Subnets: network.Subnetworks,
	}
}

type Subnetwork struct {
	Name string
	CIDR string
}

func (s Subnetwork) String() string {
	return fmt.Sprintf(
		"Subnetwork {Name: '%s', CIDR: '%v'}", s.Name, s.CIDR,
	)
}

func subnetworkFromGCP(subnetwork *computepb.Subnetwork) *Subnetwork {
	if subnetwork == nil {
		return nil
	}
	return &Subnetwork{
		Name: helper.StringPointerToString(subnetwork.Name),
		CIDR: helper.StringPointerToString(subnetwork.IpCidrRange),
	}
}

type ExternalVPNGatewayInterface struct {
	IP string
}

func (e ExternalVPNGatewayInterface) String() string {
	return fmt.Sprintf(
		"ExternalVPNGatewayInterface {IP: '%s'}", e.IP,
	)
}

type ExternalVPNGateway struct {
	Name       string
	URL        string
	Interfaces []ExternalVPNGatewayInterface
}

func (e ExternalVPNGateway) String() string {
	return fmt.Sprintf(
		"ExternalVPNGateway {Name: '%s', URL: '%s', Interfaces: %v}",
		e.Name, e.URL, e.Interfaces,
	)
}

func externalVPNGatewayFromGCP(gateway *computepb.ExternalVpnGateway) *ExternalVPNGateway {
	if gateway == nil {
		return nil
	}
	gw := &ExternalVPNGateway{
		Name:       helper.StringPointerToString(gateway.Name),
		URL:        helper.StringPointerToString(gateway.SelfLink),
		Interfaces: []ExternalVPNGatewayInterface{},
	}
	for i := range gateway.Interfaces {
		if gateway.Interfaces[i] == nil {
			continue
		}
		gw.Interfaces = append(gw.Interfaces, ExternalVPNGatewayInterface{
			IP: gateway.Interfaces[i].GetIpAddress(),
		})
	}
	return gw
}

func externalVPNGatewayToGCP(gateway *ExternalVPNGateway) *computepb.ExternalVpnGateway {
	if gateway == nil {
		return nil
	}
	gcpGateway := &computepb.ExternalVpnGateway{
		RedundancyType: helper.StringToStringPointer("FOUR_IPS_REDUNDANCY"),
		Name:           &gateway.Name,
		SelfLink:       helper.StringToStringPointer(gateway.URL),
	}
	if len(gateway.Interfaces) > 0 {
		gcpGateway.Interfaces = []*computepb.ExternalVpnGatewayInterface{}
		for i := range gateway.Interfaces {
			id := uint32(i)
			gcpGateway.Interfaces = append(gcpGateway.Interfaces, &computepb.ExternalVpnGatewayInterface{
				IpAddress: &gateway.Interfaces[i].IP,
				Id:        &id,
			})
		}
	}
	return gcpGateway
}

type VPNTunnel struct {
	Name                     string
	URL                      string
	ExternalVPNGateway       string
	ExternalGatewayInterface int
	IKEVersion               int32
	SharedSecret             string
	CloudRouter              string
	VPNGateway               string
	Interface                int
}

func (v VPNTunnel) String() string {
	return fmt.Sprintf(
		"VPNTunnel {Name: '%s', URL: '%s', ExternalVPNGateway: '%s', ExternalGatewayInterface: %d "+
			"IKEVersion: %d, SharedSecret: ***, CloudRouter: '%s', VPNGateway: '%s', Interface: %d}",
		v.Name, v.URL, v.ExternalVPNGateway, v.ExternalGatewayInterface, v.IKEVersion,
		v.CloudRouter, v.VPNGateway, v.Interface,
	)
}

func vpnTunnelFromGCP(tunnel *computepb.VpnTunnel) *VPNTunnel {
	if tunnel == nil {
		return nil
	}
	vpnTunnel := &VPNTunnel{
		Name:               helper.StringPointerToString(tunnel.Name),
		ExternalVPNGateway: helper.StringPointerToString(tunnel.PeerExternalGateway),
		SharedSecret:       helper.StringPointerToString(tunnel.SharedSecret),
		CloudRouter:        helper.StringPointerToString(tunnel.Router),
		VPNGateway:         helper.StringPointerToString(tunnel.VpnGateway),
		URL:                helper.StringPointerToString(tunnel.SelfLink),
	}
	if tunnel.PeerExternalGatewayInterface != nil {
		vpnTunnel.ExternalGatewayInterface = int(*tunnel.PeerExternalGatewayInterface)
	}
	if tunnel.IkeVersion != nil {
		vpnTunnel.IKEVersion = *tunnel.IkeVersion
	}
	if tunnel.VpnGatewayInterface != nil {
		vpnTunnel.Interface = int(*tunnel.VpnGatewayInterface)
	}
	return vpnTunnel
}

func vpnTunnelToGCP(tunnel *VPNTunnel) *computepb.VpnTunnel {
	if tunnel == nil {
		return nil
	}
	externalGatewayInterface := int32(tunnel.ExternalGatewayInterface)
	vpnGatewayInterface := int32(tunnel.Interface)

	gcpTunnel := &computepb.VpnTunnel{
		Name:                         &tunnel.Name,
		IkeVersion:                   &tunnel.IKEVersion,
		PeerExternalGateway:          &tunnel.ExternalVPNGateway,
		PeerExternalGatewayInterface: &externalGatewayInterface,
		Router:                       &tunnel.CloudRouter,
		VpnGateway:                   &tunnel.VPNGateway,
		VpnGatewayInterface:          &vpnGatewayInterface,
		SharedSecret:                 &tunnel.SharedSecret,
	}
	if tunnel.URL != "" {
		gcpTunnel.SelfLink = &tunnel.URL
	}
	return gcpTunnel
}

// If the string has the form of GCP URL, the
// function will retrieve the resource name only.
// If the string seems to be rather name, it will
// return it.
func nameFromURL(url string) string {
	chunks := strings.Split(url, "/")
	return chunks[len(chunks)-1]
}
