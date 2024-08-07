package main

import (
	"fmt"
	"testing"
)

func TestTemplateCases(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{2, 3, 4},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := template(tt.input)
			// template(tt.input)

			// if !reflect.DeepEqual(got, tt.expected) {
			// if !reflect.DeepEqual(tt.input, tt.expected) {
			// if tt.input != tt.expected {
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
				// t.Errorf("got %v; expected %v", tt.input, tt.expected)
			}
		})
	}
}
