package main

import (
	"fmt"
	"testing"
)

func TestLargestPalindromicNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "323211444",
			expected: "432141234",
		},
		{
			input:    "998877",
			expected: "987789",
		},
		{
			input:    "54321",
			expected: "5",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := largestPalindromicNumber(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
