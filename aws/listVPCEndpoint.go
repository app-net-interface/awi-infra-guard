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
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListVPCEndpoints called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListVPCEndpoints: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListVPCEndpoints: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Infof("[AccountID: %s] ListVPCEndpoints called. Region: %s", operationalAccountID, params.GetRegion())

	builder := newFilterBuilder()
	for k, v := range params.Labels {
		builder.withTag(k, v)
	}
	// Add VPC ID filter if provided in params, as VPCEndpoints are VPC-specific
	if params.GetVpcId() != "" {
		builder.withVPC(params.GetVpcId())
	}
	filters := builder.build()

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allVPCEs      []types.VPCEndpoint
			allErrors     []error
			resultChannel = make(chan regionResult)
		)
		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListVPCEndpoints: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListVPCEndpoints: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions {
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListVPCEndpoints: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getVPCEsForRegion
				vpces, err := c.getVPCEsForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					vpces:  vpces,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListVPCEndpoints: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListVPCEndpoints: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allVPCEs = append(allVPCEs, result.vpces...)
			}
		}
		c.logger.Infof("[AccountID: %s] Found %d VPC Endpoints across %d regions", operationalAccountID, len(allVPCEs), len(regions))

		if len(allErrors) > 0 {
			return allVPCEs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allVPCEs, nil
	}
	c.logger.Debugf("[AccountID: %s] ListVPCEndpoints: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getVPCEsForRegion
	return c.getVPCEsForRegion(ctx, params.Region, filters, operationalAccountID)
}

func (c *Client) getVPCEsForRegion(ctx context.Context, regionName string, filters []awstypes.Filter, operationalAccountID string) ([]types.VPCEndpoint, error) {
	c.logger.Debugf("[AccountID: %s] getVPCEsForRegion: Region %s", operationalAccountID, regionName)
	var veps []types.VPCEndpoint

	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getVPCEsForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{
		Filters: filters,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx) // Use ctx from parameters
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getVPCEsForRegion: DescribeVpcEndpoints NextPage failed for region %s: %v", operationalAccountID, regionName, err)
			return nil, err
		}
		for _, vep := range page.VpcEndpoints {
			var name, state, vepType, serviceName string
			labels := make(map[string]string)

			switch vep.VpcEndpointType {
			case awstypes.VpcEndpointTypeGateway:
				vepType = "Gateway"
			case awstypes.VpcEndpointTypeInterface:
				vepType = "Interface"
			case awstypes.VpcEndpointTypeGatewayLoadBalancer:
				vepType = "GatewayLoadbalancer"
			default:
				vepType = string(vep.VpcEndpointType) // Handle unknown types
			}

			// Extracting Name from Tags
			for _, tag := range vep.Tags {
				labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
				if aws.ToString(tag.Key) == "Name" || aws.ToString(tag.Key) == "name" {
					name = aws.ToString(tag.Value)
				}
			}
			serviceName = aws.ToString(vep.ServiceName)
			state = string(vep.State)

			// Extract Security Group IDs
			var securityGroupIDs []string
			for _, group := range vep.Groups {
				securityGroupIDs = append(securityGroupIDs, aws.ToString(group.GroupId))
			}

			veps = append(veps, types.VPCEndpoint{
				ID:               aws.ToString(vep.VpcEndpointId),
				Provider:         c.GetName(),
				AccountID:        aws.ToString(vep.OwnerId), // Owner of the resource
				Name:             name,
				VPCId:            aws.ToString(vep.VpcId),
				Region:           regionName,
				State:            state,
				Labels:           labels,
				Type:             vepType,
				ServiceName:      serviceName,
				SubnetIds:        vep.SubnetIds,      // Assign the slice directly
				RouteTableIds:    vep.RouteTableIds,    // Assign the slice directly
				SecurityGroupIDs: securityGroupIDs,   // Assign the extracted slice
				CreatedAt:        vep.CreationTimestamp,
				SelfLink:         fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#EndpointDetails:vpcEndpointId=%s", regionName, regionName, aws.ToString(vep.VpcEndpointId)),
			})
		}
	}

	p2 := ec2.NewDescribeInstanceConnectEndpointsPaginator(client, &ec2.DescribeInstanceConnectEndpointsInput{
		// Filters for InstanceConnectEndpoints if needed, though the request object doesn't have specific ones for these.
		// MaxResults: nil,
		// NextToken: nil,
	})

	for p2.HasMorePages() {
		page, err := p2.NextPage(ctx) // Use ctx from parameters
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getVPCEsForRegion: DescribeInstanceConnectEndpoints NextPage failed for region %s: %v", operationalAccountID, regionName, err)
			return nil, err
		}
		for _, vep := range page.InstanceConnectEndpoints {
			var name string
			labels := make(map[string]string)

			// Extracting Name from Tags
			for _, tag := range vep.Tags {
				labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
				if aws.ToString(tag.Key) == "Name" || aws.ToString(tag.Key) == "name" {
					name = aws.ToString(tag.Value)
				}
			}

			// Instance Connect Endpoints are a type of VPC Endpoint.
			// We can represent them using the same types.VPCEndpoint struct,
			// filling in fields as appropriate.
			veps = append(veps, types.VPCEndpoint{
				Provider:  c.GetName(),
				ID:        aws.ToString(vep.InstanceConnectEndpointId),
				AccountID: aws.ToString(vep.OwnerId), // Owner of the resource
				Name:      name,
				VPCId:     aws.ToString(vep.VpcId),
				Region:    regionName,
				State:     string(vep.State), // Assuming vep.State is of type awstypes.Ec2InstanceStateName or similar
				Labels:    labels,
				Type:      "InstanceConnect", // Specific type for these endpoints
				SubnetIds: []string{aws.ToString(vep.SubnetId)}, // ICE has one subnet
				CreatedAt: vep.CreatedAt,
				SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#InstanceConnectEndpointDetails:instanceConnectEndpointId=%s", regionName, regionName, aws.ToString(vep.InstanceConnectEndpointId)),
				// Other fields like ServiceName, RouteTableIds, SecurityGroupIDs might not be directly applicable or available
				// for Instance Connect Endpoints in the same way as for standard VPC Endpoints.
				// SecurityGroupIDs: vep.SecurityGroupIds, // If available and needed
			})
		}
	}
	c.logger.Debugf("[AccountID: %s] getVPCEsForRegion: Found %d VPC Endpoints (including Instance Connect Endpoints) in region %s.", operationalAccountID, len(veps), regionName)
	return veps, nil
}
