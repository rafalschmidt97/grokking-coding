package main

import (
	"fmt"
	"sort"
)

// Time complexity: TODO: change
// Space complexity: TODO: change
func largestPalindromicNumber(input string) string {
	inputArray := []rune(input)
	sort.Slice(inputArray, func(i, j int) bool {
		return inputArray[i] > inputArray[j]
	})

	front, back := "", ""
	previous := inputArray[0]
	for i := 1; i < len(inputArray); i++ {
		current := inputArray[i]
		if current == previous {
			previous = 'X' // mark as taken if matching
			front += string(current)
			back = string(current) + back
		} else {
			previous = current
		}
	}

	if len(front) == 0 && len(back) == 0 {
		return string(inputArray[0])
	} else {
		return front + back
	}
}

func main() {
	input := "24321"
	fmt.Println(largestPalindromicNumber(input))
}
