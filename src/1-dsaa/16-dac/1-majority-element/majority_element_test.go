package main

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				1, 2, 2, 3, 2,
			},
			expected: 2,
		},
		{
			input: []int{
				4, 4, 4, 4, 7, 4, 4,
			},
			expected: 4,
		},
		{
			input: []int{
				9, 9, 1, 1, 9, 1, 9, 9,
			},
			expected: 9,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := majorityElement(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
