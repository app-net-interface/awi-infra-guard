package accesscontrol

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

// RemoveRulesFromSecurityGroup goes through all Network Security Rules
// attached to that particular Security Group and removes these entries
// which match given Rule Names.
//
// The function does not send updating request - it prepares SecurityGroup
// object making it ready to pass to the Azure Client.
//
// The function returns a boolean informing if any rule was actually
// removed. If not, it is pointless to update the Security Group.
func RemoveRulesFromSecurityGroup(
	nsg *armnetwork.SecurityGroup, rules RuleNames,
) (bool, error) {
	if nsg == nil {
		return false, errors.New(
			"cannot remove rules from nil Network Security Group",
		)
	}
	if nsg.Properties == nil {
		return false, fmt.Errorf(
			"cannot remove rules from nil Network Security Group %s due to nil properties",
			helper.StringPointerToString(nsg.ID),
		)
	}

	securityRulesWithoutVPCRules := make([]*armnetwork.SecurityRule, 0, len(nsg.Properties.SecurityRules))

	anyRuleRemoved := false

	for i := range nsg.Properties.SecurityRules {
		if nsg.Properties.SecurityRules[i] == nil {
			continue
		}
		if !securityRuleMatchesRuleGroup(nsg.Properties.SecurityRules[i], rules) {
			// We want to preserve only rules that are not matching our expected name.
			securityRulesWithoutVPCRules = append(securityRulesWithoutVPCRules, nsg.Properties.SecurityRules[i])
			continue
		}
		anyRuleRemoved = true
	}

	if anyRuleRemoved {
		nsg.Properties.SecurityRules = securityRulesWithoutVPCRules
	}

	return anyRuleRemoved, nil
}

func securityRuleMatchesRuleGroup(rule *armnetwork.SecurityRule, rules []ruleName) bool {
	if rule == nil || rule.Name == nil {
		return false
	}
	for _, r := range rules {
		if strings.HasPrefix(*rule.Name, string(r)) {
			return true
		}
	}
	return false
}
