package main

import (
	"fmt"
	"testing"
)

func TestMaximumCountOfPositiveIntegerAndNegativeInteger(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				-4, -3, -1, 0, 1, 3, 5, 7,
			},
			expected: 4,
		},
		{
			input: []int{
				-8, -7, -5, -4, 0, 0, 0,
			},
			expected: 4,
		},
		{
			input: []int{
				0, 2, 2, 3, 3, 3, 4,
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maximumCountPositiveNegative(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
