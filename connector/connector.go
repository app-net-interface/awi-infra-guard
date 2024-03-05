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
	"github.com/app-net-interface/awi-infra-guard/connector/secret"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/sirupsen/logrus"
)

type Connector struct {
	providerManager *ProviderManager
	db              *db.Client
	config          *Config
	logger          *logrus.Entry
	mainLogger      *logrus.Logger
}

func NewConnector(ctx context.Context, logger *logrus.Logger, config *Config) (*Connector, error) {
	providerManager, err := NewProviderManager(logger.WithField("logger", "providerManager"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Provider Manager: %w", err)
	}
	dbClient, err := db.NewClient(db.DefaultDBFile, logger.WithField("logger", "db"))
	if err != nil {
		return nil, fmt.Errorf("could not create DB Client: %w", err)
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

	var MaxNumberOfTunnels uint8
	if source.MaxNumberOfTunnels != 0 && destination.MaxNumberOfTunnels != 0 {
		MaxNumberOfTunnels = min(source.MaxNumberOfTunnels, destination.MaxNumberOfTunnels)
	} else {
		MaxNumberOfTunnels = max(source.MaxNumberOfTunnels, destination.MaxNumberOfTunnels)
	}
	if MaxNumberOfTunnels != 0 {
		source.NumberOfInterfaces = min(MaxNumberOfTunnels, source.NumberOfInterfaces)
		destination.NumberOfInterfaces = min(MaxNumberOfTunnels, destination.NumberOfInterfaces)
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

type gateways struct {
	SourceGateway       types.Gateway
	DestinationGateway  types.Gateway
	SourceProvider      provider.Provider
	DestinationProvider provider.Provider
}

func (c *Connector) generateSecrets(numberOfInterfaces uint8) ([]string, error) {
	secrets := make([]string, numberOfInterfaces)
	for i := range secrets {
		var err error
		secrets[i], err = secret.GeneratePSK()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to generate Shared Secret: %w", err,
			)
		}
	}
	return secrets, nil
}

func (c *Connector) createConnectionWithStaticRouting(
	ctx context.Context,
	gateways gateways,
	sourceInterfaces,
	destinationInterfaces []string,
	numberOfTunnels uint8,
	secrets []string,
) error {
	return errors.New("NOT IMPLEMENTED")
}

type connectionType int

const (
	connectionTypeBGP connectionType = iota
	connectionTypeStatic
	connectionTypeNone
)

// checkConnectionTypeThatCanBeEstablished verifies setting of
// both sides of the possible connection, checks if the connection
// can be established and if so it returns what kind of connection
// can be established and how many tunnels.
func (c *Connector) checkConnectionTypeThatCanBeEstablished(
	ctx context.Context,
	gateways gateways,
) (connectionType, uint8, error) {
	sourceSetting, err := gateways.SourceProvider.GetGatewayConnectionSettings(
		ctx, gateways.SourceGateway)
	if err != nil {
		return connectionTypeNone, 0, fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			gateways.SourceGateway.Name, gateways.SourceProvider.Name(), err,
		)
	}
	destinationSetting, err := gateways.DestinationProvider.GetGatewayConnectionSettings(
		ctx, gateways.DestinationGateway)
	if err != nil {
		return connectionTypeNone, 0, fmt.Errorf(
			"failed to get connection settings for gateway %s from provider %s: %w",
			gateways.DestinationGateway.Name, gateways.DestinationProvider.Name(), err,
		)
	}
	numberOfTunnels, err := calculateNumberOfTunnels(sourceSetting, destinationSetting)
	if err != nil {
		return connectionTypeNone, 0, fmt.Errorf("failed to calculate number of desired tunnels: %w", err)
	}

	if sourceSetting.BGPSetting != nil && destinationSetting.BGPSetting != nil {
		scenario := types.BGPScenarioFromBothConfigs(
			&sourceSetting.BGPSetting.Addressing, &destinationSetting.BGPSetting.Addressing)
		if scenario != nil {
			return connectionTypeBGP, numberOfTunnels, nil
		}
		c.logger.Infof(
			"both sides of connection support BGP but their configuration "+
				"do not match with each other. Trying Static Routing Setting."+
				"Source Gateway: %s:%s and DestinationGateway: %s:%s",
			gateways.SourceProvider.Name(),
			gateways.SourceGateway.Name,
			gateways.DestinationProvider.Name(),
			gateways.DestinationGateway.Name,
		)
	}

	if sourceSetting.StaticRoutingSetting != nil && destinationSetting.StaticRoutingSetting != nil {
		return connectionTypeStatic, numberOfTunnels, nil
	}

	return connectionTypeNone, 0, nil
}

func (c *Connector) createConnection(
	ctx context.Context,
	gateways gateways,
) error {
	connType, numberOfTunnels, err := c.checkConnectionTypeThatCanBeEstablished(ctx, gateways)
	if err != nil {
		return fmt.Errorf(
			"failed to verify what kind of connection can be established between gateways "+
				"%s:%s and %s:%s: %w",
			gateways.SourceGateway.CloudProvider, gateways.SourceGateway.Name,
			gateways.DestinationGateway.CloudProvider, gateways.DestinationGateway.Name,
			err,
		)
	}
	if connType == connectionTypeNone {
		return fmt.Errorf(
			"cannot establish a connection between gateways %s:%s and %s:%s. Based on "+
				"gateways settings its not possible to establish either Connection with "+
				"static routing nor dynamic routing using BGP",
			gateways.SourceGateway.CloudProvider, gateways.SourceGateway.Name,
			gateways.DestinationGateway.CloudProvider, gateways.DestinationGateway.Name,
		)
	}

	err = gateways.SourceProvider.InitializeCreation(
		ctx, gateways.SourceGateway, gateways.DestinationGateway,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s: %w",
			gateways.SourceGateway.Name, err)
	}

	err = gateways.DestinationProvider.InitializeCreation(
		ctx, gateways.DestinationGateway, gateways.SourceGateway,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s: %w",
			gateways.DestinationGateway.Name, err)
	}

	sourceInterfaces, err := gateways.SourceProvider.InitializeGatewayInterfaces(
		ctx, gateways.SourceGateway, gateways.DestinationGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s interfaces: %w",
			gateways.SourceGateway.Name, err)
	}
	destInterfaces, err := gateways.DestinationProvider.InitializeGatewayInterfaces(
		ctx, gateways.DestinationGateway, gateways.SourceGateway)
	if err != nil {
		return fmt.Errorf(
			"failed to initialize Gateway %s interfaces: %w",
			gateways.DestinationGateway.Name, err)
	}

	secrets, err := c.generateSecrets(numberOfTunnels)
	if err != nil {
		return fmt.Errorf(
			"failed to prepare pre-shared secrets: %w", err,
		)
	}

	if connType == connectionTypeBGP {
		return c.createConnectionWithBGP(
			ctx,
			gateways,
			sourceInterfaces,
			destInterfaces,
			numberOfTunnels,
			secrets,
		)
	}
	return c.createConnectionWithStaticRouting(
		ctx,
		gateways,
		sourceInterfaces,
		destInterfaces,
		numberOfTunnels,
		secrets,
	)
}

func (c *Connector) Connect(ctx context.Context, request types.Request) error {
	var connectionID string
	var connectionDetails types.ConnectionDetails

	connection, err := getConnectionEntryIfExists(
		ctx,
		c.db,
		c.logger,
		request,
	)
	if err != nil {
		return fmt.Errorf(
			"got error while trying to find a Connection entry in the DB for configuration: %v",
			c.config)
	}
	if connection == nil {
		c.logger.Debugf("connection does not exist yet")
		if request.ConnectionDetails == nil {
			return errors.New("cannot create a connection without connection ID or details provided")
		}
		connectionDetails = *request.ConnectionDetails
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
		connectionID = connection.ID
		connectionDetails = connectionDetailsFromConnection(*connection)
	}

	c.logger.Debugf("Starting to create a Connection: %v", request)

	sourceProvider, destinationProvider, err := c.initializeProviders(ctx, connectionDetails)
	if err != nil {
		return fmt.Errorf("cannot create connection due to provider failure: %w", err)
	}
	sourceGW, destGW, err := c.getGateways(ctx, connectionDetails, sourceProvider, destinationProvider)
	if err != nil {
		return fmt.Errorf("validation of requested gateways to be connected failed: %w", err)
	}

	if connectionID == "" {
		connectionID, err = createConnectionEntryInDB(
			ctx, c.db, c.logger, sourceGW, destGW,
		)
		if err != nil {
			return fmt.Errorf("connection could not be stored in the DB: %w", err)
		}
	}

	err = c.createConnection(
		ctx, gateways{
			SourceGateway:       sourceGW,
			DestinationGateway:  destGW,
			SourceProvider:      sourceProvider,
			DestinationProvider: destinationProvider,
		})
	if err != nil {
		return fmt.Errorf(
			"failed to create a connection between %s:%s and %s:%s: %w",
			sourceProvider.Name(), sourceGW.Name, destinationProvider.Name(), destGW.Name,
			err)
	}

	sourceCIDRs, err := sourceProvider.GetCIDRs(ctx, sourceGW)
	if err != nil {
		return fmt.Errorf(
			"failed to obtain list of CIDRs behind Source Gateway %s:%s: %w",
			sourceProvider.Name(), sourceGW.Name, err)
	}
	destinationCIDRs, err := destinationProvider.GetCIDRs(ctx, destGW)
	if err != nil {
		return fmt.Errorf(
			"failed to obtain list of CIDRs behind Destination Gateway %s:%s: %w",
			destinationProvider.Name(), destGW.Name, err,
		)
	}

	if err = updateConnectionEntryWithGatewayCIDRsInDB(
		ctx, c.db, c.logger, connectionID, sourceCIDRs, destinationCIDRs,
	); err != nil {
		c.logger.Warn("Resources were create successfully but failed to update DB Connection Entry with CIDRs")
		return fmt.Errorf("failed to update connection DB Entry with CIDRs: %w", err)
	}

	if err = updateConnectionEntryStateInDB(ctx, c.db, c.logger, connectionID, db.StateActive, ""); err != nil {
		c.logger.Warn("Resources were create successfully but failed to update DB Connection Entry with Active state")
		return fmt.Errorf("failed to update connection DB Entry with Active state: %w", err)
	}

	c.logger.Info("Connection created successfully")
	return nil
}

func (c *Connector) initializeProviders(
	ctx context.Context, connectionDetails types.ConnectionDetails,
) (provider.Provider, provider.Provider, error) {
	sourceConfig := ""
	sourceProviderName := connectionDetails.Source.Provider
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
	destProviderName := connectionDetails.Destination.Provider
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

func connectionDetailsFromConnection(connection db.Connection) types.ConnectionDetails {
	// TODO: Source/Destination Details are not part of a DB.
	//
	// Adjust model to properly use that field or remove it.
	return types.ConnectionDetails{
		Source: types.GatewayIdentifier{
			GatewayID: connection.SourceID,
			Region:    connection.SourceRegion,
			Provider:  connection.SourceProvider,
		},
		Destination: types.GatewayIdentifier{
			GatewayID: connection.DestinationID,
			Region:    connection.DestinationRegion,
			Provider:  connection.DestinationProvider,
		},
	}
}

func (c *Connector) Disconnect(ctx context.Context, request types.Request) error {
	c.logger.Debugf("Starting to delete a Connection with the configuration: %v", c.config)

	var connectionDetails types.ConnectionDetails
	var connectionID string

	if request.ConnectionID != nil {
		connection, err := getConnectionEntryIfExists(ctx, c.db, c.logger, request)
		if err != nil {
			return fmt.Errorf(
				"cannot remove the connection as obtaining connection with ID '%s'"+
					"from DB failed. If the connection is not present in the DB you can provide "+
					"source and destination details directly instead of providing connection ID: %w",
				*request.ConnectionID, err,
			)
		}
		if connection == nil {
			return fmt.Errorf(
				"cannot remove the connection as obtaining connection with ID '%s'"+
					"from DB failed - got nil connection. If the connection is not present "+
					"in the DB you can provide source and destination details directly instead "+
					"of providing connection ID",
				*request.ConnectionID,
			)
		}
		connectionDetails = connectionDetailsFromConnection(*connection)
		connectionID = connection.ID
	} else {
		if request.ConnectionDetails == nil {
			return errors.New(
				"cannot remove the connection without Connection ID or details provided")
		}
		connection, err := c.db.GetConnectionWithGateways(
			request.ConnectionDetails.Source.GatewayID,
			request.ConnectionDetails.Source.Provider,
			request.ConnectionDetails.Destination.GatewayID,
			request.ConnectionDetails.Destination.Provider,
		)
		if err != nil {
			c.logger.Warnf(
				"failed to find a connection for %s:%s and %s:%s in the database. "+
					"Will attempt to delete a connection without DB entry. Got error: %v",
				request.ConnectionDetails.Source.Provider,
				request.ConnectionDetails.Source.GatewayID,
				request.ConnectionDetails.Destination.Provider,
				request.ConnectionDetails.Destination.GatewayID,
				err,
			)
		} else {
			if connection == nil {
				c.logger.Warnf(
					"failed to find a connection for %s:%s and %s:%s in the database. "+
						"Will attempt to delete a connection without DB entry. Got nil object",
					request.ConnectionDetails.Source.Provider,
					request.ConnectionDetails.Source.GatewayID,
					request.ConnectionDetails.Destination.Provider,
					request.ConnectionDetails.Destination.GatewayID,
				)
			} else {
				connectionID = connection.ID
			}
		}
		connectionDetails = *request.ConnectionDetails
	}

	sourceProvider, destinationProvider, err := c.initializeProviders(ctx, connectionDetails)
	if err != nil {
		return fmt.Errorf("cannot create connection due to provider failure: %w", err)
	}
	sourceGW, destGW, err := c.getGateways(ctx, connectionDetails, sourceProvider, destinationProvider)
	if err != nil {
		return fmt.Errorf("validation of requested gateways to be connected failed: %w", err)
	}

	if connectionID != "" {
		err = updateConnectionEntryStateInDB(
			ctx, c.db, c.logger, connectionID, db.StateDeletionInProgress, "",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to update deletion state in the DB of connection: %s: %w",
				connectionID, err,
			)
		}
	}

	if err := sourceProvider.DeleteConnectionResources(ctx, sourceGW, destGW); err != nil {
		return fmt.Errorf("failed to create resources for Source Gateway %s: %w", sourceGW.Name, err)
	}
	if err := destinationProvider.DeleteConnectionResources(ctx, destGW, sourceGW); err != nil {
		return fmt.Errorf("failed to create resources for Destination Gateway %s: %w", destGW.Name, err)
	}

	if connectionID != "" {
		if err := deleteConnectionEntryInDB(ctx, c.db, c.logger, connectionID); err != nil {
			c.logger.Warn("Resources were removed successfully but failed to remove DB Connection Entry")
			return fmt.Errorf("failed to remove connection DB Entry: %w", err)
		}
		c.logger.Infof("Connection '%s' removed successfully", connectionID)
	} else {
		c.logger.Infof("Connection %s:%s - %s:%s removed successfully",
			request.ConnectionDetails.Source.Provider,
			request.ConnectionDetails.Source.GatewayID,
			request.ConnectionDetails.Destination.Provider,
			request.ConnectionDetails.Destination.GatewayID,
		)
	}

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
	connectionDetails types.ConnectionDetails,
	sourceProvider, destProvider provider.Provider,
) (types.Gateway, types.Gateway, error) {
	c.logger.Debug("Checking the presence of requested Gateways")
	sourceGW, err := sourceProvider.GetGateway(ctx, connectionDetails.Source)
	if err != nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"the Source Gateway %v not found: %w",
			connectionDetails.Source,
			err)
	}
	if sourceGW == nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"the Source Gateway %v not found",
			connectionDetails.Source)
	}
	destGW, err := destProvider.GetGateway(ctx, connectionDetails.Destination)
	if err != nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"the Destination Gateway %v not found: %w",
			connectionDetails.Destination,
			err)
	}
	if destGW == nil {
		return types.Gateway{}, types.Gateway{}, fmt.Errorf(
			"the Destination Gateway %v not found",
			connectionDetails.Destination)
	}
	return *sourceGW, *destGW, nil
}
