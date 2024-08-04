package main

import "fmt"

// Time complexity: TODO: change
// Space complexity: TODO: change
func longestSubstringWithoutRepeatingChar(input string) int {
	repeatingChar := make(map[byte]struct{})
	maxLength, start, end := 0, 0, 0

	for end < len(input) {
		if _, ok := repeatingChar[input[end]]; !ok {
			repeatingChar[input[end]] = struct{}{}
			maxLength = max(maxLength, end-start+1)
			end++
		} else {
			delete(repeatingChar, input[start])
			start++
		}
	}

	return maxLength
}

func main() {
	input := "abcdaef"
	fmt.Println(longestSubstringWithoutRepeatingChar(input))
}
