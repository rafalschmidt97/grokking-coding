package main

import (
	"fmt"
	"testing"
)

func TestRemovingStarsFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "abc*de*f",
			expected: "abdf",
		},
		{
			input:    "a*b*c*d",
			expected: "d",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := removeStarsContainers(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
