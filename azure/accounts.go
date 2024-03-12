package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// ListAccounts returns a slice of types.Account, each representing an Azure subscription
func (c *Client) ListAccounts() []types.Account {
	client, err := armsubscriptions.NewClient(c.cred, nil)
	if err != nil {
		fmt.Printf("failed to create subscriptions client: %w", err)
		return nil
	}

	var accounts []types.Account // Declare a slice to hold the mapped accounts
	ctx := context.Background()
	pager := client.NewListPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			fmt.Printf("failed to get the next page of subscriptions: %w", err)
			return nil
		}

		// Iterate through the subscriptions and map them to types.Account
		for _, sub := range resp.SubscriptionListResult.Value {
			// Ensure ID and DisplayName are not nil before dereferencing
			var id, name string
			if sub.SubscriptionID != nil {
				id = *sub.SubscriptionID
			}
			if sub.DisplayName != nil {
				name = *sub.DisplayName
			}

			account := types.Account{
				ID:   id,
				Name: name,
			}
			accounts = append(accounts, account) // Append the mapped account to the slice
		}
	}

	return accounts
}
