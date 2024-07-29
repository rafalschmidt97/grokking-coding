package main

import (
	"fmt"
	"testing"
)

func TestCounteringElements(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				4, 3, 1, 5, 6,
			},
			expected: 3,
		},
		{
			input: []int{
				7, 8, 9, 10,
			},
			expected: 3,
		},
		{
			input: []int{
				11, 13, 15, 16,
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := counteringElements(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
