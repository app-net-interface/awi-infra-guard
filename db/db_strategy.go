// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
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

package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/app-net-interface/kubernetes-discovery/cluster"
)

type StrategyWithDB struct {
	logger             *logrus.Logger
	cloudProviders     map[string]*providerWithDB
	kubernetesProvider provider.Kubernetes
	realStrategy       provider.Strategy
}

func NewStrategyWithDB(dbClient Client, providerStrategy provider.Strategy, logger *logrus.Logger) *StrategyWithDB {
	strategy := &StrategyWithDB{
		cloudProviders: make(map[string]*providerWithDB),
		realStrategy:   providerStrategy,
		logger:         logger,
	}
	for _, p := range providerStrategy.GetAllProviders() {
		strategy.cloudProviders[strings.ToLower(p.GetName())] = &providerWithDB{
			realProvider: p,
			dbClient:     dbClient,
			logger:       logger,
		}
	}
	k8sProvider, _ := providerStrategy.GetKubernetesProvider()
	strategy.kubernetesProvider = &KubernetesProviderWithDB{
		realProvider: k8sProvider,
		dbClient:     dbClient,
	}
	return strategy
}

func (s *StrategyWithDB) GetProvider(ctx context.Context, cloud string) (provider.CloudProvider, error) {
	prov, ok := s.cloudProviders[strings.ToLower(cloud)]
	if !ok {
		return nil, fmt.Errorf("couldn't find provider %s", cloud)
	}
	s.logger.Infof("Using local DB for %s provider", cloud)
	return prov, nil
}

func (s *StrategyWithDB) GetAllProviders() []provider.CloudProvider {
	var providers []provider.CloudProvider
	for _, v := range s.cloudProviders {
		providers = append(providers, v)
	}
	return providers
}

func (s *StrategyWithDB) GetKubernetesProvider() (provider.Kubernetes, error) {
	return s.kubernetesProvider, nil
}

func (s *StrategyWithDB) RefreshState(ctx context.Context) error {
	return s.realStrategy.RefreshState(ctx)
}

type providerWithDB struct {
	realProvider provider.CloudProvider
	dbClient     Client
	logger       *logrus.Logger
}

func (p *providerWithDB) GetName() string {
	return p.realProvider.GetName()
}

func (p *providerWithDB) ListAccounts() []types.Account {
	return p.realProvider.ListAccounts()
}

func (p *providerWithDB) ListRegions(ctx context.Context, params *infrapb.ListRegionsRequest) ([]types.Region, error) {
	dbRegions, err := p.dbClient.ListRegions()
	if err != nil {
		return nil, err
	}
	var providersRegions []types.Region
	for _, region := range dbRegions {
		if strings.ToLower(region.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		providersRegions = append(providersRegions, *region)
	}
	return providersRegions, nil
}

func (p *providerWithDB) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {
	dbVPCs, err := p.dbClient.ListVPCs()
	if err != nil {
		return nil, err
	}
	var providersVPCs []types.VPC
	for _, vpc := range dbVPCs {
		if strings.ToLower(vpc.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != vpc.AccountID {
			continue
		}
		if params.GetRegion() != "" && params.GetRegion() != vpc.Region {
			continue
		}

		match := true
		for k, v := range params.GetLabels() {
			r, ok := vpc.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}

		if match {
			providersVPCs = append(providersVPCs, *vpc)
		}

	}
	return providersVPCs, nil
}

func (p *providerWithDB) GetVPCIndex(ctx context.Context, params *infrapb.GetVPCIndexRequest) (*types.VPCIndex, error) {
	dbVPCIndex, err := p.dbClient.GetVPCIndex(types.CloudID(p.realProvider.GetName(), params.VpcId))
	if err != nil {
		fmt.Printf("Error getting VPC index: %v", err)
		return nil, err
	}
	return dbVPCIndex, nil
}

func (p *providerWithDB) ListVPCIndex(ctx context.Context, params *infrapb.GetVPCIndexRequest) ([]types.VPCIndex, error) {
	dbVPCIndexes, err := p.dbClient.ListVPCIndex()
	if err != nil {
		return nil, err
	}
	var providersVPCIndexes []types.VPCIndex
	for _, vpcIndex := range dbVPCIndexes {
		if !strings.EqualFold(vpcIndex.Provider, p.realProvider.GetName()) {
			continue
		}
		providersVPCIndexes = append(providersVPCIndexes, *vpcIndex)
	}
	return providersVPCIndexes, nil
}

func (p *providerWithDB) ListInstances(ctx context.Context, params *infrapb.ListInstancesRequest) ([]types.Instance, error) {
	dbInstances, err := p.dbClient.ListInstances()
	if err != nil {
		return nil, err
	}
	var providersInstances []types.Instance
	for _, instance := range dbInstances {
		if strings.ToLower(instance.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != instance.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != instance.Region {
				continue
			}
		}
		if params.GetVpcId() != "" && params.GetVpcId() != instance.VPCID {
			continue
		}
		if params.GetZone() != "" && params.GetZone() != instance.Zone {
			continue
		}
		match := true
		for k, v := range params.GetLabels() {
			r, ok := instance.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}
		if match {
			providersInstances = append(providersInstances, *instance)
		}

	}
	return providersInstances, nil
}

func (p *providerWithDB) ListSubnets(ctx context.Context, params *infrapb.ListSubnetsRequest) ([]types.Subnet, error) {
	dbSubnets, err := p.dbClient.ListSubnets()
	if err != nil {
		return nil, err
	}
	var providersSubnets []types.Subnet
	for _, subnet := range dbSubnets {
		if strings.ToLower(subnet.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != subnet.AccountID {
			continue
		}
		if (params.GetRegion() != "" && params.GetRegion() != "global") && params.GetRegion() != subnet.Region {
			continue
		}
		if params.GetVpcId() != "" && params.GetVpcId() != subnet.VpcId {
			continue
		}
		if params.GetZone() != "" && params.GetZone() != subnet.Zone {
			continue
		}
		if params.GetCidr() != "" && params.GetCidr() != subnet.CidrBlock {
			continue
		}

		match := true
		for k, v := range params.GetLabels() {
			r, ok := subnet.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}

		if match {
			providersSubnets = append(providersSubnets, *subnet)
		}

	}
	return providersSubnets, nil
}

func (p *providerWithDB) ListACLs(ctx context.Context, params *infrapb.ListACLsRequest) ([]types.ACL, error) {
	dbACLs, err := p.dbClient.ListACLs()
	if err != nil {
		return nil, err
	}
	var providersACLs []types.ACL
	for _, acl := range dbACLs {
		if strings.ToLower(acl.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != acl.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != acl.Region {
				continue
			}
		}

		if params.GetVpcId() != "" && params.GetVpcId() != acl.VpcID {
			continue
		}

		providersACLs = append(providersACLs, *acl)

	}
	return providersACLs, nil
}
func (p *providerWithDB) ListSecurityGroups(ctx context.Context, params *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error) {
	dbSecurityGroups, err := p.dbClient.ListSecurityGroups()
	if err != nil {
		return nil, err
	}
	var providersSecurityGroups []types.SecurityGroup
	for _, securityGroup := range dbSecurityGroups {
		if strings.ToLower(securityGroup.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != securityGroup.AccountID {
			continue
		}
		if params.GetVpcId() != "" && params.GetVpcId() != securityGroup.VpcID {
			continue
		}

		providersSecurityGroups = append(providersSecurityGroups, *securityGroup)

	}
	return providersSecurityGroups, nil
}

func (p *providerWithDB) ListRouteTables(ctx context.Context, params *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error) {
	dbRouteTables, err := p.dbClient.ListRouteTables()
	if err != nil {
		return nil, err
	}
	var providersRouteTables []types.RouteTable
	for _, routeTable := range dbRouteTables {
		if strings.ToLower(routeTable.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != routeTable.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != routeTable.Region {
				continue
			}
		}
		if params.GetVpcId() != "" && params.GetVpcId() != routeTable.VpcID {
			continue
		}
		providersRouteTables = append(providersRouteTables, *routeTable)
	}
	return providersRouteTables, nil
}

func (p *providerWithDB) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {
	dbNATGateways, err := p.dbClient.ListNATGateways()
	if err != nil {
		return nil, err
	}
	var providerNATGateways []types.NATGateway
	for _, natGateway := range dbNATGateways {
		if strings.ToLower(natGateway.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != natGateway.AccountID {
			continue
		}
		if params.GetVpcId() != "" && params.GetVpcId() != natGateway.VpcId {
			continue
		}
		providerNATGateways = append(providerNATGateways, *natGateway)
	}
	return providerNATGateways, nil
}

func (p *providerWithDB) ListRouters(ctx context.Context, params *infrapb.ListRoutersRequest) ([]types.Router, error) {
	dbRouters, err := p.dbClient.ListRouters()
	if err != nil {
		return nil, err
	}
	var providerRouters []types.Router
	for _, router := range dbRouters {
		if strings.ToLower(router.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != router.AccountID {
			continue
		}
		providerRouters = append(providerRouters, *router)
	}
	return providerRouters, nil
}

func (p *providerWithDB) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {
	dbIGWs, err := p.dbClient.ListInternetGateways()
	if err != nil {
		return nil, err
	}
	var providerIGWs []types.IGW
	for _, igw := range dbIGWs {
		if strings.ToLower(igw.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != igw.AccountID {
			continue
		}
		providerIGWs = append(providerIGWs, *igw)
	}
	return providerIGWs, nil
}

func (p *providerWithDB) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {
	dbvpce, err := p.dbClient.ListVPCEndpoints()
	if err != nil {
		return nil, err
	}
	var providerVpces []types.VPCEndpoint
	for _, vpce := range dbvpce {
		if strings.ToLower(vpce.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != vpce.AccountID {
			continue
		}
		providerVpces = append(providerVpces, *vpce)
	}
	return providerVpces, nil
}

func (p *providerWithDB) ListPublicIPs(ctx context.Context, params *infrapb.ListPublicIPsRequest) ([]types.PublicIP, error) {
	dbPublicIP, err := p.dbClient.ListPublicIPs()
	if err != nil {
		return nil, err
	}
	var providerPublicIPs []types.PublicIP
	for _, publicIP := range dbPublicIP {
		if strings.ToLower(publicIP.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != publicIP.AccountID {
			continue
		}
		providerPublicIPs = append(providerPublicIPs, *publicIP)
	}
	return providerPublicIPs, nil
}

func (p *providerWithDB) ListLBs(ctx context.Context, params *infrapb.ListLBsRequest) ([]types.LB, error) {
	dbLBs, err := p.dbClient.ListLBs()
	if err != nil {
		return nil, err
	}
	var providersLBs []types.LB
	for _, lb := range dbLBs {
		if !strings.EqualFold(lb.Provider, p.realProvider.GetName()) {
			fmt.Printf("LB provider don't match %s  --- %s \n", p.realProvider.GetName(), lb.Provider)
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != lb.AccountID {
			fmt.Printf("LB account ID don't match %s  --- %s \n", params.GetAccountId(), lb.AccountID)
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != lb.Region {
				fmt.Printf("LB region don't match %s  --- %s \n", params.GetRegion(), lb.Region)
				continue
			}
		}
		if params.GetVpcId() != "" && params.GetVpcId() != lb.VPCID {
			fmt.Printf("LB VPC ID don't match %s  --- %s \n", params.GetVpcId(), lb.VPCID)
			continue
		}
		providersLBs = append(providersLBs, *lb)
	}
	return providersLBs, nil
}

func (p *providerWithDB) ListNetworkInterfaces(ctx context.Context, params *infrapb.ListNetworkInterfacesRequest) ([]types.NetworkInterface, error) {
	dbNetworkInterfaces, err := p.dbClient.ListNetworkInterfaces()
	if err != nil {
		return nil, err
	}
	var providersNetworkInterfaces []types.NetworkInterface
	for _, ni := range dbNetworkInterfaces {
		if strings.ToLower(ni.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != ni.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != ni.Region {
				continue
			}
		}
		if params.GetVpcId() != "" && params.GetVpcId() != ni.VPCID {
			continue
		}

		providersNetworkInterfaces = append(providersNetworkInterfaces, *ni)

	}
	return providersNetworkInterfaces, nil
}

func (p *providerWithDB) ListKeyPairs(ctx context.Context, params *infrapb.ListKeyPairsRequest) ([]types.KeyPair, error) {
	dbKeyPairs, err := p.dbClient.ListKeyPairs()
	if err != nil {
		return nil, err
	}
	var providersKeyPairs []types.KeyPair
	for _, kp := range dbKeyPairs {
		if strings.ToLower(kp.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != kp.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != kp.Region {
				continue
			}
		}

		providersKeyPairs = append(providersKeyPairs, *kp)

	}
	return providersKeyPairs, nil
}

func (p *providerWithDB) GetSubnet(ctx context.Context, params *infrapb.GetSubnetRequest) (types.Subnet, error) {
	dbSubnet, err := p.dbClient.GetSubnet(types.CloudID(p.realProvider.GetName(), params.GetId()))
	if err != nil {
		return types.Subnet{}, err
	}
	if strings.ToLower(dbSubnet.Provider) != strings.ToLower(p.realProvider.GetName()) {
		return types.Subnet{}, fmt.Errorf("subnet with ID %s not found", params.GetId())
	}
	if params.GetAccountId() != "" && params.GetAccountId() != dbSubnet.AccountID {
		return types.Subnet{}, fmt.Errorf("subnet with ID %s does not belong to account %s", params.GetId(), params.GetAccountId())
	}

	if params.GetRegion() != "global" {
		if params.GetRegion() != "" && params.GetRegion() != dbSubnet.Region {
			return types.Subnet{}, fmt.Errorf("subnet with ID %s is not in region %s", params.GetId(), params.GetRegion())
		}
	}

	if params.GetVpcId() != "" && params.GetVpcId() != dbSubnet.VpcId {
		return types.Subnet{}, fmt.Errorf("subnet with ID %s does not belong to VPC with ID %s", params.GetId(), params.GetVpcId())
	}

	return *dbSubnet, nil
}

func (p *providerWithDB) GetVPCIDForCIDR(ctx context.Context, params *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	// TODO use local DB
	return p.realProvider.GetVPCIDForCIDR(ctx, params)
}

func (p *providerWithDB) GetCIDRsForLabels(ctx context.Context, params *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	// TODO use local DB
	return p.realProvider.GetCIDRsForLabels(ctx, params)
}

func (p *providerWithDB) GetIPsForLabels(ctx context.Context, params *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	// TODO use local DB
	return p.realProvider.GetIPsForLabels(ctx, params)
}

func (p *providerWithDB) GetInstancesForLabels(ctx context.Context, params *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	// TODO use local DB
	return p.realProvider.GetInstancesForLabels(ctx, params)
}

func (p *providerWithDB) GetVPCIDWithTag(ctx context.Context, params *infrapb.GetVPCIDWithTagRequest) (string, error) {
	// TODO use local DB
	return p.realProvider.GetVPCIDWithTag(ctx, params)
}

func (p *providerWithDB) ListClusters(ctx context.Context, params *infrapb.ListCloudClustersRequest) ([]types.Cluster, error) {
	dbClusters, err := p.dbClient.ListClusters()
	if err != nil {
		return nil, err
	}
	var clusters []types.Cluster
	for _, cluster := range dbClusters {
		if strings.ToLower(cluster.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != cluster.AccountID {
			continue
		}
		if params.GetRegion() != "global" {
			if params.GetRegion() != "" && params.GetRegion() != cluster.Region {
				continue
			}
		}
		if params.GetVpcId() != "" && params.GetVpcId() != cluster.VpcID {
			continue
		}

		match := true
		for k, v := range params.GetLabels() {
			r, ok := cluster.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}

		if match {
			clusters = append(clusters, *cluster)
		}
	}
	return clusters, nil
}

func (p *providerWithDB) RetrieveClustersData(ctx context.Context) ([]cluster.DiscoveredCluster, error) {
	return p.realProvider.RetrieveClustersData(ctx)
}

func (p *providerWithDB) AddInboundAllowRuleInVPC(ctx context.Context, account, region string, destinationVpcID string, cidrsToAllow []string, ruleName string, tags map[string]string) error {
	return p.realProvider.AddInboundAllowRuleInVPC(ctx, account, region, destinationVpcID, cidrsToAllow, ruleName, tags)
}

func (p *providerWithDB) AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region string, vpcID string, ruleName string, labels map[string]string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	return p.realProvider.AddInboundAllowRuleByLabelsMatch(ctx, account, region, vpcID, ruleName, labels, cidrsToAllow, protocolsAndPorts)
}
func (p *providerWithDB) AddInboundAllowRuleBySubnetMatch(ctx context.Context, account, region string, vpcID string, ruleName string, subnetCidrs []string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, subnets []types.Subnet, err error) {
	return p.realProvider.AddInboundAllowRuleBySubnetMatch(ctx, account, region, vpcID, ruleName, subnetCidrs, cidrsToAllow, protocolsAndPorts)
}

func (p *providerWithDB) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, account, region string, vpcID string, ruleName string, instancesIPs []string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	return p.realProvider.AddInboundAllowRuleByInstanceIPMatch(ctx, account, region, vpcID, ruleName, instancesIPs, cidrsToAllow, protocolsAndPorts)
}

func (p *providerWithDB) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, account, region string, loadBalancerDNS string, vpcID string, ruleName string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (loadBalancerId, ruleId string, err error) {
	return p.realProvider.AddInboundAllowRuleForLoadBalancerByDNS(ctx, account, region, loadBalancerDNS, vpcID, ruleName, cidrsToAllow, protocolsAndPorts)
}

func (p *providerWithDB) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, account, region string, vpcID string, ruleName string) error {
	return p.realProvider.RemoveInboundAllowRuleFromVPCByName(ctx, account, region, vpcID, ruleName)
}

func (p *providerWithDB) RemoveInboundAllowRulesFromVPCById(ctx context.Context, account, region string, vpcID string, instanceIDs []string, loadBalancersIDs []string, ruleId string) error {
	return p.realProvider.RemoveInboundAllowRulesFromVPCById(ctx, account, region, vpcID, instanceIDs, loadBalancersIDs, ruleId)
}

func (p *providerWithDB) RemoveInboundAllowRuleRulesByTags(ctx context.Context, account, region string, vpcID string, ruleName string, tags map[string]string) error {
	return p.realProvider.RemoveInboundAllowRuleRulesByTags(ctx, account, region, vpcID, ruleName, tags)
}

func (p *providerWithDB) RefreshInboundAllowRule(ctx context.Context, account, region string, ruleId string, cidrsToAdd []string, cidrsToRemove []string, destinationLabels map[string]string, destinationPrefixes []string, destinationVPCId string, protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error) {
	return p.realProvider.RefreshInboundAllowRule(ctx, account, region, ruleId, cidrsToAdd, cidrsToRemove, destinationLabels, destinationPrefixes, destinationVPCId, protocolsAndPorts)
}

func (p *providerWithDB) GetSyncTime(id string) (types.SyncTime, error) {
	s, err := p.dbClient.GetSyncTime(id)
	if err != nil {
		return types.SyncTime{}, err
	}
	if s == nil {
		return types.SyncTime{}, fmt.Errorf("nil sync time for id: %s", id)
	}
	return *s, nil
}

func (p *providerWithDB) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	return p.realProvider.ConnectVPCs(ctx, input)
}

func (p *providerWithDB) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {
	return p.realProvider.ConnectVPC(ctx, input)
}

func (p *providerWithDB) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	return p.realProvider.DisconnectVPCs(ctx, input)
}

func (p *providerWithDB) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	return p.realProvider.DisconnectVPC(ctx, input)
}

type KubernetesProviderWithDB struct {
	realProvider provider.Kubernetes
	dbClient     Client
}

func (p *KubernetesProviderWithDB) ListClusters(ctx context.Context) ([]types.Cluster, error) {
	dbClusters, err := p.dbClient.ListClusters()
	if err != nil {
		return nil, err
	}
	var clusters []types.Cluster
	for _, cluster := range dbClusters {
		clusters = append(clusters, *cluster)
	}
	return clusters, nil
}

func (p *KubernetesProviderWithDB) ListNamespaces(ctx context.Context, clusterName string, labels map[string]string) ([]types.Namespace, error) {
	dbNamespaces, err := p.dbClient.ListNamespaces()
	if err != nil {
		return nil, err
	}
	var namespaces []types.Namespace
	for _, namespace := range dbNamespaces {
		if clusterName != "" && namespace.Cluster != clusterName {
			continue
		}
		match := true
		for k, v := range labels {
			r, ok := namespace.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}
		if match {
			namespaces = append(namespaces, *namespace)
		}
	}
	return namespaces, nil
}

func (p *KubernetesProviderWithDB) ListPods(ctx context.Context, clusterName string, labels map[string]string) ([]types.Pod, error) {
	dbPods, err := p.dbClient.ListPods()
	if err != nil {
		return nil, err
	}
	var pods []types.Pod
	for _, pod := range dbPods {
		if clusterName != "" && pod.Cluster != clusterName {
			continue
		}
		match := true
		for k, v := range labels {
			r, ok := pod.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}
		if match {
			pods = append(pods, *pod)
		}
	}
	return pods, nil
}

func (p *KubernetesProviderWithDB) ListServices(ctx context.Context, clusterName string, labels map[string]string) ([]types.K8SService, error) {
	dbServices, err := p.dbClient.ListKubernetesServices()
	if err != nil {
		return nil, err
	}
	var services []types.K8SService
	for _, service := range dbServices {
		if clusterName != "" && service.Cluster != clusterName {
			continue
		}
		match := true
		for k, v := range labels {
			r, ok := service.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}
		if match {
			services = append(services, *service)
		}
	}
	return services, nil
}

func (p *KubernetesProviderWithDB) ListNodes(ctx context.Context, clusterName string, labels map[string]string) ([]types.K8sNode, error) {
	dbNodes, err := p.dbClient.ListKubernetesNodes()
	if err != nil {
		return nil, err
	}
	var nodes []types.K8sNode
	for _, node := range dbNodes {
		if clusterName != "" && node.Cluster != clusterName {
			continue
		}
		match := true
		for k, v := range labels {
			r, ok := node.Labels[k]
			if !ok || r != v {
				match = false
				break
			}
		}
		if match {
			nodes = append(nodes, *node)
		}
	}
	return nodes, nil
}

func (p *KubernetesProviderWithDB) ListPodsCIDRs(ctx context.Context, clusterName string) ([]string, error) {
	// TODO use local DB
	return p.realProvider.ListPodsCIDRs(ctx, clusterName)
}

func (p *KubernetesProviderWithDB) ListServicesCIDRs(ctx context.Context, clusterName string) (string, error) {
	// TODO use local DB
	return p.realProvider.ListServicesCIDRs(ctx, clusterName)
}

func (p *KubernetesProviderWithDB) UpdateServiceSourceRanges(ctx context.Context, clusterName, namespace, name string, cidrsToAdd []string, cidrsToRemove []string) error {
	return p.realProvider.UpdateServiceSourceRanges(ctx, clusterName, namespace, name, cidrsToAdd, cidrsToRemove)
}

func (p *KubernetesProviderWithDB) GetSyncTime(id string) (types.SyncTime, error) {
	s, err := p.dbClient.GetSyncTime(id)
	if err != nil {
		return types.SyncTime{}, err
	}
	if s == nil {
		return types.SyncTime{}, fmt.Errorf("nil sync time for id: %s", id)
	}
	return *s, nil
}

func (p *providerWithDB) ListVPNConcentrators(ctx context.Context, params *infrapb.ListVPNConcentratorsRequest) ([]types.VPNConcentrator, error) {
	dbVPNConcentrators, err := p.dbClient.ListVPNConcentrators()
	if err != nil {
		return nil, err
	}
	var providerVPNConcentrators []types.VPNConcentrator
	for _, vpnc := range dbVPNConcentrators {
		if strings.ToLower(vpnc.Provider) != strings.ToLower(p.realProvider.GetName()) {
			continue
		}
		if params.GetAccountId() != "" && params.GetAccountId() != vpnc.AccountID {
			continue
		}
		if params.GetRegion() != "" && params.GetRegion() != vpnc.Region {
			continue
		}

		providerVPNConcentrators = append(providerVPNConcentrators, *vpnc)
	}
	return providerVPNConcentrators, nil
}

func (p *providerWithDB) ListVpcGraphNodes(ctx context.Context, params *infrapb.ListVpcGraphNodesRequest) ([]types.VpcGraphNode, error) {
	vpcIndexKey := types.CloudID(p.realProvider.GetName(), params.GetVpcId())
	vpcIndex, err := p.dbClient.GetVPCIndex(vpcIndexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC index for %s: %w", vpcIndexKey, err)
	}
	if vpcIndex == nil {
		return nil, fmt.Errorf("VPC index not found for %s", vpcIndexKey)
	}

	var nodes []types.VpcGraphNode
	providerName := p.realProvider.GetName()
	logger := p.logger

	createNode := func(id, resourceType, name string, properties map[string]string) types.VpcGraphNode {
		if properties == nil {
			properties = make(map[string]string)
		}
		return types.VpcGraphNode{
			ID:           id,
			ResourceType: resourceType,
			Name:         name,
			Properties:   properties,
			Provider:     providerName,
			AccountId:    vpcIndex.AccountId,
			Region:       vpcIndex.Region,
		}
	}

	for _, id := range vpcIndex.InstanceIds {
		res, err := p.dbClient.GetInstance(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"privateIP": res.PrivateIP, "publicIP": res.PublicIP, "state": res.State, "type": res.Type}
			nodes = append(nodes, createNode(res.ID, types.InstanceType, res.Name, props))
		} else {
			logger.Warnf("Failed to get instance %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.SubnetIds {
		res, err := p.dbClient.GetSubnet(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"cidrBlock": res.CidrBlock, "zone": res.Zone}
			nodes = append(nodes, createNode(res.SubnetId, types.SubnetType, res.Name, props))
		} else {
			logger.Warnf("Failed to get subnet %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.RouteTableIds {
		res, err := p.dbClient.GetRouteTable(types.CloudID(providerName, id))
		if err == nil && res != nil {
			nodes = append(nodes, createNode(res.ID, types.RouteTableType, res.Name, nil))
		} else {
			logger.Warnf("Failed to get route table %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.NatGatewayIds {
		res, err := p.dbClient.GetNATGateway(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"publicIP": res.PublicIp, "privateIP": res.PrivateIp, "state": res.State}
			nodes = append(nodes, createNode(res.ID, types.NATGatewayType, res.Name, props))
		} else {
			logger.Warnf("Failed to get NAT gateway %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.IgwIds {
		res, err := p.dbClient.GetIGW(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"state": res.State}
			nodes = append(nodes, createNode(res.ID, types.IGWType, res.Name, props))
		} else {
			logger.Warnf("Failed to get internet gateway %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.SecurityGroupIds {
		res, err := p.dbClient.GetSecurityGroup(types.CloudID(providerName, id))
		if err == nil && res != nil {
			nodes = append(nodes, createNode(res.ID, types.SecurityGroupType, res.Name, nil))
		} else {
			logger.Warnf("Failed to get security group %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.AclIds {
		res, err := p.dbClient.GetACL(types.CloudID(providerName, id))
		if err == nil && res != nil {
			nodes = append(nodes, createNode(res.ID, types.ACLType, res.Name, nil))
		} else {
			logger.Warnf("Failed to get ACL %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.LbIds {
		res, err := p.dbClient.GetLB(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"dnsName": res.DNSName, "type": res.Type, "scheme": res.Scheme}
			nodes = append(nodes, createNode(res.ID, types.LBType, res.Name, props))
		} else {
			logger.Warnf("Failed to get LB %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.VpcEndpointIds {
		res, err := p.dbClient.GetVPCEndpoint(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"serviceName": res.ServiceName, "type": res.Type, "state": res.State}
			nodes = append(nodes, createNode(res.ID, types.VPCEndpointType, res.Name, props))
		} else {
			logger.Warnf("Failed to get VPC Endpoint %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.RouterIds {
		res, err := p.dbClient.GetRouter(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"asn": fmt.Sprintf("%d", res.ASN), "state": res.State}
			nodes = append(nodes, createNode(res.ID, types.RouterType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Router %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.NetworkInterfaceIds {
		res, err := p.dbClient.GetNetworkInterface(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"macAddress": res.MacAddress, "status": res.Status, "publicIP": res.PublicIP}
			nodes = append(nodes, createNode(res.ID, types.NetworkInterfaceType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Network Interface %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.VpnConcentratorIds {
		res, err := p.dbClient.GetVPNConcentrator(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"asn": fmt.Sprintf("%d", res.ASN), "state": res.State, "type": res.Type}
			nodes = append(nodes, createNode(res.ID, types.VPNConcentratorType, res.Name, props))
		} else {
			logger.Warnf("Failed to get VPN Concentrator %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.PublicIpIds {
		res, err := p.dbClient.GetPublicIP(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{"publicIP": res.PublicIP, "instanceId": res.InstanceId, "type": res.Type}
			nodes = append(nodes, createNode(res.ID, types.PublicIPType, res.ID, props))
		} else {
			logger.Warnf("Failed to get Public IP %s for graph node: %v", id, err)
		}
	}
	for _, id := range vpcIndex.ClusterIds {
		res, err := p.dbClient.GetCluster(types.CloudID(providerName, id))
		if err == nil && res != nil {
			props := map[string]string{} // Add properties like ARN if available in res
			nodes = append(nodes, createNode(res.Name, types.ClusterType, res.Name, props))
		} else {
			logger.Warnf("Failed to get Cluster %s for graph node: %v", id, err)
		}
	}

	return nodes, nil
}

func (p *providerWithDB) ListVpcGraphEdges(ctx context.Context, params *infrapb.ListVpcGraphEdgesRequest) ([]types.VpcGraphEdge, error) {
	vpcIndexKey := types.CloudID(p.realProvider.GetName(), params.GetVpcId())
	vpcIndex, err := p.dbClient.GetVPCIndex(vpcIndexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC index for %s: %w", vpcIndexKey, err)
	}
	if vpcIndex == nil {
		return nil, fmt.Errorf("VPC index not found for %s", vpcIndexKey)
	}

	var edges []types.VpcGraphEdge
	providerName := p.realProvider.GetName()
	logger := p.logger

	createEdge := func(sourceID, targetID, relationshipType string) types.VpcGraphEdge {
		return types.VpcGraphEdge{
			SourceNodeID:     sourceID,
			TargetNodeID:     targetID,
			RelationshipType: relationshipType,
			Provider:         providerName,
			AccountId:        vpcIndex.AccountId,
			Region:           vpcIndex.Region,
		}
	}

	for _, subnetId := range vpcIndex.SubnetIds {
		edges = append(edges, createEdge(vpcIndex.VpcId, subnetId, "CONTAINS"))
	}

	for _, subnetId := range vpcIndex.SubnetIds {
		subnet, err := p.dbClient.GetSubnet(types.CloudID(providerName, subnetId))
		if err != nil || subnet == nil {
			logger.Warnf("Failed to get subnet %s for edge generation: %v", subnetId, err)
			continue
		}
	}

	for _, rtId := range vpcIndex.RouteTableIds {
		rt, err := p.dbClient.GetRouteTable(types.CloudID(providerName, rtId))
		if err != nil || rt == nil {
			logger.Warnf("Failed to get route table %s for edge generation: %v", rtId, err)
			continue
		}
		for _, subnetId := range rt.SubnetIds {
			if contains(vpcIndex.SubnetIds, subnetId) {
				edges = append(edges, createEdge(rtId, subnetId, "ASSOCIATED_WITH"))
			}
		}
		for _, route := range rt.Routes {
			targetId := route.Target
			if contains(vpcIndex.IgwIds, targetId) ||
				contains(vpcIndex.NatGatewayIds, targetId) ||
				contains(vpcIndex.InstanceIds, targetId) ||
				contains(vpcIndex.VpcEndpointIds, targetId) ||
				contains(vpcIndex.NetworkInterfaceIds, targetId) {
				edges = append(edges, createEdge(rtId, targetId, "ROUTES_TO"))
			}
		}
	}

	for _, instanceId := range vpcIndex.InstanceIds {
		instance, err := p.dbClient.GetInstance(types.CloudID(providerName, instanceId))
		if err != nil || instance == nil {
			logger.Warnf("Failed to get instance %s for edge generation: %v", instanceId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, instance.SubnetID) {
			edges = append(edges, createEdge(instanceId, instance.SubnetID, "LOCATED_IN"))
		}
		for _, sgId := range instance.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edges = append(edges, createEdge(instanceId, sgId, "USES_SECURITY_GROUP"))
			}
		}
		for _, niId := range instance.InterfaceIDs {
			if contains(vpcIndex.NetworkInterfaceIds, niId) {
				edges = append(edges, createEdge(instanceId, niId, "HAS_INTERFACE"))
			}
		}
	}

	for _, niId := range vpcIndex.NetworkInterfaceIds {
		ni, err := p.dbClient.GetNetworkInterface(types.CloudID(providerName, niId))
		if err != nil || ni == nil {
			logger.Warnf("Failed to get network interface %s for edge generation: %v", niId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, ni.SubnetID) {
			edges = append(edges, createEdge(niId, ni.SubnetID, "LOCATED_IN"))
		}
		for _, sgId := range ni.SecurityGroupIDs {
			if contains(vpcIndex.SecurityGroupIds, sgId) {
				edges = append(edges, createEdge(niId, sgId, "USES_SECURITY_GROUP"))
			}
		}
	}

	for _, natId := range vpcIndex.NatGatewayIds {
		nat, err := p.dbClient.GetNATGateway(types.CloudID(providerName, natId))
		if err != nil || nat == nil {
			logger.Warnf("Failed to get NAT gateway %s for edge generation: %v", natId, err)
			continue
		}
		if contains(vpcIndex.SubnetIds, nat.SubnetId) {
			edges = append(edges, createEdge(natId, nat.SubnetId, "LOCATED_IN"))
		}
	}

	for _, igwId := range vpcIndex.IgwIds {
		igw, err := p.dbClient.GetIGW(types.CloudID(providerName, igwId))
		if err != nil || igw == nil {
			logger.Warnf("Failed to get IGW %s for edge generation: %v", igwId, err)
			continue
		}
		if igw.AttachedVpcId == vpcIndex.VpcId {
			edges = append(edges, createEdge(igwId, vpcIndex.VpcId, "ATTACHED_TO"))
		}
	}

	for _, vpceId := range vpcIndex.VpcEndpointIds {
		vpce, err := p.dbClient.GetVPCEndpoint(types.CloudID(providerName, vpceId))
		if err != nil || vpce == nil {
			logger.Warnf("Failed to get VPC Endpoint %s for edge generation: %v", vpceId, err)
			continue
		}
		if vpce.VPCId == vpcIndex.VpcId {
			edges = append(edges, createEdge(vpceId, vpcIndex.VpcId, "LOCATED_IN"))
		}
	}

	for _, lbId := range vpcIndex.LbIds {
		lb, err := p.dbClient.GetLB(types.CloudID(providerName, lbId))
		if err != nil || lb == nil {
			logger.Warnf("Failed to get LB %s for edge generation: %v", lbId, err)
			continue
		}
		for _, instanceId := range lb.InstanceIDs {
			if contains(vpcIndex.InstanceIds, instanceId) {
				edges = append(edges, createEdge(lbId, instanceId, "LOAD_BALANCES_TO"))
			}
		}
	}

	for _, aclId := range vpcIndex.AclIds {
		acl, err := p.dbClient.GetACL(types.CloudID(providerName, aclId))
		if err != nil || acl == nil {
			logger.Warnf("Failed to get ACL %s for edge generation: %v", aclId, err)
			continue
		}
		for _, subnetId := range acl.Subnets {
			if contains(vpcIndex.SubnetIds, subnetId) {
				edges = append(edges, createEdge(aclId, subnetId, "ASSOCIATED_WITH"))
			}
		}
	}

	return edges, nil
}

func (p *providerWithDB) GetVpcConnectivityGraph(ctx context.Context, params *infrapb.GetVpcConnectivityGraphRequest) ([]types.VpcGraphNode, []types.VpcGraphEdge, error) {
	nodesReq := &infrapb.ListVpcGraphNodesRequest{
		Provider:  params.Provider,
		AccountId: params.AccountId,
		Region:    params.Region,
		VpcId:     params.VpcId,
		Creds:     params.Creds,
	}
	nodes, err := p.ListVpcGraphNodes(ctx, nodesReq)
	if err != nil {
		p.logger.Errorf("Failed to list VPC graph nodes for graph request (VPC %s): %v", params.VpcId, err)
		return nil, nil, fmt.Errorf("failed to get nodes: %w", err)
	}

	edgesReq := &infrapb.ListVpcGraphEdgesRequest{
		Provider:  params.Provider,
		AccountId: params.AccountId,
		Region:    params.Region,
		VpcId:     params.VpcId,
		Creds:     params.Creds,
	}
	edges, err := p.ListVpcGraphEdges(ctx, edgesReq)
	if err != nil {
		p.logger.Errorf("Failed to list VPC graph edges for graph request (VPC %s): %v", params.VpcId, err)
		return nil, nil, fmt.Errorf("failed to get edges: %w", err)
	}

	return nodes, edges, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
