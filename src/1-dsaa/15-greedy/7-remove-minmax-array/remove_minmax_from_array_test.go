package main

import (
	"fmt"
	"testing"
)

func TestRemovingMinimumAndMaximumFromArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 2, 5, 1, 4},
			expected: 3,
		},
		{
			input:    []int{7, 5, 6, 8, 1},
			expected: 2,
		},
		{
			input:    []int{2, 4, 10, 1, 3, 5},
			expected: 4,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := removeMinMaxFromArray(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
