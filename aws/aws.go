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

	//"github.com/app-net-interface/awi-infra-guard/connector/aws"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"

	"github.com/app-net-interface/awi-infra-guard/types"
)

const providerName = "AWS"

type Client struct {
	defaultRegion    string
	defaultAccountID string
	accountID        string
	profiles         []types.Account
	clients          map[string]awsRegionalClientSet

	defaultAWSClient awsClient
	logger           *logrus.Logger
}

type awsRegionalClientSet map[string]awsClient

type awsClient struct {
	ec2Client *ec2.Client
	lbClient  *elasticloadbalancing.Client
	eksClient *eks.Client
}

func GetProfiles(ctx context.Context, cfg aws.Config, logger *logrus.Logger) []types.Account {

	profiles := make([]types.Account, 0)
	configFile, err := ini.Load(config.DefaultSharedCredentialsFilename())
	if err != nil {
		logger.Errorf("Failed to determine AWS config profiles, using default profile...")
		client := sts.NewFromConfig(cfg)
		req, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
		if err != nil {
			logger.Errorf("Failed to determine Account ID for defualt profile...")
		} else {
			profiles = []types.Account{{
				Name:     "default",
				ID:       convertString(req.Account),
				Provider: providerName,
			},
			}
		}
	} else {
		for _, v := range configFile.Sections() {
			if len(v.Keys()) == 0 {
				continue
			}
			profileName := v.Name()
			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region), config.WithSharedConfigProfile(profileName))
			if err != nil {
				logger.Errorf("Failed to load config for profile %s: %v", profileName, err)
				continue
			}
			client := sts.NewFromConfig(cfg)
			req, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
			if err != nil {
				logger.Errorf("Failed to determine Account ID for profile %s: %v", profileName, err)
				continue
			}
			profiles = append(profiles, types.Account{
				Name: profileName,
				ID:   convertString(req.Account),
			})
		}
	}
	return profiles
}

func NewClient(ctx context.Context, logger *logrus.Logger) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get AWS config: %v", err)
	}
	if cfg.Region == "" {
		reg := "us-east-1"
		logger.Warnf("Default AWS region is not specified, falling back to %s", reg)
		cfg.Region = reg
	}
	stsclient := sts.NewFromConfig(cfg)
	var defaultAccountID string
	req, err := stsclient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		logger.Errorf("Failed to determine Account ID for default profile: %v", err)
	} else {
		defaultAccountID = convertString(req.Account)
	}

	profiles := GetProfiles(ctx, cfg, logger)

	client := ec2.NewFromConfig(cfg)
	lbClient := elasticloadbalancing.NewFromConfig(cfg)
	eksClient := eks.NewFromConfig(cfg)

	allRegions, err := client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}
	clients, err := getAllClients(ctx, allRegions.Regions, profiles)
	if err != nil {
		return nil, err
	}
	return &Client{
		defaultRegion:    cfg.Region,
		defaultAccountID: defaultAccountID,
		defaultAWSClient: awsClient{
			ec2Client: client,
			lbClient:  lbClient,
			eksClient: eksClient,
		},
		clients:  clients,
		logger:   logger,
		profiles: profiles,
	}, nil
}

func (c *Client) GetName() string {
	return providerName
}

func (c *Client) ListAccounts() []types.Account {
	accounts := make([]types.Account, 0, len(c.profiles))
	for _, v := range c.profiles {
		accounts = append(accounts, v)
	}
	return accounts
}

func useDefaultConfig(region, defaultRegion, profile string) bool {
	if (region == "" || region == defaultRegion) && profile == "" {
		return true
	}
	return false
}

func (c *Client) getClientsForProfileAndRegion(profile, region string) (awsClient, error) {
	if profile == "" {
		profile = "default"
	}
	clients, ok := c.clients[profile]
	if !ok {
		found := false
		for _, v := range c.profiles {
			if v.ID == profile || v.Name == profile {
				clients, ok = c.clients[v.Name]
				if !ok {
					return awsClient{}, fmt.Errorf("couldn't find client configuration for profile: %s", profile)
				}
				found = true
			}
		}
		if !found {
			return awsClient{}, fmt.Errorf("couldn't find client configuration for profile: %s", profile)
		}
	}
	if region == "" {
		region = c.defaultRegion
	}
	regionalClient, ok := clients[region]
	if !ok {
		return awsClient{}, fmt.Errorf("couldn't find client configuration for region: %s", region)
	}
	return regionalClient, nil
}

func (c *Client) getAllClientsForProfile(profile string) (awsRegionalClientSet, error) {
	if profile == "" {
		profile = "default"
	}
	clients, ok := c.clients[profile]
	if !ok {
		found := false
		for _, v := range c.profiles {
			if v.ID == profile || v.Name == profile {
				clients, ok = c.clients[v.Name]
				if !ok {
					return nil, fmt.Errorf("couldn't find client configuration for profile: %s", profile)
				}
				found = true
				return clients, nil
			}
		}
		if !found {
			return nil, fmt.Errorf("couldn't find client configuration for profile: %s", profile)
		}
	}
	return clients, nil
}

func (c *Client) getEC2Client(ctx context.Context, account, region string) (*ec2.Client, error) {
	if useDefaultConfig(region, c.defaultRegion, account) {
		return c.defaultAWSClient.ec2Client, nil
	}
	client, err := c.getClientsForProfileAndRegion(account, region)
	if err != nil {
		return nil, err
	}
	return client.ec2Client, nil
}

func (c *Client) getELBClient(ctx context.Context, account, region string) (*elasticloadbalancing.Client, error) {
	if useDefaultConfig(region, c.defaultRegion, account) {
		return c.defaultAWSClient.lbClient, nil
	}
	client, err := c.getClientsForProfileAndRegion(account, region)
	if err != nil {
		return nil, err
	}
	return client.lbClient, nil
}

func (c *Client) getEKSClient(ctx context.Context, account, region string) (*eks.Client, error) {
	if useDefaultConfig(region, c.defaultRegion, account) {
		return c.defaultAWSClient.eksClient, nil
	}
	client, err := c.getClientsForProfileAndRegion(account, region)
	if err != nil {
		return nil, err
	}
	return client.eksClient, nil
}

func (c *Client) getAllRegions(ctx context.Context) ([]awstypes.Region, error) {
	allRegions, err := c.defaultAWSClient.ec2Client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}
	return allRegions.Regions, nil
}

func (c *Client) GetVPCIDForCIDR(ctx context.Context, params *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	builder := newFilterBuilder()
	builder.withCIDR(params.GetCidr())
	subnets, err := c.getSubnets(ctx, params.GetAccountId(), params.GetRegion(), builder)
	if err != nil {
		return "", fmt.Errorf("could not find subnet: %v", err)
	}
	if len(subnets) != 1 {
		return "", fmt.Errorf("expected one subnet for %s found %d", params.GetCidr(), len(subnets))
	}
	return *subnets[0].VpcId, nil
}

func (c *Client) GetCIDRsForLabels(ctx context.Context, params *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	var cidrs map[string]bool
	condition := params.GetLabels()[types.ConditionLabel]
	if condition == types.AndCondition {
		var err error
		cidrs, err = c.andConditionSubnet(ctx, params.GetAccountId(), params.GetRegion(), params.GetLabels())
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		cidrs, err = c.orConditionSubnet(ctx, params.GetAccountId(), params.GetRegion(), params.GetLabels())
		if err != nil {
			return nil, err
		}
	}
	result := make([]string, 0, len(cidrs))
	for cidr := range cidrs {
		result = append(result, cidr)
	}
	return result, nil
}

func (c *Client) GetInstancesForLabels(ctx context.Context, params *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	awsInstances, err := c.getInstancesForLabels(ctx, params.GetAccountId(), params.GetRegion(), params.GetLabels(), params.GetVpcId())
	if err != nil {
		return nil, err
	}
	return convertInstances(c.defaultAccountID, c.defaultRegion, params.GetAccountId(), params.GetRegion(), awsInstances), nil
}

func (c *Client) GetIPsForLabels(ctx context.Context, params *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	var ips map[string]struct{}
	var instances []types.Instance
	condition := params.GetLabels()[types.ConditionLabel]
	if condition == types.OrCondition {
		var err error
		instances, err = c.orConditionInstance(ctx, params.GetAccountId(), params.GetRegion(), params.GetLabels())
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		instances, err = c.andConditionInstance(ctx, params.GetAccountId(), params.GetRegion(), params.GetLabels())
		if err != nil {
			return nil, err
		}
	}

	ips = make(map[string]struct{}, len(instances))
	for _, instance := range instances {
		ips[instance.PrivateIP] = struct{}{}
	}
	result := make([]string, 0, len(ips))
	for cidr := range ips {
		result = append(result, cidr)
	}
	return result, nil
}

func (c *Client) GetVPCIDWithTag(ctx context.Context, params *infrapb.GetVPCIDWithTagRequest) (string, error) {
	vpc, err := c.getVPCWithTag(ctx, params.GetAccountId(), params.GetRegion(), params.GetKey(), params.GetValue())
	if err != nil {
		return "", err
	}
	return *vpc.VpcId, nil
}

func (c *Client) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}

func (c *Client) andConditionSubnet(ctx context.Context, account, region string, labels map[string]string) (map[string]bool, error) {
	builder := newFilterBuilder()
	for k, v := range labels {
		if k == types.ConditionLabel {
			continue
		}
		builder.withTag(k, v)
	}
	subnets, err := c.getSubnets(ctx, account, region, builder)
	if err != nil {
		return nil, err
	}
	cidrs := make(map[string]bool, len(subnets))
	for _, subnet := range subnets {
		cidrs[*subnet.CidrBlock] = true
	}
	return cidrs, nil
}

func (c *Client) andConditionInstance(ctx context.Context, account, region string, labels map[string]string) ([]types.Instance, error) {
	builder := newFilterBuilder()
	for k, v := range labels {
		if k == types.ConditionLabel {
			continue
		}
		builder.withTag(k, v)
	}
	awsInstances, err := c.getInstances(ctx, account, region, builder)
	if err != nil {
		return nil, err
	}
	instances := convertInstances(c.defaultAccountID, c.defaultRegion, account, region, awsInstances)
	return instances, nil
}

func (c *Client) orConditionSubnet(ctx context.Context, account, region string, labels map[string]string) (map[string]bool, error) {
	cidrs := make(map[string]bool)
	for k, v := range labels {
		if k == types.ConditionLabel {
			continue
		}
		builder := newFilterBuilder()
		builder.withTag(k, v)
		subnets, err := c.getSubnets(ctx, account, region, builder)
		if err != nil {
			return nil, err
		}
		for _, subnet := range subnets {
			cidrs[*subnet.CidrBlock] = true
		}
	}
	return cidrs, nil
}

func (c *Client) orConditionInstance(ctx context.Context, account, region string, labels map[string]string) ([]types.Instance, error) {
	var instances []types.Instance
	for k, v := range labels {
		if k == types.ConditionLabel {
			continue
		}
		builder := newFilterBuilder()
		builder.withTag(k, v)
		awsInstances, err := c.getInstances(ctx, account, region, builder)
		if err != nil {
			return nil, err
		}
		// TODO this should handle duplicated (case when instance have multiple matching labels)
		instances = append(instances, convertInstances(c.defaultAccountID, c.defaultRegion, account, region, awsInstances)...)
	}
	return instances, nil
}

func (c *Client) getSubnets(ctx context.Context, account, region string, builder *filterBuilder) ([]awstypes.Subnet, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}
	input := &ec2.DescribeSubnetsInput{
		Filters: builder.build(),
	}
	subnets, err := client.DescribeSubnets(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("could not find subnet: %v", err)
	}
	return subnets.Subnets, nil
}

func (c *Client) getInstances(ctx context.Context, account, region string, builder *filterBuilder) ([]awstypes.Reservation, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}
	input := &ec2.DescribeInstancesInput{
		Filters: builder.build(),
	}
	instances, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("could not find instance: %v", err)
	}
	return instances.Reservations, nil
}

func (c *Client) getVPCWithTag(ctx context.Context, account, region, key, value string) (*awstypes.Vpc, error) {
	client, err := c.getEC2Client(ctx, account, region)
	if err != nil {
		return nil, err
	}
	builder := newFilterBuilder()
	builder.withTag(key, value)
	input := &ec2.DescribeVpcsInput{
		Filters: builder.build(),
	}
	vpcs, err := client.DescribeVpcs(ctx, input)
	if err != nil {
		return nil, err
	}
	if len(vpcs.Vpcs) != 1 {
		return nil, fmt.Errorf("expected one VPC with tag key: %s, value %s, found %d", key, value, len(vpcs.Vpcs))
	}
	return &vpcs.Vpcs[0], nil
}

func mapToTagSpecfication(m map[string]string, resourceType awstypes.ResourceType) awstypes.TagSpecification {
	tagsSpec := awstypes.TagSpecification{
		ResourceType: resourceType,
		Tags:         make([]awstypes.Tag, 0, len(m)),
	}
	for k, v := range m {
		tagsSpec.Tags = append(tagsSpec.Tags, awstypes.Tag{
			Key:   &k,
			Value: &v,
		})
	}
	return tagsSpec
}
