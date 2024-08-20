package main

import "fmt"

// Time Complexity: O(N * L + M), combining the time to build the Trie and perform searches.
//
//	Building the Trie: (O(N x L)), where (N) is the number of products and (L) is the average length of the products.
//	Searching for Each Prefix: (O(M x K)), where (M) is the length of the
//	search word and (K) is the time taken for the DFS, which is limited to 3 (constant time).
//	Overall, it's approximately (O(M)).
//
// Space complexity is dominated by the trie storage, which is (O(N x L)).
//
//	Trie Storage: (O(N x L)), as each product of average length (L) is stored in the trie.
//	Search Results: (O(M)), as we store up to 3 suggestions for each character in the search word.
func searchSugestions(search string, words []string) [][]string {
	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	searchResults := [][]string{}

	prefix := ""
	for _, charSuggest := range search {
		prefix += string(charSuggest)

		charResults := []string{}
		searchForWords(prefix, &charResults, root)

		searchResults = append(searchResults, charResults)
	}

	return searchResults
}

func searchForWords(prefix string, results *[]string, rootNode *TrieNode) {
	currentNode := rootNode
	for _, charInner := range prefix {
		index := charInner - 'a'
		if currentNode.children[index] == nil {
			return
		}
		currentNode = currentNode.children[index]
	}
	dfsForWords(prefix, results, currentNode)
}

func dfsForWords(prefix string, results *[]string, currentNode *TrieNode) {
	if len(*results) == 3 {
		return
	}

	if currentNode.endWord {
		*results = append(*results, prefix)
	}

	for char := 'a'; char <= 'z'; char++ {
		index := char - 'a'
		if currentNode.children[index] != nil {
			newPrefix := prefix + string(char)
			dfsForWords(newPrefix, results, currentNode.children[index])
		}
	}
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
