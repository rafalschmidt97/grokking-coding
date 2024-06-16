package main

import (
	"fmt"
	"testing"
)

func TestRangeSumBst(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		min      int
		max      int
		expected int
	}{
		{
			input: &TreeNode{
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
			},
			min:      7,
			max:      15,
			expected: 32,
		},
		{
			input: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Value: 3,
					},
					Right: &TreeNode{
						Value: 10,
					},
					Value: 5,
				},
				Right: &TreeNode{
					Value: 25,
				},
				Value: 20,
			},
			min:      3,
			max:      10,
			expected: 18,
		},
		{
			input: &TreeNode{
				Right: &TreeNode{
					Left: &TreeNode{
						Value: 32,
					},
					Value: 35,
				},
				Value: 30,
			},
			min:      30,
			max:      34,
			expected: 62,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := rangeSumBst(tt.input, tt.min, tt.max)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
