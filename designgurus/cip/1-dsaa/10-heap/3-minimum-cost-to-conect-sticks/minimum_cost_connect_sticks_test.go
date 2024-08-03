package main

import (
	"fmt"
	"testing"
)

func TestMinimumCostToConnectSticks(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				2, 4, 3,
			},
			expected: 14,
		},
		{
			input: []int{
				1, 8, 2, 5,
			},
			expected: 27,
		},
		{
			input: []int{
				5, 5, 5, 5,
			},
			expected: 40,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minimumCostToConnectSticks(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
