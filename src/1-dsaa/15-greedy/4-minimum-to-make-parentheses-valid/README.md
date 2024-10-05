# Minimum Add to Make Parentheses Valid (medium)

## Problem Statement

Given a string str containing '(' and ')' characters, find the minimum number of
parentheses that need to be added to a string of parentheses to make it valid.

A valid string of parentheses is one where each opening parenthesis '(' has a
corresponding closing parenthesis ')' and vice versa. The goal is to determine
the least amount of additions needed to achieve this balance.

## Example

```text
Example 1:
Input: "(()"
Expected Output: 1
Justification: The string has two opening parentheses and one closing parenthesis.
Adding one closing parenthesis at the end will balance it.

Example 2:
Input: "))(("
Expected Output: 4
Justification: There are two closing parentheses at the beginning and two
opening at the end. We need two opening parentheses before the first closing
and two closing parentheses after the last opening to balance the string.

Example 3:
Input: "(()())("
Expected Output: 1
Justification: The string has three opening parentheses and three closing
parentheses, with an additional opening parenthesis at the end. Adding one
closing parenthesis at the end will balance it.
```

## Constraints

```text
1 <= s.length <= 1000
s[i] is either '(' or ')'.
```
