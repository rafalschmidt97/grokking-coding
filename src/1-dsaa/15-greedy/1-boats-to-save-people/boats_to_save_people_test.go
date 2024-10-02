package main

import (
	"fmt"
	"testing"
)

func TestBodyToSavePeople(t *testing.T) {
	tests := []struct {
		people   []int
		limit    int
		expected int
	}{
		{
			people:   []int{10, 55, 70, 20, 90, 85},
			limit:    100,
			expected: 4,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.people, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := boatsToSavePeople(tt.people, tt.limit)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
