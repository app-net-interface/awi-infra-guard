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

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListInstances(ctx context.Context, input *infrapb.ListInstancesRequest) ([]types.Instance, error) {

	var instances []types.Instance
	var vNetName string
	if input.VpcId != "" {
		parts := strings.Split(input.VpcId, "/")
		if len(parts) > 0 {
			vNetName = parts[len(parts)-1]
		}
	}
	c.logger.Infof("Retrieving instances for account %s and VPC %s", input.AccountId, input.VpcId)
	if input.AccountId == "" {
		return nil, fmt.Errorf("account ID is required")
	}

	vmClient, err := armcompute.NewVirtualMachinesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM client: %w", err)
	}
	nicClient, err := armnetwork.NewInterfacesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network interfaces client: %w", err)
	}
	publicIPClient, err := armnetwork.NewPublicIPAddressesClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create public IP addresses client: %w", err)
	}

	// Example logic for listing all VMs (simplified for demonstration)
	vmPager := vmClient.NewListAllPager(nil)
	for vmPager.More() {
		vmResult, err := vmPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of VMs: %w", err)
		}
		for _, vm := range vmResult.Value {
			if vm.Properties.NetworkProfile == nil {
				continue
			}
			var nsgId []string
			var interfaceIds []string
			for _, nicRef := range vm.Properties.NetworkProfile.NetworkInterfaces {
				if nicRef.ID != nil {
					interfaceIds = append(interfaceIds, *nicRef.ID)
				}

				nic, err := nicClient.Get(ctx, getResourceGroupName(*nicRef.ID), getResourceName(*nicRef.ID), nil)
				//nic.Interface.Properties.NetworkSecurityGroup.
				if err != nil {
					fmt.Printf("Failed to get NIC: %v\n", err)
					continue
				}
				if nic.Properties == nil || nic.Properties.IPConfigurations == nil {
					continue
				}
				if nic.Properties.NetworkSecurityGroup == nil {
					c.logger.Debugf("No security group associated with the NIC")
				} else {
					nsgId = append(nsgId, *nic.Properties.NetworkSecurityGroup.ID)
					c.logger.Debugf("Network Security Group ID: %s", nsgId)
				}
				for _, ipConf := range nic.Properties.IPConfigurations {
					if ipConf.Properties == nil || ipConf.Properties.Subnet == nil || !strings.Contains(*ipConf.Properties.Subnet.ID, vNetName) {
						continue
					}
					// Extract VNet ID from the subnet ID
					subnetID := *ipConf.Properties.Subnet.ID
					vNetID := extractVNetIDFromSubnetID(subnetID)

					privateIP := ""
					if ipConf.Properties.PrivateIPAddress != nil {
						privateIP = *ipConf.Properties.PrivateIPAddress
					}
					publicIP := ""
					if ipConf.Properties.PublicIPAddress != nil {
						pip, err := publicIPClient.Get(ctx, getResourceGroupName(*ipConf.Properties.PublicIPAddress.ID), getResourceName(*ipConf.Properties.PublicIPAddress.ID), nil)
						if err == nil && pip.Properties != nil && pip.Properties.IPAddress != nil {
							publicIP = *pip.Properties.IPAddress
						}
					}
					var vmStatus string = "Unknown"
					if vm.Properties.InstanceView != nil {
						for _, status := range vm.Properties.InstanceView.Statuses {
							if status.Code != nil && *status.Code == "PowerState/running" {
								vmStatus = "Running"
							} else if status.Code != nil && *status.Code == "PowerState/deallocated" {
								vmStatus = "Terminated"
							} else if status.Code != nil && *status.Code == "PowerState/stopped" {
								vmStatus = "Stopped"
							} else if status.Code != nil && *status.Code == "PowerState/starting" {
								vmStatus = "Starting"
							} else if status.Code != nil && *status.Code == "PowerState/stopping" {
								vmStatus = "Stopping"
							}
						}
					}
					// nicID := *vm.Properties.NetworkProfile.NetworkInterfaces[0].ID
					// nicName := getResourceName(nicID)

					// // Create a new network client
					// resourceGroupName := getResourceGroupName(*nicRef.ID)
					// networkClient, err := armnetwork.NewInterfacesClient(input.AccountId, c.cred, nil)
					// if err != nil {
					// 	c.logger.Errorf("failed to create network client: %v", err)
					// }

					// // Get the network interface
					// nic, err := networkClient.Get(ctx, resourceGroupName, nicName, nil)
					// if err != nil {
					// 	c.logger.Errorf("failed to get network interface: %v", err)
					// }

					// // Get the network security group ID
					// if nic.Properties.NetworkSecurityGroup == nil {
					// 	c.logger.Debugf("No security group associated with the NIC")
					// }
					// nsgID := *nic.Properties.NetworkSecurityGroup.ID
					// c.logger.Debugf("Network Security Group ID: %s", nsgID)

					// Construct and append the instance
					instance := types.Instance{
						ID:        *vm.ID,
						Name:      *vm.Name,
						Type:      string(*vm.Properties.HardwareProfile.VMSize),
						PublicIP:  publicIP,
						PrivateIP: privateIP,
						SubnetID:  *ipConf.Properties.Subnet.ID,
						VPCID:     vNetID,
						Labels:    convertToStringMap(vm.Tags),
						//State:            string(*vm.Properties.ProvisioningState),
						State:            vmStatus,
						Region:           *vm.Location,
						Provider:         "Azure",
						AccountID:        input.AccountId,
						SecurityGroupIDs: nsgId,
						InterfaceIDs:     interfaceIds,

						//SelfLink: *vm.Properties.,
						// LastSyncTime and Zone fields require additional logic or assumptions
					}

					instances = append(instances, instance)
					c.logger.Debugf("Azure Instance = %v", instance)
					break // Assuming a single NIC per VM for simplicity
				}
			}
		}
	}
	return instances, nil
}
