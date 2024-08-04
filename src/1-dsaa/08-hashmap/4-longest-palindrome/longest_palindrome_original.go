package main

// Time complexity: O(N+M)
//
// Space complexity: O(1), as it doesn't grow.
// For a string with only lowercase English letters,
// the maximum number of unique characters is 26. ASCII 128/256
func longestPalindromeOriginal(input string) int {
	lettersMap := make(map[rune]int, 0)
	for _, letter := range input {
		lettersMap[letter]++
	}
	length := 0
	oddMarker := false
	for _, occurances := range lettersMap {
		if occurances%2 == 0 {
			length += occurances
		} else {
			length += occurances - 1
			oddMarker = true
		}
	}

	if oddMarker {
		length += 1
	}

	return length
}
