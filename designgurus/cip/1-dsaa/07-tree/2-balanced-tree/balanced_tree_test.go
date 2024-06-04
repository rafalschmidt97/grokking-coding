package main

import (
	"fmt"
	"testing"
)

func TestMaxDepthBinaryTree(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
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
			},
			expected: true,
		},
		// TODO: add more
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := isBinaryTreeBalanced(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
