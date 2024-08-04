# Kth Smallest Element in a BST (medium)

## Problem Statement

Given a root node of the Binary Search Tree (BST) and integer 'k'. Return the
Kth smallest element among all node values of the binary tree.

## Example

```text
Example 1:
Input:

    8
   / \
  3   10
 / \    \
1   6    14
   /  \  /
  4   7  13
k = 4
Expected Output: 6
Justification: The in-order traversal of the tree is [1, 3, 4, 6, 7, 8, 10, 13, 14].
The 4th element in this sequence is 6.

Example 2:
Input:

    5
   / \
  2   6
 /
1
k = 3
Expected Output: 5
Justification: The in-order traversal of the tree is [1, 2, 5, 6]. The 3rd element in this sequence is 5.

Example 3:
Input:

1
 \
  3
 /
2
k = 2
Expected Output: 2
Justification: The in-order traversal of the tree is [1, 2, 3]. The 2nd element in this sequence is 2.


```

## Constraints

```text
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104
```
