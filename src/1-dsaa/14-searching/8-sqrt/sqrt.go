package main

import "fmt"

// Time complexity: O(logn)
// Space complexity: O(1)
func sqrt(x int) int {
	if x < 2 {
		return x // return x if it is 0 or 1
	}

	left, right := 2, x/2
	var mid int
	var num int64
	for left <= right { // binary search for the square root
		mid = left + (right-left)/2
		num = int64(mid) * int64(mid)
		if num > int64(x) {
			right = mid - 1 // if mid * mid is greater than x, set right to mid - 1
		} else if num < int64(x) {
			left = mid + 1 // if mid * mid is less than x, set left to mid + 1
		} else {
			return mid // if mid * mid is equal to x, return mid
		}
	}

	return right
}

func main() {
	input := 8
	fmt.Println(sqrt(input))
}
