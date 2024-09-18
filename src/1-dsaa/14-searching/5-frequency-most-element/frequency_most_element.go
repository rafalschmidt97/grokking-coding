package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(nlogn)
// Space complexity: O(n)
func frequencyOfMostFrequentElement(nums []int, k int) int {
	// Sort the input array
	sort.Ints(nums)

	left := 0
	sum := 0
	maxFreq := 1

	// Iterate through the array with a sliding window approach
	for right := 0; right < len(nums); right++ {
		sum += nums[right] // Add the current number to the sum

		// Check if the window from left to right can be made uniform with nums[right]
		// If not, move the left pointer to the right to reduce the window size
		for nums[right]*(right-left+1) > sum+k {
			sum -= nums[left] // Remove the element at the left of the window
			left++
		}

		// Update the maximum frequency of the most frequent element
		maxFreq = max(maxFreq, right-left+1)
	}

	return maxFreq
}

func main() {
	input := []int{
		1, 2, 3,
	}
	k := 3
	fmt.Println(frequencyOfMostFrequentElement(input, k))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
