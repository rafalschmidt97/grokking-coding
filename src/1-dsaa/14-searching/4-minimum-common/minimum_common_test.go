package main

import (
	"fmt"
	"testing"
)

func TestMinimumCommonValue(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{
			nums1: []int{
				1, 3, 5, 7,
			},
			nums2: []int{
				3, 4, 5, 6, 8, 10,
			},
			expected: 3,
		},
		{
			nums1: []int{
				2, 4, 6,
			},
			nums2: []int{
				1, 3, 5,
			},
			expected: -1,
		},
		{
			nums1: []int{
				1, 2, 2, 3,
			},
			nums2: []int{
				2, 2, 4,
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.nums1, tt.expected)
		t.Run(testname, func(t *testing.T) {
			got := minimumCommon(tt.nums1, tt.nums2)
			if got != tt.expected {
				t.Errorf("got %v; expected %v", got, tt.expected)
			}
		})
	}
}
