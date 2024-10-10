package main

import "fmt"

// Time complexity: O(n log n), recursion and count
// Space complexity: O(log n), depth of call stack
func majorityElement(array []int) int {
	// Base case: if the segment contains only one element
	if len(array) == 1 {
		return array[0]
	}

	// Find the middle index of the current segment
	mid := len(array) / 2

	// Recursively find the majority element in the left and right halves
	leftMajority := majorityElement(array[:mid])
	rightMajority := majorityElement(array[mid:])

	// If both halves agree on the majority element, return it
	if leftMajority == rightMajority {
		return leftMajority
	}

	// Count the occurrences of the left and right majority elements in the current segment
	leftCount := countOccurrences(array, leftMajority)
	rightCount := countOccurrences(array, rightMajority)

	// Return the element with the higher count in the current segment
	if leftCount > rightCount {
		return leftMajority
	}
	return rightMajority // also return right one when it's equal to left
}

// Helper function to count the occurrences of a number in the array
func countOccurrences(array []int, number int) int {
	count := 0
	for _, value := range array {
		if value == number {
			count++
		}
	}
	return count
}

func main() {
	input := []int{
		9, 9, 1, 1, 9, 1, 9, 9,
	}
	fmt.Println(majorityElement(input))
}
