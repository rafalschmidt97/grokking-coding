package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(n + mlogm)
// Space complexity: O(n + m)
func sortVowels(input string) string {
	vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	vowelsMap := map[rune]struct{}{}
	for _, v := range vowels {
		vowelsMap[v] = struct{}{}
	}

	inputVowels := []rune{}
	for _, v := range input {
		if _, ok := vowelsMap[v]; ok {
			inputVowels = append(inputVowels, v)
		}
	}

	sort.Slice(inputVowels, func(i, j int) bool {
		return inputVowels[i] <= inputVowels[j]
	})

	iv := 0
	newInput := "" // consider string builder
	for _, v := range input {
		if _, ok := vowelsMap[v]; ok {
			newInput += string(inputVowels[iv]) // consonants
			iv++
		} else {
			newInput += string(v) // consonants
		}
	}

	return newInput
}

func main() {
	input := "gamE"
	fmt.Println(sortVowels(input))
}
