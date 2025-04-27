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
	"strconv"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	awslbtypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go"
)

var allPorts = types.Ports{""}
var allProtocolsAndPorts = types.ProtocolsAndPorts{"-1": allPorts}

func (c *Client) AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region,
	vpcID string, securityGroupName string, labels map[string]string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (securityGroupId string, instances []types.Instance, err error) {
	allInstances, err := c.getInstancesForLabels(ctx, account, region, labels, vpcID)
	if err != nil {
		return "", nil, err
	}
	return c.updateSecurityGroupsForInstances(ctx, account, region,
		vpcID, securityGroupName, cidrsToAllow,
		protocolsAndPorts, allInstances)
}

func (c *Client) AddInboundAllowRuleBySubnetMatch(ctx context.Context, account, region,
	vpcID string, securityGroupName string, subnetCidrs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (securityGroupId string, instances []types.Instance, subnets []types.Subnet, err error) {

	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to get EC2 client for subnet conversion: %w", err)
	}

	allInstances, allSubnetsAws, err := c.getInstancesForPrefixes(ctx, account, region, subnetCidrs, vpcID)
	if err != nil {
		return "", nil, nil, err
	}
	securityGroupId, instances, err = c.updateSecurityGroupsForInstances(ctx, account, region,
		vpcID, securityGroupName, cidrsToAllow,
		protocolsAndPorts, allInstances)
	if err != nil {
		return securityGroupId, instances, nil, err
	}

	subnets = c.convertSubnets(ctx, client, c.defaultAccountID, c.defaultRegion, account, region, allSubnetsAws)
	return securityGroupId, instances, subnets, nil
}

func (c *Client) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, account, region,
	vpcID string, securityGroupName string, instancesIPs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (securityGroupId string, instances []types.Instance, err error) {
	allInstances, err := c.getInstancesForPrivateIPs(ctx, account, region, instancesIPs, vpcID)
	if err != nil {
		return "", nil, err
	}
	return c.updateSecurityGroupsForInstances(ctx, account, region,
		vpcID, securityGroupName, cidrsToAllow,
		protocolsAndPorts, allInstances)
}

func (c *Client) updateSecurityGroupsForInstances(ctx context.Context, account, region,
	vpcID string, securityGroupName string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts, allInstances []awsec2types.Reservation) (securityGroupId string, instances []types.Instance, err error) {
	c.logger.Infof("updating AWS security groups for %d instances", len(allInstances))

	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withSecurityGroupName(securityGroupName)
	getParam := &ec2.DescribeSecurityGroupsInput{
		Filters: builder.build(),
	}
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return "", nil, err
	}
	getOut, err := client.DescribeSecurityGroups(ctx, getParam)
	if err == nil && getOut != nil && len(getOut.SecurityGroups) > 0 {
		securityGroupId = *getOut.SecurityGroups[0].GroupId
	} else {
		// try to create
		param := &ec2.CreateSecurityGroupInput{
			VpcId:       &vpcID,
			Description: &securityGroupName,
			GroupName:   &securityGroupName,
		}
		out, err := client.CreateSecurityGroup(ctx, param)
		if err != nil {
			return "", nil, err
		}
		securityGroupId = *out.GroupId
	}

	for _, instanceIP := range cidrsToAllow {
		if err := c.updateIngressSecurityGroupWithACL(ctx, account, region, protocolsAndPorts,
			instanceIP, securityGroupId, nil); err != nil {
			return "", nil, err
		}
	}
	instances = append(instances, convertInstances(c.defaultAccountID, c.defaultRegion, account, region, allInstances)...)
	for _, instance := range allInstances {
		for _, inst := range instance.Instances {
			groups := make([]string, 0, len(inst.SecurityGroups)+1)
			for _, group := range inst.SecurityGroups {
				groups = append(groups, *group.GroupId)
			}
			groups = append(groups, securityGroupId)
			p := &ec2.ModifyNetworkInterfaceAttributeInput{
				NetworkInterfaceId: inst.NetworkInterfaces[0].NetworkInterfaceId,
				Groups:             groups,
			}
			_, err := client.ModifyNetworkInterfaceAttribute(ctx, p)
			if err != nil {
				return "", nil, err
			}
		}
	}
	return
}

func (c *Client) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context,
	account, region string,
	loadBalancerDNS string,
	vpcID string,
	securityGroupName string,
	cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts,
) (loadBalancerId, securityGroupId string, err error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return "", "", err
	}
	lbClient, err := c.getELBClient(ctx, account, region)
	if err != nil {
		return "", "", err
	}
	var lb awslbtypes.LoadBalancerDescription
	// check if this is internal load balancer with default DNS 'internal-<LB_NAME>-xxxxx.<REGION>.elb.amazonaws.com:
	splitted := strings.Split(loadBalancerDNS, "-")
	if len(splitted) > 1 && splitted[0] == "internal" {
		getParam := &elasticloadbalancing.DescribeLoadBalancersInput{
			LoadBalancerNames: []string{splitted[1]},
		}
		lbs, err := lbClient.DescribeLoadBalancers(ctx, getParam)
		if err != nil {
			return "", "", err
		}
		if len(lbs.LoadBalancerDescriptions) == 0 {
			return "", "", fmt.Errorf("couldn't find load balancer with DNS %s", loadBalancerDNS)
		}
		lb = lbs.LoadBalancerDescriptions[0]
	} else {
		// otherwise check all load balancers and look for matching one:
		lbs, err := lbClient.DescribeLoadBalancers(ctx, &elasticloadbalancing.DescribeLoadBalancersInput{})
		if err != nil {
			return "", "", err
		}
		for _, lb := range lbs.LoadBalancerDescriptions {
			if lb.DNSName != nil && *lb.DNSName == loadBalancerDNS {
				lb = lbs.LoadBalancerDescriptions[0]
				break
			}
		}
	}

	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withSecurityGroupName(securityGroupName)
	getParam := &ec2.DescribeSecurityGroupsInput{
		Filters: builder.build(),
	}
	getOut, err := client.DescribeSecurityGroups(ctx, getParam)
	if err == nil && getOut != nil && len(getOut.SecurityGroups) > 0 {
		securityGroupId = *getOut.SecurityGroups[0].GroupId
	} else {
		// try to create
		param := &ec2.CreateSecurityGroupInput{
			VpcId:       &vpcID,
			Description: &securityGroupName,
			GroupName:   &securityGroupName,
		}
		out, err := client.CreateSecurityGroup(ctx, param)
		if err != nil {
			return "", "", err
		}
		securityGroupId = *out.GroupId
	}

	for _, instanceIP := range cidrsToAllow {
		if err := c.updateIngressSecurityGroupWithACL(ctx, account, region, protocolsAndPorts,
			instanceIP, securityGroupId, nil); err != nil {
			return "", "", err
		}
	}

	_, err = lbClient.ApplySecurityGroupsToLoadBalancer(ctx,
		&elasticloadbalancing.ApplySecurityGroupsToLoadBalancerInput{
			LoadBalancerName: lb.LoadBalancerName,
			SecurityGroups:   append(lb.SecurityGroups, securityGroupId),
		})
	if err != nil {
		return "", "", err
	}

	return *lb.LoadBalancerName, securityGroupId, nil
}

func (c *Client) RemoveInboundAllowRulesFromVPCById(ctx context.Context, account, region string,
	vpcId string, instanceIDs []string, loadBalancersIds []string, securityGroupId string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	lbClient, err := c.getELBClient(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcId)
	input := &ec2.DescribeInstancesInput{
		Filters: builder.build(),
	}
	instances, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return err
	}
	for _, instance := range instances.Reservations {
		for _, inst := range instance.Instances {
			found := false
			for _, id := range instanceIDs {
				if *inst.InstanceId == id {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			groups := make([]string, 0, len(inst.SecurityGroups)-1)
			for _, group := range inst.SecurityGroups {
				if *group.GroupId != securityGroupId {
					groups = append(groups, *group.GroupId)
				}
			}
			p := &ec2.ModifyNetworkInterfaceAttributeInput{
				NetworkInterfaceId: inst.NetworkInterfaces[0].NetworkInterfaceId,
				Groups:             groups,
			}
			_, err := client.ModifyNetworkInterfaceAttribute(ctx, p)
			if err != nil {
				return err
			}
		}
	}

	getParam := &elasticloadbalancing.DescribeLoadBalancersInput{
		LoadBalancerNames: loadBalancersIds,
	}
	lbs, err := lbClient.DescribeLoadBalancers(ctx, getParam)
	if err != nil {
		return err
	}
	for _, lb := range lbs.LoadBalancerDescriptions {
		updatedSecurityGroups := make([]string, 0, len(lb.SecurityGroups)-1)
		found := false
		for _, id := range lb.SecurityGroups {
			if id == securityGroupId {
				found = true
			} else {
				updatedSecurityGroups = append(updatedSecurityGroups, id)
			}
		}
		if !found {
			continue
		}
		_, err = lbClient.ApplySecurityGroupsToLoadBalancer(ctx,
			&elasticloadbalancing.ApplySecurityGroupsToLoadBalancerInput{
				LoadBalancerName: lb.LoadBalancerName,
				SecurityGroups:   updatedSecurityGroups,
			})
		if err != nil {
			return err
		}
	}

	param := &ec2.DeleteSecurityGroupInput{
		GroupId: &securityGroupId,
	}
	if _, err := client.DeleteSecurityGroup(ctx, param); err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return err
		}

		if "InvalidGroup.NotFound" == apiErr.ErrorCode() {
			c.logger.Infof("Security group %s already removed", securityGroupId)
			return nil
		}
		return err
	}
	c.logger.Infof("Deleted security group %s", securityGroupId)
	return nil
}

func (c *Client) AddInboundAllowRuleInVPC(ctx context.Context, account, region string,
	vpcID string, cidrs []string, securityGroupName string, tags map[string]string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	builder.withSecurityGroupName(securityGroupName)
	getParam := &ec2.DescribeSecurityGroupsInput{
		Filters: builder.build(),
	}
	getOut, err := client.DescribeSecurityGroups(ctx, getParam)
	var groupId string
	if err == nil && getOut != nil && len(getOut.SecurityGroups) > 0 {
		groupId = *getOut.SecurityGroups[0].GroupId
	} else {
		// try to create
		param := &ec2.CreateSecurityGroupInput{
			VpcId:       &vpcID,
			Description: &securityGroupName,
			GroupName:   &securityGroupName,
		}
		out, err := client.CreateSecurityGroup(ctx, param)
		if err != nil {
			return err
		}
		groupId = *out.GroupId
	}

	c.logger.Infof("Created security group %s", groupId)
	for _, cidr := range cidrs {
		if err = c.updateIngressSecurityGroupWithACL(ctx, account, region, nil, cidr, groupId, tags); err != nil {
			return err
		}
		c.logger.Infof("Updated security group %s with rule allowing cidr %s", groupId, cidr)
	}

	instances, err := c.getInstancesForVPC(ctx, account, region, vpcID)
	if err != nil {
		return err
	}
	for _, instance := range instances {
		for _, inst := range instance.Instances {
			groups := make([]string, 0, len(inst.SecurityGroups)+1)
			for _, group := range inst.SecurityGroups {
				groups = append(groups, *group.GroupId)
			}
			groups = append(groups, groupId)
			p := &ec2.ModifyNetworkInterfaceAttributeInput{
				NetworkInterfaceId: inst.NetworkInterfaces[0].NetworkInterfaceId,
				Groups:             groups,
			}
			_, err = client.ModifyNetworkInterfaceAttribute(ctx, p)
			if err != nil {
				return err
			}
			c.logger.Infof("Updated instance with security group %s", groupId)
		}
	}

	return nil
}

func (c *Client) RemoveInboundAllowRuleRulesByTags(ctx context.Context, account, region, vpcID, securityGroupName string, tags map[string]string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	sgbuilder := newFilterBuilder()
	sgbuilder.withVPC(vpcID)
	sgbuilder.withSecurityGroupName(securityGroupName)
	getParam := &ec2.DescribeSecurityGroupsInput{
		Filters: sgbuilder.build(),
	}
	getOut, err := client.DescribeSecurityGroups(ctx, getParam)
	if err != nil {
		return err
	}

	if len(getOut.SecurityGroups) == 0 {
		c.logger.Infof("there are no matching security groups to remove rules from, security group name %s", securityGroupName)
		return nil
	}

	builder := newFilterBuilder()
	for k, v := range tags {
		builder.withTag(k, v)
	}
	builder.withSecurityGroupId(*getOut.SecurityGroups[0].GroupId)
	param := &ec2.DescribeSecurityGroupRulesInput{
		Filters: builder.build(),
	}
	out, err := client.DescribeSecurityGroupRules(ctx, param)
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(out.SecurityGroupRules))
	for _, rule := range out.SecurityGroupRules {
		ids = append(ids, *rule.SecurityGroupRuleId)
	}
	revokeParams := &ec2.RevokeSecurityGroupIngressInput{
		GroupId:              getOut.SecurityGroups[0].GroupId,
		SecurityGroupRuleIds: ids,
	}

	_, err = client.RevokeSecurityGroupIngress(ctx, revokeParams)
	if err != nil {
		return err
	}

	getOut, err = client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{*getOut.SecurityGroups[0].GroupId},
	})
	if err != nil {
		return err
	}

	if len(getOut.SecurityGroups[0].IpPermissions) == 0 {
		err := c.RemoveInboundAllowRuleFromVPCByName(ctx, account, region, vpcID, *getOut.SecurityGroups[0].GroupId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, account, region, vpcID string, securityGroupID string) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	input := &ec2.DescribeInstancesInput{
		Filters: builder.build(),
	}
	instances, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return err
	}
	for _, instance := range instances.Reservations {
		for _, inst := range instance.Instances {
			// assuming instance may or may not contain given security group or making room for all existing
			groups := make([]string, 0, len(inst.SecurityGroups))
			foundGroup := false
			for _, group := range inst.SecurityGroups {
				if *group.GroupId != securityGroupID {
					groups = append(groups, *group.GroupId)
				} else {
					foundGroup = true
				}
			}
			if foundGroup {
				p := &ec2.ModifyNetworkInterfaceAttributeInput{
					NetworkInterfaceId: inst.NetworkInterfaces[0].NetworkInterfaceId,
					Groups:             groups,
				}
				_, err := client.ModifyNetworkInterfaceAttribute(ctx, p)
				if err != nil {
					return err
				}
				c.logger.Infof("Removed security group %s from instance", securityGroupID)
			}
		}
	}
	param := &ec2.DeleteSecurityGroupInput{
		GroupId: &securityGroupID,
	}
	if _, err := client.DeleteSecurityGroup(ctx, param); err != nil {
		return err
	}
	c.logger.Infof("Deleted security group %s", securityGroupID)
	return nil
}

func (c *Client) RefreshInboundAllowRule(ctx context.Context,
	account, region string,
	securityGroupId string,
	cidrsToAdd []string,
	cidrsToRemove []string,
	destinationLabels map[string]string,
	destinationPrefixes []string,
	destinationVPCId string,
	protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error) {

	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, nil, err
	}
	for _, cidr := range cidrsToAdd {
		if err := c.updateIngressSecurityGroupWithACL(ctx, account, region, protocolsAndPorts, cidr, securityGroupId, nil); err != nil {
			return nil, nil, err
		}
	}
	for _, cidr := range cidrsToRemove {
		if err := c.deleteIngressSecurityGroupWithACL(ctx, client, cidr, securityGroupId); err != nil {
			return nil, nil, err
		}
	}

	// updating new destination instances with ACL security group
	var destInstances []awsec2types.Reservation
	var destSubnetsAws []awsec2types.Subnet
	if len(destinationLabels) > 0 {
		subnetWithVPCID := make(map[string]struct{})
		for _, instance := range destInstances {
			for _, i := range instance.Instances {
				subnetWithVPCID[convertString(i.VpcId)+"/"+convertString(i.SubnetId)] = struct{}{}
			}
		}
		for subnetVPCID := range subnetWithVPCID {
			split := strings.Split(subnetVPCID, "/")
			subnet, err := c.GetSubnet(ctx, &infrapb.GetSubnetRequest{
				VpcId:     split[0],
				Id:        split[1],
				Region:    region,
				AccountId: account,
			})
			if err != nil {
				return nil, nil, err
			}
			subnets = append(subnets, subnet)
		}
	} else if len(destinationPrefixes) > 0 {
		destInstances, destSubnetsAws, err = c.getInstancesForPrefixes(ctx, account, region, destinationPrefixes, destinationVPCId)
		if err != nil {
			return nil, nil, err
		}
		subnets = c.convertSubnets(ctx, client, c.defaultAccountID, c.defaultRegion, account, region, destSubnetsAws)
	} else {
		return nil, nil, nil
	}

	instances = append(instances, convertInstances(c.defaultAccountID, c.defaultRegion, account, region, destInstances)...)
	for _, resrv := range destInstances {
		for _, instance := range resrv.Instances {
			found := false
			for _, sg := range instance.SecurityGroups {
				if *sg.GroupId == securityGroupId {
					found = true
					break
				}
			}
			if found {
				continue
			}
			c.logger.Infof("ACL Refresh: adding security group %s to instance with IP %s",
				securityGroupId, *instance.PrivateIpAddress)

			groups := make([]string, 0, len(instance.SecurityGroups)+1)
			for _, group := range instance.SecurityGroups {
				groups = append(groups, *group.GroupId)
			}
			groups = append(groups, securityGroupId)
			p := &ec2.ModifyNetworkInterfaceAttributeInput{
				NetworkInterfaceId: instance.NetworkInterfaces[0].NetworkInterfaceId,
				Groups:             groups,
			}
			_, err = client.ModifyNetworkInterfaceAttribute(ctx, p)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return instances, subnets, nil
}

func (c *Client) getInstancesForLabels(ctx context.Context, account, region string, labels map[string]string, vpcID string) ([]awsec2types.Reservation, error) {
	builder := newFilterBuilder()
	builder.withVPC(vpcID)
	condition := labels[types.ConditionLabel]
	if condition != types.OrCondition {
		condition = types.AndCondition
	}
	var allInstances []awsec2types.Reservation
	if condition == types.AndCondition {
		for k, v := range labels {
			if k == types.ConditionLabel {
				continue
			}
			builder.withTag(k, v)
		}
		instances, err := c.getInstances(ctx, account, region, builder)
		if err != nil {
			return nil, err
		}
		allInstances = append(allInstances, instances...)
	} else {
		for k, v := range labels {
			if k == types.ConditionLabel {
				continue
			}
			builder = newFilterBuilder()
			builder.withTag(k, v)
			builder.withVPC(vpcID)
			instances, err := c.getInstances(ctx, account, region, builder)
			if err != nil {
				return nil, err
			}
			allInstances = append(allInstances, instances...)
		}
	}
	return allInstances, nil
}

func (c *Client) getInstancesForPrefixes(ctx context.Context, account, region string, prefixes []string, vpcID string) ([]awsec2types.Reservation, []awsec2types.Subnet, error) {
	var allInstances []awsec2types.Reservation
	var allSubnets []awsec2types.Subnet
	for _, prefix := range prefixes {
		builder := newFilterBuilder()
		builder.withVPC(vpcID)
		builder.withCIDR(prefix)
		subnets, err := c.getSubnets(ctx, account, region, builder)
		if err != nil {
			return nil, nil, fmt.Errorf("could not find subnet: %v", err)
		}
		allSubnets = append(allSubnets, subnets...)

		for _, subnet := range subnets {
			builder = newFilterBuilder()
			builder.withVPC(vpcID)
			builder.withSubnet(*subnet.SubnetId)
			instances, err := c.getInstances(ctx, account, region, builder)
			if err != nil {
				return nil, nil, err
			}
			allInstances = append(allInstances, instances...)
		}
	}
	return allInstances, allSubnets, nil
}

func (c *Client) getInstancesForPrivateIPs(ctx context.Context, account, region string, ips []string, vpcID string) ([]awsec2types.Reservation, error) {
	var allInstances []awsec2types.Reservation
	for _, ip := range ips {
		builder := newFilterBuilder()
		builder.withVPC(vpcID)
		builder.withPrivateIPAddress(ip)
		instances, err := c.getInstances(ctx, account, region, builder)
		if err != nil {
			return nil, fmt.Errorf("could not find instance: %v", err)
		}
		allInstances = append(allInstances, instances...)
	}
	return allInstances, nil
}

func (c *Client) getInstancesForVPC(ctx context.Context, account, region, vpcID string) ([]awsec2types.Reservation, error) {
	builder := newFilterBuilder()
	builder.withVPC(vpcID)

	return c.getInstances(ctx, account, region, builder)
}

func (c *Client) updateIngressSecurityGroupWithACL(
	ctx context.Context,
	account, region string,
	aclProtocolsAndPorts types.ProtocolsAndPorts,
	instanceIP string,
	securityGroupID string,
	tags map[string]string,
) error {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return err
	}
	if len(aclProtocolsAndPorts) == 0 {
		aclProtocolsAndPorts = allProtocolsAndPorts
	}
	sgData := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       aws.String(securityGroupID),
		IpPermissions: []awsec2types.IpPermission{},
	}
	if len(tags) > 0 {
		sgData.TagSpecifications = []awsec2types.TagSpecification{
			mapToTagSpecfication(tags, awsec2types.ResourceTypeSecurityGroupRule),
		}
	}
	for protocol, ports := range aclProtocolsAndPorts {
		if len(ports) == 0 {
			ports = allPorts
		}
		for _, port := range ports {
			var fromPort, toPort int32
			var err error
			fromPort, toPort, err = getPortRange(port)
			if err != nil {
				return err
			}

			sgData.IpPermissions = append(sgData.IpPermissions,
				awsec2types.IpPermission{
					IpProtocol: aws.String(protocol),
					IpRanges: []awsec2types.IpRange{
						{
							CidrIp: aws.String(instanceIP),
						},
					},
					FromPort: &fromPort,
					ToPort:   &toPort,
				})
		}
	}
	if _, err := client.AuthorizeSecurityGroupIngress(ctx, sgData); err != nil {
		if !strings.Contains(err.Error(), "InvalidPermission.Duplicate") {
			return err
		}
	}
	return nil
}

func (c *Client) deleteIngressSecurityGroupWithACL(ctx context.Context, client *ec2.Client, instanceIP, securityGroupID string) error {
	cidrIP := instanceIP
	if !strings.Contains(instanceIP, "/") {
		cidrIP = fmt.Sprintf("%s/32", instanceIP)
	}
	sgData := &ec2.RevokeSecurityGroupIngressInput{
		GroupId: aws.String(securityGroupID),
		CidrIp:  aws.String(cidrIP),
		// TODO this should be exact ports and protocols https://github.com/app-net-interface/awi-grpc-catalyst-sdwan/issues/31
		IpProtocol: aws.String("-1"),
	}
	if _, err := client.RevokeSecurityGroupIngress(ctx, sgData); err != nil {
		return err
	}
	return nil
}

func getPortRange(port string) (int32, int32, error) {
	if port == "" {
		return -1, -1, nil
	}
	ports := strings.Split(port, "-")
	if len(ports) == 1 {
		num, err := strconv.Atoi(port)
		if err != nil {
			return 0, 0, fmt.Errorf("expected number as port, got: %s", port)
		}
		return int32(num), int32(num), nil
	}
	if len(ports) != 2 {
		return 0, 0, fmt.Errorf("expected port in format 'from-to', got: %s", port)
	}
	from, err := strconv.Atoi(ports[0])
	if err != nil {
		return 0, 0, fmt.Errorf("expected number as port, got: %s", ports[0])
	}
	to, err := strconv.Atoi(ports[1])
	if err != nil {
		return 0, 0, fmt.Errorf("expected number as port, got: %s", ports[1])
	}
	return int32(from), int32(to), nil
}
