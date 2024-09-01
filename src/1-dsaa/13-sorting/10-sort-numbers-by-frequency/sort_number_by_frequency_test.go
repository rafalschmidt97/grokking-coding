package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortNumbersByFrequency(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				4, 4, 6, 2, 2, 2,
			},
			expected: []int{
				6, 4, 4, 2, 2, 2,
			},
		},
		{
			input: []int{
				0, -1, -1, -1, 5,
			},
			expected: []int{
				5, 0, -1, -1, -1,
			},
		},
		{
			input: []int{
				10, 10, 10, 20, 20, 30,
			},
			expected: []int{
				30, 20, 20, 10, 10, 10,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			sortByFrequency(tt.input)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v; expected %v", tt.input, tt.expected)
			}
		})
	}
}
