package main

import "fmt"

// Time complexity: O(logm + logn)
// Space complexity: O(1)
func search2dMatrix(array [][]int, target int) bool {
	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middleIndex := lowIndex + (highIndex-lowIndex)/2

		if array[middleIndex][len(array[middleIndex])-1] < target {
			lowIndex = middleIndex + 1 // Target is on the right half
		} else {
			highIndex = middleIndex - 1 // Target is on the left half
		}
	}

	return binarySearch(array[lowIndex], target)
}

func binarySearch(array []int, target int) bool {
	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middleIndex := lowIndex + (highIndex-lowIndex)/2

		if array[middleIndex] == target {
			return true // Target found, return its index
		} else if array[middleIndex] < target {
			lowIndex = middleIndex + 1 // Target is on the right half
		} else {
			highIndex = middleIndex - 1 // Target is on the left half
		}
	}
	return false
}

func main() {
	target := 5
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(search2dMatrix(matrix, target))
}
