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
	"context"
	"errors"
	"fmt"
	"slices"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/sirupsen/logrus"
	computeServ "google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
)

type Client struct {
	instancesClient           *compute.InstancesClient
	subnetsClient             *compute.SubnetworksClient
	networksClient            *compute.NetworksClient
	routersClient             *compute.RoutersClient
	vpnGatewaysClient         *compute.VpnGatewaysClient
	vpnTunnelsClient          *compute.VpnTunnelsClient
	externalVPNGatewaysClient *compute.ExternalVpnGatewaysClient
	computeService            *computeServ.Service
	logger                    *logrus.Entry
	activeProject             string
	closeTasks                []func() error
}

// NewClient instantiates a new GCP Client instance and sets
// an active project.
// If the project string is empty, the GCP Client will check how many projects are available,
// if there is only one, it will be picked. Otherwise, an error will be returned that the
// client does not know which project should be used.
func NewClient(ctx context.Context, logger *logrus.Entry, project string) (*Client, error) {
	project, err := chooseProject(ctx, logger, project)
	if err != nil {
		return nil, err
	}

	closeTasks := []func() error{}

	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, instancesClient.Close)

	subnetsClient, err := compute.NewSubnetworksRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, subnetsClient.Close)

	networksClient, err := compute.NewNetworksRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, networksClient.Close)

	routersClient, err := compute.NewRoutersRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, routersClient.Close)

	vpnGatewayClient, err := compute.NewVpnGatewaysRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, vpnGatewayClient.Close)

	vpnTunnelClient, err := compute.NewVpnTunnelsRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, vpnTunnelClient.Close)

	externalVPNGatewaysClient, err := compute.NewExternalVpnGatewaysRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	closeTasks = append(closeTasks, externalVPNGatewaysClient.Close)

	computeService, err := computeServ.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		instancesClient:           instancesClient,
		subnetsClient:             subnetsClient,
		networksClient:            networksClient,
		routersClient:             routersClient,
		vpnGatewaysClient:         vpnGatewayClient,
		vpnTunnelsClient:          vpnTunnelClient,
		externalVPNGatewaysClient: externalVPNGatewaysClient,
		computeService:            computeService,
		logger:                    logger,
		activeProject:             project,
		closeTasks:                closeTasks,
	}, nil
}

func chooseProject(ctx context.Context, logger *logrus.Entry, project string) (string, error) {
	availableProjects, err := findProjects(ctx, logger)
	if err != nil {
		return "", fmt.Errorf(
			"failed to choose GCP Project due to error with listing available GCP Projects: %w", err,
		)
	}
	if len(availableProjects) == 0 {
		return "", errors.New("the user is not authenticated to any GCP Project. Did you set proper credentials?")
	}
	if project == "" && len(availableProjects) > 1 {
		return "", fmt.Errorf(
			"cannot choose a GCP Project to choose as multiple are available."+
				" Please specify one of the followings %v", availableProjects,
		)
	}
	if project == "" {
		return availableProjects[0], nil
	}
	if !slices.Contains(availableProjects, project) {
		return "", fmt.Errorf(
			"the requested project ID '%s' was not found among available projects: %v.",
			project, availableProjects,
		)
	}
	return project, nil
}

func findProjects(ctx context.Context, logger *logrus.Entry) ([]string, error) {
	projectsClient, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return nil, err
	}

	var projectIDs []string
	iter := projectsClient.SearchProjects(ctx, &resourcemanagerpb.SearchProjectsRequest{})
	for {
		project, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		logger.Infof("Found GCP project %s (ID %s)", project.GetDisplayName(), project.GetProjectId())
		projectIDs = append(projectIDs, project.GetProjectId())
	}

	return projectIDs, nil
}

func (c *Client) GetVPNGateway(ctx context.Context, region, gatewayID string) (*VPNGateway, error) {
	gatewayID = nameFromURL(gatewayID)
	gateway, err := c.vpnGatewaysClient.Get(ctx, &computepb.GetVpnGatewayRequest{
		Project:    c.activeProject,
		Region:     region,
		VpnGateway: gatewayID,
	})
	if err != nil {
		return nil, err
	}
	return vpnGatewayFromGCP(gateway), nil
}

// Close closes all internal clients used
// by this client.
//
// Call it as soon as your main Client is not
// in use anymore.
func (c *Client) Close() error {
	for _, closeTask := range c.closeTasks {
		if err := closeTask(); err != nil {
			return fmt.Errorf("failed to close the client: %w", err)
		}
	}
	c.logger.Debug("The client has been closed successfully")
	return nil
}

func (c *Client) FindVPNGatewaysForNetwork(ctx context.Context, region, networkID string) ([]*VPNGateway, error) {
	networkID = nameFromURL(networkID)
	filteredGateways := []*VPNGateway{}

	iter := c.vpnGatewaysClient.List(ctx, &computepb.ListVpnGatewaysRequest{
		Project: c.activeProject,
		Region:  region,
	})
	if iter == nil {
		return nil, errors.New("Could not obtain iterator for VPN Gateways. Do you have proper access?")
	}

	for gateway, err := iter.Next(); err != iterator.Done; gateway, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP VPN Gateways due to: %w", err)
		}
		if gateway == nil {
			return nil, errors.New("Unexpected empty GCP VPN Gateway object")
		}
		if gateway.Network != nil && nameFromURL(*gateway.Network) == networkID {
			filteredGateways = append(filteredGateways, vpnGatewayFromGCP(gateway))
		}
	}
	return filteredGateways, nil
}

func (c *Client) ListVPNGateways(ctx context.Context, region string) ([]*VPNGateway, error) {
	iter := c.vpnGatewaysClient.List(ctx, &computepb.ListVpnGatewaysRequest{
		Project: c.activeProject,
		Region:  region,
	})
	if iter == nil {
		return nil, errors.New("Could not obtain iterator for VPN Gateways. Do you have proper access?")
	}

	gateways := []*VPNGateway{}
	for gateway, err := iter.Next(); err != iterator.Done; gateway, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP VPN Gateways due to: %w", err)
		}
		if gateway == nil {
			return nil, errors.New("Unexpected empty GCP VPN Gateway object")
		}
		gateways = append(gateways, vpnGatewayFromGCP(gateway))
	}
	return gateways, nil
}

func (c *Client) ListVPNTunnels(ctx context.Context, region string) ([]*VPNTunnel, error) {
	iter := c.vpnTunnelsClient.List(ctx, &computepb.ListVpnTunnelsRequest{
		Project: c.activeProject,
		Region:  region,
	})
	if iter == nil {
		return nil, errors.New("Could not obtain iterator for VPN Tunnels. Do you have proper access?")
	}

	tunnels := []*VPNTunnel{}
	for tunnel, err := iter.Next(); err != iterator.Done; tunnel, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP VPN Tunnel due to: %w", err)
		}
		if tunnel == nil {
			return nil, errors.New("Unexpected empty GCP VPN Tunnel object")
		}
		tunnels = append(tunnels, vpnTunnelFromGCP(tunnel))
	}
	return tunnels, nil
}

func (c *Client) GetVPNTunnel(ctx context.Context, region, vpnTunnelID string) (*VPNTunnel, error) {
	vpnTunnelID = nameFromURL(vpnTunnelID)
	tunnel, err := c.vpnTunnelsClient.Get(ctx, &computepb.GetVpnTunnelRequest{
		VpnTunnel: vpnTunnelID,
		Project:   c.activeProject,
		Region:    region,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get external vpn gateway '%s' due to: %w", vpnTunnelID, err)
	}
	return vpnTunnelFromGCP(tunnel), nil
}

func (c *Client) CreateVPNTunnel(
	ctx context.Context, tunnel *VPNTunnel, region string,
) (*VPNTunnel, error) {
	if tunnel == nil {
		return nil, errors.New(
			"cannot create VPN Tunnel as the requested object is empty")
	}
	op, err := c.vpnTunnelsClient.Insert(ctx, &computepb.InsertVpnTunnelRequest{
		VpnTunnelResource: vpnTunnelToGCP(tunnel),
		Project:           c.activeProject,
		Region:            region,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create an VPN Tunnel '%v' due to %w",
			*tunnel, err)
	}
	if op == nil {
		return nil, fmt.Errorf(
			"creating VPN Tunnel '%v' did not return any operation object",
			*tunnel)
	}
	if err = op.Wait(ctx); err != nil {
		return nil, fmt.Errorf(
			"failed while waiting for creation of an VPN Tunnel '%v' due to %w. Operation ID: %v",
			*tunnel, err, op.Name())
	}
	createdTunnel, err := c.GetVPNTunnel(ctx, region, tunnel.Name)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain freshly created VPN Tunnel '%s': %w", tunnel.Name, err)
	}
	return createdTunnel, nil
}

func (c *Client) DeleteVPNTunnel(
	ctx context.Context, region, name string,
) error {
	op, err := c.vpnTunnelsClient.Delete(ctx, &computepb.DeleteVpnTunnelRequest{
		VpnTunnel: name,
		Project:   c.activeProject,
		Region:    region,
	})
	if err != nil {
		return fmt.Errorf(
			"failed to delete an VPN Tunnel '%s' due to %w",
			name, err)
	}
	if op == nil {
		return fmt.Errorf(
			"deleting VPN Tunnel '%s' did not return any operation object",
			name)
	}
	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf(
			"failed while waiting for deletion of an VPN Tunnel '%s' due to %w. Operation ID: %v",
			name, err, op.Name())
	}
	return nil
}

func (c *Client) ListNetworks(ctx context.Context) ([]*Network, error) {
	iter := c.networksClient.List(ctx, &computepb.ListNetworksRequest{
		Project: c.activeProject,
	})
	if iter == nil {
		return nil, errors.New("Could not obtain iterator for GCP Networks. Do you have proper access?")
	}

	networks := []*Network{}
	for network, err := iter.Next(); err != iterator.Done; network, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP Networks due to: %w", err)
		}
		if network == nil {
			return nil, errors.New("Unexpected empty GCP Network object")
		}
		networks = append(networks, networkFromGCP(network))
	}
	return networks, nil
}

func (c *Client) GetNetwork(ctx context.Context, networkID string) (*Network, error) {
	networkID = nameFromURL(networkID)
	network, err := c.networksClient.Get(ctx, &computepb.GetNetworkRequest{
		Project: c.activeProject,
		Network: networkID,
	})
	if err != nil {
		return nil, err
	}
	c.logger.Debugf("Found network: %s", network.String())
	return networkFromGCP(network), nil
}

func (c *Client) GetSubnetwork(ctx context.Context, region, subnetworkID string) (*Subnetwork, error) {
	subnetworkID = nameFromURL(subnetworkID)
	subnetwork, err := c.subnetsClient.Get(ctx, &computepb.GetSubnetworkRequest{
		Project:    c.activeProject,
		Region:     region,
		Subnetwork: subnetworkID,
	})
	if err != nil {
		return nil, err
	}
	c.logger.Debugf("Found subnetwork: %s", subnetwork.String())
	return subnetworkFromGCP(subnetwork), nil
}

func (c *Client) GetRouter(ctx context.Context, region, routerID string) (*CloudRouter, error) {
	routerID = nameFromURL(routerID)
	router, err := c.routersClient.Get(ctx, &computepb.GetRouterRequest{
		Project: c.activeProject,
		Region:  region,
		Router:  routerID,
	})
	if err != nil {
		return nil, err
	}
	c.logger.Debugf("Found Cloud Router: %s", router.String())
	return cloudRouterFromGCP(router), nil
}

func (c *Client) AddRouterInterface(
	ctx context.Context, region, routerID string, iface CloudRouterInterface,
) (*CloudRouter, error) {
	router, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get Cloud Router with ID '%s': %w", routerID, err)
	}
	if router == nil {
		return nil, fmt.Errorf("got empty Cloud Router object for ID: %s", routerID)
	}
	router.Interfaces = append(router.Interfaces, iface)
	op, err := c.routersClient.Patch(ctx, &computepb.PatchRouterRequest{
		Project:        c.activeProject,
		Region:         region,
		Router:         routerID,
		RouterResource: cloudRouterToGCP(router),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to add Cloud Router Interface '%v' due to %w",
			iface, err)
	}
	if op == nil {
		return nil, fmt.Errorf(
			"adding Cloud Router Interface '%v' did not return any operation object",
			iface)
	}
	if err = op.Wait(ctx); err != nil {
		return nil, fmt.Errorf(
			"failed while waiting for creation of a Cloud Router Interface '%v' due to %w. Operation ID: %v",
			iface, err, op.Name())
	}
	updatedRouter, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain freshly updated Cloud Router '%s': %w", routerID, err)
	}
	return updatedRouter, nil
}

func (c *Client) DeleteRouterInterface(
	ctx context.Context, region, routerID, name string,
) error {
	router, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return fmt.Errorf("failed to get Cloud Router with ID '%s': %w", routerID, err)
	}
	if router == nil {
		return fmt.Errorf("got empty Cloud Router object for ID: %s", routerID)
	}
	updatedInterfaces := []CloudRouterInterface{}
	for _, iface := range router.Interfaces {
		if iface.Name != name {
			updatedInterfaces = append(updatedInterfaces, iface)
		}
		c.logger.Debugf("excluding Interface '%v' from router ID '%s'", iface, routerID)
	}
	router.Interfaces = updatedInterfaces
	op, err := c.routersClient.Update(ctx, &computepb.UpdateRouterRequest{
		Project:        c.activeProject,
		Region:         region,
		Router:         routerID,
		RouterResource: cloudRouterToGCP(router),
	})
	if err != nil {
		return fmt.Errorf(
			"failed to remove Cloud Router Interface '%s' due to %w",
			name, err)
	}
	if op == nil {
		return fmt.Errorf(
			"removing Cloud Router Interface '%s' did not return any operation object",
			name)
	}
	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf(
			"failed while waiting for deletion of a Cloud Router Interface '%s' due to %w. Operation ID: %v",
			name, err, op.Name())
	}
	return nil
}

func (c *Client) AddRouterBGPPeer(
	ctx context.Context, region, routerID string, bgpPeer CloudRouterBGPPeer,
) (*CloudRouter, error) {
	router, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get Cloud Router with ID '%s': %w", routerID, err)
	}
	if router == nil {
		return nil, fmt.Errorf("got empty Cloud Router object for ID: %s", routerID)
	}
	router.BGPPeers = append(router.BGPPeers, bgpPeer)
	op, err := c.routersClient.Patch(ctx, &computepb.PatchRouterRequest{
		Project:        c.activeProject,
		Region:         region,
		Router:         routerID,
		RouterResource: cloudRouterToGCP(router),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to add Cloud Router BGP Peer '%v' due to %w",
			bgpPeer, err)
	}
	if op == nil {
		return nil, fmt.Errorf(
			"adding Cloud Router BGP Peer '%v' did not return any operation object",
			bgpPeer)
	}
	if err = op.Wait(ctx); err != nil {
		return nil, fmt.Errorf(
			"failed while waiting for creation of a Cloud Router BGP Peer '%v' due to %w. Operation ID: %v",
			bgpPeer, err, op.Name())
	}
	updatedRouter, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain freshly updated Cloud Router '%s': %w", routerID, err)
	}
	return updatedRouter, nil
}

func (c *Client) DeleteRouterBGPPeer(
	ctx context.Context, region, routerID, name string,
) error {
	router, err := c.GetRouter(ctx, region, routerID)
	if err != nil {
		return fmt.Errorf("failed to get Cloud Router with ID '%s': %w", routerID, err)
	}
	if router == nil {
		return fmt.Errorf("got empty Cloud Router object for ID: %s", routerID)
	}
	updatedPeers := []CloudRouterBGPPeer{}
	for _, peer := range router.BGPPeers {
		if peer.Name != name {
			updatedPeers = append(updatedPeers, peer)
		}
		c.logger.Debugf("excluding BGP Peer '%v' from router ID '%s'", peer, routerID)
	}
	router.BGPPeers = updatedPeers
	op, err := c.routersClient.Update(ctx, &computepb.UpdateRouterRequest{
		Project:        c.activeProject,
		Region:         region,
		Router:         routerID,
		RouterResource: cloudRouterToGCP(router),
	})
	if err != nil {
		return fmt.Errorf(
			"failed to remove Cloud Router BGP Peer '%s' due to %w",
			name, err)
	}
	if op == nil {
		return fmt.Errorf(
			"removing Cloud Router BGP Peer '%s' did not return any operation object",
			name)
	}
	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf(
			"failed while waiting for deletion of a Cloud Router BGP Peer '%s' due to %w. Operation ID: %v",
			name, err, op.Name())
	}
	return nil
}

func (c *Client) ListRouters(ctx context.Context, region string) ([]*CloudRouter, error) {
	iter := c.routersClient.List(ctx, &computepb.ListRoutersRequest{
		Project: c.activeProject,
		Region:  region,
	})
	if iter == nil {
		return nil, errors.New("Could not obtain iterator for GCP Routers. Do you have proper access?")
	}

	routers := []*CloudRouter{}
	for router, err := iter.Next(); err != iterator.Done; router, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP Networks due to: %w", err)
		}
		if router == nil {
			return nil, errors.New("Unexpected empty GCP Network object")
		}
		routers = append(routers, cloudRouterFromGCP(router))
	}
	return routers, nil
}

func (c *Client) ListExternalVPNGateway(ctx context.Context) ([]*ExternalVPNGateway, error) {
	iter := c.externalVPNGatewaysClient.List(ctx, &computepb.ListExternalVpnGatewaysRequest{
		Project: c.activeProject,
	})
	if iter == nil {
		return nil, errors.New(
			"Could not obtain iterator for External VPN Gateway. Do you have proper access?")
	}

	gateways := []*ExternalVPNGateway{}
	for gateway, err := iter.Next(); err != iterator.Done; gateway, err = iter.Next() {
		if err != nil {
			return nil, fmt.Errorf("Could not list GCP External VPN Gateways due to: %w", err)
		}
		if gateway == nil {
			return nil, errors.New("Unexpected empty External VPN Gateway object")
		}
		gateways = append(gateways, externalVPNGatewayFromGCP(gateway))
	}
	return gateways, nil
}

func (c *Client) GetExternalVPNGateway(ctx context.Context, externalVPNGatewayID string) (*ExternalVPNGateway, error) {
	externalVPNGatewayID = nameFromURL(externalVPNGatewayID)
	gateway, err := c.externalVPNGatewaysClient.Get(ctx, &computepb.GetExternalVpnGatewayRequest{
		ExternalVpnGateway: externalVPNGatewayID,
		Project:            c.activeProject,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get external vpn gateway '%s' due to: %w", externalVPNGatewayID, err)
	}
	return externalVPNGatewayFromGCP(gateway), nil
}

func (c *Client) CreateExternalVPNGateway(
	ctx context.Context, gw *ExternalVPNGateway,
) (*ExternalVPNGateway, error) {
	if gw == nil {
		return nil, errors.New(
			"cannot create External VPN Gateway as the requested object is empty")
	}
	op, err := c.externalVPNGatewaysClient.Insert(ctx, &computepb.InsertExternalVpnGatewayRequest{
		ExternalVpnGatewayResource: externalVPNGatewayToGCP(gw),
		Project:                    c.activeProject,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create an External VPN Gateway '%v' due to %w",
			*gw, err)
	}
	if op == nil {
		return nil, fmt.Errorf(
			"creating External VPN Gateway '%v' did not return any operation object",
			*gw)
	}
	if err = op.Wait(ctx); err != nil {
		return nil, fmt.Errorf(
			"failed while waiting for creation of an External VPN Gateway '%v' due to %w. Operation ID: %v",
			*gw, err, op.Name())
	}
	createdGateway, err := c.GetExternalVPNGateway(ctx, gw.Name)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain freshly created External VPN Gateway '%s': %w", gw.Name, err)
	}
	return createdGateway, nil
}

func (c *Client) DeleteExternalVPNGateway(
	ctx context.Context, name string,
) error {
	op, err := c.externalVPNGatewaysClient.Delete(ctx, &computepb.DeleteExternalVpnGatewayRequest{
		ExternalVpnGateway: name,
		Project:            c.activeProject,
	})
	if err != nil {
		return fmt.Errorf(
			"failed to delete an External VPN Gateway '%s' due to %w",
			name, err)
	}
	if op == nil {
		return fmt.Errorf(
			"deleting External VPN Gateway '%s' did not return any operation object",
			name)
	}
	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf(
			"failed while waiting for deletion of an External VPN Gateway '%s' due to %w. Operation ID: %v",
			name, err, op.Name())
	}
	return nil
}

func (c *Client) GetName() string {
	return "GCP"
}
