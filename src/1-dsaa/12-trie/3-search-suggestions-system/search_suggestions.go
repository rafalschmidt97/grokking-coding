package main

import "fmt"

func searchSugestions(search string, words []string) [][]string {
	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	searchResults := [][]string{}

	prefix := ""
	for _, char := range search {
		prefix += string(char)

		charResults := []string{}
		// TODO: trie DFS search

		searchResults = append(searchResults, charResults)
	}

	return searchResults
}

func main() {
	search := "app"
	input := []string{
		"apple", "apricot", "application",
	}
	fmt.Println(searchSugestions(search, input))
}

// Boilerplate

type TrieNode struct {
	children [26]*TrieNode
	endWord  bool
}

func (node *TrieNode) insert(word string) {
	current := node
	for _, char := range word {
		index := char - 'a'
		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.endWord = true
}
