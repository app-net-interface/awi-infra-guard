package accesscontrol

import (
	"fmt"
	"slices"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

type rule interface {
	Name() ruleName
	ToNetworkSecurityGroupRule(priority uint) (armnetwork.SecurityRule, error)
}

type access string

const (
	AccessAllow access = "ALLOW"
	AccessDeny  access = "DENY"
)

type VPCRule struct {
	namePrefix ruleName
	access     access
	sourceCIDR string
}

func (r VPCRule) Name() ruleName {
	return r.namePrefix
}

func (r VPCRule) ToNetworkSecurityGroupRule(priority uint) (armnetwork.SecurityRule, error) {
	name, err := nameWithPriority(r.namePrefix, priority)
	if err != nil {
		return armnetwork.SecurityRule{}, fmt.Errorf(
			"failed to generate a name for Security Rule %v: %w",
			r, err,
		)
	}
	return armnetwork.SecurityRule{
		Name: to.Ptr(name),
		Properties: newSecurityRuleProperties(
			r.access,
			r.sourceCIDR,
			"",
			priority,
			armnetwork.SecurityRuleProtocolAsterisk,
			[]string{},
		),
	}, nil
}

type DirectedVPCRule struct {
	namePrefix ruleName
	access     access
	sourceCIDR string
}

func (r DirectedVPCRule) Name() ruleName {
	return r.namePrefix
}

func (r DirectedVPCRule) ToNetworkSecurityGroupRule(priority uint) (armnetwork.SecurityRule, error) {
	name, err := nameWithPriority(r.namePrefix, priority)
	if err != nil {
		return armnetwork.SecurityRule{}, fmt.Errorf(
			"failed to generate a name for Security Rule %v: %w",
			r, err,
		)
	}
	return armnetwork.SecurityRule{
		Name: to.Ptr(name),
		Properties: newSecurityRuleProperties(
			r.access,
			r.sourceCIDR,
			"",
			priority,
			armnetwork.SecurityRuleProtocolAsterisk,
			[]string{},
		),
	}, nil
}

type CustomRule struct {
	namePrefix      ruleName
	access          access
	sourceCIDR      string
	destinationCIDR string
	protocol        armnetwork.SecurityRuleProtocol
	ports           []string
	// Subnets are not a part of actual Network Security
	// Group Rule but its a helper to provide a context
	// if the rules should be applied for a particular
	// NSG.
	subnets []string
}

func (r CustomRule) Name() ruleName {
	return r.namePrefix
}

func (r CustomRule) MatchesSubnet(subnet string) bool {
	return slices.Contains(r.subnets, subnet)
}

func (r CustomRule) ToNetworkSecurityGroupRule(priority uint) (armnetwork.SecurityRule, error) {
	name, err := nameWithPriority(r.namePrefix, priority)
	if err != nil {
		return armnetwork.SecurityRule{}, fmt.Errorf(
			"failed to generate a name for Security Rule %v: %w",
			r, err,
		)
	}
	return armnetwork.SecurityRule{
		Name: to.Ptr(name),
		Properties: newSecurityRuleProperties(
			r.access,
			r.sourceCIDR,
			r.destinationCIDR,
			priority,
			r.protocol,
			r.ports,
		),
	}, nil
}

func newSecurityRuleProperties(
	access access,
	sourceCIDR string,
	destinationCIDR string,
	priority uint,
	protocol armnetwork.SecurityRuleProtocol,
	ports []string,
) *armnetwork.SecurityRulePropertiesFormat {
	ruleAccess := armnetwork.SecurityRuleAccessAllow
	if access == AccessDeny {
		ruleAccess = armnetwork.SecurityRuleAccessDeny
	}

	var destCIDR string
	if destinationCIDR != "" {
		destCIDR = destinationCIDR
	}

	azurePriority := int32(priority)

	portRanges := make([]*string, len(ports))
	for i := range ports {
		portRanges[i] = &ports[i]
	}

	return &armnetwork.SecurityRulePropertiesFormat{
		Access:                   &ruleAccess,
		Priority:                 &azurePriority,
		SourceAddressPrefix:      &sourceCIDR,
		DestinationAddressPrefix: &destCIDR,
		Protocol:                 &protocol,
		DestinationPortRanges:    portRanges,
	}
}
