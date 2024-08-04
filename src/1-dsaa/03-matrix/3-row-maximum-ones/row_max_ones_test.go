package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxWithOnes(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			input: [][]int{
				{1, 0},
				{1, 1},
				{0, 1},
			},
			expected: []int{1, 2},
		},
		{
			input: [][]int{
				{0, 1, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
			expected: []int{2, 3},
		},
		{
			input: [][]int{
				{1, 0, 1},
				{0, 0, 1},
				{1, 1, 0},
			},
			expected: []int{0, 2},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maxWithOnes(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
