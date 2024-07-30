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

// Transit Gateway related functions

package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// Create a new transit gateway if we don't find one in the region.
// TBD : What if we find multiple transit gateways ??
// We should use the one that has AWI tag, and if there are multiple transit gateway with AWI tag,
// let's use the first one in the list
func (c *Client) createTransitGateway(ctx context.Context, account, region string) (*string, error) {
	c.logger.Infof("Creating a new TGW in account %s and for region %s", account, region)
	tgwID, err := c.findTransitGatewayByName(ctx, transitGatewayName, account, region)
	if err != nil && err != NotFoundError {
		return nil, err
	}
	if tgwID != nil {
		return tgwID, nil
	}
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}
	createTgwInput := &ec2.CreateTransitGatewayInput{
		TagSpecifications: []types.TagSpecification{
			mapToTagSpecfication(map[string]string{"Name": transitGatewayName}, types.ResourceTypeTransitGateway),
		},
		Description: aws.String("AWI Transit Gateway"),
	}

	tgwOutput, err := client.CreateTransitGateway(ctx, createTgwInput)
	if err != nil {
		return nil, err
	}

	return tgwOutput.TransitGateway.TransitGatewayId, nil
}

// TBD: Delete the transit gateway with id .
// func (c *Client) deleteTransitGateway(ctx context.Context, account, TransitGatewayId string) (e error) {
// 	return e
// }

func (c *Client) findTransitGatewayByName(ctx context.Context, transitGatewayName string, account, region string) (*string, error) {
	// Describe Transit Gateways with filters to find by name
	input := &ec2.DescribeTransitGatewaysInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{transitGatewayName},
			},
		},
	}

	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}

	// Describe the Transit Gateways with the specified name
	output, err := client.DescribeTransitGateways(ctx, input)
	if err != nil {
		return nil, err
	}

	// Check if any Transit Gateways were found
	if len(output.TransitGateways) == 0 {
		return nil, NotFoundError
	}
	for _, v := range output.TransitGateways {
		if v.State != types.TransitGatewayStateDeleted && v.State != types.TransitGatewayStateDeleting {
			return v.TransitGatewayId, nil
		}
	}
	return nil, NotFoundError
}

func (c *Client) createTransitGatewayAttachment(ctx context.Context, account, region string,
	transitGatewayID *string, vpcID string) (*string, *string, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, nil, err
	}

	vpcInput := &ec2.DescribeVpcsInput{VpcIds: []string{vpcID}}

	vpcs, err := client.DescribeVpcs(ctx, vpcInput)
	if err != nil {
		return nil, nil, err
	}
	if len(vpcs.Vpcs) == 0 {
		return nil, nil, fmt.Errorf("couldn't find VPC %s", vpcID)
	}

	vpcCidr := vpcs.Vpcs[0].CidrBlock

	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	inputSubnet := &ec2.DescribeSubnetsInput{
		Filters: builder.build(),
	}
	subnets, err := client.DescribeSubnets(ctx, inputSubnet)
	if err != nil {
		return nil, nil, err
	}
	var subnetIDs []string
	for _, v := range subnets.Subnets {
		subnetIDs = append(subnetIDs, aws.ToString(v.SubnetId))
	}

	name := fmt.Sprintf("AWI attachment VPC %v", vpcID)
	input := &ec2.DescribeTransitGatewayAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{name},
			},
		},
	}
	out, err := client.DescribeTransitGatewayAttachments(ctx, input)
	if err != nil {
		return nil, nil, err
	}
	if len(out.TransitGatewayAttachments) > 0 {
		for _, attachment := range out.TransitGatewayAttachments {
			if attachment.State != types.TransitGatewayAttachmentStateDeleted &&
				attachment.State != types.TransitGatewayAttachmentStateDeleting &&
				attachment.State != types.TransitGatewayAttachmentStateFailed {
				return attachment.TransitGatewayAttachmentId, vpcCidr, nil
			}
		}
	}
	attachVpcInput := &ec2.CreateTransitGatewayVpcAttachmentInput{
		TagSpecifications: []types.TagSpecification{
			mapToTagSpecfication(map[string]string{"Name": name}, types.ResourceTypeTransitGatewayAttachment),
		},
		TransitGatewayId: transitGatewayID,
		VpcId:            aws.String(vpcID),
		SubnetIds:        subnetIDs,
	}

	outCr, err := client.CreateTransitGatewayVpcAttachment(ctx, attachVpcInput)
	if err != nil && !strings.Contains(err.Error(), "DuplicateTransitGatewayAttachment") {
		c.logger.Errorf("Error attaching VPC to Transit Gateway: %v", err)
		return nil, nil, err
	}

	c.logger.Infof("VPC %s attached to Transit Gateway %s.", vpcID, aws.ToString(transitGatewayID))
	return outCr.TransitGatewayVpcAttachment.TransitGatewayAttachmentId, vpcCidr, nil
}

func (c *Client) associateRouteTable(ctx context.Context, account, region string, routeTableID, tgwAttachmentID *string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}

	associateRouteTableInput := &ec2.AssociateTransitGatewayRouteTableInput{
		TransitGatewayAttachmentId: tgwAttachmentID,
		TransitGatewayRouteTableId: routeTableID,
	}

	_, err = client.AssociateTransitGatewayRouteTable(ctx, associateRouteTableInput)
	if err != nil && !strings.Contains(err.Error(), "Resource.AlreadyAssociated") {
		return err
	}

	return nil
}

func (c *Client) createTransitGatewayRoutes(ctx context.Context, account, region string,
	routeTableId, toTgwAttachmentID *string, cidr *string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}

	_, err = client.CreateTransitGatewayRoute(ctx, &ec2.CreateTransitGatewayRouteInput{
		TransitGatewayRouteTableId: routeTableId,
		DestinationCidrBlock:       cidr,
		TransitGatewayAttachmentId: toTgwAttachmentID,
	})
	if err != nil && !strings.Contains(err.Error(), "RouteAlreadyExists") {
		return err
	}

	return nil
}
