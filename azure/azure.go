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

package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

const providerName = "Azure"

type Client struct {
	cred       *azidentity.DefaultAzureCredential
	logger     *logrus.Logger
	vnetClient **armnetwork.VirtualNetworksClient
}

// NewClient initializes a new Azure client with all necessary clients for compute, network, and subscriptions.
func NewClient(ctx context.Context, logger *logrus.Logger) (*Client, error) {
	// Subscription ID from environment variable
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to obtain a credential:", err)
		return nil, err
	}
	client := &Client{
		cred:   cred,
		logger: logger,
	}

	return client, nil
}



func (c *Client) GetName() string {
	return providerName
}

func (c *Client) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}

func (c *Client) GetSubnet(ctx context.Context, input *infrapb.GetSubnetRequest) (types.Subnet, error) {
	// TBD
	return types.Subnet{}, nil
}

func (c *Client) GetVPCIDForCIDR(ctx context.Context, input *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) GetCIDRsForLabels(ctx context.Context, input *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetIPsForLabels(ctx context.Context, input *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetInstancesForLabels(ctx context.Context, input *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetVPCIDWithTag(ctx context.Context, input *infrapb.GetVPCIDWithTagRequest) (string, error) {
	// TBD
	return "", nil
}

/*
func getSubscriptionToken(ctx context.Context, subscriptionID string, credential *auth.DefaultAzureCredential) (string, error) {
    // Create a subscription-specific credential using azidentity
    subscriptionCredential, err := auth.NewSubscriptionCredential(credential, subscriptionID)
    if err != nil {
        return "", fmt.Errorf("failed to create subscription credential: %w", err)
    }

    // Acquire a token for the subscription
    token, err := subscriptionCredential.GetToken(ctx, "https://management.azure.com")
    if err != nil {
        return "", fmt.Errorf("failed to get token for subscription %s: %w", subscriptionID, err)
    }

    return token.AccessToken, nil
}

func getSubscriptionFactory (ctx context.Context) (subs []*armsubscriptions.Subscription, error) {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
	}
}

func getToken(ctx context.Context) (map[string]string, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to obtain a credential:", err)
		return
	}

	subscriptionsClient, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		fmt.Println("Failed to create subscriptions client:", err)
		return
	}

	ctx := context.Background()

	pager := subscriptionsClient.NewListPager(nil)

	fmt.Println("Listing all VNets across all subscriptions:")

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			fmt.Println("Failed to get the next page of subscriptions:", err)
			return
		}
		for _, sub := range resp.SubscriptionListResult.Value {
			fmt.Printf("Subscription: %s\n", *sub.SubscriptionID)
		}
	}
}*/
