package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDifferenceArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 5, 1, 6}, expected: []int{12, 5, 1, 8}},
		{input: []int{3, 3, 3}, expected: []int{6, 0, 6}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{14, 11, 6, 1, 10}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := findDifferenceArray(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v; input %v", got, tt.expected, tt.input)
			}
		})
	}
}
