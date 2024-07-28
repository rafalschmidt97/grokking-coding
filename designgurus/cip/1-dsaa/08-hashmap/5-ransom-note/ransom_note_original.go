package main

// Time complexity: O(N+M)
// Space complexity: O(1)
func ransomNoteOriginal(note string, input string) bool {
	inputMap := make(map[rune]int, 0)
	for _, l := range input {
		inputMap[l]++
	}
	for _, l := range note {
		if inputMap[l] == 0 {
			return false
		}
		inputMap[l]--
	}
	return true
}
