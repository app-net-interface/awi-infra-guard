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

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
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

		// List all regions to ensure Internet Gateways from every region are considered
		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}

		for _, region := range regions {
			wg.Add(1)
			go func(regionName string) {
				defer wg.Done()
				regIgws, err := c.getInternetGatewaysForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: *region.RegionName,
					igws:   regIgws,
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
				allIGWs = append(allIGWs, result.igws...)
			}
		}
		c.logger.Infof("In account %s Found %d IGWs across %d regions", c.accountID, len(allIGWs), len(regions))

		if len(allErrors) > 0 {
			return allIGWs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
	}
	return c.getInternetGatewaysForRegion(ctx, params.Region, filters)
}

func (c *Client) getInternetGatewaysForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter) ([]types.IGW, error) {
	var igws []types.IGW
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	paginator := ec2.NewDescribeInternetGatewaysPaginator(client, &ec2.DescribeInternetGatewaysInput{
		Filters: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, igw := range page.InternetGateways {
			var vpcId, name string
			var state string = "unattached"
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range igw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			if len(igw.Attachments) > 0 {
				vpcId = *igw.Attachments[0].VpcId
				state = string(igw.Attachments[0].State)
			}

			igws = append(igws, types.IGW{
				ID:            *igw.InternetGatewayId,
				Provider:      c.GetName(),
				AccountID:     *igw.OwnerId,
				Name:          name,
				AttachedVpcId: vpcId,
				Region:        regionName,
				State:         state,
				Labels:        labels,
				SelfLink:      fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#InternetGateway:internetGatewayId=%s", regionName, regionName, *igw.InternetGatewayId),
			})
		}
	}
	return igws, nil
}
