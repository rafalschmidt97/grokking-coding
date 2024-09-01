package main

import "bytes"

// Helper method to check if a character is a vowel
func isVowel(c byte) bool {
	vowels := "aeiouAEIOU"
	return bytes.IndexByte([]byte(vowels), c) != -1
}

// Sorts the vowels in a string while maintaining the position of consonants
//
// Time Complexity
// - Counting Vowels: Traversing the string to count vowel frequencies requires O(n),
// where (n) is the length of the string.
// - Building the Result: Constructing the final string, including searching for
// the next vowel in the sorted order with a count greater than zero, theoretically
// could reach O(n x m), where (m) is the maximum length of the sorted vowel string
// (sortedVowelOrder). However, since the length of sortedVowelOrder is constant
// and small (10), the effective complexity for this part remains O(n).
// - Overall Time Complexity: Therefore, the total time complexity is O(n)
// since the operations inside the string traversal are bounded by a constant.
//
// Space Complexity
// - Frequency Map: The space complexity for the frequency map is O(n) since the number
// of distinct vowels (keys) is constant (10).
// - StringBuilder: The space used by the StringBuilder is O(n) as it stores the
// rearranged version of the input string.
// - Overall Space Complexity: Thus, the total space complexity of the algorithm is O(n)
// considering the storage required for the result string.
func sortVowelsOriginal(str string) string {
	frequencyMap := make(map[byte]int)

	// Count each vowel's frequency
	for i := 0; i < len(str); i++ {
		c := str[i]
		if isVowel(c) {
			frequencyMap[c]++
		}
	}

	// Predefined sorted order of vowels based on ASCII values
	sortedVowelOrder := "AEIOUaeiou"
	result := ""
	index := 0 // Pointer for sortedVowelOrder

	// Construct the result string with sorted vowels in their positions
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !isVowel(c) {
			result += string(c) // Append consonant directly
		} else {
			// Move to the next vowel with remaining count
			for index < len(sortedVowelOrder) &&
				frequencyMap[sortedVowelOrder[index]] == 0 {
				index++
			}

			// Check if index is within bounds
			if index < len(sortedVowelOrder) {
				result += string(sortedVowelOrder[index])
				frequencyMap[sortedVowelOrder[index]]--
			}
		}
	}

	return result
}
