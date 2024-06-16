package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func closestBstValue(_ *TreeNode, target float64) int {
	return 0
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
