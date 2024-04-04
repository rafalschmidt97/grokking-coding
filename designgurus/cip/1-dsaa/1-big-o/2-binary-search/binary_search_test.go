package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array []int
		x     int
		want  int
	}{
		{array: []int{2, 3, 4, 5, 10, 40}, x: 3, want: 1},
		{array: []int{2, 3, 4, 5, 10, 40}, x: 10, want: 4},
		{array: []int{2, 3, 4, 5, 10, 40}, x: 2, want: 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.array, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := binarySearch(tt.array, tt.x)

			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
