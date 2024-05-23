package main

import (
	"fmt"
	"strings"
)

// Time complexity: O(n)
// Space complexity: O(n)
func palindromeCheck(input string) bool {
	normalizedInput := strings.ReplaceAll(strings.ToLower(input), " ", "")

	for len(normalizedInput) > 1 {
		if normalizedInput[0] != normalizedInput[len(normalizedInput)-1] {
			return false
		}
		normalizedInput = normalizedInput[1 : len(normalizedInput)-1] // dequeue from each side
	}

	return true
}

func main() {
	input := "madam"
	fmt.Println(palindromeCheck(input))
}
