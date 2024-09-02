package main

import (
	"fmt"
	"testing"
)

func TestTakeGiftsFromTheRichestPile(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{
			input: []int{
				4, 9, 16,
			},
			k:        2,
			expected: 11,
		},
		{
			input: []int{
				1, 2, 3,
			},
			k:        1,
			expected: 4,
		},
		{
			input: []int{
				25, 36, 49,
			},
			k:        3,
			expected: 18,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := takeGiftsFromRichestMaxContainers(tt.input, tt.k)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
