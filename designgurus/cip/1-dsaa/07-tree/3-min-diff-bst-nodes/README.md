# Minimum Difference Between BST Nodes (easy)

## Problem Statement

Given a Binary Search Tree (BST), you are required to find the smallest
difference between the values of any two different nodes.

In a BST, the nodes are arranged in a manner where the value of nodes on the
left is less than or equal to the root, and the value of nodes on the right is
greater than the root.

## Example

```text
Example 1:

Input:
    4
   / \
  2   6
 / \
1   3
Expected Output: 1
Justification: The pairs (1,2), (2,3), or (3,4) have the smallest difference of 1.
Example 2:

Input:
    10
   /  \
  5   15
 / \    \
2   7    18
Expected Output: 2
Justification: The pair (5,7) has the smallest difference of 2.
Example 3:

Input:
  40
   \
    70
   /  \
  50  90
Expected Output: 10
Justification: The pair (40,50) has the smallest difference of 10.
```

## Constraints

```text
The number of nodes in the tree is in the range [2, 104].
0 <= Node.val <= 105
```
