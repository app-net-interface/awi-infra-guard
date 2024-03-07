package azure

import (
	"context"

	"github.com/app-net-interface/awi-infra-guard/types"
)

// AccessControl interface implementation
func (c *Client) AddInboundAllowRuleInVPC(ctx context.Context, account, region string, destinationVpcID string, cidrsToAllow []string, ruleName string,
	tags map[string]string) error {
	// TBD
	return nil
}

func (c *Client) AddInboundAllowRuleByLabelsMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, labels map[string]string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	// TBD
	return "", nil, nil
}

func (c *Client) AddInboundAllowRuleBySubnetMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, subnetCidrs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, subnets []types.Subnet, err error) {
	// TBD
	return "", nil, nil, nil
}

func (c *Client) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, account, region string,
	vpcID string, ruleName string, instancesIPs []string, cidrsToAllow []string,
	protocolsAndPorts types.ProtocolsAndPorts) (ruleId string, instances []types.Instance, err error) {
	// TBD
	return "", nil, nil
}

func (c *Client) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, account, region string, loadBalancerDNS string, vpcID string,
	ruleName string, cidrsToAllow []string, protocolsAndPorts types.ProtocolsAndPorts) (loadBalancerId, ruleId string, err error) {
	// TBD
	return "", "", nil
}

func (c *Client) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, account, region string, vpcID string, ruleName string) error {
	// TBD
	return nil
}

func (c *Client) RemoveInboundAllowRulesFromVPCById(ctx context.Context, account, region string, vpcID string, instanceIDs []string,
	loadBalancersIDs []string, ruleId string) error {
	// TBD
	return nil
}

func (c *Client) RemoveInboundAllowRuleRulesByTags(ctx context.Context, account, region string, vpcID string, ruleName string, tags map[string]string) error {
	// TBD
	return nil
}

func (c *Client) RefreshInboundAllowRule(ctx context.Context, account, region string, ruleId string, cidrsToAdd []string, cidrsToRemove []string,
	destinationLabels map[string]string, destinationPrefixes []string, destinationVPCId string,
	protocolsAndPorts types.ProtocolsAndPorts) (instances []types.Instance, subnets []types.Subnet, err error) {
	// TBD
	return nil, nil, nil
}