package main

import (
	"fmt"
	"testing"
)

func TestBalancedParantheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "{[()]}",
			expected: true,
		},
		{
			input:    "{[}]",
			expected: false,
		},
		{
			input:    "(]",
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := verifyBalancedParentheses(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
