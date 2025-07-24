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
	"fmt"
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

func (c *Client) ListLBs(ctx context.Context, input *infrapb.ListLBsRequest) ([]types.LB, error) {
	c.logger.Infof("Listing Load Balancers for project %s", input.AccountId)

	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create compute service: %w", err)
	}

	var lbs []types.LB
	req := computeService.ForwardingRules.AggregatedList(input.AccountId)
	if err := req.Pages(ctx, func(page *compute.ForwardingRuleAggregatedList) error {
		for region, list := range page.Items {
			if input.Region != "" && input.Region != "all" && !strings.HasSuffix(region, "/"+input.Region) {
				continue
			}
			for _, fr := range list.ForwardingRules {
				// We are interested in load balancers, not other forwarding rule types like PSC
				if strings.Contains(fr.Target, "serviceAttachments") {
					continue
				}

				vpcID := fr.Network
				if input.VpcId != "" && vpcID != input.VpcId {
					continue
				}

				lb, err := c.convertGCPForwardingRuleToLB(ctx, computeService, input.AccountId, fr)
				if err != nil {
					c.logger.Warnf("could not convert forwarding rule %s to LB: %v", fr.Name, err)
					continue
				}
				lbs = append(lbs, lb)
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to list forwarding rules: %w", err)
	}

	c.logger.Infof("Successfully listed %d Load Balancers for project %s", len(lbs), input.AccountId)
	return lbs, nil
}

func (c *Client) convertGCPForwardingRuleToLB(ctx context.Context, client *compute.Service, projectID string, fr *compute.ForwardingRule) (types.LB, error) {
	var publicIPs, privateIPs []string
	if fr.IPAddress != "" {
		// GCP doesn't have a simple is_private field. We assume if it's not a global IP, it's private.
		// A more robust check would involve parsing the IP. For now, we classify based on scheme.
		if fr.LoadBalancingScheme == "EXTERNAL" || fr.LoadBalancingScheme == "EXTERNAL_MANAGED" {
			publicIPs = append(publicIPs, fr.IPAddress)
		} else {
			privateIPs = append(privateIPs, fr.IPAddress)
		}
	}

	var listeners []types.LBListener
	if len(fr.Ports) > 0 {
		for _, portStr := range fr.Ports {
			port, err := strconv.ParseInt(portStr, 10, 32)
			if err != nil {
				c.logger.Warnf("could not parse port '%s' for forwarding rule %s", portStr, fr.Name)
				continue
			}
			listeners = append(listeners, types.LBListener{
				Protocol: fr.IPProtocol,
				Port:     int32(port),
			})
		}
	} else if fr.PortRange != "" {
		// Handle port ranges if necessary, for now, we just record the protocol
		listeners = append(listeners, types.LBListener{
			Protocol: fr.IPProtocol,
		})
	}

	instanceIDs, err := c.getInstanceIDsForGCPBackend(ctx, client, projectID, fr.Target)
	if err != nil {
		c.logger.Warnf("Failed to get instance IDs for backend %s: %v", fr.Target, err)
	}

	return types.LB{
		ID:        fmt.Sprintf("%d", fr.Id),
		Name:      fr.Name,
		Provider:  c.GetName(),
		AccountID: projectID,
		Region:    extractResourceID(fr.Region),
		VPCID:     extractResourceID(fr.Network),
		SubnetIDs: []string{extractResourceID(fr.Subnetwork)},
		//State:         fr.Status,
		Type:          getGCPlbType(fr),
		Scheme:        fr.LoadBalancingScheme,
		PublicIPs:     publicIPs,
		PrivateIPs:    privateIPs,
		Listeners:     listeners,
		InstanceIDs:   instanceIDs,
		Labels:        fr.Labels,
		SelfLink:      fr.SelfLink,
		IPAddressType: fr.IpVersion,
		DNSName:       fr.Name, // GCP LBs don't have a DNS name in the same way as AWS/Azure
	}, nil
}

func (c *Client) getInstanceIDsForGCPBackend(ctx context.Context, client *compute.Service, projectID string, targetURL string) ([]string, error) {
	if targetURL == "" {
		return nil, nil
	}

	backendServiceURL, err := c.getBackendServiceURL(ctx, client, projectID, targetURL)
	if err != nil {
		return nil, err
	}
	if backendServiceURL == "" {
		// This could be a legacy network load balancer using a TargetPool
		if strings.Contains(targetURL, "targetPools") {
			return c.getInstancesFromTargetPool(ctx, client, projectID, targetURL)
		}
		c.logger.Debugf("Unhandled target type for LB instance resolution: %s", targetURL)
		return nil, nil
	}

	// Handle backend services (for modern LBs)
	isGlobal := strings.Contains(backendServiceURL, "/global/")
	var backendService *compute.BackendService
	if isGlobal {
		backendService, err = client.BackendServices.Get(projectID, extractResourceID(backendServiceURL)).Context(ctx).Do()
	} else {
		region := extractResourceID(extractParentURL(backendServiceURL))
		backendService, err = client.RegionBackendServices.Get(projectID, region, extractResourceID(backendServiceURL)).Context(ctx).Do()
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get backend service %s: %w", backendServiceURL, err)
	}

	var instanceIDs []string
	for _, backend := range backendService.Backends {
		instanceGroupURL := backend.Group
		ids, err := c.listInstancesInGroup(ctx, client, projectID, instanceGroupURL)
		if err != nil {
			c.logger.Warnf("Failed to list instances for group %s: %v", instanceGroupURL, err)
			continue
		}
		instanceIDs = append(instanceIDs, ids...)
	}

	// Deduplicate
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range instanceIDs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}

func (c *Client) getBackendServiceURL(ctx context.Context, client *compute.Service, projectID, targetURL string) (string, error) {
	targetName := extractResourceID(targetURL)
	isGlobal := !strings.Contains(targetURL, "/regions/")

	if strings.Contains(targetURL, "targetHttpProxies") {
		var proxy *compute.TargetHttpProxy
		var err error
		if isGlobal {
			proxy, err = client.TargetHttpProxies.Get(projectID, targetName).Context(ctx).Do()
		} else {
			region := extractResourceID(extractParentURL(targetURL))
			proxy, err = client.RegionTargetHttpProxies.Get(projectID, region, targetName).Context(ctx).Do()
		}
		if err != nil {
			return "", err
		}
		return proxy.UrlMap, nil
	}
	if strings.Contains(targetURL, "targetHttpsProxies") {
		var proxy *compute.TargetHttpsProxy
		var err error
		if isGlobal {
			proxy, err = client.TargetHttpsProxies.Get(projectID, targetName).Context(ctx).Do()
		} else {
			region := extractResourceID(extractParentURL(targetURL))
			proxy, err = client.RegionTargetHttpsProxies.Get(projectID, region, targetName).Context(ctx).Do()
		}
		if err != nil {
			return "", err
		}
		return proxy.UrlMap, nil
	}
	if strings.Contains(targetURL, "backendServices") {
		return targetURL, nil
	}
	// TargetPools are handled separately, so we don't resolve a backend service URL for them here.
	if !strings.Contains(targetURL, "targetPools") {
		c.logger.Debugf("Unhandled target type for LB instance resolution: %s", targetURL)
	}
	return "", nil
}

func getGCPlbType(fr *compute.ForwardingRule) string {
	// Infer type based on protocol, similar to how we classify AWS/Azure LBs.
	// HTTP(S) LBs are ALBs. TCP/UDP LBs are NLBs.
	if strings.Contains(fr.Target, "targetHttpProxies") || strings.Contains(fr.Target, "targetHttpsProxies") {
		return "ALB"
	}
	switch fr.IPProtocol {
	case "TCP", "UDP", "SCTP", "ESP", "ICMP", "AH", "L3_DEFAULT":
		return "NLB"
	default:
		// Fallback to the scheme name if it's not a clear ALB/NLB case.
		return fr.LoadBalancingScheme
	}
}

func (c *Client) getInstancesFromTargetPool(ctx context.Context, client *compute.Service, projectID, targetPoolURL string) ([]string, error) {
	poolName := extractResourceID(targetPoolURL)
	poolRegion := extractResourceID(extractParentURL(targetPoolURL))

	pool, err := client.TargetPools.Get(projectID, poolRegion, poolName).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get target pool %s: %w", targetPoolURL, err)
	}
	// The 'Instances' field in a TargetPool response contains the full URLs of the instances.
	var instanceIDs []string
	for _, instanceURL := range pool.Instances {
		instanceIDs = append(instanceIDs, extractResourceID(instanceURL))
	}
	return instanceIDs, nil
}

func (c *Client) listInstancesInGroup(ctx context.Context, client *compute.Service, projectID, instanceGroupURL string) ([]string, error) {
	var instanceIDs []string
	groupName := extractResourceID(instanceGroupURL)
	isZonal := strings.Contains(instanceGroupURL, "/zones/")

	if isZonal {
		groupZone := extractResourceID(extractParentURL(instanceGroupURL))
		req := client.InstanceGroups.ListInstances(projectID, groupZone, groupName, &compute.InstanceGroupsListInstancesRequest{})
		if err := req.Pages(ctx, func(page *compute.InstanceGroupsListInstances) error {
			for _, item := range page.Items {
				instanceIDs = append(instanceIDs, extractResourceID(item.Instance))
			}
			return nil
		}); err != nil {
			return nil, fmt.Errorf("failed to list instances in zonal group %s: %w", instanceGroupURL, err)
		}
	} else { // Regional Instance Group
		groupRegion := extractResourceID(extractParentURL(instanceGroupURL))
		req := client.RegionInstanceGroups.ListInstances(projectID, groupRegion, groupName, &compute.RegionInstanceGroupsListInstancesRequest{})
		if err := req.Pages(ctx, func(page *compute.RegionInstanceGroupsListInstances) error {
			for _, item := range page.Items {
				instanceIDs = append(instanceIDs, extractResourceID(item.Instance))
			}
			return nil
		}); err != nil {
			return nil, fmt.Errorf("failed to list instances in regional group %s: %w", instanceGroupURL, err)
		}
	}

	return instanceIDs, nil
}

func extractParentURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 2 {
		return strings.Join(parts[:len(parts)-2], "/")
	}
	return ""
}
