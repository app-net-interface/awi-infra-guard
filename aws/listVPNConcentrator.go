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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPNConcentrators(ctx context.Context, params *infrapb.ListVPNConcentratorsRequest) ([]types.VPNConcentrator, error) {
	c.creds = params.Creds
	c.accountID = params.AccountId
	if c.creds != nil && c.creds.GetRoleBasedAuth() != nil && c.creds.GetRoleBasedAuth().GetAwsRole().RoleArn != "" {
		c.accountID = ExtractAccountID(c.creds.GetRoleBasedAuth().GetAwsRole().RoleArn)
	}

	var regions []string

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		allRegions, err := c.getAllRegions(ctx)
		if err != nil {
			return nil, err
		}
		for _, region := range allRegions {
			regions = append(regions, *region.RegionName)
		}
	} else {
		regions = []string{params.GetRegion()}
	}

	resultChan := make(chan regionResult, len(regions))

	for _, region := range regions {
		go func(region string) {
			vpncs, err := c.listRegionalVPNConcentrator(ctx, c.accountID, region)
			resultChan <- regionResult{
				region: region,
				vpncs:  vpncs,
				err:    err,
			}
		}(region)
	}

	var vpnConcentrators []types.VPNConcentrator
	for i := 0; i < len(regions); i++ {
		result := <-resultChan
		if result.err != nil {
			c.logger.Errorf("Error listing VPN concentrators in region %s: %v", result.region, result.err)
			continue
		}
		vpnConcentrators = append(vpnConcentrators, result.vpncs...)
	}

	return vpnConcentrators, nil
}

func (c *Client) listRegionalVPNConcentrator(ctx context.Context, account, region string) ([]types.VPNConcentrator, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}

	input := &ec2.DescribeVpnGatewaysInput{}
	output, err := client.DescribeVpnGateways(ctx, input)
	if err != nil {
		return nil, err
	}

	vpnConcentrators := make([]types.VPNConcentrator, 0, len(output.VpnGateways))
	for _, vpnGateway := range output.VpnGateways {
		vpnConcentrators = append(vpnConcentrators, convertVPNConcentrator(vpnGateway, account, region))
	}

	return vpnConcentrators, nil
}

func convertVPNConcentrator(vpnGateway awstypes.VpnGateway, account, region string) types.VPNConcentrator {
	vpnc := types.VPNConcentrator{}
	labels := make(map[string]string)

	for _, tag := range vpnGateway.Tags {
		labels[*tag.Key] = *tag.Value
	}

	vpnc = types.VPNConcentrator{
		ID:        *vpnGateway.VpnGatewayId,
		AccountID: account,
		Region:    region,
		Provider:  providerName,
		Type:      string(vpnGateway.Type),
		State:     string(vpnGateway.State),
		Labels:    labels,
	}

	if vpnGateway.AmazonSideAsn != nil {
		vpnc.ASN = *vpnGateway.AmazonSideAsn
	}
	if len(vpnGateway.VpcAttachments) > 0 {
		vpnc.VpcID = *vpnGateway.VpcAttachments[0].VpcId
	}

	return vpnc
}
