package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortingStack(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{34, 3, 31, 98, 92, 23},
			expected: []int{3, 23, 31, 34, 92, 98},
		},
		{
			input:    []int{4, 3, 2, 10, 12, 1, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6, 10, 12},
		},
		{
			input:    []int{20, 10, -5, -1},
			expected: []int{-5, -1, 10, 20},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := sortStackContainers(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
