package main

import (
	"fmt"
	"testing"
)

func TestZizagIterator(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := next(tt.input)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
