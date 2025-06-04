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
	"github.com/aws/aws-sdk-go-v2/aws" // Import aws package
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListRouters called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListRouters: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListRouters: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Infof("[AccountID: %s] ListRouters called. Region: %s", operationalAccountID, params.GetRegion())

	// REMOVED: c.accountID = params.AccountId
	// REMOVED: c.creds = params.Creds

	builder := newFilterBuilder()
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allRouters    []types.Router
			allErrors     []error
			resultChannel = make(chan regionResult)
		)
		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListRouters: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions {
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListRouters: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getRoutersForRegion
				routers, err := c.getRoutersForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region:  regionName,
					routers: routers,
					err:     err,
				}
			}(aws.ToString(region.RegionName), operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListRouters: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allRouters = append(allRouters, result.routers...)
			}
		}
		c.logger.Infof("[AccountID: %s] Found %d routers across %d regions", operationalAccountID, len(allRouters), len(regions))

		if len(allErrors) > 0 {
			return allRouters, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allRouters, nil
	}
	// Pass operationalAccountID to getRoutersForRegion
	return c.getRoutersForRegion(ctx, params.Region, filters, operationalAccountID)
}

func (c *Client) getRoutersForRegion(ctx context.Context, region string, filters []awsTypes.Filter, operationalAccountID string) ([]types.Router, error) {
	c.logger.Debugf("[AccountID: %s] getRoutersForRegion: Region %s", operationalAccountID, region)
	var routers []types.Router
	var CIDRBlock string

	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, region)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getRoutersForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, region, err)
		return nil, err
	}

	paginator := ec2.NewDescribeTransitGatewaysPaginator(client, &ec2.DescribeTransitGatewaysInput{
		Filters: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx) // Use ctx from parameters
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getRoutersForRegion: DescribeTransitGateways NextPage failed for region %s: %v", operationalAccountID, region, err)
			return nil, err
		}
		for _, tgw := range page.TransitGateways {
			var name string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range tgw.Tags {
				key := aws.ToString(tag.Key)
				value := aws.ToString(tag.Value)
				if key == "Name" || key == "name" {
					name = value
				}
				labels[key] = value
			}

			CIDRBlock = "" // Reset for each TGW
			if tgw.Options != nil && len(tgw.Options.TransitGatewayCidrBlocks) > 0 {
				CIDRBlock = tgw.Options.TransitGatewayCidrBlocks[0] // Slices are not pointers
			}

			routers = append(routers, types.Router{
				ID:        aws.ToString(tgw.TransitGatewayId),
				Provider:  c.GetName(),
				Name:      name,
				Region:    region,
				State:     string(tgw.State),
				Labels:    labels,
				CIDRBlock: CIDRBlock,
				AccountID: aws.ToString(tgw.OwnerId), // Owner of the resource
				CreatedAt: aws.ToTime(tgw.CreationTime),
				SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/vpc/home?region=%s#TransitGateways:transitGatewayId=%s", region, region, aws.ToString(tgw.TransitGatewayId)),
			})
		}
	}
	c.logger.Debugf("[AccountID: %s] getRoutersForRegion: Found %d routers in region %s.", operationalAccountID, len(routers), region)
	return routers, nil
}
