# Find the Median of a Number Stream (medium)

## Problem Statement

Design a class to calculate the median of a number stream. The class should have
the following two methods:

insertNum(int num): stores the number in the class findMedian(): returns the
median of all numbers inserted in the class If the count of numbers inserted in
the class is even, the median will be the average of the middle two numbers.

## Example

```text
1. insertNum(3)
2. insertNum(1)
3. findMedian() -> output: 2
4. insertNum(5)
5. findMedian() -> output: 3
6. insertNum(4)
7. findMedian() -> output: 3.5
```

## Constraints

```text
-10^5 <= num <= 10^5
There will be at least one element in the data structure before calling findMedian.
At most 5 * 104 calls will be made to insertNum and findMedian.
```
