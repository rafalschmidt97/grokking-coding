package main

import (
	"fmt"
	"math"
)

// DFS approach
//
// Time complexity: O(n)
// Space complexity: O(d) where d is the depth
func isBinaryTreeBalanced(input *TreeNode) bool {
	return depth(input) != -1
}

func depth(input *TreeNode) int {
	if input == nil {
		return 0 // base case for recursiveness
	}

	leftNodeDepth := depth(input.Left)
	rightNodeDepth := depth(input.Right)

	if leftNodeDepth == -1 || rightNodeDepth == -1 {
		return -1 // current node is not balanced, populate
	}

	if math.Abs(float64(leftNodeDepth-rightNodeDepth)) > 1 {
		return -1
	}

	// return int(math.Max(float64(leftNodeDepth), float64(rightNodeDepth))) + 1
	if leftNodeDepth >= rightNodeDepth {
		return leftNodeDepth + 1 // add one for the current node
	} else {
		return rightNodeDepth + 1
	}
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Value: 9,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Value: 15,
			},
			Right: &TreeNode{
				Value: 7,
			},
			Value: 20,
		},
		Value: 3,
	}
	fmt.Println(isBinaryTreeBalanced(input))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
