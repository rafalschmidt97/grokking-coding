package main

import (
	"fmt"
	"sort"
)

// Time complexity: O(N)
// Space complexity: O(1)
func boatsToSavePeople(people []int, limit int) int {
	sort.Ints(people)    // Sort the slice in ascending order
	i := 0               // Pointer for the lightest person
	j := len(people) - 1 // Pointer for the heaviest person
	boats := 0           // Count of boats

	for i <= j {
		if people[i]+people[j] <= limit {
			// If the lightest and heaviest person can share a boat
			i++ // Move to the next lightest person
		}
		j--     // Move to the next heaviest person
		boats++ // Increment boat count
	}

	// When last person is left.
	if i == j {
		boats += 1
	}

	return boats // Return the total number of boats needed}
}

func main() {
	people := []int{10, 55, 70, 20, 90, 85}
	limit := 100
	fmt.Println(boatsToSavePeople(people, limit))
}
