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

func (c *Client) ListKeyPairs(ctx context.Context, input *infrapb.ListKeyPairsRequest) ([]types.KeyPair, error) {
	c.creds = input.Creds
	c.accountID = input.AccountId

	builder := newFilterBuilder()
	builder.withVPC(input.GetVpcId())
	for k, v := range input.GetLabels() {
		builder.withTag(k, v)
	}
	filters := builder.build()

	if input.GetRegion() == "" || input.GetRegion() == "all" {
		var (
			allKeyPairs   []types.KeyPair
			allErrors     []error
			wg            sync.WaitGroup
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
				regKeyPairs, err := c.getKeyPairsForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					kps:    regKeyPairs,
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
				allKeyPairs = append(allKeyPairs, result.kps...)
			}
		}

		c.logger.Infof("In account %s Found %d KeyPairs across %d regions", c.accountID, len(allKeyPairs), len(regions))

		if len(allErrors) > 0 {
			return allKeyPairs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		PrintResources(allKeyPairs, "types.KeyPair")

		return allKeyPairs, nil
	}

	return c.getKeyPairsForRegion(ctx, input.Region, filters)
}

func (c *Client) getKeyPairsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter) ([]types.KeyPair, error) {
	ec2Client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}

	input := &ec2.DescribeKeyPairsInput{
		Filters: filters,
	}

	output, err := ec2Client.DescribeKeyPairs(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error describing key pairs: %v", err)
	}
	labels := make(map[string]string)

	keyPairs := make([]types.KeyPair, len(output.KeyPairs))
	for i, kp := range output.KeyPairs {
		keyName := aws.ToString(kp.KeyName)
		for _, tag := range kp.Tags {
			if *tag.Key == "Name" || *tag.Key == "name" {
				if kp.KeyName == nil && tag.Value != nil {
					keyName = *tag.Value
				}
			}
			labels[*tag.Key] = *tag.Value
		}
		
		keyPairs[i] = types.KeyPair{
			ID:          aws.ToString(kp.KeyPairId),
			Name:        keyName,
			Fingerprint: aws.ToString(kp.KeyFingerprint),
			PublicKey:   aws.ToString(kp.PublicKey),
			CreatedAt:   aws.ToTime(kp.CreateTime),
			Labels:      labels,
			Provider:    c.GetName(),
			Region:      regionName,
			AccountID:   c.accountID,
			KeyPairType: string(kp.KeyType),
		}
	}

	return keyPairs, nil
}
