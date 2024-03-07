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
