package main

import (
	"fmt"
	"testing"
)

func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "BbCcXxY",
			expected: "BbCcXx",
		},
		{
			input:    "aZAbcD",
			expected: "",
		},
		{
			input:    "qQwWeErR",
			expected: "qQwWeErR",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := longestNiceSubstring(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
