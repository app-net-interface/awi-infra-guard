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
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListInstances called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListInstances: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListInstances: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListInstances called for vpc %s and region %s ", operationalAccountID, params.VpcId, params.Region)

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())
	for k, v := range params.GetLabels() {
		builder.withTag(k, v)
	}
	builder.withAvailabilityZone(params.GetZone())
	filters := builder.build()
	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allInstances  []types.Instance
			allErrors     []error
			resultChannel = make(chan regionResult)
		)
		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListInstances: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListInstances: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListInstances: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getInstancesForRegion
				instances, err := c.getInstancesForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region:    regionName,
					instances: instances,
					err:       err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListInstances: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListInstances: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allInstances = append(allInstances, result.instances...)
			}
		}
		c.logger.Infof("[AccountID: %s] ListInstances: Found %d instances across %d regions", operationalAccountID, len(allInstances), len(regions))

		if len(allErrors) > 0 {
			return allInstances, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allInstances, nil
	}
	c.logger.Debugf("[AccountID: %s] ListInstances: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getInstancesForRegion
	return c.getInstancesForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getInstancesForRegion now accepts operationalAccountID
func (c *Client) getInstancesForRegion(ctx context.Context, regionName string, filters []awstypes.Filter, operationalAccountID string) ([]types.Instance, error) {
	c.logger.Debugf("[AccountID: %s] getInstancesForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getInstancesForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Call DescribeInstances operation
	resp, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: filters,
	})
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getInstancesForRegion: DescribeInstances failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Pass operationalAccountID as the 'account' parameter to convertInstances
	return convertInstances(c.defaultAccountID, c.defaultRegion, operationalAccountID, regionName, resp.Reservations), nil
}

// convertInstances's 'account' parameter will now be the operationalAccountID
func convertInstances(defaultAccount, defaultRegion, account, region string, reservations []awstypes.Reservation) []types.Instance {
	if region == "" {
		region = defaultRegion
	}
	// The 'account' parameter is now the operationalAccountID passed from getInstancesForRegion.
	// If it's empty, it falls back to defaultAccount, but ideally, operationalAccountID should always be set.
	if account == "" {
		account = defaultAccount
	}
	instances := make([]types.Instance, 0) // Initialize to avoid nil if no reservations/instances
	for _, reservation := range reservations {
		for _, inst := range reservation.Instances { // Iterate over all instances in a reservation
			name := getTagName(inst.Tags)
			instanceLink := fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", region, region, aws.ToString(inst.InstanceId))

			secGroups := make([]string, len(inst.SecurityGroups))
			for i, group := range inst.SecurityGroups {
				secGroups[i] = aws.ToString(group.GroupId)
			}

			networkInterfaces := make([]string, len(inst.NetworkInterfaces))
			for j, iface := range inst.NetworkInterfaces {
				networkInterfaces[j] = aws.ToString(iface.NetworkInterfaceId)
			}

			instance := types.Instance{
				ID:               aws.ToString(inst.InstanceId),
				Name:             aws.ToString(name), // name is already *string or nil
				PrivateIP:        aws.ToString(inst.PrivateIpAddress),
				PublicIP:         aws.ToString(inst.PublicIpAddress),
				SubnetID:         aws.ToString(inst.SubnetId),
				VPCID:            aws.ToString(inst.VpcId),
				Type:             string(inst.InstanceType), // Corrected: InstanceType is not a pointer
				Labels:           convertTags(inst.Tags),
				State:            convertState(inst.State),
				Region:           region,
				Zone:             aws.ToString(inst.Placement.AvailabilityZone),
				AccountID:        account, // This is now the operationalAccountID
				Provider:         providerName,
				SelfLink:         instanceLink,
				SecurityGroupIDs: secGroups,
				InterfaceIDs:     networkInterfaces,
				CreatedAt:        inst.LaunchTime,
				UpdatedAt:        inst.UsageOperationUpdateTime, // Assuming LaunchTime is the creation time; adjust if needed
			}
			instances = append(instances, instance)
		}
	}
	return instances
}

func convertTags(tags []awstypes.Tag) map[string]string {
	labels := make(map[string]string, len(tags))
	for _, t := range tags {
		// Assuming convertString handles nil pointers gracefully, similar to aws.ToString
		// If convertString is not defined or doesn't, use aws.ToString
		labels[aws.ToString(t.Key)] = aws.ToString(t.Value)
	}
	return labels
}

func convertState(state *awstypes.InstanceState) string {
	if state == nil {
		return ""
	}
	return string(state.Name)
}

func getTagName(tags []awstypes.Tag) *string { // Returns *string or nil
	for _, tag := range tags {
		if aws.ToString(tag.Key) == "Name" { // Use aws.ToString for safety with *string
			return tag.Value // tag.Value is *string
		}
	}
	return nil
}

func convertClusterTags(tags map[string]*string) map[string]string {
	m := make(map[string]string, len(tags))
	for k, v := range tags {
		// Assuming convertString handles nil pointers gracefully, similar to aws.ToString
		// If convertString is not defined or doesn't, use aws.ToString
		m[k] = aws.ToString(v)
	}
	return m
}

// If convertString is a helper you have, ensure it's robust.
// Otherwise, using aws.ToString is generally safer for AWS SDK types.
// For example:
/*
func convertString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}
*/
