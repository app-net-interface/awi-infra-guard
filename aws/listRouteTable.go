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

func (c *Client) ListRouteTables(ctx context.Context, params *infrapb.ListRouteTablesRequest) ([]types.RouteTable, error) {

	c.creds = params.Creds
	c.accountID = params.AccountId

	builder := newFilterBuilder()
	builder.withVPC(params.GetVpcId())

	filters := builder.build()

	if params.Region == "" || params.GetRegion() == "all" {
		var (
			wg             sync.WaitGroup
			allRouteTables []types.RouteTable
			allErrors      []error
			resultChannel  = make(chan regionResult)
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
				rts, err := c.getRouteTablesForRegion(ctx, regionName, filters)
				resultChannel <- regionResult{
					region: regionName,
					rts:    rts,
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
				allRouteTables = append(allRouteTables, result.rts...)
			}
		}
		c.logger.Infof("In account %s Found %d route tables across %d regions", c.accountID, len(allRouteTables), len(regions))

		if len(allErrors) > 0 {
			return allRouteTables, fmt.Errorf("errors occurred in some regions: %v", allErrors)
		}
		return allRouteTables, nil
	}
	return c.getRouteTablesForRegion(ctx, params.Region, filters)
}

func (c *Client) getRouteTablesForRegion(ctx context.Context, regionName string, filters []awsTypes.Filter) ([]types.RouteTable, error) {
	client, err := c.getEC2Client(ctx, c.accountID, regionName)
	if err != nil {
		return nil, err
	}
	// Call DescribeVpcs operation
	resp, err := client.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	return convertRouteTables(c.defaultRegion, regionName, resp.RouteTables), nil
}

func convertRouteTables(defaultRegion, region string, awsRts []awsTypes.RouteTable) []types.RouteTable {
	if region == "" {
		region = defaultRegion
	}

	out := make([]types.RouteTable, 0, len(awsRts))
	for _, rt := range awsRts {
		routes := make([]types.Route, 0, len(rt.Routes))
		rtLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#RouteTableDetails:routeTableId=%s", region, region, aws.ToString(rt.RouteTableId))

		for _, r := range rt.Routes {
			var destination string
			if r.DestinationCidrBlock != nil {
				destination = convertString(r.DestinationCidrBlock)
			} else if r.DestinationIpv6CidrBlock != nil {
				destination = convertString(r.DestinationIpv6CidrBlock)
			}

			target := ""
			if r.CarrierGatewayId != nil {
				target = convertString(r.CarrierGatewayId)
			} else if r.CoreNetworkArn != nil {
				target = convertString(r.CoreNetworkArn)
			} else if r.EgressOnlyInternetGatewayId != nil {
				target = convertString(r.EgressOnlyInternetGatewayId)
			} else if r.GatewayId != nil {
				target = convertString(r.GatewayId)
			} else if r.InstanceId != nil {
				target = convertString(r.InstanceId)
			} else if r.InstanceOwnerId != nil {
				target = convertString(r.InstanceOwnerId)
			} else if r.LocalGatewayId != nil {
				target = convertString(r.LocalGatewayId)
			} else if r.NatGatewayId != nil {
				target = convertString(r.NatGatewayId)
			} else if r.NetworkInterfaceId != nil {
				target = convertString(r.NetworkInterfaceId)
			} else if r.TransitGatewayId != nil {
				target = convertString(r.TransitGatewayId)
			} else if r.VpcPeeringConnectionId != nil {
				target = convertString(r.VpcPeeringConnectionId)
			}

			routes = append(routes, types.Route{
				Destination: destination,
				Status:      string(r.State),
				Target:      target,
			})
		}
		out = append(out, types.RouteTable{
			Name:      convertString(getTagName(rt.Tags)),
			ID:        convertString(rt.RouteTableId),
			Provider:  providerName,
			VpcID:     convertString(rt.VpcId),
			Region:    region,
			AccountID: *rt.OwnerId,
			Labels:    convertTags(rt.Tags),
			Routes:    routes,
			SelfLink:  rtLink,
		})
	}
	return out
}
