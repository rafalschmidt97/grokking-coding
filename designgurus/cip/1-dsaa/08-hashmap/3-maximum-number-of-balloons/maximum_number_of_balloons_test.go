package main

import (
	"fmt"
	"testing"
)

func TestMaximumNumberOfBallons(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "balloonballoon",
			expected: 2,
		},
		{
			input:    "bbaall",
			expected: 0,
		},
		{
			input:    "balloonballoooon",
			expected: 2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maximumNumberOfBallons(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
