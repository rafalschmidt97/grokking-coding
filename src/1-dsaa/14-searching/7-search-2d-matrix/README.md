# Search a 2D Matrix II (medium)

## Problem Statement

Given a 2D grid of size m x n matrix containing integers, and integer target,
return true if target value exists in the matrix. Otherwise, return false.

The matrix has the following properties:

Values in each column are sorted in non-decreasing order from top to bottom.
Values in each row are sorted in non-decreasing order from left to right.

## Example

```text
Example 1:
Input: target = 5, matrix =
[[1,2,3],
 [4,5,6],
 [7,8,9]]
Expected Output: true
Justification: The number 5 is located in the second row and second
column of the matrix, thus the output is true.

Example 2:
Input: target = 19, matrix =
[[10,11,12,13],
 [11,12,13,17],
 [14,19,22,24],
 [22,23,24,25]]
Expected Output: true
Justification: 19 is present in the third row and second column
, confirming its presence in the matrix.

Example 3:
Input: target = 6, matrix =
[[1,3,5],
 [7,9,11],
 [13,15,17]]
Expected Output: false
Justification: 6 does not appear anywhere in the matrix, so the
function should return false.
```
