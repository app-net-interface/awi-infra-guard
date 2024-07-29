package aws

import (
	"context"
	"fmt"
	"sync"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *Client) ListACLs(ctx context.Context, params *infrapb.ListACLsRequest) ([]types.ACL, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())

	filters := builder.build()

	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg            sync.WaitGroup
			allACLs       []types.ACL
			allErrors     []error
			resultChannel = make(chan regionResult)
		)
		regions, err := c.getAllRegions(ctx)
		if err != nil {
			c.logger.Errorf("Unable to describe regions, %v", err)
			return nil, err
		}
		for _, region := range regions {
			wg.Add(1)
			go func(regionName string) {
				defer wg.Done()
				acls, err := c.getACLsForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					acls:   acls,
					err:    err,
				}
			}(*region.RegionName)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("Error in region %s: %v", result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allACLs = append(allACLs, result.acls...)
			}
		}

		if len(allErrors) > 0 {
			return allACLs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allACLs, nil
	}
	return c.getACLsForRegion(ctx, params.Region, filters)
}

func (c *Client) getACLsForRegion(ctx context.Context, regionName string, filters []awstypes.Filter) ([]types.ACL, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeVpcs operation
	resp, err := client.DescribeNetworkAcls(ctx, &ec2.DescribeNetworkAclsInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertACLs(c.defaultAccountID, c.defaultRegion, c.accountID, regionName, resp.NetworkAcls), nil
}

func convertACLs(defaultAccount, defaultRegion, account, region string, awsACLs []awsTypes.NetworkAcl) []types.ACL {
	if region == "" {
		region = defaultRegion
	}
	if account == "" {
		account = defaultAccount
	}
	out := make([]types.ACL, 0, len(awsACLs))
	for _, acl := range awsACLs {
		rules := make([]types.ACLRule, 0, len(acl.Entries))
		for _, r := range acl.Entries {

			rule := types.ACLRule{
				Number:            0,
				Protocol:          convertString(r.Protocol),
				PortRange:         "",
				SourceRanges:      nil,
				DestinationRanges: nil,
				Action:            string(r.RuleAction),
				Direction:         "",
			}
			if r.RuleNumber != nil {
				rule.Number = int(*r.RuleNumber)
			}
			if r.Egress != nil {
				if *r.Egress == true {
					rule.Direction = "Egress"
				} else {
					rule.Direction = "Ingress"
				}
			}
			if rule.Protocol == "-1" {
				rule.Protocol = "all"
			}
			if r.PortRange != nil {
				if r.PortRange.From != nil {
					rule.PortRange = fmt.Sprintf("%d", r.PortRange.From)
				}
				if r.PortRange.To != nil {
					rule.PortRange += fmt.Sprintf("- %d", r.PortRange.To)
				}
			}

			var cidrs []string
			if r.CidrBlock != nil {
				cidrs = append(cidrs, convertString(r.CidrBlock))
			}
			if r.Ipv6CidrBlock != nil {
				cidrs = append(cidrs, convertString(r.Ipv6CidrBlock))
			}
			if rule.Direction == "Egress" {
				rule.DestinationRanges = cidrs
			}
			if rule.Direction == "Ingress" {
				rule.SourceRanges = cidrs
			}

			rules = append(rules, rule)
		}
		aclLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#NetworkAclDetails:networkAclId=%s", region, region, aws.ToString(acl.NetworkAclId))
		out = append(out, types.ACL{
			Name:      convertString(getTagName(acl.Tags)),
			ID:        convertString(acl.NetworkAclId),
			Provider:  providerName,
			VpcID:     convertString(acl.VpcId),
			Region:    region,
			AccountID: *acl.OwnerId,
			Labels:    convertTags(acl.Tags),
			Rules:     rules,
			SelfLink:  aclLink,
		})
	}
	return out
}
