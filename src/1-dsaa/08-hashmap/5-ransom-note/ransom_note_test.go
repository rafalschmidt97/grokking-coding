package main

import (
	"fmt"
	"testing"
)

func TestRansomNote(t *testing.T) {
	tests := []struct {
		note     string
		input    string
		expected bool
	}{
		{
			note:     "hello",
			input:    "hellworld",
			expected: true,
		},
		{
			note:     "notes",
			input:    "stoned",
			expected: true,
		},
		{
			note:     "apple",
			input:    "pale",
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := ransomNoteOriginal(tt.note, tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
