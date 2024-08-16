# Problem 2: Number of Provinces (medium)

## Problem Statement

Imagine a country with several cities. Some cities have direct roads connecting
them, while others might be connected through a sequence of intermediary cities.
Using a matrix representation, `if matrix[i][j]` holds the value 1, it indicates
that city i is directly linked to city j and vice versa. If it holds 0, then
there's no direct link between them.

Determine the number of separate city clusters (or provinces).

A province is defined as a collection of cities that can be reached from each
other either directly or through other cities in the same province.

## Example

```text
1
Input: [[1,1,0],[1,1,0],[0,0,1]]
Expected Output: 2
Justification: There are two provinces: cities 0 and 1 form one province,
and city 2 forms its own province.

2. Input: [[1,0,0,1],[0,1,1,0],[0,1,1,0],[1,0,0,1]]
Expected Output:2
Justification: There are two provinces: cities 0 and 3 are
interconnected forming one province, and cities 1 and 2 form another.

3
Input: [[1,0,0],[0,1,0],[0,0,1]]
Expected Output: 3
Justification: Each city stands alone and is not connected to any other city.
Thus, we have three provinces.
```

## Constraints

```text
1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] is 1 or 0.
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]
```
