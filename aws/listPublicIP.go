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
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListPublicIPs called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListPublicIPs: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListPublicIPs: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}

	// REMOVED: c.creds = params.Creds
	// REMOVED: c.accountID = params.AccountId
	// REMOVED: accountId := params.GetAccountId()
	// REMOVED: if c.accountID == "" {
	// REMOVED: 	accountId = c.defaultAccountID
	// REMOVED: }

	//builder := newFilterBuilder() // Filters are not directly applied to DescribeAddresses
	//filters := builder.build()

	c.logger.Infof("[AccountID: %s] Requesting Public IP for account id %s", operationalAccountID, operationalAccountID)
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allPublicIPs  []types.PublicIP
			allErrors     []error
			resultChannel = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListPublicIPs: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListPublicIPs: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions {
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListPublicIPs: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getPublicIPForRegion
				pips, err := c.getPublicIPForRegion(ctx, regionName, accID)
				resultChannel <- regionResult{
					region: regionName,
					pips:   pips,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListPublicIPs: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListPublicIPs: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allPublicIPs = append(allPublicIPs, result.pips...)
			}
		}
		c.logger.Infof("[AccountID: %s] Found %d public ips across %d regions", operationalAccountID, len(allPublicIPs), len(regions))

		if len(allErrors) > 0 {
			return allPublicIPs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allPublicIPs, nil
	}
	// Pass operationalAccountID to getPublicIPForRegion
	return c.getPublicIPForRegion(ctx, params.Region, operationalAccountID)
}

// getPublicIPForRegion now accepts operationalAccountID
func (c *Client) getPublicIPForRegion(ctx context.Context, regionName string, operationalAccountID string) ([]types.PublicIP, error) {
	c.logger.Debugf("[AccountID: %s] getPublicIPForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getPublicIPForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Call DescribeAddresses - Do not pass instance filters here
	resp, err := client.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{})
	if err != nil {
		// Log the specific error
		c.logger.Errorf("[AccountID: %s] DescribeAddresses failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, fmt.Errorf("could not get AWS public IPs for region %s, account %s: %w", regionName, operationalAccountID, err)
	}

	// Call getInstances - decide if ANY filters are needed here. For now, call without filters.
	builder := newFilterBuilder()
	// If specific instance filters ARE needed (e.g., only running), add them here:
	// builder.add("instance-state-name", "running")
	// Use operationalAccountID when calling getInstances
	reservations, err := c.getInstances(ctx, operationalAccountID, regionName, builder) // Pass builder directly
	if err != nil {
		// Log warning but continue, as we might still process EIPs without instance correlation
		c.logger.Warnf("[AccountID: %s] Failed to get instances in region %s: %v. VPCId might be missing for some EIPs.", operationalAccountID, regionName, err)
		reservations = []awstypes.Reservation{} // Ensure reservations is an empty slice, not nil
	}
	// Pass operationalAccountID to convertPublicIPs
	return c.convertPublicIPs(operationalAccountID, regionName, resp.Addresses, reservations)
}

// convertPublicIPs's 'account' parameter will now be the operationalAccountID
func (c *Client) convertPublicIPs(account, region string, addresses []awstypes.Address, reservations []awstypes.Reservation) ([]types.PublicIP, error) {
	c.logger.Debugf("[AccountID: %s] convertPublicIPs: Converting public IPs for region %s.", account, region)
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
					if privateIP == "" && len(eni.PrivateIpAddresses) > 0 { // Check len again before accessing [0]
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
			if details.PrivateIP != "" { // Prioritize private IP from instance ENI if available
				privateIP = details.PrivateIP
			}

			// Sanity check InstanceId (EIP association vs Instance data)
			if instanceID != "" && details.InstanceId != "" && instanceID != details.InstanceId {
				c.logger.Warnf("[AccountID: %s] Mismatch: EIP %s associated with instance %s (from DescribeAddresses), but instance details map shows instance %s for region %s", account, publicIP, instanceID, details.InstanceId, region)
				instanceID = details.InstanceId // Prefer instance data if there's a discrepancy and both exist
			} else if instanceID == "" && details.InstanceId != "" {
				instanceID = details.InstanceId
			}
		} else if instanceID != "" {
			// This EIP is associated with an instance, but we didn't find that instance's public IP in the DescribeInstances scan.
			// This could happen if the instance was terminated between DescribeInstances and DescribeAddresses, or if DescribeInstances filters missed it.
			// We'll use the VPC ID from the EIP's NetworkInterface if available.
			if address.NetworkInterfaceId != nil {
				// To get VPC ID for an EIP not directly found on an instance, you might need to DescribeNetworkInterfaces
				// using address.NetworkInterfaceId. For simplicity here, we'll leave VPCID potentially blank if not found via instance.
				c.logger.Warnf("[AccountID: %s] Elastic IP %s associated with instance %s (per DescribeAddresses), but instance details not found or public IP not detected on instance's ENIs in DescribeInstances result for region %s. VPCId might be missing or derived from EIP's ENI if logic was added.", account, publicIP, instanceID, region)
			}
		}

		allPublicIPsMap[publicIP] = types.PublicIP{
			ID:                 aws.ToString(address.AllocationId), // Use AllocationId as the primary ID for EIPs
			Region:             region,
			InstanceId:         instanceID,
			PublicIP:           publicIP,
			Provider:           providerName,
			AccountID:          account, // This is now the operationalAccountID
			Type:               ipType,
			Labels:             convertTags(address.Tags), // Ensure convertTags is accessible
			PrivateIP:          privateIP,
			VPCId:              vpcID,
			NetworkInterfaceId: aws.ToString(address.NetworkInterfaceId),
		}
	}

	// 3. Add any remaining ephemeral IPs (found on instances but not matching any EIP)
	for publicIP, details := range instanceDetailsByPublicIP {
		if _, exists := allPublicIPsMap[publicIP]; !exists {
			// This public IP was found on an instance but is not an EIP (not in DescribeAddresses output)
			allPublicIPsMap[publicIP] = types.PublicIP{
				ID:                 details.InstanceId + "_" + publicIP, // Create a synthetic ID for ephemeral IPs
				VPCId:              details.VPCId,
				Region:             region,
				InstanceId:         details.InstanceId,
				PublicIP:           publicIP,
				PrivateIP:          details.PrivateIP,
				Provider:           providerName,
				AccountID:          account, // This is now the operationalAccountID
				Type:               "ephemeral",
				NetworkInterfaceId: "", // Ephemeral IPs are directly on instance ENIs, not separate EIP ENIs
				// Labels for ephemeral IPs are typically the instance's tags.
				// If you need specific labels for the ephemeral IP itself, this would require different logic.
			}
		}
	}

	// 4. Convert map back to slice
	finalPublicIPs := make([]types.PublicIP, 0, len(allPublicIPsMap))
	for _, pip := range allPublicIPsMap {
		finalPublicIPs = append(finalPublicIPs, pip)
	}
	c.logger.Debugf("[AccountID: %s] convertPublicIPs: Converted %d public IPs for region %s.", account, len(finalPublicIPs), region)
	return finalPublicIPs, nil
}

// Assuming getInstances and convertTags are defined elsewhere and accessible.
// If convertTags is not, a simple version:
// func convertTags(tags []awstypes.Tag) map[string]string {
// 	labels := make(map[string]string)
// 	for _, t := range tags {
// 		labels[aws.ToString(t.Key)] = aws.ToString(t.Value)
// 	}
// 	return labels
// }
