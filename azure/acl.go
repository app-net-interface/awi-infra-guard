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

package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types" // Adjust the import path according to your project structure
)

func (c *Client) ListACLs(ctx context.Context, input *infrapb.ListACLsRequest) ([]types.ACL, error) {

	// List All ACLS (Irrespective of VPC, subnet,NIC attached)
	acls, err := c.ListAllACLs(ctx, input)
	// Step 2: List all VNets and their subnets, noting any route table associations.
	va, err := ListVNetSubnetAssociations(ctx, input.AccountId, c.cred)
	if err != nil {
		return nil, err
	}
	// Step 3: Compare both lists and update the RouteTables list with VPCId and subnet.
	for i, acl := range acls {
		if association, ok := va.NsgAssociations[acl.ID]; ok {
			acls[i].VpcID = association.VNetID // Update with VNet ID
			//routeTables[i].Subnets = association.SubnetIDs // Update with associated subnet IDs
		}
		// Note: Route tables without no subnet (VPC) association will simply not be updated.
	}
	return acls, nil
}

func (c *Client) ListAllACLs(ctx context.Context, input *infrapb.ListACLsRequest) ([]types.ACL, error) {

	var acls []types.ACL

	// Creating an instance of the NSG client
	nsgClient, err := armnetwork.NewSecurityGroupsClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network security group client: %w", err)
	}

	// List all NSGs in the subscription
	pager := nsgClient.NewListAllPager(nil)
	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of network security groups: %w", err)
		}

		for _, nsg := range result.Value {
			labels := make(map[string]string)
			if nsg.Tags != nil {
				for k, v := range nsg.Tags {
					labels[k] = *v
				}
			}

			// Mapping NSG details to types.ACL
			acl := types.ACL{
				Name:         *nsg.Name,
				ID:           *nsg.ID,
				Provider:     c.GetName(),
				VpcID:        "Not Attached", // VNet association would need additional logic
				Region:       *nsg.Location,
				Labels:       labels,
				AccountID:    input.AccountId,
				Rules:        []types.ACLRule{}, // Rules extraction would need additional logic
				LastSyncTime: "",                // Populate this with the current time or another relevant timestamp
			}
			acls = append(acls, acl)
		}
	}
	return acls, nil
}
