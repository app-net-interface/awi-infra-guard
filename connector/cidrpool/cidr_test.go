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

package cidrpool

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func Test_invert32Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "nil slice should return an error",
			bytes:   nil,
			wantErr: true,
		},
		{
			name:    "too short slice of bytes should return an error",
			bytes:   []byte{255, 255, 255},
			wantErr: true,
		},
		{
			name:    "too long slice of bytes should return an error",
			bytes:   []byte{255, 255, 255, 255, 255},
			wantErr: true,
		},
		{
			name:    "byte in form of IPv6 should return an error",
			bytes:   []byte{255, 255, 255, 255, 255, 255},
			wantErr: true,
		},
		{
			name:  "test reversing regular IP address",
			bytes: []byte{192, 168, 0, 1},
			want:  []byte{63, 87, 255, 254},
		},
		{
			name:  "test reversing popular mask",
			bytes: []byte{255, 255, 255, 0},
			want:  []byte{0, 0, 0, 255},
		},
		{
			name:  "test 0 address",
			bytes: []byte{0, 0, 0, 0},
			want:  []byte{255, 255, 255, 255},
		},
		{
			name:  "test full address",
			bytes: []byte{255, 255, 255, 255},
			want:  []byte{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := invert32Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("invert32Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invert32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increaseIPBy32Bytes(t *testing.T) {
	tests := []struct {
		ip         net.IP
		bytesToAdd []byte
		want       net.IP
		wantErr    bool
	}{
		{
			ip:         nil,
			bytesToAdd: nil,
			wantErr:    true,
		},
		{
			ip:         []byte{10, 0, 0, 0},
			bytesToAdd: nil,
			wantErr:    true,
		},
		{
			ip:         nil,
			bytesToAdd: []byte{10, 0, 0, 0},
			wantErr:    true,
		},
		{
			ip:         []byte{10, 0, 0},
			bytesToAdd: []byte{10, 0, 0, 0},
			wantErr:    true,
		},
		{
			ip:         []byte{10, 0, 0, 0},
			bytesToAdd: []byte{10, 0, 0},
			wantErr:    true,
		},
		{
			ip:         []byte{10, 0, 0, 0},
			bytesToAdd: []byte{10, 0, 0, 0},
			want:       []byte{20, 0, 0, 0},
		},
		{
			ip:         []byte{0, 0, 0, 0},
			bytesToAdd: []byte{0, 0, 0, 0},
			want:       []byte{0, 0, 0, 0},
		},
		{
			ip:         []byte{1, 2, 3, 4},
			bytesToAdd: []byte{0, 0, 0, 0},
			want:       []byte{1, 2, 3, 4},
		},
		{
			ip:         []byte{0, 0, 0, 0},
			bytesToAdd: []byte{1, 2, 3, 4},
			want:       []byte{1, 2, 3, 4},
		},
		{
			ip:         []byte{0, 0, 0, 255},
			bytesToAdd: []byte{0, 0, 0, 1},
			want:       []byte{0, 0, 1, 0},
		},
		{
			ip:         []byte{0, 0, 0, 1},
			bytesToAdd: []byte{0, 0, 0, 255},
			want:       []byte{0, 0, 1, 0},
		},
		{
			ip:         []byte{0, 255, 255, 255},
			bytesToAdd: []byte{0, 0, 0, 1},
			want:       []byte{1, 0, 0, 0},
		},
		{
			ip:         []byte{0, 0, 0, 1},
			bytesToAdd: []byte{0, 255, 255, 255},
			want:       []byte{1, 0, 0, 0},
		},
		{
			ip:         []byte{255, 255, 255, 255},
			bytesToAdd: []byte{0, 0, 0, 1},
			want:       []byte{0, 0, 0, 0},
		},
		{
			ip:         []byte{255, 255, 255, 255},
			bytesToAdd: []byte{255, 255, 255, 255},
			want:       []byte{255, 255, 255, 254},
		},
		{
			ip:         []byte{192, 168, 3, 0},
			bytesToAdd: []byte{70, 32, 255, 16},
			want:       []byte{6, 201, 2, 16},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := increaseIPBy32Bytes(tt.ip, tt.bytesToAdd)
			if (err != nil) != tt.wantErr {
				t.Errorf("increaseIPBy32Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increaseIPBy32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseCIDR(cidr string) net.IPNet {
	_, c, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(fmt.Sprintf("Got error while parsing CIDR: %v", err))
	}
	return *c
}

func Test_nextCIDR(t *testing.T) {
	tests := []struct {
		cidr    net.IPNet
		want    net.IPNet
		wantErr bool
	}{
		{
			cidr: net.IPNet{
				IP:   nil,
				Mask: nil,
			},
			wantErr: true,
		},
		{
			cidr: net.IPNet{
				IP:   []byte{192, 168, 1, 1},
				Mask: nil,
			},
			wantErr: true,
		},
		{
			cidr: net.IPNet{
				IP:   nil,
				Mask: []byte{255, 255, 255, 0},
			},
			wantErr: true,
		},
		{
			cidr: parseCIDR("192.168.0.0/24"),
			want: parseCIDR("192.168.1.0/24"),
		},
		{
			cidr: parseCIDR("192.168.0.0/0"),
			want: parseCIDR("193.168.0.0/0"),
		},
		{
			cidr: parseCIDR("192.168.0.0/28"),
			want: parseCIDR("192.168.0.16/28"),
		},
		{
			cidr: parseCIDR("192.168.0.16/28"),
			want: parseCIDR("192.168.0.32/28"),
		},
		{
			cidr: parseCIDR("192.168.0.0/32"),
			want: parseCIDR("192.168.0.1/32"),
		},
		{
			cidr: parseCIDR("192.168.255.0/24"),
			want: parseCIDR("192.169.0.0/24"),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := nextCIDR(tt.cidr)
			if (err != nil) != tt.wantErr {
				t.Errorf("nextCIDR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareIPs(t *testing.T) {
	tests := []struct {
		ip1  net.IP
		ip2  net.IP
		want comparisonResult
	}{
		{
			ip1:  net.IPv4(192, 168, 0, 1),
			ip2:  net.IPv4(192, 168, 0, 0),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(192, 168, 0, 0),
			ip2:  net.IPv4(192, 168, 0, 1),
			want: LesserThan,
		},
		{
			ip1:  net.IPv4(192, 168, 0, 1),
			ip2:  net.IPv4(192, 168, 0, 1),
			want: Equal,
		},
		{
			ip1:  net.IPv4(1, 2, 3, 4),
			ip2:  net.IPv4(4, 3, 2, 1),
			want: LesserThan,
		},
		{
			ip1:  net.IPv4(0, 0, 0, 1),
			ip2:  net.IPv4(0, 0, 0, 0),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(0, 0, 1, 0),
			ip2:  net.IPv4(0, 0, 0, 0),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(0, 0, 1, 0),
			ip2:  net.IPv4(0, 0, 0, 255),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(0, 1, 0, 0),
			ip2:  net.IPv4(0, 0, 0, 0),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(0, 1, 0, 0),
			ip2:  net.IPv4(0, 0, 255, 255),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(1, 0, 0, 0),
			ip2:  net.IPv4(0, 0, 0, 0),
			want: GreaterThan,
		},
		{
			ip1:  net.IPv4(1, 0, 0, 0),
			ip2:  net.IPv4(0, 255, 255, 255),
			want: GreaterThan,
		},
		{
			ip1:  []byte{0, 0, 0, 0},
			ip2:  []byte{0, 0, 0},
			want: Incomparable,
		},
		{
			ip1:  []byte{0, 0, 0},
			ip2:  []byte{0, 0, 0},
			want: Equal,
		},
		{
			ip1:  []byte{0, 0, 1},
			ip2:  []byte{0, 0, 0},
			want: GreaterThan,
		},
		{
			ip1:  []byte{0, 0, 0},
			ip2:  []byte{0, 0, 0, 0},
			want: Incomparable,
		},
		{
			ip1:  []byte{0, 0, 0, 0},
			ip2:  nil,
			want: Incomparable,
		},
		{
			ip1:  nil,
			ip2:  []byte{0, 0, 0, 0},
			want: Incomparable,
		},
		{
			ip1:  nil,
			ip2:  nil,
			want: Equal,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := compareIPs(tt.ip1, tt.ip2); got != tt.want {
				t.Errorf("compareIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCIDRContainsCIDR(t *testing.T) {
	tests := []struct {
		name  string
		cidrA net.IPNet
		cidrB net.IPNet
		want  bool
	}{
		{
			name: "Cidr A with nil IP and Mask will return false",
			cidrA: net.IPNet{
				IP:   nil,
				Mask: nil,
			},
			cidrB: parseCIDR("10.0.0.0/24"),
			want:  false,
		},
		{
			name:  "Cidr B with nil IP and Mask will return false",
			cidrA: parseCIDR("10.0.0.0/24"),
			cidrB: net.IPNet{
				IP:   nil,
				Mask: nil,
			},
			want: false,
		},
		{
			name: "Cidr A with nil IP and proper Mask will return false",
			cidrA: net.IPNet{
				IP:   nil,
				Mask: net.CIDRMask(16, 32),
			},
			cidrB: parseCIDR("10.0.0.0/24"),
			want:  false,
		},
		{
			name: "Cidr A with nil Mask and proper IP will return false",
			cidrA: net.IPNet{
				IP:   net.IPv4(192, 168, 0, 1),
				Mask: nil,
			},
			cidrB: parseCIDR("192.168.0.1/32"),
			want:  false,
		},
		{
			name:  "Cidr A doesn't contain non overlapping Cidr B",
			cidrA: parseCIDR("192.168.0.0/24"),
			cidrB: parseCIDR("10.0.0.0/24"),
			want:  false,
		},
		{
			name:  "Cidr A contains the same Cidr B",
			cidrA: parseCIDR("192.168.0.0/24"),
			cidrB: parseCIDR("192.168.0.0/24"),
			want:  true,
		},
		{
			name:  "Cidr A doesn't contain Cidr B with same IP but smaller mask (it doesn't hold it entirely)",
			cidrA: parseCIDR("192.168.0.0/24"),
			cidrB: parseCIDR("192.168.0.0/16"),
			want:  false,
		},
		{
			name:  "Cidr A contains Cidr B with same IP but bigger mask (Cidr B is a subset of A)",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.168.0.0/24"),
			want:  true,
		},
		{
			name:  "Cidr A contains its subset - Cidr B with different IP",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.168.3.0/24"),
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CIDRContainsCIDR(tt.cidrA, tt.cidrB); got != tt.want {
				t.Errorf("CIDRContainsCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCIDRsOverlap(t *testing.T) {
	tests := []struct {
		name  string
		cidrA net.IPNet
		cidrB net.IPNet
		want  bool
	}{
		{
			name:  "Two exact same CIDRs will return true",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.168.0.0/16"),
			want:  true,
		},
		{
			name:  "Two CIDRs with the same IP will return true",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.168.0.0/32"),
			want:  true,
		},
		{
			name:  "Two CIDRs with the same IP will return true (reversed order)",
			cidrA: parseCIDR("192.168.0.0/32"),
			cidrB: parseCIDR("192.168.0.0/16"),
			want:  true,
		},
		{
			name:  "Two non overlapping CIDRs will return false",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.169.0.0/16"),
			want:  false,
		},
		{
			name:  "Two CIDRs with different IPs where one is a subset of second will return true",
			cidrA: parseCIDR("192.168.0.0/16"),
			cidrB: parseCIDR("192.168.1.0/24"),
			want:  true,
		},
		{
			name:  "Two CIDRs with different IPs where one is a subset of second will return true (reversed order)",
			cidrA: parseCIDR("192.168.1.0/24"),
			cidrB: parseCIDR("192.168.0.0/16"),
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CIDRsOverlap(tt.cidrA, tt.cidrB); got != tt.want {
				t.Errorf("CIDRsOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTestCIDRV4Pool(cidr string) *CIDRV4Pool {
	pool, err := NewCIDRV4Pool(cidr)
	if err != nil {
		panic(fmt.Sprintf("cannot create new testing CIDR pool: %v", err))
	}
	return pool
}

func newTestCIDRV4PoolWithTakenCIDRs(cidr string, taken []net.IPNet) *CIDRV4Pool {
	pool, err := NewCIDRV4Pool(cidr)
	if err != nil {
		panic(fmt.Sprintf("cannot create new testing CIDR pool: %v", err))
	}
	for i := range taken {
		if err = pool.ExcludeCIDRFromPool(taken[i]); err != nil {
			panic(fmt.Sprintf("cannot take a CIDR from testing pool: %v", err))
		}
	}
	return pool
}

func TestCIDRV4Pool_Get(t *testing.T) {
	tests := []struct {
		name      string
		pool      *CIDRV4Pool
		size      int
		want      net.IPNet
		wantEmpty bool
		wantErr   bool
	}{
		{
			name:    "testing requesting invalid size",
			pool:    newTestCIDRV4Pool("192.168.1.0/24"),
			size:    -1,
			wantErr: true,
		},
		{
			name:    "testing requesting invalid size (bigger than 32)",
			pool:    newTestCIDRV4Pool("192.168.1.0/24"),
			size:    64,
			wantErr: true,
		},
		{
			name:    "testing getting CIDR from nil pool",
			pool:    nil,
			size:    32,
			wantErr: true,
		},
		{
			name: "testing getting regular CIDR",
			pool: newTestCIDRV4Pool("192.168.1.0/24"),
			size: 32,
			want: parseCIDR("192.168.1.0/32"),
		},
		{
			name:      "testing getting too big CIDR",
			pool:      newTestCIDRV4Pool("192.168.1.0/24"),
			size:      16,
			wantEmpty: true,
		},
		{
			name: "testing getting CIDR of equal size",
			pool: newTestCIDRV4Pool("192.168.1.0/24"),
			size: 24,
			want: parseCIDR("192.168.1.0/24"),
		},
		{
			name: "testing getting CIDR when there is one already taken",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.1.0/24",
				[]net.IPNet{
					parseCIDR("192.168.1.0/28"),
				},
			),
			size: 28,
			want: parseCIDR("192.168.1.16/28"),
		},
		{
			name: "testing getting small CIDR when there is one big already taken",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.1.0/24",
				[]net.IPNet{
					parseCIDR("192.168.1.0/28"),
				},
			),
			size: 30,
			want: parseCIDR("192.168.1.16/30"),
		},
		{
			name: "testing getting big CIDR when there is one small already taken",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.1.0/24",
				[]net.IPNet{
					parseCIDR("192.168.1.0/28"),
				},
			),
			size: 26,
			want: parseCIDR("192.168.1.64/26"),
		},
		{
			name: "testing getting CIDR when the remaining CIDRs are not big enough",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.1.0/24",
				[]net.IPNet{
					parseCIDR("192.168.1.0/28"),
				},
			),
			size:      24,
			wantEmpty: true,
		},
		{
			name: "testing getting last remaining CIDR",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.1.0/24",
				[]net.IPNet{
					parseCIDR("192.168.1.0/25"),
				},
			),
			size: 25,
			want: parseCIDR("192.168.1.128/25"),
		},
		{
			name: "testing getting CIDR from pool with many different taken CIDRs",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.0.0/16",
				[]net.IPNet{
					parseCIDR("192.168.0.0/24"),
					parseCIDR("192.168.1.0/24"),
					parseCIDR("192.168.2.0/24"),
					parseCIDR("192.168.3.0/28"),
					parseCIDR("192.168.4.0/24"),
				},
			),
			size: 26,
			want: parseCIDR("192.168.3.64/26"),
		},
		{
			name: "testing getting CIDR from pool with many different taken CIDRs (getting last CIDR)",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.0.0/16",
				[]net.IPNet{
					parseCIDR("192.168.0.0/24"),
					parseCIDR("192.168.1.0/24"),
					parseCIDR("192.168.2.0/24"),
					parseCIDR("192.168.3.0/26"),
					parseCIDR("192.168.3.64/26"),
					parseCIDR("192.168.3.128/26"),
					parseCIDR("192.168.3.192/26"),
					parseCIDR("192.168.4.0/24"),
				},
			),
			size: 25,
			want: parseCIDR("192.168.5.0/25"),
		},
		{
			name: "testing getting CIDR from pool with many different taken CIDRs (getting the first CIDR)",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.0.0/16",
				[]net.IPNet{
					parseCIDR("192.168.1.0/24"),
					parseCIDR("192.168.2.0/24"),
					parseCIDR("192.168.3.0/26"),
					parseCIDR("192.168.3.64/26"),
					parseCIDR("192.168.3.128/26"),
				},
			),
			size: 25,
			want: parseCIDR("192.168.0.0/25"),
		},
		{
			name: "testing getting CIDR from pool with many different taken CIDRs (getting the second CIDR)",
			pool: newTestCIDRV4PoolWithTakenCIDRs(
				"192.168.0.0/16",
				[]net.IPNet{
					parseCIDR("192.168.0.0/25"),
					parseCIDR("192.168.2.0/24"),
					parseCIDR("192.168.3.0/26"),
					parseCIDR("192.168.3.64/26"),
					parseCIDR("192.168.3.128/26"),
				},
			),
			size: 25,
			want: parseCIDR("192.168.0.128/25"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pool.Get(tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("CIDRV4Pool.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.wantEmpty {
				if got != nil {
					t.Errorf("CIDRV4Pool.Get() wantEmpty but got %v", got)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("CIDRV4Pool.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Fix a bug with invalid obtained value.
func TestCIDRV4Pool_MultipleGets(t *testing.T) {
	pool := newTestCIDRV4PoolWithTakenCIDRs(
		"192.168.0.0/16",
		[]net.IPNet{
			parseCIDR("192.168.0.0/25"),
			parseCIDR("192.168.2.0/24"),
			parseCIDR("192.168.3.0/26"),
			parseCIDR("192.168.3.64/26"),
			parseCIDR("192.168.3.128/26"),
		},
	)
	cidr, err := pool.Get(16)
	if err != nil || cidr != nil {
		t.Fatalf("expected empty CIDR. Got err: %v, cidr: %v", err, cidr)
	}
	cidr, err = pool.Get(24)
	expectedCIDR := parseCIDR("192.168.1.0/24")
	if err != nil || !reflect.DeepEqual(*cidr, expectedCIDR) {
		t.Fatalf("expected %v CIDR. Got err: %v, cidr: %v", expectedCIDR, err, cidr)
	}
	cidr, err = pool.Get(25)
	expectedCIDR = parseCIDR("192.168.0.128/25")
	if err != nil || !reflect.DeepEqual(*cidr, expectedCIDR) {
		t.Fatalf("expected %v CIDR. Got err: %v, cidr: %v", expectedCIDR, err, cidr)
	}
	cidr, err = pool.Get(24)
	expectedCIDR = parseCIDR("192.168.4.0/24")
	if err != nil || !reflect.DeepEqual(*cidr, expectedCIDR) {
		t.Fatalf("expected %v CIDR. Got err: %v, cidr: %v", expectedCIDR, err, cidr)
	}
	cidr, err = pool.Get(32)
	expectedCIDR = parseCIDR("192.168.3.192/32")
	if err != nil || !reflect.DeepEqual(*cidr, expectedCIDR) {
		t.Fatalf("expected %v CIDR. Got err: %v, cidr: %v", expectedCIDR, err, cidr)
	}
	cidr, err = pool.Get(32)
	expectedCIDR = parseCIDR("192.168.3.193/32")
	if err != nil || !reflect.DeepEqual(*cidr, expectedCIDR) {
		t.Fatalf("expected %v CIDR. Got err: %v, cidr: %v", expectedCIDR, err, cidr)
	}
}

func Test_getBiggerCIDR(t *testing.T) {
	tests := []struct {
		name  string
		cidr1 net.IPNet
		cidr2 net.IPNet
		want  net.IPNet
	}{
		{
			name:  "when two equal CIDRs then the output will be the same",
			cidr1: parseCIDR("192.168.1.0/24"),
			cidr2: parseCIDR("192.168.1.0/24"),
			want:  parseCIDR("192.168.1.0/24"),
		},
		{
			name:  "when two CIDRs with equal size non overlapping then the first one is returned",
			cidr1: parseCIDR("192.168.2.0/24"),
			cidr2: parseCIDR("192.168.1.0/24"),
			want:  parseCIDR("192.168.2.0/24"),
		},
		{
			name:  "picking bigger one",
			cidr1: parseCIDR("192.168.1.0/28"),
			cidr2: parseCIDR("192.168.1.0/24"),
			want:  parseCIDR("192.168.1.0/24"),
		},
		{
			name:  "picking bigger one (reversed order)",
			cidr1: parseCIDR("192.168.1.0/24"),
			cidr2: parseCIDR("192.168.1.0/28"),
			want:  parseCIDR("192.168.1.0/24"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBiggerCIDR(tt.cidr1, tt.cidr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBiggerCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Adjust tests to the newer CIDR form.
//
//func constructList(CIDRs []net.IPNet) *cidrNode {
// 	if len(CIDRs) == 0 {
// 		return nil
// 	}
// 	root := &cidrNode{
// 		cidr: CIDRs[0],
// 	}
// 	current := root
// 	for i := range CIDRs[1:] {
// 		current.InsertAfter(CIDRs[i+1])
// 		current = current.Next()
// 	}
// 	return root
// }
//
// func listEqual(rootA, rootB *cidrNode) bool {
// 	for rootA != nil {
// 		if rootB == nil {
// 			return false
// 		}
// 		if !reflect.DeepEqual(rootA.cidr, rootB.cidr) {
// 			return false
// 		}
// 		rootA = rootA.Next()
// 		rootB = rootB.Next()
// 	}
// 	return rootB == nil
// }
//
// func listToSlice(root *cidrNode) []net.IPNet {
// 	result := []net.IPNet{}
// 	for root != nil {
// 		result = append(result, root.cidr)
// 		root = root.Next()
// 	}
// 	return result
// }
//
// func TestCIDRV4Pool_takeCIDR(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		pool         *CIDRV4Pool
// 		cidrToTake   *net.IPNet
// 		expectedPool *CIDRV4Pool
// 		wantErr      bool
// 	}{
// 		{
// 			name:       "working with empty pool",
// 			pool:       nil,
// 			cidrToTake: parseCIDR("192.168.0.0/28"),
// 			wantErr:    true,
// 		},
// 		{
// 			name: "working with empty cidr",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 			},
// 			cidrToTake: nil,
// 			wantErr:    true,
// 		},
// 		{
// 			name:       "working with empty pool and cidr",
// 			pool:       nil,
// 			cidrToTake: nil,
// 			wantErr:    true,
// 		},
// 		{
// 			name: "picking first CIDR from the subset will just store that CIDR",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 			},
// 			cidrToTake: parseCIDR("192.168.0.0/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.0.0/28"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "take method ignores CIDR being outside of the pool",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 			},
// 			cidrToTake: parseCIDR("192.0.0.0/24"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.0.0.0/24"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "picking the same CIDR as the one already taken will duplicate it",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.0.0/28"),
// 				}),
// 			},
// 			cidrToTake: parseCIDR("192.168.0.0/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.0.0/28"),
// 					parseCIDR("192.168.0.0/28"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "picking the CIDR with the same IP will place it after it",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.0.0/30"),
// 				}),
// 			},
// 			cidrToTake: parseCIDR("192.168.0.0/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.0.0/30"),
// 					parseCIDR("192.168.0.0/28"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "placing the taken CIDR in the middle",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/16"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.1.0/24"),
// 					parseCIDR("192.168.3.0/24"),
// 				}),
// 			},
// 			cidrToTake: parseCIDR("192.168.2.0/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.1.0/24"),
// 					parseCIDR("192.168.2.0/28"),
// 					parseCIDR("192.168.3.0/24"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "placing the taken CIDR in the front",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/16"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.2.16/28"),
// 					parseCIDR("192.168.3.0/24"),
// 				}),
// 			},
// 			cidrToTake: parseCIDR("192.168.2.0/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.2.0/28"),
// 					parseCIDR("192.168.2.16/28"),
// 					parseCIDR("192.168.3.0/24"),
// 				}),
// 			},
// 		},
// 		{
// 			name: "placing the taken CIDR in unsorted list will place it before the first greater one",
// 			pool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/16"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.4.0/28"),
// 					parseCIDR("192.168.3.0/24"),
// 					parseCIDR("192.168.2.48/28"),
// 					parseCIDR("192.168.4.32/28"),
// 					parseCIDR("192.168.5.0/28"),
// 					parseCIDR("192.168.1.0/24"),
// 				}),
// 			},
// 			cidrToTake: parseCIDR("192.168.4.16/28"),
// 			expectedPool: &CIDRV4Pool{
// 				CIDR: parseCIDR("192.168.0.0/24"),
// 				cidrsInUse: constructList([]net.IPNet{
// 					parseCIDR("192.168.4.0/28"),
// 					parseCIDR("192.168.3.0/24"),
// 					parseCIDR("192.168.2.48/28"),
// 					parseCIDR("192.168.4.16/28"),
// 					parseCIDR("192.168.4.32/28"),
// 					parseCIDR("192.168.5.0/28"),
// 					parseCIDR("192.168.1.0/24"),
// 				}),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.pool.takeCIDR(tt.cidrToTake); (err != nil) != tt.wantErr {
// 				t.Errorf("CIDRV4Pool.takeCIDR() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !tt.wantErr && !listEqual(tt.pool.cidrsInUse, tt.expectedPool.cidrsInUse) {
// 				t.Errorf(
// 					"CIDRV4Pool taken cidrs = %v, want %v",
// 					listToSlice(tt.pool.cidrsInUse), listToSlice(tt.expectedPool.cidrsInUse))
// 			}
// 		})
// 	}
// }
//
// func TestCIDRV4Pool_isAlreadyTaken(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		pool        *CIDRV4Pool
// 		cidrToCheck *net.IPNet
// 		want        bool
// 		wantErr     bool
// 	}{
// 		{
// 			name:    "Empty pool and CIDR to check should return error",
// 			wantErr: true,
// 		},
// 		{
// 			name:        "Empty pool and should return error",
// 			cidrToCheck: parseCIDR("192.168.0.0/24"),
// 			wantErr:     true,
// 		},
// 		{
// 			name:    "Empty CIDR to check should return error",
// 			pool:    newTestCIDRV4Pool("192.168.0.0/24"),
// 			wantErr: true,
// 		},
// 		{
// 			name:        "Checking not taken yet CIDR from the pool",
// 			pool:        newTestCIDRV4Pool("192.168.0.0/24"),
// 			cidrToCheck: parseCIDR("192.168.0.0/24"),
// 			want:        false,
// 		},
// 		{
// 			name:        "Checking CIDR outside of the pool returns false because it doesn't check pool boundaries",
// 			pool:        newTestCIDRV4Pool("192.168.0.0/24"),
// 			cidrToCheck: parseCIDR("192.169.0.0/24"),
// 			want:        true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.pool.isAlreadyTaken(tt.cidrToCheck)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CIDRV4Pool.isAlreadyTaken() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !tt.wantErr && got != tt.want {
// 				t.Errorf("CIDRV4Pool.isAlreadyTaken() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
