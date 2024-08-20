# Design Add and Search Words Data Structure (medium)

## Problem Statement

Design a data structure that supports the addition of new words and the ability
to check if a string matches any previously added word.

Implement the Solution class:

Solution() Initializes the object. void addWord(word) Inserts word into the data
structure, making it available for future searches. bool search(word) Checks if
there is any word in the data structure that matches word. The method returns
true if such a match exists, otherwise returns false. Note: In the search query
word, the character '.' can represent any single letter, effectively serving as
a wildcard character.

## Example

```text
Example 1:

Input:
  ["Solution", "addWord", "addWord", "search", "search"]
  [[], ["apple"], ["banana"], ["apple"], ["......"]]
Expected Output: [-1, -1, -1, 1, 1]
Justification: After adding the words "apple" and "banana", searching for
"apple" will return true since "apple" is in the data structure.
Searching for "......" will also return true as "banana" match the pattern.

Example 2:
Input:
  ["Solution", "addWord", "addWord", "search", "search"]
  [[], ["cat"], ["dog"], ["c.t"], ["d..g"]]
Expected Output: [-1, -1, -1, 1, 0]
Justification: "c.t" matches "cat" and "d..g" doesn't matches "dog".

Example 3:
Input:
  ["Solution", "addWord", "search", "search"]
  [[], ["hello"], ["h.llo"], ["h...o"]]
Expected Output: [-1, -1, 1, 1]
Justification: "h.llo" and "h...o" both match "hello".
```

## Constraints

```text
1 <= word.length <= 25
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 10^4 calls will be made to addWord and search.
```
