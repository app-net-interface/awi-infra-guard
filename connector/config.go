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
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// The Configurations for providers to be used.
	Providers map[string]interface{}
	// The configuration for Logging System in CSP
	// Connector.
	Logger LoggerConfig
}

type LogLevel string

const (
	LogLevelTrace    LogLevel = "TRACE"
	LogLevelDebug    LogLevel = "DEBUG"
	LogLevelInfo     LogLevel = "INFO"
	LogLevelWarn     LogLevel = "WARN"
	LogLevelError    LogLevel = "ERROR"
	LogLevelCritical LogLevel = "CRITICAL"
	LogLevelNone     LogLevel = "NONE"
)

type LoggerConfig struct {
	// The global Log Level inherited by every
	// component which has no explicit level configured.
	Level LogLevel
	// The map of Log Levels per given Component.
	// The Component name is picked as follows:
	// <PROVIDER NAME>:<LOGGER NAME>
	// So, for example for AWS Provider with Client
	// component, the name will be as follows
	ComponentLevels map[string]LogLevel
	// Configuring where the logs should be presented
	// or directed.
	Output LoggerOutputConfig
}

type LoggerOutputConfig struct {
	// If set to true, the Logger will print logs to
	// standard output.
	StdOut bool
	// The FilePath of an output logger file.
	//
	// If the FilePath is empty, there will be no
	// attempt to log to any file.
	//
	// Invalid FilePath will result in an error of the
	// CLI.
	FilePath string
}

func LoadConfig(filepath string) (Config, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading YAML file: %w", err)
	}

	config := Config{}
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return Config{}, fmt.Errorf("Error parsing YAML file: %w", err)
	}

	for key, value := range config.Providers {
		marshalled, err := yaml.Marshal(value)
		if err != nil {
			return Config{}, fmt.Errorf(
				"Error parsing configuration for provider %s: %w", key, err,
			)
		}
		config.Providers[key] = string(marshalled)
	}

	return config, nil
}
