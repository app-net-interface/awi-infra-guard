package gcp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
)

// TO DO: Implement the ListSecurityGroups method
func (c *Client) ListSecurityGroups(ctx context.Context, input *infrapb.ListSecurityGroupsRequest) ([]types.SecurityGroup, error) {
	secGroups := make([]types.SecurityGroup, 0)
	projectID := input.GetAccountId()

	var net network
	var err error
	if input.GetVpcId() != "" {
		net, err = c.vpcIdToSingleNetwork(ctx, input.GetAccountId(), input.GetVpcId())
		if err != nil {
			return nil, err
		}
	}
	networks, err := c.ListVPC(ctx, &infrapb.ListVPCRequest{AccountId: net.project})
	if err != nil {
		return nil, err
	}
	iter, err := c.computeService.Firewalls.List(projectID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	for _, fw := range iter.Items {
		//check if item has a target tag
		if len(fw.TargetTags) == 0 {
			continue
		} else {
			// Check if the target tag is present in an instance
			// Find all instances in the project match the target tags
			it := c.instancesClient.AggregatedList(ctx, &computepb.AggregatedListInstancesRequest{
				Project: input.AccountId,
			})

			for {
				pair, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return nil, err
				}
				//func getNetwork(networks []types.VPC, subnetworks []types.Subnet, networkInterfaces []*computepb.NetworkInterface) (privateIP, publicIP, vpcID, subnetID string) {
				// Check if the instance has the target tag

				for _, gcpInstance := range pair.Value.Instances {
					_, _, vpcId, _ := getNetwork(networks, nil, gcpInstance.NetworkInterfaces)
					if len(gcpInstance.Tags.Items) > 0 && ContainsAny(gcpInstance.Tags.Items, fw.TargetTags) {
						c.logger.Tracef("Instance %s has the target tag %v", *gcpInstance.Name, fw.TargetTags)
						secGroups = append(secGroups, convertSecurityGroup(projectID, vpcId, fw))
					}
				}
			}
		}
	}
	return secGroups, nil
}

func convertSecurityGroup(projectID string, vpcId string, firewall *compute.Firewall) types.SecurityGroup {

	consoleAccessLink := fmt.Sprintf("https://console.cloud.google.com/net-security/firewall-manager/firewall-policies/details/%s?project=%s", firewall.Name, projectID)

	rules := make([]types.SecurityGroupRule, 0, len(firewall.Allowed)+len(firewall.Denied))
	for _, rule := range firewall.Allowed {
		rules = append(rules, types.SecurityGroupRule{
			Protocol:  rule.IPProtocol,
			PortRange: strings.Join(rule.Ports, ","),
			Source:    firewall.SourceRanges,
			Direction: firewall.Direction,
		})
	}
	for _, rule := range firewall.Denied {
		rules = append(rules, types.SecurityGroupRule{
			Protocol:  rule.IPProtocol,
			PortRange: strings.Join(rule.Ports, ","),
			Source:    firewall.SourceRanges,
			Direction: firewall.Direction,
		})
	}

	return types.SecurityGroup{
		ID:        strconv.FormatUint(firewall.Id, 10),
		Name:      firewall.Name,
		Provider:  "gcp",
		AccountID: projectID,
		VpcID:     vpcId,
		Rules:     rules,
		SelfLink:  consoleAccessLink,
	}
}
