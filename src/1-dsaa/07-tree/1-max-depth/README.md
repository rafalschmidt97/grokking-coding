# Maximum Depth (or Height) of Binary Tree (easy)

## Problem Statement

Determine the depth (or height) of a binary tree, which refers to the number of
nodes along the longest path from the root node to the farthest leaf node. If
the tree is empty, the depth is 0.

## Example

```text
Input
     1
    / \
   2   3
  / \
 4   5
Output: 3
Explanation: The longest path is 1->2->4 or 1->2->5 with 3 nodes.

Input:
1
 \
  2
   \
    3
Expected Output: 3
Justification: There's only one path 1->2->3 with 3 nodes.
Input:
    1
   / \
  2   3
 / \
4   7
     \
      9
Expected Output: 4
Justification: The longest path is 1->2->7->9 with 4 nodes.

```

## Constraints

```text
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
```
