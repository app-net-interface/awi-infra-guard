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

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// Struct representing Pool for IPv4 CIDRs.
//
// The usage is to type an available CIDR pool
// (for example 192.168.0.0/16) and exclude forbidden CIDRs
// (you may want to ensure that the Pool won't return subsets
// of following CIDRs [192.168.1.0/24, 192.168.12.16/28] etc).
//
// After doing so, the Get(mask_size) method will provide the
// first available CIDR with a given mask size. For example,
// if we specified a pool 192.168.0.0/16 and we ask to Get a
// CIDR with mask size 28, we will get CIDR: 192.168.0.0/28.
// If we ask for such CIDR again, we will get 192.168.0.16/28.
type CIDRV4Pool struct {
	CIDR       *net.IPNet
	cidrsInUse *cidrNode
}

func NewCIDRV4Pool(cidr string) (*CIDRV4Pool, error) {
	_, parsed, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	if strings.Contains(cidr, ":") {
		return nil, errors.New("CIDR IPv6 are not supported")
	}
	return &CIDRV4Pool{
		CIDR: parsed,
	}, nil
}

func (c *CIDRV4Pool) Size() int {
	size, _ := c.CIDR.Mask.Size()
	return size
}

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
)

func compareIPs(ip1, ip2 net.IP) comparisonResult {
	if len(ip1) > len(ip2) {
		return GreaterThan
	}
	if len(ip1) < len(ip2) {
		return LesserThan
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

type cidrNode struct {
	cidr *net.IPNet
	next *cidrNode
}

func (c *cidrNode) InsertAfter(cidr *net.IPNet) {
	if c == nil {
		return
	}
	currentNext := c.next
	c.next = &cidrNode{
		cidr: cidr,
		next: currentNext,
	}
}

func (c *cidrNode) RemoveNext() {
	if c == nil {
		return
	}
	if c.next == nil {
		return
	}
	c.next = c.next.next
}

func (c *cidrNode) Next() *cidrNode {
	if c == nil {
		return nil
	}
	return c.next
}

func cidrSize(cidr string) (int, error) {
	_, parsed, err := net.ParseCIDR(cidr)
	if err != nil {
		return 0, fmt.Errorf("invalid CIDR form. %w", err)
	}
	size, _ := parsed.Mask.Size()
	return size, nil
}

// This is to register a CIDR as unavailable to the pool.
//
// It assumes that the CIDR do not overlap with any already
// taken CIDR and that it is a part of global CIDR.
func (c *CIDRV4Pool) takeCIDR(cidr *net.IPNet) error {
	if c == nil {
		return errors.New("cannot insert anything to empty pool")
	}
	if cidr == nil {
		return errors.New("cannot insert empty CIDR")
	}
	currentNode := c.cidrsInUse
	if currentNode == nil {
		c.cidrsInUse = &cidrNode{
			cidr: cidr,
		}
		return nil
	}
	if compareIPs(cidr.IP, currentNode.cidr.IP) == LesserThan {
		c.cidrsInUse = &cidrNode{
			cidr: cidr,
			next: currentNode,
		}
		return nil
	}
	for currentNode.Next() != nil {
		if compareIPs(cidr.IP, currentNode.Next().cidr.IP) == LesserThan {
			break
		}
		currentNode = currentNode.Next()
	}
	currentNode.InsertAfter(cidr)
	return nil
}

// Returns true if cidrB lies within boundaries of cidrA.
//
// The cidrB is considered to be contained by the cidrA if all
// IP Addresses that belong to cidrB belong to cidrA as well.
func CIDRContainsCIDR(cidrA *net.IPNet, cidrB *net.IPNet) bool {
	if cidrA == nil {
		return false
	}
	if cidrB == nil {
		return false
	}
	if !cidrA.Contains(cidrB.IP) {
		return false
	}
	sizeA, _ := cidrA.Mask.Size()
	sizeB, _ := cidrB.Mask.Size()
	return sizeA <= sizeB
}

// Returns true if there is at least one IP Address which
// belongs to both cidrA and cidrB.
func CIDRsOverlap(cidrA *net.IPNet, cidrB *net.IPNet) bool {
	if cidrA == nil {
		return false
	}
	if cidrB == nil {
		return false
	}
	if cidrA.Contains(cidrB.IP) {
		return true
	}
	if cidrB.Contains(cidrA.IP) {
		return true
	}
	return false
}

// Returns true if taking the given CIDR would not change anything.
func (c *CIDRV4Pool) isAlreadyTaken(cidr *net.IPNet) (bool, error) {
	if c == nil {
		return false, errors.New("checking empty pool")
	}
	if cidr == nil {
		return false, errors.New("cannot check for empty CIDR")
	}
	currentNode := c.cidrsInUse
	for currentNode != nil {
		if CIDRContainsCIDR(c.cidrsInUse.cidr, cidr) {
			return true, nil
		}
		currentNode = currentNode.Next()
	}
	return false, nil
}

// Removes all CIDRs that are within provided CIDR. It is for cleaning
// purposes since there is no need in storing those CIDRs anymore as
// we have broader one already preventing from obtaining those CIDRs.
func (c *CIDRV4Pool) removeCIDRsThatArePartOf(cidr *net.IPNet) error {
	if c == nil {
		return errors.New("trying to remove CIDRs from empty pool")
	}
	if cidr == nil {
		return errors.New("provided CIDR is empty/invalid")
	}
	currentNode := c.cidrsInUse
	for currentNode != nil && CIDRContainsCIDR(cidr, currentNode.cidr) {
		currentNode = currentNode.Next()
		c.cidrsInUse = currentNode
	}
	if currentNode == nil {
		return nil
	}
	for currentNode.Next() != nil {
		if CIDRContainsCIDR(cidr, currentNode.Next().cidr) {
			currentNode.RemoveNext()
		} else {
			currentNode = currentNode.Next()
		}
	}
	return nil
}

// This method forcefully removes a given CIDR from the pool.
//
// The method ignores previously loaned CIDRs. If the CIDR
// is already taken, no action will happen. If a subset of that
// CIDR is taken, the Pool will extend it to match the given
// pool.
//
// If the CIDR is out of range of this pool or is greater than
// this pool, the error will be returned.
func (c *CIDRV4Pool) ExcludeCIDRFromPool(cidr *net.IPNet) error {
	if c == nil {
		return errors.New("cannot insert anything to empty pool")
	}
	if cidr == nil {
		return errors.New("cannot insert empty CIDR")
	}
	if !CIDRContainsCIDR(c.CIDR, cidr) {
		return fmt.Errorf(
			"CIDR: %s is out of range from CIDR: %s", cidr.String(), c.CIDR.String(),
		)
	}
	taken, err := c.isAlreadyTaken(cidr)
	if err != nil {
		return fmt.Errorf(
			"Cannot verify if the CIDR %s is not already taken from the pool: %w", cidr.String(), err,
		)
	}
	if taken {
		return nil
	}
	if err = c.removeCIDRsThatArePartOf(cidr); err != nil {
		return fmt.Errorf(
			"Cannot clean up CIDRs from the pool that are no longer needed: %w", err,
		)
	}
	if err = c.takeCIDR(cidr); err != nil {
		return fmt.Errorf(
			"error while trying to take the CIDR %s: %w", cidr.String(), err,
		)
	}
	return nil
}

func getBiggerCIDR(cidr1, cidr2 *net.IPNet) *net.IPNet {
	if cidr1 == nil && cidr2 == nil {
		return nil
	}
	if cidr1 == nil {
		return cidr2
	}
	if cidr2 == nil {
		return cidr1
	}
	size1, _ := cidr1.Mask.Size()
	size2, _ := cidr2.Mask.Size()
	if size1 <= size2 {
		return cidr1
	}
	return cidr2
}

// Looks for a new potential CIDR when the previous one turned out to be overlapping
// with already taken CIDR.
//
// The logic behind proposing new CIDR is as follows:
//   - pick greater CIDR from either previous proposal or the overlapping taken CIDR
//   - find a first IP Address, greater than that CIDR not belonging to that CIDR
//   - propose a new CIDR consisting of that IP Address and the Mask from previous
//     proposal
func (c *CIDRV4Pool) proposeNewCIDR(olderProposal, overlappedCIDR *net.IPNet) (*net.IPNet, error) {
	biggerCIDR := getBiggerCIDR(olderProposal, overlappedCIDR)
	if biggerCIDR == nil {
		return nil, errors.New("Received empty CIDR")
	}
	nextCidr, err := nextCIDR(*biggerCIDR)
	if err != nil {
		return nil, fmt.Errorf("Error while finding new CIDR: %w", err)
	}
	return &net.IPNet{
		IP:   nextCidr.IP,
		Mask: olderProposal.Mask,
	}, nil
}

// Obtains first available CIDR with a given size from the pool.
//
// If a requested CIDR size cannot be obtained as there is no
// big enough CIDR left available, the method will return nil
// pointer and no error.
//
// If something else fails, the method will return an error.
func (c *CIDRV4Pool) Get(size int) (*net.IPNet, error) {
	if c == nil {
		return nil, errors.New("Cannot request a CIDR from a pool as it is nil")
	}
	if size < 0 || size > 32 {
		return nil, fmt.Errorf("invalid requested size. Must be within range 0-32. Got %d", size)
	}
	if size < c.Size() {
		return nil, nil
	}

	proposedCIDR := &net.IPNet{
		IP:   c.CIDR.IP,
		Mask: net.CIDRMask(size, 32),
	}

	if c.cidrsInUse == nil {
		return proposedCIDR, c.takeCIDR(proposedCIDR)
	}

	currentNode := c.cidrsInUse

	for currentNode != nil {
		if !CIDRContainsCIDR(c.CIDR, proposedCIDR) {
			// We have breached our main CIDR. It means no available CIDR
			// was found.
			return nil, nil
		}
		if currentNode.Next() != nil {
			result := compareIPs(currentNode.Next().cidr.IP, proposedCIDR.IP)
			if result == LesserThan || result == Equal {
				// While checking the proposed CIDR we want to fit properly with the
				// IPs of other CIDRs, so we skip to the closest CIDR which IP is lower
				// or equal to ours.
				// If the most close CIDR on our left side doesn't overlap with our
				// CIDR then all other CIDRs on the left side won't overlap as well
				// since they are not overlapping among themselves.
				currentNode = currentNode.Next()
				continue
			}
		}
		if !CIDRsOverlap(proposedCIDR, currentNode.cidr) {
			if currentNode.Next() == nil {
				return proposedCIDR, c.takeCIDR(proposedCIDR)
			}
			if !CIDRsOverlap(proposedCIDR, currentNode.Next().cidr) {
				return proposedCIDR, c.takeCIDR(proposedCIDR)
			}
			currentNode = currentNode.Next()
		}
		var err error
		proposedCIDR, err = c.proposeNewCIDR(proposedCIDR, currentNode.cidr)
		if err != nil {
			return nil, fmt.Errorf(
				"an error occurred while looking for new possible CIDR: %w", err)
		}
		currentNode = currentNode.Next()
	}

	if !CIDRContainsCIDR(c.CIDR, proposedCIDR) {
		// We have breached our main CIDR. It means no available CIDR
		// was found.
		return nil, nil
	}
	return proposedCIDR, nil

}
