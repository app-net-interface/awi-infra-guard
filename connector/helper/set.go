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

package helper

type Set[T comparable] struct {
	values map[T]struct{}
}

func SetFromSlice[T comparable](s []T) Set[T] {
	v := make(map[T]struct{})
	for i := range s {
		v[s[i]] = struct{}{}
	}
	return Set[T]{values: v}
}

func (s *Set[T]) Set(v T) {
	s.values[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.values[v]
	return ok
}
