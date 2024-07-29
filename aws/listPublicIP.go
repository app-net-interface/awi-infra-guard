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
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListPublicIPs(ctx context.Context, params *infrapb.ListPublicIPsRequest) ([]types.PublicIP, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId
	accountId := params.GetAccountId()

	if c.accountID == "" {
		accountId = c.defaultAccountID
	}

	builder := newFilterBuilder()
	filters := builder.build()

	c.logger.Infof("*******Requesting Public IP for account id %s ****", accountId)
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allPublicIPs  []types.PublicIP
			allErrors     []error
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
				pips, err := c.getPublicIPForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					pips:   pips,
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
				allPublicIPs = append(allPublicIPs, result.pips...)
			}
		}

		if len(allErrors) > 0 {
			return allPublicIPs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allPublicIPs, nil
	}
	return c.getPublicIPForRegion(ctx, params.Region, filters)
}

func (c *Client) getPublicIPForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.PublicIP, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeVpcs operation
	resp, err := client.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{
		Filters: filters,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get AWS public IPs: %v", err)
	}
	builder := newFilterBuilder()
	builder.filters = filters
	reservations, err := c.getInstances(ctx, c.accountID, regionName, builder)
	if err != nil {
		c.logger.Warnf("Failed to get instances %s", err)
	}
	return c.convertPublicIPs(c.accountID, regionName, resp.Addresses, reservations)
}

func (c *Client) convertPublicIPs(account, region string, addresses []awstypes.Address, reservations []awstypes.Reservation) ([]types.PublicIP, error) {

	staticIPs := make([]types.PublicIP, 0, len(addresses))
	var instancePublicIPs []types.PublicIP
	var allPublicIPs []types.PublicIP
	for _, address := range addresses {

		staticIPs = append(staticIPs, types.PublicIP{
			ID:         aws.ToString(address.AllocationId),
			Region:     region,
			InstanceId: aws.ToString(address.InstanceId),
			PublicIP:   aws.ToString(address.PublicIp),
			Provider:   providerName,
			AccountID:  c.accountID,
			Type:       "static",
			Labels:     convertTags(address.Tags),
			PrivateIP:  aws.ToString(address.PrivateIpAddress),
			SelfLink:   fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#ElasticIpDetails:AllocationId=%s", region, region, aws.ToString(address.AllocationId)),
		})
	}

	for _, reservation := range reservations {
		if len(reservation.Instances) == 0 {
			continue
		}
		inst := reservation.Instances[0]
		publicIP := aws.ToString(inst.PublicIpAddress)

		var vpcId string
		if inst.VpcId != nil {
			vpcId = *inst.VpcId
		}

		if publicIP != "" {
			instancePublicIPs = append(instancePublicIPs, types.PublicIP{
				ID:         "0",
				VPCID:      vpcId,
				Region:     region,
				InstanceId: aws.ToString(inst.InstanceId),
				PublicIP:   publicIP,
				PrivateIP:  aws.ToString(inst.PrivateIpAddress),
				Provider:   providerName,
				AccountID:  account,
				Type:       "ephimeral",
			})
		}
	}
	allPublicIPs = mergePublicIPs(staticIPs, instancePublicIPs)

	/*
		for _, staticIP := range staticIPs {
			fmt.Printf("******* Elastic ip %v VPC %s Region %s Instance %s Account %s ****** \n", staticIP.PublicIP, staticIP.VPCID, staticIP.Region, staticIP.InstanceId, staticIP.AccountID)
		}
		for _, instancePublicIP := range instancePublicIPs {
			fmt.Printf("******* Instance Public IP %v Region %s VPC %s Instance %s Account %s ****** \n", instancePublicIP.PublicIP, instancePublicIP.VPCID, instancePublicIP.Region, instancePublicIP.InstanceId, instancePublicIP.AccountID)
		}


		for _, publicIP := range allPublicIPs {
			fmt.Printf("******* After Merge Public IP %v VPC %s Region %s Instance %s Account %s ****** \n", publicIP.PublicIP, publicIP.VPCID, publicIP.Region, publicIP.InstanceId, publicIP.AccountID)
		}
	*/

	return allPublicIPs, nil
}

// Merge static(elastic) and dynamic(instance) public IPs
func mergePublicIPs(elasticIPs []types.PublicIP, instanceIPs []types.PublicIP) []types.PublicIP {
	// Create a map to keep track of unique items in slice1 by Name
	itemMap := make(map[string]types.PublicIP)

	// Populate the map with items from slice1
	for _, item := range elasticIPs {
		itemMap[item.PublicIP] = item
	}

	// Process items in slice2
	var remainingSlice []types.PublicIP
	for _, instanceIP := range instanceIPs {
		if existing, found := itemMap[instanceIP.PublicIP]; found {
			// If duplicate is found, update the entry
			//fmt.Printf("*****Found matching entry for %s: --%s--%s \n", existing.PublicIP, existing.VPCID, instanceIP.VPCID)
			existing.VPCID = instanceIP.VPCID
			existing.Region = instanceIP.Region
			existing.PrivateIP = instanceIP.PrivateIP
			itemMap[instanceIP.PublicIP] = existing

		} else {
			remainingSlice = append(remainingSlice, instanceIP)
		}
	}

	// Convert the map back to a slice for slice1
	elasticIPs = []types.PublicIP{}
	for _, item := range itemMap {
		//fmt.Printf("Account Id = %s Public IP = %s VPC ID = %s \n", item.AccountID, item.PublicIP, item.VPCID)
		elasticIPs = append(elasticIPs, item)
	}

	// Append all remaining unique entries of slice2 to slice1
	elasticIPs = append(elasticIPs, remainingSlice...)
	return elasticIPs
}
