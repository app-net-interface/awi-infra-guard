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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/app-net-interface/awi-infra-guard/connector/db"
	t "github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/sirupsen/logrus"
)

const transitGatewayName = "AWI-DEV-TGW"
const transitGatewayRouteTableName = "AWI-TGW-ROUTE-TABLE"
const awiRouteTagPrefix = "AWI-ROUTE_"

const awiRouteTableTagKey = "transit_route"

var awiRouteTableTagValues = []string{"true", "True", "TRUE", "yes", "Yes", "YES", "1",
	"ok", "OK", "Ok", "Success", "success", "SUCCESS"}

var NotFoundError error

func (c *Client) ConnectVPCs(ctx context.Context, input t.VPCConnectionParams) (t.VPCConnectionOutput, error) {
	var err error
	region1 := input.Region1
	if region1 == "" {
		region1, err = c.findVPCRegion(ctx, input.Vpc1ID)
		if err != nil {
			return t.VPCConnectionOutput{}, err
		}
	}
	region2 := input.Region2
	if region2 == "" {
		region2, err = c.findVPCRegion(ctx, input.Vpc2ID)
		if err != nil {
			return t.VPCConnectionOutput{}, err
		}
	}

	timeout := 5 * time.Minute
	// Create a timer channel to handle the timeout
	timer := time.NewTimer(timeout)
	// Perform the operation in a loop
	for {
		select {
		case <-timer.C:
			// Timeout reached
			msg := fmt.Sprintf("failed to finish connect VPCs operation within timeout")
			c.logger.Errorf(msg)
			return t.VPCConnectionOutput{}, fmt.Errorf(msg)
		default:
			// Try the operation
			err := c.connectVPCsOp(ctx, input.ConnID, input.Vpc1ID, input.Vpc2ID, "", region1, region2)
			if err != nil {
				if !strings.Contains(err.Error(), "is in invalid state") {
					c.logger.Errorf("Failed to connect VPCs: %v", err)
					return t.VPCConnectionOutput{}, err
				}
				c.logger.Infof("Connect VPCs failed because of: %v, retrying...", err)
			} else {
				return t.VPCConnectionOutput{
					Region1: region1,
					Region2: region2,
				}, nil
			}
		}
	}
}

func (c *Client) ConnectVPC(ctx context.Context, input t.SingleVPCConnectionParams) (t.SingleVPCConnectionOutput, error) {
	var err error
	if input.Region == "" {
		input.Region, err = c.findVPCRegion(ctx, input.VpcID)
		if err != nil {
			return t.SingleVPCConnectionOutput{}, err
		}
	}

	timeout := 5 * time.Minute
	// Create a timer channel to handle the timeout
	timer := time.NewTimer(timeout)
	// Perform the operation in a loop
	for {
		select {
		case <-timer.C:
			// Timeout reached
			msg := fmt.Sprintf("failed to finish connect VPC operation within timeout")
			c.logger.Errorf(msg)
			return t.SingleVPCConnectionOutput{}, fmt.Errorf(msg)
		default:
			// Try the operation
			err := c.connectOneSideVPCOp(ctx, input, "")
			if err != nil {
				if !strings.Contains(err.Error(), "is in invalid state") {
					c.logger.Errorf("Failed to connect VPC: %v", err)
					return t.SingleVPCConnectionOutput{}, err
				}
				c.logger.Infof("Connect VPC failed because of: %v, retrying...", err)
			} else {
				return t.SingleVPCConnectionOutput{
					Region: input.Region,
				}, nil
			}
		}
	}
}

func closeCSPDB(client *db.Client, logger *logrus.Logger) {
	if err := client.Close(); err != nil {
		logger.Warnf("failed to close CSP Connection DB: %v", err)
	}
}

// Returns true if the existing CSP Connection matches the requested
// VPCs that are about to get created.
func (c *Client) connectionMatches(connection db.Connection, input t.SingleVPCConnectionParams) bool {
	sourceProvider := strings.ToLower(connection.SourceProvider)
	destProvider := strings.ToLower(connection.DestinationProvider)
	if connection.State != db.StateActive {
		return false
	}
	if sourceProvider != "aws" && destProvider != "aws" {
		return false
	}
	if sourceProvider == "aws" && destProvider == "aws" {
		// We do not expect the connection between the same provider
		return false
	}
	if sourceProvider == "aws" {
		// TODO: Currently, source ID is the ID of Transit Gateway
		// and not a name as expected here. it needs to be calculated
		// properly.
		// if connection.SourceID != transitGatewayName {
		// 	return false
		// }
		if connection.SourceRegion != input.Region {
			return false
		}
		// TODO: Calculation of Destination region is not supported yet.
		// The ConnectVPC method needs to get information about second
		// side of connection.
		// if connection.DestinationRegion != input.Destination.Region {
		// 	return false
		// }

		// TODO: VPC Name, ID and URL for GCP are mixed
		// if connection.DestinationVPC != input.Destination.VPC {
		// 	return false
		// }
		if destProvider != strings.ToLower(input.Destination.Provider) {
			return false
		}
		return true
	}
	// TODO: Currently, destination ID is the ID of Transit Gateway
	// and not a name as expected here. it needs to be calculated
	// properly.
	if connection.DestinationID != transitGatewayName {
		return false
	}
	// TODO: Calculation of Destination region is not supported yet.
	// The ConnectVPC method needs to get information about second
	// side of connection.
	// if connection.DestinationRegion != input.Region {
	// 	return false
	// }
	if connection.SourceRegion != input.Destination.Region {
		return false
	}
	// TODO: VPC Name, ID and URL for GCP are mixed
	// if connection.SourceVPC != input.Destination.VPC {
	// 	return false
	// }
	if sourceProvider != strings.ToLower(input.Destination.Provider) {
		return false
	}
	return true
}

func (c *Client) getCIDRsFromVPCOfTheSecondProvider(ctx context.Context, input t.SingleVPCConnectionParams) ([]string, error) {
	client, err := db.NewClient(db.DefaultDBFile, c.logger.WithField("logger", "cloud-connection-db"))
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the Client for DB storing CSP Connections: %w", err)
	}
	defer closeCSPDB(client, c.logger)

	connections, err := client.ListConnections()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the Client for DB storing CSP Connections: %w", err)
	}

	for _, conn := range connections {
		if !c.connectionMatches(conn, input) {
			continue
		}
		if strings.ToLower(conn.SourceProvider) == "aws" {
			return conn.DestinationCIDRs, nil
		}
		return conn.SourceCIDRs, nil
	}

	return nil, fmt.Errorf(
		"could not find a connection between Transit Gateway %s and Destination VPC: %v",
		transitGatewayName, input.Destination,
	)
}

func (c *Client) connectOneSideVPCOp(ctx context.Context, input t.SingleVPCConnectionParams, account string) error {
	// 1. create transit gateway
	tgwID, err := c.createTransitGateway(ctx, account, input.Region)
	if err != nil {
		return err
	}
	// 2. create transit gateway attachment
	tgwAttachmentId, _, err := c.createTransitGatewayAttachment(ctx, account, input.Region, tgwID, input.VpcID)
	if err != nil {
		return err
	}
	// 3. create transit gateway routes
	tgwRT, err := c.createTransitGatewayRoutesTables(ctx, account, input.Region, input.VpcID, tgwID)
	if err != nil {
		return err
	}
	// 4. associate RouteTable with transit gateway attachments
	err = c.associateRouteTable(ctx, account, input.Region, tgwRT, tgwAttachmentId)
	if err != nil {
		return err
	}
	cidrs, err := c.getCIDRsFromVPCOfTheSecondProvider(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to obtain CIDR for the AWS Transit Gateway: %w", err)
	}
	if len(cidrs) == 0 {
		return errors.New("Got empty CIDR list from the second provider. Cannot create proper Route in Transit Gateway")
	}
	// TODO: Handle creating routes and adding proper tags when there are multiple CIDRs
	if len(cidrs) > 1 {
		return errors.New("NOT IMPLEMENTED - Gotta handle more CIDRs than 1")
	}
	err = c.attachRoutesToVPC(ctx, account, input.Region, input.VpcID, tgwID, &cidrs[0])
	if err != nil {
		return err
	}
	err = c.addCIDRsTagToVPCRouteTable(ctx, input.ConnID, account, input.Region, input.VpcID, &cidrs[0])
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) connectVPCsOp(ctx context.Context, connID string, vpc1ID, vpc2ID string, account, region1, region2 string) error {
	if region1 != region2 {
		return fmt.Errorf("transit gateway connection can only be used for single region connections")
	}
	region := region1
	// 1. create transit gateway
	tgwID, err := c.createTransitGateway(ctx, account, region)
	if err != nil {
		return err
	}
	// 2. create transit gateway attachments
	tgwAttachment1Id, vpc1CIDR, err := c.createTransitGatewayAttachment(ctx, account, region, tgwID, vpc1ID)
	if err != nil {
		return err
	}
	tgwAttachment2Id, vpc2CIDR, err := c.createTransitGatewayAttachment(ctx, account, region, tgwID, vpc2ID)
	if err != nil {
		return err
	}
	// 3. create transit gateway routes
	tgwRT1, err := c.createTransitGatewayRoutesTables(ctx, account, region, vpc1ID, tgwID)
	if err != nil {
		return err
	}
	tgwRT2, err := c.createTransitGatewayRoutesTables(ctx, account, region, vpc2ID, tgwID)
	if err != nil {
		return err
	}
	// 4. associate RouteTable with transit gateway attachments
	err = c.associateRouteTable(ctx, account, region, tgwRT1, tgwAttachment1Id)
	if err != nil {
		return err
	}
	err = c.associateRouteTable(ctx, account, region, tgwRT2, tgwAttachment2Id)
	if err != nil {
		return err
	}
	// 5. add routes to route tables
	err = c.createTransitGatewayRoutes(ctx, account, region, tgwRT1, tgwAttachment2Id, vpc2CIDR)
	if err != nil {
		return err
	}
	err = c.createTransitGatewayRoutes(ctx, account, region, tgwRT2, tgwAttachment1Id, vpc1CIDR)
	if err != nil {
		return err
	}
	// 6. attach routes to VPC
	err = c.attachRoutesToVPC(ctx, account, region, vpc1ID, tgwID, vpc2CIDR)
	if err != nil {
		return err
	}
	err = c.attachRoutesToVPC(ctx, account, region, vpc2ID, tgwID, vpc1CIDR)
	if err != nil {
		return err
	}
	err = c.addCIDRsTagToVPCRouteTable(ctx, connID, account, region, vpc1ID, vpc2CIDR)
	if err != nil {
		return err
	}
	err = c.addCIDRsTagToVPCRouteTable(ctx, connID, account, region, vpc2ID, vpc1CIDR)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) attachRoutesToVPC(ctx context.Context, account, region string, vpcID string, tgwID *string,
	cidr *string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withTagManyValues(awiRouteTableTagKey, awiRouteTableTagValues)

	in := &ec2.DescribeRouteTablesInput{
		Filters: builder.build(),
	}

	res, err := client.DescribeRouteTables(ctx, in)
	if err != nil {
		return err
	}

	if len(res.RouteTables) == 0 {
		return fmt.Errorf(
			"no route table with matching tag with key: %s, possible values: %s found for vpc %s",
			awiRouteTableTagKey, awiRouteTableTagValues, vpcID)
	}

	for _, rt := range res.RouteTables {
		input := &ec2.CreateRouteInput{
			DestinationCidrBlock: cidr,
			GatewayId:            tgwID,
			RouteTableId:         rt.RouteTableId,
		}

		_, err = client.CreateRoute(ctx, input)
		if err != nil && !strings.Contains(err.Error(), "RouteAlreadyExists") {
			return err
		}
	}

	return nil
}

func (c *Client) addCIDRsTagToVPCRouteTable(ctx context.Context, connID, account, region string,
	vpcID string, cidr *string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withAssociationMain()

	in := &ec2.DescribeRouteTablesInput{
		Filters: builder.build(),
	}

	res, err := client.DescribeRouteTables(ctx, in)
	if err != nil {
		return err
	}

	if len(res.RouteTables) == 0 {
		return fmt.Errorf("main route table not found for vpc %s", vpcID)
	}
	input := &ec2.CreateTagsInput{
		Resources: []string{aws.ToString(res.RouteTables[0].RouteTableId)},
		Tags: []types.Tag{{
			Key:   aws.String(routeTableConnTag(connID)),
			Value: cidr,
		},
		},
	}

	_, err = client.CreateTags(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) disconnectVPCs(ctx context.Context, connID string, vpc1ID, vpc2ID string, account, region1, region2 string) error {
	var err error
	if region1 == "" {
		region1, err = c.findVPCRegion(ctx, vpc1ID)
		if err != nil {
			return err
		}
	}
	if region2 == "" {
		region2, err = c.findVPCRegion(ctx, vpc2ID)
		if err != nil {
			return err
		}
	}

	err = c.findCIDRsTagAndRemoveRoutes(ctx, connID, account, region1, vpc1ID)
	if err != nil {
		return err
	}
	err = c.findCIDRsTagAndRemoveRoutes(ctx, connID, account, region2, vpc2ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) disconnectOneSideVPC(ctx context.Context, connID, vpcID, account, region string) error {
	// TODO: Handle removing leftover routes.
	return nil
}

func (c *Client) findCIDRsTagAndRemoveRoutes(ctx context.Context, connID, account, region, vpcID string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withAssociationMain()

	in := &ec2.DescribeRouteTablesInput{
		Filters: builder.build(),
	}

	res, err := client.DescribeRouteTables(ctx, in)
	if err != nil {
		return err
	}

	if len(res.RouteTables) == 0 {
		return fmt.Errorf("main route table not found for vpc %s", vpcID)
	}
	var cidr string
	key := routeTableConnTag(connID)
	routeTable := res.RouteTables[0]
	for _, tag := range routeTable.Tags {
		if aws.ToString(tag.Key) == key {
			cidr = aws.ToString(tag.Value)
			break
		}
	}

	input := &ec2.DeleteRouteInput{
		DestinationCidrBlock: &cidr,
		RouteTableId:         res.RouteTables[0].RouteTableId,
	}

	_, err = client.DeleteRoute(ctx, input)
	if err != nil && !strings.Contains(err.Error(), "no route with destination-cidr-block") {
		return err
	}

	inputTag := &ec2.DeleteTagsInput{
		Resources: []string{aws.ToString(routeTable.RouteTableId)},
		Tags: []types.Tag{{
			Key: aws.String(routeTableConnTag(connID)),
		},
		},
	}

	err = c.deleteTransitGatewayRoutes(ctx, account, region, vpcID, cidr)
	if err != nil {
		return err
	}
	_, err = client.DeleteTags(ctx, inputTag)
	if err != nil {
		return err
	}
	c.logger.Infof("Removed routes, vpc: %s, connection ID: %s.", vpcID, connID)
	return nil
}

func (c *Client) createTransitGatewayRoutesTables(ctx context.Context, account, region, vpcID string, transitGatewayID *string) (*string, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}
	name := getTransitGatewayRouteTableName(vpcID)

	input := &ec2.DescribeTransitGatewayRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{name},
			},
		},
	}

	// Describe the Transit Gateways RT with the specified name
	output, err := client.DescribeTransitGatewayRouteTables(ctx, input)
	if err != nil {
		return nil, err
	}

	// Check if any Transit Gateways RT were found
	if len(output.TransitGatewayRouteTables) > 0 {
		c.logger.Infof("AWI Transit Gateway Route Table already exists")
		return output.TransitGatewayRouteTables[0].TransitGatewayRouteTableId, nil
	}

	inputRT := &ec2.CreateTransitGatewayRouteTableInput{
		TagSpecifications: []types.TagSpecification{
			mapToTagSpecfication(map[string]string{"Name": name}, types.ResourceTypeTransitGatewayRouteTable),
		},
		TransitGatewayId: transitGatewayID,
	}

	out, err := client.CreateTransitGatewayRouteTable(ctx, inputRT)
	if err != nil {
		c.logger.Errorf("Error creating transit gateway route table: %v", err)
		return nil, err
	}
	c.logger.Infof("Successfully created AWI Transit Gateway Route Table")
	return out.TransitGatewayRouteTable.TransitGatewayRouteTableId, nil
}

func (c *Client) deleteTransitGatewayRoutes(ctx context.Context, account, region string, vpcID string, cidr string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	name := getTransitGatewayRouteTableName(vpcID)

	input := &ec2.DescribeTransitGatewayRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{name},
			},
		},
	}

	// Describe the Transit Gateways RT with the specified name
	output, err := client.DescribeTransitGatewayRouteTables(ctx, input)
	if err != nil {
		return err
	}

	// Check if any Transit Gateways RT were found
	if len(output.TransitGatewayRouteTables) == 0 {
		c.logger.Warnf("TransitGatway route table %s not found", name)
		return nil
	}

	id := output.TransitGatewayRouteTables[0].TransitGatewayRouteTableId

	_, err = client.DeleteTransitGatewayRoute(ctx, &ec2.DeleteTransitGatewayRouteInput{
		TransitGatewayRouteTableId: id,
		DestinationCidrBlock:       &cidr,
	})
	if err != nil && !strings.Contains(err.Error(), "no route with destination-cidr-block") {
		return err
	}

	return nil
}

func routeTableConnTag(connID string) string {
	return awiRouteTagPrefix + connID
}

func getTransitGatewayRouteTableName(vpcID string) string {
	return transitGatewayRouteTableName + " " + vpcID
}

func (c *Client) findVPCRegion(ctx context.Context, vpcID string) (string, error) {
	regionalClients, err := c.getAllClientsForProfile("")
	if err != nil {
		return "", err
	}
	input := &ec2.DescribeVpcsInput{VpcIds: []string{vpcID}}
	for regionName, awsRegionClient := range regionalClients {
		vpcs, err := awsRegionClient.ec2Client.DescribeVpcs(ctx, input)
		if err != nil || len(vpcs.Vpcs) == 0 {
			continue
		}
		c.logger.Infof("Found region of VPC %s: %s", vpcID, regionName)
		return regionName, nil
	}

	return "", fmt.Errorf("failed to find region of vpc %s", vpcID)
}

func (c *Client) DisconnectVPCs(ctx context.Context, input t.VPCDisconnectionParams) (t.VPCDisconnectionOutput, error) {
	c.logger.Infof("Discconnecting %s and %s, connection ID: %s", input.Vpc1ID, input.Vpc2ID, input.ConnID)
	err := c.disconnectVPCs(ctx, input.ConnID, input.Vpc1ID, input.Vpc2ID, "", input.Region1, input.Region2)
	return t.VPCDisconnectionOutput{}, err
}

func (c *Client) DisconnectVPC(ctx context.Context, input t.SingleVPCDisconnectionParams) (t.VPCDisconnectionOutput, error) {
	c.logger.Infof("Discconnecting %s, connection ID: %s", input.VpcID, input.ConnID)
	err := c.disconnectOneSideVPC(ctx, input.ConnID, input.VpcID, "", input.Region)
	return t.VPCDisconnectionOutput{}, err
}
