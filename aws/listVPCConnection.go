// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) createPeeringConnection(pc awstypes.VpcPeeringConnection, region string) types.VPCConnection {
	name := ""
	for _, tag := range pc.Tags {
		if aws.ToString(tag.Key) == "Name" {
			name = aws.ToString(tag.Value)
			break
		}
	}

	return types.VPCConnection{
		Provider:         providerName,
		ID:               aws.ToString(pc.VpcPeeringConnectionId),
		Name:             name,
		Account:          aws.ToString(pc.AccepterVpcInfo.OwnerId),
		Region:           region,
		FromVpcId:        aws.ToString(pc.RequesterVpcInfo.VpcId),
		ToVpcId:          aws.ToString(pc.AccepterVpcInfo.VpcId),
		FromVpcAccountId: aws.ToString(pc.RequesterVpcInfo.OwnerId),
		ToVpcAccountId:   aws.ToString(pc.AccepterVpcInfo.OwnerId),
		FromVpcRegion:    aws.ToString(pc.RequesterVpcInfo.Region),
		ToVpcRegion:      aws.ToString(pc.AccepterVpcInfo.Region),
		Status:           string(pc.Status.Code),
		ConnectionType:   "peering",
	}
}

func (c *Client) ListVpcConnections(ctx context.Context, input *infrapb.ListVpcConnectionsRequest) ([]types.VPCConnection, error) {
	var allVpcConns []types.VPCConnection

	//c.logger.Debugf("Starting VPC connection enumeration for account %s", input.AccountId)
	//c.logger.Debugf("Client configuration - Default Account: %s, Default Region: %s", c.defaultAccountID, c.defaultRegion)

	// Get all available AWS regions
	regions, err := c.getAllRegions(ctx, input.AccountId)
	if err != nil {
		c.logger.Errorf("Failed to get AWS regions: %v", err)
		return nil, err
	}

	//c.logger.Debugf("Processing %d AWS regions", len(regions))
	for _, region := range regions {
		regionName := aws.ToString(region.RegionName)
		//c.logger.Debugf("Processing region %s", regionName)

		// Get EC2 client for this region
		client, err := c.getEC2Client(ctx, input.AccountId, regionName)
		if err != nil {
			c.logger.Warnf("Failed to get EC2 client for region %s: %v", regionName, err)
			continue
		}

		// Process VPC peering connections
		//c.logger.Debugf("Calling DescribeVpcPeeringConnections for account %s in region %s", input.AccountId, regionName)
		peeringConnections, err := client.DescribeVpcPeeringConnections(ctx, &ec2.DescribeVpcPeeringConnectionsInput{})
		if err != nil {
			c.logger.Warnf("Failed to list VPC peering connections in region %s: %v", regionName, err)
		} else {
			//c.logger.Debugf("Found %d VPC peering connections in region %s", len(peeringConnections.VpcPeeringConnections), regionName)
			for _, pc := range peeringConnections.VpcPeeringConnections {
				//c.logger.Debugf("Processing VPC peering connection: ID=%s, RequesterVPC=%s, AccepterVPC=%s, Status=%s",
				//aws.ToString(pc.VpcPeeringConnectionId),
				//aws.ToString(pc.RequesterVpcInfo.VpcId),
				//aws.ToString(pc.AccepterVpcInfo.VpcId),
				//string(pc.Status.Code))
				allVpcConns = append(allVpcConns, c.createPeeringConnection(pc, regionName))
			}
		}

		// Process Transit Gateway VPC attachments
		tgwAttachments, err := client.DescribeTransitGatewayAttachments(ctx, &ec2.DescribeTransitGatewayAttachmentsInput{})
		if err != nil {
			c.logger.Warnf("Failed to list transit gateway attachments in region %s: %v", regionName, err)
			continue
		}

		// Group attachments by transit gateway
		tgwAttachmentMap := make(map[string][]awstypes.TransitGatewayAttachment)
		for _, att := range tgwAttachments.TransitGatewayAttachments {
			if att.ResourceType == awstypes.TransitGatewayAttachmentResourceTypeVpc {
				tgwID := aws.ToString(att.TransitGatewayId)
				tgwAttachmentMap[tgwID] = append(tgwAttachmentMap[tgwID], att)
			}
		}

		// Get Transit Gateway route tables
		tgwRouteTables, err := client.DescribeTransitGatewayRouteTables(ctx, &ec2.DescribeTransitGatewayRouteTablesInput{})
		if err != nil {
			c.logger.Warnf("Failed to list transit gateway route tables in region %s: %v", regionName, err)
			continue
		}

		// Build a map of VPC attachments by TGW for quick lookup
		attachmentsByTGW := make(map[string][]awstypes.TransitGatewayAttachment)
		for _, att := range tgwAttachments.TransitGatewayAttachments {
			if att.ResourceType == awstypes.TransitGatewayAttachmentResourceTypeVpc {
				tgwID := aws.ToString(att.TransitGatewayId)
				attachmentsByTGW[tgwID] = append(attachmentsByTGW[tgwID], att)
			}
		}

		// For each TGW, check route tables to find actual VPC-to-VPC connectivity
		for tgwID, attachments := range attachmentsByTGW {
			//c.logger.Debugf("Processing TGW %s with %d VPC attachments", tgwID, len(attachments))

			// Get route tables for this TGW
			for _, rt := range tgwRouteTables.TransitGatewayRouteTables {
				if aws.ToString(rt.TransitGatewayId) != tgwID {
					continue
				}

				// Get routes in this route table
				routes, err := client.SearchTransitGatewayRoutes(ctx, &ec2.SearchTransitGatewayRoutesInput{
					TransitGatewayRouteTableId: rt.TransitGatewayRouteTableId,
					Filters: []awstypes.Filter{
						{
							Name:   aws.String("type"),
							Values: []string{"static", "propagated"},
						},
					},
				})
				if err != nil {
					c.logger.Warnf("Failed to get routes for TGW route table %s: %v",
						aws.ToString(rt.TransitGatewayRouteTableId), err)
					continue
				}

				// In a Transit Gateway, attachments in the same route table can communicate if:
				// 1. They are both active attachments
				// 2. There are routes in the route table (which we've already fetched)

				activeAttachments := make(map[string]awstypes.TransitGatewayAttachment)
				for _, attachment := range attachments {
					if attachment.State == awstypes.TransitGatewayAttachmentStateAvailable {
						attID := aws.ToString(attachment.TransitGatewayAttachmentId)
						activeAttachments[attID] = attachment
					}
				}

				//c.logger.Debugf("Route table %s has %d active attachments",
				//	aws.ToString(rt.TransitGatewayRouteTableId),
				//	len(activeAttachments))

				// If there are routes in this table, all active attachments can potentially communicate
				if len(routes.Routes) > 0 {
					// Create connections between all active attachments
					for _, att1 := range activeAttachments {
						att1ID := aws.ToString(att1.TransitGatewayAttachmentId)

						for _, att2 := range activeAttachments {
							att2ID := aws.ToString(att2.TransitGatewayAttachmentId)

							// Skip self-connections
							if att1ID == att2ID {
								continue
							}

							// Skip if we've already created this connection (in reverse)
							if att1ID > att2ID {
								continue
							}

							//c.logger.Debugf("Creating TGW connection between attachments: %s and %s", att1ID, att2ID)
							allVpcConns = append(allVpcConns, types.VPCConnection{
								Provider:         providerName,
								ID:               fmt.Sprintf("%s:%s:%s", tgwID, att1ID, att2ID),
								Name:             fmt.Sprintf("tgw-%s-route", tgwID),
								Account:          aws.ToString(att1.ResourceOwnerId),
								Region:           regionName,
								FromVpcId:        aws.ToString(att1.ResourceId),
								ToVpcId:          aws.ToString(att2.ResourceId),
								FromVpcAccountId: aws.ToString(att1.ResourceOwnerId),
								ToVpcAccountId:   aws.ToString(att2.ResourceOwnerId),
								FromVpcRegion:    regionName,
								ToVpcRegion:      regionName,
								Status:           "inactive",
								ConnectionType:   "transit_gateway",
							})
							//c.logger.Debugf("Found valid routes between VPCs: %s and %s", aws.ToString(att1.ResourceId), aws.ToString(att2.ResourceId))
							allVpcConns = append(allVpcConns, types.VPCConnection{
								Provider:         providerName,
								ID:               fmt.Sprintf("%s-%s-%s", tgwID, att1ID, att2ID),
								Name:             fmt.Sprintf("tgw-%s-route", tgwID),
								Account:          aws.ToString(att1.ResourceOwnerId),
								Region:           regionName,
								FromVpcId:        aws.ToString(att1.ResourceId),
								ToVpcId:          aws.ToString(att2.ResourceId),
								FromVpcAccountId: aws.ToString(att1.ResourceOwnerId),
								ToVpcAccountId:   aws.ToString(att2.ResourceOwnerId),
								FromVpcRegion:    regionName,
								ToVpcRegion:      regionName,
								Status:           "active", // If routes exist, connection is active
								ConnectionType:   "transit_gateway",
							})
						}
					}
				}
			}
		}
	}

	//c.logger.Debugf("Found %d VPC connections across all regions for account %s", len(allVpcConns), input.AccountId)
	return allVpcConns, nil
}
