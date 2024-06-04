package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func isBinaryTreeBalanced(input *TreeNode) bool {
	return false
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
