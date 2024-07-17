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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListSecurityGroups(ctx context.Context, input *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error) {
	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId())
	getParam := &ec2.DescribeSecurityGroupsInput{
		Filters: builder.build(),
	}
	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allSgs        []types.SecurityGroup
			resultChannel = make(chan []types.SecurityGroup)
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

				out, err := awsRegionClient.ec2Client.DescribeSecurityGroups(ctx, getParam)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS security groups: %v", err)
					return
				}

				securityGroups := convertSecurityGroups(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), regionName, out.SecurityGroups)
				resultChannel <- securityGroups
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for sgs := range resultChannel {
			allSgs = append(allSgs, sgs...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return allSgs, nil
	}
	client, err := c.getEC2Client(ctx, input.GetAccountId(), input.GetRegion())
	if err != nil {
		return nil, err
	}
	out, err := client.DescribeSecurityGroups(ctx, getParam)
	if err != nil {
		return nil, err
	}
	return convertSecurityGroups(c.defaultAccountID, c.defaultRegion, input.GetAccountId(), input.GetRegion(), out.SecurityGroups), nil
}

func convertSecurityGroups(defaultAccount, defaultRegion, account, region string, awsSGs []awsTypes.SecurityGroup) []types.SecurityGroup {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	out := make([]types.SecurityGroup, 0, len(awsSGs))
	for _, sg := range awsSGs {
		sgLink := fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#SecurityGroup:groupId=%s", region, region, aws.ToString(sg.GroupId))
		out = append(out, types.SecurityGroup{
			Name:      convertString(sg.GroupName),
			ID:        convertString(sg.GroupId),
			Provider:  providerName,
			VpcID:     convertString(sg.VpcId),
			Region:    region,
			AccountID: account,
			Labels:    convertTags(sg.Tags),
			Rules:     convertSecurityGroupRules(sg.IpPermissions, sg.IpPermissionsEgress),
			SelfLink:  sgLink,
		})
	}
	return out
}

func convertSecurityGroupRules(ingress []awsTypes.IpPermission, egress []awsTypes.IpPermission) []types.SecurityGroupRule {
	rules := make([]types.SecurityGroupRule, 0, len(ingress)+len(egress))

	f := func(direction string, v awsTypes.IpPermission) {
		var ipRanges []string
		for _, r := range v.IpRanges {
			ipRanges = append(ipRanges, convertString(r.CidrIp))
		}
		for _, r := range v.Ipv6Ranges {
			ipRanges = append(ipRanges, convertString(r.CidrIpv6))
		}
		portRange := "all"
		if v.FromPort != nil {
			portRange = fmt.Sprintf("%d", v.FromPort)
		}
		if v.ToPort != nil {
			portRange += fmt.Sprintf("- %d", v.ToPort)
		}
		protocol := convertString(v.IpProtocol)
		if protocol == "-1" {
			protocol = "all"
		}
		rules = append(rules, types.SecurityGroupRule{
			Protocol:  protocol,
			PortRange: portRange,
			Source:    ipRanges,
			Direction: direction,
		})
	}
	for _, v := range ingress {
		f("Ingress", v)
	}
	for _, v := range egress {
		f("Egress", v)
	}
	return rules
}
