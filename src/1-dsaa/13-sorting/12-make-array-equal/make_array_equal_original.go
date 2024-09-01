package main

import (
	"sort"
)

func makeArrayEqualOriginal(nums []int) int {
	sort.Ints(nums) // Sort the array in ascending order
	operations := 0
	count := 1
	for i := len(nums) - 1; i > 0; i-- { // Loop through the array from the end
		if nums[i] != nums[i-1] { // If the current number is different from the previous one
			operations += count // Increment the total operations by the current count
		}
		count++ // Increment the count for the next number
	}
	return operations // Return the total number of operations
}
