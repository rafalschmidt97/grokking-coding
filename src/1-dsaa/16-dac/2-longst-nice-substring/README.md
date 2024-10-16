# Longest Nice Substring (easy)

## Problem Statement

Given a string str, return the longest nice substring of a given string.

A substring is considered nice if for every lowercase letter in the substring,
its uppercase counterpart is also present, and vice versa.

If no such string exists, return an empty string.

## Example

```text
Example 1:
Input: "BbCcXxY"
Expected Output: "BbCcXx"
Justification: Here, "BbCcXx" is the longest substring where each letter's
uppercase and lowercase forms are present.

Example 2:
Input: "aZAbcD"
Expected Output: ""
Justification: There is no contiguous substring where each character exists
in both its uppercase and lowercase forms.

Example 3:
Input: "qQwWeErR"
Expected Output: "qQwWeErR"
Justification: The entire string is the longest nice substring since
every letter exists in both uppercase and lowercase forms.
```

## Constraints

```text
1 <= s.length <= 100
s consists of uppercase and lowercase English letters.
```
