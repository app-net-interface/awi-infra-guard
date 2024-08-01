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
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {
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
			allNGWs       []types.NATGateway
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
				regNgws, err := c.getNATGatewaysForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: *region.RegionName,
					ngws:   regNgws,
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
				allNGWs = append(allNGWs, result.ngws...)
			}
		}
		c.logger.Infof("In account %s Found %d NGWs across %d regions", c.accountID, len(allNGWs), len(regions))

		if len(allErrors) > 0 {
			return allNGWs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
	}
	return c.getNATGatewaysForRegion(ctx, params.Region, filters)
}

func (c *Client) getNATGatewaysForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.NATGateway, error) {
	var natGateways []types.NATGateway
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{
		Filter: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, ngw := range page.NatGateways {
			var name, publicIp, privateIp string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range ngw.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}

			// Assuming the first address is the public one (if exists)
			if len(ngw.NatGatewayAddresses) > 0 {
				if ngw.NatGatewayAddresses[0].PublicIp != nil {
					publicIp = *ngw.NatGatewayAddresses[0].PublicIp
				}
				if ngw.NatGatewayAddresses[0].PrivateIp != nil {
					privateIp = *ngw.NatGatewayAddresses[0].PrivateIp
				}
			}

			natGateways = append(natGateways, types.NATGateway{
				ID:        *ngw.NatGatewayId,
				Provider:  c.GetName(),
				Name:      name,
				VpcId:     *ngw.VpcId,
				Region:    regionName,
				State:     string(ngw.State),
				PublicIp:  publicIp,
				PrivateIp: privateIp,
				SubnetId:  convertString(ngw.SubnetId),
				AccountID: c.accountID,
				Labels:    labels,
				CreatedAt: *ngw.CreateTime,
				SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#NatGatewayDetails:natGatewayId=%s", regionName, regionName, *ngw.NatGatewayId),
			})
		}
	}
	return natGateways, nil
}
