# Remove Duplicate Letters (medium)

## Problem Statement

Given a string s, remove all duplicate letters from the input string while
maintaining the original order of the letters.

Additionally, the returned string should be the smallest in lexicographical
order among all possible results.

A string is in the smallest lexicographical order if it appears first in a
dictionary. For example, "abc" is smaller than "acb" because "abc" comes first
alphabetically.

## Example

```text
1.
  Input: "babac"
  Expected Output: "abc"
  Justification:
  After removing 1 b and 1 a from the input string, we can get bac, and abc
strings. The final answer is 'abc', which is the smallest lexicographical
string without duplicate letters.

2.
  Input: "zabccde"
  Expected Output: "zabcde"
  Justification: Removing one of the 'c's forms 'zabcde', the smallest string
in lexicographical order without duplicates.

3.
  Input: "mnopmn"
  Expected Output: "mnop"
  Justification: Removing the second 'm' and 'n' gives 'mnop', which is the
smallest possible string without duplicate characters.
```

## Constraints

```text
1 <= s.length <= 104
s consists of lowercase English letters.
```
