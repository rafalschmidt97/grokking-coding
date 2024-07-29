# Problem 2: Jewels and Stones (easy)

## Problem Statement

Given two strings. The first string represents types of jewels, where each
character is a unique type of jewel. The second string represents stones you
have, where each character represents a type of stone. Determine how many of the
stones you have are also jewels.

## Example

```text
Example 1:
Input: Jewels = "abc", Stones = "aabbcc"
Expected Output: 6
Justification: All the stones are jewels.

Example 2:
Input: Jewels = "aA", Stones = "aAaZz"
Expected Output: 3
Justification: There are 2 'a' and 1 'A' stones that are jewels.

Example 3:
Input: Jewels = "zZ", Stones = "zZzZzZ"
Expected Output: 6
Justification: All the stones are jewels.
```

## Constraints

```text
1 <= jewels.length, stones.length <= 50
jewels and stones consist of only English letters.
All the characters of jewels are unique.
```
