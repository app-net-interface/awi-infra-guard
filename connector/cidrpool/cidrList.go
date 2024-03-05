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
	"errors"
	"fmt"
	"net"
)

// cidrList is a helper for storing a list of
// non-overlapping and non-adjacent CIDRs in
// an ascending order.
//
// The CIDR list is focused on inserting new
// CIDR addresses which may result in removing
// older entries that are a subset of a new cidr
// or not doing anything if the cidr is already
// present within the list.
//
// CIDRs, which are a part of this list, are
// sorted in an ascending order.
//
// CIDRs do not overlap. Any insertion of a CIDR
// overlapping or adjacenting with other CIDRs
// will result in node merging.
type cidrList struct {
	head *cidrNode
}

func (c *cidrList) Head() *cidrNode {
	if c == nil {
		return nil
	}
	return c.head
}

// Insert ensures that the provided CIDR is covered by
// the list.
//
// The Insert method checks if the cidr is already covered
// by the list (if all IP Addresses are within CIDR addresses
// already present in the node). If the CIDR is covered no
// additional insertion is made.
//
// If the CIDR is not covered, the cidr is inserted as a
// new Node at the proper position to maintain ascending order
// of CIDR IP Addresses.
//
// To ensure there are no overlaps, the insertion is preceeded
// with the removal of any CIDR that is a subset of that
// CIDR. Additionally, after insertion is done, the list checks
// if neighbour CIDRs can be replaced with one bigger CIDR.
func (c *cidrList) Insert(cidr net.IPNet) error {
	if c == nil {
		return errors.New("cannot insert to non existing list")
	}

	if c.isCovered(cidr) {
		// Nothing to do.
		return nil
	}

	if err := c.removeSubsetsOf(cidr); err != nil {
		return fmt.Errorf(
			"failed to remove existing CIDRs that were a part of inserted cidr %s: %w",
			cidr.String(), err,
		)
	}

	if err := c.insertAndMergeAdjacent(cidr); err != nil {
		return fmt.Errorf(
			"failed to insert CIDR %s: %w",
			cidr.String(), err,
		)
	}

	return nil
}

// insertAndMergeAdjacent places the CIDR to fit the ascending order of
// CIDRs in the list. It doesn't check if the CIDR is
// overlapping with other CIDRs so that should be handled
// by the different method.
//
// The method, however, checks if the inserted CIDR is adjacent
// to its neighbours - if it is, it checks if both CIDRs can be
// merged into one broader cidr and it does if that is possible.
//
// Example of cidrs that will be merged:
//   - 192.168.0.0/24 and 192.168.1.0/24 can be replaced with
//     192.168.0.0/23
//
// Examples of cidrs that cannot be merged:
//
//   - 192.168.0.0/24 and 192.168.1.0/25 - different masks cause
//     fractions - it is either 3 portions of /25 CIDRs or 1.5
//     portion of /24. The mask must be equal.
//
//   - 192.167.255.0/24 and 192.168.0.0/24 - these two CIDRs share
//     the same mask and are adjacent to each other but the cidr
//     with mask /23 cannot start at 192.167.255.0 - it can start
//     either at 192.167.254.0 which would not cover 192.168.0.0/24
//     or at 192.168.0.0 which would not cover 192.167.255.0/24.
//     The lower CIDR needs to start at the same position as the
//     CIDR with mask decreased by one.
func (c *cidrList) insertAndMergeAdjacent(cidr net.IPNet) error {
	if c == nil {
		return errors.New("cannot insert to non existing list")
	}

	// Insert node at the proper position and get the node
	// that preceeds it.
	currentNode, err := c.insert(cidr)
	if err != nil {
		return fmt.Errorf(
			"CIDR Node insertion failed: %w", err,
		)
	}

	if currentNode != nil {
		// Check if new node can be merged with left neighbour.
		merged, err := c.mergeWithNextIfPossible(currentNode)
		if err != nil {
			return fmt.Errorf("merging with left adjacent node failed: %w", err)
		}
		if !merged {
			// If the merge happened, the current node points now to an
			// updated, bigger CIDR. We will want to see if that new node
			// can be merged with its new right neighbour.
			// If the merge did not happen, we want to check the situation
			// of our new node with its right neighbour so we need to switch
			// currentNode to the new node.
			currentNode = currentNode.next
		}
	} else {
		// If the currentNode was nil it means that the insert
		// placed a node at a head position. There is no left
		// neighbour and we can start checking if the merge
		// can happen with the right neighbour.
		currentNode = c.head
	}

	// Merge repeatedly with possible adjacent neighbours till
	// no more merges are possible.
	//
	// Example scenario, we have the following CIDRs:
	// * 192.168.0.0/24
	// * 192.168.2.0/23
	// * 192.168.4.0/22
	//
	// If we insert 192.168.1.0/24 it will result in possible merge
	// of 192.168.0.0/24 and 192.168.1.0/24 into 192.168.0.0/23.
	// The new CIDR gives an opportunity to merge with 192.168.2.0/23
	// producing CIDR 192.168.0.0/22 which can be merged with
	// 192.168.4.0/22 resulting in one CIDR 192.168.0.0/21.
	//
	// This is why, the process may be repeated multiple times - if
	// any merge occurs it means that another merge opportunity could
	// appear.
	//
	// Thanks to that, we avoid CIDR fragmentation. Otherwise, checking
	// if CIDR 192.168.0.0/21 can be lend would be troublesome.
	mergeOccured := true
	for mergeOccured {
		currentNode, mergeOccured, err = c.mergeWithNeighboursIfPossible(
			currentNode,
		)
		if err != nil {
			return fmt.Errorf(
				"failed during an attempt to merge with neighbours: %w", err,
			)
		}
	}

	return nil
}

// Checks if merge can happen between node and its neighbours.
//
// Returns the current node (if the merging with left neighbour
// happens then the current node is removed and the left
// neighbour is the one that matters).
func (c *cidrList) mergeWithNeighboursIfPossible(node *cidrNode) (*cidrNode, bool, error) {
	if c == nil {
		return node, false, nil
	}
	prev := node.Prev()
	mergedWithLeft, err := c.mergeWithNextIfPossible(prev)
	if err != nil {
		return node, false, fmt.Errorf(
			"failed to merge with left neighbour: %w", err,
		)
	}
	if mergedWithLeft {
		node = prev
	}
	mergedWithRight, err := c.mergeWithNextIfPossible(node)
	if err != nil {
		return node, false, fmt.Errorf(
			"failed to merge with left neighbour: %w", err,
		)
	}
	return node, mergedWithLeft || mergedWithRight, nil
}

// mergeWithNextIfPossible checks if the provided node and
// the next node are adjacent and if can be merged.
//
// Returns true if merging happened.
func (c *cidrList) mergeWithNextIfPossible(node *cidrNode) (bool, error) {
	if node == nil || node.next == nil {
		return false, nil
	}
	toMerge, err := AreCIDRsAdjacentAndMergeable(node.cidr, node.Next().cidr)
	if err != nil {
		return false, fmt.Errorf(
			"failed verification and potential merge of diplicates from the CIDR list: %w",
			err)
	}
	if !toMerge {
		return false, nil
	}
	node.RemoveNext()
	node.cidr, err = GrowCIDRByOne(node.cidr)
	if err != nil {
		return false, fmt.Errorf("failed to merge adjacent CIDRs: %w", err)
	}
	return true, nil
}

// insert inserts the cidr in the list of nodes to keep the
// ascending order.
//
// The insert method simply puts the new cidr node before the
// first node which has equal or bigger CIDR IP.
//
// Returns the inserted node.
func (c *cidrList) insert(cidr net.IPNet) (*cidrNode, error) {
	if c == nil {
		return nil, errors.New("cannot insert to non existing list")
	}

	currentNode := c.head
	if currentNode == nil {
		c.head = &cidrNode{
			cidr: cidr,
		}
		return c.head, nil
	}

	if compareIPs(cidr.IP, currentNode.cidr.IP) == LesserThan {
		c.head = &cidrNode{
			cidr: cidr,
			next: currentNode,
		}
		return c.head, nil
	}

	for currentNode.Next() != nil {
		if compareIPs(cidr.IP, currentNode.Next().cidr.IP) == LesserThan {
			break
		}
		currentNode = currentNode.Next()
	}

	currentNode.InsertAfter(cidr)
	return currentNode.Next(), nil
}

// Returns true if the CIDR address is entirely
// covered by the list.
func (c *cidrList) isCovered(cidr net.IPNet) bool {
	if c == nil {
		return false
	}
	currentNode := c.head
	for currentNode != nil {
		if CIDRContainsCIDR(currentNode.cidr, cidr) {
			return true
		}
		currentNode = currentNode.Next()
	}
	return false
}

// RemoveSubsetsOf removes nodes that are a subset of
// a given CIDR.
func (c *cidrList) removeSubsetsOf(cidr net.IPNet) error {
	if c == nil {
		return errors.New("cannot remove anything from nil CIDR List")
	}

	currentNode := c.head

	// Go to first overlapping CIDR.
	for currentNode != nil && !CIDRContainsCIDR(cidr, currentNode.cidr) {
		currentNode = currentNode.Next()
	}

	if currentNode == nil {
		// Overlapping CIDRs not found. Nothing to do.
		return nil
	}

	lastNonOverlappingNode := currentNode.Prev()
	// Go through all overlapping nodes.
	for currentNode != nil && CIDRContainsCIDR(cidr, currentNode.cidr) {
		currentNode = currentNode.Next()
	}

	// The list first node was already overlapping - we need
	// to set a new head (if we reached the end of the list, we
	// will set nil indicating that there are no more entries in
	// the list).
	if lastNonOverlappingNode == nil {
		c.head = currentNode
		c.head.prev = nil
		return nil
	}

	lastNonOverlappingNode.next = currentNode
	lastNonOverlappingNode.next.prev = lastNonOverlappingNode
	return nil
}

type cidrNode struct {
	cidr net.IPNet
	next *cidrNode
	prev *cidrNode
}

func (c *cidrNode) InsertAfter(cidr net.IPNet) {
	if c == nil {
		return
	}
	currentNext := c.next
	c.next = &cidrNode{
		cidr: cidr,
		next: currentNext,
		prev: c,
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
	c.next.prev = c
}

func (c *cidrNode) Next() *cidrNode {
	if c == nil {
		return nil
	}
	return c.next
}

func (c *cidrNode) Prev() *cidrNode {
	if c == nil {
		return nil
	}
	return c.prev
}
