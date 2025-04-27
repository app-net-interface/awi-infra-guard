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

package db

import (
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/types"
	//"github.com/boltdb/bolt"
)

// Updated UpdateVPCIndex to add debug logs at function entry and exit.
func (client *boltClient) UpdateVPCIndex(provider, vpcID, resourceId, resourceType string) error {
	// Add check for empty vpcID early
	if vpcID == "" {
		fmt.Printf("DEBUG: Skipping VPCIndex update for ResourceID: %s, ResourceType: %s because vpcID is empty\n", resourceId, resourceType)
		return nil // Not an error, just nothing to do
	}

	fmt.Printf("DEBUG: Updating VPCIndex for provider: %s VPCID: %s, ResourceID: %s, ResourceType: %s\n", provider, vpcID, resourceId, resourceType)
	key := provider + ":" + vpcID
	vpcIndex, err := client.GetVPCIndex(key)

	// Handle errors from GetVPCIndex first
	if err != nil {
		fmt.Printf("ERROR: Failed to get VPCIndex for key %s during update: %v\n", key, err)
		return err // Return the actual database error
	}

	// Handle case where index is not found (GetVPCIndex returned (nil, nil))
	if vpcIndex == nil {
		// This should ideally not happen often after SyncVPCIndexes Step 3 creates all base indexes.
		// It might happen if a resource references a VPC that wasn't discovered by ListVPCs.
		fmt.Printf("WARN: VPCIndex not found for key %s when trying to add resource %s (%s). Skipping update.\n", key, resourceId, resourceType)
		return nil // Return nil error to allow SyncVPCIndexes to continue
	}

	// Proceed with updating the found index
	switch resourceType {
	case "instance":
		if resourceId != "" && !contains(vpcIndex.InstanceIds, resourceId) {
			vpcIndex.InstanceIds = append(vpcIndex.InstanceIds, resourceId)
		}
	case "acl":
		if resourceId != "" && !contains(vpcIndex.AclIds, resourceId) {
			vpcIndex.AclIds = append(vpcIndex.AclIds, resourceId)
		}
	case "securityGroup":
		if resourceId != "" && !contains(vpcIndex.SecurityGroupIds, resourceId) {
			vpcIndex.SecurityGroupIds = append(vpcIndex.SecurityGroupIds, resourceId)
		}
	case "natgw":
		if resourceId != "" && !contains(vpcIndex.NatGatewayIds, resourceId) {
			vpcIndex.NatGatewayIds = append(vpcIndex.NatGatewayIds, resourceId)
		}
	case "subnet":
		if resourceId != "" && !contains(vpcIndex.SubnetIds, resourceId) {
			vpcIndex.SubnetIds = append(vpcIndex.SubnetIds, resourceId)
		}
	case "rt":
		if resourceId != "" && !contains(vpcIndex.RouteTableIds, resourceId) {
			vpcIndex.RouteTableIds = append(vpcIndex.RouteTableIds, resourceId)
		}
	case "publicIP":
		if resourceId != "" && !contains(vpcIndex.PublicIpIds, resourceId) {
			vpcIndex.PublicIpIds = append(vpcIndex.PublicIpIds, resourceId)
		}
	case "igw":
		if resourceId != "" && !contains(vpcIndex.IgwIds, resourceId) {
			vpcIndex.IgwIds = append(vpcIndex.IgwIds, resourceId)
		}
	case "vpce":
		if resourceId != "" && !contains(vpcIndex.VpcEndpointIds, resourceId) {
			vpcIndex.VpcEndpointIds = append(vpcIndex.VpcEndpointIds, resourceId)
		}
	case "lb":
		if resourceId != "" && !contains(vpcIndex.LbIds, resourceId) {
			vpcIndex.LbIds = append(vpcIndex.LbIds, resourceId)
		}
	case "vpnc":
		if resourceId != "" && !contains(vpcIndex.VpnConcentratorIds, resourceId) {
			vpcIndex.VpnConcentratorIds = append(vpcIndex.VpnConcentratorIds, resourceId)
		}
	case "netif":
		if resourceId != "" && !contains(vpcIndex.NetworkInterfaceIds, resourceId) {
			vpcIndex.NetworkInterfaceIds = append(vpcIndex.NetworkInterfaceIds, resourceId)
		}
	case "router":
		if resourceId != "" && !contains(vpcIndex.RouterIds, resourceId) {
			vpcIndex.RouterIds = append(vpcIndex.RouterIds, resourceId)
		}
	default:
		fmt.Printf("WARN: Unknown resource type: %s\n", resourceType)
	}
	//fmt.Printf("DEBUG: Updated VPCIndex: %+v for VPCID %s\n", vpcIndex, vpcID)
	return client.PutVPCIndex(vpcIndex) // Update the index in the DB
}

// SyncVPCIndexes synchronizes VPCIndexes after all resource sync.
func (client *boltClient) SyncVPCIndexes() error {
	fmt.Println("INFO: Starting VPC Index synchronization.")

	// 1. Fetch all VPCs and Routers
	vpcs, err := client.ListVPCs()
	if err != nil {
		fmt.Printf("ERROR: Failed to list VPCs for index sync: %v\n", err)
		return fmt.Errorf("failed to list VPCs for index sync: %w", err)
	}
	fmt.Printf("INFO: Fetched %d VPCs.\n", len(vpcs))

	allRouters, err := client.ListRouters()
	if err != nil {
		fmt.Printf("WARN: Failed to list Routers for index sync: %v\n", err)
		allRouters = []*types.Router{}
	}
	fmt.Printf("INFO: Fetched %d Routers.\n", len(allRouters))

	// 2. Group AWS Router IDs by Region (Key = lowercase region string)
	awsRoutersByRegion := make(map[string][]string) // Key: "us-west-2", Value: ["tgw-1", "tgw-2"]
	fmt.Printf("DEBUG: Processing %d routers for AWS region grouping.\n", len(allRouters))
	for _, router := range allRouters {
		fmt.Printf("DEBUG: Checking router: ID=%s, Provider=%s, Region=%s\n", router.ID, router.Provider, router.Region)
		// Only process AWS routers
		if strings.ToLower(router.Provider) == "aws" {
			if router.Region == "" {
				fmt.Printf("WARN: AWS Router %s missing Region, skipping for VPC index injection\n", router.ID)
				continue
			}
			// Use lowercase region as the key directly (NO "aws:" prefix)
			regionKey := strings.ToLower(router.Region)
			fmt.Printf("DEBUG: Adding AWS router %s to region key '%s'\n", router.ID, regionKey)
			awsRoutersByRegion[regionKey] = append(awsRoutersByRegion[regionKey], router.ID)
		}
	}
	fmt.Printf("INFO: Grouped AWS Routers by region. Map size: %d\n", len(awsRoutersByRegion))
	fmt.Printf("DEBUG: Content of awsRoutersByRegion map: %+v\n", awsRoutersByRegion)

	// 3. Recreate VPCIndexes from all VPCs, injecting regional AWS Router IDs conditionally
	processedVPCs := make(map[string]bool)
	fmt.Println("INFO: Starting VPC index creation/update loop.")
	for i, vpc := range vpcs {
		if vpc.Provider == "" || vpc.ID == "" {
			fmt.Printf("WARN: Skipping VPC at index %d due to missing Provider or ID. VPC Details: %+v\n", i, vpc)
			continue
		}
		vpcKey := vpc.Provider + ":" + vpc.ID
		if processedVPCs[vpcKey] {
			continue
		}
		processedVPCs[vpcKey] = true

		fmt.Printf("DEBUG: Processing VPC Index for key: %s (Region: %s)\n", vpcKey, vpc.Region)

		// Initialize the index
		index := &types.VPCIndex{
			VpcId:               vpc.ID,
			Provider:            vpc.Provider,
			AccountId:           vpc.AccountID,
			Region:              vpc.Region,
			RouterIds:           []string{}, // Start empty
			InstanceIds:         []string{},
			SubnetIds:           []string{},
			RouteTableIds:       []string{},
			NatGatewayIds:       []string{},
			IgwIds:              []string{},
			SecurityGroupIds:    []string{},
			AclIds:              []string{},
			LbIds:               []string{},
			VpcEndpointIds:      []string{},
			NetworkInterfaceIds: []string{},
			VpnConcentratorIds:  []string{},
			PublicIpIds:         []string{},
			ClusterIds:          []string{},
		}

		// Inject AWS regional routers specifically for AWS VPCs
		if strings.ToLower(vpc.Provider) == "aws" { // Check if it's an AWS VPC
			if vpc.Region != "" {
				// Use lowercase region for lookup (NO "aws:" prefix)
				regionKey := strings.ToLower(vpc.Region)
				fmt.Printf("DEBUG: VPC is AWS. Looking up region key '%s' in awsRoutersByRegion map\n", regionKey)
				regionalRouterIDs := awsRoutersByRegion[regionKey] // THE LOOKUP (using just region)
				if regionalRouterIDs != nil {
					fmt.Printf("DEBUG: Injecting %d AWS Router IDs into index for %s\n", len(regionalRouterIDs), vpcKey)
					index.RouterIds = regionalRouterIDs // THE ASSIGNMENT
				} else {
					fmt.Printf("DEBUG: No AWS routers found for region key '%s' (VPC %s)\n", regionKey, vpcKey)
					index.RouterIds = []string{} // Ensure empty slice
				}
			} else {
				fmt.Printf("WARN: AWS VPC %s missing Region, cannot inject regional TGWs\n", vpc.ID)
			}
		}

		// Save the initial index
		fmt.Printf("DEBUG: Putting initial VPC Index for %s with RouterIDs: %v\n", vpcKey, index.RouterIds)
		if err := client.PutVPCIndex(index); err != nil {
			fmt.Printf("ERROR: Failed to put initial VPC Index for %s: %v\n", vpcKey, err)
			// Consider returning the error if this failure is critical
			// return fmt.Errorf("failed to put initial VPC Index for %s: %w", vpcKey, err)
		}
	}
	fmt.Println("INFO: Finished VPC index creation/update loop.")

	// 4. For each other resource type... update the VPCIndex
	fmt.Println("INFO: Starting resource ID addition to VPC indexes.") // Log start of updates

	// *** ADD BACK INSTANCE PROCESSING ***
	instances, err := client.ListInstances()
	if err != nil {
		return fmt.Errorf("failed listing instances for index sync: %w", err)
	}
	fmt.Printf("INFO: Updating indexes for %d instances.\n", len(instances))
	for _, inst := range instances {
		if err := client.UpdateVPCIndex(inst.Provider, inst.VPCID, inst.ID, "instance"); err != nil {
			fmt.Printf("ERROR: Failed updating index for instance %s (VPC %s): %v\n", inst.ID, inst.VPCID, err)
			// return err // Decide if fatal
		}
	}
	// *** END ADD BACK INSTANCE PROCESSING ***

	subnets, err := client.ListSubnets()
	if err != nil {
		return fmt.Errorf("failed listing subnets for index sync: %w", err)
	}
	fmt.Printf("INFO: Updating indexes for %d subnets.\n", len(subnets))
	for _, s := range subnets {
		if err := client.UpdateVPCIndex(s.Provider, s.VpcId, s.SubnetId, "subnet"); err != nil {
			fmt.Printf("ERROR: Failed updating index for subnet %s (VPC %s): %v\n", s.SubnetId, s.VpcId, err)
		}
	}

	natgws, err := client.ListNATGateways()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d NAT gateways.\n", len(natgws)) // Added log
	for _, ng := range natgws {
		if err := client.UpdateVPCIndex(ng.Provider, ng.VpcId, ng.ID, "natgw"); err != nil {
			fmt.Printf("ERROR: Failed updating index for NAT gateway %s (VPC %s): %v\n", ng.ID, ng.VpcId, err)
			// return err // Decide if fatal
		}
	}

	rts, err := client.ListRouteTables()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d route tables.\n", len(rts)) // Added log
	for _, rt := range rts {
		if err := client.UpdateVPCIndex(rt.Provider, rt.VpcID, rt.ID, "rt"); err != nil {
			fmt.Printf("ERROR: Failed updating index for route table %s (VPC %s): %v\n", rt.ID, rt.VpcID, err)
			// return err // Decide if fatal
		}
	}

	igws, err := client.ListInternetGateways()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d internet gateways.\n", len(igws)) // Added log
	for _, igw := range igws {
		if err := client.UpdateVPCIndex(igw.Provider, igw.AttachedVpcId, igw.ID, "igw"); err != nil {
			fmt.Printf("ERROR: Failed updating index for internet gateway %s (VPC %s): %v\n", igw.ID, igw.AttachedVpcId, err)
			// return err // Decide if fatal
		}
	}

	pubIPs, err := client.ListPublicIPs()
	if err != nil {
		return fmt.Errorf("failed listing public IPs for index sync: %w", err)
	}
	fmt.Printf("INFO: Updating indexes for %d public IPs.\n", len(pubIPs)) // Added log
	for _, pubIP := range pubIPs {
		if pubIP.VPCId != "" {
			if err := client.UpdateVPCIndex(pubIP.Provider, pubIP.VPCId, pubIP.ID, "publicIP"); err != nil {
				fmt.Printf("ERROR: Failed to update VPC index for Public IP %s (VPC %s): %v\n", pubIP.ID, pubIP.VPCId, err)
				// return err // Decide if fatal
			}
		}
	}

	vpceps, err := client.ListVPCEndpoints()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d VPC endpoints.\n", len(vpceps)) // Added log
	for _, vpce := range vpceps {
		if err := client.UpdateVPCIndex(vpce.Provider, vpce.VPCId, vpce.ID, "vpce"); err != nil {
			fmt.Printf("ERROR: Failed updating index for VPC endpoint %s (VPC %s): %v\n", vpce.ID, vpce.VPCId, err)
			// return err // Decide if fatal
		}
	}

	secGroups, err := client.ListSecurityGroups()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d security groups.\n", len(secGroups)) // Added log
	for _, sg := range secGroups {
		if err := client.UpdateVPCIndex(sg.Provider, sg.VpcID, sg.ID, "securityGroup"); err != nil {
			fmt.Printf("ERROR: Failed updating index for security group %s (VPC %s): %v\n", sg.ID, sg.VpcID, err)
			// return err // Decide if fatal
		}
	}

	acls, err := client.ListACLs()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d ACLs.\n", len(acls)) // Added log
	for _, acl := range acls {
		if err := client.UpdateVPCIndex(acl.Provider, acl.VpcID, acl.ID, "acl"); err != nil {
			fmt.Printf("ERROR: Failed updating index for ACL %s (VPC %s): %v\n", acl.ID, acl.VpcID, err)
			// return err // Decide if fatal
		}
	}

	lbs, err := client.ListLBs()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d load balancers.\n", len(lbs)) // Added log
	for _, lb := range lbs {
		if err := client.UpdateVPCIndex(lb.Provider, lb.VPCID, lb.ID, "lb"); err != nil {
			fmt.Printf("ERROR: Failed updating index for load balancer %s (VPC %s): %v\n", lb.ID, lb.VPCID, err)
			// return err // Decide if fatal
		}
	}

	vpncs, err := client.ListVPNConcentrators()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d VPN concentrators.\n", len(vpncs)) // Added log
	for _, vpnc := range vpncs {
		if err := client.UpdateVPCIndex(vpnc.Provider, vpnc.VpcID, vpnc.ID, "vpnc"); err != nil {
			fmt.Printf("ERROR: Failed updating index for VPN concentrator %s (VPC %s): %v\n", vpnc.ID, vpnc.VpcID, err)
			// return err // Decide if fatal
		}
	}

	netifs, err := client.ListNetworkInterfaces()
	if err != nil {
		return err // Consider logging error context
	}
	fmt.Printf("INFO: Updating indexes for %d network interfaces.\n", len(netifs)) // Added log
	for _, netif := range netifs {
		if err := client.UpdateVPCIndex(netif.Provider, netif.VPCID, netif.ID, "netif"); err != nil {
			fmt.Printf("ERROR: Failed updating index for network interface %s (VPC %s): %v\n", netif.ID, netif.VPCID, err)
			// return err // Decide if fatal
		}
	}

	// 5. Handle Routers based on Provider
	fmt.Println("INFO: Starting non-AWS router index updates.") // Log start
	for _, router := range allRouters {
		if router.Provider != "aws" {
			if router.VPCId != "" {
				fmt.Printf("DEBUG: Updating index for non-AWS router %s (VPC %s)\n", router.ID, router.VPCId)
				if err := client.UpdateVPCIndex(router.Provider, router.VPCId, router.ID, "router"); err != nil {
					fmt.Printf("ERROR: Failed to update VPC index for non-AWS router %s (VPC %s): %v\n", router.ID, router.VPCId, err)
				}
			}
		}
	}

	fmt.Println("INFO: VPC Index synchronization complete.") // Log end
	return nil
}

// VPCIndex
// Updated PutVPCIndex to log the key being stored.
func (client *boltClient) PutVPCIndex(vpcIndex *types.VPCIndex) error {
	key := vpcIndex.DbId()
	//fmt.Printf("DEBUG: Putting VPCIndex with key: %s, value: %+v\n", key, vpcIndex)
	err := update(client, vpcIndex, key, vpcIndexTable)
	if err != nil {
		fmt.Printf("ERROR: Failed to put VPCIndex %+v: %v\n", vpcIndex, err)
	}
	return err
}

// Updated GetVPCIndex to log the exact key used.
func (client *boltClient) GetVPCIndex(id string) (*types.VPCIndex, error) {
	//fmt.Printf("DEBUG: Calling get for VPCIndex with key: %s\n", id)
	vpcIndex, err := get[types.VPCIndex](client, id, vpcIndexTable)
	if err != nil {
		fmt.Printf("ERROR: Failed to retrieve VPCIndex for ID %s: %v\n", id, err)
	} else if vpcIndex == nil {
		fmt.Printf("DEBUG: GetVPCIndex returned nil for ID %s error = %v \n", id, err)
	}
	return vpcIndex, err
}

func (client *boltClient) ListVPCIndex() ([]*types.VPCIndex, error) {
	return list[types.VPCIndex](client, vpcIndexTable)
}

func (client *boltClient) DeleteVPCIndex(id string) error {
	return delete_(client, id, vpcIndexTable)
}

// New function to call SyncVPCIndexes after all resource sync is complete.
func (client *boltClient) FinalizeSync() error {
	return client.SyncVPCIndexes()
}
