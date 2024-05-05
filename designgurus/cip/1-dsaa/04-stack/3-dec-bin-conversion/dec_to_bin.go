package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Time complexity: O(logn) due to division by 2 at each step
// Space complexity: O(logn)
func decToBin(input int) string {
	stack := make([]int, 0)
	leftAcum := input

	for leftAcum != 0 {
		modulo := leftAcum % 2
		stack = append(stack, modulo)
		leftAcum /= 2
	}

	var builder strings.Builder
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		builder.WriteString(strconv.Itoa(pop))
		stack = stack[:len(stack)-1]
	}
	return builder.String()
}

func main() {
	input := 2
	fmt.Println(decToBin(input))
}
