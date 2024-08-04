# Problem 2: Matrix Diagonal Sum (easy)

## Problem Statement

Given a square matrix (2D array), calculate the sum of its two diagonals.

The two diagonals in consideration are the primary diagonal that spans from the
top-left to the bottom-right and the secondary diagonal that spans from
top-right to bottom-left. If a number is part of both diagonals (which occurs
only for odd-sized matrices), it should be counted only once in the sum.

## Example

```text
Example 1:
Input:
[[1,2,3],
 [4,5,6],
 [7,8,9]]
Expected Output: 25
Justification: Summing up the two diagonals (1+5+9+3+7), we get 25. o
Please note that the element at [1][1] = 5 is counted only once.

Example 2:
Input:
[[1,0],
 [0,1]]
Expected Output: 2
Justification: The sum of the two diagonals is 1+1 = 2.

Example 3:
Input:
[[5]]
Expected Output: 5
Justification: Since there's only one element, it is the sum itself.
```

## Constraints

```text
n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100
```
