package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndexPairsOfString(t *testing.T) {
	tests := []struct {
		text     string
		words    []string
		expected [][]int
	}{
		{
			text: "bluebirdskyscraper",
			words: []string{
				"blue", "bird", "sky",
			},
			expected: [][]int{
				{0, 3},
				{4, 7},
				{8, 10},
			},
		},
		{
			text: "programmingisfun",
			words: []string{
				"pro", "is", "fun", "gram",
			},
			expected: [][]int{
				{0, 2},
				{3, 6},
				{11, 12},
				{13, 15},
			},
		},
		{
			text: "interstellar",
			words: []string{
				"stellar", "star", "inter",
			},
			expected: [][]int{
				{0, 4},
				{5, 11},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.text, tt.words, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := indexPairsOfString(tt.text, tt.words)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
