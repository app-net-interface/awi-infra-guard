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
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListKeyPairs(ctx context.Context, params *infrapb.ListKeyPairsRequest) ([]types.KeyPair, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListKeyPairs called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListKeyPairs: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListKeyPairs: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListKeyPairs called with VPC ID: %s, Region: %s", operationalAccountID, params.GetVpcId(), params.GetRegion())

	// REMOVED: c.creds = params.Creds
	// REMOVED: c.accountID = params.AccountId

	builder := newFilterBuilder()
	// KeyPairs are not directly associated with VPCs in AWS EC2 in a filterable way.
	// If VpcId is provided, it's likely for context or future use, but DescribeKeyPairs doesn't filter by VPC.
	// builder.withVPC(params.GetVpcId()) // This filter won't apply to DescribeKeyPairs
	for k, v := range params.GetLabels() {
		builder.withTag(k, v) // KeyPairs can be filtered by tags
	}
	filters := builder.build()

	if params.GetRegion() == "" || params.GetRegion() == "all" {
		var (
			allKeyPairs   []types.KeyPair
			allErrors     []error
			wg            sync.WaitGroup
			resultChannel = make(chan regionResult)
		)

		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListKeyPairs: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListKeyPairs: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListKeyPairs: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getKeyPairsForRegion
				regKeyPairs, err := c.getKeyPairsForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					kps:    regKeyPairs,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListKeyPairs: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListKeyPairs: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allKeyPairs = append(allKeyPairs, result.kps...)
			}
		}

		c.logger.Infof("[AccountID: %s] ListKeyPairs: Found %d KeyPairs across %d regions", operationalAccountID, len(allKeyPairs), len(regions))

		if len(allErrors) > 0 {
			return allKeyPairs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allKeyPairs, nil
	}
	c.logger.Debugf("[AccountID: %s] ListKeyPairs: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getKeyPairsForRegion
	return c.getKeyPairsForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getKeyPairsForRegion now accepts operationalAccountID
func (c *Client) getKeyPairsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter, operationalAccountID string) ([]types.KeyPair, error) {
	c.logger.Debugf("[AccountID: %s] getKeyPairsForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	ec2Client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getKeyPairsForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}

	input := &ec2.DescribeKeyPairsInput{
		Filters: filters,
		// IncludePublicKey: aws.Bool(true), // Uncomment if you need the public key material
	}

	output, err := ec2Client.DescribeKeyPairs(ctx, input)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getKeyPairsForRegion: DescribeKeyPairs failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, fmt.Errorf("error describing key pairs in region %s for account %s: %v", regionName, operationalAccountID, err)
	}

	keyPairs := make([]types.KeyPair, 0, len(output.KeyPairs)) // Initialize with 0 length, capacity len(output.KeyPairs)
	for _, kp := range output.KeyPairs {
		keyName := aws.ToString(kp.KeyName)
		labels := make(map[string]string) // Initialize labels for each key pair
		for _, tag := range kp.Tags {
			// KeyName from tag is usually not how AWS works for KeyPairs, KeyName is a primary identifier.
			// If a "Name" tag exists, it's just a tag.
			labels[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
		}
		// If KeyName was empty but a "Name" tag exists, you might choose to use it.
		// However, kp.KeyName should generally be populated.
		if keyName == "" {
			if nameTagVal, ok := labels["Name"]; ok {
				keyName = nameTagVal
			} else if nameTagVal, ok := labels["name"]; ok {
				keyName = nameTagVal
			}
		}

		keyPairs = append(keyPairs, types.KeyPair{
			ID:          aws.ToString(kp.KeyPairId),
			Name:        keyName,
			Fingerprint: aws.ToString(kp.KeyFingerprint),
			PublicKey:   aws.ToString(kp.PublicKey), // Only populated if IncludePublicKey was true in input
			CreatedAt:   aws.ToTime(kp.CreateTime),
			Labels:      labels,
			Provider:    c.GetName(),
			Region:      regionName,
			AccountID:   operationalAccountID, // Use operationalAccountID
			KeyPairType: string(kp.KeyType),   // kp.KeyType is of type types.KeyType (e.g., "rsa", "ed25519")
		})
	}

	return keyPairs, nil
}
