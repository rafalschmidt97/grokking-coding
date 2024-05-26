package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseListNode(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				3, 5, 2,
			},
			expected: []int{
				2, 5, 3,
			},
		},
		{
			input: []int{
				7,
			},
			expected: []int{
				7,
			},
		},
		{
			input: []int{
				-1, 0, 1,
			},
			expected: []int{
				1, 0, -1,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := reverse(arrayToListNode(tt.input))
			if !reflect.DeepEqual(listNodeToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
