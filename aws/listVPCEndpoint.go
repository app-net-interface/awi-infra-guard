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
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId
	builder := newFilterBuilder()
	for k, v := range params.Labels {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allVPCEs      []types.VPCEndpoint
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
				vpces, err := c.getVPCEsForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					vpces:  vpces,
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
				allVPCEs = append(allVPCEs, result.vpces...)
			}
		}

		if len(allErrors) > 0 {
			return allVPCEs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allVPCEs, nil
	}
	return c.getVPCEsForRegion(ctx, params.Region, nil)
}

func (c *Client) getVPCEsForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.VPCEndpoint, error) {
	var veps []types.VPCEndpoint

	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, vep := range page.VpcEndpoints {
			var name, state, serviceName, subnetIds, routeTableIds string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range vep.Tags {
				if *tag.Key == "Name" || *tag.Key == "name" {
					name = *tag.Value
				}
				labels[*tag.Key] = *tag.Value
			}
			if vep.ServiceName != nil {
				serviceName = *vep.ServiceName
			}

			//var subnetIds, routeTableIds []string
			subnetIds = strings.Join(vep.SubnetIds, ",")
			routeTableIds = strings.Join(vep.RouteTableIds, ",")

			veps = append(veps, types.VPCEndpoint{
				ID:            *vep.VpcEndpointId,
				Provider:      c.GetName(),
				AccountID:     *vep.OwnerId,
				Name:          name,
				VPCId:         *vep.VpcId,
				Region:        regionName,
				State:         state,
				Labels:        labels,
				ServiceName:   serviceName,
				SubnetIds:     subnetIds,
				RouteTableIds: routeTableIds,
				CreatedAt:     vep.CreationTimestamp,
				SelfLink:      fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#EndpointDetails:vpcEndpointId=%s", regionName, regionName, *vep.VpcEndpointId),
			})
		}
	}

	return veps, nil
}
