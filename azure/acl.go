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
