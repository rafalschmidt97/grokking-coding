package main

import "fmt"

// Time complexity: O(n+m)
// Space complexity: O(n)
func medianSortedArraysSimple(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0)

	num1i := 0
	num2i := 0
	for len(nums1) > num1i && len(nums2) > num2i {
		if nums1[num1i] < nums2[num2i] {
			merged = append(merged, nums1[num1i])
			num1i++
		} else {
			merged = append(merged, nums2[num2i])
			num2i++
		}
	}
	for len(nums1) > num1i {
		merged = append(merged, nums1[num1i])
		num1i++
	}
	for len(nums2) > num2i {
		merged = append(merged, nums2[num2i])
		num2i++
	}

	mid := (len(merged) - 1) / 2
	if len(merged)%2 == 0 {
		return float64(merged[mid]+merged[mid+1]) / 2
	}

	return float64(merged[mid])
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 3}
	fmt.Println(medianSortedArraysSimple(nums1, nums2))
}
