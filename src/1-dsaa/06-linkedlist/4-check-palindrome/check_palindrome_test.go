package main

import (
	"fmt"
	"testing"
)

func TestMergeLists(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input: []int{
				1, 2, 3, 2, 1,
			},
			expected: true,
		},
		{
			input: []int{
				1, 2, 2, 3,
			},
			expected: false,
		},
		{
			input: []int{
				1, 2, 2,
			},
			expected: false,
		},
		{
			input: []int{
				1, 2,
			},
			expected: false,
		},
		{
			input: []int{
				1, 1, 1, 1,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := checkPalindrome(arrayToListNode(tt.input))
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
