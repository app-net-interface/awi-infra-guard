// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package connector

import (
	"context"
	"errors"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/cidrpool"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

type createConnectionBGPFullConfig struct {
	Gateway    types.Gateway
	Provider   provider.Provider
	Interfaces []string
	Secrets    []string
	ASN        uint64
}

func (c *Connector) createConnectionWithBGPAuthoritarian(
	ctx context.Context,
	firstConfig createConnectionBGPFullConfig,
	secondConfig createConnectionBGPFullConfig,
	numberOfTunnels uint8,
	secrets []string,
	bgpCIDRPools []*cidrpool.CIDRV4Pool,
) error {
	resultFromFirstProvider, err := firstConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		firstConfig.Gateway,
		secondConfig.Gateway,
		types.AttachModeGenerateBothIPs,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: secondConfig.Interfaces,
			ASN:               firstConfig.ASN,
			PeerASN:           secondConfig.ASN,
			NumberOfTunnels:   numberOfTunnels,
			SharedSecrets:     secrets,
			BGPCIDRPools:      bgpCIDRPools,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			firstConfig.Gateway.Name, types.AttachModeGenerateBothIPs, err,
		)
	}

	if len(firstConfig.Interfaces) == 0 {
		c.logger.Info("Late Interfaces. Trying to obtain from previous operation")
		if len(resultFromFirstProvider.Interfaces) == 0 {
			return fmt.Errorf(
				"cannot create attachment for gateway %s as no interfaces were returned "+
					"from first gateway %s", secondConfig.Gateway.Name, firstConfig.Gateway.Name,
			)
		}
		firstConfig.Interfaces = resultFromFirstProvider.Interfaces
	}

	_, err = secondConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		secondConfig.Gateway,
		firstConfig.Gateway,
		types.AttachModeAcceptOtherIP,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: firstConfig.Interfaces,
			ASN:               secondConfig.ASN,
			PeerASN:           firstConfig.ASN,
			NumberOfTunnels:   numberOfTunnels,
			BGPAddresses:      resultFromFirstProvider.PeerBGPAddresses,
			PeerBGPAddresses:  resultFromFirstProvider.BGPAddresses,
			SharedSecrets:     resultFromFirstProvider.SharedSecrets,
			BGPCIDRPools:      bgpCIDRPools,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			secondConfig.Gateway.Name, types.AttachModeAcceptOtherIP, err,
		)
	}
	return nil
}

func (c *Connector) createConnectionWithBGPCooperative(
	ctx context.Context,
	firstConfig createConnectionBGPFullConfig,
	secondConfig createConnectionBGPFullConfig,
	numberOfTunnels uint8,
	secrets []string,
	bgpCIDRPools []*cidrpool.CIDRV4Pool,
) error {
	result, err := firstConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		firstConfig.Gateway,
		secondConfig.Gateway,
		types.AttachModeGenerateIP,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: secondConfig.Interfaces,
			ASN:               firstConfig.ASN,
			PeerASN:           secondConfig.ASN,
			NumberOfTunnels:   numberOfTunnels,
			SharedSecrets:     secrets,
			BGPCIDRPools:      bgpCIDRPools,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			firstConfig.Gateway.Name, types.AttachModeGenerateIP, err,
		)
	}

	// TODO: Handle verifying interfaces - should they be picked from
	// the config or from the first result of attachment.
	result, err = secondConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		secondConfig.Gateway,
		firstConfig.Gateway,
		types.AttachModeGenerateIPAndAcceptOtherIP,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: firstConfig.Interfaces,
			ASN:               secondConfig.ASN,
			PeerASN:           firstConfig.ASN,
			NumberOfTunnels:   numberOfTunnels,
			SharedSecrets:     result.SharedSecrets,
			BGPCIDRPools:      bgpCIDRPools,
			PeerBGPAddresses:  result.BGPAddresses,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			secondConfig.Gateway.Name, types.AttachModeGenerateIPAndAcceptOtherIP, err,
		)
	}

	// TODO: Handle verifying interfaces - should they be picked from
	// the config or from the first result of attachment.
	_, err = firstConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		firstConfig.Gateway,
		secondConfig.Gateway,
		types.AttachModeAcceptOtherIP,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: secondConfig.Interfaces,
			ASN:               firstConfig.ASN,
			PeerASN:           secondConfig.ASN,
			NumberOfTunnels:   numberOfTunnels,
			SharedSecrets:     secrets,
			BGPCIDRPools:      bgpCIDRPools,
			PeerBGPAddresses:  result.BGPAddresses,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			firstConfig.Gateway.Name, types.AttachModeAcceptOtherIP, err,
		)
	}

	return nil
}

func (c *Connector) prepareCIDRForBGPAddressing(
	ctx context.Context,
	gateways gateways,
) (cidrpool.CIDRV4Pools, error) {
	sourceSettings, err := gateways.SourceProvider.GetGatewayConnectionSettings(
		ctx, gateways.SourceGateway)
	if err != nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"failed to load Source Gateway Settings for Gateway %s: %w",
			gateways.SourceGateway.Name, err,
		)
	}
	destSettings, err := gateways.DestinationProvider.GetGatewayConnectionSettings(
		ctx, gateways.DestinationGateway,
	)
	if err != nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"failed to load Destination Gateway Settings for Gateway %s: %w",
			gateways.DestinationGateway.Name, err,
		)
	}

	if sourceSettings.BGPSetting == nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"source gateway %s doesn't support BGP",
			gateways.SourceGateway.Name,
		)
	}
	if destSettings.BGPSetting == nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"destination gateway %s doesn't support BGP",
			gateways.DestinationGateway.Name,
		)
	}

	intersectingCIDRs, err := cidrpool.IntersectingCIDRs(
		sourceSettings.BGPSetting.AllowedIPRanges,
		destSettings.BGPSetting.AllowedIPRanges,
	)
	if err != nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"failed to process allowed CIDRs: %w", err,
		)
	}

	pool, err := cidrpool.NewCIDRV4Pools(intersectingCIDRs)
	if err != nil {
		return cidrpool.CIDRV4Pools{}, fmt.Errorf(
			"failed to create CIDR pool: %w", err,
		)
	}

	for _, forbiddenCIDR := range sourceSettings.BGPSetting.ExcludedIPRanges {
		if err = pool.ExcludeCIDRFromPools(forbiddenCIDR); err != nil {
			return cidrpool.CIDRV4Pools{}, fmt.Errorf(
				"failed to exclude CIDR %s from pool: %w", forbiddenCIDR, err,
			)
		}
	}
	for _, forbiddenCIDR := range destSettings.BGPSetting.ExcludedIPRanges {
		if err = pool.ExcludeCIDRFromPools(forbiddenCIDR); err != nil {
			return cidrpool.CIDRV4Pools{}, fmt.Errorf(
				"failed to exclude CIDR %s from pool: %w", forbiddenCIDR, err,
			)
		}
	}

	if pool.Full() {
		return cidrpool.CIDRV4Pools{}, errors.New(
			"failed to prepare a CIDR pool for BGP Addressing. " +
				"There are no overlapping addresses that can be used by both gateways",
		)
	}
	return pool, nil
}

func (c *Connector) createConnectionWithBGP(
	ctx context.Context,
	gateways gateways,
	sourceInterfaces,
	destinationInterfaces []string,
	numberOfTunnels uint8,
	secrets []string,
) error {
	sourceASN, err := gateways.SourceProvider.InitializeASN(
		ctx, gateways.SourceGateway, gateways.DestinationGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s ASN: %w",
			gateways.SourceGateway.Name, err)
	}
	destASN, err := gateways.DestinationProvider.InitializeASN(
		ctx, gateways.DestinationGateway, gateways.SourceGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s ASN: %w",
			gateways.DestinationGateway.Name, err)
	}

	bgpScenario, err := c.getBGPScenario(ctx, gateways)
	if err != nil {
		return fmt.Errorf(
			"failed to obtain a scenario for creating BGP Session: %w",
			err,
		)
	}

	configs := []createConnectionBGPFullConfig{
		{
			Gateway:    gateways.SourceGateway,
			Provider:   gateways.SourceProvider,
			Interfaces: sourceInterfaces,
			ASN:        sourceASN,
		},
		{
			Gateway:    gateways.DestinationGateway,
			Provider:   gateways.DestinationProvider,
			Interfaces: destinationInterfaces,
			ASN:        destASN,
		},
	}
	if !bgpScenario.SourceSideStarts {
		configs[0], configs[1] = configs[1], configs[0]
	}

	pool, err := c.prepareCIDRForBGPAddressing(ctx, gateways)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare CIDR Pools for BGP Addressing: %w", err,
		)
	}

	pools := make([]*cidrpool.CIDRV4Pool, 0, numberOfTunnels)
	for i := uint8(0); i < numberOfTunnels; i++ {
		cidr, err := pool.Get(30)
		if err != nil {
			return fmt.Errorf(
				"failed to obtain CIDR with mask 30: %w", err,
			)
		}
		if cidr == nil {
			return fmt.Errorf(
				"failed to obtain CIDR with mask 30. No more available pools with such mask",
			)
		}
		newPool, err := cidrpool.NewCIDRV4Pool(cidr.String())
		if err != nil {
			return fmt.Errorf(
				"failed to create a pool from CIDR %s: %w", cidr.String(), err,
			)
		}
		if newPool == nil {
			return fmt.Errorf(
				"failed to create a pool from CIDR %s. Got nil", cidr.String(),
			)
		}
		pools = append(pools, newPool)
	}

	if bgpScenario.StarterGeneratesBothAddresses {
		return c.createConnectionWithBGPAuthoritarian(
			ctx,
			configs[0],
			configs[1],
			numberOfTunnels,
			secrets,
			pools,
		)
	}
	return c.createConnectionWithBGPCooperative(
		ctx,
		configs[0],
		configs[1],
		numberOfTunnels,
		secrets,
		pools,
	)
}

func (c *Connector) getBGPScenario(
	ctx context.Context,
	gateways gateways,
) (types.BGPScenario, error) {
	sourceSetting, err := gateways.SourceProvider.GetGatewayConnectionSettings(
		ctx, gateways.SourceGateway)
	if err != nil {
		return types.BGPScenario{}, fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			gateways.SourceGateway.Name, gateways.SourceProvider.Name(), err,
		)
	}
	if sourceSetting.BGPSetting == nil {
		return types.BGPScenario{}, fmt.Errorf(
			"failed to get BGP settings for gateway %s from provider %s. "+
				"Options don't specify any BGP setting: %v",
			gateways.SourceGateway.Name, gateways.SourceProvider.Name(),
			sourceSetting,
		)
	}
	destinationSetting, err := gateways.DestinationProvider.GetGatewayConnectionSettings(
		ctx, gateways.DestinationGateway)
	if err != nil {
		return types.BGPScenario{}, fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			gateways.DestinationGateway.Name, gateways.DestinationProvider.Name(), err,
		)
	}
	if destinationSetting.BGPSetting == nil {
		return types.BGPScenario{}, fmt.Errorf(
			"failed to get BGP settings for gateway %s from provider %s. "+
				"Options don't specify any BGP setting: %v",
			gateways.DestinationGateway.Name, gateways.DestinationProvider.Name(),
			destinationSetting,
		)
	}
	scenario := types.BGPScenarioFromBothConfigs(
		&sourceSetting.BGPSetting.Addressing, &destinationSetting.BGPSetting.Addressing)
	if scenario == nil {
		return types.BGPScenario{}, fmt.Errorf(
			"failed to obtain BGP Scenario between gateways. It seems their "+
				"BGP settings are incompatible. Source %s:%s has setting: %v "+
				"and Destination %s:%s has setting: %v",
			gateways.SourceProvider.Name(),
			gateways.SourceGateway.Name,
			sourceSetting.BGPSetting,
			gateways.DestinationProvider.Name(),
			gateways.DestinationGateway.Name,
			destinationSetting.BGPSetting,
		)
	}
	return *scenario, nil
}
