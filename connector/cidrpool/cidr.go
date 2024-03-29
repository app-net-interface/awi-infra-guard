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

package cidrpool

import (
	"fmt"
	"net"
)

// Used to invert IP form represented by Mask.
func invert32Bytes(bytes []byte) ([]byte, error) {
	newBytes := make([]byte, 4)
	if len(bytes) != 4 {
		return nil, fmt.Errorf("invalid IPv4 to invert. Expected 4 bytes, got: %v", bytes)
	}
	for i := range bytes {
		newBytes[i] = 255 - bytes[i]
	}
	return newBytes, nil
}

func increaseIPBy32Bytes(ip net.IP, bytes []byte) (net.IP, error) {
	newBytes := make([]byte, 4)
	if len(ip) != 4 {
		return nil, fmt.Errorf("invalid IPv4. got: %v", ip)
	}
	if len(bytes) != 4 {
		return nil, fmt.Errorf("invalid IPv4 to add. Expected 4 bytes, got: %v", bytes)
	}
	// TODO: Handle rejecting IPv6
	carry := byte(0)
	for i := range bytes {
		i = 3 - i
		newBytes[i] = (ip[i] + bytes[i] + carry)
		if int(ip[i])+int(bytes[i])+int(carry) > 255 {
			carry = byte(1)
		} else {
			carry = byte(0)
		}
	}

	return newBytes, nil
}

// Returns the new CIDR with the same mask and the next
// IP for a given Network Mask.
//
// For example, for a CIDR 192.168.1.0/24 the nextCIDR
// will return 192.168.2.0/24 and for CIDR 192.168.1.0/28
// the nextCIDR will return 192.168.1.16/28.
func nextCIDR(cidr net.IPNet) (net.IPNet, error) {
	invertedMask, err := invert32Bytes(cidr.Mask)
	if err != nil {
		return net.IPNet{}, err
	}
	ipOffset, err := increaseIPBy32Bytes(invertedMask, []byte{0, 0, 0, 1})
	if err != nil {
		return net.IPNet{}, err
	}
	newIP, err := increaseIPBy32Bytes(cidr.IP, ipOffset)
	if err != nil {
		return net.IPNet{}, err
	}
	return net.IPNet{
		IP:   newIP,
		Mask: cidr.Mask,
	}, nil
}

type comparisonResult int

const (
	GreaterThan comparisonResult = iota
	LesserThan
	Equal
	Incomparable
)

// compareIPs returns information how ip1 is related to ip2.
// The greater IP Address means it is closer to 255.255.255.255
// and further from 0.0.0.0
//
// For ip1 = 192.168.0.0 and ip2 192.168.2.0 the result will be
// LesserThan.
//
// If IP Address versions do not match, the method will return
// Incomparable.
func compareIPs(ip1, ip2 net.IP) comparisonResult {
	if len(ip1) != len(ip2) {
		return Incomparable
	}
	for i := range ip1 {
		if ip1[i] > ip2[i] {
			return GreaterThan
		} else if ip1[i] < ip2[i] {
			return LesserThan
		}
	}
	return Equal
}

// Returns true if both CIDRs are adjacent and
// can be merged into single CIDR covering exactly
// and only both of them. Read more in cidrList insert method.
//
// Adjacent CIDRs can be merged only if the lower CIDR IP is
// the same as for the IP of a CIDR with mask decreased by one.
// For example, CIDR 192.168.0.0/24 has the same IP as CIDR
// 192.168.0.0/23 but same cannot apply for CIDR 192.168.1.0/24
// as there is no such CIDR for mask 23 (either 192.168.0.0/23
// or 192.168.2.0/23).
func AreCIDRsAdjacentAndMergeable(lowerCIDR, higherCIDR net.IPNet) (bool, error) {
	adjacent, err := CIDRsAreAdjacent(lowerCIDR, higherCIDR)
	if err != nil {
		return false, fmt.Errorf(
			"failed to check if CIDRs %s and %s are adjacent",
			lowerCIDR.String(), higherCIDR.String(),
		)
	}
	if !adjacent {
		return false, nil
	}

	cidrSize, _ := lowerCIDR.Mask.Size()
	if cidrSize == 0 {
		return false, nil
	}

	_, biggerCIDR, err := net.ParseCIDR(
		fmt.Sprintf("%s/%d", lowerCIDR.IP.String(), cidrSize-1),
	)
	if err != nil {
		return false, fmt.Errorf(
			"failed to check bigger CIDR for CIDR %s: %w",
			lowerCIDR.String(), err,
		)
	}

	return lowerCIDR.IP.Equal(biggerCIDR.IP), nil
}

// GrowCIDRByOne returns CIDR with a mask lower
// by one.
//
// Example: For CIDR 192.168.1.0/24 the resulting CIDR
// will be 192.168.0.0/23.
func GrowCIDRByOne(cidr net.IPNet) (net.IPNet, error) {
	cidrSize, _ := cidr.Mask.Size()
	if cidrSize <= 0 {
		return net.IPNet{}, fmt.Errorf("cannot decrease CIDR mask by one with CIDR: %s", cidr.String())
	}
	_, biggerCIDR, err := net.ParseCIDR(
		fmt.Sprintf("%s/%d", cidr.IP.String(), cidrSize-1),
	)
	if err != nil {
		return net.IPNet{}, fmt.Errorf("failed to calculate bigger cidr: %w", err)
	}
	if biggerCIDR == nil {
		return net.IPNet{}, fmt.Errorf("got nil CIDR from parsing %s",
			fmt.Sprintf("%s/%d", cidr.IP.String(), cidrSize-1))
	}
	return *biggerCIDR, nil
}

// Returns true if cidrB lies within boundaries of cidrA.
//
// The cidrB is considered to be contained by the cidrA if all
// IP Addresses that belong to cidrB belong to cidrA as well.
func CIDRContainsCIDR(cidrA net.IPNet, cidrB net.IPNet) bool {
	if !cidrA.Contains(cidrB.IP) {
		return false
	}
	sizeA, _ := cidrA.Mask.Size()
	sizeB, _ := cidrB.Mask.Size()
	return sizeA <= sizeB
}

// Returns true if both CIDRs are exactly next to each other.
//
// Examples of adjacenting CIDRs:
// * 192.168.0.0/24 and 192.168.1.0/24
// * 192.168.0.0/32 and 192.168.0.1/32
// * 192.168.1.0/28 and 192.168.1.16/28
//
// Adjacenting is not considered across overflow. CIDRs such
// as 255.255.255.0/24 and 0.0.0.0/24 are not being compared.
func CIDRsAreAdjacent(cidrA net.IPNet, cidrB net.IPNet) (bool, error) {
	if compareIPs(cidrA.IP, cidrB.IP) == LesserThan {
		cidrA, cidrB = cidrB, cidrA
	}
	next, err := nextCIDR(cidrB)
	if err != nil {
		return false, fmt.Errorf(
			"cannot calculate adjacent CIDR to CIDR %v: %w",
			cidrB, err,
		)
	}
	return next.IP.Equal(cidrA.IP), nil
}

// Returns true if there is at least one IP Address which
// belongs to both cidrA and cidrB.
func CIDRsOverlap(cidrA net.IPNet, cidrB net.IPNet) bool {
	if cidrA.Contains(cidrB.IP) || cidrB.Contains(cidrA.IP) {
		return true
	}
	return false
}

// Returns the CIDR containing more IP Addresses.
//
// If both CIDRs are equal, returns the first one.
func getBiggerCIDR(cidr1, cidr2 net.IPNet) net.IPNet {
	size1, _ := cidr1.Mask.Size()
	size2, _ := cidr2.Mask.Size()
	if size1 <= size2 {
		return cidr1
	}
	return cidr2
}

func cidrWithin(cidr net.IPNet, cidrs []net.IPNet) bool {
	for _, checked := range cidrs {
		if CIDRContainsCIDR(checked, cidr) {
			return true
		}
	}
	return false
}

// IntersectingCIDRs returns CIDRs that are a part of
// left CIDRs and right CIDRs (if a CIDR in left is a
// part of bigger CIDR in right that counts as well).
func IntersectingCIDRs(leftCIDRs, rightCIDRs []string) ([]string, error) {
	leftParsed := make([]net.IPNet, 0, len(leftCIDRs))
	rightParsed := make([]net.IPNet, 0, len(rightCIDRs))

	for _, cidr := range leftCIDRs {
		_, parsed, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to process CIDR '%s' from left CIDRs: %w",
				cidr, err,
			)
		}
		if parsed == nil {
			return nil, fmt.Errorf(
				"failed to process CIDR '%s' from left CIDRs. Got nil CIDR",
				cidr,
			)
		}
		leftParsed = append(leftParsed, *parsed)
	}

	for _, cidr := range rightCIDRs {
		_, parsed, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to process CIDR '%s' from right CIDRs: %w",
				cidr, err,
			)
		}
		if parsed == nil {
			return nil, fmt.Errorf(
				"failed to process CIDR '%s' from right CIDRs. Got nil CIDR",
				cidr,
			)
		}
		rightParsed = append(rightParsed, *parsed)
	}

	matched := make([]string, 0, len(leftCIDRs)+len(rightCIDRs))
	for _, cidr := range leftParsed {
		if cidrWithin(cidr, rightParsed) {
			matched = append(matched, cidr.String())
		}
	}

	for _, cidr := range rightParsed {
		if cidrWithin(cidr, leftParsed) {
			matched = append(matched, cidr.String())
		}
	}

	return matched, nil
}
