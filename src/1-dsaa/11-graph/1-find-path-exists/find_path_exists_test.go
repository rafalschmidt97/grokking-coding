package main

import (
	"fmt"
	"testing"
)

func TestFindIfPathExists(t *testing.T) {
	tests := []struct {
		edges    [][]int
		nodes    int
		start    int
		end      int
		expected bool
	}{
		{
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3},
			},
			nodes:    4,
			start:    0,
			end:      3,
			expected: true,
		},
		{
			edges: [][]int{
				{0, 1}, {2, 3},
			},
			nodes:    4,
			start:    0,
			end:      3,
			expected: false,
		},
		{
			edges: [][]int{
				{0, 1}, {3, 4},
			},
			nodes:    5,
			start:    0,
			end:      4,
			expected: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.edges, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := findIfPathExists(tt.nodes, tt.edges, tt.start, tt.end)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
