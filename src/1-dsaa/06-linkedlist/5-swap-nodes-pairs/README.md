# Problem 5: Swap Nodes in Pairs (medium)

## Problem Statement

Given a singly linked list, swap every two adjacent nodes and return the head of
the modified list.

If the total number of nodes in the list is odd, the last node remains in place.
Every node in the linked list contains a single integer value.

## Example

```text
Input: [1, 2, 3, 4]
Output: [2, 1, 4, 3]
Justification: Pairs (1,2) and (3,4) are swapped.

Input: [7, 8, 9, 10, 11]
Output: [8, 7, 10, 9, 11]
Justification: Pairs (7,8) and (9,10) are swapped.
11 remains in its place as it has no adjacent node to swap with.

Input: [5, 6]
Output: [6, 5]
Justification: The pair (5,6) is swapped.

```

## Constraints

```text
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
```
