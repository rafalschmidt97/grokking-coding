# Extra Characters in a String (medium)

## Problem Statement

Given a string s and an array of words words. Break string s into multiple
non-overlapping substrings such that each substring should be part of the words.
There are some characters left which are not part of any substring.

Return the minimum number of remaining characters in s, which are not part of
any substring after string break-up.

## Example

```text
Example 1:
Input: s = "amazingracecar", words = ["race", "car"]
Expected Output: 7
Justification: The string s can be rearranged to form "racecar",
leaving 'a', 'm', 'a', 'z', 'i', 'n', 'g' as extra.

Example 2:
Input: s = "bookkeeperreading", words = ["keep", "read"]
Expected Output: 9
Justification: The words "keep" and "read" can be formed from s,
but 'b', 'o', 'o', 'k', 'e', 'r', 'i', 'n', 'g' are extra.

Example 3:
Input: s = "thedogbarksatnight", words = ["dog", "bark", "night"]
Expected Output: 6
Justification: The words "dog", "bark", and "night" can be formed,
leaving 't', 'h', 'e', 's', 'a', 't' as extra characters.star" is not found.
```

## Constraints

```text
1 <= str.length <= 50
1 <= dictionary.length <= 50
1 <= dictionary[i].length <= 50
dictionary[i] and s consists of only lowercase English letters
dictionary contains distinct words
```
