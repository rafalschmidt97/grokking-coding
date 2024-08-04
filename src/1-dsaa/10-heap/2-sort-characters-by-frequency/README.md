# Sort Characters By Frequency (easy)

## Problem Statement

Given a string, arrange its characters in descending order based on the
frequency of each character. If two characters have the same frequency, their
relative order in the output string can be arbitrary.

## Example

```text
Input: "apple"
Expected Output: "ppale" or "ppela"
Justification: The character 'p' appears twice, while 'a', 'l', and 'e'
each appear once. Thus, 'p' should appear before the other characters in the output.

Input: "banana"
Expected Output: "aaannb".
Justification: The character 'a' appears three times, 'n' twice and 'b' once.

Input: "aabb"
Expected Output: "aabb" or "bbaa"
Justification: Both 'a' and 'b' appear twice, so they can appear in any order in the output.
```

## Constraints

```text
1 <= s.length <= 5 * 105
s consists of uppercase and lowercase English letters and digits.
```
