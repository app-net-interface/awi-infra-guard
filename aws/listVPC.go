// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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
	"strings"
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {
	// The panic below will compare the incoming params.AccountId against the
	// c.accountID that the client was initialized with or its last sticky state.
	// If the design intends for one client to handle multiple accounts sequentially
	// by changing c.accountID, this panic will fire.
	// For true parallelism without race conditions, c.accountID should not be modified per call.
	// We will proceed by NOT modifying c.accountID in this function.
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		// This panic condition might need re-evaluation based on whether c.accountID
		// is a "default" or "sticky operational" ID. For a stateless operational model,
		// this panic might be too strict if c.accountID is just a default.
		// However, given the previous logs, it was firing because c.accountID was stale
		// from a PREVIOUS operation. By not setting c.accountID here, we avoid that specific race.
		//c.logger.Warnf("ListVPC called with AccountID: %s, while client's c.accountID is: %s. Proceeding with params.AccountId.", params.AccountId, c.accountID)
		panic(fmt.Sprintf("ListVPC called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		// If no accountId in params, consider using the client's default/initial accountID.
		// This depends on desired behavior if params.AccountId is empty.
		c.logger.Infof("ListVPC: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListVPC: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}

	c.logger.Debugf("[AccountID: %s] ListVPC called with Region: '%s', Labels: %v. HasCreds: %t",
		operationalAccountID, params.GetRegion(), params.GetLabels(), params.GetCreds() != nil)

	// REMOVED: originalAccountID := c.accountID
	// REMOVED: originalCreds := c.creds (user said to ignore creds for this fix)

	if params == nil { // Should ideally not happen if operationalAccountID is derived from params.
		c.logger.Warnf("[AccountID: %s] ListVPC called with nil params, this path might be problematic.", operationalAccountID)
		params = &infrapb.ListVPCRequest{} // This will result in empty region, labels etc.
	}
	// REMOVED: Lines that set c.creds and c.accountID
	// else {
	// c.logger.Infof("[AccountID: %s] Listing VPC for region %s.", operationalAccountID, params.Region)
	// c.creds = params.Creds // Not changing c.creds based on user feedback for this fix
	// c.accountID = params.AccountId // DO NOT DO THIS - source of race condition
	// }

	builder := newFilterBuilder()
	for k, v := range params.Labels {
		builder.withTag(k, v)
	}
	filters := builder.build()
	c.logger.Debugf("[AccountID: %s] ListVPC: Applied filters: %v", operationalAccountID, filters)

	if params.Region == "" || params.Region == "all" {
		c.logger.Infof("[AccountID: %s] ListVPC: Iterating all regions.", operationalAccountID)
		var (
			wg            sync.WaitGroup
			allVPCs       []types.VPC
			allErrors     []error
			resultChannel = make(chan regionResult)
		)

		c.logger.Debugf("[AccountID: %s] ListVPC: Calling c.getAllRegions.", operationalAccountID)
		regions, err := c.getAllRegions(ctx, operationalAccountID) // Pass operationalAccountID
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListVPC: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Infof("[AccountID: %s] ListVPC: Found %d regions to process.", operationalAccountID, len(regions))

		for _, region := range regions {
			wg.Add(1)
			// Pass operationalAccountID (derived from params.AccountId) to the goroutine.
			// Do not use c.accountID here as it's shared and racy.
			// operationalCreds := params.GetCreds() // If creds were to be passed

			go func(regionName string, accountIDForOperation string /*, credsForOperation *infrapb.Credentials*/) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListVPC goroutine started for Region: '%s'.", accountIDForOperation, regionName)

				var vpcs []types.VPC
				var err error
				// Pass accountIDForOperation to getVPCsForRegion
				vpcs, err = c.getVPCsForRegion(ctx, regionName, filters, accountIDForOperation /*, credsForOperation */)

				c.logger.Debugf("[AccountID: %s] ListVPC goroutine for Region: '%s': getVPCsForRegion returned. Error: %v, VPCs found: %d.", accountIDForOperation, regionName, err, len(vpcs))
				resultChannel <- regionResult{
					region: regionName,
					vpcs:   vpcs,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID /*, operationalCreds */)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListVPC: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		processedRegions := 0
		for result := range resultChannel {
			processedRegions++
			c.logger.Debugf("[AccountID: %s] ListVPC: Received result for region '%s'. Error: %v, VPCs: %d", operationalAccountID, result.region, result.err, len(result.vpcs))
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListVPC: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allVPCs = append(allVPCs, result.vpcs...)
			}
		}
		c.logger.Debugf("[AccountID: %s] ListVPC: Finished collecting results from %d regions.", operationalAccountID, processedRegions)

		if len(allErrors) > 0 {
			c.logger.Errorf("[AccountID: %s] ListVPC: Errors occurred in %d regions: %v", operationalAccountID, len(allErrors), allErrors)
			// REMOVED: c.accountID = originalAccountID
			// REMOVED: c.creds = originalCreds
			return allVPCs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		c.logger.Infof("[AccountID: %s] ListVPC: Found %d VPCs across %d regions successfully.", operationalAccountID, len(allVPCs), len(regions))
		// REMOVED: c.accountID = originalAccountID
		// REMOVED: c.creds = originalCreds
		return allVPCs, nil
	}

	c.logger.Infof("[AccountID: %s] ListVPC: Processing specific region: '%s'.", operationalAccountID, params.Region)
	// Pass operationalAccountID to getVPCsForRegion
	vpcs, err := c.getVPCsForRegion(ctx, params.Region, filters, operationalAccountID /*, params.GetCreds() */)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] ListVPC, Region: '%s': Error from getVPCsForRegion: %v", operationalAccountID, params.Region, err)
		// REMOVED: c.accountID = originalAccountID
		// REMOVED: c.creds = originalCreds
		return nil, err
	}
	c.logger.Infof("[AccountID: %s] ListVPC, Region: '%s': Found %d VPCs.", operationalAccountID, params.Region, len(vpcs))
	// REMOVED: c.accountID = originalAccountID
	// REMOVED: c.creds = originalCreds
	return vpcs, nil
}

func (c *Client) getVPCsForRegion(ctx context.Context, region string, filters []awstypes.Filter, operationalAccountID string /*, operationalCreds *infrapb.Credentials */) ([]types.VPC, error) {
	// Use the passed-in operationalAccountID
	c.logger.Debugf("[AccountID: %s] getVPCsForRegion started for Region: '%s'. Filters: %v", operationalAccountID, region, filters)

	// Pass operationalAccountID to getEC2Client.
	// getEC2Client must be updated to accept operationalAccountID and use it
	// instead of relying on c.accountID. It should also take operationalCreds if needed.
	client, err := c.getEC2Client(ctx, operationalAccountID, region /*, operationalCreds */)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getVPCsForRegion, Region: '%s': Failed to get EC2 client: %v", operationalAccountID, region, err)
		return nil, err
	}
	c.logger.Debugf("[AccountID: %s] getVPCsForRegion, Region: '%s': Successfully obtained EC2 client.", operationalAccountID, region)

	c.logger.Debugf("[AccountID: %s] getVPCsForRegion, Region: '%s': Calling DescribeVpcs API.", operationalAccountID, region)
	resp, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: filters,
	})
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getVPCsForRegion, Region: '%s': DescribeVpcs API call failed: %v", operationalAccountID, region, err)
		return nil, err
	}
	c.logger.Infof("[AccountID: %s] getVPCsForRegion, Region: '%s': DescribeVpcs API successful. Found %d VPCs.", operationalAccountID, region, len(resp.Vpcs))
	return convertVPCs(resp.Vpcs, c.defaultRegion, region), nil
}

func convertVPCs(vpcs []awstypes.Vpc, defaultRegion string, region string) []types.VPC {
	if region == "" {
		region = defaultRegion
	}

	result := make([]types.VPC, 0, len(vpcs))
	for _, vpc := range vpcs {
		var ipv6CIDR string
		if len(vpc.Ipv6CidrBlockAssociationSet) > 0 {
			for _, ipv6Association := range vpc.Ipv6CidrBlockAssociationSet {
				ipv6CIDR = fmt.Sprintf("%s,%s", *ipv6Association.Ipv6CidrBlock, ipv6CIDR)
			}
		}
		vpcLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#VpcDetails:VpcId=%s", region, region, aws.ToString(vpc.VpcId))
		project := ""
		for _, tag := range vpc.Tags {
			if strings.ToLower(*tag.Key) == "project" {
				project = *tag.Value
				break
			}
		}
		result = append(result, types.VPC{
			Name:      convertString(getTagName(vpc.Tags)),
			ID:        aws.ToString(vpc.VpcId),
			Region:    region,
			Labels:    convertTags(vpc.Tags),
			IPv4CIDR:  *vpc.CidrBlock,
			IPv6CIDR:  ipv6CIDR,
			AccountID: *vpc.OwnerId,
			Provider:  providerName,
			Project:   project,
			SelfLink:  vpcLink,
		})
	}
	return result
}
