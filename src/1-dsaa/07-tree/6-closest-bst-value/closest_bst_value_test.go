package main

import (
	"fmt"
	"testing"
)

func TestClosestBstValue(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		target   float64
		expected int
	}{
		{
			input: &TreeNode{
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
			},
			target:   6.4,
			expected: 6,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Value: 10,
				},
				Right: &TreeNode{
					Value: 30,
				},
				Value: 20,
			},
			target:   25,
			expected: 20,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Value: 1,
				},
				Right: &TreeNode{
					Value: 3,
				},
				Value: 2,
			},
			target:   2.9,
			expected: 3,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := closestBstValueOriginal(tt.input, tt.target)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
