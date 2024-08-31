package main

import "fmt"

func radixSort(arr []int) []int {
	maximum := 0
	for _, num := range arr {
		if num > maximum {
			maximum = num
		}
	}

	// Do counting sort for every digit.
	// The exp is 10^i where i is the current digit number
	for exp := 1; maximum/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}

	return arr
}

// Function to perform counting sort on the array according to the
// digit represented by exp (10^i)
func countSort(arr []int, exp int) {
	output := make([]int, len(arr)) // output array
	count := make([]int, 10)
	var i int

	// Store count of occurrences in count[]
	for i = 0; i < len(arr); i++ {
		count[(arr[i]/exp)%10]++
	}

	// Change count[i] so that it now contains actual position of this digit in output[]
	for i = 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i = len(arr) - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	// Copy the output array to arr[], so that arr[] now contains sorted
	// numbers according to current digit
	for i = 0; i < len(arr); i++ {
		arr[i] = output[i]
	}
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := radixSort(inputArray)
	fmt.Println(sortedArray)
}
