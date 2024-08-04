package main

import "fmt"

// Time complexity: O(N):
//
// Space complexity: O(1), as it doesn't grow.
// For a string with only lowercase English letters,
// the maximum number of unique characters is 26. ASCII 128/256
func firstNonRepeatingChar(input string) int {
	charsMap := make(map[rune]int, 0)

	for _, char := range input {
		charsMap[char]++
	}

	for index, char := range input {
		if charsMap[char] == 1 {
			return index
		}
	}

	return -1
}

func firstNonRepeatingCharComplex(input string) int {
	charsMap := make(map[rune]Character, 0)

	for i, char := range input {
		if v, ok := charsMap[char]; ok {
			charsMap[char] = Character{
				occurances: v.occurances + 1,
				index:      v.index,
			}
		} else {
			charsMap[char] = Character{
				occurances: 1,
				index:      i,
			}
		}
	}

	for _, char := range charsMap {
		if char.occurances == 1 {
			return char.index
		}
	}

	return -1
}

type Character struct {
	occurances int
	index      int
}

func main() {
	input := "apple"
	fmt.Println(firstNonRepeatingChar(input))
}
