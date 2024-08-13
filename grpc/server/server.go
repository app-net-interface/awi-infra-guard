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

package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	"github.com/app-net-interface/awi-infra-guard/db"
	"github.com/app-net-interface/awi-infra-guard/grpc/config"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/provider"
	"github.com/app-net-interface/awi-infra-guard/sync"
	"github.com/app-net-interface/awi-infra-guard/types"
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
	}

	c := parseConfig(logger)
	fmt.Printf("Provider Config: %+v\n", c.Providers)

	providerStrategy, err := provider.NewRealProviderStrategy(ctx, logger, c.Providers, c.KubernetesSupported)

	if err != nil {
		logger.Warnf("Initialized with error %v", err)
	}

	var usedStrategy provider.Strategy
	usedStrategy = providerStrategy
	if c.UseLocalDB {
		logger.Infof("Initializing local database")
		dbClient := db.NewBoltClient()
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

		strategyWithDB := db.NewStrategyWithDB(dbClient, providerStrategy, logger)
		usedStrategy = strategyWithDB

		syncer := sync.NewSyncer(logger, dbClient, providerStrategy, &c.SyncConfig)
		go syncer.SyncPeriodically(ctx)
	}

	s := &Server{
		logger:   logger,
		strategy: usedStrategy,
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

	go s.refreshClusters(ctx, time.Second*60)
	logger.Infof("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) refreshClusters(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			err := s.strategy.RefreshState(ctx)
			if err != nil {
				s.logger.Errorf("Failed to refresh state: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *Server) ListAccounts(ctx context.Context, in *infrapb.ListAccountsRequest) (*infrapb.ListAccountsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	accounts := cloudProvider.ListAccounts()
	if err != nil {
		return nil, err
	}
	return &infrapb.ListAccountsResponse{
		Accounts: typesAccountsToGrpc(accounts),
	}, nil
}

func (s *Server) ListRegions(ctx context.Context, in *infrapb.ListRegionsRequest) (*infrapb.ListRegionsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	regions, err := cloudProvider.ListRegions(ctx, in)
	if err != nil {
		return nil, err
	}
	return &infrapb.ListRegionsResponse{
		Regions: typesRegionsToGrpc(regions),
	}, nil
}

func (s *Server) ListVPC(ctx context.Context, in *infrapb.ListVPCRequest) (*infrapb.ListVPCResponse, error) {
	var errorMessage string
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	vpcs, err := cloudProvider.ListVPC(ctx, in)
	if err != nil {
		errorMessage = err.Error()
	}
	syncTime, e := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.VPCType))
	if e != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.VPCType, cloudProvider.GetName())
	}
	return &infrapb.ListVPCResponse{
		LastSyncTime: syncTime.Time,
		Vpcs:         typesVpcsToGrpc(vpcs),
		Err: &infrapb.Error{
			Code:         100,
			ErrorMessage: errorMessage,
			Serverity:    "Severe",
		},
	}, err
}

func (s *Server) ListInstances(ctx context.Context, in *infrapb.ListInstancesRequest) (*infrapb.ListInstancesResponse, error) {
	var errorMessage string

	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	instances, err := cloudProvider.ListInstances(ctx, in)
	if err != nil {
		errorMessage = err.Error()
		return nil, err
	}
	syncTime, e := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.InstanceType))
	if e != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.InstanceType, cloudProvider.GetName())
	}

	return &infrapb.ListInstancesResponse{
		LastSyncTime: syncTime.Time,
		Instances:    typesInstanceToGrpc(instances),
		Err: &infrapb.Error{
			Code:         100,
			ErrorMessage: errorMessage,
			Serverity:    "Severe",
		},
	}, err
}
func (s *Server) ListSubnets(ctx context.Context, in *infrapb.ListSubnetsRequest) (*infrapb.ListSubnetsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	subnets, err := cloudProvider.ListSubnets(ctx, in)
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.SubnetType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.SubnetType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}

	return &infrapb.ListSubnetsResponse{
		LastSyncTime: t,
		Subnets:      typesSubnetsToGrpc(subnets),
	}, nil
}

func (s *Server) ListACLs(ctx context.Context, in *infrapb.ListACLsRequest) (*infrapb.ListACLsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListACLs(ctx, in)
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.ACLType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.ACLType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListACLsResponse{
		LastSyncTime: t,
		Acls:         typesACLsToGrpc(l),
	}, nil
}

func (s *Server) ListSecurityGroups(ctx context.Context, in *infrapb.ListSecurityGroupsRequest) (*infrapb.ListSecurityGroupsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListSecurityGroups(ctx, in)
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.SecurityGroupType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.SecurityGroupType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListSecurityGroupsResponse{
		LastSyncTime:   t,
		SecurityGroups: typesSgsToGrpc(l),
	}, nil
}

func (s *Server) ListRouteTables(ctx context.Context, in *infrapb.ListRouteTablesRequest) (*infrapb.ListRouteTablesResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListRouteTables(ctx, in)
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.RouteTableType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.RouteTableType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListRouteTablesResponse{
		LastSyncTime: t,
		RouteTables:  typesRouteTableToGrpc(l),
	}, nil
}

func (s *Server) ListNATGateways(ctx context.Context, in *infrapb.ListNATGatewaysRequest) (*infrapb.ListNATGatewaysResponse, error) {

	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListNATGateways(ctx, in)
	if err != nil {
		s.logger.Errorf("Failure to retreive NAT GW %s", err.Error())
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.NATGatewayType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.NATGatewayType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListNATGatewaysResponse{
		LastSyncTime: t,
		NatGateways:  typesNATGatewaysToGrpc(l),
	}, nil
}

func (s *Server) ListRouters(ctx context.Context, in *infrapb.ListRoutersRequest) (*infrapb.ListRoutersResponse, error) {

	s.logger.Infof("Listing routers from user query")
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListRouters(ctx, in)
	if err != nil {
		s.logger.Errorf("Failure to retreive Router %s", err.Error())
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.RouterType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.RouterType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListRoutersResponse{
		LastSyncTime: t,
		Routers:      typesRoutersToGrpc(l),
	}, nil
}

func (s *Server) ListInternetGateways(ctx context.Context, in *infrapb.ListInternetGatewaysRequest) (*infrapb.ListInternetGatewaysResponse, error) {

	s.logger.Infof("Listing routers from user query")
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListInternetGateways(ctx, in)
	if err != nil {
		s.logger.Errorf("Failure to retreive Router %s", err.Error())
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.IGWType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.IGWType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListInternetGatewaysResponse{
		LastSyncTime: t,
		Igws:         typesIGWsToGrpc(l),
	}, nil
}

func (s *Server) ListVPCEndpoints(ctx context.Context, in *infrapb.ListVPCEndpointsRequest) (*infrapb.ListVPCEndpointsResponse, error) {

	s.logger.Infof("Listing routers from user query")
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListVPCEndpoints(ctx, in)
	if err != nil {
		s.logger.Errorf("Failure to retreive Router %s", err.Error())
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.VPCEndpointType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.VPCEndpointType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListVPCEndpointsResponse{
		LastSyncTime: t,
		Veps:         typesVPCEndpointsToGrpc(l),
	}, nil
}

// server/server.go
func (s *Server) ListPublicIPs(ctx context.Context, in *infrapb.ListPublicIPsRequest) (*infrapb.ListPublicIPsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	l, err := cloudProvider.ListPublicIPs(ctx, in)
	if err != nil {
		s.logger.Errorf("Failure to retreive Router %s", err.Error())
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.PublicIPType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.PublicIPType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListPublicIPsResponse{
		LastSyncTime: t,
		PublicIps:    typesPublicIPsToGrpc(l),
	}, nil
}

func (s *Server) GetVPCIDForCIDR(ctx context.Context, in *infrapb.GetVPCIDForCIDRRequest) (*infrapb.GetVPCIDForCIDRResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	vpcId, err := cloudProvider.GetVPCIDForCIDR(ctx, in)
	if err != nil {
		return nil, err
	}

	return &infrapb.GetVPCIDForCIDRResponse{
		VpcId: vpcId,
	}, nil
}

// GetCIDRsForLabels returns CIDRs of subnets with given labels.
func (s *Server) GetCIDRsForLabels(ctx context.Context, in *infrapb.GetCIDRsForLabelsRequest) (*infrapb.GetCIDRsForLabelsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	cidrs, err := cloudProvider.GetCIDRsForLabels(ctx, in)
	if err != nil {
		return nil, err
	}

	return &infrapb.GetCIDRsForLabelsResponse{
		Cidrs: cidrs,
	}, nil
}

// GetIPsForLabels returns IPs of instances with given labels.
func (s *Server) GetIPsForLabels(ctx context.Context, in *infrapb.GetIPsForLabelsRequest) (*infrapb.GetIPsForLabelsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	ips, err := cloudProvider.GetIPsForLabels(ctx, in)
	if err != nil {
		return nil, err
	}

	return &infrapb.GetIPsForLabelsResponse{
		Ips: ips,
	}, nil
}

func (s *Server) GetVPCIDWithTag(ctx context.Context, in *infrapb.GetVPCIDWithTagRequest) (*infrapb.GetVPCIDWithTagResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	vpcId, err := cloudProvider.GetVPCIDWithTag(ctx, in)
	if err != nil {
		return nil, err
	}

	return &infrapb.GetVPCIDWithTagResponse{
		VpcId: vpcId,
	}, nil
}

func (s *Server) AddInboundAllowRuleInVPC(ctx context.Context, in *infrapb.AddInboundAllowRuleInVPCRequest) (*infrapb.AddInboundAllowRuleInVPCResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	err = cloudProvider.AddInboundAllowRuleInVPC(ctx, in.AccountId, in.Region, in.DestinationVpcId, in.CidrsToAllow, in.RuleName, in.Tags)
	return &infrapb.AddInboundAllowRuleInVPCResponse{}, err
}

func (s *Server) AddInboundAllowRuleByLabelsMatch(ctx context.Context, in *infrapb.AddInboundAllowRuleByLabelsMatchRequest) (*infrapb.AddInboundAllowRuleByLabelsMatchResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	ruleId, instances, err := cloudProvider.AddInboundAllowRuleByLabelsMatch(ctx, in.AccountId, in.Region, in.VpcId, in.RuleName, in.Labels, in.CidrsToAllow, grpcProtocolsAndPortToTypes(in.ProtocolsAndPorts))
	if err != nil {
		return nil, err
	}
	return &infrapb.AddInboundAllowRuleByLabelsMatchResponse{
		RuleId:    ruleId,
		Instances: typesInstanceToGrpc(instances),
	}, nil
}

func (s *Server) AddInboundAllowRuleBySubnetMatch(ctx context.Context, in *infrapb.AddInboundAllowRuleBySubnetMatchRequest) (*infrapb.AddInboundAllowRuleBySubnetMatchResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	ruleId, instances, subnets, err := cloudProvider.AddInboundAllowRuleBySubnetMatch(ctx, in.AccountId, in.Region, in.VpcId, in.RuleName, in.SubnetCidrs, in.CidrsToAllow, grpcProtocolsAndPortToTypes(in.ProtocolsAndPorts))
	if err != nil {
		return nil, err
	}
	return &infrapb.AddInboundAllowRuleBySubnetMatchResponse{
		RuleId:    ruleId,
		Instances: typesInstanceToGrpc(instances),
		Subnets:   typesSubnetsToGrpc(subnets),
	}, nil
}

func (s *Server) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, in *infrapb.AddInboundAllowRuleByInstanceIPMatchRequest) (*infrapb.AddInboundAllowRuleByInstanceIPMatchResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	ruleId, instances, err := cloudProvider.AddInboundAllowRuleByInstanceIPMatch(ctx, in.AccountId, in.Region, in.VpcId, in.RuleName, in.InstancesIps, in.CidrsToAllow, grpcProtocolsAndPortToTypes(in.ProtocolsAndPorts))
	if err != nil {
		return nil, err
	}
	return &infrapb.AddInboundAllowRuleByInstanceIPMatchResponse{
		RuleId:    ruleId,
		Instances: typesInstanceToGrpc(instances),
	}, nil
}

func (s *Server) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, in *infrapb.AddInboundAllowRuleForLoadBalancerByDNSRequest) (*infrapb.AddInboundAllowRuleForLoadBalancerByDNSResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	loadBalancerIds, ruleId, err := cloudProvider.AddInboundAllowRuleForLoadBalancerByDNS(ctx, in.AccountId, in.Region, in.LoadBalancerDns, in.VpcId, in.RuleName, in.CidrsToAllow, grpcProtocolsAndPortToTypes(in.ProtocolsAndPorts))
	if err != nil {
		return nil, err
	}
	return &infrapb.AddInboundAllowRuleForLoadBalancerByDNSResponse{
		LoadBalancerId: loadBalancerIds,
		RuleId:         ruleId,
	}, nil
}

func (s *Server) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, in *infrapb.RemoveInboundAllowRuleFromVPCByNameRequest) (*infrapb.RemoveInboundAllowRuleFromVPCByNameResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	err = cloudProvider.RemoveInboundAllowRuleFromVPCByName(ctx, in.AccountId, in.Region, in.VpcId, in.RuleName)
	if err != nil {
		return nil, err
	}
	return &infrapb.RemoveInboundAllowRuleFromVPCByNameResponse{}, nil
}

func (s *Server) RemoveInboundAllowRulesFromVPCById(ctx context.Context, in *infrapb.RemoveInboundAllowRulesFromVPCByIdRequest) (*infrapb.RemoveInboundAllowRulesFromVPCByIdResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	err = cloudProvider.RemoveInboundAllowRulesFromVPCById(ctx, in.AccountId, in.Region, in.VpcId, in.InstanceIds, in.LoadBalancerIds, in.RuleId)
	if err != nil {
		return nil, err
	}
	return &infrapb.RemoveInboundAllowRulesFromVPCByIdResponse{}, nil
}

func (s *Server) RemoveInboundAllowRuleRulesByTags(ctx context.Context, in *infrapb.RemoveInboundAllowRuleRulesByTagsRequest) (*infrapb.RemoveInboundAllowRuleRulesByTagsResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	err = cloudProvider.RemoveInboundAllowRuleRulesByTags(ctx, in.AccountId, in.Region, in.VpcId, in.RuleName, in.Tags)
	if err != nil {
		return nil, err
	}
	return &infrapb.RemoveInboundAllowRuleRulesByTagsResponse{}, err
}

func (s *Server) RefreshInboundAllowRule(ctx context.Context, in *infrapb.RefreshInboundAllowRuleRequest) (*infrapb.RefreshInboundAllowRuleResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	instances, subnets, err := cloudProvider.RefreshInboundAllowRule(ctx, in.AccountId, in.Region, in.RuleId, in.CidrsToAdd, in.CidrsToRemove, in.DestinationLabels, in.DestinationPrefixes, in.DestinationVpcId, grpcProtocolsAndPortToTypes(in.ProtocolsAndPorts))
	if err != nil {
		return nil, err
	}
	return &infrapb.RefreshInboundAllowRuleResponse{
		Instances: typesInstanceToGrpc(instances),
		Subnets:   typesSubnetsToGrpc(subnets),
	}, nil
}

func (s *Server) ListCloudClusters(ctx context.Context, in *infrapb.ListCloudClustersRequest) (*infrapb.ListCloudClustersResponse, error) {
	cloudProvider, err := s.strategy.GetProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	clusters, err := cloudProvider.ListClusters(ctx, in)
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := cloudProvider.GetSyncTime(types.SyncTimeKey(cloudProvider.GetName(), types.ClusterType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, provider %s", types.ClusterType, cloudProvider.GetName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListCloudClustersResponse{
		LastSyncTime: t,
		Clusters:     typesClustersToGrpc(clusters),
	}, nil
}

func (s *Server) ListClusters(ctx context.Context, in *infrapb.ListClustersRequest) (*infrapb.ListClustersResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	clusters, err := k8sProvider.ListClusters(ctx)
	if err != nil {
		return nil, err
	}
	return &infrapb.ListClustersResponse{Clusters: typesClustersToGrpc(clusters)}, nil
}

func (s *Server) ListPods(ctx context.Context, in *infrapb.ListPodsRequest) (*infrapb.ListPodsResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	pods, err := k8sProvider.ListPods(ctx, in.GetClusterName(), in.GetLabels())
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := k8sProvider.GetSyncTime(types.SyncTimeKey(in.GetClusterName(), types.PodsType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, cluster %s", types.PodsType, in.GetClusterName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListPodsResponse{
		Pods:         typesPodsToGrpc(pods),
		LastSyncTime: t,
	}, nil
}
func (s *Server) ListServices(ctx context.Context, in *infrapb.ListServicesRequest) (*infrapb.ListServicesResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	services, err := k8sProvider.ListServices(ctx, in.GetClusterName(), in.GetLabels())
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := k8sProvider.GetSyncTime(types.SyncTimeKey(in.GetClusterName(), types.K8sServiceType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, cluster %s", types.K8sServiceType, in.GetClusterName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListServicesResponse{
		Services:     typesServicesToGrpc(services),
		LastSyncTime: t,
	}, nil
}

func (s *Server) ListNodes(ctx context.Context, in *infrapb.ListNodesRequest) (*infrapb.ListNodesResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	nodes, err := k8sProvider.ListNodes(ctx, in.GetClusterName(), in.GetLabels())
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := k8sProvider.GetSyncTime(types.SyncTimeKey(in.GetClusterName(), types.K8sNodeType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, cluster %s", types.K8sNodeType, in.GetClusterName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListNodesResponse{
		Nodes:        typesNodesToGrpc(nodes),
		LastSyncTime: t,
	}, nil
}

func (s *Server) ListNamespaces(ctx context.Context, in *infrapb.ListNamespacesRequest) (*infrapb.ListNamespacesResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	namespaces, err := k8sProvider.ListNamespaces(ctx, in.GetClusterName(), in.GetLabels())
	if err != nil {
		return nil, err
	}
	var t string
	syncTime, err := k8sProvider.GetSyncTime(types.SyncTimeKey(in.GetClusterName(), types.NamespaceType))
	if err != nil {
		s.logger.Errorf("Failed to get sync time for %s, cluster %s", types.NamespaceType, in.GetClusterName())
	} else {
		t = syncTime.Time
	}
	return &infrapb.ListNamespacesResponse{
		Namespaces:   typesNamespacesToGrpc(namespaces),
		LastSyncTime: t,
	}, nil
}

func (s *Server) ListPodsCIDRs(ctx context.Context, in *infrapb.ListPodsCIDRsRequest) (*infrapb.ListPodsCIDRsResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	cidrs, err := k8sProvider.ListPodsCIDRs(ctx, in.GetClusterName())
	if err != nil {
		return nil, err
	}
	return &infrapb.ListPodsCIDRsResponse{Cidrs: cidrs}, nil
}

func (s *Server) ListServicesCIDRs(ctx context.Context, in *infrapb.ListServicesCIDRsRequest) (*infrapb.ListServicesCIDRsResponse, error) {
	k8sProvider, err := s.strategy.GetKubernetesProvider()
	if err != nil {
		return nil, err
	}
	cidrs, err := k8sProvider.ListServicesCIDRs(ctx, in.GetClusterName())
	if err != nil {
		return nil, err
	}
	return &infrapb.ListServicesCIDRsResponse{Cidr: cidrs}, nil
}

// unaryServerInterceptor logs the details of the unary RPC calls
func (s *Server) unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		s.logger.Infof("Unary Request - Method:%s, Peer:%s\n", info.FullMethod, p.Addr)
	} else {
		s.logger.Infof("Unary Request - Method:%s\n", info.FullMethod)
	}

	// Call the handler to complete the RPC
	resp, err := handler(ctx, req)

	// Log response
	s.logger.Infof("Request = %+v", req)
	s.logger.Debugf("Unary Response - Method:%s, Response:%v, Error:%v\n", info.FullMethod, resp, err)

	return resp, err
}
