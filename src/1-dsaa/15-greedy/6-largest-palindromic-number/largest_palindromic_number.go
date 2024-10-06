package main

import (
	"fmt"
	"strconv"
)

// Time complexity: O(n)
// Space complexity: O(n)
func largestPalindromicNumber(input string) string {
	frequencies := make([]int, 10)

	for _, ch := range input {
		val := ch - '0'
		frequencies[val]++
	}

	firstHalf := ""
	middle := ""

	for i := 9; i >= 0; i-- {
		count := frequencies[i]
		if count > 0 {
			digit := strconv.Itoa(i)
			// Append half of the count to the first half
			for j := 0; j < count/2; j++ {
				firstHalf += digit
			}
			// If there's an odd count and no middle digit yet, set the middle digit
			if count%2 == 1 && middle == "" {
				middle = digit
			}
		}
	}

	// Construct the second half by reversing the first half
	secondHalf := ""
	for i := len(firstHalf) - 1; i >= 0; i-- {
		secondHalf += string(firstHalf[i])
	}

	// Combine first half, middle (if any), and second half to form the palindrome
	return firstHalf + middle + secondHalf
}

func main() {
	input := "323211444"
	fmt.Println(largestPalindromicNumber(input))
}
