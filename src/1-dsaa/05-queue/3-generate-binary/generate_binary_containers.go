package main

import (
	"container/list"
	"strconv"
	"strings"
)

func generateBinaryContainers(input int) []string {
	if input <= 0 {
		panic("not supported")
	}

	generatedNumbers := make([]string, 0, input)

	for i := 1; i <= input; i++ {
		vi := i
		queue := list.New() // Use list as a queue

		// Generate the binary representation by pushing elements to the front
		for vi >= 1 {
			rest := vi % 2
			queue.PushFront(strconv.Itoa(rest)) // Push in front
			vi /= 2
		}

		// Build the binary number string from the queue
		var binaryBuilder strings.Builder
		for e := queue.Front(); e != nil; e = e.Next() {
			binaryBuilder.WriteString(e.Value.(string))
		}

		// Append the generated binary number to the result
		generatedNumbers = append(generatedNumbers, binaryBuilder.String())
	}

	return generatedNumbers
}
