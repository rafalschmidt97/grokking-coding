# Minimize the Maximum of Two Arrays (medium)

## Problem Statement

Given divisor1, divisor2, uniqueCnt1, and uniqueCnt2 integers, find the smallest
possible maximum integer that could be present in either array after they are
filled according to the below conditions.

You can take two arrays arr1 and arr2 which are initially empty. arr1 contains
total uniqueCnt1 different positive integers, each of them is not divisible by
divisor1. arr2 contains total uniqueCnt2 different positive integers, each of
them is not divisible by divisor2. There are no common integers in both arrays.

## Example

```text
Example 1:
Input: uniqueCnt1 = 2, divisor1 = 2, uniqueCnt2 = 2, divisor2 = 3
Expected Output: 4
Explanation: The optimal arrays could be arr1 = [1, 3] (numbers not divisible
by 2) and arr2 = [2, 4] (numbers not divisible by 3). The maximum number among
both arrays is 4.

Example 2:
Input: uniqueCnt1 = 3, divisor1 = 3, uniqueCnt2 = 4, divisor2 = 4
Expected Output: 7
Explanation: Possible arrays are arr1 = [1, 2, 4] and arr2 = [3, 5, 6, 7].
The highest integer used is 7.

Example 3:
Input: uniqueCnt1 = 1, divisor1 = 7, uniqueCnt2 = 1, divisor2 = 10
Expected Output: 2
Explanation: We can use arr1 = [1] (since it's not divisible by 7) and arr2 = [2]
(since it's not divisible by 10). The highest integer here is 2.
```

## Constraints

```text
2 <= divisor1, divisor2 <= 10^5
1 <= uniqueCnt1, uniqueCnt2 < 10^9
2 <= uniqueCnt1 + uniqueCnt2 <= 10^9
```
