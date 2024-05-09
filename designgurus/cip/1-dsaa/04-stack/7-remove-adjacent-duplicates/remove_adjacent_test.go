package main

import (
	"fmt"
	"testing"
)

func TestRemoveAdjacentDuplicates(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "abbaca",
			expected: "ca",
		},
		{
			input:    "azxxzy",
			expected: "ey",
		},
		{
			input:    "abba",
			expected: "",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := removeAdjacent(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
