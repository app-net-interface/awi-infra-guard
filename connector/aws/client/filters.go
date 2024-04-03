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

package client

import (
	"github.com/app-net-interface/awi-infra-guard/connector/helper"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListCustomerGatewayFilter interface {
	GetCustomerGatewayListFilter() types.Filter
}

type ListCustomerGatewayFilterIP struct {
	Value string
}

func (f ListCustomerGatewayFilterIP) GetCustomerGatewayListFilter() types.Filter {
	return types.Filter{
		Name:   helper.StringToStringPointer("ip-address"),
		Values: []string{f.Value},
	}
}

type ListVPNConnectionFilter interface {
	GetVPNConnectionListFilter() types.Filter
}

type ListVPNConnectionFilterCustomerGatewayID struct {
	Value string
}

func (f ListVPNConnectionFilterCustomerGatewayID) GetVPNConnectionListFilter() types.Filter {
	return types.Filter{
		Name:   helper.StringToStringPointer("customer-gateway-id"),
		Values: []string{f.Value},
	}
}

type ListVPNConnectionFilterTransitGatewayID struct {
	Value string
}

func (f ListVPNConnectionFilterTransitGatewayID) GetVPNConnectionListFilter() types.Filter {
	return types.Filter{
		Name:   helper.StringToStringPointer("transit-gateway-id"),
		Values: []string{f.Value},
	}
}
