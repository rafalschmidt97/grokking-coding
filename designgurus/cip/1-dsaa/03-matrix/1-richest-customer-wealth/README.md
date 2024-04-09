# Problem 1: Richest Customer Wealth (easy)

## Problem Statement

You are given an `m x n` matrix accounts where `accounts[i][j]` is the amount of
money the `ith` customer has in the `jth` bank. Return the wealth that the
richest customer has. Imagine every customer has multiple bank accounts, with
each account holding a certain amount of money. The total wealth of a customer
is calculated by summing all the money across all their multiple.

## Example

```text
Example 1:

Input: [[5,2,3],[0,6,7]]
Expected Output: 13
Justification: The total wealth of the first customer is 10 and of the second customer is 13. So, the output is 13 as it's the maximum among all customers.
Example 2:

Input: [[1,2],[3,4],[5,6]]
Expected Output: 11
Justification: Total wealth for each customer is [3, 7, 11]. Maximum of these is 11.
Example 3:

Input: [[10,100],[200,20],[30,300]]
Expected Output: 330
Justification: Total wealth for each customer is [110, 220, 330]. The wealthiest customer has 330.
```

## Constraints

```text
m == accounts.length
n == accounts[i].length
1 <= m, n <= 50
1 <= accounts[i][j] <= 100
```
