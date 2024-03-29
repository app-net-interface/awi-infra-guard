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

	"github.com/app-net-interface/awi-infra-guard/connector/db"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func getConnectionEntryIfExists(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	request types.Request,
) (*db.Connection, error) {
	if request.ConnectionID != nil {
		return dbClient.GetConnection(*request.ConnectionID)
	}
	return dbClient.GetConnectionWithGateways(
		request.ConnectionDetails.Source.GatewayID,
		request.ConnectionDetails.Source.Provider,
		request.ConnectionDetails.Destination.GatewayID,
		request.ConnectionDetails.Destination.Provider,
	)
}

func createConnectionEntryInDB(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	sourceGateway, destGateway types.Gateway,
) (string, error) {
	logger.Debugf("Creating new entry for the connection in DB")

	conn := connectionFromGateways(
		sourceGateway, destGateway, db.Connection{},
	)
	conn.State = db.StateCreationInProgress
	err := dbClient.PutConnection(conn)
	if err != nil {
		return "", fmt.Errorf("failed to create entry about the connection in the DB: %w", err)
	}
	return conn.ID, nil
}

func deleteConnectionEntryInDB(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	connectionID string,
) error {
	if connectionID == "" {
		return errors.New("cannot delete connection without ID provided")
	}
	logger.Debugf("Deleting entry '%s' for the connection in DB", connectionID)
	err := dbClient.DeleteConnection(connectionID)
	if err != nil {
		return fmt.Errorf("failed to delete entry '%s' in the DB: %w", connectionID, err)
	}
	return nil
}

func updateConnectionEntryWithGatewaysInDB(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	connectionID string,
	sourceGateway, destGateway types.Gateway) error {
	if connectionID == "" {
		return errors.New("cannot update connection without ID provided")
	}
	conn, err := dbClient.GetConnection(connectionID)
	if err != nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: %w", connectionID, err,
		)
	}
	if conn == nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: got nil object", connectionID,
		)
	}
	err = dbClient.PutConnection(connectionFromGateways(
		sourceGateway, destGateway, *conn,
	))
	if err != nil {
		return fmt.Errorf("failed to update entry about the connection in the DB: %w", err)
	}
	return nil
}

func updateConnectionEntryWithGatewayCIDRsInDB(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	connectionID string,
	sourceCIDRs, destCIDRs []string) error {
	if connectionID == "" {
		return errors.New("cannot update connection without ID provided")
	}
	conn, err := dbClient.GetConnection(connectionID)
	if err != nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: %w", connectionID, err,
		)
	}
	if conn == nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: got nil object", connectionID,
		)
	}
	conn.SourceCIDRs = sourceCIDRs
	conn.DestinationCIDRs = destCIDRs
	err = dbClient.PutConnection(*conn)
	if err != nil {
		return fmt.Errorf("failed to update entry about the connection in the DB: %w", err)
	}
	return nil
}

func updateConnectionEntryStateInDB(
	ctx context.Context,
	dbClient *db.Client,
	logger *logrus.Entry,
	connectionID string,
	state db.ConnectionState,
	stateError string,
) error {
	if connectionID == "" {
		return errors.New("cannot update connection without ID provided")
	}
	conn, err := dbClient.GetConnection(connectionID)
	if err != nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: %w", connectionID, err,
		)
	}
	if conn == nil {
		return fmt.Errorf(
			"cannot get connection with ID %s: got nil object", connectionID,
		)
	}
	conn.State = state
	conn.Error = stateError
	err = dbClient.PutConnection(*conn)
	if err != nil {
		return fmt.Errorf("failed to update state about the connection in the DB: %w", err)
	}
	return nil
}

func connectionFromGateways(
	sourceGateway, destGateway types.Gateway,
	oldConnection db.Connection,
) db.Connection {
	if oldConnection.ID == "" {
		oldConnection.ID = string(uuid.New().String())
	}
	oldConnection.SourceID = sourceGateway.Name
	oldConnection.SourceProvider = sourceGateway.CloudProvider
	oldConnection.SourceVPC = sourceGateway.VPC
	oldConnection.SourceRegion = sourceGateway.Region
	oldConnection.DestinationID = destGateway.Name
	oldConnection.DestinationProvider = destGateway.CloudProvider
	oldConnection.DestinationVPC = destGateway.VPC
	oldConnection.DestinationRegion = destGateway.Region
	return oldConnection
}
