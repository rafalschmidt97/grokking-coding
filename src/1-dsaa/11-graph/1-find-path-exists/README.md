# Problem 1: Find if Path Exists in Graph(easy)

## Problem Statement

Given an undirected graph, represented as a list of edges. Each edge is
illustrated as a pair of integers [u, v], signifying that there's a mutual
connection between node u and node v.

You are also given starting node start, and a destination node end, return true
if a path exists between the starting node and the destination node. Otherwise,
return false.

## Example

```text
Example 1:
Input: 4, [[0,1],[1,2],[2,3]], 0, 3
Expected Output: true
Justification: There's a path from node 0 -> 1 -> 2 -> 3.

Example 2:
Input: 4, [[0,1],[2,3]], 0, 3
Expected Output: false
Justification: Nodes 0 and 3 are not connected, so no path exists between them.

Example 3:
Input: 5, [[0,1],[3,4]], 0, 4
Expected Output: false
Justification: Nodes 0 and 4 are not connected in any manner.
```

## Constraints

```text
1 <= n <= 2 * 10^5
0 <= edges.length <= 2 * 10^5
edges[i].length == 2
0 <= ui, vi <= n - 1
ui != vi
0 <= source, destination <= n - 1
There are no duplicate edges.
There are no self edges.
```
