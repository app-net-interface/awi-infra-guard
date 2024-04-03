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

package secret

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GeneratePSK generates a pre-shared key for VPN tunnels.
//
// TODO: GeneratePSK generates keys that sometimes don't
// meet AWS criterias - inspect that.
func GeneratePSK() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate PreSharedKey: %v", err)
	}
	return hex.EncodeToString(bytes), nil
}
