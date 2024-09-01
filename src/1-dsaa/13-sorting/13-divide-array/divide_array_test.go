package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDivideArrayIntoArraysWithMaxDifference(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
		k        int
	}{
		{
			input: []int{
				2, 6, 4, 9, 3, 7, 3, 4, 1,
			},
			k: 3,
			expected: [][]int{
				{1, 2, 3}, {3, 4, 4}, {6, 7, 9},
			},
		},
		{
			input: []int{
				10, 12, 15, 20, 25, 30,
			},
			k: 10,
			expected: [][]int{
				{10, 12, 15}, {20, 25, 30},
			},
		},
		{
			input: []int{
				1, 2, 4, 5, 9, 10,
			},
			k:        1,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := divideArray(tt.input, tt.k)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
