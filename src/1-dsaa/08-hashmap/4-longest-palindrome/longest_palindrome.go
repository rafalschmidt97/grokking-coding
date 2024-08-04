package main

import "fmt"

// Time complexity: O(N)
//
// Space complexity: O(1), as it doesn't grow.
// For a string with only lowercase English letters,
// the maximum number of unique characters is 26. ASCII 128/256
func longestPalindrome(input string) int {
	lettersMap := make(map[rune]int, 0)
	for _, letter := range input {
		lettersMap[letter]++
	}
	length := 0
	for _, letter := range input {
		occurances := lettersMap[letter]
		length += occurances / 2
	}

	return length
}

func main() {
	input := "bananas"
	fmt.Println(longestPalindrome(input))
}
