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

	//compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"google.golang.org/api/iterator"

	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) GetSubnet(ctx context.Context, params *infrapb.GetSubnetRequest) (types.Subnet, error) {
	if params.GetVpcId() == "" || params.GetId() == "" {
		return types.Subnet{}, fmt.Errorf("wrong paramters of GetSubnet function, vpcID and id are expected")
	}
	net, err := c.vpcIdToSingleNetwork(ctx, params.GetAccountId(), params.GetVpcId())
	if err != nil {
		return types.Subnet{}, err
	}
	iter := c.subnetsClient.AggregatedList(ctx, &computepb.AggregatedListSubnetworksRequest{
		Project: net.project,
	})

	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: params.GetAccountId()})
	if err != nil {
		return types.Subnet{}, err
	}

	for {
		pair, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return types.Subnet{}, err
		}
		for _, gcpSubnet := range pair.Value.Subnetworks {
			if gcpSubnet.GetName() == params.GetId() || strconv.FormatUint(gcpSubnet.GetId(), 10) == params.GetId() {
				return convertSubnet(net.project, networks, gcpSubnet), nil
			}
		}
	}
	return types.Subnet{}, fmt.Errorf("could not find subnet %s", params.GetId())
}

func (c *Client) ListSubnets(ctx context.Context, request *infrapb.ListSubnetsRequest) ([]types.Subnet, error) {
	filter := newFilterBuilder()
	if len(request.GetLabels()) > 0 {
		return nil, fmt.Errorf("GCP subnets don't support tags or labels")
	}
	if request.GetVpcId() != "" {
		net, err := c.vpcIdToSingleNetwork(ctx, request.GetAccountId(), request.GetVpcId())
		if err != nil {
			return nil, err
		}
		filter.withNetwork(net.fullUrl)
	}
	if request.GetCidr() != "" {
		filter.withIPCIDRRange(request.GetCidr())
	}
	// subnets in GCP have regional scope
	if request.GetRegion() != "" {
		filter.withRegion(request.GetRegion())
	}

	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: request.GetAccountId()})
	if err != nil {
		return nil, err
	}

	subnets := make([]types.Subnet, 0)
	f := func(projectId string) error {
		iter := c.subnetsClient.AggregatedList(ctx, &computepb.AggregatedListSubnetworksRequest{
			Filter:  filter.build(),
			Project: projectId,
		})

		for {
			pair, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			for _, gcpSubnet := range pair.Value.Subnetworks {
				subnets = append(subnets, convertSubnet(projectId, networks, gcpSubnet))
			}
		}
		return nil
	}
	if request.GetAccountId() == "" {
		for projectId := range c.projectIDs {
			err := f(projectId)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := c.checkProject(request.GetAccountId())
		if err != nil {
			return nil, err
		}
		err = f(request.GetAccountId())
		if err != nil {
			return nil, err
		}
	}
	return subnets, nil
}

func (c *Client) getSubnetsByNetworkAndCidr(ctx context.Context, net network, cidr string) ([]*computepb.Subnetwork, error) {
	filter := newFilterBuilder()
	if net.url != "" {
		filter.withNetwork(net.url)
	}
	if cidr != "" {
		filter.withIPCIDRRange(cidr)
	}
	iter := c.subnetsClient.AggregatedList(ctx, &computepb.AggregatedListSubnetworksRequest{
		Filter:  filter.build(),
		Project: net.project,
	})

	var subnets []*computepb.Subnetwork
	for {
		pair, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, gcpSubnet := range pair.Value.Subnetworks {
			subnets = append(subnets, gcpSubnet)
		}
	}
	return subnets, nil
}

func (c *Client) GetCIDRsForLabels(_ context.Context, _ *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	return nil, fmt.Errorf("GCP subnets don't support labels")
}

func convertSubnet(projectID string, networks []types.VPC, subnetwork *computepb.Subnetwork) types.Subnet {
	if subnetwork == nil {
		return types.Subnet{}
	}

	subnet := types.Subnet{
		SubnetId:  strconv.FormatUint(subnetwork.GetId(), 10),
		CidrBlock: subnetwork.GetIpCidrRange(),
		Name:      subnetwork.GetName(),
		AccountID: projectID,
		Provider:  providerName,
	}

	network := strings.Split(subnetwork.GetNetwork(), "/")
	if len(network) != 0 {
		name := network[len(network)-1]
		for _, v := range networks {
			if v.Name == name || v.ID == name {
				subnet.VpcId = v.ID
				break
			}
		}
	}
	region := strings.Split(subnetwork.GetRegion(), "/")
	if len(region) != 0 {
		subnet.Zone = region[len(region)-1]
	}

	return subnet
}
