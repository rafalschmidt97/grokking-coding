package main

import "fmt"

// Time complexity: O(nlogm)
// Space complexity: O(1)
func minimumCommon(nums1 []int, nums2 []int) int {
	smallerNums, biggerNums := nums1, nums2

	if len(nums2) > len(nums1) {
		smallerNums, biggerNums = nums2, nums1
	}

	for _, n := range smallerNums {
		if binarySearch(biggerNums, n) != -1 {
			return n
		}
	}

	return -1
}

func binarySearch(array []int, target int) (index int) {
	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middleIndex := lowIndex + (highIndex-lowIndex)/2

		if array[middleIndex] == target {
			return middleIndex // Target found, return its index
		} else if array[middleIndex] < target {
			lowIndex = middleIndex + 1 // Target is on the right half
		} else {
			highIndex = middleIndex - 1 // Target is on the left half
		}
	}
	return -1 // Target not found
}

func main() {
	nums1 := []int{
		1, 3, 5, 7,
	}
	nums2 := []int{
		3, 4, 5, 6, 8, 10,
	}
	fmt.Println(minimumCommon(nums1, nums2))
}
