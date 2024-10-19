# Median of Two Sorted Arrays (medium)

## Problem Statement

Given two sorted arrays, nums1 and nums2 of different sizes in the ascending
order, return the median of two sorted arrays after merging them.

The median is the middle value in an ordered set, or the average of the two
middle values if the set has an even number of elements.

## Example

```text
Example 1:
Input: [1, 3, 5] and [2, 4, 6]
Expected Output: 3.5
Justification: When merged, the array becomes [1, 2, 3, 4, 5, 6].
The median is the average of the middle two values, (3 + 4) / 2 = 3.5.

Example 2:
Input: [1, 1, 1] and [2, 2, 2]
Expected Output: 1.5
Justification: The merged array is [1, 1, 1, 2, 2, 2].
The median is (1 + 2) / 2 = 1.5.

Example 3:
Input: [2, 3, 5, 8] and [1, 4, 6, 7, 9]
Expected Output: 5
Justification: The merged array would be [1,2,3,4,5,6,7,8,9].
```

## Constraints

```text
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
```
