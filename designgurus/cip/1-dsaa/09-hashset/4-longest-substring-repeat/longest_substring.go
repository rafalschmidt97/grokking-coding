package main

import "fmt"

// Time complexity: TODO: change
// Space complexity: TODO: change
func longestSubstringWithoutRepeatingChar(input string) int {
	repeatingCharSet := make(map[rune]bool, 0)
	for _, char := range input {
		if _, ok := repeatingCharSet[char]; ok {
			repeatingCharSet[char] = true
		} else {
			repeatingCharSet[char] = false
		}
	}
	counter := 1
	for _, char := range input {
		if repeating := repeatingCharSet[char]; !repeating {
			counter++
		} else {
			if counter > 1 {
				counter--
			}
		}
	}
	return counter
}

func main() {
	input := "abcdaef"
	fmt.Println(longestSubstringWithoutRepeatingChar(input))
}
