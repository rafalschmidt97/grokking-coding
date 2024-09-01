package main

import (
	"fmt"
	"testing"
)

func TestSortVowelsInString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "gamE",
			expected: "gEma",
		},
		{
			input:    "aEiOu",
			expected: "EOaiu",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := sortVowelsOriginal(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
