package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSubarray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		k        int
	}{
		{
			input:    []int{1, 2, 3, 1, 4, 5, 2, 3, 6},
			k:        3,
			expected: []int{3, 3, 4, 5, 5, 5, 6},
		},
		{
			input:    []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13},
			k:        4,
			expected: []int{10, 10, 10, 15, 15, 90, 90},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 4, 5},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := maxOfSubarray(tt.input, tt.k)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
