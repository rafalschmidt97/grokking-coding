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
	input := []int{1, 2, 3, 4, 5, 6, 7}
	x := 7
	r, e := linearSearch(&input, x)
	fmt.Printf("%v (%v)\n", r, e)
}
