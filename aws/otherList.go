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
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListACLs(ctx context.Context, input *infrapb.ListACLsRequest) ([]types.ACL, error) {
	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId())
	getParam := &ec2.DescribeNetworkAclsInput{
		Filters: builder.build(),
	}

	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allAcls       []types.ACL
			resultChannel = make(chan []types.ACL)
			errorChannel  = make(chan error)
		)

		regionalClients, err := c.getAllClientsForProfile(input.GetAccountId())
		if err != nil {
			return nil, err
		}
		for regionName, awsRegionClient := range regionalClients {
			wg.Add(1)
			go func(regionName string, awsRegionClient awsClient) {
				defer wg.Done()

				out, err := awsRegionClient.ec2Client.DescribeNetworkAcls(ctx, getParam)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS ACLs: %v", err)
					return
				}

				acls := convertACLs(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), regionName, out.NetworkAcls)
				resultChannel <- acls
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for acls := range resultChannel {
			allAcls = append(allAcls, acls...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return allAcls, nil
	}

	client, err := c.getEC2Client(ctx, input.GetAccountId(), input.GetRegion())
	if err != nil {
		return nil, err
	}
	out, err := client.DescribeNetworkAcls(ctx, getParam)
	if err != nil {
		return nil, err
	}
	return convertACLs(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), input.GetRegion(), out.NetworkAcls), nil
}

func convertACLs(defaultAccount, defaultRegion, account, region string, awsACLs []awsTypes.NetworkAcl) []types.ACL {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	out := make([]types.ACL, 0, len(awsACLs))
	for _, acl := range awsACLs {
		rules := make([]types.ACLRule, 0, len(acl.Entries))
		for _, r := range acl.Entries {

			rule := types.ACLRule{
				Number:            0,
				Protocol:          convertString(r.Protocol),
				PortRange:         "",
				SourceRanges:      nil,
				DestinationRanges: nil,
				Action:            string(r.RuleAction),
				Direction:         "",
			}
			if r.RuleNumber != nil {
				rule.Number = int(*r.RuleNumber)
			}
			if r.Egress != nil {
				if *r.Egress == true {
					rule.Direction = "Egress"
				} else {
					rule.Direction = "Ingress"
				}
			}
			if rule.Protocol == "-1" {
				rule.Protocol = "all"
			}
			if r.PortRange != nil {
				if r.PortRange.From != nil {
					rule.PortRange = fmt.Sprintf("%d", r.PortRange.From)
				}
				if r.PortRange.To != nil {
					rule.PortRange += fmt.Sprintf("- %d", r.PortRange.To)
				}
			}

			var cidrs []string
			if r.CidrBlock != nil {
				cidrs = append(cidrs, convertString(r.CidrBlock))
			}
			if r.Ipv6CidrBlock != nil {
				cidrs = append(cidrs, convertString(r.Ipv6CidrBlock))
			}
			if rule.Direction == "Egress" {
				rule.DestinationRanges = cidrs
			}
			if rule.Direction == "Ingress" {
				rule.SourceRanges = cidrs
			}

			rules = append(rules, rule)
		}
		aclLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#NetworkAclDetails:networkAclId=%s", region, region, aws.ToString(acl.NetworkAclId))
		out = append(out, types.ACL{
			Name:      convertString(getTagName(acl.Tags)),
			ID:        convertString(acl.NetworkAclId),
			Provider:  providerName,
			VpcID:     convertString(acl.VpcId),
			Region:    region,
			AccountID: account,
			Labels:    convertTags(acl.Tags),
			Rules:     rules,
			SelfLink:  aclLink,
		})
	}
	return out
}

func (c *Client) ListRouteTables(ctx context.Context, input *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error) {
	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId())
	getParam := &ec2.DescribeRouteTablesInput{
		Filters: builder.build(),
	}

	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			routeTables   []types.RouteTable
			resultChannel = make(chan []types.RouteTable)
			errorChannel  = make(chan error)
		)

		regionalClients, err := c.getAllClientsForProfile(input.GetAccountId())
		if err != nil {
			return nil, err
		}
		for regionName, awsRegionClient := range regionalClients {
			wg.Add(1)
			go func(regionName string, awsRegionClient awsClient) {
				defer wg.Done()

				out, err := awsRegionClient.ec2Client.DescribeRouteTables(ctx, getParam)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS RouteTables: %v", err)
					return
				}

				routeTs := convertRouteTables(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), regionName, out.RouteTables)
				resultChannel <- routeTs
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for rts := range resultChannel {
			routeTables = append(routeTables, rts...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return routeTables, nil
	}

	client, err := c.getEC2Client(ctx, input.GetAccountId(), input.GetRegion())
	if err != nil {
		return nil, err
	}
	out, err := client.DescribeRouteTables(ctx, getParam)
	if err != nil {
		return nil, err
	}
	return convertRouteTables(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), input.GetRegion(), out.RouteTables), nil
}

func convertRouteTables(defaultAccount, defaultRegion, account, region string, awsRts []awsTypes.RouteTable) []types.RouteTable {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	out := make([]types.RouteTable, 0, len(awsRts))
	for _, rt := range awsRts {
		routes := make([]types.Route, 0, len(rt.Routes))
		rtLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#RouteTableDetails:routeTableId=%s", region, region, aws.ToString(rt.RouteTableId))

		for _, r := range rt.Routes {
			var destination string
			if r.DestinationCidrBlock != nil {
				destination = convertString(r.DestinationCidrBlock)
			} else if r.DestinationIpv6CidrBlock != nil {
				destination = convertString(r.DestinationIpv6CidrBlock)
			}

			target := ""
			if r.CarrierGatewayId != nil {
				target = convertString(r.CarrierGatewayId)
			} else if r.CoreNetworkArn != nil {
				target = convertString(r.CoreNetworkArn)
			} else if r.EgressOnlyInternetGatewayId != nil {
				target = convertString(r.EgressOnlyInternetGatewayId)
			} else if r.GatewayId != nil {
				target = convertString(r.GatewayId)
			} else if r.InstanceId != nil {
				target = convertString(r.InstanceId)
			} else if r.InstanceOwnerId != nil {
				target = convertString(r.InstanceOwnerId)
			} else if r.LocalGatewayId != nil {
				target = convertString(r.LocalGatewayId)
			} else if r.NatGatewayId != nil {
				target = convertString(r.NatGatewayId)
			} else if r.NetworkInterfaceId != nil {
				target = convertString(r.NetworkInterfaceId)
			} else if r.TransitGatewayId != nil {
				target = convertString(r.TransitGatewayId)
			} else if r.VpcPeeringConnectionId != nil {
				target = convertString(r.VpcPeeringConnectionId)
			}

			routes = append(routes, types.Route{
				Destination: destination,
				Status:      string(r.State),
				Target:      target,
			})
		}
		out = append(out, types.RouteTable{
			Name:      convertString(getTagName(rt.Tags)),
			ID:        convertString(rt.RouteTableId),
			Provider:  providerName,
			VpcID:     convertString(rt.VpcId),
			Region:    region,
			AccountID: account,
			Labels:    convertTags(rt.Tags),
			Routes:    routes,
			SelfLink:  rtLink,
		})
	}
	return out
}
