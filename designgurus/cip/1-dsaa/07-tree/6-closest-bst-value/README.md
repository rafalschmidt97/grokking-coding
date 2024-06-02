# Closest Binary Search Tree Value (medium)

## Problem Statement

Given a binary search tree (BST) and a target number, you need to find a node
value in the BST that is closest to the given target. A BST is a tree where for
every node, the values in the left subtree are smaller than the node, and the
values in the right subtree are greater.

## Example

```text
Example 1:

Input:

    Tree:
       5
     /   \
    3     8
   / \   / \
  1   4 6   9
Target: 6.4

Expected Output: 6

Justification: The values 6 and 8 are the closest numbers to 6.4 in the tree. Among them, 6 is closer.

Example 2:

Input:

    Tree:
       20
     /    \
    10     30
Target: 25

Expected Output: 20

Justification: 20 and 30 are the closest numbers to 25. However, 20 is closer than 30.

Example 3:

Input:

    Tree:
       2
     /   \
    1     3
Target: 2.9

Expected Output: 3

Justification: 3 is the closest value to 2.9 in the tree.
```

## Constraints

```text
The number of nodes in the tree is in the range [1, 104].
0 <= Node.val <= 109
-109 <= target <= 109
```
