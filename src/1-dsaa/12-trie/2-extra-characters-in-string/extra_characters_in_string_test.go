package main

import (
	"fmt"
	"testing"
)

func TestExtraCharactersInString(t *testing.T) {
	tests := []struct {
		text     string
		words    []string
		expected int
	}{
		{
			text: "amazingracecar",
			words: []string{
				"race", "car",
			},
			expected: 7,
		},
		{
			text: "bookkeeperreading",
			words: []string{
				"keep", "read",
			},
			expected: 9,
		},
		{
			text: "thedogbarksatnight",
			words: []string{
				"dog", "bark", "night",
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.text, tt.words, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minExtraChar(tt.text, tt.words)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
