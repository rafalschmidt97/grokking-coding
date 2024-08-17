package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode // Represents each letter of the alphabet.
	isEnd    bool          // Flag to represent if the node is the end of a word.
}

func processTrieOperations(operations [][]string) []int {
	root := TrieNode{
		children: [26]*TrieNode{},
		isEnd:    false,
	}

	results := []int{}

	for _, operation := range operations {
		switch operation[0] {
		case "insert":
			result := insert(&root, operation[1])
			results = append(results, result)
		case "search":
			result := search(&root, operation[1])
			results = append(results, result)
		case "startsWith":
			result := startsWith(&root, operation[1])
			results = append(results, result)
		}
	}

	return results
}

// N, N
func insert(root *TrieNode, word string) int {
	currentNode := root

	for _, char := range word {
		encodedChar := char - 'a'

		if currentNode.children[encodedChar] == nil {
			currentNode.children[encodedChar] = &TrieNode{
				children: [26]*TrieNode{},
				isEnd:    false,
			}
		}

		currentNode = currentNode.children[encodedChar] // move the pointer
	}

	currentNode.isEnd = true

	return -1 // void, no need to indicate successful insertion or detect if a word was already present
}

// N
func search(root *TrieNode, word string) int {
	currentNode := root

	for _, char := range word {
		encodedChar := char - 'a'

		if currentNode.children[encodedChar] == nil {
			return 0
		}

		currentNode = currentNode.children[encodedChar] // move the pointer
	}

	if currentNode.isEnd {
		return 1
	}

	return 0
}

// N
func startsWith(root *TrieNode, prefix string) int {
	currentNode := root

	for _, char := range prefix {
		encodedChar := char - 'a'

		if currentNode.children[encodedChar] == nil {
			return 0
		}

		currentNode = currentNode.children[encodedChar] // move the pointer
	}

	return 1
}

func main() {
	operations := [][]string{
		{"insert", "apple"},
		{"search", "apple"},
		{"startsWith", "app"},
	}
	fmt.Println(processTrieOperations(operations))
}
