# Index Pairs of a String (easy)

Given a string text and a list of strings words, identify all [i, j] index pairs
such that the substring text[i...j] is in words.

These index pairs should be returned in ascending order, first by the start
index, then by the end index. Find every occurrence of each word within the
text, ensuring that overlapping occurrences are also identified.

## Problem Statement

## Example

```text
1
Input: text = "bluebirdskyscraper", words = ["blue", "bird", "sky"]
Expected Output: [[0, 3], [4, 7], [8, 10]]
Justification: The word "blue" is found from index 0 to 3, "bird"
from 4 to 7, and "sky" from 8 to 10 in the string.

2
Input: text = "programmingisfun", words = ["pro", "is", "fun", "gram"]
Expected Output: [[0, 2], [3, 6], [11, 12], [13, 15]]
Justification: "pro" is found from 0 to 2, "gram" from 3 to 6,
"is" from 11 to 12, and "fun" from 13 to 15.

3
Input: text = "interstellar", words = ["stellar", "star", "inter"]
Expected Output: [[0, 4], [5, 11]]
Justification: "inter" is found from 0 to 4, and "stellar" from 5 to 11.
"star" is not found.
```

## Constraints

```text
1 <= text.length <= 100
1 <= words.length <= 20
1 <= words[i].length <= 50
text and words[i] consist of lowercase English letters.
All the strings of words are unique.
```
