package main

import (
	"fmt"
	"testing"
)

func TestUniqueNumberOccurances(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input: []int{
				4, 5, 4, 6, 6, 6,
			},
			expected: true,
		},
		{
			input: []int{
				7, 8, 8, 9, 9, 9, 10, 10,
			},
			expected: false,
		},
		{
			input: []int{
				11, 12, 13, 14, 14, 15, 15, 15,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := uniqueNumbers(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
