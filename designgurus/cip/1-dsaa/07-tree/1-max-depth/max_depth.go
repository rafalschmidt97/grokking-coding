package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func maxDepthBinaryTree(input *TreeNode) int {
	return 1
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