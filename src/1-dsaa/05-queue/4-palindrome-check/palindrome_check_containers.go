package main

import (
	"container/list"
	"strings"
)

func palindromeCheckContainers(input string) bool {
	normalizedInput := strings.ReplaceAll(strings.ToLower(input), " ", "")
	deque := list.New()
	for _, char := range normalizedInput {
		deque.PushBack(char)
	}

	for deque.Len() > 1 {
		if deque.Remove(deque.Front()) != deque.Remove(deque.Back()) {
			return false
		}
	}

	return true
}
