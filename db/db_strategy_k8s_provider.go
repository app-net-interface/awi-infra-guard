// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/types"
)

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
