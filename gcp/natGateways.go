package gcp

import (
	"context"
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
	return nil, nil

	cc, err := compute.NewRoutersRESTClient(ctx)
	if err != nil {
		c.logger.Errorf("Failed to create client: %v", err)
		return natGateways, err
	}
	defer cc.Close()

	// Regions client
	var partialResponse bool = true

	rc, err := compute.NewRegionsRESTClient(ctx)
	if err != nil {
		c.logger.Errorf("Failed to create regions client: %v", err)
	}
	defer rc.Close()
	iter := rc.List(ctx, &computepb.ListRegionsRequest{Project: params.AccountId, ReturnPartialSuccess: &partialResponse}, nil)

	for {
		// List all routers in the project
		region, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			c.logger.Warnf("Region iteration error %v ", err)
			break
		}
		it := cc.List(ctx, &computepb.ListRoutersRequest{
			Project: params.AccountId,
			Region:  *region.Name,
		})
		c.logger.Infof("Listing NAT Gateway for GCP project %s ", params.AccountId)
	LOOP:
		for {
			router, err := it.Next()
			if err != nil {
				c.logger.Errorf("Error listing routers %v", err)
				break LOOP
			}
			c.logger.Infof("Router = %+v ", router)
			// A router that doesn't have Nats configuration is not a NAT router
			if len(router.Nats) == 0 {
				break
			}
			var subnetName, routerName string
			for _, nat := range router.Nats {
				// Assuming the first subnet is the primary one for simplicity
				if len(nat.Subnetworks) > 0 {
					subnetResourceID := nat.Subnetworks[0]
					subnetName = extractResourceID(*subnetResourceID.Name)
				}
			}
			if router.Name != nil {
				routerName = *router.Name
			}
			natGateway := types.NATGateway{
				ID:        string(*router.Id),
				Provider:  c.GetName(),
				Name:      routerName,
				AccountId: params.AccountId,
				VpcId:     *router.Network,
				Region:    *router.Region,
				State:     "Unknown", // Assuming ACTIVE
				//CreatedAt:    timestamppb.New(router.GetCreationTimestamp()),
				LastSyncTime: time.Now().Format(time.RFC3339),
				SubnetId:     subnetName,
			}

			c.logger.Infof("NAT Gateway: %+v\n", natGateway)
			natGateways = append(natGateways, natGateway)
		}
	}
	return natGateways, err
}

// Helper function to extract the last segment of a resource URL, typically the ID or name.
func extractResourceID(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
