package main

import (
	"fmt"
)

// If there are no elements to the left/right of i, the corresponding sum should be
// taken as 0.
// differenceArray[i] = | leftSum[i] - rightSum[i] |
//
// Time complexity: O(n)
// Space complexity: O(n)
func findDifferenceArray(input []int) []int {
	differenceArray := make([]int, len(input)) // consider inplace

	leftSum, rightSum := make([]int, len(input)), make([]int, len(input))

	for i := 0; i < len(input); i++ {
		if i == 0 { // left edges
			leftSum[i] = 0 // consider starting from 1 and doing leftSum[0] = 0
		} else {
			leftSum[i] = leftSum[i-1] + input[i-1]
		}

		rightIndex := len(input) - 1 - i
		if rightIndex == len(input)-1 { // right edge
			rightSum[rightIndex] = 0
		} else {
			rightSum[rightIndex] = rightSum[rightIndex+1] + input[rightIndex+1]
		}
	}

	for i := 0; i < len(input); i++ {
		sum := leftSum[i] - rightSum[i]
		if sum < 0 { // math.Abs works with floats, requires int(math.Abs(xxx))
			differenceArray[i] = -sum
		} else {
			differenceArray[i] = sum
		}
	}
	return differenceArray
}

func main() {
	input := []int{2, 5, 1, 6}
	fmt.Println(findDifferenceArray(input))
}
