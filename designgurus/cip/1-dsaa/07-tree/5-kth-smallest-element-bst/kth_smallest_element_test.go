package main

import (
	"fmt"
	"testing"
)

func TestKthSmallestElementBst(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		k        int
		expected int
	}{
		{
			input: &TreeNode{
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
			},
			k:        4,
			expected: 6,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Value: 1,
					},
					Value: 2,
				},
				Right: &TreeNode{
					Value: 6,
				},
				Value: 5,
			},
			k:        3,
			expected: 5,
		},
		{
			input: &TreeNode{
				Right: &TreeNode{
					Left: &TreeNode{
						Value: 2,
					},
					Value: 3,
				},
				Value: 1,
			},
			k:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := kthSmallestElementBst(tt.input, tt.k)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
