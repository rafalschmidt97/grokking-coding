package main

import (
	"fmt"
	"strings"
)

// Time complexity: O(n)
// Space complexity: O(n)
func makeGreat(input string) string {
	stack := []rune{}

	for _, v := range input {
		if len(stack) > 0 &&
			strings.EqualFold(string(stack[len(stack)-1]), string(v)) &&
			stack[len(stack)-1] != v {

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}

func main() {
	input := "AaBbCcDdEeff"
	fmt.Println(makeGreat(input))
}
