package main

import (
	"sort"
)

func distributeApplesOriginal(apples []int, boxCapacities []int) int {
	// Sort the box capacities in ascending order to use the largest boxes last
	sort.Ints(boxCapacities)
	totalApples := 0

	// Calculate the total number of apples
	for _, count := range apples {
		totalApples += count
	}

	// Start from the largest box and subtract its capacity from the total apples until all apples are accounted for
	i := len(boxCapacities) - 1
	for i >= 0 && totalApples > 0 {
		totalApples -= boxCapacities[i]
		i--
	}

	// The number of boxes used is the difference between the total number of boxes and the remaining boxes
	return len(boxCapacities) - i - 1
}
