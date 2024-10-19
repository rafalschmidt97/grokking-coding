package main

import (
	"fmt"
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1: []int{
				1, 3, 5,
			},
			nums2: []int{
				2, 4, 6,
			},
			expected: 3.5,
		},
		{
			nums1: []int{
				1, 1, 1,
			},
			nums2: []int{
				2, 2, 2,
			},
			expected: 1.5,
		},
		{
			nums1: []int{
				2, 3, 5, 8,
			},
			nums2: []int{
				1, 4, 6, 7, 9,
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.nums1, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := medianSortedArraysWithMergeSort(tt.nums1, tt.nums2)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
