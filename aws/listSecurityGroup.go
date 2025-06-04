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
	if c.accountID != "" && param.AccountId != "" && c.accountID != param.AccountId {
		panic(fmt.Sprintf("ListSecurityGroups called with different AccountID: %s, expected: %s", param.AccountId, c.accountID))
	}

	operationalAccountID := param.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListSecurityGroups: param.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListSecurityGroups: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Infof("[AccountID: %s] ListSecurityGroups called for VPC %s, Region %s", operationalAccountID, param.GetVpcId(), param.GetRegion())

	// REMOVED: c.creds = param.Creds
	// REMOVED: c.accountID = param.AccountId

	builder := newFilterBuilder()
	builder.withVPC(param.GetVpcId())
	for k, v := range param.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if param.GetRegion() == "" || param.GetRegion() == "all" {
		var (
			wg                sync.WaitGroup
			allSecurityGroups []types.SecurityGroup
			allErrors         []error
			resultChannel     = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListSecurityGroups: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListSecurityGroups: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions {
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListSecurityGroups: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getSecurityGroupsForRegion
				sgs, err := c.getSecurityGroupsForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					sgs:    sgs,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListSecurityGroups: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListSecurityGroups: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allSecurityGroups = append(allSecurityGroups, result.sgs...)
			}
		}
		c.logger.Infof("[AccountID: %s] Found %d security groups across %d regions", operationalAccountID, len(allSecurityGroups), len(regions))

		if len(allErrors) > 0 {
			return allSecurityGroups, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allSecurityGroups, nil
	}
	c.logger.Debugf("[AccountID: %s] ListSecurityGroups: Processing specific region: %s", operationalAccountID, param.Region)
	// Pass operationalAccountID to getSecurityGroupsForRegion
	return c.getSecurityGroupsForRegion(ctx, param.Region, filters, operationalAccountID)
}

func (c *Client) getSecurityGroupsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter, operationalAccountID string) ([]types.SecurityGroup, error) {
	c.logger.Debugf("[AccountID: %s] getSecurityGroupsForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getSecurityGroupsForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Call DescribeSecurityGroups operation
	resp, err := client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: filters,
	})
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getSecurityGroupsForRegion: DescribeSecurityGroups failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Pass operationalAccountID to convertSecurityGroups
	return convertSecurityGroups(c.defaultRegion, regionName, operationalAccountID, resp.SecurityGroups), nil
}

func convertSecurityGroups(defaultRegion, region, accountID string, awsSGs []awsTypes.SecurityGroup) []types.SecurityGroup {
	if region == "" {
		region = defaultRegion
	}

	out := make([]types.SecurityGroup, 0, len(awsSGs))
	for _, sg := range awsSGs {
		sgLink := fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#SecurityGroup:groupId=%s", region, region, aws.ToString(sg.GroupId))
		out = append(out, types.SecurityGroup{
			Name:      aws.ToString(sg.GroupName),
			ID:        aws.ToString(sg.GroupId),
			Provider:  providerName,
			VpcID:     aws.ToString(sg.VpcId),
			Region:    region,
			AccountID: accountID, // Use the passed operationalAccountID
			Labels:    convertTags(sg.Tags), // Assuming convertTags uses aws.ToString or is safe
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
			ipRanges = append(ipRanges, aws.ToString(r.CidrIp))
		}
		for _, r := range v.Ipv6Ranges {
			ipRanges = append(ipRanges, aws.ToString(r.CidrIpv6))
		}
		// Handle UserIdGroupPairs (source/destination SGs)
		for _, r := range v.UserIdGroupPairs {
			sourceSG := aws.ToString(r.GroupId)
			if aws.ToString(r.UserId) != "" { // Check if UserId is present (cross-account)
				sourceSG = fmt.Sprintf("%s/%s", aws.ToString(r.UserId), sourceSG)
			}
			ipRanges = append(ipRanges, sourceSG)
		}

		portRange := "all" // Default for "all" traffic or if FromPort is nil
		if v.FromPort != nil {
			portRange = fmt.Sprintf("%d", *v.FromPort) // Dereference since FromPort is *int32
			if v.ToPort != nil && *v.ToPort != *v.FromPort { // Only append ToPort if it's different
				portRange += fmt.Sprintf("-%d", *v.ToPort) // Dereference since ToPort is *int32
			}
		}

		protocol := aws.ToString(v.IpProtocol)
		if protocol == "-1" { // AWS SDK uses "-1" for "all" protocols
			protocol = "all"
		}
		if protocol == "" { // If IpProtocol is nil or empty string, it implies "all" for some rules (e.g. SG source)
			protocol = "all"
		}

		rules = append(rules, types.SecurityGroupRule{
			Protocol:  protocol,
			PortRange: portRange,
			Source:    ipRanges, // This now includes CIDRs and SGs
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
