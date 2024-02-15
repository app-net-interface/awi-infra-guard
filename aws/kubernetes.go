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

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/app-net-interface/kubernetes-discovery/cluster"
	eksdiscovery "github.com/app-net-interface/kubernetes-discovery/cluster/eks"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	awstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
)

func (c *Client) ListClusters(ctx context.Context, params *infrapb.ListCloudClustersRequest) ([]types.Cluster, error) {
	if params.GetRegion() == "" || params.GetRegion() == "all" {
		return c.listClustersInAllRegions(ctx, params.GetAccountId(), params.GetVpcId(), params.GetLabels())
	}

	return c.listClustersInSingleRegion(ctx, params.GetAccountId(), params.GetRegion(), params.GetVpcId(), params.GetLabels())
}

func (c *Client) listClustersInAllRegions(ctx context.Context, account, vpcID string, labels map[string]string) ([]types.Cluster, error) {
	var (
		wg            sync.WaitGroup
		allClusters   []types.Cluster
		resultChannel = make(chan []types.Cluster)
		errorChannel  = make(chan error)
	)

	regionalClients, err := c.getAllClientsForProfile(account)
	if err != nil {
		return nil, err
	}

	for regionName, awsRegionClient := range regionalClients {
		wg.Add(1)
		go func(regionName string, awsRegionClient awsClient) {
			defer wg.Done()
			clustersOut, err := awsRegionClient.eksClient.ListClusters(ctx, &eks.ListClustersInput{})
			if err != nil {
				errorChannel <- fmt.Errorf("could not get AWS clusters: %v", err)
				return
			}
			clusters, err := filterAndConvertClusters(ctx, awsRegionClient.eksClient, c.defaultAccountID, c.defaultRegion, account, regionName, clustersOut.Clusters, vpcID, labels)
			if err != nil {
				errorChannel <- fmt.Errorf("could not get AWS cluster: %v", err)
				return
			}
			resultChannel <- clusters
		}(regionName, awsRegionClient)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
		close(errorChannel)
	}()

	for subnets := range resultChannel {
		allClusters = append(allClusters, subnets...)
	}

	if err := <-errorChannel; err != nil {
		return nil, err
	}

	return allClusters, nil
}

func (c *Client) listClustersInSingleRegion(ctx context.Context, account, region, vpcID string, labels map[string]string) ([]types.Cluster, error) {
	eksClient, err := c.getEKSClient(ctx, account, region)
	if err != nil {
		return nil, err
	}

	clustersOut, err := eksClient.ListClusters(ctx, &eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}

	clusters, err := filterAndConvertClusters(ctx, eksClient, c.defaultAccountID, c.defaultRegion, account, region, clustersOut.Clusters, vpcID, labels)
	return clusters, nil
}

func filterAndConvertClusters(ctx context.Context, client *eks.Client, defaultAccountID, defaultRegion, account, region string, awsClusters []string, vpcID string, labels map[string]string) ([]types.Cluster, error) {
	clusters := make([]types.Cluster, 0, len(awsClusters))

clusterLoop:
	for _, awsCluster := range awsClusters {
		clusterOut, err := client.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: aws.String(awsCluster)})
		if err != nil {
			return nil, err
		}
		cluster := awsClusterToTypesCluster(defaultAccountID, defaultRegion, account, region, clusterOut.Cluster)
		if vpcID != "" && cluster.VpcID != vpcID {
			continue clusterLoop
		}
		for k, v := range labels {
			if cluster.Labels[k] != v {
				continue clusterLoop
			}
		}

		clusters = append(clusters, cluster)
	}

	return clusters, nil
}

func awsClusterToTypesCluster(defaultAccountID, defaultRegion, account, region string, in *awstypes.Cluster) types.Cluster {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccountID
	}
	return types.Cluster{
		Name:      convertString(in.Name),
		VpcID:     convertString(in.ResourcesVpcConfig.VpcId),
		Arn:       convertString(in.Arn),
		Labels:    in.Tags,
		Region:    region,
		FullName:  convertString(in.Name) + "." + region + ".eksctl.io",
		AccountID: account,
		Provider:  providerName,
		Id:        convertString(in.Name),
	}
}

func (c *Client) RetrieveClustersData(ctx context.Context) ([]cluster.DiscoveredCluster, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get AWS config: %v", err)
	}
	eksDiscovery := eksdiscovery.NewClustersRetriever(cfg)
	clusters, err := eksDiscovery.Retrieve(ctx, cluster.WithRegions("-"))
	if err != nil {
		return nil, err
	}
	return clusters.DiscoveredClusters, nil
}
