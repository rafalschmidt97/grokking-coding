# Problem 3: Left and Right Sum Differences (easy)

## Problem Statement

Given an input array of integers nums, find an integer array, let's call it
differenceArray, of the same length as an input integer array.

Each element of differenceArray, i.e., differenceArray[i], should be calculated
as follows: take the sum of all elements to the left of index i in array nums
(denoted as leftSum[i]), and subtract it from the sum of all elements to the
right of index i in array nums (denoted as rightSum[i]), taking the absolute
value of the result. Formally:

```text
differenceArray[i] = | leftSum[i] - rightSum[i] |
```

If there are no elements to the left/right of i, the corresponding sum should be
taken as 0.

## Examples

```text
Example 1:

Input: [2, 5, 1, 6]
Expected Output: [12, 5, 1, 8]
Explanation:
For i=0: |(0) - (5+1+6)| = |0 - 12| = 12
For i=1: |(2) - (1+6)| = |2 - 7| = 5
For i=2: |(2+5) - (6)| = |7 - 6| = 1
For i=3: |(2+5+1) - (0)| = |8 - 0| = 8
Example 2:

Input: [3, 3, 3]
Expected Output: [6, 0, 6]
Explanation:
For i=0: |(0) - (3+3)| = 6
For i=1: |(3) - (3)| = 0
For i=2: |(3+3) - (0)| = 6
Example 3:

Input: [1, 2, 3, 4, 5]
Expected Output: [14, 11, 6, 1, 10]
Explanation:
Calculations for each index i will follow the above-mentioned logic.
```

## Constraints

- 1 <= nums.length <= 1000
- 1 <= nums[i] <= 105
