package main

import "fmt"

// Time complexity: O(N+M)
// Space complexity: O(1)
func ransomNote(note string, input string) bool {
	inputMap := make(map[rune]int, 0)
	noteMap := make(map[rune]int, 0)
	for _, l := range input {
		inputMap[l]++
	}
	for _, l := range note {
		noteMap[l]++
	}
	for l, occurances := range noteMap {
		if occurances > inputMap[l] {
			return false
		}
	}
	return true
}

func main() {
	note := "hello"
	input := "hellworld"
	fmt.Println(ransomNote(note, input))
}
