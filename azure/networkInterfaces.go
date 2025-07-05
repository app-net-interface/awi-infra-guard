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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListNetworkInterfaces(ctx context.Context, input *infrapb.ListNetworkInterfacesRequest) ([]types.NetworkInterface, error) {
	c.logger.Infof("Listing network interfaces for account %s", input.AccountId)
	var networkInterfaces []types.NetworkInterface

	client, err := armnetwork.NewInterfacesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network interfaces client: %w", err)
	}

	pager := client.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get next page of network interfaces: %w", err)
		}

		for _, nic := range page.Value {
			if nic.Properties == nil {
				continue
			}

			var securityGroupIds []string
			if nic.Properties.NetworkSecurityGroup != nil && nic.Properties.NetworkSecurityGroup.ID != nil {
				securityGroupIds = append(securityGroupIds, *nic.Properties.NetworkSecurityGroup.ID)
			}

			var instanceId string
			if nic.Properties.VirtualMachine != nil && nic.Properties.VirtualMachine.ID != nil {
				instanceId = *nic.Properties.VirtualMachine.ID
			}

			labels := make(map[string]string)
			if nic.Tags != nil {
				for k, v := range nic.Tags {
					if v != nil {
						labels[k] = *v
					}
				}
			}

			var privateIps []string
			var publicIp, subnetId, vpcId string
			for _, ipConfig := range nic.Properties.IPConfigurations {
				if ipConfig.Properties != nil {
					if ipConfig.Properties.PrivateIPAddress != nil {
						privateIps = append(privateIps, *ipConfig.Properties.PrivateIPAddress)
					}
					if subnetId == "" && ipConfig.Properties.Subnet != nil && ipConfig.Properties.Subnet.ID != nil {
						subnetId = *ipConfig.Properties.Subnet.ID
						// Extract VPC ID from Subnet ID
						parts := strings.Split(subnetId, "/subnets/")
						if len(parts) > 0 {
							vpcId = parts[0]
						}
					}
					if ipConfig.Properties.PublicIPAddress != nil && ipConfig.Properties.PublicIPAddress.Properties != nil && ipConfig.Properties.PublicIPAddress.Properties.IPAddress != nil {
						publicIp = *ipConfig.Properties.PublicIPAddress.Properties.IPAddress
					}
				}
			}

			ni := types.NetworkInterface{
				ID:               *nic.ID,
				Name:             *nic.Name,
				Region:           *nic.Location,
				VPCID:            vpcId,
				SubnetID:         subnetId,
				PrivateIPs:       privateIps, // Primary private IP
				PublicIP:         publicIp,   // This is the ID of the Public IP resource
				MacAddress:       *nic.Properties.MacAddress,
				Status:           string(*nic.Properties.ProvisioningState),
				InstanceID:       instanceId,
				SecurityGroupIDs: securityGroupIds,
				Provider:         c.GetName(),
				AccountID:        input.AccountId,
				Labels:           labels,
			}

			networkInterfaces = append(networkInterfaces, ni)
		}
	}

	c.logger.Infof("Successfully listed %d network interfaces for account %s", len(networkInterfaces), input.AccountId)
	return networkInterfaces, nil
}
