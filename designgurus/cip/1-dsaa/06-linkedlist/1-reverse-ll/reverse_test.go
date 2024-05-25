package main

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestReverseLL(t *testing.T) {
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
			input: []int{},
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
			got := reverse(arrayToLinkedList(tt.input))
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
