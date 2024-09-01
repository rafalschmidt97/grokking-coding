# Reduction Operations to Make the Array Elements Equal (medium)

## Problem Statement

Given an array of integers nums, return the number of operations required to
make all elements in nums equal. To perform one operation, you can follow the
below steps:

- Select the maximum element of nums. If there are multiple occurrences of the
  maximum element, choose the element which has lowest index i.
- Select the second largest element of nums.
- Replace the element at index i with the second largest element.

## Example

```text
Example 1:
Input: [3, 5, 5, 2]
Expected output: 5
Justification:
- The largest element is 5. Reducing both 5s to 3 requires two operations.
- Update array will be [3, 3, 3, 2].
- Three more operations are needed to reduce the 3 to 2. The updated array
will be [2, 2, 2, 2].
- A total five operations make all elements equal to 2.

Example 2:
Input: [11, 9, 7, 5, 3]
Expected output: 10
Justification:
- Each number needs to be reduced stepwise to the next smaller number until all
are equal to the smallest number 3.
- 1 operation is required to convert 11 to 9.
The updated array will be [9, 9, 7, 5, 3].
- 2 operations are required to convert 9 to 7.
The updated array will be [7, 7, 7, 5, 3].
- 3 operations are required to convert 7 to 5.
The updated array will be [5, 5, 5, 5, 3].
- 4 operations are required to convert 5 to 3.
The updated array will be [3, 3, 3, 3, 3].
- Total numbers of operations: 1 + 2 + 3 + 4 = 10.

Example 3:
Input: [8, 8, 8, 8]
Expected output: 0
Justification: All elements are already equal, so no operations are needed.
```

## Constraints

```text
1 <= nums.length <= 5 * 10^4
1 <= nums[i] <= 5 * 10^4
```
