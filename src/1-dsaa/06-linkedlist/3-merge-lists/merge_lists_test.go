package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeLists(t *testing.T) {
	tests := []struct {
		inputOne []int
		inputTwo []int
		expected []int
	}{
		{
			inputOne: []int{
				1, 3, 5,
			},
			inputTwo: []int{
				2, 4, 6,
			},
			expected: []int{
				1, 2, 3, 4, 5, 6,
			},
		},
		{
			inputOne: []int{
				2, 4, 6,
			},
			inputTwo: []int{
				1, 3, 5,
			},
			expected: []int{
				1, 2, 3, 4, 5, 6,
			},
		},
		{
			inputOne: []int{
				1, 2, 3,
			},
			inputTwo: []int{
				4, 5, 6,
			},
			expected: []int{
				1, 2, 3, 4, 5, 6,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.inputOne, tt.inputTwo, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := mergeLists(arrayToListNode(tt.inputOne), arrayToListNode(tt.inputTwo))
			if !reflect.DeepEqual(listNodeToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
