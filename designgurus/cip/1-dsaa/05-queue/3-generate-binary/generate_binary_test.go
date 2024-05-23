package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateBinary(t *testing.T) {
	tests := []struct {
		expected []string
		input    int
	}{
		{
			input:    2,
			expected: []string{"1", "10"},
		},
		{
			input:    3,
			expected: []string{"1", "10", "11"},
		},
		{
			input:    5,
			expected: []string{"1", "10", "11", "100", "101"},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := generateBinaryOriginal(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
