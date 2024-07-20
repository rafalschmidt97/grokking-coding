package main

import (
	"math"
)

// Time complexity: O(logn)
// Space complexity: O(H)
func closestBstValueV2(input *TreeNode, target float64) int {
	solution := Solution{
		closest: math.MinInt32,
	}

	inorderTraversalWithSearchForClosest(input, target, &solution)
	return solution.closest
}

func inorderTraversalWithSkipSearchForClosest(input *TreeNode, target float64, solution *Solution) {
	if input == nil {
		return
	}

	diffNextClosest := target - float64(input.Value)
	diffCurrentClosest := target - float64(solution.closest)

	if diffNextClosest == 0 {
		return
	}

	if math.Abs(diffNextClosest) < math.Abs(diffCurrentClosest) {
		solution.closest = input.Value
	}

	if diffNextClosest < 0 {
		inorderTraversalWithSkipSearchForClosest(input.Left, target, solution)
	} else {
		inorderTraversalWithSkipSearchForClosest(input.Right, target, solution)
	}
}
