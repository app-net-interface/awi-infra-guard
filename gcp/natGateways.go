package gcp

import (
	"context"
	"strconv"
	"strings"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/iterator"
)

func (c *Client) ListNATGateways(ctx context.Context, params *infrapb.ListNATGatewaysRequest) ([]types.NATGateway, error) {

	var natGateways []types.NATGateway

	client, err := compute.NewRoutersRESTClient(ctx)
	if err != nil {
		c.logger.Errorf("compute.NewRoutersRESTClient: %v", err)
		return natGateways, err
	}
	defer client.Close()

	// List all routers in the project
	req := &computepb.AggregatedListRoutersRequest{
		Project: params.AccountId,
	}

	it := client.AggregatedList(ctx, req)

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.logger.Errorf("Failed to list routers: %v", err)
			return natGateways, err
		}

		for _, router := range resp.Value.GetRouters() {
			//c.logger.Infof("Router [%d] = %+v %+v ", *router.Name, *&router.Id)
			for i, nat := range router.GetNats() {
				if nat != nil {
					var routerName, subnetName string
					if router.Name != nil {
						routerName = *router.Name
					}
					if len(nat.Subnetworks) > 0 {
						subnetResourceID := nat.Subnetworks[0]
						subnetName = extractResourceID(*subnetResourceID.Name)
					}
					natGateway := types.NATGateway{
						ID:        strconv.FormatUint(*router.Id, 10),
						Provider:  c.GetName(),
						Name:      routerName,
						AccountId: params.AccountId,
						VpcId:     *router.Network,
						Region:    *router.Region,
						State:     "Available", // Assuming ACTIVE
						//CreatedAt:    timestamppb.New(router.GetCreationTimestamp()),
						LastSyncTime: time.Now().Format(time.RFC3339),
						SubnetId:     subnetName,
					}
					c.logger.Debugf("GCP NAT GW [%d]  = %+v ", i, natGateway)
					natGateways = append(natGateways, natGateway)
				}
			}

		}
	}
	return natGateways, err
}

// Helper function to extract the last segment of a resource URL, typically the ID or name.
func extractResourceID(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
