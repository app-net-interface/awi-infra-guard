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

package aws

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListNetworkInterfaces(ctx context.Context, params *infrapb.ListNetworkInterfacesRequest) ([]types.NetworkInterface, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListNetworkInterfaces called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListNetworkInterfaces: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListNetworkInterfaces: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] Listing network interfaces for vpc %s and region %s ", operationalAccountID, params.VpcId, params.Region)

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	//builder.withAvailabilityZone(params.GetZone()) // Zone is not a direct filter for DescribeNetworkInterfaces
	filters := builder.build()

	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg                   sync.WaitGroup
			allNetworkInterfaces []types.NetworkInterface
			allErrors            []error
			resultChannel        = make(chan regionResult)
		)
		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListNetworkInterfaces: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListNetworkInterfaces: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListNetworkInterfaces: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getNetworkInterfacesForRegion
				networkInterfaces, err := c.getNetworkInterfacesForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					nifs:   networkInterfaces,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListNetworkInterfaces: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListNetworkInterfaces: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allNetworkInterfaces = append(allNetworkInterfaces, result.nifs...)
			}
		}
		c.logger.Infof("[AccountID: %s] ListNetworkInterfaces: Found %d network interfaces across %d regions", operationalAccountID, len(allNetworkInterfaces), len(regions))

		if len(allErrors) > 0 {
			return allNetworkInterfaces, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allNetworkInterfaces, nil
	}
	c.logger.Debugf("[AccountID: %s] ListNetworkInterfaces: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getNetworkInterfacesForRegion
	return c.getNetworkInterfacesForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getNetworkInterfacesForRegion now accepts operationalAccountID
func (c *Client) getNetworkInterfacesForRegion(ctx context.Context, regionName string, filters []awstypes.Filter, operationalAccountID string) ([]types.NetworkInterface, error) {
	c.logger.Debugf("[AccountID: %s] getNetworkInterfacesForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getNetworkInterfacesForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	resp, err := client.DescribeNetworkInterfaces(ctx, &ec2.DescribeNetworkInterfacesInput{
		Filters: filters,
	})
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getNetworkInterfacesForRegion: DescribeNetworkInterfaces failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Pass operationalAccountID as the 'account' parameter to convertNetworkInterfaces
	return convertNetworkInterfaces(c.defaultAccountID, c.defaultRegion, operationalAccountID, regionName, resp.NetworkInterfaces), nil
}

// convertNetworkInterfaces's 'account' parameter will now be the operationalAccountID
func convertNetworkInterfaces(defaultAccount, defaultRegion, account, region string, nis []awstypes.NetworkInterface) []types.NetworkInterface {
	if region == "" {
		region = defaultRegion
	}
	// The 'account' parameter is now the operationalAccountID passed from getNetworkInterfacesForRegion.
	// If it's empty, it falls back to defaultAccount, but ideally, operationalAccountID should always be set.
	if account == "" {
		account = defaultAccount
	}
	networkInterfaces := make([]types.NetworkInterface, 0, len(nis))
	for _, ni := range nis {
		name := getTagName(ni.TagSet) // TagSet is the correct field for tags on NetworkInterface
		var privateIPs []string
		for _, ip := range ni.PrivateIpAddresses {
			if ip.PrivateIpAddress != nil {
				privateIPs = append(privateIPs, aws.ToString(ip.PrivateIpAddress))
			}
		}

		var publicIP string
		if ni.Association != nil && ni.Association.PublicIp != nil {
			publicIP = *ni.Association.PublicIp
		}

		var securityGroups []string
		for _, sg := range ni.Groups {
			securityGroups = append(securityGroups, aws.ToString(sg.GroupId))
		}
		publicDNSName := ""
		if ni.Association != nil {
			publicDNSName = aws.ToString(ni.Association.PublicDnsName)
		}

		// Determine Interface Type (primary/secondary/unattached)
		interfaceType := "unattached" // Default if not attached
		if ni.Attachment != nil {
			if ni.Attachment.DeviceIndex != nil {
				if aws.ToInt32(ni.Attachment.DeviceIndex) == 0 { // aws.ToInt32 handles *int32
					interfaceType = "primary"
				} else {
					interfaceType = "secondary"
				}
			} else {
				// If attached but DeviceIndex is somehow nil, treat as secondary? Or "unknown"?
				interfaceType = "secondary" // Assuming non-primary if index is missing but attached
			}
		}

		networkInterface := types.NetworkInterface{
			ID:               aws.ToString(ni.NetworkInterfaceId),
			Name:             aws.ToString(name), // name is *string
			Provider:         providerName,
			AccountID:        account, // This is now the operationalAccountID
			VPCID:            aws.ToString(ni.VpcId),
			SubnetID:         aws.ToString(ni.SubnetId),
			AvailabilityZone: aws.ToString(ni.AvailabilityZone),
			Region:           region,
			PrivateIPs:       privateIPs,
			PublicIP:         publicIP,
			SecurityGroupIDs: securityGroups,
			MacAddress:       aws.ToString(ni.MacAddress),
			PrivateDNSName:   aws.ToString(ni.PrivateDnsName), // This field exists on awstypes.NetworkInterface
			PublicDNSName:    publicDNSName,
			Description:      aws.ToString(ni.Description),
			Labels:           getTags(ni.TagSet), // Use getTags with ni.TagSet
			Status:           string(ni.Status),
			InterfaceType:    interfaceType, // Assign the determined type here
		}
		networkInterfaces = append(networkInterfaces, networkInterface)
	}
	return networkInterfaces
}

func getTags(tagSet []awstypes.Tag) map[string]string { // Renamed from convertTags to avoid conflict if another exists
	labels := make(map[string]string) // Initialize to avoid nil map
	for _, tag := range tagSet {
		labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
	}
	return labels
}

// getTagName was already defined in listInstance.go, ensure it's accessible or redefine if necessary
// For clarity, if it's specific to this package and not shared, it could be:
// func getTagName(tags []awstypes.Tag) *string {
// 	for _, tag := range tags {
// 		if aws.ToString(tag.Key) == "Name" {
// 			return tag.Value
// 		}
// 	}
// 	return nil
// }
