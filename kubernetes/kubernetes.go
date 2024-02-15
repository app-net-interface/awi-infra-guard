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

package kubernetes

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/app-net-interface/awi-infra-guard/types"
)

type KubernetesClient struct {
	mtx     sync.Mutex
	logger  *logrus.Logger
	clients map[string]*kubernetes.Clientset
}

func (k *KubernetesClient) getClient(cluster string) (*kubernetes.Clientset, error) {
	if cluster == "" {
		return nil, fmt.Errorf("cluster name not provided")
	}
	k.mtx.Lock()
	defer k.mtx.Unlock()
	client, ok := k.clients[cluster]
	if !ok {
		return nil, fmt.Errorf("cluster not found: %s", cluster)
	}
	return client, nil
}

func NewKubernetesClient(logger *logrus.Logger, kubeConfigFileName string) (*KubernetesClient, error) {
	clients, err := parseK8sConfig(logger, kubeConfigFileName)
	if err != nil {
		logger.Errorf("Failed to parse kube config: %v", err)
	}

	return &KubernetesClient{
		logger:  logger,
		clients: clients,
	}, nil
}

func parseK8sConfig(logger *logrus.Logger, kubeConfigFileName string) (map[string]*kubernetes.Clientset, error) {
	if kubeConfigFileName == "" {
		ok := false
		home := homedir.HomeDir()
		if home != "" {
			kubeConfigFileName = filepath.Join(home, ".kube", "config")
			_, err := os.Stat(kubeConfigFileName)
			if err == nil {
				ok = true
			}
		}
		if !ok {
			logger.Warnf("kube config file not provided in kubeConfigFileName parameter and not present" +
				" in HOME/.kube/config, kuberentes clusters won't be watched")
			return nil, nil
		}
	}

	kubeconfigBytes, err := os.ReadFile(kubeConfigFileName)
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.Load(kubeconfigBytes)
	if err != nil {
		return nil, err
	}

	clientMap := make(map[string]*kubernetes.Clientset)
	for ctxName, ctx := range config.Contexts {
		cfg, err := buildConfigFromFlags(ctxName, kubeConfigFileName)
		if err != nil {
			return nil, err
		}
		k8sClient, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			return nil, err
		}
		clientMap[ctx.Cluster] = k8sClient
	}

	return clientMap, nil
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func (k *KubernetesClient) AddClients(clients map[string]*kubernetes.Clientset) {
	k.mtx.Lock()
	defer k.mtx.Unlock()
	if len(clients) == 0 {
		return
	}
	if k.clients == nil {
		k.clients = clients
		return
	}
	for name, cl := range clients {
		k.clients[name] = cl
	}
	return
}

func (k *KubernetesClient) ListClusters(ctx context.Context) ([]types.Cluster, error) {
	clusters := make([]types.Cluster, 0, len(k.clients))
	k.mtx.Lock()
	defer k.mtx.Unlock()
	for cluster := range k.clients {
		clusters = append(clusters, types.Cluster{Name: cluster})
	}
	return clusters, nil
}

func (k *KubernetesClient) ListPods(ctx context.Context, clusterName string, labels map[string]string) (pods []types.Pod, err error) {
	k8sClient, err := k.getClient(clusterName)
	if err != nil {
		return nil, err
	}
	var labelSelector string
	for k, v := range labels {
		labelSelector += fmt.Sprintf("%s=%s,", k, v)
	}
	labelSelector = strings.TrimSuffix(labelSelector, ",")
	k8sPodsList, err := k8sClient.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	return k8sPodsToTypes(clusterName, k8sPodsList.Items), nil
}

func (k *KubernetesClient) ListServices(ctx context.Context, clusterName string, labels map[string]string) ([]types.K8SService, error) {
	var k8sClient *kubernetes.Clientset
	var err error
	if clusterName != "" {
		k8sClient, err = k.getClient(clusterName)
		if err != nil {
			return nil, err
		}
	}

	var labelSelector string
	for k, v := range labels {
		labelSelector += fmt.Sprintf("%s=%s,", k, v)
	}
	labelSelector = strings.TrimSuffix(labelSelector, ",")
	var services []types.K8SService
	if clusterName != "" {
		k8sServicesList, err := k8sClient.CoreV1().Services("").List(ctx, metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			return nil, err
		}
		services = k8sServicesToTypes(clusterName, k8sServicesList.Items)
	} else {
		k.mtx.Lock()
		defer k.mtx.Unlock()
		for cluster, client := range k.clients {
			k8sServicesList, err := client.CoreV1().Services("").List(ctx, metav1.ListOptions{
				LabelSelector: labelSelector,
			})
			if err != nil {
				return nil, err
			}
			services = append(services, k8sServicesToTypes(cluster, k8sServicesList.Items)...)
		}
	}

	return services, nil
}

func (k *KubernetesClient) ListNodes(ctx context.Context, clusterName string, labels map[string]string) ([]types.K8sNode, error) {
	var k8sClient *kubernetes.Clientset
	var err error
	if clusterName != "" {
		k8sClient, err = k.getClient(clusterName)
		if err != nil {
			return nil, err
		}
	}

	var labelSelector string
	for k, v := range labels {
		labelSelector += fmt.Sprintf("%s=%s,", k, v)
	}
	labelSelector = strings.TrimSuffix(labelSelector, ",")
	var nodes []types.K8sNode
	if clusterName != "" {
		k8sNodesList, err := k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			return nil, err
		}
		nodes = k8sNodesToTypes(clusterName, k8sNodesList.Items)
	} else {
		k.mtx.Lock()
		defer k.mtx.Unlock()
		for cluster, client := range k.clients {
			k8sNodesList, err := client.CoreV1().Nodes().List(ctx, metav1.ListOptions{
				LabelSelector: labelSelector,
			})
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, k8sNodesToTypes(cluster, k8sNodesList.Items)...)
		}
	}

	return nodes, nil
}

func (k *KubernetesClient) ListNamespaces(ctx context.Context, clusterName string, labels map[string]string) ([]types.Namespace, error) {
	var k8sClient *kubernetes.Clientset
	var err error
	if clusterName != "" {
		k8sClient, err = k.getClient(clusterName)
		if err != nil {
			return nil, err
		}
	}

	var labelSelector string
	for k, v := range labels {
		labelSelector += fmt.Sprintf("%s=%s,", k, v)
	}
	labelSelector = strings.TrimSuffix(labelSelector, ",")
	var namespaces []types.Namespace
	if clusterName != "" {
		k8sNamespaceList, err := k8sClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			return nil, err
		}
		namespaces = k8sNamespacesToTypes(clusterName, k8sNamespaceList.Items)
	} else {
		k.mtx.Lock()
		defer k.mtx.Unlock()
		for cluster, client := range k.clients {
			k8sNamespaceList, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{
				LabelSelector: labelSelector,
			})
			if err != nil {
				return nil, err
			}
			namespaces = append(namespaces, k8sNamespacesToTypes(cluster, k8sNamespaceList.Items)...)
		}
	}

	return namespaces, nil
}

func (k *KubernetesClient) UpdateServiceSourceRanges(ctx context.Context, clusterName, namespace, name string, cidrsToAdd []string, cidrsToRemove []string) error {
	k8sClient, err := k.getClient(clusterName)
	if err != nil {
		return err
	}
	svc, err := k8sClient.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	newSourceRanges := make([]string, 0)
	if len(cidrsToRemove) == 0 {
		newSourceRanges = svc.Spec.LoadBalancerSourceRanges
	} else {
		for _, cidr := range svc.Spec.LoadBalancerSourceRanges {
			toBeRemoved := false
			for _, toRemove := range cidrsToRemove {
				if cidr == toRemove {
					toBeRemoved = true
					break
				}
			}
			if !toBeRemoved {
				newSourceRanges = append(newSourceRanges, cidr)
			}
		}
	}
	if len(cidrsToAdd) > 0 {
		newSourceRanges = append(newSourceRanges, cidrsToAdd...)
	}
	svc.Spec.LoadBalancerSourceRanges = newSourceRanges
	_, err = k8sClient.CoreV1().Services(svc.GetNamespace()).Update(ctx, svc, metav1.UpdateOptions{})
	return err
}

func (k *KubernetesClient) ListPodsCIDRs(ctx context.Context, clusterName string) ([]string, error) {
	k8sClient, err := k.getClient(clusterName)
	if err != nil {
		return nil, err
	}
	podCidrsMap := make(map[string]struct{})
	nodesList, err := k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, node := range nodesList.Items {
		podCidrsMap[node.Spec.PodCIDR] = struct{}{}
	}
	podCidrsList := make([]string, 0, len(podCidrsMap))
	for cidr := range podCidrsMap {
		podCidrsList = append(podCidrsList, cidr)
	}
	return podCidrsList, nil
}

func (k *KubernetesClient) ListServicesCIDRs(ctx context.Context, clusterName string) (string, error) {
	k8sClient, err := k.getClient(clusterName)
	if err != nil {
		return "", err
	}
	// apply Service with invalid IP, then parse error message for services CIDR block
	// based on https://stackoverflow.com/questions/44190607/how-do-you-find-the-cluster-service-cidr-of-a-kubernetes-cluster
	_, err = k8sClient.CoreV1().Services("default").Create(ctx, &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "dummy",
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "0.0.0.0",
		},
		Status: v1.ServiceStatus{},
	}, metav1.CreateOptions{})
	if err == nil {
		return "", fmt.Errorf("unexpectedly no error in creation of invalid service")
	}
	pattern := `The range of valid IPs is (\S+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(err.Error())

	if len(match) > 1 {
		ipRange := match[1]
		return ipRange, nil
	} else {
		return "", fmt.Errorf("no IP address range found")
	}
}

func (k *KubernetesClient) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}
