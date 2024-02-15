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

package types

// The abstraction over Cloud edges that can be used for attaching
// a connection.
type Gateway struct {
	// The Name of an actual resource used beneath the abstraction.
	// For example for AWS it would be the ID of given TransitGateway.
	Name string
	// The actual resource kind used by a given Gateway. For an AWS it
	// may be a TransitGateway.
	Kind string
	// The CSP owner of that particular Gateway.
	// AWS or GCP.
	CloudProvider string
	// The name of VPC/Network the Gateway is attached to.
	VPC string
	// The ASN Number for the Gateway.
	ASN string
	// The Region where resources were created.
	Region string
	// Number of Connections using that particular Gateway.
	ConnectionsCount uint64
	// The limit of Connections that can use this particular Gateway.
	// 0 indicates no limit.
	ConnectionsLimit uint64
	// Additional information available per provider.
	Data map[string]string
}
