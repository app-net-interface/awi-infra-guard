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

	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func getAllClients(ctx context.Context, allRegions []awstypes.Region, allProfiles []types.Account) (clients map[string]awsRegionalClientSet, err error) {
	clients = make(map[string]awsRegionalClientSet, len(allProfiles))
	for _, account := range allProfiles {
		clientSet := make(map[string]awsClient, len(allRegions))
		for _, region := range allRegions {
			regionName := convertString(region.RegionName)
			builder := &configBuilder{}
			builder.withProfile(account.Name)
			builder.withRegion(regionName)
			cfg, err := config.LoadDefaultConfig(ctx, builder.build()...)
			if err != nil {
				return nil, err
			}
			clientSet[regionName] = awsClient{
				ec2Client:   ec2.NewFromConfig(cfg),
				lbClient:    elasticloadbalancing.NewFromConfig(cfg),
				eksClient:   eks.NewFromConfig(cfg),
				elbv2Client: elbv2.NewFromConfig(cfg),
			}
		}
		clients[account.Name] = clientSet
	}

	return clients, nil
}

type configBuilder struct {
	LoadConfigFuns []func(*config.LoadOptions) error
}

func (b *configBuilder) withRegion(region string) {
	if region != "" {
		b.LoadConfigFuns = append(b.LoadConfigFuns, config.WithRegion(region))
	}
}

func (b *configBuilder) withProfile(profile string) {
	if profile != "" {
		b.LoadConfigFuns = append(b.LoadConfigFuns, config.WithSharedConfigProfile(profile))
	}
}

func (b *configBuilder) build() []func(*config.LoadOptions) error {
	return b.LoadConfigFuns
}
