package main

import (
	"fmt"
)

// Time complexity: O(min(N, M))
// Space complexity: O(min(N, M)):
func mergeTwoBsts(inputOne *TreeNode, inputTwo *TreeNode) *TreeNode {
	if inputOne == nil {
		return inputTwo
	}

	if inputTwo == nil {
		return inputOne
	}

	newNode := &TreeNode{Value: inputOne.Value + inputTwo.Value}
	newNode.Left = mergeTwoBsts(inputOne.Left, inputTwo.Left)
	newNode.Right = mergeTwoBsts(inputOne.Right, inputTwo.Right)

	return newNode
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
