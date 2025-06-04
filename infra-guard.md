# awi-infra-guard Additional Documentation

## 1. Overview

`awi-infra-guard` is an SDK and gRPC service designed to fetch infrastructure resource information and push updates to multiple infrastructure providers, including AWS, Google Cloud Platform (GCP), Azure, VMware, and ACI. It also supports operations on Kubernetes clusters.

Reference: [`README.md`](README.md)

## 2. Supported Infrastructure Providers

Currently, `awi-infra-guard` officially supports:
*   AWS (Amazon Web Services)
*   Google Cloud Platform (GCP)
*   Azure (Support in progress, some functionalities might be placeholders)

Reference: [`README.md`](README.md), `aws/aws.go`, `gcp/gcp.go`, `azure/azure.go`

## 3. Kubernetes Support

The system supports operations on Kubernetes clusters. It can automatically discover EKS (AWS) and GKE (GCP) clusters. Information for other Kubernetes clusters can optionally be provided via a kubeconfig file (typically `~/.kube/config`).

Reference: [`README.md`](README.md)

## 4. Usage

`awi-infra-guard` can be utilized in two main ways:
*   As an imported Go library.
*   As a standalone gRPC service.

### 4.1. Credentials Configuration

#### AWS Credentials
Set up your AWS credentials via the `~/.aws/credentials` and `~/.aws/config` files or by using environment variables as per the AWS SDK [guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials). Multiple AWS accounts are supported through profiles.

#### GCP Credentials
Setup application default credentials based on instructions from GCP [guide](https://cloud.google.com/docs/authentication/application-default-credentials). Multiple projects are supported; for instructions on how to specify them, check the "awi-infra-guard as a library" and "awi-infra-guard as a service" sections.

Reference: [`README.md`](README.md)

### 4.2. As a Go Library

Import the package `github.com/app-net-interface/awi-infra-guard`.

```sh
go get github.com/app-net-interface/awi-infra-guard@develop
```

Example usage:
```go
package main

import (
    "context"
    "fmt"

    "github.com/sirupsen/logrus"
    "github.com/app-net-interface/awi-infra-guard/provider"
    "github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb" // Assuming infrapb is the generated package
)

func main() {
    ctx := context.Background()
    providerStrategy := provider.NewRealProviderStrategy(ctx, logrus.New(), "")

    awsProvider, err := providerStrategy.GetProvider(context.TODO(), "aws")
    if err != nil {
        panic(err)
    }
    instances, err := awsProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in AWS:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }

    gcpProvider, err := providerStrategy.GetProvider(context.TODO(), "gcp")
    if err != nil {
        panic(err)
    }
    instances, err = gcpProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in GCP:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }
}
```
Reference: [`README.md`](README.md)

### 4.3. As a gRPC Service

Run the service using:
```sh
make run
```
The server typically listens on `[::]:50052`. You can interact with it using `grpc_cli` or a custom gRPC client.

Example `grpc_cli` calls:
```sh
# List AWS Instances in a specific VPC
grpc_cli call localhost:50052 ListInstances "provider: 'aws', vpc_id: 'vpc-04a1eaad3aa81310f'"

# List all discovered Kubernetes clusters
grpc_cli call localhost:50052 ListClusters ""

# List Pods in a specific EKS cluster
grpc_cli call localhost:50052 ListPods "cluster_name: 'eks-awi-demo'"
```
Reference: [`README.md`](README.md)

## 5. Docker Instructions

### Building and Pushing Image
```sh
# Build
make docker-build IMG=<your-repo>/<name>
# Push
make docker-push IMG=<your-repo>/<name>
# Build and Push
make docker-build docker-push IMG=<your-repo>/<name>
```

### Running Docker Image
Mount necessary configuration files into the container:
*   `/root/config/config.yaml`: Main configuration.
*   `/root/.aws/credentials`: AWS credentials.
*   `/app/gcp-key/gcp-key.json`: GCP service account key.
*   `/root/.kube/config`: Kubernetes configuration.

Reference: [`README.md`](README.md), `Makefile`

## 6. API Reference (gRPC)

The gRPC API is defined across several `.proto` files located in `grpc/proto`. The primary services are `CloudProviderService`, `KubernetesService`, and `AccessControlService`.

Reference: `grpc/proto/cloud.proto`, `grpc/proto/kubernetes.proto`, `grpc/proto/access_control.proto`, `grpc/proto/types.proto`, `grpc/proto/types_k8s.proto`

### 6.1. Core Functionalities & How It Works

`awi-infra-guard` operates by:

1.  **Resource Discovery and Synchronization:**
    *   It connects to configured cloud providers (AWS, GCP, etc.) and Kubernetes clusters.
    *   A `Syncer` component (see `sync/sync.go`) periodically fetches resource information (VPCs, instances, subnets, Kubernetes deployments, pods, etc.).
    *   This data is then stored in a local BoltDB database. This allows for faster subsequent reads and reduces direct API calls to providers.
    *   The database client logic can be seen in files like `db/bolt_client.go`.

2.  **VPC Indexing:**
    *   To efficiently query resources within a VPC and understand their relationships, `awi-infra-guard` builds and maintains a `VPCIndex`.
    *   The `VPCIndex` (defined in `grpc/proto/types.proto` and managed in `db/bolt_client_vpc_index.go`) stores lists of IDs for various resource types (instances, subnets, route tables, etc.) belonging to a specific VPC.
    *   This index is crucial for quickly assembling data for VPC-centric views, like connectivity graphs. The `SyncVPCIndexes` function in `db/bolt_client_vpc_index.go` is responsible for creating and updating these indexes.

3.  **Connectivity Graph Generation:**
    *   `awi-infra-guard` can generate connectivity graphs for VPCs and instances.
    *   **VPC Connectivity Graph:**
        *   `ListVpcGraphNodes`: Fetches all relevant resources within a VPC (instances, subnets, gateways, etc.) as nodes. It uses the `VPCIndex` to identify these resources. (Implementation: `db/db_strategy_graph.go`)
        *   `ListVpcGraphEdges`: Determines the relationships (edges) between these nodes (e.g., an instance is `LOCATED_IN` a subnet, a subnet `USES_ROUTE_TABLE`). (Implementation: `db/db_strategy_graph.go`)
    *   **Instance Connectivity Graph:**
        *   `GetInstanceConnectivityGraph`: Provides a focused view of resources directly and indirectly connected to a specific instance, including its network interface, subnet, security groups, route table, and the targets of its routes. (Implementation: `db/db_strategy_graph.go`)

4.  **Access Control Management:**
    *   The `AccessControlService` allows for programmatic modification of network security rules, such as adding inbound allow rules to VPCs or security groups based on various criteria.

### 6.2. Services and Methods

#### 6.2.1. `CloudProviderService`
*   Defined in: `grpc/proto/cloud.proto`
*   Purpose: Handles operations related to discovering and listing cloud provider resources.

**Key RPC Methods:**

*   `ListAccounts(ListAccountsRequest) returns (ListAccountsResponse)`
    *   Lists configured cloud provider accounts.
    *   Request: `ListAccountsRequest` (can be empty).
    *   Response: `ListAccountsResponse` (contains a list of `Account` messages).
*   `ListRegions(ListRegionsRequest) returns (ListRegionsResponse)`
    *   Lists regions available for a given cloud provider and account.
    *   Request: `ListRegionsRequest` (specifies `cloud_provider`, `account_id`, `creds`).
    *   Response: `ListRegionsResponse` (contains a list of `Region` messages).
*   `ListVPC(ListVPCRequest) returns (ListVPCResponse)`
    *   Lists Virtual Private Clouds (VPCs).
    *   Request: `ListVPCRequest` (specifies `cloud_provider`, `account_id`, `region`, `creds`).
    *   Response: `ListVPCResponse` (contains a list of `VPC` messages).
*   `GetVPCIndex(GetVPCIndexRequest) returns (GetVPCIndexResponse)`
    *   Retrieves an index of resources within a specific VPC.
    *   Request: `GetVPCIndexRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `GetVPCIndexResponse` (contains a `VPCIndex` message).
*   `ListInstances(ListInstancesRequest) returns (ListInstancesResponse)`
    *   Lists virtual machine instances.
    *   Request: `ListInstancesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `subnet_id`, `labels`, `creds`).
    *   Response: `ListInstancesResponse` (contains a list of `Instance` messages).
*   `ListSubnets(ListSubnetsRequest) returns (ListSubnetsResponse)`
    *   Lists subnets.
    *   Request: `ListSubnetsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListSubnetsResponse` (contains a list of `Subnet` messages).
*   `ListACLs(ListACLsRequest) returns (ListACLsResponse)`
    *   Lists Access Control Lists (Network ACLs).
    *   Request: `ListACLsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListACLsResponse` (contains a list of `ACL` messages).
*   `ListSecurityGroups(ListSecurityGroupsRequest) returns (ListSecurityGroupsResponse)`
    *   Lists security groups.
    *   Request: `ListSecurityGroupsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListSecurityGroupsResponse` (contains a list of `SecurityGroup` messages).
*   `ListRouteTables(ListRouteTablesRequest) returns (ListRouteTablesResponse)`
    *   Lists route tables.
    *   Request: `ListRouteTablesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListRouteTablesResponse` (contains a list of `RouteTable` messages).
*   `ListNATGateways(ListNATGatewaysRequest) returns (ListNATGatewaysResponse)`
    *   Lists NAT Gateways.
    *   Request: `ListNATGatewaysRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListNATGatewaysResponse` (contains a list of `NATGateway` messages).
*   `ListRouters(ListRoutersRequest) returns (ListRoutersResponse)`
    *   Lists routers (e.g., AWS Transit Gateways, GCP Cloud Routers).
    *   Request: `ListRoutersRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListRoutersResponse` (contains a list of `Router` messages).
*   `ListInternetGateways(ListInternetGatewaysRequest) returns (ListInternetGatewaysResponse)`
    *   Lists Internet Gateways.
    *   Request: `ListInternetGatewaysRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListInternetGatewaysResponse` (contains a list of `IGW` messages).
*   `ListVPCEndpoints(ListVPCEndpointsRequest) returns (ListVPCEndpointsResponse)`
    *   Lists VPC Endpoints.
    *   Request: `ListVPCEndpointsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVPCEndpointsResponse` (contains a list of `VPCEndpoint` messages).
    *   Reference: `List-New-Resource.md`
*   `ListPublicIPs(ListPublicIPsRequest) returns (ListPublicIPsResponse)`
    *   Lists public IP addresses.
    *   Request: `ListPublicIPsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListPublicIPsResponse` (contains a list of `PublicIP` messages).
*   `ListLBs(ListLBsRequest) returns (ListLBsResponse)`
    *   Lists Load Balancers.
    *   Request: `ListLBsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListLBsResponse` (contains a list of `LB` messages).
*   `ListNetworkInterfaces(ListNetworkInterfacesRequest) returns (ListNetworkInterfacesResponse)`
    *   Lists network interfaces.
    *   Request: `ListNetworkInterfacesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `instance_id`, `creds`).
    *   Response: `ListNetworkInterfacesResponse` (contains a list of `NetworkInterface` messages).
*   `ListKeyPairs(ListKeyPairsRequest) returns (ListKeyPairsResponse)`
    *   Lists SSH key pairs.
    *   Request: `ListKeyPairsRequest` (filters by `cloud_provider`, `account_id`, `region`, `creds`).
    *   Response: `ListKeyPairsResponse` (contains a list of `KeyPair` messages).
*   `ListVpcGraphNodes(ListVpcGraphNodesRequest) returns (ListVpcGraphNodesResponse)`
    *   Lists all resource nodes within a specified VPC for graph visualization.
    *   Request: `ListVpcGraphNodesRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVpcGraphNodesResponse` (contains a list of `VpcGraphNode` messages).
*   `ListVpcGraphEdges(ListVpcGraphEdgesRequest) returns (ListVpcGraphEdgesResponse)`
    *   Lists the relationships (edges) between resource nodes within a specified VPC.
    *   Request: `ListVpcGraphEdgesRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVpcGraphEdgesResponse` (contains a list of `VpcGraphEdge` messages).
*   `GetVpcConnectivityGraph(GetVpcConnectivityGraphRequest) returns (GetVpcConnectivityGraphResponse)`
    *   Combines nodes and edges to provide a full connectivity graph for a VPC.
    *   Request: `GetVpcConnectivityGraphRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `GetVpcConnectivityGraphResponse` (contains lists of `VpcGraphNode` and `VpcGraphEdge` messages).
*   `GetInstanceConnectivityGraph(GetInstanceConnectivityGraphRequest) returns (GetInstanceConnectivityGraphResponse)`
    *   Provides a connectivity graph centered around a specific instance.
    *   Request: `GetInstanceConnectivityGraphRequest` (specifies `cloud_provider`, `account_id`, `region`, `instance_id`, `creds`).
    *   Response: `GetInstanceConnectivityGraphResponse` (contains lists of `InstanceGraphNode` and `InstanceGraphEdge` messages).

#### 6.2.2. `KubernetesService`
*   Defined in: `grpc/proto/kubernetes.proto`
*   Purpose: Handles operations related to Kubernetes cluster resources.

**Key RPC Methods:**

*   `ListClusters(ListClustersRequest) returns (ListClustersResponse)`
    *   Lists all discovered/configured Kubernetes clusters.
    *   Request: `ListClustersRequest` (empty).
    *   Response: `ListClustersResponse` (contains a list of `Cluster` messages).
*   `ListNamespaces(ListNamespacesRequest) returns (ListNamespacesResponse)`
    *   Lists namespaces within a Kubernetes cluster.
    *   Request: `ListNamespacesRequest` (filters by `cluster_name`, `labels`).
    *   Response: `ListNamespacesResponse` (contains a list of `Namespace` messages).
*   `ListNodes(ListNodesRequest) returns (ListNodesResponse)`
    *   Lists nodes in a Kubernetes cluster.
    *   Request: `ListNodesRequest` (filters by `cluster_name`, `labels`).
    *   Response: `ListNodesResponse` (contains a list of `Node` messages from `types_k8s.proto`).
*   `ListPods(ListPodsRequest) returns (ListPodsResponse)`
    *   Lists pods in a Kubernetes cluster.
    *   Request: `ListPodsRequest` (filters by `cluster_name`, `namespace`, `labels`).
    *   Response: `ListPodsResponse` (contains a list of `Pod` messages).
*   `ListServices(ListServicesRequest) returns (ListServicesResponse)`
    *   Lists services in a Kubernetes cluster.
    *   Request: `ListServicesRequest` (filters by `cluster_name`, `namespace`, `labels`).
    *   Response: `ListServicesResponse` (contains a list of `K8sService` messages).
*   `ListPodsCIDRs(ListPodsCIDRsRequest) returns (ListPodsCIDRsResponse)`
    *   Lists the Pod CIDR blocks for a cluster.
    *   Request: `ListPodsCIDRsRequest` (specifies `cluster_name`).
    *   Response: `ListPodsCIDRsResponse` (contains a list of CIDR strings).
*   `ListServicesCIDRs(ListServicesCIDRsRequest) returns (ListServicesCIDRsResponse)`
    *   Lists the Service CIDR block for a cluster.
    *   Request: `ListServicesCIDRsRequest` (specifies `cluster_name`).
    *   Response: `ListServicesCIDRsResponse` (contains a CIDR string).

#### 6.2.3. `AccessControlService`
*   Defined in: `grpc/proto/access_control.proto`
*   Purpose: Manages network access control rules.

**Key RPC Methods:**

*   `AddInboundAllowRuleInVPC(AddInboundAllowRuleInVPCRequest) returns (AddInboundAllowRuleInVPCResponse)`
    *   Adds an inbound allow rule to all instances in a VPC, typically by modifying security groups.
    *   Request: `AddInboundAllowRuleInVPCRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleInVPCResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleByLabelsMatch(AddInboundAllowRuleByLabelsMatchRequest) returns (AddInboundAllowRuleByLabelsMatchResponse)`
    *   Adds an inbound allow rule to instances matching specified labels.
    *   Request: `AddInboundAllowRuleByLabelsMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `labels_to_match_instance`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleByLabelsMatchResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleBySubnetMatch(AddInboundAllowRuleBySubnetMatchRequest) returns (AddInboundAllowRuleBySubnetMatchResponse)`
    *   Adds an inbound allow rule to instances within a specified subnet.
    *   Request: `AddInboundAllowRuleBySubnetMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `subnet_id`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleBySubnetMatchResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleByInstanceIPMatch(AddInboundAllowRuleByInstanceIPMatchRequest) returns (AddInboundAllowRuleByInstanceIPMatchResponse)`
    *   Adds an inbound allow rule to specific instances identified by their IP addresses.
    *   Request: `AddInboundAllowRuleByInstanceIPMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `instance_ips`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleByInstanceIPMatchResponse` (contains `rule_id`, `matched_instances_ids`).

### 6.3. Key Data Types (Messages)

These are some of the fundamental data types used in the gRPC API, primarily defined in `grpc/proto/types.proto` and `grpc/proto/types_k8s.proto`.

#### 6.3.1. Cloud Resource Types (`types.proto`)

*   **`Account`**: Represents a cloud provider account.
    *   Fields: `id`, `name`, `provider`, `project_id` (for GCP).
*   **`Region`**: Represents a cloud provider region.
    *   Fields: `id`, `name`, `provider`, `account_id`.
*   **`VPC`**: Represents a Virtual Private Cloud.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `cidr_block`, `is_default`, `state`, `tags`, `last_sync_time`, `self_link`.
*   **`VPCIndex`**: An index of resources within a VPC.
    *   Fields: `vpc_id`, `provider`, `account_id`, `region`, `router_ids`, `instance_ids`, `subnet_ids`, `route_table_ids`, `nat_gateway_ids`, `igw_ids`, `security_group_ids`, `acl_ids`, `lb_ids`, `vpc_endpoint_ids`, `network_interface_ids`, `vpn_concentrator_ids`, `public_ip_ids`, `cluster_ids`.
*   **`Instance`**: Represents a virtual machine instance.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `zone`, `vpc_id`, `subnet_id`, `private_ip`, `public_ip`, `state`, `type` (instance type), `image_id`, `key_name`, `security_group_ids`, `network_interface_ids`, `tags`, `labels`, `last_sync_time`, `self_link`.
*   **`Subnet`**: Represents a subnet within a VPC.
    *   Fields: `subnet_id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `zone`, `cidr_block`, `available_ips`, `state`, `is_default`, `map_public_ip_on_launch`, `tags`, `route_table_id`, `acl_id`, `last_sync_time`, `self_link`.
*   **`ACL` (Access Control List)**: Represents a network ACL.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `rules` (list of `ACLRule`), `is_default`, `tags`, `associated_subnet_ids`, `last_sync_time`, `self_link`.
*   **`ACLRule`**: A rule within an ACL.
    *   Fields: `number`, `protocol`, `action` (allow/deny), `direction` (inbound/outbound), `port_range`, `source_ranges`, `destination_ranges`.
*   **`SecurityGroup`**: Represents a security group.
    *   Fields: `id`, `name`, `description`, `vpc_id`, `provider`, `account_id`, `region`, `rules` (list of `SecurityGroupRule`), `tags`, `last_sync_time`, `self_link`.
*   **`SecurityGroupRule`**: A rule within a security group.
    *   Fields: `id`, `direction` (inbound/outbound), `protocol`, `port_range`, `source` (CIDRs, SG IDs, Prefix List IDs), `description`.
*   **`RouteTable`**: Represents a route table.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `routes` (list of `Route`), `is_main`, `tags`, `associated_subnet_ids`, `last_sync_time`, `self_link`.
*   **`Route`**: A route within a route table.
    *   Fields: `destination` (CIDR), `target` (e.g., gateway ID, instance ID, NAT GW ID), `status` (active/blackhole), `next_hop_type`.
*   **`NATGateway`**: Represents a NAT Gateway.
    *   Fields: `id`, `name`, `vpc_id`, `subnet_id`, `provider`, `account_id`, `region`, `state`, `public_ip_ids`, `private_ip_ids`, `network_interface_ids`, `tags`, `last_sync_time`, `self_link`.
*   **`Router`**: Represents a generic router (e.g., TGW, Cloud Router).
    *   Fields: `id`, `name`, `vpc_id` (optional), `provider`, `account_id`, `region`, `state`, `tags`, `last_sync_time`, `self_link`.
*   **`IGW` (Internet Gateway)**: Represents an Internet Gateway.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `attached_vpc_id`, `tags`, `last_sync_time`, `self_link`.
*   **`VPCEndpoint`**: Represents a VPC Endpoint.
    *   Fields: `id`, `name`, `vpc_id`, `service_name`, `provider`, `account_id`, `region`, `state`, `type` (Interface/Gateway), `subnet_ids`, `route_table_ids`, `security_group_ids`, `network_interface_ids`, `dns_entries`, `tags`, `last_sync_time`, `self_link`.
*   **`PublicIP`**: Represents a public IP address.
    *   Fields: `id`, `name`, `ip_address`, `provider`, `account_id`, `region`, `vpc_id` (optional), `instance_id` (optional), `network_interface_id` (optional), `nat_gateway_id` (optional), `lb_id` (optional), `allocation_id` (AWS), `domain` (VPC/Standard), `association_id` (AWS), `tags`, `last_sync_time`, `self_link`.
*   **`LB` (Load Balancer)**: Represents a load balancer.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `type` (application/network/gateway), `scheme` (internal/internet-facing), `state`, `dns_name`, `listeners`, `target_groups`, `availability_zones`, `subnet_ids`, `security_group_ids`, `tags`, `last_sync_time`, `self_link`.
*   **`NetworkInterface`**: Represents a network interface.
    *   Fields: `id`, `name`, `vpc_id`, `subnet_id`, `instance_id` (optional), `provider`, `account_id`, `region`, `zone`, `mac_address`, `private_ip_address`, `public_ip_address` (optional), `security_group_ids`, `status`, `description`, `attachment_id` (if attached), `tags`, `last_sync_time`, `self_link`.
*   **`KeyPair`**: Represents an SSH key pair.
    *   Fields: `id`, `name`, `fingerprint`, `provider`, `account_id`, `region`, `tags`, `last_sync_time`, `self_link`.

#### 6.3.2. Kubernetes Resource Types (`types_k8s.proto`)

*   **`Cluster`**: Represents a Kubernetes cluster.
    *   Fields: `name`, `provider` (e.g., EKS, GKE, Kind), `account_id`, `region`, `vpc_id`, `version`, `endpoint`, `status`, `tags`, `last_sync_time`, `self_link`.
*   **`Node`**: Represents a Kubernetes node.
    *   Fields: `cluster`, `name`, `namespace`, `addresses` (internal/external IPs), `project` (GCP project), `last_sync_time`, `self_link`, `labels`, `conditions`, `capacity`, `allocatable`, `node_info`.
*   **`Namespace`**: Represents a Kubernetes namespace.
    *   Fields: `cluster`, `name`, `labels`, `project` (GCP project), `last_sync_time`, `self_link`, `status`, `annotations`.
*   **`Pod`**: Represents a Kubernetes pod.
    *   Fields: `cluster`, `namespace`, `name`, `ip_address`, `node_name`, `status`, `labels`, `containers`, `owner_references`, `creation_timestamp`, `last_sync_time`, `self_link`.
*   **`K8sService`**: Represents a Kubernetes service.
    *   Fields: `cluster`, `namespace`, `name`, `type` (ClusterIP, NodePort, LoadBalancer), `cluster_ip`, `external_ips`, `ports`, `selector`, `labels`, `load_balancer_ingress`, `creation_timestamp`, `last_sync_time`, `self_link`.

#### 6.3.3. Graph Types (`types.proto`)

*   **`VpcGraphNode`**: Represents a node in a VPC connectivity graph.
    *   Fields: `id`, `name`, `type` (e.g., "instance", "subnet"), `provider`, `account_id`, `region`, `vpc_id`, `properties` (map of string to string for additional details like CIDR, state, rules).
*   **`VpcGraphEdge`**: Represents an edge (relationship) in a VPC connectivity graph.
    *   Fields: `id`, `source_node_id`, `target_node_id`, `label` (describes the relationship, e.g., "LOCATED_IN", "ROUTES_TO", "ASSOCIATED_WITH"), `properties` (map of string to string).
*   **`InstanceGraphNode`**: Represents a node in an instance-centric connectivity graph.
    *   Fields: Similar to `VpcGraphNode`, tailored for instance context.
*   **`InstanceGraphEdge`**: Represents an edge in an instance-centric connectivity graph.
    *   Fields: Similar to `VpcGraphEdge`.

#### 6.3.4. Common and Utility Types

*   **`Credentials`**: Used to pass authentication details for providers.
    *   Fields: `aws_role`, `gcp_role`, `azure_role`, `aws_user_auth`, `gcp_user_auth`, `azure_user_auth`.
*   **`Error`**: Standard error message structure.
*   `google.protobuf.Timestamp`: Standard timestamp.
*   `google.protobuf.FieldMask`: Used for partial updates (not extensively shown in current examples but good practice for update RPCs).

This expanded documentation should provide a more comprehensive understanding of `awi-infra-guard`.
```# awi-infra-guard Documentation

## 1. Overview

`awi-infra-guard` is an SDK and gRPC service designed to fetch infrastructure resource information and push updates to multiple infrastructure providers, including AWS, Google Cloud Platform (GCP), Azure, VMware, and ACI. It also supports operations on Kubernetes clusters.

Reference: [`README.md`](README.md)

## 2. Supported Infrastructure Providers

Currently, `awi-infra-guard` officially supports:
*   AWS (Amazon Web Services)
*   Google Cloud Platform (GCP)
*   Azure (Support in progress, some functionalities might be placeholders)

Reference: [`README.md`](README.md), `aws/aws.go`, `gcp/gcp.go`, `azure/azure.go`

## 3. Kubernetes Support

The system supports operations on Kubernetes clusters. It can automatically discover EKS (AWS) and GKE (GCP) clusters. Information for other Kubernetes clusters can optionally be provided via a kubeconfig file (typically `~/.kube/config`).

Reference: [`README.md`](README.md)

## 4. Usage

`awi-infra-guard` can be utilized in two main ways:
*   As an imported Go library.
*   As a standalone gRPC service.

### 4.1. Credentials Configuration

#### AWS Credentials
Set up your AWS credentials via the `~/.aws/credentials` and `~/.aws/config` files or by using environment variables as per the AWS SDK [guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials). Multiple AWS accounts are supported through profiles.

#### GCP Credentials
Setup application default credentials based on instructions from GCP [guide](https://cloud.google.com/docs/authentication/application-default-credentials). Multiple projects are supported; for instructions on how to specify them, check the "awi-infra-guard as a library" and "awi-infra-guard as a service" sections.

Reference: [`README.md`](README.md)

### 4.2. As a Go Library

Import the package `github.com/app-net-interface/awi-infra-guard`.

```sh
go get github.com/app-net-interface/awi-infra-guard@develop
```

Example usage:
```go
package main

import (
    "context"
    "fmt"

    "github.com/sirupsen/logrus"
    "github.com/app-net-interface/awi-infra-guard/provider"
    "github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb" // Assuming infrapb is the generated package
)

func main() {
    ctx := context.Background()
    providerStrategy := provider.NewRealProviderStrategy(ctx, logrus.New(), "")

    awsProvider, err := providerStrategy.GetProvider(context.TODO(), "aws")
    if err != nil {
        panic(err)
    }
    instances, err := awsProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in AWS:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }

    gcpProvider, err := providerStrategy.GetProvider(context.TODO(), "gcp")
    if err != nil {
        panic(err)
    }
    instances, err = gcpProvider.ListInstances(context.TODO(), &infrapb.ListInstancesRequest{})
    if err != nil {
        panic(err)
    }
    fmt.Println("Instances in GCP:")
    for _, instance := range instances {
        fmt.Println(instance.VPCID, instance.Name)
    }
}
```
Reference: [`README.md`](README.md)

### 4.3. As a gRPC Service

Run the service using:
```sh
make run
```
The server typically listens on `[::]:50052`. You can interact with it using `grpc_cli` or a custom gRPC client.

Example `grpc_cli` calls:
```sh
# List AWS Instances in a specific VPC
grpc_cli call localhost:50052 ListInstances "provider: 'aws', vpc_id: 'vpc-04a1eaad3aa81310f'"

# List all discovered Kubernetes clusters
grpc_cli call localhost:50052 ListClusters ""

# List Pods in a specific EKS cluster
grpc_cli call localhost:50052 ListPods "cluster_name: 'eks-awi-demo'"
```
Reference: [`README.md`](README.md)

## 5. Docker Instructions

### Building and Pushing Image
```sh
# Build
make docker-build IMG=<your-repo>/<name>
# Push
make docker-push IMG=<your-repo>/<name>
# Build and Push
make docker-build docker-push IMG=<your-repo>/<name>
```

### Running Docker Image
Mount necessary configuration files into the container:
*   `/root/config/config.yaml`: Main configuration.
*   `/root/.aws/credentials`: AWS credentials.
*   `/app/gcp-key/gcp-key.json`: GCP service account key.
*   `/root/.kube/config`: Kubernetes configuration.

Reference: [`README.md`](README.md), `Makefile`

## 6. API Reference (gRPC)

The gRPC API is defined across several `.proto` files located in `grpc/proto`. The primary services are `CloudProviderService`, `KubernetesService`, and `AccessControlService`.

Reference: `grpc/proto/cloud.proto`, `grpc/proto/kubernetes.proto`, `grpc/proto/access_control.proto`, `grpc/proto/types.proto`, `grpc/proto/types_k8s.proto`

### 6.1. Core Functionalities & How It Works

`awi-infra-guard` operates by:

1.  **Resource Discovery and Synchronization:**
    *   It connects to configured cloud providers (AWS, GCP, etc.) and Kubernetes clusters.
    *   A `Syncer` component (see `sync/sync.go`) periodically fetches resource information (VPCs, instances, subnets, Kubernetes deployments, pods, etc.).
    *   This data is then stored in a local BoltDB database. This allows for faster subsequent reads and reduces direct API calls to providers.
    *   The database client logic can be seen in files like `db/bolt_client.go`.

2.  **VPC Indexing:**
    *   To efficiently query resources within a VPC and understand their relationships, `awi-infra-guard` builds and maintains a `VPCIndex`.
    *   The `VPCIndex` (defined in `grpc/proto/types.proto` and managed in `db/bolt_client_vpc_index.go`) stores lists of IDs for various resource types (instances, subnets, route tables, etc.) belonging to a specific VPC.
    *   This index is crucial for quickly assembling data for VPC-centric views, like connectivity graphs. The `SyncVPCIndexes` function in `db/bolt_client_vpc_index.go` is responsible for creating and updating these indexes.

3.  **Connectivity Graph Generation:**
    *   `awi-infra-guard` can generate connectivity graphs for VPCs and instances.
    *   **VPC Connectivity Graph:**
        *   `ListVpcGraphNodes`: Fetches all relevant resources within a VPC (instances, subnets, gateways, etc.) as nodes. It uses the `VPCIndex` to identify these resources. (Implementation: `db/db_strategy_graph.go`)
        *   `ListVpcGraphEdges`: Determines the relationships (edges) between these nodes (e.g., an instance is `LOCATED_IN` a subnet, a subnet `USES_ROUTE_TABLE`). (Implementation: `db/db_strategy_graph.go`)
    *   **Instance Connectivity Graph:**
        *   `GetInstanceConnectivityGraph`: Provides a focused view of resources directly and indirectly connected to a specific instance, including its network interface, subnet, security groups, route table, and the targets of its routes. (Implementation: `db/db_strategy_graph.go`)

4.  **Access Control Management:**
    *   The `AccessControlService` allows for programmatic modification of network security rules, such as adding inbound allow rules to VPCs or security groups based on various criteria.

### 6.2. Services and Methods

#### 6.2.1. `CloudProviderService`
*   Defined in: `grpc/proto/cloud.proto`
*   Purpose: Handles operations related to discovering and listing cloud provider resources.

**Key RPC Methods:**

*   `ListAccounts(ListAccountsRequest) returns (ListAccountsResponse)`
    *   Lists configured cloud provider accounts.
    *   Request: `ListAccountsRequest` (can be empty).
    *   Response: `ListAccountsResponse` (contains a list of `Account` messages).
*   `ListRegions(ListRegionsRequest) returns (ListRegionsResponse)`
    *   Lists regions available for a given cloud provider and account.
    *   Request: `ListRegionsRequest` (specifies `cloud_provider`, `account_id`, `creds`).
    *   Response: `ListRegionsResponse` (contains a list of `Region` messages).
*   `ListVPC(ListVPCRequest) returns (ListVPCResponse)`
    *   Lists Virtual Private Clouds (VPCs).
    *   Request: `ListVPCRequest` (specifies `cloud_provider`, `account_id`, `region`, `creds`).
    *   Response: `ListVPCResponse` (contains a list of `VPC` messages).
*   `GetVPCIndex(GetVPCIndexRequest) returns (GetVPCIndexResponse)`
    *   Retrieves an index of resources within a specific VPC.
    *   Request: `GetVPCIndexRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `GetVPCIndexResponse` (contains a `VPCIndex` message).
*   `ListInstances(ListInstancesRequest) returns (ListInstancesResponse)`
    *   Lists virtual machine instances.
    *   Request: `ListInstancesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `subnet_id`, `labels`, `creds`).
    *   Response: `ListInstancesResponse` (contains a list of `Instance` messages).
*   `ListSubnets(ListSubnetsRequest) returns (ListSubnetsResponse)`
    *   Lists subnets.
    *   Request: `ListSubnetsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListSubnetsResponse` (contains a list of `Subnet` messages).
*   `ListACLs(ListACLsRequest) returns (ListACLsResponse)`
    *   Lists Access Control Lists (Network ACLs).
    *   Request: `ListACLsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListACLsResponse` (contains a list of `ACL` messages).
*   `ListSecurityGroups(ListSecurityGroupsRequest) returns (ListSecurityGroupsResponse)`
    *   Lists security groups.
    *   Request: `ListSecurityGroupsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListSecurityGroupsResponse` (contains a list of `SecurityGroup` messages).
*   `ListRouteTables(ListRouteTablesRequest) returns (ListRouteTablesResponse)`
    *   Lists route tables.
    *   Request: `ListRouteTablesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListRouteTablesResponse` (contains a list of `RouteTable` messages).
*   `ListNATGateways(ListNATGatewaysRequest) returns (ListNATGatewaysResponse)`
    *   Lists NAT Gateways.
    *   Request: `ListNATGatewaysRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListNATGatewaysResponse` (contains a list of `NATGateway` messages).
*   `ListRouters(ListRoutersRequest) returns (ListRoutersResponse)`
    *   Lists routers (e.g., AWS Transit Gateways, GCP Cloud Routers).
    *   Request: `ListRoutersRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListRoutersResponse` (contains a list of `Router` messages).
*   `ListInternetGateways(ListInternetGatewaysRequest) returns (ListInternetGatewaysResponse)`
    *   Lists Internet Gateways.
    *   Request: `ListInternetGatewaysRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListInternetGatewaysResponse` (contains a list of `IGW` messages).
*   `ListVPCEndpoints(ListVPCEndpointsRequest) returns (ListVPCEndpointsResponse)`
    *   Lists VPC Endpoints.
    *   Request: `ListVPCEndpointsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVPCEndpointsResponse` (contains a list of `VPCEndpoint` messages).
    *   Reference: `List-New-Resource.md`
*   `ListPublicIPs(ListPublicIPsRequest) returns (ListPublicIPsResponse)`
    *   Lists public IP addresses.
    *   Request: `ListPublicIPsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListPublicIPsResponse` (contains a list of `PublicIP` messages).
*   `ListLBs(ListLBsRequest) returns (ListLBsResponse)`
    *   Lists Load Balancers.
    *   Request: `ListLBsRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListLBsResponse` (contains a list of `LB` messages).
*   `ListNetworkInterfaces(ListNetworkInterfacesRequest) returns (ListNetworkInterfacesResponse)`
    *   Lists network interfaces.
    *   Request: `ListNetworkInterfacesRequest` (filters by `cloud_provider`, `account_id`, `region`, `vpc_id`, `instance_id`, `creds`).
    *   Response: `ListNetworkInterfacesResponse` (contains a list of `NetworkInterface` messages).
*   `ListKeyPairs(ListKeyPairsRequest) returns (ListKeyPairsResponse)`
    *   Lists SSH key pairs.
    *   Request: `ListKeyPairsRequest` (filters by `cloud_provider`, `account_id`, `region`, `creds`).
    *   Response: `ListKeyPairsResponse` (contains a list of `KeyPair` messages).
*   `ListVpcGraphNodes(ListVpcGraphNodesRequest) returns (ListVpcGraphNodesResponse)`
    *   Lists all resource nodes within a specified VPC for graph visualization.
    *   Request: `ListVpcGraphNodesRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVpcGraphNodesResponse` (contains a list of `VpcGraphNode` messages).
*   `ListVpcGraphEdges(ListVpcGraphEdgesRequest) returns (ListVpcGraphEdgesResponse)`
    *   Lists the relationships (edges) between resource nodes within a specified VPC.
    *   Request: `ListVpcGraphEdgesRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `ListVpcGraphEdgesResponse` (contains a list of `VpcGraphEdge` messages).
*   `GetVpcConnectivityGraph(GetVpcConnectivityGraphRequest) returns (GetVpcConnectivityGraphResponse)`
    *   Combines nodes and edges to provide a full connectivity graph for a VPC.
    *   Request: `GetVpcConnectivityGraphRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `creds`).
    *   Response: `GetVpcConnectivityGraphResponse` (contains lists of `VpcGraphNode` and `VpcGraphEdge` messages).
*   `GetInstanceConnectivityGraph(GetInstanceConnectivityGraphRequest) returns (GetInstanceConnectivityGraphResponse)`
    *   Provides a connectivity graph centered around a specific instance.
    *   Request: `GetInstanceConnectivityGraphRequest` (specifies `cloud_provider`, `account_id`, `region`, `instance_id`, `creds`).
    *   Response: `GetInstanceConnectivityGraphResponse` (contains lists of `InstanceGraphNode` and `InstanceGraphEdge` messages).

#### 6.2.2. `KubernetesService`
*   Defined in: `grpc/proto/kubernetes.proto`
*   Purpose: Handles operations related to Kubernetes cluster resources.

**Key RPC Methods:**

*   `ListClusters(ListClustersRequest) returns (ListClustersResponse)`
    *   Lists all discovered/configured Kubernetes clusters.
    *   Request: `ListClustersRequest` (empty).
    *   Response: `ListClustersResponse` (contains a list of `Cluster` messages).
*   `ListNamespaces(ListNamespacesRequest) returns (ListNamespacesResponse)`
    *   Lists namespaces within a Kubernetes cluster.
    *   Request: `ListNamespacesRequest` (filters by `cluster_name`, `labels`).
    *   Response: `ListNamespacesResponse` (contains a list of `Namespace` messages).
*   `ListNodes(ListNodesRequest) returns (ListNodesResponse)`
    *   Lists nodes in a Kubernetes cluster.
    *   Request: `ListNodesRequest` (filters by `cluster_name`, `labels`).
    *   Response: `ListNodesResponse` (contains a list of `Node` messages from `types_k8s.proto`).
*   `ListPods(ListPodsRequest) returns (ListPodsResponse)`
    *   Lists pods in a Kubernetes cluster.
    *   Request: `ListPodsRequest` (filters by `cluster_name`, `namespace`, `labels`).
    *   Response: `ListPodsResponse` (contains a list of `Pod` messages).
*   `ListServices(ListServicesRequest) returns (ListServicesResponse)`
    *   Lists services in a Kubernetes cluster.
    *   Request: `ListServicesRequest` (filters by `cluster_name`, `namespace`, `labels`).
    *   Response: `ListServicesResponse` (contains a list of `K8sService` messages).
*   `ListPodsCIDRs(ListPodsCIDRsRequest) returns (ListPodsCIDRsResponse)`
    *   Lists the Pod CIDR blocks for a cluster.
    *   Request: `ListPodsCIDRsRequest` (specifies `cluster_name`).
    *   Response: `ListPodsCIDRsResponse` (contains a list of CIDR strings).
*   `ListServicesCIDRs(ListServicesCIDRsRequest) returns (ListServicesCIDRsResponse)`
    *   Lists the Service CIDR block for a cluster.
    *   Request: `ListServicesCIDRsRequest` (specifies `cluster_name`).
    *   Response: `ListServicesCIDRsResponse` (contains a CIDR string).

#### 6.2.3. `AccessControlService`
*   Defined in: `grpc/proto/access_control.proto`
*   Purpose: Manages network access control rules.

**Key RPC Methods:**

*   `AddInboundAllowRuleInVPC(AddInboundAllowRuleInVPCRequest) returns (AddInboundAllowRuleInVPCResponse)`
    *   Adds an inbound allow rule to all instances in a VPC, typically by modifying security groups.
    *   Request: `AddInboundAllowRuleInVPCRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleInVPCResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleByLabelsMatch(AddInboundAllowRuleByLabelsMatchRequest) returns (AddInboundAllowRuleByLabelsMatchResponse)`
    *   Adds an inbound allow rule to instances matching specified labels.
    *   Request: `AddInboundAllowRuleByLabelsMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `labels_to_match_instance`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleByLabelsMatchResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleBySubnetMatch(AddInboundAllowRuleBySubnetMatchRequest) returns (AddInboundAllowRuleBySubnetMatchResponse)`
    *   Adds an inbound allow rule to instances within a specified subnet.
    *   Request: `AddInboundAllowRuleBySubnetMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `subnet_id`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleBySubnetMatchResponse` (contains `rule_id`, `matched_instances_ids`).
*   `AddInboundAllowRuleByInstanceIPMatch(AddInboundAllowRuleByInstanceIPMatchRequest) returns (AddInboundAllowRuleByInstanceIPMatchResponse)`
    *   Adds an inbound allow rule to specific instances identified by their IP addresses.
    *   Request: `AddInboundAllowRuleByInstanceIPMatchRequest` (specifies `cloud_provider`, `account_id`, `region`, `vpc_id`, `instance_ips`, `source_ip_ranges`, `ports`, `creds`).
    *   Response: `AddInboundAllowRuleByInstanceIPMatchResponse` (contains `rule_id`, `matched_instances_ids`).

### 6.3. Key Data Types (Messages)

These are some of the fundamental data types used in the gRPC API, primarily defined in `grpc/proto/types.proto` and `grpc/proto/types_k8s.proto`.

#### 6.3.1. Cloud Resource Types (`types.proto`)

*   **`Account`**: Represents a cloud provider account.
    *   Fields: `id`, `name`, `provider`, `project_id` (for GCP).
*   **`Region`**: Represents a cloud provider region.
    *   Fields: `id`, `name`, `provider`, `account_id`.
*   **`VPC`**: Represents a Virtual Private Cloud.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `cidr_block`, `is_default`, `state`, `tags`, `last_sync_time`, `self_link`.
*   **`VPCIndex`**: An index of resources within a VPC.
    *   Fields: `vpc_id`, `provider`, `account_id`, `region`, `router_ids`, `instance_ids`, `subnet_ids`, `route_table_ids`, `nat_gateway_ids`, `igw_ids`, `security_group_ids`, `acl_ids`, `lb_ids`, `vpc_endpoint_ids`, `network_interface_ids`, `vpn_concentrator_ids`, `public_ip_ids`, `cluster_ids`.
*   **`Instance`**: Represents a virtual machine instance.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `zone`, `vpc_id`, `subnet_id`, `private_ip`, `public_ip`, `state`, `type` (instance type), `image_id`, `key_name`, `security_group_ids`, `network_interface_ids`, `tags`, `labels`, `last_sync_time`, `self_link`.
*   **`Subnet`**: Represents a subnet within a VPC.
    *   Fields: `subnet_id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `zone`, `cidr_block`, `available_ips`, `state`, `is_default`, `map_public_ip_on_launch`, `tags`, `route_table_id`, `acl_id`, `last_sync_time`, `self_link`.
*   **`ACL` (Access Control List)**: Represents a network ACL.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `rules` (list of `ACLRule`), `is_default`, `tags`, `associated_subnet_ids`, `last_sync_time`, `self_link`.
*   **`ACLRule`**: A rule within an ACL.
    *   Fields: `number`, `protocol`, `action` (allow/deny), `direction` (inbound/outbound), `port_range`, `source_ranges`, `destination_ranges`.
*   **`SecurityGroup`**: Represents a security group.
    *   Fields: `id`, `name`, `description`, `vpc_id`, `provider`, `account_id`, `region`, `rules` (list of `SecurityGroupRule`), `tags`, `last_sync_time`, `self_link`.
*   **`SecurityGroupRule`**: A rule within a security group.
    *   Fields: `id`, `direction` (inbound/outbound), `protocol`, `port_range`, `source` (CIDRs, SG IDs, Prefix List IDs), `description`.
*   **`RouteTable`**: Represents a route table.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `routes` (list of `Route`), `is_main`, `tags`, `associated_subnet_ids`, `last_sync_time`, `self_link`.
*   **`Route`**: A route within a route table.
    *   Fields: `destination` (CIDR), `target` (e.g., gateway ID, instance ID, NAT GW ID), `status` (active/blackhole), `next_hop_type`.
*   **`NATGateway`**: Represents a NAT Gateway.
    *   Fields: `id`, `name`, `vpc_id`, `subnet_id`, `provider`, `account_id`, `region`, `state`, `public_ip_ids`, `private_ip_ids`, `network_interface_ids`, `tags`, `last_sync_time`, `self_link`.
*   **`Router`**: Represents a generic router (e.g., TGW, Cloud Router).
    *   Fields: `id`, `name`, `vpc_id` (optional), `provider`, `account_id`, `region`, `state`, `tags`, `last_sync_time`, `self_link`.
*   **`IGW` (Internet Gateway)**: Represents an Internet Gateway.
    *   Fields: `id`, `name`, `provider`, `account_id`, `region`, `attached_vpc_id`, `tags`, `last_sync_time`, `self_link`.
*   **`VPCEndpoint`**: Represents a VPC Endpoint.
    *   Fields: `id`, `name`, `vpc_id`, `service_name`, `provider`, `account_id`, `region`, `state`, `type` (Interface/Gateway), `subnet_ids`, `route_table_ids`, `security_group_ids`, `network_interface_ids`, `dns_entries`, `tags`, `last_sync_time`, `self_link`.
*   **`PublicIP`**: Represents a public IP address.
    *   Fields: `id`, `name`, `ip_address`, `provider`, `account_id`, `region`, `vpc_id` (optional), `instance_id` (optional), `network_interface_id` (optional), `nat_gateway_id` (optional), `lb_id` (optional), `allocation_id` (AWS), `domain` (VPC/Standard), `association_id` (AWS), `tags`, `last_sync_time`, `self_link`.
*   **`LB` (Load Balancer)**: Represents a load balancer.
    *   Fields: `id`, `name`, `vpc_id`, `provider`, `account_id`, `region`, `type` (application/network/gateway), `scheme` (internal/internet-facing), `state`, `dns_name`, `listeners`, `target_groups`, `availability_zones`, `subnet_ids`, `security_group_ids`, `tags`, `last_sync_time`, `self_link`.
*   **`NetworkInterface`**: Represents a network interface.
    *   Fields: `id`, `name`, `vpc_id`, `subnet_id`, `instance_id` (optional), `provider`, `account_id`, `region`, `zone`, `mac_address`, `private_ip_address`, `public_ip_address` (optional), `security_group_ids`, `status`, `description`, `attachment_id` (if attached), `tags`, `last_sync_time`, `self_link`.
*   **`KeyPair`**: Represents an SSH key pair.
    *   Fields: `id`, `name`, `fingerprint`, `provider`, `account_id`, `region`, `tags`, `last_sync_time`, `self_link`.

#### 6.3.2. Kubernetes Resource Types (`types_k8s.proto`)

*   **`Cluster`**: Represents a Kubernetes cluster.
    *   Fields: `name`, `provider` (e.g., EKS, GKE, Kind), `account_id`, `region`, `vpc_id`, `version`, `endpoint`, `status`, `tags`, `last_sync_time`, `self_link`.
*   **`Node`**: Represents a Kubernetes node.
    *   Fields: `cluster`, `name`, `namespace`, `addresses` (internal/external IPs), `project` (GCP project), `last_sync_time`, `self_link`, `labels`, `conditions`, `capacity`, `allocatable`, `node_info`.
*   **`Namespace`**: Represents a Kubernetes namespace.
    *   Fields: `cluster`, `name`, `labels`, `project` (GCP project), `last_sync_time`, `self_link`, `status`, `annotations`.
*   **`Pod`**: Represents a Kubernetes pod.
    *   Fields: `cluster`, `namespace`, `name`, `ip_address`, `node_name`, `status`, `labels`, `containers`, `owner_references`, `creation_timestamp`, `last_sync_time`, `self_link`.
*   **`K8sService`**: Represents a Kubernetes service.
    *   Fields: `cluster`, `namespace`, `name`, `type` (ClusterIP, NodePort, LoadBalancer), `cluster_ip`, `external_ips`, `ports`, `selector`, `labels`, `load_balancer_ingress`, `creation_timestamp`, `last_sync_time`, `self_link`.

#### 6.3.3. Graph Types (`types.proto`)

*   **`VpcGraphNode`**: Represents a node in a VPC connectivity graph.
    *   Fields: `id`, `name`, `type` (e.g., "instance", "subnet"), `provider`, `account_id`, `region`, `vpc_id`, `properties` (map of string to string for additional details like CIDR, state, rules).
*   **`VpcGraphEdge`**: Represents an edge (relationship) in a VPC connectivity graph.
    *   Fields: `id`, `source_node_id`, `target_node_id`, `label` (describes the relationship, e.g., "LOCATED_IN", "ROUTES_TO", "ASSOCIATED_WITH"), `properties` (map of string to string).
*   **`InstanceGraphNode`**: Represents a node in an instance-centric connectivity graph.
    *   Fields: Similar to `VpcGraphNode`, tailored for instance context.
*   **`InstanceGraphEdge`**: Represents an edge in an instance-centric connectivity graph.
    *   Fields: Similar to `VpcGraphEdge`.

#### 6.3.4. Common and Utility Types

*   **`Credentials`**: Used to pass authentication details for providers.
    *   Fields: `aws_role`, `gcp_role`, `azure_role`, `aws_user_auth`, `gcp_user_auth`, `azure_user_auth`.
*   **`Error`**: Standard error message structure.
*   `google.protobuf.Timestamp`: Standard timestamp.
*   `google.protobuf.FieldMask`: Used for partial updates (not extensively shown in current examples but good practice for update RPCs).

This expanded documentation should provide a more comprehensive understanding of `awi-infra-guard`.