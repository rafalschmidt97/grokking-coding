# Top 'K' Frequent Numbers (medium)

## Problem Statement

Given an unsorted array of numbers, find the top ‘K’ frequently occurring
numbers in it.

## Example

```text
Example 1:
Input: [1, 3, 5, 12, 11, 12, 11], K = 2
Output: [12, 11]
Explanation: Both '11' and '12' apeared twice.

Example 2:
Input: [5, 12, 11, 3, 11], K = 2
Output: [11, 5] or [11, 12] or [11, 3]
Explanation: Only '11' appeared twice, all other numbers appeared once.
```

## Constraints

```text
1 <= nums.length <= 10^5
-105 <= nums[i] <= 10^5
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
```
