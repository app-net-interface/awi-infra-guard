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
	"github.com/aws/aws-sdk-go-v2/aws" // Added for aws.ToString
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListInternetGateways called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListInternetGateways: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListInternetGateways: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListInternetGateways called with VPC ID: %s, Region: %s", operationalAccountID, params.GetVpcId(), params.GetRegion())

	// REMOVED: c.creds = params.Creds
	// REMOVED: c.accountID = params.AccountId

	builder := newFilterBuilder()
	// IGWs can be filtered by attachment.vpc-id
	if params.GetVpcId() != "" {
		builder.withVPC(params.GetVpcId())
	}
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allIGWs       []types.IGW
			allErrors     []error
			resultChannel = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListInternetGateways: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListInternetGateways: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListInternetGateways: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getInternetGatewaysForRegion
				regIgws, err := c.getInternetGatewaysForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName, // Correctly use regionName from goroutine param
					igws:   regIgws,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}
		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListInternetGateways: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListInternetGateways: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allIGWs = append(allIGWs, result.igws...)
			}
		}
		c.logger.Infof("[AccountID: %s] ListInternetGateways: Found %d IGWs across %d regions", operationalAccountID, len(allIGWs), len(regions))

		if len(allErrors) > 0 {
			return allIGWs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allIGWs, nil // Added return for success case
	}
	c.logger.Debugf("[AccountID: %s] ListInternetGateways: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getInternetGatewaysForRegion
	return c.getInternetGatewaysForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getInternetGatewaysForRegion now accepts operationalAccountID
func (c *Client) getInternetGatewaysForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter, operationalAccountID string) ([]types.IGW, error) {
	c.logger.Debugf("[AccountID: %s] getInternetGatewaysForRegion: Region %s", operationalAccountID, regionName)
	var igws []types.IGW
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getInternetGatewaysForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	paginator := ec2.NewDescribeInternetGatewaysPaginator(client, &ec2.DescribeInternetGatewaysInput{
		Filters: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx) // Use ctx from parameters
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getInternetGatewaysForRegion: DescribeInternetGateways NextPage failed for region %s: %v", operationalAccountID, regionName, err)
			return nil, err
		}
		for _, igw := range page.InternetGateways {
			var vpcId, name string
			var state string = "detached" // Default if no attachments or state is nil
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range igw.Tags {
				if aws.ToString(tag.Key) == "Name" || aws.ToString(tag.Key) == "name" {
					name = aws.ToString(tag.Value)
				}
				labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
			}

			if len(igw.Attachments) > 0 {
				vpcId = aws.ToString(igw.Attachments[0].VpcId)
				if igw.Attachments[0].State != "" { // Check if state is not an empty string
					state = string(igw.Attachments[0].State)
				}
			}

			igws = append(igws, types.IGW{
				ID:            aws.ToString(igw.InternetGatewayId),
				Provider:      c.GetName(),
				AccountID:     aws.ToString(igw.OwnerId), // Uses OwnerId from the resource
				Name:          name,
				AttachedVpcId: vpcId,
				Region:        regionName,
				State:         state,
				Labels:        labels,
				SelfLink:      fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#InternetGateway:internetGatewayId=%s", regionName, regionName, aws.ToString(igw.InternetGatewayId)),
			})
		}
	}
	return igws, nil
}
