package main

import (
	"fmt"
	"testing"
)

func TestNumberOfProvinces(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{1, 0, 0, 1},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 1},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := numberOfProvinces(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
