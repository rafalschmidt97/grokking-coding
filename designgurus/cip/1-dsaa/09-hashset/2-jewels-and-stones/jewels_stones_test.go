package main

import (
	"fmt"
	"testing"
)

func TestJewelsStones(t *testing.T) {
	tests := []struct {
		jewels   string
		stones   string
		expected int
	}{
		{
			jewels:   "abc",
			stones:   "aabbcc",
			expected: 6,
		},
		{
			jewels:   "aA",
			stones:   "aAaZz",
			expected: 3,
		},
		{
			jewels:   "zZ",
			stones:   "zZzZzZ",
			expected: 6,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.jewels, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := jewelsStones(tt.jewels, tt.stones)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
