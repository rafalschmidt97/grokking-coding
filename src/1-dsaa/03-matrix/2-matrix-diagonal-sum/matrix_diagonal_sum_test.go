package main

import (
	"fmt"
	"testing"
)

func TestMatrixDiagonalSum(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 25,
		},
		{
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{5},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := matrixDiagonalSum(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
