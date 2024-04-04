package main

import "fmt"

// in place sorting
// time complexity: O(n^2)
// space complexity: O(1)
func bubbleSort(array []int) {
	for io := 0; io < len(array); io++ {
		for ii := io + 1; ii < len(array); ii++ {
			if array[io] > array[ii] {
				tempio := array[io]
				array[io] = array[ii]
				array[ii] = tempio
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
