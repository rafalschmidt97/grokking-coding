package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				3, 1, 2,
			},
			expected: []int{
				1, 2, 3,
			},
		},
		{
			input: []int{
				4,
			},
			expected: []int{
				4,
			},
		},
		{
			input: []int{
				9, 8, 7, 6, 5, 4, 3, 2, 1,
			},
			expected: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := sortList(arrayToListNode(tt.input))
			if !reflect.DeepEqual(listNodeToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
