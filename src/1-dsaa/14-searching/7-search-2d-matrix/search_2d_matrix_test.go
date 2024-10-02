package main

import (
	"fmt"
	"testing"
)

func TestSearch2dMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			target: 5,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: true,
		},
		{
			target: 19,
			matrix: [][]int{
				{10, 11, 12, 13},
				{11, 12, 13, 17},
				{14, 19, 22, 24},
				{22, 23, 24, 25},
			},
			expected: true,
		},
		{
			target: 6,
			matrix: [][]int{
				{1, 3, 5},
				{7, 9, 11},
				{13, 15, 17},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.matrix, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := search2dMatrix(tt.matrix, tt.target)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
