package main

import (
	"fmt"
	"testing"
)

func TestMaxDepthBinaryTree(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
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
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Right: &TreeNode{
					Right: &TreeNode{
						Value: 3,
					},
					Value: 2,
				},
				Value: 1,
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Value: 4,
					},
					Right: &TreeNode{
						Right: &TreeNode{
							Value: 9,
						},
						Value: 7,
					},
					Value: 2,
				},
				Right: &TreeNode{
					Value: 3,
				},
				Value: 1,
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maxDepthBinaryTree(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
