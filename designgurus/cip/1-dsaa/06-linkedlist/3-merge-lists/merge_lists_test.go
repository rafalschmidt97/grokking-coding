package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeLists(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				1, 1, 2,
			},
			expected: []int{
				1, 2,
			},
		},
		{
			input: []int{
				1, 2, 2, 3,
			},
			expected: []int{
				1, 2, 3,
			},
		},
		{
			input: []int{
				3, 3, 3,
			},
			expected: []int{
				3,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := mergeLists(arrayToListNode(tt.input))
			if !reflect.DeepEqual(listNodeToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
