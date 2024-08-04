package main

import (
	"fmt"
	"testing"
)

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{5, 2, 3},
				{0, 6, 7},
			},
			expected: 13,
		},
		{
			input: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			expected: 11,
		},
		{
			input: [][]int{
				{10, 100},
				{200, 20},
				{30, 300},
			},
			expected: 330,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maximumWealth(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
