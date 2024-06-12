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
package config

import (
	"strings"
	"time"
)

type Provider struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
}

type Resources struct {
	Cloud      []string `mapstructure:"cloud"`
	Kubernetes []string `mapstructure:"kubernetes"`
}

type SyncConfig struct {
	Enabled      bool          `mapstructure:"enabled"`
	SyncWaitTime time.Duration `mapstructure:"syncWaitTime"`
	DbFileName   string        `mapstructure:"dbFileName"`
	Resources    Resources     `mapstructure:"resources"`
}

type Config struct {
	Port                string     `mapstructure:"port"`
	Hostname            string     `mapstructure:"hostname"`
	UseLocalDB          bool       `mapstructure:"useLocalDB"`
	LogLevel            string     `mapstructure:"logLevel"`
	Providers           []Provider `mapstructure:"providers"`
	KubernetesSupported bool       `mapstructure:"kubernetesSupported"`
	SyncConfig          SyncConfig `mapstructure:"syncConfig"`
}

// Method to check if a specific cloud resource is present
func (sc *SyncConfig) HasCloudResource(resource string) bool {
	for _, res := range sc.Resources.Cloud {
		if strings.ToLower(res) == resource {
			return true
		}
	}
	return false
}

// Method to check if a specific Kubernetes resource is present
func (sc *SyncConfig) HasKubernetesResource(resource string) bool {
	for _, res := range sc.Resources.Kubernetes {
		if strings.ToLower(res) == resource {
			return true
		}
	}
	return false
}
