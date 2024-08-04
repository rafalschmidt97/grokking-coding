package main

import "fmt"

// DFS approach
//
// Time complexity: O(n)
// Space complexity: O(d) where d is the depth
func maxDepthBinaryTree(input *TreeNode) int {
	if input == nil {
		return 0 // base case for recursiveness
	}

	leftNodeDepth := maxDepthBinaryTree(input.Left)
	rightNodeDepth := maxDepthBinaryTree(input.Right)

	if leftNodeDepth >= rightNodeDepth {
		return leftNodeDepth + 1 // add one for the current node
	} else {
		return rightNodeDepth + 1
	}
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 4,
			},
			Right: &TreeNode{
				Value: 5,
			},
			Value: 2,
		},
		Right: &TreeNode{
			Value: 3,
		},
		Value: 1,
	}
	fmt.Println(maxDepthBinaryTree(input))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
