package main

import "fmt"

// TODO:
//
// Time complexity: TODO:
// Space complexity: TODO:
func mergeTwoBsts(inputOne *TreeNode, _ *TreeNode) *TreeNode {
	return inputOne
}

func main() {
	inputOne := &TreeNode{
		Left: &TreeNode{
			Value: 3,
		},
		Right: &TreeNode{
			Value: 2,
		},
		Value: 1,
	}
	inputTwo := &TreeNode{
		Left: &TreeNode{
			Value: 2,
		},
		Right: &TreeNode{
			Value: 3,
		},
		Value: 1,
	}
	fmt.Println(mergeTwoBsts(inputOne, inputTwo))
}

// Boilerplate for tree node

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}
