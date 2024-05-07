
This document serves as a comprehensive guide for adding support for the listing a new cloud resource using awi-infra-guard software. It outlines the required modifications across various components of the application. This example shows how a VPCEndpoint was added, but can be used an example to add any other resource.


## Prerequisites

Familiarity with Go, gRPC, and Protocol Buffers.
Set up with the complete development environment of the awi-infra-guard project.

## File Modifications

 **1. Proto File**

Update proto/cloud.proto and proto/types.proto to include definitions and service methods for VPCEndpoints.

```
// Update to proto/cloud.proto
rpc ListVPCEndpoints (ListVPCEndpointsRequest) returns (ListVPCEndpointsResponse) {}

// Update to proto/types.proto
message VPCEndpoint {
    string id = 1;
    string name = 2;
    string provider = 3;
    string account_id = 4;
    string vpc_id = 5;
    string region = 6;
    string state = 7;
    string type = 8;
    string service_name = 9;
    repeated int32 route_table_ids = 10;
    repeated int32 subnet_ids = 11;
    map<string, string> labels = 12;
    google.protobuf.Timestamp created_at = 13;
    google.protobuf.Timestamp updated_at = 14;
    string last_sync_time = 15;
}
```
Run `make generate` in the repository root directory to generate language-specific generated protobuf files.

**2. Server Implementation**

Implement the gRPC server methods in server/server.go and update server/translate.go to handle data translation.

```
// server/server.go
func (s *Server) ListVPCEndpoints(ctx context.Context, in *infrapb.ListVPCEndpointsRequest) (*infrapb.ListVPCEndpointsResponse, error) {
    // Add your implementation here
}

// server/translate.go
func typesVPCEndpointsToGrpc(in []types.VPCEndpoint) []*infrapb.VPCEndpoint {
    // Add your translation logic here
}
```

**3. Synchronization Logic**

Define synchronization logic for VPCEndpoints in sync/sync.go.

```

// sync/sync.go
func (s *Syncer) syncVPCEndpoints() {
    // Add synchronization logic here
}

func (s *Syncer) syncVPCEndpoints() {
	genericCloudSync[*types.VPCEndpoint](s, types.VPCEndpointType, func(ctx context.Context, cloudProvider provider.CloudProvider, accountID string) ([]types.VPCEndpoint, error) {

		return cloudProvider.ListVPCEndpoints(ctx, &infrapb.ListVPCEndpointsRequest{AccountId: accountID})
	}, s.logger, s.dbClient.ListVPCEndpoints, s.dbClient.PutVPCEndpoint, s.dbClient.DeleteVPCEndpoint)
}
```

**4. Type Definitions**

Add or update VPCEndpoint struct definitions in type/types.go.

```
// type/types.go
type VPCEndpoint struct {
}

func (v *VPCEndpoint) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *VPCEndpoint) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPCEndpoint) GetProvider() string {
	return v.Provider
}

```

**5. Provider Interface**

Ensure the CloudProvider interface in provider/provider.go supports the ListVPCEndpoints method.

```
// provider/provider.go
ListVPCEndpoints(ctx context.Context, input *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error)
```

**6. Database Layer**

Implement methods for VPCEndpoints in db/db.go and db/db_strategy.go.

```
// db/db.go
interface Client {
    ListVPCEndpoints() ([]*types.VPCEndpoint, error)
    PutVPCEndpoint(*types.VPCEndpoint) error
    GetVPCEndpoint(string) (*types.VPCEndpoint, error)
    DeleteVPCEndpoint(string) error
}

// db/db_strategy.go
func (p *providerWithDB) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) {
    // Implement interaction logic here
}

```

**7. BoltDB Client Implementation**

Add methods to manage VPCEndpoints in BoltDB within boltdb/bolt_client.go.

```
// boltdb/bolt_client.go
func (client *boltClient) PutVPCEndpoint(vpce *types.VPCEndpoint) error {
    // Implement put logic here
}
```

**8. Cloud Provider Specific Logic**

Implement cloud-specific logic to list VPCEndpoints in aws/listVPCEndpoint.go, azure.go, gcp.go.

```
func (c *Client) ListVPCEndpoints(ctx context.Context, params *infrapb.ListVPCEndpointsRequest) ([]types.VPCEndpoint, error) { }
```
