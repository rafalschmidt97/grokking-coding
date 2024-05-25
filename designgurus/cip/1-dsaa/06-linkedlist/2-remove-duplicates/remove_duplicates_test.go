package main

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
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
			got := removeDuplicates(arrayToLinkedList(tt.input))
			if !reflect.DeepEqual(linkedListToArray(got), tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}

func linkedListToArray(input *list.List) []int {
	array := []int{}
	for e := input.Front(); e != nil; e = e.Next() {
		array = append(array, e.Value.(int))
	}
	return array
}
