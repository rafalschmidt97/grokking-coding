# Sort Array by Increasing Frequency (easy)

## Problem Statement

Given an array `nums` containing the integers, return the resultant array after
sorting it in increasing order based on the frequency of the values. If two
numbers have the same frequency, they should be sorted in descending numerical
order.

## Example

```text
Example 1:

Input: nums = [4, 4, 6, 2, 2, 2]
ExpectedOutput: [6, 4, 4, 2, 2, 2]
Justification: Here, '6' appears once, '4' appears twice, and '2' appears
three times. Thus, numbers are first sorted by frequency and then by
value when frequencies tie.

Example 2:
Input: nums = [0, -1, -1, -1, 5]
ExpectedOutput: [5, 0, -1, -1, -1]
Justification: '5' and '0' appears once, and '-1' appears three times.
After sorting by frequency and resolving ties by sorting in descending order,
the result is obtained.

Example 3:
Input: nums = [10, 10, 10, 20, 20, 30]
ExpectedOutput: [30, 20, 20, 10, 10, 10]
Justification: Here, '30' has the lowest frequency, followed by '20',
and '10' has the highest frequency. They are sorted accordingly.
```

## Constraints

```text
1 <= nums.length <= 100
-100 <= nums[i] <= 100
```
