// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	// CuratorConfig contains configuration for curator.
	CuratorConfig struct {
		Client ClientConfig `yaml:"client"`
	}

	// ClientConfig contains configuration for Elasticsearch client.
	ClientConfig struct {
		Hosts    []string `yaml:"hosts"`
		Port     int32    `yaml:"port"`
		HTTPAuth string   `yaml:"http_auth"`
	}
)

// ReadConfig reads the yaml config file located at <configPath>
// and unmarshal the file content into CuratorConfig struct.
func ReadConfig(configPath string) (*CuratorConfig, error) {
	if configPath == "" {
		return nil, errors.New("No configuration path passed")
	}

	// Read the configuration file.
	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the configuration file into config struct.
	config := CuratorConfig{
		Client: ClientConfig{
			Hosts:    []string{"localhost"},
			Port:     9200,
			HTTPAuth: "",
		},
	}
	err = yaml.Unmarshal(configRaw, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
