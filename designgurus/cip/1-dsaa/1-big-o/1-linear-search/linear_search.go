package main

import (
	"errors"
	"fmt"
)

// returns index if exists
// time complexity: O(n)
func linearSearch(array []int, x int) (index int, err error) {
	for i, v := range array {
		if v == x {
			return i, nil
		}
	}
	return -1, errors.New("not exists")
}

func main() {
	array := []int{3, 5, 7, 9, 11}
	x := 7
	r, e := linearSearch(array, x)
	fmt.Printf("%v (%v)\n", r, e)
}
