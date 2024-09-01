# Apple Redistribution into Boxes (easy)

You are given an array apple of size n, where the apple[i] represents the number
of apples in ith pack. You are also given an array capacity of size m, where
capacity[j] is a number of apples that can be stored in the jth box.

Return the minimum number of boxes you need to use to put these all n packs of
apples into boxes.

Note: You are allowed to distribute apples from the same pack into different
boxes.

## Problem Statement

## Example

```text
Example 1:
Input: apple = [2, 3, 1], capacity = [4, 2, 5, 1]
Expected Output: 2
Explanation: Box 1 can take apples from packs 1 and 2 partially
(totaling 5 apples), and Box 2 can take the rest of 2 apples.

Example 2:
Input: apple = [4, 5, 6], capacity = [5, 10]
Expected Output: 2
Explanation: Box 1 can take apples from packs 1 and 2 partially
(totaling 5 apples), and Box 2 can take the rest of pack 2 and all of pack 3 apples.

Example 3:
Input: apple = [1, 2, 5, 6], capacity = [2, 3, 7, 4, 5, 2, 4]
Expected Output: 3
Explanation: We can use boxes of size 7, 5, and 2 to pack all apples in boxes.
```

## Constraints

```text
1 <= n == apple.length <= 50
1 <= m == capacity.length <= 50
1 <= apple[i], capacity[i] <= 50
The input is generated such that it's possible to
redistribute packs of apples into boxes.
```
