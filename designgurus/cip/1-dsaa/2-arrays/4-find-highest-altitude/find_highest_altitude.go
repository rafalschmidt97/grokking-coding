package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(1)
func findHighestAltitude(input []int) int {
	currentAltitude := 0
	highestAltitude := 0

	for _, v := range input {
		currentAltitude = currentAltitude + v

		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}

	}

	return highestAltitude
}

func main() {
	input := []int{-5, 1, 5, 0, -7}
	fmt.Println(findHighestAltitude(input))
}
