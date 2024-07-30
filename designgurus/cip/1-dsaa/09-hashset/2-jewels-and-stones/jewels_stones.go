package main

import "fmt"

// Time complexity: O(N+M)
// Space complexity: O(1), alphabet is fixed
func jewelsStones(jewels string, stones string) int {
	lewelsSet := make(map[rune]struct{}, 0)
	for _, jewelChar := range jewels {
		lewelsSet[jewelChar] = struct{}{}
	}
	counter := 0
	for _, stoneChar := range stones {
		if _, ok := lewelsSet[stoneChar]; ok {
			counter++
		}
	}
	return counter
}

func main() {
	jewels := "abc"
	stones := "aabbcc"
	fmt.Println(jewelsStones(jewels, stones))
}
