package main

import (
	"fmt"
)

// Time complexity: O(N)
// Space complexity: O(1)
func maximumNumberOfBallons(input string) int {
	lettersMap := make(map[rune]int, 0)

	for _, letter := range input {
		lettersMap[letter]++
	}

	maxTimesB := lettersMap['b']
	maxTimesA := lettersMap['a']
	maxTimesL := lettersMap['l'] / 2
	maxTimesO := lettersMap['o'] / 2
	maxTimesN := lettersMap['n']

	return min(maxTimesB, maxTimesA, maxTimesL, maxTimesO, maxTimesN)
}

func main() {
	input := "balloonballoon"
	fmt.Println(maximumNumberOfBallons(input))
}
