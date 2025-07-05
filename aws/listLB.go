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
	"net" // Import the net package
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types" // Assuming this is for EC2 filters, not directly used for LB listing here
	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbTypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

func (c *Client) ListLBs(ctx context.Context, input *infrapb.ListLBsRequest) ([]types.LB, error) {
	if c.accountID != "" && input.AccountId != "" && c.accountID != input.AccountId {
		panic(fmt.Sprintf("ListLBs called with different AccountID: %s, expected: %s", input.AccountId, c.accountID))
	}

	operationalAccountID := input.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListLBs: input.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListLBs: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListLBs called with VPC ID: %s, Region: %s", operationalAccountID, input.GetVpcId(), input.GetRegion())

	// REMOVED: c.creds = input.Creds
	// REMOVED: c.accountID = input.AccountId

	// Note: AWS Load Balancers (ALB/NLB/Classic) are not directly filtered by EC2-style filters (awsTypes.Filter)
	// in their Describe API calls. Filtering is typically done post-fetch based on attributes like VPC ID or tags.
	// The `filters` variable created by `newFilterBuilder` might not be directly usable with ELB/ELBv2 Describe calls.
	// We will keep it for now in case `applyFilters` uses it, but be aware of this.
	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId()) // This will be used in applyFilters if implemented for VPC
	for k, v := range input.GetLabels() {
		builder.withTag(k, v) // This will be used in applyFilters if implemented for Tags
	}
	filtersForApply := builder.build() // Renamed to clarify its use

	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			allLBs        []types.LB
			allErrors     []error
			wg            sync.WaitGroup
			resultChannel = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListLBs: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListLBs: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region (from EC2 DescribeRegions)
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListLBs: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getLBsForRegion
				// Pass filtersForApply for post-fetch filtering
				regLBs, err := c.getLBsForRegion(ctx, regionName, accID, filtersForApply)
				resultChannel <- regionResult{
					region: regionName,
					lbs:    regLBs,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListLBs: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListLBs: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allLBs = append(allLBs, result.lbs...)
			}
		}

		c.logger.Infof("[AccountID: %s] ListLBs: Found %d LBs across %d regions", operationalAccountID, len(allLBs), len(regions))

		if len(allErrors) > 0 {
			return allLBs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		//PrintResources(allLBs, "types.LB")

		return allLBs, nil
	}
	c.logger.Debugf("[AccountID: %s] ListLBs: Processing specific region: %s", operationalAccountID, input.Region)
	// Pass operationalAccountID to getLBsForRegion
	return c.getLBsForRegion(ctx, input.Region, operationalAccountID, filtersForApply)
}

// getLBsForRegion now accepts operationalAccountID and filtersForApply
func (c *Client) getLBsForRegion(ctx context.Context, regionName string, operationalAccountID string, filtersForApply []awsTypes.Filter) ([]types.LB, error) {
	c.logger.Debugf("[AccountID: %s] getLBsForRegion: Region %s", operationalAccountID, regionName)
	var lbs []types.LB

	// Get ELBv2 Load Balancers (ALB, NLB, GWLB)
	// Use operationalAccountID for getting the ELBv2 client
	elbv2Client, err := c.getELBv2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getLBsForRegion: Failed to get ELBv2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// It's possible for a client to be nil if the service is not supported or enabled, handle gracefully.
	if elbv2Client != nil {
		elbv2Paginator := elbv2.NewDescribeLoadBalancersPaginator(elbv2Client, &elbv2.DescribeLoadBalancersInput{})
		for elbv2Paginator.HasMorePages() {
			page, err := elbv2Paginator.NextPage(ctx)
			if err != nil {
				c.logger.Errorf("[AccountID: %s] getLBsForRegion: DescribeLoadBalancers (ELBv2) NextPage failed for region %s: %v", operationalAccountID, regionName, err)
				return nil, fmt.Errorf("error describing ELBv2 load balancers in region %s for account %s: %v", regionName, operationalAccountID, err)
			}
			for _, lb := range page.LoadBalancers {
				// Pass operationalAccountID to converter
				lbs = append(lbs, c.convertELBv2ToLoadBalancer(ctx, lb, regionName, operationalAccountID))
			}
		}
	} else {
		c.logger.Warnf("[AccountID: %s] getLBsForRegion: ELBv2 client is nil for region %s. Skipping ELBv2 LBs.", operationalAccountID, regionName)
	}

	// Get Classic ELB Load Balancers
	// Use operationalAccountID for getting the ELB client
	elbClient, err := c.getELBClient(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getLBsForRegion: Failed to get ELB client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	if elbClient != nil {
		elbInput := &elb.DescribeLoadBalancersInput{}
		elbPaginator := elb.NewDescribeLoadBalancersPaginator(elbClient, elbInput)
		for elbPaginator.HasMorePages() {
			page, err := elbPaginator.NextPage(ctx)
			if err != nil {
				c.logger.Errorf("[AccountID: %s] getLBsForRegion: DescribeLoadBalancers (Classic ELB) NextPage failed for region %s: %v", operationalAccountID, regionName, err)
				return nil, fmt.Errorf("error describing classic ELB load balancers in region %s for account %s: %v", regionName, operationalAccountID, err)
			}
			for _, lb := range page.LoadBalancerDescriptions {
				// Pass operationalAccountID to converter
				lbs = append(lbs, c.convertClassicELBToLoadBalancer(ctx, lb, regionName, operationalAccountID))
			}
		}
	} else {
		c.logger.Warnf("[AccountID: %s] getLBsForRegion: Classic ELB client is nil for region %s. Skipping Classic LBs.", operationalAccountID, regionName)
	}

	// Apply filters (post-fetch)
	// The filtersForApply are EC2-style filters. applyFilters needs to interpret them for LB properties.
	filteredLBs := c.applyFilters(lbs, filtersForApply, operationalAccountID)
	c.logger.Debugf("[AccountID: %s] getLBsForRegion: Region %s, found %d LBs, filtered to %d LBs.", operationalAccountID, regionName, len(lbs), len(filteredLBs))
	return filteredLBs, nil
}

// applyFilters now accepts operationalAccountID for logging
func (c *Client) applyFilters(lbs []types.LB, ec2Filters []awsTypes.Filter, operationalAccountID string) []types.LB {
	if len(ec2Filters) == 0 {
		return lbs // No filters to apply
	}
	c.logger.Debugf("[AccountID: %s] applyFilters: Applying %d EC2-style filters to %d LBs.", operationalAccountID, len(ec2Filters), len(lbs))

	var filteredLBs []types.LB
	for _, lb := range lbs {
		matchesAll := true
		for _, filter := range ec2Filters {
			filterName := aws.ToString(filter.Name)
			filterValues := make(map[string]bool)
			for _, val := range filter.Values {
				filterValues[val] = true
			}

			match := false
			switch filterName {
			case "vpc-id":
				if _, ok := filterValues[lb.VPCID]; ok {
					match = true
				}
			// Add cases for tag filters, e.g., "tag:YourTagName"
			// Example for a specific tag:
			// case "tag:Name":
			// 	if lbName, ok := lb.Labels["Name"]; ok {
			// 		if _, filterOk := filterValues[lbName]; filterOk {
			// 			match = true
			// 		}
			// 	}
			// More generic tag handling:
			default:
				if len(filterName) > 4 && filterName[:4] == "tag:" {
					tagName := filterName[4:]
					if lbTagValue, ok := lb.Labels[tagName]; ok {
						if _, filterOk := filterValues[lbTagValue]; filterOk {
							match = true
						}
					}
				} else {
					// If the filter isn't recognized for LBs, consider it a non-match or log a warning.
					// For now, we'll assume if a filter isn't explicitly handled, it doesn't match.
					c.logger.Warnf("[AccountID: %s] applyFilters: Unhandled filter name '%s' for LBs.", operationalAccountID, filterName)
				}
			}
			if !match {
				matchesAll = false
				break
			}
		}
		if matchesAll {
			filteredLBs = append(filteredLBs, lb)
		}
	}
	c.logger.Debugf("[AccountID: %s] applyFilters: Filtered %d LBs down to %d.", operationalAccountID, len(lbs), len(filteredLBs))
	return filteredLBs
}

// convertELBv2ToLoadBalancer now accepts operationalAccountID
func (c *Client) convertELBv2ToLoadBalancer(ctx context.Context, lb elbv2types.LoadBalancer, regionName string, operationalAccountID string) types.LB {
	var lbType string
	var ipAddressType string
	var publicIPs, privateIPs []string

	switch lb.Type {
	case elbv2types.LoadBalancerTypeEnumApplication:
		lbType = "ALB"
	case elbv2types.LoadBalancerTypeEnumNetwork:
		lbType = "NLB"
	case elbv2types.LoadBalancerTypeEnumGateway:
		lbType = "GWLB"
	default:
		lbType = string(lb.Type) // Use the raw string if unknown
		c.logger.Warnf("[AccountID: %s] convertELBv2ToLoadBalancer: Unknown ELBv2 type '%s' for LB %s", operationalAccountID, lb.Type, aws.ToString(lb.LoadBalancerArn))
	}
	switch lb.IpAddressType {
	case elbv2types.IpAddressTypeIpv4:
		ipAddressType = "ipv4"
	case elbv2types.IpAddressTypeDualstack:
		ipAddressType = "dualstack"
	// case elbv2types.IpAddressTypeDualstackWithoutPublicIpv4: // This constant might not exist in all SDK versions or is specific
	// 	ipAddressType = "ipv6" // Or "dualstack-without-public-ipv4"
	default:
		ipAddressType = string(lb.IpAddressType) // Use the raw string if unknown
		c.logger.Warnf("[AccountID: %s] convertELBv2ToLoadBalancer: Unknown IPAddressType '%s' for LB %s", operationalAccountID, lb.IpAddressType, aws.ToString(lb.LoadBalancerArn))

	}

	// Get IPs for all ELBv2 load balancer types by resolving their DNS name.
	// This is a reliable method for ALB, GWLB, and NLB.
	resolvedIPs, err := getIPsV2(lb)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] convertELBv2ToLoadBalancer: Error resolving load balancer DNS for %s: %v", operationalAccountID, aws.ToString(lb.DNSName), err)
	}

	// Classify the resolved IPs as public or private.
	for _, ipStr := range resolvedIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil && ip.IsPrivate() {
			privateIPs = append(privateIPs, ipStr)
		} else if ip != nil {
			publicIPs = append(publicIPs, ipStr)
		}
	}

	// For internet-facing NLBs, AWS might also provide static IPs directly.
	// We can add them to ensure completeness, avoiding duplicates.
	if lb.Type == elbv2types.LoadBalancerTypeEnumNetwork {
		for _, az := range lb.AvailabilityZones {
			for _, addr := range az.LoadBalancerAddresses {
				if addr.PrivateIPv4Address != nil {
					// Avoid duplicates
					found := false
					for _, existingIP := range privateIPs {
						if existingIP == *addr.PrivateIPv4Address {
							found = true
							break
						}
					}
					if !found {
						privateIPs = append(privateIPs, *addr.PrivateIPv4Address)
					}
				}
				if addr.AllocationId != nil {
					publicIP, err := c.getPublicIPFromAllocationID(ctx, *addr.AllocationId, regionName, operationalAccountID)
					if err == nil && publicIP != "" {
						// Avoid duplicates
						found := false
						for _, existingIP := range publicIPs {
							if existingIP == publicIP {
								found = true
								break
							}
						}
						if !found {
							publicIPs = append(publicIPs, publicIP)
						}
					}
				}
			}
		}
	}

	// Pass operationalAccountID to getListenersV2
	listeners, err := c.getListenersV2(ctx, aws.ToString(lb.LoadBalancerArn), regionName, operationalAccountID)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] convertELBv2ToLoadBalancer: Error getting listeners for load balancer %s: %v", operationalAccountID, aws.ToString(lb.LoadBalancerArn), err)
	}

	var subnetIDs []string
	for _, az := range lb.AvailabilityZones {
		if az.SubnetId != nil {
			subnetIDs = append(subnetIDs, *az.SubnetId)
		}
	}

	// Pass operationalAccountID to getTagsV2
	tags := c.getTagsV2(ctx, lb, regionName, operationalAccountID)

	var targetGroupARNs []string
	for _, listener := range listeners {
		if listener.TargetGroupID != "" {
			targetGroupARNs = append(targetGroupARNs, listener.TargetGroupID)
		}
	}
	instanceIDs, err := c.getInstanceIDsForTargetGroups(ctx, targetGroupARNs, regionName, operationalAccountID)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] convertELBv2ToLoadBalancer: Error getting instance IDs for LB %s: %v", operationalAccountID, aws.ToString(lb.LoadBalancerArn), err)
	}

	return types.LB{
		ID:               aws.ToString(lb.LoadBalancerArn),
		Provider:         c.GetName(),
		Name:             aws.ToString(lb.LoadBalancerName),
		Scheme:           string(lb.Scheme),
		DNSName:          aws.ToString(lb.DNSName),
		Type:             lbType,
		PublicIPs:        publicIPs,
		PrivateIPs:       privateIPs,
		IPAddressType:    ipAddressType,
		Listeners:        listeners,
		SecurityGroupIDs: lb.SecurityGroups,
		SubnetIDs:        subnetIDs,
		InstanceIDs:      instanceIDs,
		//State:            string(lb.State.Code),
		VPCID:     aws.ToString(lb.VpcId),
		AccountID: operationalAccountID, // Use operationalAccountID
		Region:    regionName,
		CreatedAt: aws.ToTime(lb.CreatedTime),
		Labels:    tags,
		SelfLink:  fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/v2/home?region=%s#LoadBalancers:search=%s", regionName, regionName, aws.ToString(lb.LoadBalancerName)),
	}
}

// getListenersV2 now accepts operationalAccountID
func (c *Client) getListenersV2(ctx context.Context, loadBalancerArn, regionName, operationalAccountID string) ([]types.LBListener, error) {
	// Use operationalAccountID for getting the ELBv2 client
	elbv2Client, err := c.getELBv2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getListenersV2: Failed to get ELBv2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	if elbv2Client == nil {
		c.logger.Warnf("[AccountID: %s] getListenersV2: ELBv2 client is nil for region %s. Cannot get listeners.", operationalAccountID, regionName)
		return nil, fmt.Errorf("ELBv2 client is nil for region %s, account %s", regionName, operationalAccountID)
	}

	input := &elbv2.DescribeListenersInput{
		LoadBalancerArn: aws.String(loadBalancerArn),
	}

	paginator := elbv2.NewDescribeListenersPaginator(elbv2Client, input)
	var listeners []types.LBListener

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] getListenersV2: DescribeListeners NextPage failed for LB ARN %s, region %s: %v", operationalAccountID, loadBalancerArn, regionName, err)
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
	c.logger.Debugf("[AccountID: %s] getListenersV2: Found %d listeners for LB ARN %s in region %s.", operationalAccountID, len(listeners), loadBalancerArn, regionName)
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

// convertClassicELBToLoadBalancer now accepts operationalAccountID
func (c *Client) convertClassicELBToLoadBalancer(ctx context.Context, lb elbTypes.LoadBalancerDescription, regionName string, operationalAccountID string) types.LB {
	var instanceIDs []string
	for _, instance := range lb.Instances {
		instanceIDs = append(instanceIDs, *instance.InstanceId)
	}
	resolvedIPs, err := getIPsV1(lb)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] convertClassicELBToLoadBalancer: Error resolving load balancer DNS for %s: %v", operationalAccountID, aws.ToString(lb.DNSName), err)
	}

	var publicIPs, privateIPs []string
	for _, ipStr := range resolvedIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil && ip.IsPrivate() {
			privateIPs = append(privateIPs, ipStr)
		} else if ip != nil {
			publicIPs = append(publicIPs, ipStr)
		}
	}

	// Pass operationalAccountID to getTagsV1
	tags := c.getTagsV1(ctx, lb, regionName, operationalAccountID)

	return types.LB{
		ID:            aws.ToString(lb.LoadBalancerName),
		Name:          aws.ToString(lb.LoadBalancerName),
		DNSName:       aws.ToString(lb.DNSName),
		Provider:      c.GetName(),
		IPAddressType: "ipv4",
		Type:          "Classic",
		Scheme:        aws.ToString(lb.Scheme),
		PublicIPs:     publicIPs,
		PrivateIPs:    privateIPs,
		VPCID:         aws.ToString(lb.VPCId),
		AccountID:     operationalAccountID, // Use operationalAccountID
		Region:        regionName,
		CreatedAt:     aws.ToTime(lb.CreatedTime),
		SelfLink:      fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#LoadBalancerDetails:loadBalancerId=%s", regionName, regionName, aws.ToString(lb.LoadBalancerName)),
		InstanceIDs:   instanceIDs,
		Zone:          getZone(lb),
		SubnetIDs:     lb.Subnets,
		Labels:        tags,
		Listeners:     convertListeners(lb.ListenerDescriptions),
	}
}

func getIPsV1(lb elbTypes.LoadBalancerDescription) ([]string, error) {
	if lb.DNSName != nil && *lb.DNSName != "" {
		return getIPsFromDNS(*lb.DNSName)
	}
	// If DNSName is nil or empty, there are no IPs to resolve via DNS.
	// Classic LBs might not always have IPs directly, they rely on DNS.
	return nil, nil // Return nil, nil if no DNS name to resolve
}

func getIPsV2(lb elbv2types.LoadBalancer) ([]string, error) {
	if lb.DNSName != nil && *lb.DNSName != "" {
		return getIPsFromDNS(*lb.DNSName)
	}
	// For NLBs, IPs might be available in lb.AvailabilityZones[].LoadBalancerAddresses
	// However, getIPsFromDNS is a generic approach. If direct IPs are needed for NLBs,
	// this function would need more specific logic for elbv2types.LoadBalancer.
	return nil, nil // Return nil, nil if no DNS name to resolve
}

func getZone(lb elbTypes.LoadBalancerDescription) string {
	if len(lb.AvailabilityZones) > 0 {
		return lb.AvailabilityZones[0] // Classic ELBs can span multiple AZs, this just picks the first.
	}
	return ""
}

// getTagsV2 now accepts context and operationalAccountID
func (c *Client) getTagsV2(ctx context.Context, lb elbv2types.LoadBalancer, regionName string, operationalAccountID string) map[string]string {
	// Use operationalAccountID for getting the ELBv2 client
	elbv2Client, err := c.getELBv2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getTagsV2: Error getting ELBv2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil
	}
	if elbv2Client == nil {
		c.logger.Warnf("[AccountID: %s] getTagsV2: ELBv2 client is nil for region %s. Cannot get tags.", operationalAccountID, regionName)
		return nil
	}

	input := &elbv2.DescribeTagsInput{
		ResourceArns: []string{aws.ToString(lb.LoadBalancerArn)}, // Corrected to use LoadBalancerArn
	}

	result, err := elbv2Client.DescribeTags(ctx, input) // Pass ctx
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getTagsV2: DescribeTags failed for LB ARN %s, region %s: %v", operationalAccountID, aws.ToString(lb.LoadBalancerArn), regionName, err)
		return nil
	}

	tags := make(map[string]string)
	if len(result.TagDescriptions) > 0 {
		for _, tag := range result.TagDescriptions[0].Tags {
			tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
		}
	}
	c.logger.Debugf("[AccountID: %s] getTagsV2: Found %d tags for LB ARN %s in region %s.", operationalAccountID, len(tags), aws.ToString(lb.LoadBalancerArn), regionName)
	return tags
}

// getTagsV1 now accepts context and operationalAccountID
func (c *Client) getTagsV1(ctx context.Context, lb elbTypes.LoadBalancerDescription, regionName string, operationalAccountID string) map[string]string {
	// Use operationalAccountID for getting the ELB client
	elbClient, err := c.getELBClient(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getTagsV1: Error getting ELB client for region %s: %v", operationalAccountID, regionName, err)
		return nil
	}
	if elbClient == nil {
		c.logger.Warnf("[AccountID: %s] getTagsV1: Classic ELB client is nil for region %s. Cannot get tags.", operationalAccountID, regionName)
		return nil
	}

	input := &elb.DescribeTagsInput{
		LoadBalancerNames: []string{aws.ToString(lb.LoadBalancerName)},
	}

	result, err := elbClient.DescribeTags(ctx, input) // Pass ctx
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getTagsV1: DescribeTags failed for LB Name %s, region %s: %v", operationalAccountID, aws.ToString(lb.LoadBalancerName), regionName, err)
		return nil
	}

	tags := make(map[string]string)
	if len(result.TagDescriptions) > 0 {
		for _, tag := range result.TagDescriptions[0].Tags {
			tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
		}
	}
	c.logger.Debugf("[AccountID: %s] getTagsV1: Found %d tags for LB Name %s in region %s.", operationalAccountID, len(tags), aws.ToString(lb.LoadBalancerName), regionName)
	return tags
}

func convertListeners(listeners []elbTypes.ListenerDescription) []types.LBListener {
	var result []types.LBListener
	for _, l := range listeners {
		if l.Listener == nil { // Add nil check for safety
			continue
		}
		result = append(result, types.LBListener{
			// ListenerID is not directly available for Classic ELB listeners in the same way as ELBv2.
			// It's identified by protocol and port.
			ListenerID: fmt.Sprintf("%s:%d", aws.ToString(l.Listener.Protocol), l.Listener.LoadBalancerPort),
			Protocol:   aws.ToString(l.Listener.Protocol),
			Port:       l.Listener.LoadBalancerPort,
			// Classic ELBs don't have TargetGroups
		})
	}
	return result
}

// getPublicIPFromAllocationID retrieves the public IP address for a given EIP Allocation ID.
func (c *Client) getPublicIPFromAllocationID(ctx context.Context, allocationID, regionName, operationalAccountID string) (string, error) {
	ec2Client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		return "", fmt.Errorf("failed to get EC2 client: %w", err)
	}
	if ec2Client == nil {
		return "", fmt.Errorf("EC2 client is nil for region %s", regionName)
	}

	input := &ec2.DescribeAddressesInput{
		AllocationIds: []string{allocationID},
	}

	result, err := ec2Client.DescribeAddresses(ctx, input)
	if err != nil {
		return "", err
	}

	if len(result.Addresses) > 0 && result.Addresses[0].PublicIp != nil {
		return *result.Addresses[0].PublicIp, nil
	}

	return "", fmt.Errorf("no public IP found for allocation ID %s", allocationID)
}

func (c *Client) getInstanceIDsForTargetGroups(ctx context.Context, targetGroupARNs []string, regionName string, operationalAccountID string) ([]string, error) {
	if len(targetGroupARNs) == 0 {
		return nil, nil
	}

	elbv2Client, err := c.getELBv2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get ELBv2 client: %w", err)
	}
	if elbv2Client == nil {
		return nil, fmt.Errorf("ELBv2 client is nil for region %s", regionName)
	}

	describeTGsInput := &elbv2.DescribeTargetGroupsInput{
		TargetGroupArns: targetGroupARNs,
	}
	tgOutput, err := elbv2Client.DescribeTargetGroups(ctx, describeTGsInput)
	if err != nil {
		return nil, fmt.Errorf("failed to describe target groups: %w", err)
	}

	var instanceIDs []string
	var ipTargets []string

	for _, tg := range tgOutput.TargetGroups {
		if tg.TargetGroupArn == nil {
			continue
		}
		tgArn := *tg.TargetGroupArn

		describeHealthInput := &elbv2.DescribeTargetHealthInput{
			TargetGroupArn: &tgArn,
		}
		healthOutput, err := elbv2Client.DescribeTargetHealth(ctx, describeHealthInput)
		if err != nil {
			c.logger.Warnf("[AccountID: %s] could not describe target health for TG %s: %v", operationalAccountID, tgArn, err)
			continue
		}

		for _, thd := range healthOutput.TargetHealthDescriptions {
			if thd.Target == nil || thd.Target.Id == nil {
				continue
			}
			c.logger.Debugf("[AccountID: %s] Target %s in TG %s has health status: %s", operationalAccountID, *thd.Target.Id, tgArn, thd.TargetHealth.State)

			switch tg.TargetType {
			case elbv2types.TargetTypeEnumInstance:
				instanceIDs = append(instanceIDs, *thd.Target.Id)
			case elbv2types.TargetTypeEnumIp:
				ipTargets = append(ipTargets, *thd.Target.Id)
			case elbv2types.TargetTypeEnumLambda, elbv2types.TargetTypeEnumAlb:
				c.logger.Debugf("[AccountID: %s] Skipping non-instance target type %s for target %s in TG %s", operationalAccountID, tg.TargetType, *thd.Target.Id, tgArn)
			}
		}
	}

	if len(ipTargets) > 0 {
		resolvedInstanceIDs, err := c.resolveIPsToInstanceIDs(ctx, ipTargets, regionName, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] Failed to resolve some IP targets to instance IDs: %v", operationalAccountID, err)
		}
		instanceIDs = append(instanceIDs, resolvedInstanceIDs...)
	}

	// Return unique instance IDs
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range instanceIDs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}

func (c *Client) resolveIPsToInstanceIDs(ctx context.Context, ips []string, regionName string, operationalAccountID string) ([]string, error) {
	ec2Client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get EC2 client: %w", err)
	}
	if ec2Client == nil {
		return nil, fmt.Errorf("EC2 client is nil for region %s", regionName)
	}

	input := &ec2.DescribeNetworkInterfacesInput{
		Filters: []awsTypes.Filter{
			{
				Name:   aws.String("addresses.private-ip-address"),
				Values: ips,
			},
		},
	}

	var instanceIDs []string
	paginator := ec2.NewDescribeNetworkInterfacesPaginator(ec2Client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to describe network interfaces for IP resolution: %w", err)
		}
		for _, ni := range page.NetworkInterfaces {
			if ni.Attachment != nil && ni.Attachment.InstanceId != nil {
				instanceIDs = append(instanceIDs, *ni.Attachment.InstanceId)
			}
		}
	}
	return instanceIDs, nil
}

// getIPsFromDNS (utility function, assuming it exists elsewhere or you want to define it)
// For demonstration, a placeholder:
// func getIPsFromDNS(dnsName string) ([]string, error) {
// 	ips, err := net.LookupHost(dnsName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ips, nil
// }
