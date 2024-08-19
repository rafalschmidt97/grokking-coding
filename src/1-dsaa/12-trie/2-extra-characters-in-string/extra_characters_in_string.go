package main

import "fmt"

// A greedy algorithm might not be optimal because it doesn't
// consider all possible ways to partition the string.
// It makes a decision based on the first matching word it
// encounters, which can sometimes lead to suboptimal results
// when a different word choice might result in fewer extra characters.
func minExtraCharGreedy(text string, words []string) int {
	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	extraChars := 0
	start := 0
	for start < len(text) {
		current := root
		matched := false
		for i := start; i < len(text); i++ {
			index := text[i] - 'a'

			if current.children[index] == nil {
				break
			}
			current = current.children[index]

			if current.endWord {
				start = i + 1
				matched = true
				break
			}
		}

		if !matched {
			extraChars++
			start++
		}
	}
	return extraChars
}

// Time Complexity
//
//	Trie Construction: (O(W x L)), where (W) is the number of words in the dictionary and (L) is the average length of these words.
//	Dynamic Programming Calculation: (O(n^2)), where (n) is the length of the input string s.
//	Total Time Complexity: (O(W x L + n^2)).
//
// Space Complexity
//
//	Trie Storage: (O(W x L)), for storing the words in the trie.
//	Dynamic Programming Array: (O(n)), for the array used in dynamic programming.
//	Total Space Complexity: (O(W x L + n)).
func minExtraChar(text string, words []string) int {
	// Step 1: Build the Trie from words
	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	// Step 2: Initialize the DP array
	dp := make([]int, len(text)+1)
	for i := 0; i <= len(text); i++ {
		dp[i] = i // Maximum extra characters can be the entire length
	}

	// Step 3: Fill DP array
	for i := 0; i < len(text); i++ {
		current := root
		for j := i; j < len(text); j++ {
			index := text[j] - 'a'
			if current.children[index] == nil {
				break
			}
			current = current.children[index]

			if current.endWord {
				// If the word ends here, update dp[j+1] considering this substring
				dp[j+1] = min(dp[j+1], dp[i])
			}
		}
		// No word was found that starts at i
		dp[i+1] = min(dp[i+1], dp[i]+1)
	}

	// Step 4: Return the result which is dp[len(text)]
	return dp[len(text)]
}

func main() {
	text := "amazingracecar"
	words := []string{
		"race", "car",
	}
	fmt.Println(minExtraCharGreedy(text, words))
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
