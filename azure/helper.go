package azure

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// VNetSubnetAssociation holds the association of a route table with its VNet and subnets
type VNetSubnetAssociation struct {
	VNetID    string   // ID of the VNet
	SubnetIDs []string // IDs of the subnets associated with the route table
}

type VNetInstanceAssociation struct {
	VNetID      string   // ID of the VNet
	InstanceIds []string // IDs of the subnets associated with the route table
}

type VNetAssociations struct {
	RtAssociations  map[string]VNetSubnetAssociation
	NsgAssociations map[string]VNetSubnetAssociation
}

func ListVNetSubnetAssociations(ctx context.Context, subscriptionID string, cred *azidentity.DefaultAzureCredential) (*VNetAssociations, error) {
	// Initialize a map to hold route table ID to VNetSubnetAssociation

	va := &VNetAssociations{
		RtAssociations:  make(map[string]VNetSubnetAssociation),
		NsgAssociations: make(map[string]VNetSubnetAssociation),
	}

	vnetsClient, err := armnetwork.NewVirtualNetworksClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create virtual networks client: %w", err)
	}

	// List all VNets in the subscription
	vnetPager := vnetsClient.NewListAllPager(nil)
	for vnetPager.More() {
		vnetResult, err := vnetPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get the next page of VNets: %w", err)
		}

		for _, vnet := range vnetResult.Value {
			// For each VNet, list all subnets
			subnetsClient, err := armnetwork.NewSubnetsClient(subscriptionID, cred, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create subnets client: %w", err)
			}

			// Extract the resource group name from the VNet ID
			rgName := ExtractResourceGroupName(*vnet.ID)
			subnetPager := subnetsClient.NewListPager(rgName, *vnet.Name, nil)
			for subnetPager.More() {
				subnetResult, err := subnetPager.NextPage(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed to list subnets for VNet '%s': %w", *vnet.Name, err)
				}

				for _, subnet := range subnetResult.Value {
					// Check if the subnet has an associated route table
					if subnet.Properties != nil && subnet.Properties.RouteTable != nil {
						rtID := *subnet.Properties.RouteTable.ID

						// Update the association map
						if assoc, ok := va.RtAssociations[rtID]; ok {
							assoc.SubnetIDs = append(assoc.SubnetIDs, *subnet.ID)
							va.RtAssociations[rtID] = assoc // Ensure the updated slice is put back
						} else {
							va.RtAssociations[rtID] = VNetSubnetAssociation{
								VNetID:    *vnet.ID,
								SubnetIDs: []string{*subnet.ID},
							}
						}
					}
					if subnet.Properties != nil && subnet.Properties.NetworkSecurityGroup != nil {
						nsgID := *subnet.Properties.NetworkSecurityGroup.ID

						// Update the association map
						if assoc, ok := va.NsgAssociations[nsgID]; ok {
							assoc.SubnetIDs = append(assoc.SubnetIDs, *subnet.ID)
							va.NsgAssociations[nsgID] = assoc // Ensure the updated slice is put back
						} else {
							va.NsgAssociations[nsgID] = VNetSubnetAssociation{
								VNetID:    *vnet.ID,
								SubnetIDs: []string{*subnet.ID},
							}
						}
					}
				}
			}
		}
	}
	return va, nil
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

// Helper functions to parse IDs and convert tags
func parseResourceGroupName(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if part == "resourceGroups" {
			return parts[i+1]
		}
	}
	return ""
}

func parseResourceName(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	return parts[len(parts)-1]
}

func convertToStringMap(tags map[string]*string) map[string]string {
	result := make(map[string]string)
	for key, value := range tags {
		if value != nil {
			result[key] = *value
		}
	}
	return result
}
func extractVNetIDFromSubnetID(subnetID string) string {
	parts := strings.Split(subnetID, "/")
	for i, part := range parts {
		if part == "virtualNetworks" && i < len(parts)-1 {
			return strings.Join(parts[:i+2], "/")
		}
	}
	return ""
}

func convertToSecurityGroupRule(rules []*armnetwork.SecurityRule) (sgRules []types.SecurityGroupRule) {
	sgRules = make([]types.SecurityGroupRule, 0, len(rules))
	for _, rule := range rules {
		if rule == nil {
			continue // Skip nil rules
		}
		sgRule := types.SecurityGroupRule{
			Protocol:  string(*rule.Properties.Protocol),
			PortRange: *rule.Properties.DestinationPortRange,
			Direction: string(*rule.Properties.Direction),
		}
		for _, sourceAddressPrefix := range rule.Properties.SourceAddressPrefixes {
			sgRule.Source = append(sgRule.Source, *sourceAddressPrefix)
		}
		sgRules = append(sgRules, sgRule)
	}
	return sgRules
}
