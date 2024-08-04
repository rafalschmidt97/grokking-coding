package main

import (
	"fmt"
	"testing"
)

func TestLargestUniqueNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				5, 7, 3, 7, 5, 8,
			},
			expected: 8,
		},
		{
			input: []int{
				1, 2, 3, 2, 1, 4, 4,
			},
			expected: 3,
		},
		{
			input: []int{
				9, 9, 8, 8, 7, 7,
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := largestUniqueNumber(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
