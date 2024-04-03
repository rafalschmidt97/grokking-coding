package main

import (
	"errors"
	"fmt"
)

// returns index if exists
// O(N) time complexity
func linearSearch(input *[]int, x int) (index int, err error) {
	for i, v := range *input {
		if v == x {
			return i, nil
		}
	}
	return -1, errors.New("not exists")
}

func main() {
	input := []int{3, 5, 7, 9, 11}
	x := 7
	r, e := linearSearch(&input, x)
	fmt.Printf("%v (%v)\n", r, e)
}
