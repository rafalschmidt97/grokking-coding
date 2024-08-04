# Take Gifts From the Richest Pile(easy)

## Problem Statement

You're presented with several piles of gifts, with each pile containing a
certain number of gifts. Every second, you'll engage in the following activity:

Pick the pile that contains the highest number of gifts. If multiple piles share
this distinction, you can select any of them. Compute the square root of the
number of gifts in the selected pile, and then leave behind that many gifts
(rounded down). Take all the other gifts from this pile. You'll do this for "k"
seconds. The objective is to find out how many gifts would still remain after
these "k" seconds.

## Example

```text
Input: gifts = [4, 9, 16], k = 2
Expected Output: 11
Justification:
Take from third pile (16 gifts): leave sqrt(16) = 4 gifts, take 12. Remaining gifts = [4, 9, 4]
Take from second pile (9 gifts): leave sqrt(9) = 3 gifts, take 6. Remaining gifts = [4, 3, 4]

Input: gifts = [1, 2, 3], k = 1
Expected Output: 4
Justification:
Take from third pile (3 gifts): leave sqrt(3) = 1 gift (rounded down), take 2. Remaining gifts = [1, 2, 1]

Input: gifts = [25, 36, 49], k = 3
Expected Output: 18
Justification:
Take from third pile (49 gifts): leave sqrt(49) = 7 gifts, take 42. Remaining gifts = [25, 36, 7]
Take from second pile (36 gifts): leave sqrt(36) = 6 gifts, take 30. Remaining gifts = [25, 6, 7]
Take from first pile (25 gifts): leave sqrt(25) = 5 gifts, take 20. Remaining gifts = [5, 6, 7]

```

## Constraints

```text
1 <= gifts.length <= 10^3
1 <= gifts[i] <= 10^9
1 <= k <= 10^3
```
