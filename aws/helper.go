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
	"fmt"
	"reflect"
	"strings"
)

func convertString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}

// ExtractAccountID extracts the account ID from an AWS ARN string
func ExtractAccountID(arn string) string {
    parts := strings.Split(arn, ":")
    
    if len(parts) >= 5 {
        return parts[4]
    }
    return ""
}

// PrintResources prints the details of any slice of AWS resources
func PrintResources(resources interface{}, resourceType string) {
    value := reflect.ValueOf(resources)
    if value.Kind() != reflect.Slice {
        fmt.Printf("Error: Expected a slice of %s, got %T\n", resourceType, resources)
        return
    }

    fmt.Printf("=== %s ===\n", resourceType)
    for i := 0; i < value.Len(); i++ {
        resource := value.Index(i)
        fmt.Printf("Resource %d:\n", i+1)
        printFields(resource)
        fmt.Println()
    }
    fmt.Printf("Total %s: %d\n", resourceType, value.Len())
}

func printFields(v reflect.Value) {
    t := v.Type()
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldName := t.Field(i).Name
        fieldValue := field.Interface()
        fmt.Printf("  %s: %v\n", fieldName, fieldValue)
    }
}
