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
)

func (c *Client) ListACLs(ctx context.Context, params *infrapb.ListACLsRequest) ([]types.ACL, error) {
	if c.accountID != "" && params.AccountId != "" && c.accountID != params.AccountId {
		panic(fmt.Sprintf("ListACLs called with different AccountID: %s, expected: %s", params.AccountId, c.accountID))
	}

	operationalAccountID := params.GetAccountId()
	if operationalAccountID == "" && c.accountID != "" {
		c.logger.Infof("[AccountID: %s] ListACLs: params.AccountId is empty, using client's default/initial accountID: %s", c.accountID, c.accountID)
		operationalAccountID = c.accountID
	}
	if operationalAccountID == "" {
		c.logger.Errorf("ListACLs: operationalAccountID is empty and c.accountID is also empty. Cannot proceed.")
		return nil, fmt.Errorf("account ID is required but was not provided and no default is set")
	}
	c.logger.Debugf("[AccountID: %s] ListACLs called with VPC ID: %s, Region: %s", operationalAccountID, params.GetVpcId(), params.GetRegion())

	// REMOVED: c.creds = params.Creds
	// REMOVED: c.accountID = params.AccountId

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
		// Pass operationalAccountID to getAllRegions
		regions, err := c.getAllRegions(ctx, operationalAccountID)
		if err != nil {
			c.logger.Errorf("[AccountID: %s] ListACLs: Unable to describe regions, %v", operationalAccountID, err)
			return nil, err
		}
		c.logger.Debugf("[AccountID: %s] ListACLs: Iterating %d regions.", operationalAccountID, len(regions))
		for _, region := range regions { // region is awstypes.Region
			wg.Add(1)
			// Pass operationalAccountID to the goroutine
			go func(regionName string, accID string) {
				defer wg.Done()
				c.logger.Debugf("[AccountID: %s] ListACLs: Goroutine for region %s started.", accID, regionName)
				// Pass accID (operationalAccountID) to getACLsForRegion
				acls, err := c.getACLsForRegion(ctx, regionName, filters, accID)
				resultChannel <- regionResult{
					region: regionName,
					acls:   acls,
					err:    err,
				}
			}(*region.RegionName, operationalAccountID)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
			c.logger.Debugf("[AccountID: %s] ListACLs: All region goroutines finished, resultChannel closed.", operationalAccountID)
		}()

		for result := range resultChannel {
			if result.err != nil {
				c.logger.Infof("[AccountID: %s] ListACLs: Error in region %s: %v", operationalAccountID, result.region, result.err)
				allErrors = append(allErrors, fmt.Errorf("region %s: %v", result.region, result.err))
			} else {
				allACLs = append(allACLs, result.acls...)
			}
		}
		c.logger.Infof("[AccountID: %s] ListACLs: Found %d ACLs across %d regions", operationalAccountID, len(allACLs), len(regions))
		if len(allErrors) > 0 {
			return allACLs, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allACLs, nil
	}
	c.logger.Debugf("[AccountID: %s] ListACLs: Processing specific region: %s", operationalAccountID, params.Region)
	// Pass operationalAccountID to getACLsForRegion
	return c.getACLsForRegion(ctx, params.Region, filters, operationalAccountID)
}

// getACLsForRegion now accepts operationalAccountID
func (c *Client) getACLsForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter, operationalAccountID string) ([]types.ACL, error) {
	c.logger.Debugf("[AccountID: %s] getACLsForRegion: Region %s", operationalAccountID, regionName)
	// Use operationalAccountID for getting the EC2 client
	client, err := c.getEC2Client(ctx, operationalAccountID, regionName)
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getACLsForRegion: Failed to get EC2 client for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// Call DescribeNetworkAcls operation
	resp, err := client.DescribeNetworkAcls(ctx, &ec2.DescribeNetworkAclsInput{
		Filters: filters,
	})
	if err != nil {
		c.logger.Errorf("[AccountID: %s] getACLsForRegion: DescribeNetworkAcls failed for region %s: %v", operationalAccountID, regionName, err)
		return nil, err
	}
	// convertACLs itself uses OwnerId from the resource for AccountID in types.ACL.
	// No need to pass operationalAccountID to convertACLs if it's deriving from resource.
	return convertACLs(c.defaultRegion, regionName, resp.NetworkAcls), nil
}

func convertACLs(defaultRegion, region string, awsACLs []awsTypes.NetworkAcl) []types.ACL {
	if region == "" {
		region = defaultRegion
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
				if *r.Egress {
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
					rule.PortRange = fmt.Sprintf("%d", *r.PortRange.From) // Dereference pointer
				}
				if r.PortRange.To != nil {
					// Ensure "From" was present before appending "-"
					if r.PortRange.From != nil {
						rule.PortRange += fmt.Sprintf("-%d", *r.PortRange.To) // Dereference pointer
					} else {
						rule.PortRange = fmt.Sprintf("%d", *r.PortRange.To) // Dereference pointer
					}
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
			AccountID: aws.ToString(acl.OwnerId), // Uses OwnerId from the resource
			Labels:    convertTags(acl.Tags),
			Rules:     rules,
			SelfLink:  aclLink,
		})
	}
	return out
}

