package main

import (
	"fmt"
	"testing"
)

func TestFirstNonRepeatingChar(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "apple",
			expected: 0,
		},
		{
			input:    "abcab",
			expected: 2,
		},
		{
			input:    "abab",
			expected: -1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := firstNonRepeatingChar(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
