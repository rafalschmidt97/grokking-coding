package main

import (
	"fmt"
	"testing"
)

func TestMaximumLengthOfPairChain(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 2}, {3, 4}, {2, 3}},
			expected: 2,
		},
		{
			input:    [][]int{{5, 6}, {1, 2}, {8, 9}, {2, 3}},
			expected: 3,
		},
		{
			input:    [][]int{{7, 8}, {5, 6}, {1, 2}, {3, 5}, {4, 5}, {2, 3}},
			expected: 3,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maxLenPairChain(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
