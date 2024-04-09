package main

import (
	"math"
)

// TODO: broken?!
func findDifferenceArrayOriginal(nums []int) []int {
	n := len(nums)
	leftSum := make([]int, n)
	rightSum := make([]int, n)
	differenceArray := make([]int, n)

	// Calculate leftSum array
	leftSum[0] = nums[0]
	for i := 1; i < n; i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
	}

	// Calculate rightSum array
	rightSum[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i]
	}

	// Calculate differenceArray
	for i := 0; i < n; i++ {
		differenceArray[i] = int(math.Abs(float64(leftSum[i] - rightSum[i])))
	}

	return differenceArray
}
