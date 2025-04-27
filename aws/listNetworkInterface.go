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
	c.logger.Debugf("Listing network interfaces for account %s, vpc %s and region %s ", params.AccountId, params.VpcId, params.Region)

	c.creds = params.Creds
	c.accountID = params.AccountId

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	//builder.withAvailabilityZone(params.GetZone())
	filters := builder.build()

	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg                   sync.WaitGroup
			allNetworkInterfaces []types.NetworkInterface
			allErrors            []error
			resultChannel        = make(chan regionResult)
		)
		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}
		for _, region := range regions {
			wg.Add(1)
			go func(regionName string) {
				defer wg.Done()
				networkInterfaces, err := c.getNetworkInterfacesForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					nifs:   networkInterfaces,
					err:    err,
				}
			}(*region.RegionName)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("Error in region %s: %v", result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allNetworkInterfaces = append(allNetworkInterfaces, result.nifs...)
			}
		}
		c.logger.Infof("In account %s Found %d network interfaces across %d regions", c.accountID, len(allNetworkInterfaces), len(regions))

		if len(allErrors) > 0 {
			return allNetworkInterfaces, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allNetworkInterfaces, nil
	}
	return c.getNetworkInterfacesForRegion(ctx, params.Region, filters)
}

func (c *Client) getNetworkInterfacesForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.NetworkInterface, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.DescribeNetworkInterfaces(ctx, &ec2.DescribeNetworkInterfacesInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertNetworkInterfaces(c.defaultAccountID, c.defaultRegion, c.accountID, regionName, resp.NetworkInterfaces), nil
}

func convertNetworkInterfaces(defaultAccount, defaultRegion, account, region string, nis []awstypes.NetworkInterface) []types.NetworkInterface {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	networkInterfaces := make([]types.NetworkInterface, 0, len(nis))
	for _, ni := range nis {
		name := getTagName(ni.TagSet)
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
				if aws.ToInt32(ni.Attachment.DeviceIndex) == 0 {
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
			Name:             aws.ToString(name),
			Provider:         providerName,
			AccountID:        account,
			VPCID:            aws.ToString(ni.VpcId),
			SubnetID:         aws.ToString(ni.SubnetId),
			AvailabilityZone: aws.ToString(ni.AvailabilityZone),
			Region:           region,
			PrivateIPs:       privateIPs,
			PublicIP:         publicIP,
			SecurityGroupIDs: securityGroups,
			MacAddress:       aws.ToString(ni.MacAddress),
			PrivateDNSName:   aws.ToString(ni.PrivateDnsName),
			PublicDNSName:    publicDNSName,
			Description:      aws.ToString(ni.Description),
			Labels:           getTags(ni.TagSet),
			Status:           string(ni.Status),
			InterfaceType:    interfaceType, // Assign the determined type here
		}
		networkInterfaces = append(networkInterfaces, networkInterface)
	}
	return networkInterfaces
}

func getTags(tagSet []awstypes.Tag) map[string]string {
	tags := make(map[string]string)
	for _, tag := range tagSet {
		tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
	}
	return tags
}
