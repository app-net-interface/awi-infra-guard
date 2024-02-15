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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type filterBuilder struct {
	filters []types.Filter
}

func newFilterBuilder() *filterBuilder {
	return &filterBuilder{}
}

func (a *filterBuilder) build() []types.Filter {
	return a.filters
}

func (a *filterBuilder) withVPC(vpcID string) {
	a.addToFilters("vpc-id", vpcID)
}

func (a *filterBuilder) withTag(key, value string) {
	a.addToFilters("tag:"+key, value)
}

func (a *filterBuilder) withTagManyValues(key string, values []string) {
	a.addToFiltersWithManyValues("tag:"+key, values)
}

func (a *filterBuilder) withAvailabilityZone(zone string) {
	a.addToFilters("availability-zone", zone)
}

func (a *filterBuilder) withCIDR(cidr string) {
	a.addToFilters("cidr-block", cidr)
}

func (a *filterBuilder) withSubnet(subnetID string) {
	a.addToFilters("subnet-id", subnetID)
}

func (a *filterBuilder) withSecurityGroupName(securityGroupName string) {
	a.addToFilters("group-name", securityGroupName)
}

func (a *filterBuilder) withSecurityGroupId(securityGroupId string) {
	a.addToFilters("group-id", securityGroupId)
}

func (a *filterBuilder) withPrivateIPAddress(privateIPAddress string) {
	a.addToFilters("private-ip-address", privateIPAddress)
}

func (a *filterBuilder) withAssociationMain() {
	a.addToFilters("association.main", "true")
}

func (a *filterBuilder) addToFilters(name, value string) {
	if value != "" {
		filter := types.Filter{
			Name:   aws.String(name),
			Values: []string{value},
		}
		a.filters = append(a.filters, filter)
	}
}

func (a *filterBuilder) addToFiltersWithManyValues(name string, values []string) {
	if len(values) > 0 {
		filter := types.Filter{
			Name:   aws.String(name),
			Values: values,
		}
		a.filters = append(a.filters, filter)
	}
}
