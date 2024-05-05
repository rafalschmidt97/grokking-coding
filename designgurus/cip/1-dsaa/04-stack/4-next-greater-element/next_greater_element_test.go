package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{4, 5, 2, 25},
			expected: []int{5, 25, 25, -1},
		},
		{
			input:    []int{13, 7, 6, 12},
			expected: []int{-1, 12, 12, -1},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5, -1},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := nextGreaterElemenet(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
