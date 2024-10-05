package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(1)
func validPalindromeByRemovingOne(input string) bool {
	left, right := 0, len(input)-1

	for left < right {
		if input[left] != input[right] {
			// Check both possibilities: removing either fi or li
			return isPalindrome(input, left+1, right) || isPalindrome(input, left, right-1)
		}
		left++
		right--
	}

	return true
}

func isPalindrome(input string, left, right int) bool {
	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	input := "racecar"
	fmt.Println(validPalindromeByRemovingOne(input))
}
