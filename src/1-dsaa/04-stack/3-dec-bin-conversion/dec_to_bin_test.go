package main

import (
	"fmt"
	"testing"
)

func TestDecToBin(t *testing.T) {
	tests := []struct {
		expected string
		input    int
	}{
		{input: 2, expected: "10"},
		{input: 7, expected: "111"},
		{input: 18, expected: "10010"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := decToBin(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
