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

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/app-net-interface/awi-infra-guard/connector"
	"github.com/app-net-interface/awi-infra-guard/connector/db"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

var emptyStr string = ""

var connectionRequest types.Request = types.Request{
	ConnectionID:      &emptyStr,
	ConnectionDetails: &types.ConnectionDetails{},
}
var connectionRequestFile string

func main() {
	ctx := context.Background()

	config, err := connector.LoadConfig("csp.yaml")
	if err != nil {
		fmt.Printf("Failed to initialize configuration: %v\n", err)
		os.Exit(1)
	}

	// TODO: Adjust logger to match the provided configuration.
	logger := configureLogger()

	command := createRootCommand(ctx, logger, &config)

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Formatter = &logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	return logger
}

func createRootCommand(
	ctx context.Context, logger *logrus.Logger, store *connector.Config,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cspconnector",
		Short: "cspconnector is a project for connecting different cloud providers",
		Long:  "cspconnector is a project for connecting different cloud providers",
	}
	addListCommand(ctx, logger, store, cmd)
	addConnectCommand(ctx, logger, store, cmd)
	addDisconnectCommand(ctx, logger, store, cmd)
	return cmd
}

func addListCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:   "list [resource_type]",
		Short: "lists given resource",
		Long:  "cspconnector is a project for connecting different cloud providers",
	}
	addListGatewaysCommand(ctx, logger, config, cmd)
	addListConnectionsCommand(ctx, logger, config, cmd)
	parent.AddCommand(cmd)
}

func addListGatewaysCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:     "gateway",
		Aliases: []string{"gateways"},
		Short:   "lists GCP Cloud Routers",
		Run: func(cmd *cobra.Command, args []string) {
			listGateways(ctx, logger, cmd, args, *config)
		},
	}
	parent.AddCommand(cmd)
}

func listGateways(
	ctx context.Context,
	logger *logrus.Logger,
	cmd *cobra.Command,
	args []string,
	config connector.Config,
) {
	c, err := connector.NewConnector(ctx, logger, &config)
	if err != nil {
		logger.Errorf("Failed to initialize Connector: %v", err)
		os.Exit(1)
	}
	defer closeConnector(c)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Gateway ID", "Cloud Provider", "Kind", "VPC ID", "ASN"})

	gateways := c.ListGateways(ctx)
	for _, g := range gateways {
		table.Append(
			[]string{
				g.Name,
				g.CloudProvider,
				g.Kind,
				g.VPC,
				g.ASN})
	}

	fmt.Println("CSP Gateways")
	table.Render()
}

// A wrapper for closing the connection and printing
// the error if returned.
func closeConnector(c *connector.Connector) {
	if err := c.Close(); err != nil {
		fmt.Printf("Got error while trying to Close the Connector: %v\n", err)
	}
}

func addListConnectionsCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:     "connection",
		Aliases: []string{"connections"},
		Short:   "lists available CSPs connections",
		Run: func(cmd *cobra.Command, args []string) {
			listConnections(ctx, logger, cmd, args, *config)
		},
	}
	parent.AddCommand(cmd)
}

func listConnections(
	ctx context.Context,
	logger *logrus.Logger,
	cmd *cobra.Command,
	args []string,
	config connector.Config,
) {
	c, err := connector.NewConnector(ctx, logger, &config)
	if err != nil {
		logger.Errorf("Failed to initialize Connector: %v", err)
		os.Exit(1)
	}
	defer closeConnector(c)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Connection ID", "Source Provider", "Source ID", "Destination Provider", "Destination ID", "State"})

	connections, err := c.ListConnections()
	if err != nil {
		logger.Errorf("failed to list connections: %v", err)
	}

	for _, conn := range connections {
		table.Append(
			[]string{
				conn.ID,
				conn.SourceProvider,
				conn.SourceID,
				conn.DestinationProvider,
				conn.DestinationID,
				string(conn.State)})
	}

	fmt.Println("CSP Connections")
	table.Render()
}

func addConnectCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "creates connection between CSPs",
		Run: func(cmd *cobra.Command, args []string) {
			connectCSPs(ctx, logger, *config)
		},
	}
	parent.AddCommand(cmd)
	addConnectionRequestFlags(cmd.Flags())

}

func connectCSPs(
	ctx context.Context,
	logger *logrus.Logger,
	config connector.Config,
) {
	c, err := connector.NewConnector(ctx, logger, &config)
	if err != nil {
		logger.Errorf("Failed to initialize Connector: %v", err)
		os.Exit(1)
	}
	defer closeConnector(c)

	req, err := parseRequest()
	if err != nil {
		logger.Errorf("Failed to parse Connection Request: %v", err)
		os.Exit(1)
	}

	if err = c.Connect(
		ctx,
		req,
	); err != nil {
		logger.Errorf("Failed to create a Connection: %v", err)
		// TODO: Handle saving the error from creating the connection.
		// dbErr := c.StoreErrorFromConnectionCreationInDB(ctx, err)
		// if dbErr != nil {
		// 	logger.Errorf("Failed to update DB with Connection error: %v", logger)
		// }
		os.Exit(1)
	}
}

func addDisconnectCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:   "disconnect",
		Short: "removes connection between CSPs",
		Run: func(cmd *cobra.Command, args []string) {
			disconnectCSPs(ctx, logger, *config)
		},
	}
	parent.AddCommand(cmd)
	addConnectionRequestFlags(cmd.Flags())
}

func gwIdentifiersFromConnection(
	logger *logrus.Logger, connectionID string,
) (types.GatewayIdentifier, types.GatewayIdentifier, error) {
	dbClient, err := db.NewClient(db.DefaultDBFile, logger.WithField("logger", "db"))
	if err != nil {
		return types.GatewayIdentifier{}, types.GatewayIdentifier{}, fmt.Errorf(
			"could not create DB Client for checking the connection: %w", err,
		)
	}
	defer dbClient.Close()
	connection, err := dbClient.GetConnection(connectionID)
	if err != nil {
		return types.GatewayIdentifier{}, types.GatewayIdentifier{}, fmt.Errorf(
			"failed to obtain information about connection '%s': %w",
			connectionID, err,
		)
	}
	if connection == nil {
		return types.GatewayIdentifier{}, types.GatewayIdentifier{}, fmt.Errorf(
			"connection with ID not found '%s'. Aborting it",
			connectionID,
		)
	}
	return types.GatewayIdentifier{
			GatewayID: connection.SourceID,
			Region:    connection.SourceRegion,
			Provider:  connection.SourceProvider,
		}, types.GatewayIdentifier{
			GatewayID: connection.DestinationID,
			Region:    connection.DestinationRegion,
			Provider:  connection.DestinationProvider,
		}, nil
}

func disconnectCSPs(
	ctx context.Context,
	logger *logrus.Logger,
	config connector.Config,
) {
	req, err := parseRequest()
	if err != nil {
		logger.Errorf("Failed to parse Connection Request: %v", err)
		os.Exit(1)
	}

	if req.ConnectionID != nil {
		source, destination, err := gwIdentifiersFromConnection(logger, *req.ConnectionID)
		if err != nil {
			logger.Errorf("Failed to load Gateway info: %v", err)
			return
		}
		req.ConnectionDetails = &types.ConnectionDetails{
			Source:      source,
			Destination: destination,
		}
	}

	c, err := connector.NewConnector(ctx, logger, &config)
	if err != nil {
		logger.Errorf("Failed to initialize Connector: %v", err)
		os.Exit(1)
	}
	defer closeConnector(c)

	if err = c.Disconnect(ctx, req); err != nil {
		logger.Errorf("Failed to remove a Connection: %v", err)
		os.Exit(1)
	}
}

func parseRequest() (types.Request, error) {
	if connectionRequestFile != "" {
		req, err := requestFromFile(connectionRequestFile)
		if err != nil {
			return types.Request{}, fmt.Errorf(
				"failed to parse Connection Request from a file: %w", err,
			)
		}
		return req, nil
	}
	if *connectionRequest.ConnectionID != "" {
		return types.Request{
			ConnectionID: connectionRequest.ConnectionID,
		}, nil
	}

	// Clean Connection ID
	connectionRequest.ConnectionID = nil
	if connectionRequest.ConnectionDetails.Source.GatewayID == "" &&
		connectionRequest.ConnectionDetails.Source.Provider == "" &&
		connectionRequest.ConnectionDetails.Source.Region == "" {
		return types.Request{}, errors.New(
			"failed to parse Connection Request - no handler for the source",
		)
	}
	if connectionRequest.ConnectionDetails.Destination.GatewayID == "" &&
		connectionRequest.ConnectionDetails.Destination.Provider == "" &&
		connectionRequest.ConnectionDetails.Destination.Region == "" {
		return types.Request{}, errors.New(
			"failed to parse Connection Request - no handler for the destination",
		)
	}
	return connectionRequest, nil
}

func addConnectionRequestFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(
		&connectionRequestFile,
		"from-file",
		"f",
		"",
		"Accepts a filepath to the yaml file containing "+
			"Connect/Disconnect request.",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Source.GatewayID,
		"src-id",
		"",
		"ID of Source Gateway for the connection",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Source.Provider,
		"src-provider",
		"",
		"The provider holding Source Gateway",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Source.Region,
		"src-region",
		"",
		"The region of the Source Gateway",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Destination.GatewayID,
		"dest-id",
		"",
		"ID of Destination Gateway for the connection",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Destination.Provider,
		"dest-provider",
		"",
		"The provider holding Destination Gateway",
	)
	flagSet.StringVar(
		&connectionRequest.ConnectionDetails.Destination.Region,
		"dest-region",
		"",
		"The region of the Destination Gateway",
	)
	flagSet.StringVar(
		connectionRequest.ConnectionID,
		"connection-id",
		"",
		"The ID of a connection to recreate/delete",
	)
}

func requestFromFile(filepath string) (types.Request, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return types.Request{}, fmt.Errorf("error reading YAML file: %w", err)
	}

	request := types.Request{}
	if err = yaml.Unmarshal(yamlFile, &request); err != nil {
		return types.Request{}, fmt.Errorf("error parsing YAML file: %w", err)
	}

	return request, nil
}
