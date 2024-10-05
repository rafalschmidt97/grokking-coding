package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "babac",
			expected: "abc",
		},
		{
			input:    "zabccde",
			expected: "zabcde",
		},
		{
			input:    "mnopmn",
			expected: "mnop",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := removeDuplicateLetters(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
