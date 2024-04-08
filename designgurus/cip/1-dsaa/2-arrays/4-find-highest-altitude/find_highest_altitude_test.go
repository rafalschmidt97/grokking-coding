package main

import (
	"fmt"
	"testing"
)

func TestTemplateCases(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{-5, 1, 5, 0, -7}, expected: 1},
		{input: []int{4, -3, 2, -1, -2}, expected: 4},
		{input: []int{2, 2, -3, -1, 2, 1, -5}, expected: 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := findHighestAltitude(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
