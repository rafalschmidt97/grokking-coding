# Frequency of the Most Frequent Element (medium)

> Feels hard.

## Problem Statement

You are given an nums array containing n integers and an integer k. In a single
operation, you can choose any index i and increment the nums[i] by 1.

Return the maximum possible frequency of any element of nums after performing at
most k operations.

## Example

```text
Example 1:
Input: nums = [1, 2, 3], k = 3
Expected Output: 3
Explanation: We can increment the number 1 two times and 2 one time.
The final array will be [3, 3, 3]. Now, the number 3 appears 3 times
in the array [3, 3, 3].

Example 2:
Input: nums = [4, 4, 4, 1], k = 2
Expected Output: 3
Explanation: We can increment the number 1 two times (1 -> 2 -> 3).
Now, the number 4 still appears 3 times, which is the maximum frequency
that can be achieved with 2 operations.

Example 3:
Input: nums = [10, 9, 9, 4, 8], k = 5
Expected Output: 4
Explanation: We can increment the number 9 one time and the number 8 two
times (9 -> 10, 9 -> 10; 8 -> 9 -> 10). The number 10 will then appear
4 times in the array [10, 10, 10, 4, 10].
```

## Constraints

```text
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
1 <= k <= 10^5
```
