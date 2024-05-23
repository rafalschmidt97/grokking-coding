package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateBinary(input int) []string {
	if input <= 0 {
		panic("not supported")
	}

	generatedNumbers := make([]string, 0, input)
	for i := 1; i <= input; i++ {
		vi := i
		queue := []string{} // could use ints but then it requires a string builder
		for vi >= 1 {
			rest := vi % 2
			queue = append([]string{strconv.Itoa(rest)}, queue...) // append in front
			vi /= 2
		}
		generatedNumbers = append(generatedNumbers, strings.Join(queue, ""))
	}

	return generatedNumbers
}

func main() {
	input := 2
	fmt.Println(generateBinary(input))
}
