package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMedianOfNumberStream(t *testing.T) {
	tests := []struct {
		input    []int
		expected []float64
	}{
		{
			input: []int{
				3, 1, -1, 5, -1, 4, -1,
			},
			expected: []float64{
				2, 3, 3.5,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := findMedian(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
