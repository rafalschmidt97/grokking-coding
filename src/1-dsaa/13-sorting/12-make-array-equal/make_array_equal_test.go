package main

import (
	"fmt"
	"testing"
)

func TestReductionOperationsToMakeTheArrayElementsEqual(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				3, 5, 5, 2,
			},
			expected: 5,
		},
		{
			input: []int{
				11, 9, 7, 5, 3,
			},
			expected: 10,
		},
		{
			input: []int{
				8, 8, 8, 8,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := makeArrayEqual(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
