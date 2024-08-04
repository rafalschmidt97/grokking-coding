# Problem 1: Counting Elements (easy)

## Problem Statement

Given a list of integers, determine the count of numbers for which there exists
another number in the list that is greater by exactly one unit.

In other words, for each number x in the list, if x + 1 also exists in the list,
then x is considered for the count.

## Example

```text
Example 1:
Input: [4, 3, 1, 5, 6]
Expected Output: 3
Justification: The numbers 4, 3, and 5 have 5, 4, and 6 respectively in the list, which are greater by exactly one unit.

Example 2:
Input: [7, 8, 9, 10]
Expected Output: 3
Justification: The numbers 7, 8, and 9 have 8, 9, and 10 respectively in the list, which are greater by exactly one unit.

Example 3:
Input: [11, 13, 15, 16]
Expected Output: 1
Justification: Only the number 15 has 16 in the list, which is greater by exactly one unit.
```

## Constraints

```text
1 <= arr.length <= 1000
0 <= arr[i] <= 1000
```
