package main

import "fmt"

// Not an inplace sorting (like bubblesort).
//
// The quick sort algorithm works by selecting a pivot element and partitioning
// the array into three subarrays: one containing elements less than the pivot,
// one containing elements equal to the pivot, and one containing elements greater
// than the pivot. It then recursively sorts the left and right subarrays.
//
// Time complexity: O(nlogn), worst case O(n) if the pivot is always the smallest
// or largest element in the array
//
// Space complexity: O(logn), determined by the number of recursive calls that are made.
// Each recursive call processes approximately n/2 elements, so the total number of
// recursive calls is approximately logn.
func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	pivot := array[mid]

	// do not know the size for arrays, can only predict upper bound (n)
	lessArray, equalArray, greaterArray := make([]int, 0), make([]int, 0), make([]int, 0)

	// split into three arrays using pivot
	for _, v := range array {
		if v < pivot {
			lessArray = append(lessArray, v)
		} else if v > pivot {
			greaterArray = append(greaterArray, v)
		} else {
			equalArray = append(equalArray, v)
		}
	}

	lessArray = quickSort(lessArray)
	// no need to do anything with equal array
	greaterArray = quickSort(greaterArray)
	return append(lessArray, append(equalArray, greaterArray...)...)
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := quickSort(inputArray)
	fmt.Println(sortedArray)
}
