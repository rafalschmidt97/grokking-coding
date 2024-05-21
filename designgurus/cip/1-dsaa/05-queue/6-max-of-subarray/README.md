# Problem 6: Max of All Subarrays of Size 'k'

## Problem Statement

Given an integer array and an integer k, design an algorithm to find the maximum
for each and every contiguous subarray of size k.

## Example

```text
Input: array = [1, 2, 3, 1, 4, 5, 2, 3, 6], k = 3
Output: [3, 3, 4, 5, 5, 5, 6]
Description: Here, subarray 1,2,3 has maximum 3, 2,3,1
has maximum 3, 3,1,4 has maximum 4, 1,4,5 has maximum 5, 4,5,2
has maximum 5, 5,2,3 has maximum 5, and 2,3,6 has maximum 6.

Input: array = [8, 5, 10, 7, 9, 4, 15, 12, 90, 13], k = 4
Output: [10, 10, 10, 15, 15, 90, 90]
Description: Here, the maximum of each subarray of
size 4 are 10, 10, 10, 15, 15, 90, 90 respectively.

Input: array = [1,2,3,4,5], k = 3
Output: [3, 4, 5]
Description: Here, the maximum of each subarray of size 3 are 3, 4, 5 respectively.
```

## Constraints

```text
1 <= arr.length <= 10^5 -10^4 <= arr[i] <= 10^4 1 <= k <= arr.length
```
