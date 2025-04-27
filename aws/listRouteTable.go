package aws

import (
	"context"
	"fmt"
	"strings" // Import strings
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
		// Initialize slices and maps for THIS route table INSIDE the loop
		routes := make([]types.Route, 0, len(rt.Routes))
		subnetIds := make([]string, 0, len(rt.Associations)) // For associated subnets
		igwIds := make([]string, 0)                          // For IGWs targeted by routes
		ngwIds := make([]string, 0)                          // For NAT Gateways targeted by routes
		vpgIds := make([]string, 0)                          // For VPGs targeted by routes
		tgwIds := make([]string, 0)                          // For TGWs targeted by routes

		// Use maps to track seen IDs within this specific route table's routes
		seenIgwIds := make(map[string]bool)
		seenNgwIds := make(map[string]bool)
		seenVpgIds := make(map[string]bool)
		seenTgwIds := make(map[string]bool)

		rtLink := fmt.Sprintf("https://%s.console.aws.amazon.com/vpcconsole/home?region=%s#RouteTableDetails:routeTableId=%s", region, region, aws.ToString(rt.RouteTableId))

		// Process Routes for this route table
		for _, r := range rt.Routes {
			var destination string
			if r.DestinationCidrBlock != nil {
				destination = convertString(r.DestinationCidrBlock)
			} else if r.DestinationIpv6CidrBlock != nil {
				destination = convertString(r.DestinationIpv6CidrBlock)
			} else if r.DestinationPrefixListId != nil {
				// Optionally handle Prefix List IDs if needed
				destination = convertString(r.DestinationPrefixListId)
			}

			target := ""
			// var targetType = "" // Keep track of target type if needed later

			// Determine the target and check for duplicates within this RT
			if r.CarrierGatewayId != nil {
				target = convertString(r.CarrierGatewayId)
				// targetType = "CarrierGateway"
			} else if r.CoreNetworkArn != nil {
				target = convertString(r.CoreNetworkArn)
				// targetType = "CoreNetwork"
			} else if r.EgressOnlyInternetGatewayId != nil {
				target = convertString(r.EgressOnlyInternetGatewayId)
				// targetType = "EgressOnlyInternetGateway" // Often considered a type of IGW
				// Decide if you want to add this to igwIds
				// if target != "" && !seenIgwIds[target] {
				// 	igwIds = append(igwIds, target)
				// 	seenIgwIds[target] = true
				// }
			} else if r.GatewayId != nil {
				target = convertString(r.GatewayId)
				if target != "" {
					// Check prefix for IGW or VPG
					if strings.HasPrefix(target, "igw-") {
						// targetType = "InternetGateway"
						if !seenIgwIds[target] {
							igwIds = append(igwIds, target)
							seenIgwIds[target] = true
						}
					} else if strings.HasPrefix(target, "vgw-") { // Note: AWS uses vgw- prefix for Virtual Private Gateway
						// targetType = "VirtualPrivateGateway"
						if !seenVpgIds[target] {
							vpgIds = append(vpgIds, target)
							seenVpgIds[target] = true
						}
					} else {
						// targetType = "Gateway" // Generic Gateway
					}
				}
			} else if r.InstanceId != nil {
				target = convertString(r.InstanceId)
				// targetType = "Instance"
			} else if r.LocalGatewayId != nil {
				target = convertString(r.LocalGatewayId)
				// targetType = "LocalGateway"
			} else if r.NatGatewayId != nil {
				target = convertString(r.NatGatewayId)
				// targetType = "NatGateway"
				if target != "" && !seenNgwIds[target] {
					ngwIds = append(ngwIds, target)
					seenNgwIds[target] = true
				}
			} else if r.NetworkInterfaceId != nil {
				target = convertString(r.NetworkInterfaceId)
				// targetType = "NetworkInterface"
			} else if r.TransitGatewayId != nil {
				target = convertString(r.TransitGatewayId)
				// targetType = "TransitGateway"
				if target != "" && !seenTgwIds[target] {
					tgwIds = append(tgwIds, target)
					seenTgwIds[target] = true
				}
			} else if r.VpcPeeringConnectionId != nil {
				target = convertString(r.VpcPeeringConnectionId)
				// targetType = "VpcPeeringConnection"
			} else if r.State == awsTypes.RouteStateBlackhole {
				target = "blackhole"
				// targetType = "Blackhole"
			}
			// Note: InstanceOwnerId is usually not the target itself

			routes = append(routes, types.Route{
				Destination: destination,
				Status:      string(r.State),
				Target:      target,
				// TargetType: targetType, // Optionally add TargetType to types.Route
			})
		}

		// Process Associations for this route table to get Subnet IDs
		seenSubnetIds := make(map[string]bool) // Track seen subnets for this RT
		for _, association := range rt.Associations {
			// We are primarily interested in subnet associations
			if association.SubnetId != nil {
				subnetId := convertString(association.SubnetId)
				if subnetId != "" && !seenSubnetIds[subnetId] {
					subnetIds = append(subnetIds, subnetId)
					seenSubnetIds[subnetId] = true
				}
			}
			// You could also capture associated Gateway IDs here if needed,
			// but they might differ from the route targets.
			// if association.GatewayId != nil { ... }
		}

		// Create the RouteTable object with IDs specific to this table
		routeTable := types.RouteTable{
			Name:      convertString(getTagName(rt.Tags)),
			ID:        convertString(rt.RouteTableId),
			Provider:  providerName,
			VpcID:     convertString(rt.VpcId),
			Region:    region,
			AccountID: convertString(rt.OwnerId), // Use convertString for safety
			Labels:    convertTags(rt.Tags),
			Routes:    routes,
			SubnetIds: subnetIds, // Use the subnetIds collected from associations
			IGWIds:    igwIds,    // Use the unique IGW IDs for this RT
			VGWIds:    vpgIds,    // Use the unique VPG IDs for this RT (Note: AWS uses vgw- prefix)
			NGWIds:    ngwIds,    // Use the unique NGW IDs for this RT
			TGWIds:    tgwIds,    // Use the unique TGW IDs for this RT
			SelfLink:  rtLink,
		}
		out = append(out, routeTable)
	}

	return out
}
