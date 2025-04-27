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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListPublicIPs(ctx context.Context, params *infrapb.ListPublicIPsRequest) ([]types.PublicIP, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId
	accountId := params.GetAccountId()

	if c.accountID == "" {
		accountId = c.defaultAccountID
	}

	//builder := newFilterBuilder()
	//filters := builder.build()

	c.logger.Infof("*******Requesting Public IP for account id %s ****", accountId)
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allPublicIPs  []types.PublicIP
			allErrors     []error
			resultChannel = make(chan regionResult)
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
				pips, err := c.getPublicIPForRegion(ctx, regionName)
				resultChannel <- regionResult{
					region: regionName,
					pips:   pips,
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
				allPublicIPs = append(allPublicIPs, result.pips...)
			}
		}
		c.logger.Infof("In account %s Found %d public ips across %d regions", c.accountID, len(allPublicIPs), len(regions))

		if len(allErrors) > 0 {
			return allPublicIPs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allPublicIPs, nil
	}
	return c.getPublicIPForRegion(ctx, params.Region)
}

func (c *Client) getPublicIPForRegion(ctx context.Context, regionName string) ([]types.PublicIP, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeAddresses - Do not pass instance filters here
	resp, err := client.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{})
	if err != nil {
		// Log the specific error
		c.logger.Errorf("DescribeAddresses failed for region %s: %v", regionName, err)
		return nil, fmt.Errorf("could not get AWS public IPs for region %s: %w", regionName, err)
	}

	// Call getInstances - decide if ANY filters are needed here. For now, call without filters.
	builder := newFilterBuilder()
	// If specific instance filters ARE needed (e.g., only running), add them here:
	// builder.add("instance-state-name", "running")
	reservations, err := c.getInstances(ctx, c.accountID, regionName, builder) // Pass builder directly
	if err != nil {
		// Log warning but continue, as we might still process EIPs without instance correlation
		c.logger.Warnf("Failed to get instances in region %s: %v. VPCId might be missing for some EIPs.", regionName, err)
		reservations = []awstypes.Reservation{} // Ensure reservations is an empty slice, not nil
	}
	return c.convertPublicIPs(c.accountID, regionName, resp.Addresses, reservations)
}

func (c *Client) convertPublicIPs(account, region string, addresses []awstypes.Address, reservations []awstypes.Reservation) ([]types.PublicIP, error) {

	allPublicIPsMap := make(map[string]types.PublicIP) // Use PublicIP address as key

	// 1. Process instance data to get VPCId/InstanceId/PrivateIP associated with ANY public IPs found
	instanceDetailsByPublicIP := make(map[string]struct {
		VPCId      string
		InstanceId string
		PrivateIP  string // Store the private IP associated with the specific ENI having the public IP
	})

	for _, reservation := range reservations {
		for _, inst := range reservation.Instances {
			instanceID := aws.ToString(inst.InstanceId)
			vpcId := ""
			if inst.VpcId != nil {
				vpcId = *inst.VpcId
			}

			// Iterate through Network Interfaces attached to the instance
			for _, eni := range inst.NetworkInterfaces {
				var publicIP, privateIP string

				// Check for associated public IP on this ENI
				if eni.Association != nil && eni.Association.PublicIp != nil {
					publicIP = aws.ToString(eni.Association.PublicIp)
				}

				// If no public IP on association, check the main PublicIpAddress field (primary ENI case)
				if publicIP == "" && aws.ToString(inst.PublicIpAddress) != "" {
					if eni.Attachment != nil && aws.ToInt32(eni.Attachment.DeviceIndex) == 0 {
						publicIP = aws.ToString(inst.PublicIpAddress)
					}
				}

				if publicIP == "" {
					continue // No public IP found for this specific ENI
				}

				// Find the primary private IP for this specific ENI
				if len(eni.PrivateIpAddresses) > 0 {
					for _, privAddr := range eni.PrivateIpAddresses {
						if aws.ToBool(privAddr.Primary) {
							privateIP = aws.ToString(privAddr.PrivateIpAddress)
							break // Found primary private IP for this ENI
						}
					}
					// Fallback if no primary is marked (shouldn't happen often)
					if privateIP == "" {
						privateIP = aws.ToString(eni.PrivateIpAddresses[0].PrivateIpAddress)
					}
				}

				// Store details keyed by the public IP found on this ENI
				if _, exists := instanceDetailsByPublicIP[publicIP]; !exists {
					instanceDetailsByPublicIP[publicIP] = struct {
						VPCId      string
						InstanceId string
						PrivateIP  string
					}{
						VPCId:      vpcId,
						InstanceId: instanceID,
						PrivateIP:  privateIP,
					}
				}
			}
		}
	}

	// 2. Process Elastic IPs (addresses)
	for _, address := range addresses {
		publicIP := aws.ToString(address.PublicIp)
		if publicIP == "" {
			continue // Skip invalid EIP entries
		}

		instanceID := aws.ToString(address.InstanceId) // Instance associated with EIP from DescribeAddresses
		vpcID := ""                                    // Default VPC ID is empty
		privateIP := aws.ToString(address.PrivateIpAddress) // Private IP from EIP association info
		ipType := "static" // Assume static (Elastic IP)

		// Check if this EIP's public IP was found on an instance's ENI
		if details, found := instanceDetailsByPublicIP[publicIP]; found {
			vpcID = details.VPCId
			if details.PrivateIP != "" {
				privateIP = details.PrivateIP
			}

			// Sanity check InstanceId (EIP association vs Instance data)
			if instanceID != "" && details.InstanceId != "" && instanceID != details.InstanceId {
				c.logger.Warnf("Mismatch: EIP %s associated with instance %s (from DescribeAddresses), but instance details map shows instance %s for region %s", publicIP, instanceID, details.InstanceId, region)
				instanceID = details.InstanceId
			} else if instanceID == "" && details.InstanceId != "" {
				instanceID = details.InstanceId
			}
		} else if instanceID != "" {
			c.logger.Warnf("Elastic IP %s associated with instance %s (per DescribeAddresses), but instance details not found or public IP not detected on instance's ENIs in DescribeInstances result for region %s. VPCId will be missing.", publicIP, instanceID, region)
		}

		allPublicIPsMap[publicIP] = types.PublicIP{
			ID:                 aws.ToString(address.AllocationId),
			Region:             region,
			InstanceId:         instanceID,
			PublicIP:           publicIP,
			Provider:           providerName,
			AccountID:          account,
			Type:               ipType,
			Labels:             convertTags(address.Tags),
			PrivateIP:          privateIP,
			VPCId:              vpcID,
			NetworkInterfaceId: aws.ToString(address.NetworkInterfaceId),
		}
	}

	// 3. Add any remaining ephemeral IPs (found on instances but not matching any EIP)
	for publicIP, details := range instanceDetailsByPublicIP {
		if _, exists := allPublicIPsMap[publicIP]; !exists {
			allPublicIPsMap[publicIP] = types.PublicIP{
				ID:                 details.InstanceId + "_" + publicIP,
				VPCId:              details.VPCId,
				Region:             region,
				InstanceId:         details.InstanceId,
				PublicIP:           publicIP,
				PrivateIP:          details.PrivateIP,
				Provider:           providerName,
				AccountID:          account,
				Type:               "ephemeral",
				NetworkInterfaceId: "",
			}
		}
	}

	// 4. Convert map back to slice
	finalPublicIPs := make([]types.PublicIP, 0, len(allPublicIPsMap))
	for _, pip := range allPublicIPsMap {
		finalPublicIPs = append(finalPublicIPs, pip)
	}

	return finalPublicIPs, nil
}
