package main

import "fmt"

// Time complexity: O(logn)
// Space complexity: O(1)
func maximumCountPositiveNegative(array []int) int {
	lowIndex, highIndex := 0, len(array)-1
	maxNegatives, maxPositives := 0, 0 // To hold the counts of negatives and positives

	// First pass to find the count of negative numbers
	for lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2
		if array[midIndex] < 0 {
			maxNegatives = midIndex + 1 // Update count of negatives
			lowIndex = midIndex + 1     // Move to the right
		} else {
			highIndex = midIndex - 1 // Continue searching in the left half
		}
	}

	lowIndex, highIndex = 0, len(array)-1

	// Second pass to find the count of positive numbers
	for lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2
		if array[midIndex] > 0 {
			maxPositives = len(array) - midIndex // Update count of positives
			highIndex = midIndex - 1             // Continue searching in the left half
		} else {
			lowIndex = midIndex + 1 // Move to the right
		}
	}

	// Return the maximum of the counts of negatives and positives
	if maxNegatives > maxPositives {
		return maxNegatives
	}
	return maxPositives
}

func main() {
	input := []int{-4, -3, -1, 0, 1, 3, 5, 7}
	fmt.Println(maximumCountPositiveNegative(input))
}
