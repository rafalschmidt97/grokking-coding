package main

import "fmt"

// Time Complexity: O(n^2),
// as there are two nested loops, making it inefficient on large lists.
//
// Space Complexity: O(1),
// because it's an in-place sort requiring no additional storage beyond
// swapping elements.
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// Assume the current index has the minimum value
		minIndex := i
		// Find the index of the minimum element in the unsorted portion of the array
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the first element of the unsorted portion
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	// array := []int{64, 34, 25, 12, 22, 11, 90}
	array := []int{64, 34, 25}
	// array := []int{64, 34, 25, 22}
	fmt.Println(array)
	selectionSort(array)
	fmt.Println(array)
}
