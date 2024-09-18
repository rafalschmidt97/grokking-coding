package main

import (
	"sort"
)

func frequencyOfMostFrequentElementOriginal(elements []int, maxOperations int) int {
	sort.Ints(elements)                           // Sort the array to facilitate the equalization to the highest element.
	cumulativeSum := make([]int64, len(elements)) // Array to store cumulative sums.
	cumulativeSum[0] = int64(elements[0])         // Initialize the first element of cumulative sum.

	// Build the cumulative sum array.
	for i := 1; i < len(elements); i++ {
		cumulativeSum[i] = int64(elements[i]) + cumulativeSum[i-1]
	}

	maximumFrequency := 0
	// Compute max frequency for each end position.
	for i := 0; i < len(elements); i++ {
		maximumFrequency = max(maximumFrequency, findMaxSubarrayLength(i, maxOperations, elements, cumulativeSum))
	}

	return maximumFrequency // Return the maximum frequency found.
}

// findMaxSubarrayLength finds the maximum length of subarray that can be made equal to `elements[index]`
// using at most `maxOperations`.
func findMaxSubarrayLength(index, maxOperations int, elements []int, cumulativeSum []int64) int {
	target := elements[index] // The target number we want to make others equal to.
	start := 0                // Start of the search range.
	end := index              // End of the search range, we consider subarrays ending at `index`.
	bestLength := index       // This will store the best start position for the longest valid subarray.

	// Binary search to find the maximum length of subarray that can be made equal to `target`.
	for start <= end {
		mid := (start + end) / 2             // Midpoint of the current search range.
		count := int64(index - mid + 1)      // Number of elements from `mid` to `index`.
		requiredSum := count * int64(target) // If all elements are `target`, this is the total they would sum to.
		var existingSum int64                // Current sum from `mid` to `index`.

		// Calculate the current sum from `mid` to `index`.
		if mid > 0 {
			existingSum = cumulativeSum[index] - cumulativeSum[mid-1]
		} else {
			existingSum = cumulativeSum[index]
		}

		operationsRequired := requiredSum - existingSum // How many increments are needed to make all equal to `target`.

		// Update the search range based on the number of operations required.
		if operationsRequired > int64(maxOperations) {
			start = mid + 1 // If more operations are required than allowed, move the start up.
		} else {
			bestLength = mid // Update bestLength as this is a valid subarray.
			end = mid - 1    // Try for a longer valid subarray.
		}
	}

	return index - bestLength + 1 // Return the length of the longest valid subarray.
}
