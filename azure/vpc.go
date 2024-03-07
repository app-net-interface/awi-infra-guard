package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// ListVPC returns a slice of VPC objects for a given subscription
func (c *Client) ListVPC(ctx context.Context, params *infrapb.ListVPCRequest) ([]types.VPC, error) {

	var vpclist []types.VPC
	accounts := c.ListAccounts()

	for _, account := range accounts {
		//subscriptionID := strings.Split(account.ID, "/")[2]
		//fmt.Printf("Subscription ID : %s\n", subscriptionID)
		vnetClient, err := armnetwork.NewVirtualNetworksClient(account.ID, c.cred, nil)
		if err != nil {
			fmt.Printf("failed to create VNet client: %w", err)
			return nil, err
		}
		pager := vnetClient.NewListAllPager(nil)

		for pager.More() {
			resp, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get the next page of VNets: %w", err)
			}
			for _, vnet := range resp.VirtualNetworkListResult.Value {
				// Convert tags to labels
				labels := make(map[string]string)
				if vnet.Tags != nil {
					for k, v := range vnet.Tags {
						labels[k] = *v
					}
				}
				vpc := types.VPC{
					ID:        *vnet.ID,
					Name:      *vnet.Name,
					Region:    *vnet.Location,
					Labels:    labels,
					Provider:  "Azure",
					AccountID: account.ID,
					// LastSyncTime is not directly available from the VNet properties; you may need a custom approach
				}
				vpclist = append(vpclist, vpc)
			}
		}
	}

	return vpclist, nil
}

// VPCConnector interface implementation
func (c *Client) ConnectVPC(ctx context.Context, input types.SingleVPCConnectionParams) (types.SingleVPCConnectionOutput, error) {
	// TBD
	return types.SingleVPCConnectionOutput{}, nil
}

func (c *Client) ConnectVPCs(ctx context.Context, input types.VPCConnectionParams) (types.VPCConnectionOutput, error) {
	// TBD
	return types.VPCConnectionOutput{}, nil
}

func (c *Client) DisconnectVPC(ctx context.Context, input types.SingleVPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}

func (c *Client) DisconnectVPCs(ctx context.Context, input types.VPCDisconnectionParams) (types.VPCDisconnectionOutput, error) {
	// TBD
	return types.VPCDisconnectionOutput{}, nil
}
