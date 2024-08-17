package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinimumNumberOfVerticesToReachAll(t *testing.T) {
	tests := []struct {
		edges    [][]int
		expected []int
		nodes    int
	}{
		{
			edges: [][]int{
				{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
			},
			nodes:    6,
			expected: []int{0, 3},
		},
		{
			edges: [][]int{
				{0, 1}, {2, 1},
			},
			nodes:    3,
			expected: []int{0, 2},
		},
		{
			edges: [][]int{
				{0, 1}, {2, 1}, {3, 4},
			},
			nodes:    5,
			expected: []int{0, 2, 3},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.edges, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := verticesToReach(tt.nodes, tt.edges)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
