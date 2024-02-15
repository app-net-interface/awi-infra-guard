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

package gcp

const (
	REQUESTED_CIDR_SIZE = 30
)

type Config struct {
	// The GCP Project ID, which will be used for interactions with
	// Google Cloud.
	//
	// If empty, the GCP Client will check the list of projects, which
	// the user has access to - if there is only one project, the script
	// will use that project automatically, otherwise, the user will be
	// instructed to specify the exact project.
	Project string
	// The Region, where GCP resources such as Routers and Subnets
	// will be created.
	Region string
}
