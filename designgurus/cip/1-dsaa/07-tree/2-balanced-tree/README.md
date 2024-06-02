# Balanced Binary Tree (easy)

## Problem Statement

Determine if a binary tree is height-balanced.

A binary tree is considered height-balanced if, for each node, the difference in
height between its left and right subtrees is no more than one.

## Example

```text

Input:
    3
   / \
  9  20
     / \
    15  7
Expected Output: true
Justification: Every node in the tree has a left and right subtree depth difference of either 0 or 1.
Input:
        1
      /   \
     2     2
    / \   / \
   3   3 3   3
  / \       / \
 4   4     4   4
Expected Output: true
Justification: Each node in the tree has a left and right subtree depth difference of either 0 or 1.
Input:
       1
      / \
     2   5
    /
   3
  /
 4
Expected Output: false
Justification: The root node has a left subtree depth of 3 and right subtree depth of 1.
The difference (3 - 1 = 2) exceeds 1, hence the tree is not balanced.
```

## Constraints

```text
The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104
```
