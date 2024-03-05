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

package aws

import (
	"github.com/app-net-interface/awi-infra-guard/connector/aws/client"
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

func (c *AWSConnector) generateTagName() string {
	return helper.CreateName(
		// TODO: Enforce one order of connection sides.
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName, "",
	)
}

func (c *AWSConnector) shouldResourceBeDeleted(tags []client.Tag) bool {
	shouldDelete := false
	for _, tag := range tags {
		if c.isTagOwnedByAnotherConnection(tag.Key) {
			return false
		}
		if !shouldDelete && c.isTagOwnedByConnection(tag.Key) {
			shouldDelete = true
		}
	}
	return shouldDelete
}

func (c *AWSConnector) isResourceOwnedByDifferentConnection(tags []client.Tag) bool {
	for _, tag := range tags {
		if c.isTagOwnedByAnotherConnection(tag.Key) {
			return true
		}
	}
	return false
}

func (c *AWSConnector) isResourceOwnedByConnection(tags []client.Tag) bool {
	for _, tag := range tags {
		if c.isTagOwnedByConnection(tag.Key) {
			return true
		}
	}
	return false
}

func (c *AWSConnector) isTagOwnedByConnection(tagName string) bool {
	return helper.NameCreatedForIdentifier(
		tagName,
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName,
	)
}

func (c *AWSConnector) isTagOwnedByAnotherConnection(tagName string) bool {
	return helper.NameCreatedByScript(tagName) && !helper.NameCreatedForIdentifier(
		tagName,
		c.transactionState.GatewayName+"/"+c.transactionState.PeerGatewayName,
	)
}
