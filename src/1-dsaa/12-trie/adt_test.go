package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdtTrie(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected []int
	}{
		{
			input: [][]string{
				{"insert", "apple"},
				{"search", "apple"},
				{"startsWith", "app"},
			},
			expected: []int{-1, 1, 1},
		},
		{
			input: [][]string{
				{"insert", "banana"},
				{"search", "apple"},
				{"startsWith", "ban"},
				{"search", "banana"},
			},
			expected: []int{-1, 0, 1, 1},
		},
		{
			input: [][]string{
				{"insert", "grape"},
				{"search", "grape"},
				{"startsWith", "grap"},
				{"startsWith", "gr"},
			},
			expected: []int{-1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := processTrieOperations(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
