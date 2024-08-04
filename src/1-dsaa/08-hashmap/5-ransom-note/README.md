# Problem 5: Ransom Note (easy)

## Problem Statement

Given two strings, one representing a ransom note and the other representing the
available letters from a magazine, determine if it's possible to construct the
ransom note using only the letters from the magazine. Each letter from the
magazine can be used only once.

## Example

```text
Example 1:
Input: Ransom Note = "hello", Magazine = "hellworld"
Expected Output: true
Justification: The word "hello" can be constructed from the letters in "hellworld".

Example 2:
Input: Ransom Note = "notes", Magazine = "stoned"
Expected Output: true
Justification: The word "notes" can be fully constructed from "stoned" from its first 5 letters.

Example 3:
Input: Ransom Note = "apple", Magazine = "pale"
Expected Output: false
Justification: The word "apple" cannot be constructed from "pale" as we are missing one 'p'.
```

## Constraints

```text
1 <= ransomNote.length, magazine.length <= 10^5
ransomNote and magazine consist of lowercase English letters.
```
