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
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/aws"
	"github.com/app-net-interface/awi-infra-guard/connector/azure"
	"github.com/app-net-interface/awi-infra-guard/connector/gcp"
	"github.com/app-net-interface/awi-infra-guard/connector/provider"
	"github.com/sirupsen/logrus"
)

type ProviderManager struct {
	availableProviders   map[string]Initializer
	initializedProviders map[string]provider.Provider
	logger               *logrus.Entry
}

type Initializer func(ctx context.Context, logger *logrus.Entry, config string) (provider.Provider, error)

func NewProviderManager(logger *logrus.Entry) (ProviderManager, error) {
	manager := ProviderManager{
		availableProviders:   map[string]Initializer{},
		initializedProviders: map[string]provider.Provider{},
		logger:               logger,
	}
	if err := manager.loadDefaultProviders(); err != nil {
		return ProviderManager{}, fmt.Errorf(
			"failed to load the list of available default providers: %w",
			err,
		)
	}
	if err := manager.loadDynamicProviders(); err != nil {
		return ProviderManager{}, fmt.Errorf(
			"failed to load the list of available dynamic providers: %w",
			err,
		)
	}
	return manager, nil
}

func (p *ProviderManager) allProvidersInitialized() bool {
	if p == nil {
		return false
	}
	return len(p.availableProviders) == len(p.initializedProviders)
}

func (p *ProviderManager) loadDefaultProviders() error {
	p.availableProviders["aws"] = aws.NewConnector
	p.availableProviders["gcp"] = gcp.NewConnector
	p.availableProviders["azure"] = azure.NewConnector
	return nil
}

func (p *ProviderManager) loadDynamicProviders() error {
	// NOT IMPLEMENTED
	// This is a placeholder for dynamic plugin Providers.
	return nil
}

func (p *ProviderManager) InitializeProvider(
	ctx context.Context,
	logger *logrus.Entry,
	providerName string,
	config interface{},
) (provider.Provider, error) {
	if p == nil {
		return nil, fmt.Errorf("cannot initialize provider '%s' as ProviderManager is nil", providerName)
	}
	p.logger.Tracef("Initializing provider '%s'", providerName)
	if provider, ok := p.initializedProviders[providerName]; ok {
		p.logger.Warnf(
			"provider '%s' was already initialized. Skipping additional initialization",
			provider.Name(),
		)
		return provider, nil
	}
	providerInitialize, ok := p.availableProviders[providerName]
	if !ok {
		return nil, fmt.Errorf("cannot initialize provider '%s'. No implementation found", providerName)
	}
	provider, err := providerInitialize(ctx, logger.WithField("logger", providerName), config.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize provider '%s': %w", providerName, err)
	}
	p.initializedProviders[providerName] = provider
	p.logger.Tracef("Provider '%s' initialized successfully", providerName)
	return provider, nil
}

// InitializeAvailableProviders initializes all providers
// which implementation was found or provided to the connector.
//
// This function doesn't break if any provider fails to initialize
// but will log the error.
func (p *ProviderManager) InitializeAvailableProviders(
	ctx context.Context,
	logger *logrus.Entry,
	configs map[string]interface{},
) {
	if p == nil {
		return
	}
	if p.allProvidersInitialized() {
		p.logger.Debug(
			"providers were already initialized. Skipping initialization")
		return
	}
	p.logger.Tracef("Initializing all available providers")
	for name := range p.availableProviders {
		providerConfig := ""
		if c, ok := configs[name]; ok {
			providerConfig = c.(string)
		} else {
			p.logger.Infof("no configuration found for provider '%s'. Using default", name)
		}
		if _, err := p.InitializeProvider(ctx, logger, name, providerConfig); err != nil {
			p.logger.Errorf("cannot initialize provider '%s': %v", name, err)
		}
	}
}

func (p *ProviderManager) Close() error {
	if p == nil {
		return nil
	}
	for name, provider := range p.initializedProviders {
		if err := provider.Close(); err != nil {
			return fmt.Errorf("failed to close '%s' provider: %w", name, err)
		}
	}
	return nil
}

func (p *ProviderManager) Provider(name string) (provider.Provider, error) {
	if p == nil {
		return nil, fmt.Errorf("cannot get provider '%s' as ProviderManager is nil", name)
	}
	provider, ok := p.initializedProviders[name]
	if !ok {
		return nil, fmt.Errorf("cannot get provider '%s' as it was not initialized", name)
	}
	return provider, nil
}

func (p *ProviderManager) Providers() []provider.Provider {
	if p == nil {
		return nil
	}
	providers := []provider.Provider{}
	for _, provider := range p.initializedProviders {
		providers = append(providers, provider)
	}
	return providers
}
