package main

import "fmt"

// In place sorting.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func bubbleSort(array []int) {
	for io := 0; io < len(array); io++ {
		for ii := io + 1; ii < len(array); ii++ {
			if array[io] > array[ii] {
				array[io], array[ii] = array[ii], array[io]
			}
		}
	}
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	bubbleSort(inputArray)
	fmt.Println(inputArray)
}
