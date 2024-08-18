package main

import "fmt"

// indexPairsOfString finds all index pairs in the text where words from the Trie occur
// Time Complexity:
// - Building the Trie: O(m * k), where m is the number of words and k is the average length of the words.
// - Searching the text: O(n * k), where n is the length of the text and k is the maximum length of any word in the Trie.
// Overall: O(m * k + n * k), which simplifies to O(n * k) as typically n is much larger than m.
//
// Space Complexity:
// - Trie Storage: O(m * k), where m is the number of words and k is the average length of the words.
// - Result Storage: O(p), where p is the number of matching pairs found.
// Overall: O(m * k + p), where the space for the Trie dominates.
func indexPairsOfString(text string, words []string) [][]int {
	// Step 1: Build the Trie from words
	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	// Step 2: Search and collect index pairs inline
	var result [][]int
	for start := 0; start < len(text); start++ {
		current := root
		for i := start; i < len(text); i++ {
			index := text[i] - 'a'

			// If there's no path in the Trie for the current character,
			// we break out of the loop because no further words can match
			// starting from this index.
			if current.children[index] == nil {
				break
			}
			current = current.children[index]

			// If we reach a node where endWord is true, it means we have found
			// a word in the Trie that matches a substring in the text.
			if current.endWord {
				result = append(result, []int{start, i})
			}
		}
	}

	// Note: The pairs are collected in the order they are found.
	// If sorting was required, you would sort the result as follows:
	// sort.Slice(result, func(i, j int) bool {
	// 	if result[i][0] == result[j][0] {
	// 		return result[i][1] < result[j][1]
	// 	}
	// 	return result[i][0] < result[j][0]
	// })

	return result
}

func main() {
	text := "bluebirdskyscraper"
	words := []string{"blue", "bird", "sky"}
	fmt.Println(indexPairsOfString(text, words))
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
