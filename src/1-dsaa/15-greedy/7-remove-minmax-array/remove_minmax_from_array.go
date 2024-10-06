package main

import (
	"fmt"
)

// Time complexity: O(n)
// Space complexity: O(1)
func removeMinMaxFromArray(input []int) int {
	n := len(input)

	minIndex, minVal, maxIndex, maxVal := 0, input[0], 0, input[0]

	// Find the indexes of the minimum and maximum elements
	for i, val := range input {
		if val < minVal {
			minIndex, minVal = i, val
		}
		if val > maxVal {
			maxIndex, maxVal = i, val
		}
	}

	// Calculate distances from both ends
	minDistStart := minIndex + 1
	minDistEnd := n - minIndex
	maxDistStart := maxIndex + 1
	maxDistEnd := n - maxIndex

	// Determine the most efficient sequence of moves
	totalMoves := min(
		max(minDistStart, maxDistStart),
		min(
			min(minDistStart+maxDistEnd, minDistEnd+maxDistStart),
			max(minDistEnd, maxDistEnd)))

	return totalMoves
}

func main() {
	input := []int{3, 2, 5, 1, 4}
	fmt.Println(removeMinMaxFromArray(input))
}
