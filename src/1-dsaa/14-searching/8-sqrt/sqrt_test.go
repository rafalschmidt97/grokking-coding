package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			input:    8,
			expected: 2,
		},
		{
			input:    4,
			expected: 2,
		},
		{
			input:    2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := sqrt(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}