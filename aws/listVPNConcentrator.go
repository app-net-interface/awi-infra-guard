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
	"github.com/aws/aws-sdk-go-v2/aws" // Import aws package
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListVPNConcentrators(ctx context.Context, params *infrapb.ListVPNConcentratorsRequest) ([]types.VPNConcentrator, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListVPNConcentrators called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListVPNConcentrators: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListVPNConcentrators: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Infof("[AccountID: %s] ListVPNConcentrators called. Region: %s", operationalAccountID, params.GetRegion())

	var regionsToProcess []string

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		// Pass operationalAccountID to getAllRegions
		allRegionsSDK, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListVPNConcentrators: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		for _, region := range allRegionsSDK {
			regionsToProcess = append(regionsToProcess, aws.ToString(region.RegionName))
		}
	} else {
		regionsToProcess = []string{params.GetRegion()}
	}
	c.logger.Debugf("[AccountID: %s] ListVPNConcentrators: Processing regions: %v", operationalAccountID, regionsToProcess)

	resultChan := make(chan regionResult, len(regionsToProcess))
	var wg sync.WaitGroup

	for _, region := range regionsToProcess {
		wg.Add(1)
		// Pass operationalAccountID to the goroutine
		go func(regionName string, accID string) {
			defer wg.Done()
			c.logger.Debugf("[AccountID: %s] ListVPNConcentrators: Goroutine for region %s started.", accID, regionName)
			// Pass accID (operationalAccountID) to listRegionalVPNConcentrator
			vpncs, err := c.listRegionalVPNConcentrator(ctx, accID, regionName)
			resultChan <- regionResult{
				region: regionName,
				vpncs:  vpncs,
				err:    err,
			}
		}(region, operationalAccountID)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		c.logger.Debugf("[AccountID: %s] ListVPNConcentrators: All region goroutines finished, resultChannel closed.", operationalAccountID)
	}()

	var vpnConcentrators []types.VPNConcentrator
	for result := range resultChan {
		if result.err != nil {
			c.logger.Errorf("[AccountID: %s] Error listing VPN concentrators in region %s: %v", operationalAccountID, result.region, result.err)
			// Decide if you want to collect errors and return them, or just log and continue
			continue
		}
		vpnConcentrators = append(vpnConcentrators, result.vpncs...)
	}
	c.logger.Infof("[AccountID: %s] Found %d VPN Concentrators across processed regions.", operationalAccountID, len(vpnConcentrators))

	return vpnConcentrators, nil
}

func (c *Client) listRegionalVPNConcentrator(ctx context.Context, operationalAccountID, region string) ([]types.VPNConcentrator, error) {
	c.logger.Debugf("[AccountID: %s] listRegionalVPNConcentrator: Region %s", operationalAccountID, region)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, region)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] listRegionalVPNConcentrator: Failed to get EC2 client for region %s: %v", operationalAccountID, region, err)
		return nil, err
	}

	input := &ec2.DescribeVpnGatewaysInput{}
	output, err := client.DescribeVpnGateways(ctx, input)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] listRegionalVPNConcentrator: DescribeVpnGateways failed for region %s: %v", operationalAccountID, region, err)
		return nil, err
	}

	vpnConcentrators := make([]types.VPNConcentrator, 0, len(output.VpnGateways))
	for _, vpnGateway := range output.VpnGateways {
		// Pass operationalAccountID to convertVPNConcentrator
		vpnConcentrators = append(vpnConcentrators, convertVPNConcentrator(vpnGateway, operationalAccountID, region))
	}
	c.logger.Debugf("[AccountID: %s] listRegionalVPNConcentrator: Found %d VPN Gateways in region %s.", operationalAccountID, len(vpnConcentrators), region)
	return vpnConcentrators, nil
}

func convertVPNConcentrator(vpnGateway awstypes.VpnGateway, operationalAccountID, region string) types.VPNConcentrator {
	labels := make(map[string]string)
	for _, tag := range vpnGateway.Tags {
		labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
	}

	vpnc := types.VPNConcentrator{
		ID:        aws.ToString(vpnGateway.VpnGatewayId),
		AccountID: operationalAccountID, // Use operationalAccountID
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
		// A VGW can be attached to only one VPC at a time.
		vpnc.VpcID = aws.ToString(vpnGateway.VpcAttachments[0].VpcId)
	}
	// The Printf statement is for debugging, consider removing or using logger in production
	// fmt.Printf("VPN Concentrator: %v\n", vpnc)

	return vpnc
}
