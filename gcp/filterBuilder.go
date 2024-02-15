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

import (
	"fmt"
	"strings"
)

type filterBuilder struct {
	filter string
}

func newFilterBuilder() filterBuilder { return filterBuilder{} }

func (f *filterBuilder) build() *string {
	if f == nil || f.filter == "" {
		return nil
	}
	return &f.filter
}

func (f *filterBuilder) addFilter(val string) {
	if f.filter != "" {
		f.filter = f.filter + fmt.Sprintf(" ")
	}
	f.filter = f.filter + val
}

func (f *filterBuilder) addLabelFilter(val, condition string) {
	if f.filter != "" {
		f.filter = f.filter + fmt.Sprintf(" %s ", condition)
	}
	f.filter = f.filter + val
}

func (f *filterBuilder) withLabel(labelKey, labelValue, condition string) {
	f.addLabelFilter(fmt.Sprintf("(labels.%s=%s)", labelKey, labelValue), condition)
}

func (f *filterBuilder) withIPCIDRRange(cidrRange string) {
	f.addFilter(fmt.Sprintf("(ipCidrRange eq %s)", cidrRange))
}

func (f *filterBuilder) withRegion(region string) {
	f.addFilter(fmt.Sprintf("(region eq .*/%s$)", region))
}

func (f *filterBuilder) withNetwork(network string) {
	if strings.HasPrefix(network, "https://www.googleapis.com/compute/v1/") {
		f.addFilter(fmt.Sprintf("(network eq %s)", network))
	} else {
		f.addFilter(fmt.Sprintf("(network eq .*/%s$)", network))
	}
}
