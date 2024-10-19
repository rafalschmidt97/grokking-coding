package main

// Time complexity: O(n+m), even though it is /2
// Space complexity: O(1)
//
// Algorithm Walkthrough
// Let's illustrate the algorithm with an example. Consider two arrays [2, 3, 5, 8] and [1, 4, 6, 7, 9].
//
// Initialize Pointers: pointer1 = 0 (for the first array), pointer2 = 0 (for the second array).
// Median Position: Total elements = 9, which is odd. Median position = (9 + 1) / 2 = 5.
// Simulate Merging:
// Compare 2 and 1, move pointer2. New pointers: pointer1 = 0, pointer2 = 1.
// Compare 2 and 4, move pointer1. New pointers: pointer1 = 1, pointer2 = 1.
// Compare 3 and 4, move pointer1. New pointers: pointer1 = 2, pointer2 = 1.
// Compare 5 and 4, move pointer2. New pointers: pointer1 = 2, pointer2 = 2.
// Compare 5 and 6, move pointer1. New pointers: pointer1 = 3, pointer2 = 2.
// Track Median Elements: The fifth element (median position) is 5.
// Determine Median Value: Since the total number is odd, the median is 5.
func medianSortedArraysWithMergeSort(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for efficient binary search
	if len(nums1) > len(nums2) {
		return medianSortedArraysWithMergeSort(nums2, nums1)
	}

	totalLength := len(nums1) + len(nums2)
	isEven := totalLength%2 == 0

	// Two pointers for each array
	var pointer1, pointer2 int
	var current, last int // To track the last two elements for median calculation

	// Iterate until the median position is reached
	for count := 0; count <= totalLength/2; count++ {
		last = current

		// Move pointers and track current value
		if pointer1 == len(nums1) {
			current = nums2[pointer2]
			pointer2++
		} else if pointer2 == len(nums2) || nums1[pointer1] < nums2[pointer2] {
			current = nums1[pointer1]
			pointer1++
		} else {
			current = nums2[pointer2]
			pointer2++
		}
	}

	// Median calculation based on total length being odd or even
	if isEven {
		return float64(current+last) / 2.0
	}
	return float64(current)
}
