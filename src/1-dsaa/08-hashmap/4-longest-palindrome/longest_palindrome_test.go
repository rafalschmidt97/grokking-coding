package main

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "applepie",
			expected: 5,
		},
		{
			input:    "aabbcc",
			expected: 6,
		},
		{
			input:    "bananas",
			expected: 5,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := longestPalindromeOriginal(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
