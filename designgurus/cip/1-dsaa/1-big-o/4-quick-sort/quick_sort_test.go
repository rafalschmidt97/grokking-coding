package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		inputArray    []int
		expectedArray []int
	}{
		{inputArray: []int{64, 34, 25, 12, 22, 11, 90}, expectedArray: []int{11, 12, 22, 25, 34, 64, 90}},
		{inputArray: []int{64, 34, 25, 12, 22, 11, 90}, expectedArray: []int{11, 12, 22, 25, 34, 64, 90}},
		{inputArray: []int{1, 2, 3, 4, 5}, expectedArray: []int{1, 2, 3, 4, 5}},
		{inputArray: []int{5, 4, 3, 2, 1}, expectedArray: []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.inputArray, tt.expectedArray)
		t.Run(testname, func(t *testing.T) {
			sortedArray := quickSort(tt.inputArray)

			if !reflect.DeepEqual(sortedArray, tt.expectedArray) {
				t.Errorf("got %v; want %v, input %v", sortedArray, tt.expectedArray, tt.inputArray)
			}
		})
	}
}
