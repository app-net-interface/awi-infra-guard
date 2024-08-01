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

func (c *Client) ListSecurityGroups(ctx context.Context, param *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error) {
	c.logger.Infof("List instances")
	c.creds = param.Creds
	c.accountID = param.AccountId
	builder := newFilterBuilder()
	builder.withVPC(param.GetVpcId())

	filters := builder.build()

	if param.GetRegion() == "" || param.GetRegion() == "all" {
		var (
			wg                sync.WaitGroup
			allSecurityGroups []types.SecurityGroup
			allErrors         []error
			resultChannel     = make(chan regionResult)
		)

		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}
		for _, region := range regions {
			wg.Add(1)
			go func(region string) {
				defer wg.Done()

				sgs, err := c.getSecurityGroupsForRegion(ctx, region, filters)
				resultChannel <- regionResult{
					region: region,
					sgs:    sgs,
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
				allSecurityGroups = append(allSecurityGroups, result.sgs...)
			}
		}
		c.logger.Infof("In account %s Found %d security groups across %d regions", c.accountID, len(allSecurityGroups), len(regions))

		if len(allErrors) > 0 {
			return allSecurityGroups, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allSecurityGroups, nil
	}
	return c.getSecurityGroupsForRegion(ctx, param.Region, filters)
}

func (c *Client) getSecurityGroupsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter) ([]types.SecurityGroup, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeVpcs operation
	resp, err := client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertSecurityGroups(c.defaultRegion, regionName, resp.SecurityGroups), nil
}

func convertSecurityGroups(defaultRegion, region string, awsSGs []awsTypes.SecurityGroup) []types.SecurityGroup {
	if region == "" {
		region = defaultRegion
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
			AccountID: *sg.OwnerId,
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
