# Maximum Count of Positive Integer and Negative Integer (easy)

## Problem Statement

Given an array nums sorted in increasing order, return the maximum between the
count of positive integers and the count of negative integers.

Note: 0 is neither positive nor negative.

## Example

```text
Example 1:
Input: nums = [-4, -3, -1, 0, 1, 3, 5, 7]
Expected Output: 4
Justification: The array contains three negative integers (-4, -3, -1) and four positive
integers (1, 3, 5, 7). The maximum count between negatives and positives is 4.

Example 2:
Input: nums = [-8, -7, -5, -4, 0, 0, 0]
Expected Output: 4
Justification: Here, there are four negative integers (-8, -7, -5, -4)
and no positives. Thus, the maximum count is 4.

Example 3:
Input: nums = [0, 2, 2, 3, 3, 3, 4]
Expected Output: 6
Justification: This input array includes zero negative integers and six
positives (2, 2, 3, 3, 3, 4). Hence, the maximum count is 6.
```

## Constraints

```text
1 <= nums.length <= 2000
-2000 <= nums[i] <= 2000
nums is sorted in a non-decreasing order.
```
