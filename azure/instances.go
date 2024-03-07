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
	c.logger.Infof("Retrieving instances for account %s and VPC %s", input.AccountId, vNetName)

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
			for _, nicRef := range vm.Properties.NetworkProfile.NetworkInterfaces {
				nic, err := nicClient.Get(ctx, parseResourceGroupName(*nicRef.ID), parseResourceName(*nicRef.ID), nil)
				//nic.Interface.Properties.NetworkSecurityGroup.
				if err != nil {
					fmt.Printf("Failed to get NIC: %v\n", err)
					continue
				}
				if nic.Properties == nil || nic.Properties.IPConfigurations == nil {
					continue
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
						pip, err := publicIPClient.Get(ctx, parseResourceGroupName(*ipConf.Properties.PublicIPAddress.ID), parseResourceName(*ipConf.Properties.PublicIPAddress.ID), nil)
						if err == nil && pip.Properties != nil && pip.Properties.IPAddress != nil {
							publicIP = *pip.Properties.IPAddress
						}
					}

					// Construct and append the instance
					instance := types.Instance{
						ID:        *vm.ID,
						Name:      *vm.Name,
						PublicIP:  publicIP,
						PrivateIP: privateIP,
						SubnetID:  *ipConf.Properties.Subnet.ID,
						VPCID:     vNetID,
						Labels:    convertToStringMap(vm.Tags),
						//State:     string(*vm.Properties.InstanceView.Statuses[0].Code),
						Region:    *vm.Location,
						Provider:  "Azure",
						AccountID: input.AccountId,
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

