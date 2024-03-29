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

package azure

import (
	"context"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/app-net-interface/kubernetes-discovery/cluster"
)

func (c *Client) ListNamespaces(ctx context.Context, clusterName string, labels map[string]string) (namespaces []types.Namespace, err error) {
	// TBD
	return nil, nil
}

func (c *Client) ListPods(ctx context.Context, clusterName string, labels map[string]string) (pods []types.Pod, err error) {
	// TBD
	return nil, nil
}

func (c *Client) ListServices(ctx context.Context, clusterName string, labels map[string]string) (services []types.K8SService, err error) {
	// TBD
	return nil, nil
}

func (c *Client) ListNodes(ctx context.Context, clusterName string, labels map[string]string) (nodes []types.K8sNode, err error) {
	// TBD
	return nil, nil
}

func (c *Client) ListPodsCIDRs(ctx context.Context, clusterName string) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) ListServicesCIDRs(ctx context.Context, clusterName string) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) UpdateServiceSourceRanges(ctx context.Context, clusterName, namespace, name string, cidrsToAdd []string, cidrsToRemove []string) error {
	// TBD
	return nil
}

func (c *Client) ListClusters(ctx context.Context, input *infrapb.ListCloudClustersRequest) ([]types.Cluster, error) {
	// TBD
	return nil, nil
}

func (c *Client) RetrieveClustersData(ctx context.Context) ([]cluster.DiscoveredCluster, error) {
	// TBD
	return nil, nil
}
