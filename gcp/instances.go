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

package gcp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"golang.org/x/exp/slices"

	"google.golang.org/api/iterator"

	"github.com/app-net-interface/awi-infra-guard/types"
)

// ListInstances returns instances matching to provided vpc ID or label or zone.
func (c *Client) ListInstances(ctx context.Context, params *infrapb.ListInstancesRequest) ([]types.Instance, error) {
	if params == nil {
		params = &infrapb.ListInstancesRequest{}
	}
	var net network
	var err error
	if params.GetVpcId() != "" {
		net, err = c.vpcIdToSingleNetwork(ctx, params.GetAccountId(), params.GetVpcId())
		if err != nil {
			return nil, err
		}
	}

	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	subnets, err := c.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}

	filter := newFilterBuilder()
	condition, ok := params.GetLabels()[types.ConditionLabel]
	if !ok {
		condition = types.AndCondition
	}
	for k, v := range params.GetLabels() {
		if k == types.ConditionLabel {
			continue
		}
		filter.withLabel(k, v, condition)
	}
	instances := make([]types.Instance, 0)
	f := func(projectID string) error {
		if params.GetZone() != "" {
			iter := c.instancesClient.List(ctx, &computepb.ListInstancesRequest{
				Filter:  filter.build(),
				Project: projectID,
				Zone:    params.GetZone(),
			})

			for {
				gcpInstance, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return err
				}
				instance := convertInstance(projectID, networks, subnets, gcpInstance)
				if params.GetVpcId() == "" || params.GetVpcId() == instance.VPCID || net.name == instance.VPCID {
					instances = append(instances, instance)
				}
			}
		} else {
			it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
				Project: projectID,
				Filter:  filter.build(),
			},
			)
			for {
				pair, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return err
				}
				for _, gcpInstance := range pair.Value.Instances {
					instance := convertInstance(projectID, networks, subnets, gcpInstance)
					if params.GetVpcId() == "" || params.GetVpcId() == instance.VPCID || net.name == instance.VPCID {
						instances = append(instances, instance)
					}
				}
			}
		}
		return nil
	}
	if params.GetAccountId() == "" {
		for projectID := range c.projectIDs {
			err := f(projectID)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}

	return instances, nil
}

func (c *Client) GetIPsForLabels(ctx context.Context, params *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	filter := newFilterBuilder()
	condition, ok := params.GetLabels()[types.ConditionLabel]
	if !ok {
		condition = types.AndCondition
	}
	for k, v := range params.GetLabels() {
		if k == types.ConditionLabel {
			continue
		}
		filter.withLabel(k, v, condition)
	}

	ipsMap := map[string]struct{}{}
	f := func(projectID string) error {
		it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
			Project: projectID,
			Filter:  filter.build(),
		},
		)
		for {
			pair, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			for _, gcpInstance := range pair.Value.Instances {
				for _, iface := range gcpInstance.NetworkInterfaces {
					ipsMap[iface.GetNetworkIP()] = struct{}{}
				}
			}
		}
		return nil
	}
	if params.GetAccountId() == "" {
		for projectID := range c.projectIDs {
			err := f(projectID)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}

	ips := make([]string, 0, len(ipsMap))
	for ip := range ipsMap {
		ips = append(ips, ip)
	}
	return ips, nil
}

func (c *Client) GetInstancesForLabels(ctx context.Context, params *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	net, err := c.vpcIdToSingleNetwork(ctx, params.GetAccountId(), params.GetVpcId())
	if err != nil {
		return nil, err
	}
	filter := newFilterBuilder()
	condition, ok := params.GetLabels()[types.ConditionLabel]
	if !ok {
		condition = types.AndCondition
	}
	for k, v := range params.GetLabels() {
		if k == types.ConditionLabel {
			continue
		}
		filter.withLabel(k, v, condition)
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	subnets, err := c.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	var instances []types.Instance
	f := func(projectID string) error {
		it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
			Project: projectID,
			Filter:  filter.build(),
		},
		)
		for {
			pair, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			for _, gcpInstance := range pair.Value.Instances {
				instance := convertInstance(projectID, networks, subnets, gcpInstance)
				if params.GetVpcId() == instance.VPCID || net.name == instance.VPCID {
					instances = append(instances, instance)
				}
			}
		}
		return nil
	}

	if params.GetAccountId() == "" {
		for projectID := range c.projectIDs {
			err := f(projectID)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(params.GetAccountId())
		if err != nil {
			return nil, err
		}
	}

	return instances, nil
}

func (c *Client) addNetworkTagToInstancesByNetworkAndLabels(ctx context.Context, tag string, net network, labels map[string]string) ([]types.Instance, error) {
	filter := newFilterBuilder()
	condition, ok := labels[types.ConditionLabel]
	if !ok {
		condition = types.AndCondition
	}
	for k, v := range labels {
		if k == types.ConditionLabel {
			continue
		}
		filter.withLabel(k, v, condition)
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	subnets, err := c.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}

	var instances []types.Instance
	it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
		Project: net.project,
		Filter:  filter.build(),
	},
	)

	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, gcpInstance := range pair.Value.Instances {
			for _, iface := range gcpInstance.NetworkInterfaces {
				if iface.GetNetwork() == net.fullUrl {
					instances = append(instances, convertInstance(net.project, networks, subnets, gcpInstance))
					err = c.addNetworkTagToInstance(ctx, tag, net.project, gcpInstance)
					if err != nil {
						return nil, err
					}
					break
				}
			}
		}
	}
	return instances, nil
}

func (c *Client) addNetworkTagToInstancesByNetworkAndSubnets(ctx context.Context, tag string, net network, subnetCidrs []string) ([]types.Instance, []types.Subnet, error) {
	var gcpSubnets []*computepb.Subnetwork
	var subnets []types.Subnet
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, nil, err
	}

	for _, cidr := range subnetCidrs {
		s, err := c.getSubnetsByNetworkAndCidr(ctx, net, cidr)
		if err != nil {
			return nil, nil, err
		}
		gcpSubnets = append(gcpSubnets, s...)
		for _, subnet := range s {
			subnets = append(subnets, convertSubnet(net.project, networks, subnet))
		}
	}

	var instancesIds []string
	var instances []types.Instance
	it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
		Project: net.project,
	},
	)
	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		for _, subnet := range gcpSubnets {
			for _, gcpInstance := range pair.Value.Instances {
				id := strconv.FormatUint(gcpInstance.GetId(), 10)
				for _, iface := range gcpInstance.NetworkInterfaces {
					if iface.GetNetwork() == net.fullUrl && iface.GetSubnetwork() == subnet.GetSelfLink() {
						// check if instance wasn't already handled in other subnet
						instanceAlreadyHandled := false
						for _, instancesId := range instancesIds {
							if instancesId == id {
								instanceAlreadyHandled = true
								break
							}
						}
						if instanceAlreadyHandled {
							break
						}

						instancesIds = append(instancesIds, id)
						instances = append(instances, convertInstance(net.project, networks, subnets, gcpInstance))
						err = c.addNetworkTagToInstance(ctx, tag, net.project, gcpInstance)
						if err != nil {
							return nil, nil, err
						}
						break
					}
				}
			}
		}
	}
	return instances, subnets, nil
}

func (c *Client) addNetworkTagToInstancesByNetworkAndIps(ctx context.Context, tag string, net network, ips []string) ([]types.Instance, error) {
	var instancesIds []string
	var instances []types.Instance
	it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
		Project: net.project,
	},
	)
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	subnets, err := c.ListSubnets(ctx, &infrapb.ListSubnetsRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, ip := range ips {
			for _, gcpInstance := range pair.Value.Instances {
				id := strconv.FormatUint(gcpInstance.GetId(), 10)
				for _, iface := range gcpInstance.NetworkInterfaces {
					if iface.GetNetwork() == net.fullUrl && iface.GetNetworkIP() == ip {
						// check if instance wasn't already handled by other IP
						instanceAlreadyHandled := false
						for _, instancesId := range instancesIds {
							if instancesId == id {
								instanceAlreadyHandled = true
								break
							}
						}
						if instanceAlreadyHandled {
							break
						}

						instancesIds = append(instancesIds, id)
						instances = append(instances, convertInstance(net.project, networks, subnets, gcpInstance))
						err = c.addNetworkTagToInstance(ctx, tag, net.project, gcpInstance)
						if err != nil {
							return nil, err
						}
						break
					}
				}
			}
		}
	}
	return instances, nil
}

func (c *Client) addNetworkTagToInstance(ctx context.Context, tag string, project string, instance *computepb.Instance) error {
	for _, t := range instance.GetTags().GetItems() {
		if t == tag {
			c.logger.Infof("Tag %s already present in instance %s", tag, instance.GetName())
			return nil
		}
	}
	var zone string
	zoneSplit := strings.Split(instance.GetZone(), "/")
	if len(zoneSplit) > 0 {
		zone = zoneSplit[len(zoneSplit)-1]
	} else {
		return fmt.Errorf("couldn't determine zone of instance %s", instance.GetName())
	}
	op, err := c.instancesClient.SetTags(ctx, &computepb.SetTagsInstanceRequest{
		Instance: instance.GetName(),
		Project:  project,
		Zone:     zone,
		TagsResource: &computepb.Tags{
			Items:       append(instance.GetTags().GetItems(), tag),
			Fingerprint: instance.Tags.Fingerprint,
		},
	})
	if err != nil {
		return err
	}
	err = op.Wait(ctx)
	if err != nil {
		return err
	}
	c.logger.Infof("Added tag %s to instance %s in project %s", tag, instance.GetName(),
		project)
	return nil
}

func (c *Client) removeNetworkTagFromInstance(ctx context.Context, tag string, project string, instance *computepb.Instance) error {
	var zone string
	zoneSplit := strings.Split(instance.GetZone(), "/")
	if len(zoneSplit) > 0 {
		zone = zoneSplit[len(zoneSplit)-1]
	} else {
		return fmt.Errorf("couldn't determine zone of instance %s", instance.GetName())
	}

	for k, v := range instance.GetTags().GetItems() {
		if v == tag {
			instance.Tags.Items = slices.Delete(instance.Tags.Items, k, k+1)
		}
	}

	op, err := c.instancesClient.SetTags(ctx, &computepb.SetTagsInstanceRequest{
		Instance: instance.GetName(),
		Project:  project,
		Zone:     zone,
		TagsResource: &computepb.Tags{
			Items:       instance.Tags.Items,
			Fingerprint: instance.Tags.Fingerprint,
		},
	})
	if err != nil {
		return err
	}
	err = op.Wait(ctx)
	if err != nil {
		return err
	}
	c.logger.Infof("Removed tag %s from instance %s in project %s", tag, instance.GetName(),
		project)
	return nil
}

func (c *Client) removeNetworkTagFromInstancesByIDs(ctx context.Context, project, tag string, instancesIDs []string) error {
	it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
		Project: project,
	},
	)
	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		for _, instancesID := range instancesIDs {
			for _, gcpInstance := range pair.Value.Instances {
				if strconv.FormatUint(gcpInstance.GetId(), 10) == instancesID {
					err := c.removeNetworkTagFromInstance(ctx, tag, project, gcpInstance)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func convertInstance(projectID string, networks []types.VPC, subnets []types.Subnet, gcpInstance *computepb.Instance) types.Instance {
	machineParts := strings.Split(gcpInstance.GetMachineType(), "/")

	newInstance := types.Instance{
		ID:        strconv.FormatUint(gcpInstance.GetId(), 10),
		Name:      gcpInstance.GetName(),
		Labels:    gcpInstance.GetLabels(),
		Zone:      gcpInstance.GetZone(),
		Type:      machineParts[len(machineParts)-1],
		AccountID: projectID,
		Provider:  providerName,
		State:     gcpInstance.GetStatus(),
		SelfLink:  *gcpInstance.SelfLink,
	}
	newInstance.PrivateIP, newInstance.PublicIP, newInstance.VPCID, newInstance.SubnetID =
		getNetwork(networks, subnets, gcpInstance.GetNetworkInterfaces())
	return newInstance
}

func getNetwork(networks []types.VPC, subnetworks []types.Subnet, networkInterfaces []*computepb.NetworkInterface) (privateIP, publicIP, vpcID, subnetID string) {
	if len(networkInterfaces) == 0 {
		return
	}
	privateIP = networkInterfaces[0].GetNetworkIP()
	if len(networkInterfaces[0].GetAccessConfigs()) != 0 {
		publicIP = networkInterfaces[0].GetAccessConfigs()[0].GetNatIP()
	}
	network := strings.Split(networkInterfaces[0].GetNetwork(), "/")
	if len(network) != 0 {
		name := network[len(network)-1]
		for _, v := range networks {
			if v.Name == name || v.ID == name {
				vpcID = v.ID
				break
			}
		}
	}
	subnetwork := strings.Split(networkInterfaces[0].GetSubnetwork(), "/")
	if len(subnetwork) != 0 {
		name := subnetwork[len(subnetwork)-1]
		for _, v := range subnetworks {
			if v.Name == name || v.SubnetId == name {
				subnetID = v.SubnetId
				break
			}
		}
	}
	return
}
