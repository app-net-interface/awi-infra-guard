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
	"fmt" // Added for panic statement

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (c *Client) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {
	// Panic statement for debugging
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListRegions called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		// If no accountId in params, consider using the client's default/initial accountID.
		c.logger.Infof("[AccountID: %s] ListRegions: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListRegions: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}

	c.logger.Debugf("[AccountID: %s] ListRegions called. HasCreds: %t",
		operationalAccountID, params.GetCreds() != nil)

	// DO NOT MODIFY c.accountID or c.creds here
	// REMOVED: c.accountID = params.AccountId
	// REMOVED: c.creds = params.Creds

	// Call getAllRegions, passing the operationalAccountID.
	// This assumes c.getAllRegions returns []awsSdkEc2Types.Region (aliased as awstypes.Region in other files)
	awsSDKRegions, err := c.getAllRegions(ctx, operationalAccountID)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] ListRegions: Error from getAllRegions: %v", operationalAccountID, err)
		return nil, err // Return the error from getAllRegions
	}

	// Convert AWS SDK regions to []types.Region
	var appRegions []types.Region
	for _, sdkRegion := range awsSDKRegions { // awsSDKRegions is []awsSdkEc2Types.Region
		if sdkRegion.RegionName == nil {
			c.logger.Warnf("[AccountID: %s] ListRegions: Found an SDK region with a nil RegionName. Skipping.", operationalAccountID)
			continue
		}
		appRegions = append(appRegions, types.Region{
			ID:        *sdkRegion.RegionName,
			Name:      *sdkRegion.RegionName,
			Provider:  providerName,           // Assuming providerName is a defined constant or variable in the package
			AccountID: operationalAccountID, // Set the operational account ID
		})
	}

	c.logger.Debugf("[AccountID: %s] Found %d regions enabled. Regions are: %+v", operationalAccountID, len(appRegions), appRegions)
	return appRegions, nil // Return the converted regions
}

// getAllRegions is assumed to be defined elsewhere in the package like this:
// func (c *Client) getAllRegions(ctx context.Context, operationalAccountID string) ([]awsSdkEc2Types.Region, error)
// It should use the operationalAccountID to get an EC2 client for DescribeRegions.
