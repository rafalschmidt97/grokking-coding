# Minimum Cost to Connect Sticks(medium)

## Problem Statement

Given a collection of sticks with different lengths. To combine any two sticks,
there's a cost involved, which is equal to the sum of their lengths.

Connect all the sticks into a single one with the minimum possible cost.
Remember, once two sticks are combined, they form a single stick whose length is
the sum of the lengths of the two original sticks.

## Example

```text
Input: [2, 4, 3]
Expected Output: 14
Justification: Combine sticks 2 and 3 for a cost of 5. Now, we have sticks [4,5].
Combine these at a cost of 9. Total cost = 5 + 9 = 14.

Input: [1, 8, 2, 5]
Expected Output: 30
Justification: Combine sticks 1 and 2 for a cost of 3. Now, we have sticks [3, 8, 5].
Combine 3 and 5 for a cost of 8. Now, we have sticks [8,8]. Combine these for a cost of 16. Total cost = 3 + 8 + 16 = 27.

Input: [5, 5, 5, 5]
Expected Output: 40
Justification: Combine two 5s for a cost of 10. Do this again for another cost of 10.
Now, we have two sticks of 10 each. Combine these for a cost of 20. Total cost = 10 + 10 + 20 = 40.

```

## Constraints

```text
1 <= sticks.length <= 10^4
1 <= sticks[i] <= 10^4
```
