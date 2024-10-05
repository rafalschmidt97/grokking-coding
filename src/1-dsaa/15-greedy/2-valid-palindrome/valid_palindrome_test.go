package main

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "racecar",
			expected: true,
		},
		{
			input:    "abccdba",
			expected: true,
		},
		{
			input:    "abcdef",
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := validPalindromeByRemovingOne(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
