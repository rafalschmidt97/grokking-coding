package main

import (
	"fmt"
	"testing"
)

func TestMeetingRooms(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{10, 15},
				{20, 25},
				{30, 35},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{10, 20},
				{15, 25},
				{24, 30},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{10, 20},
				{20, 30},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := meetingsRoom(tt.input)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
