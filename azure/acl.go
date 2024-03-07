package azure

import (
	"context"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types" // Adjust the import path according to your project structure
)

func (c *Client) ListACLs(ctx context.Context, input *infrapb.ListACLsRequest) ([]types.ACL, error) {
	
	// List All ACLS (Irrespective of VPC, subnet,NIC attached)
	acls, err := c.ListAllACLs(ctx, input)
	// Step 2: List all VNets and their subnets, noting any route table associations.
	va, err := ListVNetSubnetAssociations(ctx, input.AccountId, c.cred)
	if err != nil {
		return nil, err
	}
	// Step 3: Compare both lists and update the RouteTables list with VPCId and subnet.
	for i, acl := range acls {
		if association, ok := va.NsgAssociations[acl.ID]; ok {
			acls[i].VpcID = association.VNetID // Update with VNet ID
			//routeTables[i].Subnets = association.SubnetIDs // Update with associated subnet IDs
		}
		// Note: Route tables without no subnet (VPC) association will simply not be updated.
	}
	return acls, nil
}
