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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListInstances(ctx context.Context, params *infrapb.ListInstancesRequest) ([]types.Instance, error) {
	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	builder.withAvailabilityZone(params.GetZone())
	in := &ec2.DescribeInstancesInput{
		Filters: builder.build(),
	}
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allInstances  []types.Instance
			resultChannel = make(chan []types.Instance)
			errorChannel  = make(chan error)
		)

		regionalClients, err := c.getAllClientsForProfile(params.GetAccountId())
		if err != nil {
			return nil, err
		}
		for regionName, awsRegionClient := range regionalClients {
			wg.Add(1)
			go func(regionName string, awsRegionClient awsClient) {
				defer wg.Done()

				out, err := awsRegionClient.ec2Client.DescribeInstances(ctx, in)
				if err != nil {
					errorChannel <- fmt.Errorf("could not get AWS instances: %v", err)
					return
				}

				instances := convertInstances(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), regionName, out.Reservations)
				resultChannel <- instances
			}(regionName, awsRegionClient)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			close(errorChannel)
		}()

		for instances := range resultChannel {
			allInstances = append(allInstances, instances...)
		}

		if err := <-errorChannel; err != nil {
			return nil, err
		}
		return allInstances, nil
	}
	client, err := c.getEC2Client(ctx, params.GetAccountId(), params.GetRegion())
	if err != nil {
		return nil, err
	}
	out, err := client.DescribeInstances(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("could not get AWS instances: %v", err)
	}
	return convertInstances(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), params.GetRegion(), out.Reservations), nil
}

func convertInstances(defaultAccount, defaultRegion, account, region string, reservations []awstypes.Reservation) []types.Instance {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	instances := make([]types.Instance, 0, len(reservations))
	for _, reservation := range reservations {
		if len(reservation.Instances) == 0 {
			continue
		}
		inst := reservation.Instances[0]
		name := getTagName(inst.Tags)
		instanceLink := fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", region, region, aws.ToString(inst.InstanceId))

		secGroups := make([]string, len(inst.SecurityGroups))
		for i, group := range inst.SecurityGroups {
			secGroups[i] = *group.GroupId
		}

		networkInterfaces := make([]string, len(inst.NetworkInterfaces))
		for j, iface := range inst.NetworkInterfaces {
			networkInterfaces[j] = *iface.NetworkInterfaceId
		}

		instance := types.Instance{
			ID:               aws.ToString(inst.InstanceId),
			Name:             aws.ToString(name),
			PrivateIP:        aws.ToString(inst.PrivateIpAddress),
			PublicIP:         aws.ToString(inst.PublicIpAddress),
			SubnetID:         aws.ToString(inst.SubnetId),
			VPCID:            aws.ToString(inst.VpcId),
			Type:             aws.ToString((*string)(&inst.InstanceType)),
			Labels:           convertTags(inst.Tags),
			State:            convertState(inst.State),
			Region:           region,
			Zone:             aws.ToString(inst.Placement.AvailabilityZone),
			AccountID:        account,
			Provider:         providerName,
			SelfLink:         instanceLink,
			SecurityGroupIDs: secGroups,
			InterfaceIDs:     networkInterfaces,
		}
		instances = append(instances, instance)
	}
	return instances
}

func convertTags(tags []awstypes.Tag) map[string]string {
	labels := make(map[string]string, len(tags))
	for _, t := range tags {
		labels[convertString(t.Key)] = convertString(t.Value)
	}
	return labels
}

func convertState(state *awstypes.InstanceState) string {
	if state == nil {
		return ""
	}
	return string(state.Name)
}

func getTagName(tags []awstypes.Tag) *string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return tag.Value
		}
	}
	return nil
}

func convertClusterTags(tags map[string]*string) map[string]string {
	m := make(map[string]string, len(tags))
	for k, v := range tags {
		m[k] = convertString(v)
	}
	return m
}
