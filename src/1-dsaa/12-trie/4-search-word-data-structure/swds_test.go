package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchWordDataStructure(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected []int
	}{
		{
			input: [][]string{
				{"addWord", "apple"},
				{"addWord", "banana"},
				{"search", "apple"},
				{"search", "......"},
			},
			expected: []int{-1, -1, 1, 1},
		},
		{
			input: [][]string{
				{"addWord", "cat"},
				{"addWord", "dog"},
				{"search", "c.t"},
				{"search", "d..g"},
			},
			expected: []int{-1, -1, 1, 0},
		},
		{
			input: [][]string{
				{"addWord", "hello"},
				{"search", "h.llo"},
				{"search", "h...o"},
			},
			expected: []int{-1, 1, 1},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := processStructureOperations(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
