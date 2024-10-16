package main

import (
	"fmt"
	"unicode"
)

// Time complexity: O(n^3)
// Space complexity: O(n)
func longestNiceSubstring(str string) string {
	if len(str) <= 0 {
		return "" // base condition
	}

	if isNice(str) { // whole string is nice
		return str
	}

	longest := ""

	// Divide and conquer: Check for each split point
	for i := 0; i < len(str); i++ {
		left := str[:i]
		right := str[i+1:]

		// Recursively find the longest nice substring in the left and right parts
		leftNice := longestNiceSubstring(left)
		rightNice := longestNiceSubstring(right)

		// Conquer: Choose the longer one between leftNice and rightNice
		if len(leftNice) >= len(rightNice) {
			if len(leftNice) > len(longest) {
				longest = leftNice
			}
		} else {
			if len(rightNice) > len(longest) {
				longest = rightNice
			}
		}
	}

	return longest
}

func isNice(str string) bool {
	charMap := make(map[rune]int)

	for _, char := range str {
		charMap[char]++
	}

	for _, char := range str {
		// Check both cases for each character
		if unicode.IsLower(char) && charMap[unicode.ToUpper(char)] == 0 ||
			unicode.IsUpper(char) && charMap[unicode.ToLower(char)] == 0 {
			return false
		}
	}

	return true
}

func main() {
	input := "BbCcXxY"
	fmt.Println(longestNiceSubstring(input))
}
