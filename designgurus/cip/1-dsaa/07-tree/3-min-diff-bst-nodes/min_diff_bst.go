package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func minimumDifferenceBetweenBst(input *TreeNode) int {
	return 0
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
