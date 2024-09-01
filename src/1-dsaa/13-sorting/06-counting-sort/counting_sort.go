package main

import "fmt"

// Time Complexity: O(n + k), where n is the number of elements and k is the range of the input.
// Space Complexity: O(k), due to the additional space used by the count array.
func countingSort(arr []int) []int {
	maximum := 0
	for _, num := range arr {
		if num > maximum {
			maximum = num
		}
	}
	count := make([]int, maximum+1)
	output := make([]int, len(arr))

	// Count each element
	for _, num := range arr {
		count[num]++
	}

	// Accumulate count
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for _, num := range arr {
		output[count[num]-1] = num
		count[num]--
	}

	return output
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := countingSort(inputArray)
	fmt.Println(sortedArray)
}
