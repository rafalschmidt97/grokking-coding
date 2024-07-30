package main

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcdaef",
			expected: 6,
		},
		{
			input:    "aaaaa",
			expected: 1,
		},
		{
			input:    "abrkaabcdefghijjxxx",
			expected: 10,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := longestSubstringWithoutRepeatingChar(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
