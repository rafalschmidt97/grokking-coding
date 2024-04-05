package main

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 3, 4}, expected: false},
		{input: []int{1, 2, 3, 1}, expected: true},
		{input: []int{3, 2, 1, 1}, expected: true},
		{input: []int{3, 2, 3, 1}, expected: true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := containsDuplicate(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
