package main

import (
	"sort"
)

func divideArrayOriginal(nums []int, k int) [][]int {
	sort.Ints(nums) // Sort the array to bring closer elements near each other
	var result [][]int

	// Iterate over the array in steps of 3 since each subarray must contain exactly 3 elements
	for i := 0; i <= len(nums)-3; i += 3 {
		// Check if the max difference within the current triplet is within the allowed limit k
		if nums[i+2]-nums[i] <= k {
			subgroup := []int{nums[i], nums[i+1], nums[i+2]}
			result = append(result, subgroup) // Add the valid group to the result
		} else {
			return [][]int{} // Return an empty list if any group fails the condition
		}
	}

	return result // Return the final grouped list
}
