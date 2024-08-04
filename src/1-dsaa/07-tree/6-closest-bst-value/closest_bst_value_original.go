package main

import (
	"math"
)

// Time complexity: O(logn)
// Space complexity: O(1)
func closestBstValueOriginal(input *TreeNode, target float64) int {
	solution := Solution{
		closest: math.MinInt32,
	}

	for input != nil {
		diffNextClosest := target - float64(input.Value)
		diffCurrentClosest := target - float64(solution.closest)

		// Check if the current node's value is closer to the target than the previous closest value.
		// If so, update closestVal.
		if math.Abs(diffNextClosest) < math.Abs(diffCurrentClosest) {
			solution.closest = input.Value
		}

		// Decide the direction to traverse.
		// If the target is less than the current node's value, we move left; otherwise, move right.
		// This decision is made based on the properties of a BST.
		//
		// if diffNextClosest < 0 {
		if target < float64(input.Value) {
			input = input.Left
		} else {
			input = input.Right
		}
	}

	return solution.closest
}
