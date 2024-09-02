package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindTopKFrequentNumbers(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
		k        int
	}{
		{
			input: []int{
				1, 3, 5, 12, 11, 12, 11,
			},
			k: 2,
			expected: [][]int{
				{12, 11},
			},
		},
		{
			input: []int{
				5, 12, 11, 3, 11,
			},
			k: 2,
			expected: [][]int{
				{11, 5}, {11, 12}, {11, 3},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := findTopKFrequentNumbersContainers(tt.input, tt.k)
			for _, expected := range tt.expected { // if any
				if reflect.DeepEqual(got, expected) {
					return // t.Ok does not exists
				}
			}
			t.Errorf("got %v; expected one of %v", got, tt.expected)
		})
	}
}
