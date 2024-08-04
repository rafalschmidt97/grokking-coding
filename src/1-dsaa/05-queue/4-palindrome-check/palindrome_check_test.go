package main

import (
	"fmt"
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "madam",
			expected: true,
		},
		{
			input:    "openai",
			expected: false,
		},
		{
			input:    "A man a plan a canal Panama",
			expected: true,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := palindromeCheck(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
