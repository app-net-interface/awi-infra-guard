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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/app-net-interface/awi-infra-guard/azure"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/app-net-interface/kubernetes-discovery/cluster"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	"github.com/app-net-interface/awi-infra-guard/aws"
	"github.com/app-net-interface/awi-infra-guard/gcp"
	infra_kubernetes "github.com/app-net-interface/awi-infra-guard/kubernetes"
	"github.com/app-net-interface/awi-infra-guard/types"
)

type Strategy interface {
	GetProvider(ctx context.Context, cloud string) (CloudProvider, error)
	GetAllProviders() []CloudProvider
	GetKubernetesProvider() (Kubernetes, error)
	RefreshState(ctx context.Context) error
}

type CloudProvider interface {
	GetName() string
	ListAccounts() []types.Account
	// ListVPC returns cloud instances based on provided filters, empty filter means no filtering by this parameter.
	ListVPC(ctx context.Context, input *infrapb.ListVPCRequest) ([]types.VPC, error)
	// ListInstances returns cloud instances based on provided filters, empty filter means no filtering by this parameter.
	ListInstances(ctx context.Context, input *infrapb.ListInstancesRequest) ([]types.Instance, error)
	// ListSubnets returns cloud instances based on provided filters, empty filter means no filtering by this parameter.
	// Scope of subnet is regional in some clouds (e.g. GCP, Azure) and zonal in others (e.g. AWS), filtering is done by
	// this scope.
	ListSubnets(ctx context.Context, input *infrapb.ListSubnetsRequest) ([]types.Subnet, error)
	ListACLs(ctx context.Context, input *infrapb.ListACLsRequest) ([]types.ACL, error)
	ListSecurityGroups(ctx context.Context, input *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error)
	ListRouteTables(ctx context.Context, input *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error)
	ListNATGateways(ctx context.Context, input *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error)
	ListRouters(ctx context.Context, input *infrapb.ListRoutersRequest) ([]types.Router, error)


	// GetSubnet returns single subnet based on it's ID
	GetSubnet(ctx context.Context, input *infrapb.GetSubnetRequest) (types.Subnet, error)
	// GetVPCIDForCIDR returns ID of VPC which have subnet with given CIDR.
	GetVPCIDForCIDR(ctx context.Context, input *infrapb.GetVPCIDForCIDRRequest) (string, error)
	// GetCIDRsForLabels returns CIDRs of subnets with given labels.
	GetCIDRsForLabels(ctx context.Context, input *infrapb.GetCIDRsForLabelsRequest) ([]string, error)
	// GetIPsForLabels returns IPs of instances with given labels.
	GetIPsForLabels(ctx context.Context, input *infrapb.GetIPsForLabelsRequest) ([]string, error)
	// GetInstancesForLabels returns instances with given labels.
	GetInstancesForLabels(ctx context.Context, input *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error)
	GetVPCIDWithTag(ctx context.Context, input *infrapb.GetVPCIDWithTagRequest) (string, error)
	ListClusters(ctx context.Context, input *infrapb.ListCloudClustersRequest) ([]types.Cluster, error)
	RetrieveClustersData(ctx context.Context) ([]cluster.DiscoveredCluster, error)
	AccessControl
	DBMetadata
	VPCConnector
}

type AccessControl interface {
	// AddInboundAllowRuleInVPC allows given cidrs in all VPC instances. Security rules are created with name ruleName
	// and tags if they are allowed in given cloud.
	AddInboundAllowRuleInVPC(ctx context.Context, account, region string, destinationVpcID string, cidrsToAllow []string, ruleName string,
		tags map[string]string) error
	// AddInboundAllowRuleByLabelsMatch allows cidrsToAllow with protocolsAndPorts to all instances which match to labels
	AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region string,
		vpcID string, ruleName string, labels map[string]string, cidrsToAllow []string,
		protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error)
	// AddInboundAllowRuleBySubnetMatch allows cidrsToAllow with protocolsAndPorts to all instances which are within provided cloud subnets
	AddInboundAllowRuleBySubnetMatch(ctx context.Context, account, region string,
		vpcID string, ruleName string, subnetCidrs []string, cidrsToAllow []string,
		protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, subnets []types.Subnet, err error)
	// AddInboundAllowRuleByInstanceIPMatch allows cidrsToAllow with protocolsAndPorts to all instances which have provided instancesIPs
	AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, account, region string,
		vpcID string, ruleName string, instancesIPs []string, cidrsToAllow []string,
		protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error)
	// AddInboundAllowRuleForLoadBalancerByDNS allows cidrsToAllow with protocolsAndPorts to load balancer with given DNS
	AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, account, region string, loadBalancerDNS string, vpcID string,
		ruleName string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts,
	) (loadBalancerId, ruleId string, err error)
	RemoveInboundAllowRuleFromVPCByName(ctx context.Context, account, region string, vpcID string, ruleName string) error
	RemoveInboundAllowRulesFromVPCById(ctx context.Context, account, region string, vpcID string, instanceIDs []string,
		loadBalancersIDs []string, ruleId string) error
	RemoveInboundAllowRuleRulesByTags(ctx context.Context, account, region string, vpcID string, ruleName string, tags map[string]string) error
	// RefreshInboundAllowRule adds and removes CIDRs in rule rules and applies rule in instances matching
	// to destinationLabels or destinationPrefixes
	RefreshInboundAllowRule(ctx context.Context, account, region string, ruleId string, cidrsToAdd []string, cidrsToRemove []string,
		destinationLabels map[string]string, destinationPrefixes []string, destinationVPCId string,
		protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error)
}

type Kubernetes interface {
	ListClusters(ctx context.Context) (clusters []types.Cluster, err error)
	ListNamespaces(ctx context.Context, clusterName string, labels map[string]string) (namespaces []types.Namespace, err error)
	ListPods(ctx context.Context, clusterName string, labels map[string]string) (pods []types.Pod, err error)
	ListServices(ctx context.Context, clusterName string, labels map[string]string) (services []types.K8SService, err error)
	ListNodes(ctx context.Context, clusterName string, labels map[string]string) (nodes []types.K8sNode, err error)
	ListPodsCIDRs(ctx context.Context, clusterName string) ([]string, error)
	ListServicesCIDRs(ctx context.Context, clusterName string) (string, error)
	UpdateServiceSourceRanges(ctx context.Context, clusterName, namespace, name string, cidrsToAdd []string, cidrsToRemove []string) error
	DBMetadata
}

type DBMetadata interface {
	GetSyncTime(id string) (types.SyncTime, error)
}

type VPCConnector interface {
	ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error)
	ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error)
	DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error)
	DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error)
}

type RealProviderStrategy struct {
	awsClient   *aws.Client
	gcpClient   *gcp.Client
	azureClient *azure.Client
	k8sClient   *infra_kubernetes.KubernetesClient
	logger      *logrus.Logger
}

func (s *RealProviderStrategy) GetProvider(ctx context.Context, cloud string) (CloudProvider, error) {
	s.logger.Infof("Using %s as an infra provider", cloud)
	switch strings.ToLower(cloud) {
	case "aws":
		if s.awsClient == nil {
			return nil, fmt.Errorf("AWS client is not initialized")
		}
		return s.awsClient, nil
	case "gcp":
		if s.gcpClient == nil {
			return nil, fmt.Errorf("GCP client is not initizalized")
		}
		return s.gcpClient, nil
	case "azure":
		if s.azureClient == nil {
			return nil, fmt.Errorf("Azure client is not initizalized")
		}
		return s.azureClient, nil
	}
	return nil, fmt.Errorf("unsupported provider")
}

func (s *RealProviderStrategy) GetAllProviders() []CloudProvider {
	var providers []CloudProvider
	if s.gcpClient != nil {
		providers = append(providers, s.gcpClient)
	}
	if s.awsClient != nil {
		providers = append(providers, s.awsClient)
	}
	if s.azureClient != nil {
		providers = append(providers, s.azureClient)
	}
	return providers
}

func (s *RealProviderStrategy) GetKubernetesProvider() (Kubernetes, error) {
	if s.k8sClient == nil {
		return nil, fmt.Errorf("kubernetes client is not initizalized")
	}
	return s.k8sClient, nil
}

func (s *RealProviderStrategy) RefreshState(ctx context.Context) error {
	s.logger.Debugf("Refreshing clusters state...")
	s.RetrieveClusters(ctx)
	return nil
}

func NewRealProviderStrategy(ctx context.Context, logger *logrus.Logger, kubeConfigFileName string) *RealProviderStrategy {
	s := &RealProviderStrategy{
		logger: logger,
	}
	var err error
	s.awsClient, err = aws.NewClient(ctx, s.logger)
	if err != nil {
		logger.Warnf("Failed to init AWS client: %v", err)
	}
	s.gcpClient, err = gcp.NewClient(ctx, s.logger)
	if err != nil {
		logger.Warnf("Failed to init GCP client: %v", err)
	}
	s.azureClient, err = azure.NewClient(ctx, s.logger)
	if err != nil {
		logger.Warnf("Failed to init Azure client: %v", err)
	}

	s.k8sClient, err = infra_kubernetes.NewKubernetesClient(logger, kubeConfigFileName)
	if err != nil {
		logger.Warnf("Failed to init kubernetes clients: %v", err)
	}
	s.RetrieveClusters(ctx)
	return s
}

func (s *RealProviderStrategy) RetrieveClusters(ctx context.Context) {
	retrievedClients := make(map[string]*kubernetes.Clientset)
	for _, cloudProvider := range s.GetAllProviders() {
		if cloudProvider == nil {
			continue
		}
		clusters, err := cloudProvider.RetrieveClustersData(ctx)
		if err != nil {
			s.logger.Warnf("Failed to retrive clusters data from provider %s", cloudProvider.GetName())
			continue
		}
		for _, cl := range clusters {
			if cl == nil {
				continue
			}
			data, err := cl.GetData()
			if err != nil {
				s.logger.Warnf("Failed to get data from cluster in provider %s", cloudProvider.GetName())
				continue
			}
			if data.Name == "" {
				s.logger.Warnf("Empty name in cluster in provider %s", cloudProvider.GetName())
				continue
			}
			clientset, err := cl.GetClientset(ctx)
			if err != nil {
				s.logger.Warnf("Failed to get clientset from cluster %s in provider %s", data.Name, cloudProvider.GetName())
				continue
			}
			retrievedClients[data.Name] = clientset
		}
	}
	// TODO remove clusters which don't exist anymore
	s.k8sClient.AddClients(retrievedClients)
}
