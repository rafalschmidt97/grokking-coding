package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwapNodesPairs(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				1, 2, 3, 4,
			},
			expected: []int{
				2, 1, 4, 3,
			},
		},
		{
			input: []int{
				7, 8, 9, 10, 11,
			},
			expected: []int{
				8, 7, 10, 9, 11,
			},
		},
		{
			input: []int{
				5, 6,
			},
			expected: []int{
				6, 5,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := swapNodesPairs(arrayToListNode(tt.input))
			if !reflect.DeepEqual(listNodeToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
