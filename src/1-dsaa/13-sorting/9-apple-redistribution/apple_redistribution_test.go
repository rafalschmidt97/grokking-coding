package main

import (
	"fmt"
	"testing"
)

func TestAppleRedistribution(t *testing.T) {
	tests := []struct {
		input    []int
		capacity []int
		expected int
	}{
		{
			input: []int{
				2, 3, 1,
			},
			capacity: []int{
				4, 2, 5, 1,
			},
			expected: 2,
		},
		{
			input: []int{
				4, 5, 6,
			},
			capacity: []int{
				5, 10,
			},
			expected: 2,
		},
		{
			input: []int{
				1, 2, 5, 6,
			},
			capacity: []int{
				2, 3, 7, 4, 5, 2, 4,
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := distributeApples(tt.input, tt.capacity)

			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
