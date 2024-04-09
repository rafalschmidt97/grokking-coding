# Problem 4: Find the Highest Altitude(easy)

## Problem Statement

Given an array of integers representing a journey on a bike, wherein each number
indicates a climb or descent in altitude, measured as a gain or loss of height.

The rider starts at altitude 0 and determines the highest altitude the biker
reaches at any point during the journey.

## Example

```text
Example 1
Input: [-5, 1, 5, 0, -7]
Expected Output: 1
Justification: The altitude changes are [-5, -4, 1, 1, -6], where 1 is the highest altitude reached.

Example 2
Input: [4, -3, 2, -1, -2]
Expected Output: 4
Justification: The altitude changes are [4, 1, 3, 2, 0], where 4 is the highest altitude reached.

Example 3
Input: [2, 2, -3, -1, 2, 1, -5]
Expected Output: 4
Justification: The altitude changes are [2, 4, 1, 0, 2, 3, -2], where 4 is the highest altitude reached.
```

## Constraints

- n == gain.length
- 1 <= n <= 100
- -100 <= gain[i] <= 100
