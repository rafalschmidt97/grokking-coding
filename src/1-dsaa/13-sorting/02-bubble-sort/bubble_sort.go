package main

import "fmt"

// In place sorting.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	bubbleSort(inputArray)
	fmt.Println(inputArray)
}
