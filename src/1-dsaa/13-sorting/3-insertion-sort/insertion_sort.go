package main

import "fmt"

// Time Complexity: O(n^2)
// in the average and worst case, but O(n) in the best case
// if the array is nearly sorted.
//
// Space Complexity: O(1),
// as it performs the sorting in place.
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1], that are greater than key,
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	insertionSort(inputArray)
	fmt.Println(inputArray)
}
