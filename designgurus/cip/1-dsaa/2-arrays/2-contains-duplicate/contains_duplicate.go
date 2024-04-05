package main

import "fmt"

// Brute force approach (improvement is skipping previous)
// Time complexity: O(n^2)
// Space complexity: O(1)
func containsDuplicate(input []int) bool {
	for io := 0; io < len(input); io++ {
		for ii := io + 1; ii < len(input); ii++ {
			if input[io] == input[ii] {
				return true
			}
		}
	}

	return false
}

// Using hash sets (improvement is checking for exists before setting)
// Time complexity: O(n), because we iterate the array only once.
// Space complexity: O(n), as it stores the numbers in a set.
func containsDuplicateUsingSet(input []int) bool {
	set := make(map[int]struct{}, len(input)) // empty struct has 0 bytes, better than bool

	for _, v := range input {
		if _, exists := set[v]; exists {
			return true
		}

		set[v] = struct{}{}
	}

	// return len(set) != len(input) // no need for it
	return false
}

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(x))
}
