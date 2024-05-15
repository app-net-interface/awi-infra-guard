package accesscontrol

import (
	"crypto/sha256"
	"fmt"
	"slices"
	"strings"
)

// ruleName is unexported string created to
// enforce using exported functions from the
// package when setting up names for Access
// Control resources outside of this package.
type ruleName string

// RuleNames is an exported slice of ruleName used
// mainly to specify rules that should be removed.
type RuleNames = []ruleName

// VPCRuleName generates proper name identifier
// based on source and destination VPCs. The VPC
// rule acts bidirectional and so the order of
// VPC names will be picked by the function
// (names are sorted to keep it deterministic).
//
// TODO: Currently the name is a hash of vpc IDs,
// to keep the length of generated name fixed and
// not over accepted Azure limits, however it is
// not collision-proof. Name collision must be
// handled properly.
func VPCRuleName(vpcId1, vpcId2 string) ruleName {
	ids := []string{vpcId1, vpcId2}
	slices.Sort(ids)

	hasher := sha256.New()
	hasher.Write([]byte(strings.Join(ids, ":")))
	hashBytes := hasher.Sum(nil)

	return ruleName(fmt.Sprintf("%x", hashBytes))
}

// CustomRuleName accepts a regular name provided
// by the external entity and hashes it to keep
// the length name consistent.
//
// TODO: Currently the name is a hash of a given
// string, to keep the length of generated nam
// fixed and not over accepted Azure limits,
// however it is not collision-proof. Name
// collision must be handled properly.
func CustomRuleName(name string) ruleName {
	hasher := sha256.New()
	hasher.Write([]byte(name))
	hashBytes := hasher.Sum(nil)

	return ruleName(fmt.Sprintf("%x", hashBytes))
}

// nameWithPriority combines Rule name with fixed-length priority string.
// The priority always uses ":" character and 4 digits. For priorities
// lower than 1000, the actual priority is preceeded with 0s to match
// 4 characters length.
//
// The priority acts as a name distinguisher between rules inside the
// same Network Security Group as the name prefix may be equal but
// priority ensures uniqueness.
func nameWithPriority(name ruleName, priority uint) (string, error) {
	if priority >= 10000 {
		return "", fmt.Errorf(
			"unexpected priority value - expected 4 digits at max: %d", priority,
		)
	}
	return string(name) + fmt.Sprintf(":%04d", priority), nil
}
