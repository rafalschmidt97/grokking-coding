package main

import (
	"fmt"
	"testing"
)

func TestMinimizeMaximumOfTwoArrays(t *testing.T) {
	tests := []struct {
		uniqueCnt1 int
		divisor1   int
		uniqueCnt2 int
		divisor2   int
		expected   int
	}{
		{
			uniqueCnt1: 2,
			divisor1:   2,
			uniqueCnt2: 2,
			divisor2:   3,
			expected:   4,
		},
		{
			uniqueCnt1: 3,
			divisor1:   3,
			uniqueCnt2: 4,
			divisor2:   4,
			expected:   7,
		},
		{
			uniqueCnt1: 1,
			divisor1:   7,
			uniqueCnt2: 1,
			divisor2:   10,
			expected:   2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.divisor1, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minimizeMaximum(tt.divisor1, tt.divisor2, tt.uniqueCnt1, tt.uniqueCnt2)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
