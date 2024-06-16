package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func rangeSumBst(_ *TreeNode, min int, max int) int {
	return 0
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 3,
			},
			Right: &TreeNode{
				Value: 7,
			},
			Value: 5,
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Value: 18,
			},
			Value: 15,
		},
		Value: 10,
	}
	min := 7
	max := 15
	fmt.Println(rangeSumBst(input, min, max))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
