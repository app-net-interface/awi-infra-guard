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

	"github.com/app-net-interface/awi-infra-guard/types"
	//"github.com/boltdb/bolt"
)

// Updated UpdateVPCIndex to add debug logs at function entry and exit.
func (client *boltClient) UpdateVPCIndex(provider, vpcID, resourceId, resourceType string) error {
	//fmt.Printf("DEBUG: Updating VPCIndex for VPCID: %s, ResourceID: %s, ResourceType: %s\n", vpcID, resourceId, resourceType)
	key := provider + ":" + vpcID
	vpcIndex, err := client.GetVPCIndex(key)
	if vpcIndex == nil {
		fmt.Printf("WARN: No existing VPCIndex for VPCID %s, adding new one for resource %s \n", vpcID, resourceType)
		// vpcIndex = &types.VPCIndex{
		// 	VpcId:            vpcID,
		// 	Provider:         provider,
		// 	InstanceIds:      []string{},
		// 	AclIds:           []string{},
		// 	SecurityGroupIds: []string{},
		// 	NatGatewayIds:    []string{},
		// 	VpcEndpointIds:   []string{},
		// 	LbIds:            []string{},
		// }
		return err
	}
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
	case "rts":
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
	default:
		fmt.Printf("WARN: Unknown resource type: %s\n", resourceType)
	}
	//fmt.Printf("DEBUG: Updated VPCIndex: %+v for VPCID %s\n", vpcIndex, vpcID)
	return client.PutVPCIndex(vpcIndex)
}

// New function to synchronize VPCIndexes after all resource sync.
func (client *boltClient) SyncVPCIndexes() error {
	// Remove all existing VPCIndexes.
	indexes, err := client.ListVPCIndex()
	if err != nil {
		return err
	}
	for _, idx := range indexes {
		if err := client.DeleteVPCIndex(idx.DbId()); err != nil {
			fmt.Printf("ERROR : Deleting VPC Index %s ", err)
			return err
		}
	}

	// Recreate VPCIndexes from all VPCs.
	vpcs, err := client.ListVPCs()
	if err != nil {
		return err
	}
	for _, vpc := range vpcs {
		index := &types.VPCIndex{
			VpcId:            vpc.ID,
			Provider:         vpc.Provider,
			InstanceIds:      []string{},
			AclIds:           []string{},
			SecurityGroupIds: []string{},
			NatGatewayIds:    []string{},
			VpcEndpointIds:   []string{},
			LbIds:            []string{},
		}
		if err := client.PutVPCIndex(index); err != nil {
			fmt.Printf("ERROR : Put VPC Index %s ", err)
			return err
		}
	}

	// For each resource type, update the VPCIndex accordingly.
	instances, err := client.ListInstances()
	if err != nil {
		return err
	}
	for _, instance := range instances {
		if err := client.UpdateVPCIndex(instance.Provider, instance.VPCID, instance.ID, "instance"); err != nil {
			return err
		}
	}

	natgws, err := client.ListNATGateways()
	if err != nil {
		return err
	}
	for _, ng := range natgws {
		if err := client.UpdateVPCIndex(ng.Provider, ng.VpcId, ng.ID, "natgw"); err != nil {
			return err
		}
	}

	subnets, err := client.ListSubnets()
	if err != nil {
		return err
	}
	for _, s := range subnets {
		if err := client.UpdateVPCIndex(s.Provider, s.VpcId, s.SubnetId, "subnet"); err != nil {
			return err
		}
	}

	rts, err := client.ListRouteTables()
	if err != nil {
		return err
	}
	for _, rt := range rts {
		if err := client.UpdateVPCIndex(rt.Provider, rt.VpcID, rt.ID, "rt"); err != nil {
			return err
		}
	}

	igws, err := client.ListInternetGateways()
	if err != nil {
		return err
	}
	for _, igw := range igws {
		if err := client.UpdateVPCIndex(igw.Provider, igw.AttachedVpcId, igw.ID, "igw"); err != nil {
			return err
		}
	}

	pubIPs, err := client.ListPublicIPs()
	if err != nil {
		return err
	}
	for _, pubIP := range pubIPs {
		if err := client.UpdateVPCIndex(pubIP.Provider, pubIP.VPCID, pubIP.ID, "publicIP"); err != nil {
			return err
		}
	}

	vpceps, err := client.ListVPCEndpoints()
	if err != nil {
		return err
	}
	for _, vpce := range vpceps {
		if err := client.UpdateVPCIndex(vpce.Provider, vpce.VPCId, vpce.ID, "vpce"); err != nil {
			return err
		}
	}

	secGroups, err := client.ListSecurityGroups()
	if err != nil {
		return err
	}
	for _, sg := range secGroups {
		if err := client.UpdateVPCIndex(sg.Provider, sg.VpcID, sg.ID, "securityGroup"); err != nil {
			return err
		}
	}

	acls, err := client.ListACLs()
	if err != nil {
		return err
	}
	for _, acl := range acls {
		if err := client.UpdateVPCIndex(acl.Provider, acl.VpcID, acl.ID, "acl"); err != nil {
			return err
		}
	}

	lbs, err := client.ListLBs()
	if err != nil {
		return err
	}
	for _, lb := range lbs {
		if err := client.UpdateVPCIndex(lb.Provider, lb.VPCID, lb.ID, "lb"); err != nil {
			return err
		}
	}

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
