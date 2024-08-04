package main

import (
	"fmt"
	"testing"
)

func TestMinimumDifferenceBetweenBst(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
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
			},
			expected: 1,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Value: 2,
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
			},
			expected: 2,
		},
		{
			input: &TreeNode{
				Right: &TreeNode{
					Left: &TreeNode{
						Value: 50,
					},
					Right: &TreeNode{
						Value: 90,
					},
					Value: 70,
				},
				Value: 40,
			},
			expected: 10,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minimumDifferenceBetweenBst(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
