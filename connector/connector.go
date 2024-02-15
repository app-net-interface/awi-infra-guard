// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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

	"github.com/app-net-interface/awi-infra-guard/connector/db"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Connector struct {
	providerManager *ProviderManager
	db              *db.Client
	config          *Config
	logger          *logrus.Entry
	mainLogger      *logrus.Logger
	// Variables obtained during Connection Creation
	connectionID     string
	sourceCIDRs      []string
	destinationCIDRs []string
}

func NewConnector(ctx context.Context, logger *logrus.Logger, config *Config) (*Connector, error) {
	providerManager, err := NewProviderManager(logger.WithField("logger", "providerManager"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Provider Manager: %w", err)
	}
	dbClient, err := db.NewClient(db.DefaultDBFile, logger.WithField("logger", "db"))
	if err != nil {
		return nil, fmt.Errorf("Could not create DB Client: %w", err)
	}
	return &Connector{
		providerManager: &providerManager,
		db:              dbClient,
		config:          config,
		logger:          logger.WithField("logger", "main"),
		mainLogger:      logger,
	}, nil
}

func (c *Connector) Close() error {
	if err := c.providerManager.Close(); err != nil {
		return fmt.Errorf("an error occured while closing Provider Manager: %w", err)
	}
	c.providerManager = nil
	if err := c.db.Close(); err != nil {
		return fmt.Errorf("an error occured while closing DB client: %w", err)
	}
	c.db = nil
	c.logger.Debug("Closed the Connector successfully")
	return nil
}

func calculateNumberOfTunnels(source, destination types.GatewayConnectionSettings) (uint8, error) {
	if source.NumberOfInterfaces == 0 || destination.NumberOfInterfaces == 0 {
		return 0, fmt.Errorf(
			"cannot establish a connection when one side defines 0 interfaces. "+
				"Source interfaces: %d. Destination interfaces: %d",
			source.NumberOfInterfaces, destination.NumberOfInterfaces,
		)
	}
	if destination.NumberOfInterfaces > source.NumberOfInterfaces {
		if destination.NumberOfInterfaces%source.NumberOfInterfaces != 0 {
			return 0, fmt.Errorf(
				"the greater number of interfaces must be multiple of lesser value."+
					"Source interfaces: %d. Destination interfaces: %d",
				source.NumberOfInterfaces, destination.NumberOfInterfaces,
			)
		}
		return destination.NumberOfInterfaces, nil
	}
	if source.NumberOfInterfaces%destination.NumberOfInterfaces != 0 {
		return 0, fmt.Errorf(
			"the greater number of interfaces must be multiple of lesser value."+
				"Source interfaces: %d. Destination interfaces: %d",
			source.NumberOfInterfaces, destination.NumberOfInterfaces,
		)
	}
	return source.NumberOfInterfaces, nil
}

type createConnectionBGPFullConfig struct {
	Gateway    types.Gateway
	Provider   provider.Provider
	Setting    types.GatewayConnectionSettings
	Interfaces []string
	ASN        uint64
}

func (c *Connector) createConnectionWithBGPMasterMind(
	ctx context.Context,
	sourceConfig createConnectionBGPFullConfig,
	destConfig createConnectionBGPFullConfig,
	NumberOfTunnels uint8,
) error {
	var firstConfig, secondConfig createConnectionBGPFullConfig

	if sourceConfig.Setting.BGPSetting.PickOtherIPAddress {
		firstConfig = sourceConfig
		secondConfig = destConfig
	} else {
		firstConfig = destConfig
		secondConfig = sourceConfig
	}
	resultFromFirstProvider, err := firstConfig.Provider.AttachToExternalGatewayWithBGP(
		ctx,
		firstConfig.Gateway,
		secondConfig.Gateway,
		types.AttachModeGenerateBothIPs,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: secondConfig.Interfaces,
			ASN:               firstConfig.ASN,
			PeerASN:           secondConfig.ASN,
			NumberOfTunnels:   NumberOfTunnels,
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
		types.AttachModeAttachOtherIP,
		types.CreateBGPConnectionConfig{
			OutsideInterfaces: firstConfig.Interfaces,
			ASN:               secondConfig.ASN,
			PeerASN:           firstConfig.ASN,
			NumberOfTunnels:   NumberOfTunnels,
			BGPAddresses:      resultFromFirstProvider.PeerBGPAddresses,
			PeerBGPAddresses:  resultFromFirstProvider.BGPAddresses,
			SharedSecrets:     resultFromFirstProvider.SharedSecrets,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create attachment resources for gateway %s with mode %d: %w",
			secondConfig.Gateway.Name, types.AttachModeAttachOtherIP, err,
		)
	}
	return nil
}

func (c *Connector) createConnectionWithBGPRegular(
	ctx context.Context,
	sourceConfig createConnectionBGPFullConfig,
	destConfig createConnectionBGPFullConfig,
	NumberOfTunnels uint8,
) error {
	// TODO: Implement the logic.
	return errors.New("createConnectionWithBGPRegular NOT IMPLEMENTED")
}

func (c *Connector) createConnectionWithBGP(
	ctx context.Context,
	sourceGateway,
	destinationGateway types.Gateway,
	sourceProvider,
	destinationProvider provider.Provider,
	sourceSetting,
	destinationSetting types.GatewayConnectionSettings,
	numberOfTunnels uint8,
) error {
	sourceInterfaces, err := sourceProvider.InitializeGatewayInterfaces(ctx, sourceGateway, destinationGateway)
	if err != nil {
		return fmt.Errorf("failed to initialize Gateway %s interfaces: %w", sourceGateway.Name, err)
	}
	destInterfaces, err := destinationProvider.InitializeGatewayInterfaces(ctx, destinationGateway, sourceGateway)
	if err != nil {
		return fmt.Errorf("failed to initialize Gateway %s interfaces: %w", destinationGateway.Name, err)
	}

	sourceASN, err := sourceProvider.InitializeASN(ctx, sourceGateway, destinationGateway)
	if err != nil {
		return fmt.Errorf("failed to initialize Gateway %s ASN: %w", sourceGateway.Name, err)
	}
	destASN, err := destinationProvider.InitializeASN(ctx, destinationGateway, sourceGateway)
	if err != nil {
		return fmt.Errorf("failed to initialize Gateway %s ASN: %w", destinationGateway.Name, err)
	}

	sourceConfig := createConnectionBGPFullConfig{
		Gateway:    sourceGateway,
		Provider:   sourceProvider,
		Setting:    sourceSetting,
		Interfaces: sourceInterfaces,
		ASN:        sourceASN,
	}
	destConfig := createConnectionBGPFullConfig{
		Gateway:    destinationGateway,
		Provider:   destinationProvider,
		Setting:    destinationSetting,
		Interfaces: destInterfaces,
		ASN:        destASN,
	}

	if (sourceSetting.BGPSetting.PickOwnIPAddress && sourceSetting.BGPSetting.PickOtherIPAddress) ||
		(destinationSetting.BGPSetting.PickOwnIPAddress && destinationSetting.BGPSetting.PickOtherIPAddress) {
		return c.createConnectionWithBGPMasterMind(ctx, sourceConfig, destConfig, numberOfTunnels)
	}
	return c.createConnectionWithBGPRegular(ctx, sourceConfig, destConfig, numberOfTunnels)
}

func (c *Connector) createConnectionWithStaticRouting(
	ctx context.Context,
	sourceGateway,
	destinationGateway types.Gateway,
	sourceProvider,
	destinationProvider provider.Provider,
	sourceSetting,
	destinationSetting types.GatewayConnectionSettings,
	numberOfTunnels uint8,
) error {
	return errors.New("NOT IMPLEMENTED")
}

func (c *Connector) createConnection(
	ctx context.Context,
	sourceGateway,
	destinationGateway types.Gateway,
	sourceProvider,
	destinationProvider provider.Provider,
) error {
	sourceSetting, err := sourceProvider.GetGatewayConnectionSettings(ctx, sourceGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			sourceGateway.Name, sourceProvider.Name(), err,
		)
	}
	destinationSetting, err := destinationProvider.GetGatewayConnectionSettings(ctx, destinationGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			destinationGateway.Name, destinationProvider.Name(), err,
		)
	}
	numberOfTunnels, err := calculateNumberOfTunnels(sourceSetting, destinationSetting)
	if err != nil {
		return fmt.Errorf("failed to calculate number of desired tunnels: %w", err)
	}
	if sourceSetting.BGPSetting != nil && destinationSetting.BGPSetting != nil {
		return c.createConnectionWithBGP(
			ctx,
			sourceGateway,
			destinationGateway,
			sourceProvider,
			destinationProvider,
			sourceSetting,
			destinationSetting,
			numberOfTunnels,
		)
	}
	if sourceSetting.StaticRoutingSetting != nil && destinationSetting.StaticRoutingSetting != nil {
		return c.createConnectionWithStaticRouting(
			ctx,
			sourceGateway,
			destinationGateway,
			sourceProvider,
			destinationProvider,
			sourceSetting,
			destinationSetting,
			numberOfTunnels,
		)
	}
	return fmt.Errorf(
		"cannot create a connection between providers since they cannot match routing setting. "+
			"Source gateway: %s, source provider: %s, destination gateway: %s, destination provider: %s",
		sourceGateway.Name, sourceProvider.Name(), destinationGateway.Name, destinationProvider.Name(),
	)
}

func (c *Connector) Connect(ctx context.Context, request types.Request) error {
	connection, err := c.getConnectionEntryIfExists(ctx, request)
	if err != nil {
		return fmt.Errorf(
			"got error while trying to find a Connection entry in the DB for configuration: %v",
			c.config)
	}
	if connection == nil {
		c.logger.Debugf("connection does not exist yet")
	} else if connection.State == db.StateActive {
		c.logger.Infof("connection already exists and has active state. Nothing to do")
		// TODO: Consider returning an error here.
		// Right now attempt to create a connection, which already exists, is
		// not considered an error but it may be a wrong assumption.
		return nil
	} else {
		c.logger.Infof(
			"found connection with state: %s. Trying to reestablish the connection",
			connection.State)
		c.connectionID = connection.ID
	}

	c.logger.Debugf("Starting to create a Connection: %v", request)

	sourceProvider, destinationProvider, err := c.initializeProviders(ctx, request)
	if err != nil {
		return fmt.Errorf("cannot create connection due to provider failure: %w", err)
	}
	sourceGW, destGW, err := c.getGateways(ctx, request, sourceProvider, destinationProvider)
	if err != nil {
		return fmt.Errorf("validation of requested gateways to be connected failed: %w", err)
	}

	if err := c.createConnectionEntryInDB(ctx, sourceGW, destGW, sourceProvider, destinationProvider); err != nil {
		return fmt.Errorf("connection could not be stored in the DB: %w", err)
	}

	err = c.createConnection(ctx, sourceGW, destGW, sourceProvider, destinationProvider)
	if err != nil {
		return fmt.Errorf(
			"failed to create a connection between %s:%s and %s:%s: %w",
			sourceProvider.Name(), sourceGW.Name, destinationProvider.Name(), destGW.Name,
			err)
	}

	if err = c.updateConnectionEntryInDBWithCIDRs(ctx, sourceGW, destGW, sourceProvider, destinationProvider); err != nil {
		c.logger.Warn("Resources were create successfully but failed to update DB Connection Entry with CIDRs")
		return fmt.Errorf("failed to update connection DB Entry with CIDRs: %w", err)
	}

	if err = c.updateConnectionCreationEntryInDB(ctx, db.StateActive, "", sourceGW, destGW, sourceProvider, destinationProvider); err != nil {
		c.logger.Warn("Resources were create successfully but failed to update DB Connection Entry with Active state")
		return fmt.Errorf("failed to update connection DB Entry with Active state: %w", err)
	}

	c.logger.Info("Connection created successfully")
	return nil
}

func (c *Connector) getConnectionEntryIfExists(
	ctx context.Context, request types.Request,
) (*db.Connection, error) {
	if request.ConnectionID != nil {
		return c.db.GetConnection(*request.ConnectionID)
	}
	return c.db.GetConnectionWithGateways(
		request.ConnectionDetails.Source.GatewayID,
		request.ConnectionDetails.Source.Provider,
		request.ConnectionDetails.Destination.GatewayID,
		request.ConnectionDetails.Destination.Provider,
	)
}

func (c *Connector) initializeProviders(
	ctx context.Context, request types.Request,
) (provider.Provider, provider.Provider, error) {
	sourceConfig := ""
	sourceProviderName := request.ConnectionDetails.Source.Provider
	if config, ok := c.config.Providers[sourceProviderName]; ok {
		sourceConfig = config.(string)
	} else {
		c.logger.Infof("No configuration found for Source Provider '%s'. Using defaults", sourceProviderName)
	}
	sourceProvider, err := c.providerManager.InitializeProvider(
		ctx, c.logger, sourceProviderName, sourceConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize source provider '%s': %w", sourceProviderName, err)
	}
	destConfig := ""
	destProviderName := request.ConnectionDetails.Destination.Provider
	if config, ok := c.config.Providers[destProviderName]; ok {
		destConfig = config.(string)
	} else {
		c.logger.Infof("No configuration found for Destination Provider '%s'. Using defaults", destProviderName)
	}
	destProvider, err := c.providerManager.InitializeProvider(
		ctx, c.logger, destProviderName, destConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize destination provider '%s': %w", destProviderName, err)
	}
	return sourceProvider, destProvider, nil
}

func (c *Connector) SetConnectionID(id string) {
	c.connectionID = id
}

func (c *Connector) updateConnectionEntryInDBWithCIDRs(
	ctx context.Context,
	sourceGateway, destGateway types.Gateway,
	sourceProvider, destProvider provider.Provider,
) error {
	sourceCIDRs, err := sourceProvider.GetCIDRs(ctx, sourceGateway)
	if err != nil {
		return fmt.Errorf("failed to get Connection CIDRs for source Provider: %w", err)
	}
	destinationCIDRs, err := destProvider.GetCIDRs(ctx, destGateway)
	if err != nil {
		return fmt.Errorf("failed to get Connection CIDRs for Destination Provider: %w", err)
	}

	c.sourceCIDRs = sourceCIDRs
	c.destinationCIDRs = destinationCIDRs

	if err = c.updateConnectionCreationEntryInDB(
		ctx,
		db.StateCreationInProgress,
		"",
		sourceGateway,
		destGateway,
		sourceProvider,
		destProvider); err != nil {
		return fmt.Errorf(
			"failed to perform an update for connection in DB with connection CIDRs: %w",
			err)
	}

	return nil
}

func (c *Connector) createConnectionEntryInDB(
	ctx context.Context,
	sourceGateway, destGateway types.Gateway,
	sourceProvider, destProvider provider.Provider,
) error {
	if c.connectionID != "" {
		c.logger.Debugf("connection already exists. Skipping creating new entry in DB")
	} else {
		c.logger.Debugf("Creating new entry for the connection in DB")
		c.connectionID = string(uuid.New().String())
	}
	err := c.updateConnectionCreationEntryInDB(
		ctx,
		db.StateCreationInProgress,
		"",
		sourceGateway,
		destGateway,
		sourceProvider,
		destProvider,
	)
	if err != nil {
		return fmt.Errorf("failed to create entry about the connection in the DB: %w", err)
	}
	return nil
}

func (c *Connector) updateConnectionCreationEntryInDB(
	ctx context.Context,
	state db.ConnectionState,
	creationError string,
	sourceGateway, destGateway types.Gateway,
	sourceProvider, destProvider provider.Provider) error {
	if c.connectionID == "" {
		return errors.New("trying to update connection when connection ID is not set yet")
	}
	if state != db.StateCreationInProgress && state != db.StatePartiallyCreated && state != db.StateActive {
		return fmt.Errorf(
			"this method should be used only for updates regarding Connection Creation."+
				"Got state: %s", state,
		)
	}
	if creationError != "" && state != db.StatePartiallyCreated {
		return fmt.Errorf(
			"trying to store information about error for different state than partially created. "+
				"If the state is either in the progress or it was created successfully then no error "+
				"should appear. Got state: %s and error: %s",
			state, creationError,
		)
	}
	sourceVPC, err := sourceProvider.GetVPCForGateway(ctx, sourceGateway)
	if err != nil {
		c.logger.Warnf(
			"failed to obtain the ID of VPC attached to Gateway: %v", err)
	}
	destinationVPC, err := destProvider.GetVPCForGateway(ctx, destGateway)
	if err != nil {
		c.logger.Warnf(
			"failed to obtain the ID of VPC attached to Gateway: %v", err)
	}
	err = c.db.PutConnection(db.Connection{
		ID:                  c.connectionID,
		SourceID:            sourceGateway.Name,
		SourceProvider:      sourceGateway.CloudProvider,
		SourceVPC:           sourceVPC,
		SourceCIDRs:         c.sourceCIDRs,
		SourceRegion:        sourceGateway.Region,
		DestinationID:       destGateway.Name,
		DestinationProvider: destGateway.CloudProvider,
		DestinationVPC:      destinationVPC,
		DestinationCIDRs:    c.destinationCIDRs,
		DestinationRegion:   destGateway.Region,
		State:               state,
		Error:               creationError,
	})
	if err != nil {
		return fmt.Errorf("failed to update entry about the connection in the DB: %w", err)
	}
	return nil
}

func (c *Connector) updateConnectionDeletionEntryInDB(
	ctx context.Context,
	state db.ConnectionState,
	deletionError string,
	sourceGateway, destGateway types.Gateway,
	sourceProvider, destProvider provider.Provider,
) error {
	if c.connectionID == "" {
		return errors.New("cannot figure out which connection is being updated. Missing connectionID")
	}
	if state != db.StateDeletionInProgress && state != db.StatePartiallyDeleted && state != db.StateDeleted {
		return fmt.Errorf(
			"this method should be used only for updates regarding Connection Deletion."+
				"Got state: %s", state,
		)
	}
	if deletionError != "" && state != db.StatePartiallyDeleted {
		return fmt.Errorf(
			"trying to store information about error for different state than partially deleted. "+
				"If the state is either in the progress or it was deleted successfully then no error "+
				"should appear. Got state: %s and error: %s",
			state, deletionError,
		)
	}
	sourceVPC, err := sourceProvider.GetVPCForGateway(ctx, sourceGateway)
	if err != nil {
		c.logger.Warnf(
			"failed to obtain the ID of VPC attached to Gateway: %v", err)
	}
	destinationVPC, err := destProvider.GetVPCForGateway(ctx, destGateway)
	if err != nil {
		c.logger.Warnf(
			"failed to obtain the ID of VPC attached to Gateway: %v", err)
	}
	err = c.db.PutConnection(db.Connection{
		ID:                  c.connectionID,
		SourceID:            sourceGateway.Name,
		SourceProvider:      sourceGateway.CloudProvider,
		SourceVPC:           sourceVPC,
		SourceCIDRs:         c.sourceCIDRs,
		SourceRegion:        sourceGateway.Region,
		DestinationID:       destGateway.Name,
		DestinationProvider: destGateway.CloudProvider,
		DestinationVPC:      destinationVPC,
		DestinationCIDRs:    c.destinationCIDRs,
		DestinationRegion:   destGateway.Region,
		State:               state,
		Error:               deletionError,
	})
	if err != nil {
		return fmt.Errorf("failed to update entry about the connection in the DB: %w", err)
	}
	if state == db.StateDeleted {
		if err := c.db.DeleteConnection(c.connectionID); err != nil {
			return fmt.Errorf(
				"cannot delete a connection entry %s from DB: %w",
				c.connectionID, err,
			)
		}
	}
	return nil
}

func (c *Connector) StoreErrorFromConnectionCreationInDB(
	ctx context.Context,
	creationError error,
	sourceGateway, destGateway types.Gateway,
	sourceProvider, destProvider provider.Provider,
) error {
	if c.connectionID == "" {
		c.logger.Debug(
			"Connection was not registered in the DB. Nothing to update",
		)
	}
	err := c.updateConnectionCreationEntryInDB(
		ctx,
		db.StatePartiallyCreated,
		creationError.Error(),
		sourceGateway,
		destGateway,
		sourceProvider,
		destProvider,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update DB with information about received error. Got issue: %w",
			err)
	}
	return nil
}

func (c *Connector) Disconnect(ctx context.Context, request types.Request) error {
	c.logger.Debugf("Starting to delete a Connection with the configuration: %v", c.config)

	sourceProvider, destinationProvider, err := c.initializeProviders(ctx, request)
	if err != nil {
		return fmt.Errorf("cannot create connection due to provider failure: %w", err)
	}
	sourceGW, destGW, err := c.getGateways(ctx, request, sourceProvider, destinationProvider)
	if err != nil {
		return fmt.Errorf("validation of requested gateways to be connected failed: %w", err)
	}

	if err := c.updateConnectionDeletionEntryInDB(
		ctx,
		db.StateDeletionInProgress,
		"",
		sourceGW,
		destGW,
		sourceProvider,
		destinationProvider); err != nil {
		return fmt.Errorf("failed to mark connection DB Entry as in process of deletion: %w", err)
	}

	if err := sourceProvider.DeleteConnectionResources(ctx, sourceGW, destGW); err != nil {
		return fmt.Errorf("failed to create resources for Source Gateway %s: %w", sourceGW.Name, err)
	}
	if err := destinationProvider.DeleteConnectionResources(ctx, destGW, sourceGW); err != nil {
		return fmt.Errorf("failed to create resources for Destination Gateway %s: %w", destGW.Name, err)
	}

	if err := c.updateConnectionDeletionEntryInDB(
		ctx,
		db.StateDeleted,
		"",
		sourceGW,
		destGW,
		sourceProvider,
		destinationProvider); err != nil {
		c.logger.Warn("Resources were removed successfully but failed to remove DB Connection Entry")
		return fmt.Errorf("failed to remove connection DB Entry: %w", err)
	}
	c.logger.Info("Connection removed successfully")
	return nil
}

func (c *Connector) ListGateways(ctx context.Context) []types.Gateway {
	c.providerManager.InitializeAvailableProviders(ctx, c.logger, c.config.Providers)

	gateways := []types.Gateway{}
	for _, provider := range c.providerManager.Providers() {
		providerGateways, err := provider.ListGateways(ctx)
		if err != nil {
			c.logger.Errorf("cannot list %s Gateways: %v", provider.Name(), err)
		}
		gateways = append(gateways, providerGateways...)
	}
	return gateways
}

func (c *Connector) ListConnections() ([]db.Connection, error) {
	connections, err := c.db.ListConnections()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to list DB Connections: %w", err,
		)
	}
	return connections, nil
}

func (c *Connector) getGateways(
	ctx context.Context,
	request types.Request,
	sourceProvider, destProvider provider.Provider,
) (types.Gateway, types.Gateway, error) {
	c.logger.Debug("Checking the presence of requested Gateways")
	sourceGW, err := sourceProvider.GetGateway(ctx, request.ConnectionDetails.Source)
	if err != nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"The Source Gateway %v not found: %w",
			request.ConnectionDetails.Source,
			err)
	}
	if sourceGW == nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"The Source Gateway %v not found",
			request.ConnectionDetails.Source)
	}
	destGW, err := destProvider.GetGateway(ctx, request.ConnectionDetails.Destination)
	if err != nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"The Destination Gateway %v not found: %w",
			request.ConnectionDetails.Destination,
			err)
	}
	if destGW == nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"The Destination Gateway %v not found",
			request.ConnectionDetails.Destination)
	}
	return *sourceGW, *destGW, nil
}
