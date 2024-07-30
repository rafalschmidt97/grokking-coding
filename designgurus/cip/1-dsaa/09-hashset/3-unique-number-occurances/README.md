# Problem 3: Unique Number of Occurrences (easy)

## Problem Statement

Given an array of integers, determine if the number of times each distinct
integer appears in the array is unique.

In other words, the occurrences of each integer in the array should be distinct
from the occurrences of every other integer.

## Example

```text
1.
  Input: [4, 5, 4, 6, 6, 6]
  Expected Output: true
  Justification: The number 4 appears 2 times, 5 appears 1 time,
  and 6 appears 3 times. All these occurrences (1, 2, 3) are unique.

2.
  Input: [7, 8, 8, 9, 9, 9, 10, 10]
  Expected Output: false
  Justification: The number 7 appears 1 time, 8 appears 2 times,
  9 appears 3 times, and 10 appears 2 times. The occurrences are
  not unique since the number 2 appears twice.


3.
  Input: [11, 12, 13, 14, 14, 15, 15, 15]
  Expected Output: false
  Justification: The number 11 appears 1 time, 12 appears 1 time,
  13 appears 1 time, 14 appears 2 times, and 15 appears 3 times.
  These occurrences are not unique.
```

## Constraints

```text
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
```
