package main

import "fmt"

// Time complexity: O(n*m)
// Space complexity: O(1)
func maximumWealth(input [][]int) int {
	maximumWealth := 0
	for _, c := range input {
		customerWealth := 0
		for _, b := range c {
			customerWealth = customerWealth + b
		}

		if customerWealth > maximumWealth {
			maximumWealth = customerWealth
		}
	}

	return maximumWealth
}

func main() {
	input := [][]int{{1, 2, 3}}
	fmt.Println(maximumWealth(input))
}
