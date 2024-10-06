package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(1)
func removeDuplicateLetters(input string) string {
	charFrequencyMap := make(map[rune]int)
	presentInResult := make(map[rune]bool)
	var result []rune

	// Count the frequency of each character
	for _, c := range input {
		charFrequencyMap[c]++
	}

	for _, char := range input {
		if !presentInResult[char] {
			// Ensure smallest lexicographical order
			// Remove characters from the end of the result if:
			// 1. The current character is smaller than the last character in the result
			// 2. The last character in the result appears later in the input string
			for len(result) > 0 && char < result[len(result)-1] && charFrequencyMap[result[len(result)-1]] > 0 {
				// Mark the last character in the result as not present
				delete(presentInResult, result[len(result)-1])
				// Remove the last character from the result
				result = result[:len(result)-1]
			}
			result = append(result, char)
			presentInResult[char] = true
		}
		charFrequencyMap[char]-- // Decrease the frequency
	}

	return string(result)
}

func main() {
	input := "babac"
	fmt.Println(removeDuplicateLetters(input))
}
