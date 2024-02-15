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
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/app-net-interface/awi-infra-guard/connector"
	"github.com/app-net-interface/awi-infra-guard/connector/db"
	"github.com/app-net-interface/awi-infra-guard/connector/types"
)

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
		Use:   "connect <Connection Name> <Source CSP> <Source ID> <Source Region> <Destination CSP> <Destination ID> <Destination Region>",
		Short: "creates connection between CSPs",
		Args:  cobra.ExactArgs(6),
		Run: func(cmd *cobra.Command, args []string) {
			connectCSPs(ctx, logger, cmd, args, *config)
		},
	}
	parent.AddCommand(cmd)
}

func connectCSPs(
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
	source := types.GatewayIdentifier{
		Provider:  args[0],
		GatewayID: args[1],
		Region:    args[2],
	}
	dest := types.GatewayIdentifier{
		Provider:  args[3],
		GatewayID: args[4],
		Region:    args[5],
	}

	if err = c.Connect(
		ctx,
		types.Request{
			ConnectionDetails: &types.ConnectionDetails{
				Source:      source,
				Destination: dest,
			},
		},
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

// TODO: Handle disconnecting CSP when
// there is no entry in the Database
func addDisconnectCommand(
	ctx context.Context,
	logger *logrus.Logger,
	config *connector.Config,
	parent *cobra.Command,
) {
	cmd := &cobra.Command{
		Use:   "disconnect <Connection ID>",
		Short: "removes connection between CSPs",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			disconnectCSPs(ctx, logger, cmd, args, *config)
		},
	}
	parent.AddCommand(cmd)
}

func disconnectCSPs(
	ctx context.Context,
	logger *logrus.Logger,
	cmd *cobra.Command,
	args []string,
	config connector.Config,
) {
	dbClient, err := db.NewClient(db.DefaultDBFile, logger.WithField("logger", "db"))
	if err != nil {
		logger.Errorf("Could not create DB Client for checking the connection: %v", err)
		return
	}
	connection, err := dbClient.GetConnection(args[0])
	if err != nil {
		logger.Errorf(
			"failed to obtain information about connection '%s': %v",
			args[0], err,
		)
		return
	}
	dbClient.Close()
	if connection == nil {
		logger.Errorf(
			"connection with ID not found '%s'. Aborting it",
			args[0],
		)
	}

	c, err := connector.NewConnector(ctx, logger, &config)
	if err != nil {
		logger.Errorf("Failed to initialize Connector: %v", err)
		os.Exit(1)
	}
	defer closeConnector(c)
	c.SetConnectionID(connection.ID)
	if err = c.Disconnect(
		ctx,
		types.Request{
			ConnectionDetails: &types.ConnectionDetails{
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
			},
		}); err != nil {
		logger.Errorf("Failed to remove a Connection: %v", err)
		os.Exit(1)
	}
}
