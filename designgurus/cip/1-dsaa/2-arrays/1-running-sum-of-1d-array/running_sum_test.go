package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRunningSumOf1dArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 3, 5, 1, 6}, expected: []int{2, 5, 10, 11, 17}},
		{input: []int{1, 1, 1, 1, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{-1, 2, -3, 4, -5}, expected: []int{-1, 1, -2, 2, -3}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			runningSum(tt.input)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v; expected %v", tt.input, tt.expected)
			}
		})
	}
}
