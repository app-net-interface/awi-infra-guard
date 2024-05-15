package accesscontrol

import (
	"errors"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/connector/helper"
)

// generatePriorities generates a slice of integers that
// can be used as priorities by AccessControlRuleSet.
//
// It accepts the number of requested VPC Rules, Directed
// Rules and Custom Rules and generates priorities to fit
// already used priorities.
//
// For instance, requesting priorities for 3 VPC Rules,
// 2 Directed Rules and 4 Custom Rules with following
// priorities already in use: [3600, 3602, 100, 200] will
// generate following output:
//
// [3601, 3603, 3604, 2800, 2801, 101, 102, 103]
//
// The order of rules goes as follows:
// 1. VPC Rules
// 2. Directed Rules
// 3. Custom Rules
func generatePriorities(
	prioritiesInUse helper.Set[uint],
	numberOfVPCRules,
	numberOfDirectedRules,
	numberOfCustomRules uint,
) (priorityPool, error) {
	vpcPriorities, err := generatePrioritiesForRuleGroup(
		3600, 4096, numberOfVPCRules, prioritiesInUse,
	)
	if err != nil {
		return priorityPool{}, fmt.Errorf(
			"failed to generate priorities for VPC Rules: %w", err,
		)
	}
	directedPriorities, err := generatePrioritiesForRuleGroup(
		2800, 3599, numberOfDirectedRules, prioritiesInUse,
	)
	if err != nil {
		return priorityPool{}, fmt.Errorf(
			"failed to generate priorities for Directed VPC Rules: %w", err,
		)
	}
	customPriorities, err := generatePrioritiesForRuleGroup(
		100, 2799, numberOfCustomRules, prioritiesInUse,
	)
	if err != nil {
		return priorityPool{}, fmt.Errorf(
			"failed to generate priorities for Custom Rules: %w", err,
		)
	}

	priorities := make([]uint, 0, len(vpcPriorities)+len(directedPriorities)+len(customPriorities))
	priorities = append(priorities, vpcPriorities...)
	priorities = append(priorities, directedPriorities...)
	priorities = append(priorities, customPriorities...)

	return newPriorityPool(priorities), nil
}

func generatePrioritiesForRuleGroup(
	minPriority, maxPriority, numberOfPriorities uint,
	prioritiesInUse helper.Set[uint],
) ([]uint, error) {
	current := minPriority
	priorities := make([]uint, 0, numberOfPriorities)

	for current < maxPriority {
		if len(priorities) == int(numberOfPriorities) {
			break
		}

		if prioritiesInUse.Has(current) {
			continue
		}
		priorities = append(priorities, current)
		current++
	}

	if len(priorities) != int(numberOfPriorities) {
		return nil, fmt.Errorf(
			"could not produce %d priorities as the slots are already taken",
			numberOfPriorities,
		)
	}

	return priorities, nil
}

// priorityPool is a simple structure for picking a next
// available priority without using priority already in use.
type priorityPool struct {
	current             uint
	availablePriorities []uint
}

func newPriorityPool(availablePriorities []uint) priorityPool {
	return priorityPool{
		current:             0,
		availablePriorities: availablePriorities,
	}
}

func (p *priorityPool) Next() (uint, error) {
	if int(p.current) >= len(p.availablePriorities) {
		return 0, errors.New(
			"no more priorities left",
		)
	}
	priority := p.availablePriorities[p.current]
	p.current++
	return priority, nil
}
