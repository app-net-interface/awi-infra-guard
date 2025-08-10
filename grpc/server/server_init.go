// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/app-net-interface/awi-infra-guard/db"
	sqlitedb "github.com/app-net-interface/awi-infra-guard/db/sqlite"
	"github.com/app-net-interface/awi-infra-guard/grpc/config"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/sync"
)

const configPath = "config.yaml"

type Server struct {
	logger *logrus.Logger
	infrapb.UnimplementedCloudProviderServiceServer
	infrapb.UnimplementedAccessControlServiceServer
	infrapb.UnimplementedKubernetesServiceServer
	strategy provider.Strategy
}

func setLoggingLevel(config config.Config, logger *logrus.Logger) error {
	switch config.LogLevel {
	case "PANIC":
		logger.SetLevel(logrus.PanicLevel)
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "TRACE":
		logger.SetLevel(logrus.TraceLevel)
	default:
		{
			return fmt.Errorf(
				"invalid log level '%s' in configuration. Supported levels are: "+
					"PANIC, FATAL, ERROR, WARN, INFO, DEBUG, TRACE",
				config.LogLevel,
			)
		}
	}
	return nil
}

func parseConfig(logger *logrus.Logger) config.Config {

	config := config.Config{
		Hostname:   "",
		Port:       "50052",
		UseLocalDB: true,
		LogLevel:   "INFO",
		SyncConfig: config.SyncConfig{
			DbFileName:   "infra.db",
			SyncWaitTime: time.Second * 300,
			Resources: config.Resources{
				Cloud:      []string{},
				Kubernetes: []string{},
			},
		},
		Providers:           []config.Provider{},
		KubernetesSupported: false,
	}

	err := initConfig(configPath, &config)
	if err != nil {
		logger.Errorf("Failed to parse config: %v using default values...", err)
	}
	if err = setLoggingLevel(config, logger); err != nil {
		logger.Errorf("Failed to set logging level: %v", err)
	}
	logger.Infof("Using configuration: %+v", config)
	return config
}

func initConfig(configFilePath string, c *config.Config) error {
	viper.AutomaticEnv()
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		if _, match := err.(viper.UnsupportedConfigError); match || errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("unsupported Config or File doesn't exist")
		}
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to decode into struct: %v", err)
	}

	// Get the configuration values using Viper
	// Overriding values from environment variables if set
	if port := viper.GetString("port"); port != "" {
		c.Port = port
	}
	if logLevel := viper.GetString("logLevel"); logLevel != "" {
		c.LogLevel = logLevel
	}
	if hostname := viper.GetString("hostname"); hostname != "" {
		c.Hostname = hostname
	}
	if useDB := viper.GetBool("useLocalDB"); useDB {
		c.UseLocalDB = useDB
	}

	if err := viper.UnmarshalKey("providers", &c.Providers); err != nil {
		return fmt.Errorf("unable to decode providers into struct: %v", err)
	}

	if c.UseLocalDB {
		var syncConfig config.SyncConfig
		if err := viper.UnmarshalKey("syncConfig", &syncConfig); err != nil {
			return fmt.Errorf("unable to decode syncConfig into struct: %v", err)
		}
		fmt.Printf("Config: %+v\n", syncConfig)
		c.SyncConfig = syncConfig
	}

	return nil
}

func Run() {
	ctx := context.Background()
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	}

	c := parseConfig(logger)
	fmt.Printf("Provider Config: %+v\n", c.Providers)

	providerStrategy, err := provider.NewRealProviderStrategy(logger, c.Providers, c.KubernetesSupported)

	if err != nil {
		logger.Warnf("Initialized with error %v", err)
	}

	var responseStrategy provider.Strategy
	responseStrategy = providerStrategy
	if c.UseLocalDB {
		logger.Infof("Initializing local database")
		dbClient := sqlitedb.NewSQLiteClient().(db.Client)
		if err := dbClient.Open(c.SyncConfig.DbFileName); err != nil {
			logger.Errorf("could not opend db: %v", err)
			return
		}
		defer func(dbClient db.Client) {
			err := dbClient.Close()
			if err != nil {
				logger.Errorf("Failed to close db: %v", err)
			}
		}(dbClient)

		strategyWithDB := db.NewStrategyWithDB(dbClient, providerStrategy, logger, c.KubernetesSupported)
		responseStrategy = strategyWithDB

		if c.SyncConfig.Enabled {
			syncer := sync.NewSyncer(logger, dbClient, providerStrategy, &c)
			go syncer.SyncPeriodically(ctx)
		}
	}

	s := &Server{
		logger:   logger,
		strategy: responseStrategy,
	}

	lis, err := net.Listen("tcp", c.Hostname+":"+c.Port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(s.unaryServerInterceptor),
	)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	infrapb.RegisterCloudProviderServiceServer(grpcServer, s)
	infrapb.RegisterAccessControlServiceServer(grpcServer, s)
	infrapb.RegisterKubernetesServiceServer(grpcServer, s)

	logger.Infof("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
	if c.KubernetesSupported {
		go s.refreshClusters(ctx, time.Second*60)
	}

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Application started. Press Ctrl+C to exit.")
	}()

	<-stop // Wait for interrupt signal

	logger.Info("Shutting down server...")

	// Stop the provider strategy's periodic reconciler
	//if providerStrategy != nil {
	//	logger.Info("Stopping provider strategy reconciler...")
	//	providerStrategy.StopPeriodicReconciliation() // Call the new stop method
	//}

	// Stop the gRPC server gracefully (example)
	// if grpcServer != nil {
	// 	logger.Info("Stopping gRPC server...")
	// 	grpcServer.GracefulStop()
	// }

	// If syncer was started with a cancellable context, cancel it.
	// cancelMainCtx() // Already handled by defer if Run function exits

	logger.Info("Server gracefully stopped.")
}
