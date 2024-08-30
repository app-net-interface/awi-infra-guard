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
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbTypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

func (c *Client) ListLBs(ctx context.Context, input *infrapb.ListLBsRequest) ([]types.LB, error) {
	c.creds = input.Creds
	c.accountID = input.AccountId

	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId())
	for k, v := range input.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			allLBs        []types.LB
			allErrors     []error
			wg            sync.WaitGroup
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
				regLBs, err := c.getLBsForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					lbs:    regLBs,
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
				allLBs = append(allLBs, result.lbs...)
			}
		}

		c.logger.Infof("In account %s Found %d LBs across %d regions", c.accountID, len(allLBs), len(regions))

		if len(allErrors) > 0 {
			return allLBs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		PrintResources(allLBs, "types.LB")

		return allLBs, nil
	}

	return c.getLBsForRegion(ctx, input.Region, filters)
}

func (c *Client) getLBsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter) ([]types.LB, error) {
	var lbs []types.LB

	// Get ELBv2 Load Balancers (ALB, NLB)
	elbv2Client, err := c.getELBv2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	if elbv2Client == nil {
		return nil, fmt.Errorf("ELBv2 client is nil")
	}

	elbv2Paginator := elbv2.NewDescribeLoadBalancersPaginator(elbv2Client, &elbv2.DescribeLoadBalancersInput{})
	if elbv2Paginator == nil {
		return nil, fmt.Errorf("failed to create DescribeLoadBalancersPaginator")
	}
	for elbv2Paginator.HasMorePages() {

		page, err := elbv2Paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("error describing ELBv2 load balancers: %v", err)
		}

		for _, lb := range page.LoadBalancers {
			lbs = append(lbs, c.convertELBv2ToLoadBalancer(ctx, lb, regionName))
		}
	}

	// Get Classic ELB Load Balancers
	elbClient, err := c.getELBClient(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}

	elbInput := &elb.DescribeLoadBalancersInput{}
	elbPaginator := elb.NewDescribeLoadBalancersPaginator(elbClient, elbInput)

	for elbPaginator.HasMorePages() {
		page, err := elbPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("error describing classic ELB load balancers: %v", err)
		}

		for _, lb := range page.LoadBalancerDescriptions {
			lbs = append(lbs, c.convertClassicELBToLoadBalancer(lb, regionName))
		}
	}

	// Apply filters
	filteredLBs := c.applyFilters(lbs, filters)

	return filteredLBs, nil
}

func (c *Client) applyFilters(lbs []types.LB, filters []awsTypes.Filter) []types.LB {
	// Implement filter logic here
	// This is a placeholder and should be implemented based on your specific requirements
	return lbs
}

func (c *Client) convertELBv2ToLoadBalancer(ctx context.Context, lb elbv2types.LoadBalancer, regionName string) types.LB {
	var lbType string
	var ipAddressType string
	switch lb.Type {
	case elbv2types.LoadBalancerTypeEnumApplication:
		lbType = "ALB"
	case elbv2types.LoadBalancerTypeEnumNetwork:
		lbType = "NLB"
	case elbv2types.LoadBalancerTypeEnumGateway:
		lbType = "GWLB"
	default:
		lbType = "Unknown"
	}
	switch lb.IpAddressType {
	case elbv2types.IpAddressTypeIpv4:
		ipAddressType = "ipv4"
	case elbv2types.IpAddressTypeDualstack:
		ipAddressType = "dualstack"
	case elbv2types.IpAddressTypeDualstackWithoutPublicIpv4:
		ipAddressType = "ipv6"
	default:
		ipAddressType = "unknown"
	}
	ips, err := getIPsV2(lb)
	if err != nil {
		c.logger.Errorf("Error resolving load balancer DNS: %v", err)
	}

	listeners, err := c.getListenersV2(ctx, aws.ToString(lb.LoadBalancerArn), regionName)
	if err != nil {
		c.logger.Errorf("Error getting listeners for load balancer %s: %v", aws.ToString(lb.LoadBalancerArn), err)
	}

	return types.LB{
		ID:            aws.ToString(lb.LoadBalancerArn),
		Provider:      c.GetName(),
		Name:          aws.ToString(lb.LoadBalancerName),
		Scheme:        string(lb.Scheme),
		DNSName:       aws.ToString(lb.DNSName),
		Type:          lbType,
		IPAddressType: ipAddressType,
		IPAddresses:   ips,
		Listeners:     listeners,
		//State:         string(lb.State.Code),
		VPCID:     aws.ToString(lb.VpcId),
		AccountID: c.accountID,
		Region:    regionName,
		CreatedAt: aws.ToTime(lb.CreatedTime),
		Labels:    c.getTagsV2(lb, regionName),
		SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/v2/home?region=%s#LoadBalancers:search=%s", regionName, regionName, aws.ToString(lb.LoadBalancerName)),
	}
}

func (c *Client) getListenersV2(ctx context.Context, loadBalancerArn, regionName string) ([]types.LBListener, error) {
	elbv2Client, err := c.getELBv2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}

	input := &elbv2.DescribeListenersInput{
		LoadBalancerArn: aws.String(loadBalancerArn),
	}

	paginator := elbv2.NewDescribeListenersPaginator(elbv2Client, input)
	var listeners []types.LBListener

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, listener := range output.Listeners {
			var port int32
			if listener.Port != nil {
				port = *listener.Port
			}
			listeners = append(listeners, types.LBListener{
				ListenerID:    aws.ToString(listener.ListenerArn),
				Protocol:      string(listener.Protocol),
				Port:          port,
				TargetGroupID: getDefaultTargetGroupArn(listener),
			})
		}
	}

	return listeners, nil
}

func getDefaultTargetGroupArn(listener elbv2types.Listener) string {
	for _, action := range listener.DefaultActions {
		if action.Type == elbv2types.ActionTypeEnumForward && action.TargetGroupArn != nil {
			return aws.ToString(action.TargetGroupArn)
		}
	}
	return ""
}

func (c *Client) convertClassicELBToLoadBalancer(lb elbTypes.LoadBalancerDescription, regionName string) types.LB {
	var instanceIDs []string
	for _, instance := range lb.Instances {
		instanceIDs = append(instanceIDs, *instance.InstanceId)
	}
	ips, err := getIPsV1(lb)
	if err != nil {
		c.logger.Errorf("Error resolving load balancer DNS: %v", err)
	}
	return types.LB{
		ID:            aws.ToString(lb.LoadBalancerName),
		Name:          aws.ToString(lb.LoadBalancerName),
		DNSName:       aws.ToString(lb.DNSName),
		Provider:      c.GetName(),
		IPAddressType: "ipv4",
		Type:          "Classic",
		Scheme:        aws.ToString(lb.Scheme),
		VPCID:         aws.ToString(lb.VPCId),
		AccountID:     c.accountID,
		Region:        regionName,
		CreatedAt:     aws.ToTime(lb.CreatedTime),
		SelfLink:      fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#LoadBalancerDetails:loadBalancerId=%s", regionName, regionName, aws.ToString(lb.LoadBalancerName)),
		InstanceIDs:   instanceIDs,
		IPAddresses:   ips,
		Zone:          getZone(lb),
		Labels:        c.getTagsV1(lb, regionName),
		Listeners:     convertListeners(lb.ListenerDescriptions),
	}
}


func getIPsV1(lb elbTypes.LoadBalancerDescription) ([]string, error) {
	if lb.DNSName != nil {
		return getIPsFromDNS(*lb.DNSName)
	} else {
		return nil, fmt.Errorf("DNSName is nil")
	}
}

func getIPsV2(lb elbv2types.LoadBalancer) ([]string, error) {
	if lb.DNSName != nil {
		return getIPsFromDNS(*lb.DNSName)
	} else {
		return nil, fmt.Errorf("DNSName is nil")
	}
}

func getZone(lb elbTypes.LoadBalancerDescription) string {
	if len(lb.AvailabilityZones) > 0 {
		return lb.AvailabilityZones[0]
	}
	return ""
}

func (c *Client) getTagsV2(lb elbv2types.LoadBalancer, regionName string) map[string]string {
	// Classic ELBs don't include tags in the DescribeLoadBalancers call
	// We need to make a separate DescribeTags API call to get this information
	input := &elbv2.DescribeTagsInput{
		ResourceArns: []string{aws.ToString(lb.LoadBalancerName)},
	}

	elbv2Client, err := c.getELBv2Client(context.TODO(), c.accountID, regionName)
	if err != nil {
		c.logger.Errorf("Error getting ELBv2 client: %v", err)
		return nil
	}
	result, err := elbv2Client.DescribeTags(context.TODO(), input)
	if err != nil {
		// Handle error, perhaps log it
		return nil
	}

	tags := make(map[string]string)
	if len(result.TagDescriptions) > 0 {
		for _, tag := range result.TagDescriptions[0].Tags {
			tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
		}
	}

	return tags
}

func (c *Client) getTagsV1(lb elbTypes.LoadBalancerDescription, regionName string) map[string]string {
	// Classic ELBs don't include tags in the DescribeLoadBalancers call
	// We need to make a separate DescribeTags API call to get this information
	input := &elb.DescribeTagsInput{
		LoadBalancerNames: []string{aws.ToString(lb.LoadBalancerName)},
	}

	elbClient, err := c.getELBClient(context.TODO(), c.accountID, regionName)
	if err != nil {
		c.logger.Errorf("Error getting ELB client: %v", err)
		return nil
	}
	result, err := elbClient.DescribeTags(context.TODO(), input)
	if err != nil {
		// Handle error, perhaps log it
		return nil
	}

	tags := make(map[string]string)
	if len(result.TagDescriptions) > 0 {
		for _, tag := range result.TagDescriptions[0].Tags {
			tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
		}
	}

	return tags
}

func convertListeners(listeners []elbTypes.ListenerDescription) []types.LBListener {
	var result []types.LBListener
	for _, l := range listeners {
		result = append(result, types.LBListener{
			ListenerID: "",
			Protocol:   aws.ToString(l.Listener.Protocol),
			Port:       l.Listener.LoadBalancerPort,
			// Classic ELBs don't have TargetGroups
		})
	}
	return result
}
