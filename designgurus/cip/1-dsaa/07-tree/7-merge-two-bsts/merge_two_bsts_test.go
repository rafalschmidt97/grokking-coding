package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoBst(t *testing.T) {
	tests := []struct {
		inputOne *TreeNode
		inputTwo *TreeNode
		expected *TreeNode
	}{
		{
			inputOne: &TreeNode{
				Left: &TreeNode{
					Value: 3,
				},
				Right: &TreeNode{
					Value: 2,
				},
				Value: 1,
			},
			inputTwo: &TreeNode{
				Left: &TreeNode{
					Value: 2,
				},
				Right: &TreeNode{
					Value: 3,
				},
				Value: 1,
			},
			expected: &TreeNode{
				Left: &TreeNode{
					Value: 5,
				},
				Right: &TreeNode{
					Value: 5,
				},
				Value: 2,
			},
		},
		{
			inputOne: &TreeNode{
				Left: &TreeNode{
					Value: 4,
				},
				Right: &TreeNode{
					Value: 7,
				},
				Value: 5,
			},
			inputTwo: &TreeNode{
				Left: &TreeNode{
					Value: 2,
				},
				Right: &TreeNode{
					Value: 6,
				},
				Value: 3,
			},
			expected: &TreeNode{
				Left: &TreeNode{
					Value: 6,
				},
				Right: &TreeNode{
					Value: 13,
				},
				Value: 8,
			},
		},
		{
			inputOne: &TreeNode{
				Right: &TreeNode{
					Value: 5,
				},
				Value: 2,
			},
			inputTwo: &TreeNode{
				Left: &TreeNode{
					Value: 3,
				},
				Value: 2,
			},
			expected: &TreeNode{
				Left: &TreeNode{
					Value: 3,
				},
				Right: &TreeNode{
					Value: 5,
				},
				Value: 4,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.inputOne, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := mergeTwoBsts(tt.inputOne, tt.inputTwo)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
