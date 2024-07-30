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
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {

	c.accountID = params.AccountId
	c.creds = params.Creds
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
		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}
		for _, region := range regions {
			wg.Add(1)
			go func(regionName string) {
				defer wg.Done()
				routers, err := c.getRoutersForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region:  regionName,
					routers: routers,
					err:     err,
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
				allRouters = append(allRouters, result.routers...)
			}
		}

		if len(allErrors) > 0 {
			return allRouters, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allRouters, nil
	}
	return c.getRoutersForRegion(ctx, params.Region, filters)
}

func (c *Client) getRoutersForRegion(ctx context.Context, region string, filters []awsTypes.Filter) ([]types.Router, error) {

	var routers []types.Router
	var CIDRBlock string

	client, err := c.getEC2Client(ctx, c.accountID, region)
	if err != nil {
		return nil, err
	}

	paginator := ec2.NewDescribeTransitGatewaysPaginator(client, &ec2.DescribeTransitGatewaysInput{
		Filters: filters,
	})

	for paginator.HasMorePages() {

		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, tgw := range page.TransitGateways {
			var name string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range tgw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			if tgw.Options != nil && len(tgw.Options.TransitGatewayCidrBlocks) > 0 {
				CIDRBlock = tgw.Options.TransitGatewayCidrBlocks[0]
			}
			//tgwLink : =

			routers = append(routers, types.Router{
				ID:        *tgw.TransitGatewayId,
				Provider:  c.GetName(),
				Name:      name,
				Region:    region,
				State:     string(tgw.State),
				Labels:    labels,
				CIDRBlock: CIDRBlock,
				AccountID: *tgw.OwnerId,
				CreatedAt: *tgw.CreationTime,
				SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/vpc/home?region=%s#TransitGateways:transitGatewayId=%s", region, region, *tgw.TransitGatewayId),
			})
		}
	}

	return routers, nil
}
