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

package gcp

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/iterator"
)

func (c *Client) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {
	c.logger.Debugf("Syncing GCP VPCs")

	if params == nil {
		params = &infrapb.ListVPCRequest{}
	}
	if params.Region != "" {
		return nil, fmt.Errorf("VPCs in GCP have global scope")
	}
	if len(params.Labels) > 0 {
		return nil, fmt.Errorf("VPCs in GCP don't have labels")
	}
	vpcs := make([]types.VPC, 0)
	f := func(projectID string, vpcs []types.VPC) ([]types.VPC, error) {
		iter := c.networksClient.List(ctx, &computepb.ListNetworksRequest{
			Project: projectID,
		})
		for {
			net, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			vpcs = append(vpcs, convertVPC(projectID, net))
		}
		return vpcs, nil
	}
	if params.AccountId == "" {
		for projectID := range c.projectIDs {
			newVPCs, err := f(projectID, vpcs)
			if err != nil {
				return nil, err
			}
			vpcs = append(vpcs, newVPCs...)
		}
	} else {
		return f(params.AccountId, vpcs)
	}

	return vpcs, nil
}

func (c *Client) GetVPCIDForCIDR(ctx context.Context, request *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	subnets, err := c.ListSubnets(ctx, &infrapb.ListSubnetsRequest{
		Cidr:      request.GetCidr(),
		Region:    request.GetRegion(),
		AccountId: request.GetAccountId(),
	})
	if err != nil {
		return "", err
	}
	if len(subnets) > 1 {
		return "", fmt.Errorf("expected to find one subnet with cidr %s, found %v", request.GetCidr(), len(subnets))
	}
	return subnets[0].VpcId, nil
}

func (c *Client) GetVPCIDWithTag(_ context.Context, _ *infrapb.GetVPCIDWithTagRequest) (string, error) {
	return "", fmt.Errorf("VPC tags or labels are not supported in GCP")
}

type network struct {
	// project name
	project string
	name    string
	// url in form projects/<project_name>/global/networks/<network_name>
	url string
	// url in form https://www.googleapis.com/compute/v1/projects/<project_name>/global/networks/<network_name>
	fullUrl string
	id      string
}

// findNetwork looks for network (VPC), based on provided input.
// Supported VPC input format:
// 1. URL of the network resource for this firewall rule with project name information. For example:
//   - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network
//   - projects/myproject/global/networks/my-network -
//
// 2. Name or ID, for example:
//   - my-network
//   - 235083625034176684.
//     In this case given network will be looked up in specified 'project' or
//     in all projects specified in client.projectIDs if 'project' param is empty and all matching
//     will be returned.
func (c *Client) findNetwork(ctx context.Context, project, vpc string) ([]network, error) {
	var networksWithProject []network
	fullUrlStarter := "https://www.googleapis.com/compute/v1/"

	// try to determine what kind of ID this is
	if strings.Contains(vpc, "projects/") {
		// it's a URL
		path := strings.Split(vpc, "/")
		for i, v := range path {
			if v == "projects" && i != len(path)-1 {
				url := strings.Join(path[i:], "/")
				foundProject := path[i+1]
				if project != "" && foundProject != project {
					return nil, fmt.Errorf("project in vpc path: %s doesn't match project in function paramter: %s", foundProject, project)
				}
				net, err := c.networksClient.Get(ctx, &computepb.GetNetworkRequest{
					Network: path[len(path)-1],
					Project: foundProject,
				})
				if err != nil {
					return nil, err
				}
				networksWithProject = []network{
					{
						name:    net.GetName(),
						url:     url,
						project: strconv.FormatUint(net.GetId(), 10),
						fullUrl: fullUrlStarter + url,
					},
				}
				break
			}
		}
	} else {
		// it's a name or ID, we'll look for it in all configured projects
		found := false
		findNet := func(projectName string) error {
			iter := c.networksClient.List(ctx, &computepb.ListNetworksRequest{
				Project: projectName,
			})
			for {
				net, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return err
				}
				id := strconv.FormatUint(net.GetId(), 10)
				if net.GetName() == vpc || id == vpc {
					found = true
					url := fmt.Sprintf("projects/%s/global/networks/%s", projectName, net.GetName())
					networksWithProject = append(networksWithProject, network{
						name:    net.GetName(),
						url:     url,
						project: projectName,
						fullUrl: fullUrlStarter + url,
						id:      id,
					})
				}
			}
			return nil
		}
		if project != "" {
			err := c.checkProject(project)
			if err != nil {
				return nil, err
			}
			err = findNet(project)
			if err != nil {
				return nil, err
			}
		} else {
			for project := range c.projectIDs {
				err := findNet(project)
				if err != nil {
					return nil, err
				}
			}
		}
		if !found {
			return nil, fmt.Errorf("couldn't find network %s in any configured project", vpc)
		}
	}
	return networksWithProject, nil
}

func (c *Client) vpcIdToSingleNetwork(ctx context.Context, project, vpcID string) (network, error) {
	networks, err := c.findNetwork(ctx, project, vpcID)
	if err != nil {
		return network{}, err
	}
	if len(networks) > 1 {
		return network{}, fmt.Errorf("found more than one VPC (network) matching to provided ID: %s", vpcID)
	}
	if len(networks) == 0 {
		return network{}, fmt.Errorf("couldn't find network matching to provided ID: %s", vpcID)
	}
	return networks[0], nil
}

func convertVPC(projectID string, network *computepb.Network) types.VPC {
	var ipv6Range string
	if network.InternalIpv6Range != nil {
		ipv6Range = *network.InternalIpv6Range
	}
	consoleAccesLInk := fmt.Sprintf("https://console.cloud.google.com/networking/networks/details/%s?project=%s", network.GetName(), projectID)
	return types.VPC{
		ID:     strconv.FormatUint(network.GetId(), 10),
		Name:   network.GetName(),
		Region: "global",

		AccountID: projectID,
		Provider:  providerName,
		SelfLink:  consoleAccesLInk,
		IPv4CIDR:  network.GetIPv4Range(),
		IPv6CIDR:  ipv6Range,
	}
}

func isIPv4CIDR(cidr string) bool {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return false // Not a valid CIDR notation
	}
	return ip.To4() != nil // Returns true if CIDR is IPv4
}
