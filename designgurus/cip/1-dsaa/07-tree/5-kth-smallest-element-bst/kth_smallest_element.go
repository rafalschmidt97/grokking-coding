package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func kthSmallestElementBst(_ *TreeNode, k int) int {
	return 0
}

func main() {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Value: 1,
			},
			Right: &TreeNode{
				Left: &TreeNode{
					Value: 4,
				},
				Right: &TreeNode{
					Value: 7,
				},
				Value: 6,
			},
			Value: 3,
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Left: &TreeNode{
					Value: 13,
				},
				Value: 14,
			},
			Value: 10,
		},
		Value: 8,
	}
	k := 4
	fmt.Println(kthSmallestElementBst(input, k))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
