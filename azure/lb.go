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

func (c *Client) ListLBs(ctx context.Context, input *infrapb.ListLBsRequest) ([]types.LB, error) {
	c.logger.Infof("Listing Load Balancers for account %s", input.AccountId)
	var lbs []types.LB

	client, err := armnetwork.NewLoadBalancersClient(input.AccountId, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create load balancers client: %w", err)
	}

	pager := client.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get next page of load balancers: %w", err)
		}

		for _, lb := range page.Value {
			if lb.ID == nil || lb.Name == nil || lb.Location == nil || lb.Properties == nil {
				continue
			}

			var vpcId string
			var subnetIds []string
			scheme := "internal" // Default to internal
			var publicIPs []string
			var privateIPs []string

			if lb.Properties.FrontendIPConfigurations != nil {
				for _, frontend := range lb.Properties.FrontendIPConfigurations {
					if frontend.Properties == nil {
						continue
					}
					if frontend.Properties.PublicIPAddress != nil && frontend.Properties.PublicIPAddress.ID != nil {
						scheme = "internet-facing"
						publicIPs = append(publicIPs, *frontend.Properties.PublicIPAddress.ID)
					}
					if frontend.Properties.Subnet != nil && frontend.Properties.Subnet.ID != nil {
						subnetId := *frontend.Properties.Subnet.ID
						subnetIds = append(subnetIds, subnetId)
						if vpcId == "" {
							parts := strings.Split(subnetId, "/subnets/")
							if len(parts) > 0 {
								vpcId = parts[0]
							}
						}
						if frontend.Properties.PrivateIPAddress != nil {
							privateIPs = append(privateIPs, *frontend.Properties.PrivateIPAddress)
						}
					}
				}
			}

			var listeners []types.LBListener
			if lb.Properties.LoadBalancingRules != nil {
				for _, rule := range lb.Properties.LoadBalancingRules {
					if rule.Name == nil || rule.Properties == nil || rule.Properties.Protocol == nil || rule.Properties.FrontendPort == nil {
						continue
					}
					listener := types.LBListener{
						Protocol: string(*rule.Properties.Protocol),
						Port:     *rule.Properties.FrontendPort,
					}
					if rule.Properties.BackendAddressPool != nil && rule.Properties.BackendAddressPool.ID != nil {
						listener.TargetGroupID = *rule.Properties.BackendAddressPool.ID
					}
					listeners = append(listeners, listener)
				}
			}

			var targetGroupIds []string
			var instanceIDs []string
			if lb.Properties.BackendAddressPools != nil {
				for _, pool := range lb.Properties.BackendAddressPools {
					if pool.ID != nil {
						targetGroupIds = append(targetGroupIds, *pool.ID)
					}
				}
				// Resolve backend pools to instance IDs
				resolvedInstanceIDs, err := c.getInstanceIDsForBackendPools(ctx, input.AccountId, lb.Properties.BackendAddressPools)
				if err != nil {
					c.logger.Warnf("Failed to get instance IDs for LB %s: %v", *lb.ID, err)
				}
				instanceIDs = resolvedInstanceIDs
			}

			labels := make(map[string]string)
			if lb.Tags != nil {
				for k, v := range lb.Tags {
					if v != nil {
						labels[k] = *v
					}
				}
			}

			lbType := ""
			if lb.SKU != nil && lb.SKU.Name != nil {
				lbType = string(*lb.SKU.Name)
			}

			provisioningState := ""
			if lb.Properties.ProvisioningState != nil {
				provisioningState = string(*lb.Properties.ProvisioningState)
			}

			newLB := types.LB{
				ID:             *lb.ID,
				Name:           *lb.Name,
				Region:         *lb.Location,
				VPCID:          vpcId,
				SubnetIDs:      subnetIds,
				State:          provisioningState,
				Type:           lbType,
				Scheme:         scheme,
				PublicIPs:      publicIPs,
				PrivateIPs:     privateIPs,
				Listeners:      listeners,
				TargetGroupIDs: targetGroupIds,
				InstanceIDs:    instanceIDs,
				Provider:       c.GetName(),
				AccountID:      input.AccountId,
				Labels:         labels,
			}
			lbs = append(lbs, newLB)
		}
	}

	c.logger.Infof("Successfully listed %d Load Balancers for account %s", len(lbs), input.AccountId)
	return lbs, nil
}

// getInstanceIDsForBackendPools resolves the instance IDs associated with a load balancer's backend pools.
func (c *Client) getInstanceIDsForBackendPools(ctx context.Context, accountID string, pools []*armnetwork.BackendAddressPool) ([]string, error) {
	nicIDs := make(map[string]struct{})
	vmssIDs := make(map[string]struct{})

	for _, pool := range pools {
		if pool.Properties == nil {
			continue
		}
		// Handle older, instance-based backend pools via BackendIPConfigurations
		if pool.Properties.BackendIPConfigurations != nil {
			for _, ipConfig := range pool.Properties.BackendIPConfigurations {
				if ipConfig.ID != nil {
					if nicID := getNicIDFromIPConfigID(*ipConfig.ID); nicID != "" {
						nicIDs[nicID] = struct{}{}
					}
				}
			}
		}
		// Handle newer, IP-based backend pools via LoadBalancerBackendAddresses
		if pool.Properties.LoadBalancerBackendAddresses != nil {
			for _, addr := range pool.Properties.LoadBalancerBackendAddresses {
				if addr.Properties != nil && addr.Properties.NetworkInterfaceIPConfiguration != nil && addr.Properties.NetworkInterfaceIPConfiguration.ID != nil {
					ipConfigID := *addr.Properties.NetworkInterfaceIPConfiguration.ID
					// The IP config can belong to a VMSS instance or a standalone VM's NIC.
					if vmssID := getVmssIDFromIPConfigID(ipConfigID); vmssID != "" {
						vmssIDs[vmssID] = struct{}{}
					} else if nicID := getNicIDFromIPConfigID(ipConfigID); nicID != "" {
						nicIDs[nicID] = struct{}{}
					}
				}
			}
		}
	}

	var instanceIDs []string

	// Resolve NICs to VM IDs
	if len(nicIDs) > 0 {
		resolvedIDs, err := c.resolveNICsToVMIDs(ctx, accountID, nicIDs)
		if err != nil {
			c.logger.Warnf("Failed to resolve some NICs to VM IDs: %v", err)
		}
		instanceIDs = append(instanceIDs, resolvedIDs...)
	}

	// Resolve VMSS to VM IDs
	if len(vmssIDs) > 0 {
		resolvedIDs, err := c.resolveVMSSsToVMIDs(ctx, accountID, vmssIDs)
		if err != nil {
			c.logger.Warnf("Failed to resolve some VMSSs to VM IDs: %v", err)
		}
		instanceIDs = append(instanceIDs, resolvedIDs...)
	}

	// Return unique instance IDs
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

func (c *Client) resolveNICsToVMIDs(ctx context.Context, accountID string, nicIDs map[string]struct{}) ([]string, error) {
	nicClient, err := armnetwork.NewInterfacesClient(accountID, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create network interfaces client: %w", err)
	}

	var instanceIDs []string
	for nicID := range nicIDs {
		parts := strings.Split(nicID, "/")
		if len(parts) < 9 || !strings.EqualFold(parts[3], "resourceGroups") || !strings.EqualFold(parts[7], "networkInterfaces") {
			continue
		}
		rgName := parts[4]
		nicName := parts[8]

		nic, err := nicClient.Get(ctx, rgName, nicName, nil)
		if err != nil {
			c.logger.Warnf("Failed to get network interface %s: %v", nicID, err)
			continue
		}

		if nic.Properties != nil && nic.Properties.VirtualMachine != nil && nic.Properties.VirtualMachine.ID != nil {
			instanceIDs = append(instanceIDs, *nic.Properties.VirtualMachine.ID)
		}
	}
	return instanceIDs, nil
}

func (c *Client) resolveVMSSsToVMIDs(ctx context.Context, accountID string, vmssIDs map[string]struct{}) ([]string, error) {
	vmssClient, err := armcompute.NewVirtualMachineScaleSetVMsClient(accountID, c.cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create VMSS client: %w", err)
	}

	var instanceIDs []string
	for vmssID := range vmssIDs {
		parts := strings.Split(vmssID, "/")
		if len(parts) < 9 || !strings.EqualFold(parts[3], "resourceGroups") || !strings.EqualFold(parts[7], "virtualMachineScaleSets") {
			continue
		}
		rgName := parts[4]
		vmssName := parts[8]

		pager := vmssClient.NewListPager(rgName, vmssName, nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				c.logger.Warnf("Failed to list VMs for VMSS %s: %v", vmssID, err)
				break
			}
			for _, vm := range page.Value {
				if vm.ID != nil {
					instanceIDs = append(instanceIDs, *vm.ID)
				}
			}
		}
	}
	return instanceIDs, nil
}

func getNicIDFromIPConfigID(ipConfigID string) string {
	parts := strings.Split(ipConfigID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "networkInterfaces") && i+1 < len(parts) {
			return strings.Join(parts[:i+2], "/")
		}
	}
	return ""
}

func getVmssIDFromIPConfigID(ipConfigID string) string {
	parts := strings.Split(ipConfigID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "virtualMachineScaleSets") && i+1 < len(parts) {
			return strings.Join(parts[:i+2], "/")
		}
	}
	return ""
}
