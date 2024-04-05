package main

import "fmt"

// Top-down implementation using lists (not bottom up)
//
// Merge sort continuously cuts down a list into multiple sublists until each has
// only one item, then merges those sublists into a sorted list.
//
// Time complexity: O(nlogn)
//
// Space complexity: O(n), as it creates two additional arrays (left and right)
// to store the left and right halves of the input during each recursive call.
// The memory required to store these arrays increases linearly with the size of the input.
func mergeSort(array []int) {
	// Base case. A list of zero or one elements is sorted
	if len(array) <= 1 {
		return
	}

	// Recursive case. First, divide the list into two consisting
	// of the first half and second half of the list.

	// Can predict the size for arrays precisely
	mid := len(array) / 2
	leftArray, rightArray := make([]int, mid), make([]int, len(array)-mid)
	copy(leftArray, array[:mid])
	copy(rightArray, array[mid:])

	// Recursively sort both sides
	mergeSort(leftArray)
	mergeSort(rightArray)

	// Then merge the now-sorted lists (left and right) with in place replace for array
	// missing = left out after left and right loop
	leftIndex, rightIndex, arrayIndex := 0, 0, 0

	for leftIndex < len(leftArray) && rightIndex < len(rightArray) {
		if leftArray[leftIndex] < rightArray[rightIndex] {
			array[arrayIndex] = leftArray[leftIndex]
			leftIndex++
		} else {
			array[arrayIndex] = rightArray[rightIndex]
			rightIndex++
		}
		arrayIndex++
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)
	for leftIndex < len(leftArray) {
		array[arrayIndex] = leftArray[leftIndex]
		leftIndex++
		arrayIndex++
	}

	for rightIndex < len(rightArray) {
		array[arrayIndex] = rightArray[rightIndex]
		rightIndex++
		arrayIndex++
	}
}

func main() {
	// array := []int{64, 34, 25, 12, 22, 11, 90}
	array := []int{64, 34, 25}
	// array := []int{64, 34, 25, 22}
	fmt.Println(array)
	mergeSort(array)
	fmt.Println(array)
}
