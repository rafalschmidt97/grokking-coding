package main

import (
	"fmt"
	"testing"
)

func TestFrequencyOfMostFrequentElement(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{
			input: []int{
				1, 2, 3,
			},
			k:        3,
			expected: 3,
		},
		{
			input: []int{
				4, 4, 4, 1,
			},
			k:        2,
			expected: 3,
		},
		{
			input: []int{
				10, 9, 9, 4, 8,
			},
			k:        4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := frequencyOfMostFrequentElement(tt.input, tt.k)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
