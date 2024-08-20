package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchSuggestions(t *testing.T) {
	tests := []struct {
		search   string
		input    []string
		expected [][]string
	}{
		{
			search: "app",
			input: []string{
				"apple", "apricot", "application",
			},
			expected: [][]string{
				{"apple", "application", "apricot"},
				{"apple", "application", "apricot"},
				{"apple", "application"},
			},
		},
		{
			search: "ki",
			input: []string{
				"king", "kingdom", "kit",
			},
			expected: [][]string{
				{"king", "kingdom", "kit"},
				{"king", "kingdom", "kit"},
			},
		},
		{
			search: "farm",
			input: []string{
				"fantasy", "fast", "festival",
			},
			expected: [][]string{
				{"fantasy", "fast", "festival"},
				{"fantasy", "fast"},
				{},
				{},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.search, tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := searchSugestions(tt.search, tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
