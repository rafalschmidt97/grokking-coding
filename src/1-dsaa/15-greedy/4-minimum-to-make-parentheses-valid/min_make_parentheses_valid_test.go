package main

import (
	"fmt"
	"testing"
)

func TestMinimumAddToMakeParenthesesValid(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "(()",
			expected: 1,
		},
		{
			input:    "))((",
			expected: 4,
		},
		{
			input:    "(()())(",
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minMakeParenthesesValidStack(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
