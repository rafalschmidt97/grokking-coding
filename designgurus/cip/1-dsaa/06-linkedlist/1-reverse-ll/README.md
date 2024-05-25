# Problem 1: Reverse Linked List (easy)

## Problem Statement

Given the head of a singly linked list, your task is to reverse the list and
return its head. The singly linked list has nodes, and each node contains an
integer and a pointer to the next node. The last node in the list points to
null, indicating the end of the list.

## Example

```text
Example 1:
Input: [3, 5, 2]
Expected Output: [2, 5, 3]
Justification: Reversing the list [3, 5, 2] gives us [2, 5, 3].
Example 2:
Input: [7]
Expected Output: [7]
Justification: Since there is only one element in the list, the reversed list remains the same.
Example 3:
Input: [-1, 0, 1]
Expected Output: [1, 0, -1]
Justification: The list is reversed, so the elements are in the order [1, 0, -1].
```

## Constraints

```text
The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000
```
