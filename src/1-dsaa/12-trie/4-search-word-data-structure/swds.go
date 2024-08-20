package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode // Represents each letter of the alphabet.
	isEnd    bool          // Flag to represent if the node is the end of a word.
}

func processStructureOperations(operations [][]string) []int {
	root := TrieNode{
		children: [26]*TrieNode{},
		isEnd:    false,
	}

	results := []int{}

	for _, operation := range operations {
		switch operation[0] {
		case "addWord":
			result := addWord(&root, operation[1])
			results = append(results, result)
		case "search":
			result := searchWithWildcardSupport(&root, operation[1])
			results = append(results, result)
		}
	}

	return results
}

func addWord(root *TrieNode, word string) int {
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

// Search: O(n * m) in the worst case, where n is the length of the word and m is the total number
// of nodes in the Trie. This happens when the search word contains dots ('.').
// However, for words without dots, the search is O(n).
//
// Space Complexity: O(m * n), where m is the total number of Trie nodes and n is the
// average number of characters in the words. Each Trie node has up to 26 children
// (for each letter of the alphabet). In the worst case, when no nodes
// are shared, the space complexity is O(m * n).
func searchWithWildcardSupport(root *TrieNode, word string) int {
	result := searchInNode(word, root)
	if result {
		return 1
	}
	return 0
}

func searchInNode(word string, node *TrieNode) bool {
	for i, ch := range word {
		if ch == '.' {
			for _, child := range node.children {
				if child != nil && searchInNode(word[i+1:], child) {
					return true
				}
			}
			return false
		} else {
			index := ch - 'a'
			if node.children[index] == nil {
				return false
			}
			node = node.children[index]
		}
	}
	return node.isEnd
}

func main() {
	input := [][]string{
		{"addWord", "apple"},
		{"addWord", "banana"},
		{"search", "apple"},
		{"search", "......"},
	}
	fmt.Println(processStructureOperations(input))
}
