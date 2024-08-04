package main

import (
	"fmt"
	"math"
)

type Solution struct {
	closest int
}

// Time complexity: O(n)
// Space complexity: O(H)
func closestBstValue(input *TreeNode, target float64) int {
	solution := Solution{
		closest: math.MinInt32,
	}

	inorderTraversalWithSearchForClosest(input, target, &solution)
	return solution.closest
}

func inorderTraversalWithSearchForClosest(input *TreeNode, target float64, solution *Solution) {
	if input == nil {
		return
	}

	inorderTraversalWithSearchForClosest(input.Left, target, solution)

	diffCurrentClosest := math.Abs(math.Abs(float64(solution.closest)) - target)
	diffNextClosest := math.Abs(math.Abs(float64(input.Value)) - target)

	if diffNextClosest < diffCurrentClosest {
		solution.closest = input.Value
	}

	inorderTraversalWithSearchForClosest(input.Right, target, solution)

	diffCurrentClosest = math.Abs(math.Abs(float64(solution.closest)) - target)
	diffNextClosest = math.Abs(math.Abs(float64(input.Value)) - target)

	if diffNextClosest < diffCurrentClosest {
		solution.closest = input.Value
	}
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Value: 4,
			},
			Value: 3,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 9,
			},
			Value: 8,
		},
		Value: 5,
	}
	target := 6.4
	fmt.Println(closestBstValue(input, target))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
