package main

import (
	"fmt"
	"testing"
)

func TestSortCharactersByFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "apple",
			expected: []string{"ppale", "ppela"},
		},
		{
			input:    "banana",
			expected: []string{"aaannb"},
		},
		{
			input:    "aabb",
			expected: []string{"aabb", "bbaa"},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := sortByFrequency(tt.input)
			for _, expected := range tt.expected { // if any
				if got == expected {
					return // t.Ok does not exists
				}
			}
			t.Errorf("got %v; expected one of %v", got, tt.expected)
		})
	}
}
