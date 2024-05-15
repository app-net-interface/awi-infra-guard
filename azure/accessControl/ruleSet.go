package accesscontrol

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/app-net-interface/awi-infra-guard/types"
)

// AccessControlRuleSet is a helper structure for
// preparing desired set of Network Security Group
// Rules that will be added to proper Network
// Security Groups.
type AccessControlRuleSet struct {
	VPCRules         []VPCRule
	DirectedVPCRules []DirectedVPCRule
	CustomRules      []CustomRule
}

func (a *AccessControlRuleSet) NewVPCRules(
	name ruleName,
	access access,
	sourceCIDRs []string,
) {
	for _, cidr := range sourceCIDRs {
		a.VPCRules = append(a.VPCRules, VPCRule{
			namePrefix: name,
			access:     access,
			sourceCIDR: cidr,
		})
	}
}

func (a *AccessControlRuleSet) NewDirectedVPCRules(
	name ruleName,
	access access,
	sourceCIDRs []string,
) {
	for _, cidr := range sourceCIDRs {
		a.DirectedVPCRules = append(a.DirectedVPCRules, DirectedVPCRule{
			namePrefix: name,
			access:     access,
			sourceCIDR: cidr,
		})
	}
}

func translateProtocolToAzureProtocol(protocol string) (armnetwork.SecurityRuleProtocol, error) {
	p := strings.ToLower(protocol)
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolAh)) {
		return armnetwork.SecurityRuleProtocolAh, nil
	}
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolEsp)) {
		return armnetwork.SecurityRuleProtocolEsp, nil
	}
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolIcmp)) {
		return armnetwork.SecurityRuleProtocolIcmp, nil
	}
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolTCP)) {
		return armnetwork.SecurityRuleProtocolTCP, nil
	}
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolUDP)) {
		return armnetwork.SecurityRuleProtocolUDP, nil
	}
	if p == strings.ToLower(string(armnetwork.SecurityRuleProtocolAsterisk)) || p == "" || p == "-1" {
		return armnetwork.SecurityRuleProtocolAsterisk, nil
	}
	return armnetwork.SecurityRuleProtocol(""), fmt.Errorf(
		"unsupported protocol '%s'. Azure supports following protocols: ["+
			"%s, %s, %s, %s, %s, %s]",
		protocol,
		string(armnetwork.SecurityRuleProtocolAh),
		string(armnetwork.SecurityRuleProtocolEsp),
		string(armnetwork.SecurityRuleProtocolIcmp),
		string(armnetwork.SecurityRuleProtocolTCP),
		string(armnetwork.SecurityRuleProtocolUDP),
		string(armnetwork.SecurityRuleProtocolAsterisk),
	)
}

func (a *AccessControlRuleSet) NewCustomRules(
	name ruleName,
	access access,
	subnets []string,
	sourceCIDRs []string,
	destinationCIDRs []string,
	protocolsAndPorts types.ProtocolsAndPorts,
) error {
	// TODO: Verify if Azure Security Group Rule can refer to multiple
	// source/destination CIDRs so that the entire slice below can be
	// merged into single rule.
	for _, sourceCIDR := range sourceCIDRs {
		for _, destinationCIDR := range destinationCIDRs {
			// TODO: Verify if these can be merged into single rule.
			for protocol, ports := range protocolsAndPorts {
				azProtocol, err := translateProtocolToAzureProtocol(protocol)

				if err != nil {
					return fmt.Errorf(
						"failed to translate given protocol %s: %w",
						protocol, err,
					)
				}

				a.CustomRules = append(a.CustomRules, CustomRule{
					namePrefix:      name,
					access:          access,
					sourceCIDR:      sourceCIDR,
					destinationCIDR: destinationCIDR,
					subnets:         subnets,
					protocol:        azProtocol,
					ports:           ports,
				})
			}
		}
	}

	return nil
}

func (a *AccessControlRuleSet) GenerateSecurityGroupRulesForVPC(
	prioritiesInUse helper.Set[uint],
) ([]armnetwork.SecurityRule, error) {
	return a.GenerateSecurityGroupRulesForSubnet(prioritiesInUse, nil)
}

func (a *AccessControlRuleSet) GenerateSecurityGroupRulesForSubnet(
	prioritiesInUse helper.Set[uint],
	subnet *string,
) ([]armnetwork.SecurityRule, error) {
	rules := a.collectAllRulesTogether(subnet)

	customRules := a.customRulesForSubnet(subnet)

	priorities, err := generatePriorities(
		prioritiesInUse,
		uint(len(a.VPCRules)),
		uint(len(a.DirectedVPCRules)),
		uint(len(customRules)),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to generate available priorities for rules: %w", err,
		)
	}

	securityGroupRules := make([]armnetwork.SecurityRule, 0, len(rules))

	for i := range rules {
		priority, err := priorities.Next()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to pick a priority for a rule: %w", err,
			)
		}
		rule, err := rules[i].ToNetworkSecurityGroupRule(priority)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to generate rule: %w", err,
			)
		}

		securityGroupRules = append(securityGroupRules, rule)
	}

	return securityGroupRules, nil
}

func (a *AccessControlRuleSet) RuleNamesForVPC() RuleNames {
	names := helper.Set[ruleName]{}
	rules := a.collectAllRulesTogether(nil)

	for _, rule := range rules {
		names.Set(rule.Name())
	}

	return names.Keys()
}

func (a *AccessControlRuleSet) RuleNamesForSubnet(subnet string) RuleNames {
	names := helper.Set[ruleName]{}
	rules := a.collectAllRulesTogether(&subnet)

	for _, rule := range rules {
		names.Set(rule.Name())
	}

	return names.Keys()
}

func (a *AccessControlRuleSet) customRulesForSubnet(subnet *string) []rule {
	if subnet == nil {
		return nil
	}

	rules := []rule{}

	for _, rule := range a.CustomRules {
		if rule.MatchesSubnet(*subnet) {
			rules = append(rules, rule)
		}
	}

	return rules
}

func (a *AccessControlRuleSet) collectAllRulesTogether(
	subnet *string,
) []rule {
	var customRulesForSubnet []rule
	if subnet != nil {
		customRulesForSubnet = a.customRulesForSubnet(subnet)
	}

	rules := make(
		[]rule, 0, len(a.VPCRules)+len(a.DirectedVPCRules)+len(customRulesForSubnet),
	)

	for _, r := range a.VPCRules {
		rules = append(rules, r)
	}
	for _, r := range a.DirectedVPCRules {
		rules = append(rules, r)
	}

	return append(rules, customRulesForSubnet...)
}
