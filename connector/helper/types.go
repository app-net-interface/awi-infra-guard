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

package helper

import "strconv"

// Returns dereferenced string - if the pointer
// is a nil, returns empty string.
func StringPointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// If string is empty, returns nil, otherwise
// returns a pointer to that particular string.
func StringToStringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// If string is empty, returns nil, otherwise
// returns a pointer to uint32 variable casted
// from a given string.
func StringToUInt32Pointer(value string) *uint32 {
	// TODO: Consider handling integer overload.
	if value == "" {
		return nil
	}
	converted, err := strconv.Atoi(value)
	if err != nil {
		// Consider returning error here.
		return nil
	}
	converted32 := uint32(converted)
	return &converted32
}

// If the value is nil, returns empty string,
// otherwise converts the uint32 to a string
// number.
func UInt32PointerToString(value *uint32) string {
	if value == nil {
		return ""
	}
	return strconv.FormatUint(uint64(*value), 10)
}
