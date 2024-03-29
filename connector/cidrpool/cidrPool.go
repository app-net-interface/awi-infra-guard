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
	"errors"
	"fmt"
	"net"
	"strings"
)

// Wrapper for multiple CIDRV4Pool which allows holding
// multiple CIDR addresses as a main entry.
//
// The CIDRV4Pools gets a slice of CIDR addresses and
// parses them to maintain a pool as requested.
//
// The CIDRs provided as an argument may overlap or be
// duplicated, the purpose of this structure is to handle
// that situation carefully, remove any duplicates and
// provide a clean pool.
type CIDRV4Pools struct {
	Pools []*CIDRV4Pool
}

func NewCIDRV4Pools(cidrs []string) (CIDRV4Pools, error) {
	parsed, err := separateCIDRs(cidrs)
	if err != nil {
		return CIDRV4Pools{}, fmt.Errorf(
			"could not create CIDR V4 Pools due to CIDR parsing issue: %w", err,
		)
	}
	output := CIDRV4Pools{}
	for _, parsedCIDR := range parsed {
		cidrPool, err := NewCIDRV4Pool(parsedCIDR)
		if err != nil {
			return CIDRV4Pools{}, fmt.Errorf(
				"could not create CIDR V4 Pools as parsing '%s' cidr failed: %w",
				parsedCIDR, err)
		}
		if cidrPool == nil {
			return CIDRV4Pools{}, fmt.Errorf(
				"could not create CIDR V4 Pools as parsing '%s' cidr resulted in nil object",
				parsedCIDR)
		}
		output.Pools = append(output.Pools, cidrPool)
	}
	return output, nil
}

// separateCIDRs goes through the slice of cidrs, removes
// duplicates, removes overlapping (by removing CIDRs that
// are subsets of other CIDRs) and merges neighbours that
// can be merged.
func separateCIDRs(cidrs []string) ([]string, error) {
	listOfCidrs := cidrList{}

	for _, cidr := range cidrs {
		_, parsed, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse '%s' to a proper CIDR: %w",
				cidr, err,
			)
		}
		if parsed == nil {
			return nil, fmt.Errorf(
				"got empty CIDR from parsing '%s'", cidr,
			)
		}
		if err = listOfCidrs.Insert(*parsed); err != nil {
			return nil, fmt.Errorf(
				"failed to add a parsed CIDR '%s' to a list of CIDRs "+
					"in order to separate CIDRs: %v", cidr, err,
			)
		}
	}

	separated := make([]string, 0, len(cidrs))

	currNode := listOfCidrs.Head()
	for currNode != nil {
		separated = append(separated, currNode.cidr.String())
		currNode = currNode.Next()
	}

	return separated, nil
}

// Returns true if there is no available IPs.
func (c *CIDRV4Pools) Full() bool {
	for _, pool := range c.Pools {
		if !pool.Full() {
			return false
		}
	}
	return true
}

func (c *CIDRV4Pools) Get(size int) (*net.IPNet, error) {
	for i := range c.Pools {
		if c.Pools[i] == nil {
			continue
		}
		pool, err := c.Pools[i].Get(size)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain a cidr pool: %w", err)
		}
		if pool != nil {
			return pool, nil
		}
	}
	return nil, nil
}

func (c *CIDRV4Pools) GetIP() (string, error) {
	cidr, err := c.Get(32)
	if err != nil {
		return "", fmt.Errorf(
			"failed to obtain a CIDR with size 32: %w", err,
		)
	}
	if cidr == nil {
		return "", nil
	}
	return cidr.IP.String(), nil
}

func (c *CIDRV4Pools) ExcludeCIDRFromPools(cidr string) error {
	if c == nil {
		return nil
	}
	_, parsedCIDR, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf(
			"failed to parse CIDR %s: %v", cidr, err,
		)
	}
	if parsedCIDR == nil {
		return fmt.Errorf(
			"failed to parse CIDR %s. Got empty CIDR", cidr,
		)
	}
	for i := range c.Pools {
		if c.Pools[i] == nil {
			continue
		}
		if !CIDRContainsCIDR(c.Pools[i].CIDR, *parsedCIDR) {
			continue
		}
		if err := c.Pools[i].ExcludeCIDRFromPool(*parsedCIDR); err != nil {
			return fmt.Errorf(
				"failed to remove CIDR %s from pools: %w",
				cidr, err,
			)
		}
		return nil
	}
	return nil
}

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
	CIDR       net.IPNet
	cidrsInUse cidrList
}

func NewCIDRV4Pool(cidr string) (*CIDRV4Pool, error) {
	_, parsed, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf(
			"got error while parsing cidr '%s': %w",
			cidr, err,
		)
	}
	if parsed == nil {
		return nil, fmt.Errorf(
			"got nil from parsing cidr: %s", cidr,
		)
	}
	if strings.Contains(cidr, ":") {
		return nil, errors.New("CIDR IPv6 are not supported")
	}
	return &CIDRV4Pool{
		CIDR: *parsed,
	}, nil
}

// Returns true if there are no more available
// IPs in the pool.
func (c *CIDRV4Pool) Full() bool {
	if c == nil {
		// If pool is nil there are no more IPs available :)
		return true
	}
	head := c.cidrsInUse.Head()
	return head != nil && head.Next() == nil && head.cidr.String() == c.CIDR.String()
}

func (c *CIDRV4Pool) CIDRsInUse() *cidrNode {
	if c == nil {
		return nil
	}
	return c.cidrsInUse.Head()
}

func (c *CIDRV4Pool) Size() int {
	size, _ := c.CIDR.Mask.Size()
	return size
}

// Obtains first available IP within the pool.
//
// If there is no more IPs available within the pool,
// the method will return empty string.
//
// Getting IP is equal to calling Get() with size 32.
func (c *CIDRV4Pool) GetIP() (string, error) {
	if c == nil {
		return "", errors.New("cannot request a CIDR from a pool as it is nil")
	}
	cidr, err := c.Get(32)
	if err != nil {
		return "", fmt.Errorf("failed to obtain CIDR with size 32: %w", err)
	}
	if cidr == nil {
		return "", nil
	}
	return cidr.IP.String(), nil
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
		return nil, errors.New("cannot request a CIDR from a pool as it is nil")
	}
	if size < 0 || size > 32 {
		return nil, fmt.Errorf("invalid requested size. Must be within range 0-32. Got %d", size)
	}
	if size < c.Size() {
		return nil, nil
	}

	proposedCIDR := net.IPNet{
		IP:   c.CIDR.IP,
		Mask: net.CIDRMask(size, 32),
	}

	if c.CIDRsInUse() == nil {
		return &proposedCIDR, c.cidrsInUse.Insert(proposedCIDR)
	}

	currentNode := c.CIDRsInUse()

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
				return &proposedCIDR, c.cidrsInUse.Insert(proposedCIDR)
			}
			if !CIDRsOverlap(proposedCIDR, currentNode.Next().cidr) {
				return &proposedCIDR, c.cidrsInUse.Insert(proposedCIDR)
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
	return &proposedCIDR, c.cidrsInUse.Insert(proposedCIDR)
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
func (c *CIDRV4Pool) ExcludeCIDRFromPool(cidr net.IPNet) error {
	if c == nil {
		return nil
	}
	if !CIDRContainsCIDR(c.CIDR, cidr) {
		return fmt.Errorf(
			"CIDR: %s is out of range from CIDR: %s", cidr.String(), c.CIDR.String(),
		)
	}
	if err := c.cidrsInUse.Insert(cidr); err != nil {
		return fmt.Errorf(
			"failed to update list of blocked CIDRs: %w", err,
		)
	}
	return nil
}

// Looks for a new potential CIDR when the previous one turned out to be overlapping
// with already taken CIDR.
//
// The logic behind proposing new CIDR is as follows:
//   - pick greater CIDR from either previous proposal or the overlapping taken CIDR
//   - find a first IP Address, greater than that CIDR not belonging to that CIDR
//   - propose a new CIDR consisting of that IP Address and the Mask from previous
//     proposal
func (c *CIDRV4Pool) proposeNewCIDR(olderProposal, overlappedCIDR net.IPNet) (net.IPNet, error) {
	biggerCIDR := getBiggerCIDR(olderProposal, overlappedCIDR)
	nextCidr, err := nextCIDR(biggerCIDR)
	if err != nil {
		return net.IPNet{}, fmt.Errorf("error while finding new CIDR: %w", err)
	}
	return net.IPNet{
		IP:   nextCidr.IP,
		Mask: olderProposal.Mask,
	}, nil
}
