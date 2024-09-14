package main

import (
	"fmt"
	"testing"
)

func TestMakeStringGreat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "AaBbCcDdEeff",
			expected: "ff",
		},
		{
			input:    "abBA",
			expected: "",
		},
		{
			input:    "s",
			expected: "s",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := makeGreatContainers(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
