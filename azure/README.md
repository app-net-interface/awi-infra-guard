# VNet Peering Connectivity

The connectivity across two Azure VNets is achieved with use of
VNet Peering. To control the traffic, all subnets within both
VNets have either created or updated Network Security Group to
maintain traffic policies.

## Network Security Groups

Currently, there are 3 layers of Security Groups

1. Rules blocking connection between each other 
1. Rules allowing one side connection
1. Rules allowing custom traffic

Layers below override rules defined by those above.

### Blocking connection between each other

Creating VNet Peering specifies an option to allow the full traffic
between both Virtual Networks or disables it completely. In order to
be able to create any `allow` rules, that option needs to be enabled.

As a result, the VNet Peering initially allows the traffic between
subnets from both Virtual Networks so the traffic needs to be controlled
from the start with the use of Network Security Groups.

The approach is as follows: right before creating VNet peering, the
AWI Infra Guard goes through subnets from both VNets and creates/updates
a Network Security Group denying an access from Address Space belonging
to the other Virtual Network.

If we have VNet A and VNet B, these rules will be applied in the same
way regardless if VNet A is the source or VNet B is the source or if
we have bidirectional connection. As soon as peering exists, there rules
need to exist as well.

These Security Rules will use priorities from range `3600-4096` so there
is plenty of space left for rules that will override those.

### Allowing one side of connection

The AWI GRPC Catalyst SDWAN controller allows creating a VPC connection
with option to enable the traffic from source to destination. The second
layer of NSGs is for handling this scenario.

If the connection is created with this option set to true, then the
Network Security Groups in all subnets belonging to Destination VNet are
created/updated with rules allowing the traffic from Source Subnet Addresses.

As opposite to rules blocking connection between two Virtual Networks,
rules differ depending on if a certain VNet is a Source or a Destination as
these rules are applied only on the Destination VNet.

There is one more important difference between these rules and rules from layer above -
rules for blocking the connection are created for Address Spaces from VNets
and these rules are created to allow traffic for addresses from subnets.

To visualise it, let's say we have a Source VNet A
```yaml
AddressSpace: 10.0.0.0/16
Subnets:
- Name: Subnet A
  Address: 10.0.1.0/24
- Name: Subnet B
  Address: 10.0.2.0/24
- Name: Subnet C
  Address: 10.0.3.0/24
```

If we have created VPC Connection with VNet B as a destination, that allows
the traffic from A to B, the Virtual Network B would have following rules:

```yaml
Rules:
- prefix: 10.0.1.0/24
  action: allow
- prefix: 10.0.2.0/24
  action: allow
- prefix: 10.0.3.0/24
  action: allow
- prefix: 10.0.0.0/16
  action: block
```

Therefore, after creating such connection if the user creates a new subnet
within VNet A with address `10.0.4.0/24`, it won't be able to reach subnets
from VNet B unless VPC rules are refreshed. However, a new subnet created
within VNet B will be reachable from previous subnets in VNet A.

These Security Rules will use priorities from range `2800-3599` - since
a single rule in Network Security Group can be associated with only one
prefix, the subnet addressing approach will result in more rules than
VNet addressing (on the other hand, blocking rules from the first layer
are created for all peering connections and those allowing rules are
only for a subset of connections chosen by the administrator - and most
likely won't be a common choose due to security reasons).

Network Security Rules affected by these rules will be tagged with `ruleName`
parameter from AWI GRPC Catalyst SDWAN controller provided by
`AddInboundAllowRuleInVPC` function.

**Note:** The `AddInboundAllowRuleInVPC` function is called before actual
VPC Connection (AWI GRPC Catalyst SDWAN logic) and so if those rules are
applied, they will be created first - it is important that blocking rules
cannot accidentally override previous rules.

### Custom Traffic Rules

Those are rules reserved for `App Connections`.

Custom Traffic Rules will use priority range from `100-2799`.

## Network Security Rules Identification

In order to track what Network Security Rules are created by AWI Connector,
the NSGs have tags identifying their connections.

The tag keys identify the Network Domain Connection/VPC Policy/App Connection
Policy that required particular Security Groups and tag values contain the
name prefix for rules allowing to distinguish them from other rules.

That way, we can easily keep track of which rules needs to be removed/updated
whenever new policies are being applied.

## Limitations

VNet peering is fairly simple configuration and can be cheaper than other
solutions but it comes with a few issues which make it sufficient only for
a few connections at max.

### New subnets are not automatically associated with VPC Policies

As the Network Security Groups are being associated with subnets rather than
Virtual Networks, creating VPC Policies such as initial connection blocking
between two VNets means creating proper Network Security Rules for each
subnet within these Networks.

Currently, there is no process in the background that would constantly refresh
that policy, which means that any new created subnet will not be automatically
associated with these rules. As a result, when creating a connection between
two Virtual Networks, all subnets will have a rule attached that blocks the
connection to the subnets from other VNet. However, a new subnet won't have
such blocking rule applied and thus it will be able to reach subnets from other
Network which can be a serious security violation.

This can be solved by:

* implementing the logic in AWI GRPC Catalyst SDWAN that will constantly confirm
    that all subnets within VNets are associated with proper Network Security Groups
* utilization of Azure solutions such as `Azure Policies` to instantiate processes
    that will automatically attach proper NSGs to newly created subnets (solution
    similar to the one above, but uses already provided solutions by the Cloud)
* replacing VNet Peering with more complex structure such as HUB

### Running out of rule priorities

One Network Security Group allows assigning priorities from range `100:4096`.
Two rules within the same NSG cannot share the same priority so we can create
maximally 3996 rules. Additionally, a single NSG Rule can be associated with
only one address prefix so creating VPC Allow rules will require as many rules
as subnets are within the second VNet. Not to mention that Network Security
Group can have already multiple rules defined by other participants.

If there are no many VNet peers, the limit is not a very big issue, but it
may become a bottleneck for a Mesh.
