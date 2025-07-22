// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
)

// ListVPCEndpoints lists all GCP Private Service Connect endpoints (forwarding rules).
func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {
	c.logger.Infof("Listing Private Service Connect Endpoints for project %s", params.AccountId)

	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create compute service: %w", err)
	}

	var endpoints []types.VPCEndpoint

	aggListReq := computeService.ForwardingRules.AggregatedList(params.AccountId)
	if err := aggListReq.Pages(ctx, func(page *compute.ForwardingRuleAggregatedList) error {
		for region, list := range page.Items {
			// If a specific region is requested and it's not the current one, skip.
			if params.Region != "" && params.Region != "all" && !strings.HasSuffix(region, "/"+params.Region) {
				continue
			}
			for _, fr := range list.ForwardingRules {
				// Private Service Connect endpoints are forwarding rules with a target that is a service attachment.
				if fr.LoadBalancingScheme != "INTERNAL" || !strings.Contains(fr.Target, "serviceAttachments") {
					continue
				}

				// Filter by VPC ID
				vpcID := fr.Network
				if params.VpcId != "" && vpcID != params.VpcId {
					continue
				}

				endpoints = append(endpoints, types.VPCEndpoint{
					ID:          fmt.Sprintf("%d", fr.Id),
					Provider:    c.GetName(),
					AccountID:   params.AccountId,
					Name:        fr.Name,
					VPCId:       vpcID,
					Region:      extractResourceID(region),
					State:       fr.PscConnectionStatus,
					Labels:      fr.Labels,
					Type:        "Private Service Connect",
					ServiceName: fr.Target,
					SubnetIds:   []string{fr.Subnetwork},
					// SecurityGroupIDs are not directly associated with forwarding rules in GCP.
					// They are applied at the instance level via network tags.
					SecurityGroupIDs: []string{},
					SelfLink:         fr.SelfLink,
				})
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to list forwarding rules: %w", err)
	}

	c.logger.Infof("Successfully listed %d Private Service Connect Endpoints for project %s", len(endpoints), params.AccountId)
	return endpoints, nil
}
