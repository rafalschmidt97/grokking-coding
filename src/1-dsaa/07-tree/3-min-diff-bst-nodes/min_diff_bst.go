package main

import (
	"fmt"
	"math"
)

// A binary search tree has an property where if you perform an in-order traversal,
// the result is a list of numbers in ascending order.
//
// Time complexity: O(n)
// Space complexity: O(n)

type Solution struct {
	Nodes []int
}

func minimumDifferenceBetweenBst(input *TreeNode) int {
	solution := Solution{
		Nodes: make([]int, 0),
	}

	inorderTraversal(input, &solution)

	// Initialize the minimum difference to the largest possible value.
	minDiff := math.MaxInt32

	// Loop through the list and find the difference between each consecutive pair.
	for i := 1; i < len(solution.Nodes); i++ {
		minDiff = min(minDiff, solution.Nodes[i]-solution.Nodes[i-1])
	}

	return minDiff
}

func inorderTraversal(input *TreeNode, solution *Solution) {
	if input == nil {
		return
	}
	inorderTraversal(input.Left, solution)
	solution.Nodes = append(solution.Nodes, input.Value)
	inorderTraversal(input.Right, solution)
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Value: 3,
			},
			Value: 2,
		},
		Right: &TreeNode{
			Value: 6,
		},
		Value: 4,
	}
	fmt.Println(minimumDifferenceBetweenBst(input))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
