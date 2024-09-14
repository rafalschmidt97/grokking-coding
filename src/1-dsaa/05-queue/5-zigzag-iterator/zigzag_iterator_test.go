package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZizagIterator(t *testing.T) {
	tests := []struct {
		input    [2][]int
		expected []int
	}{
		{
			input: [2][]int{
				{1, 2},
				{3, 4, 5, 6},
			},
			expected: []int{1, 3, 2, 4, 5, 6},
		},
		{
			input: [2][]int{
				{1, 2, 3, 4},
				{5, 6},
			},
			expected: []int{1, 5, 2, 6, 3, 4},
		},
		{
			input: [2][]int{
				{1, 2},
				{},
			},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := zigzagIteratorContainers(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
