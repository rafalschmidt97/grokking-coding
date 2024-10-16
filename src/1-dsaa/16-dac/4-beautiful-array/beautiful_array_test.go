package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBeautifulArray(t *testing.T) {
	tests := []struct {
		expected []int
		input    int
	}{
		{
			input:    4,
			expected: []int{2, 1, 4, 3},
		},
		{
			input:    3,
			expected: []int{1, 3, 2},
		},
		{
			input:    8,
			expected: []int{1, 5, 3, 7, 2, 6, 4, 8},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := beautifulArray(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
