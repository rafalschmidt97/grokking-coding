# Solution Walkthrough

The goal is to create two sets of integers (`arr1` and `arr2`) that meet the
given divisibility requirements. To efficiently find the smallest possible
maximum integer across both sets, we employ a binary search technique. This
method ensures we can quickly determine the minimum valid maximum value by
progressively narrowing down the range of potential values.

### High-Level Strategy:

We avoid generating all possible combinations of arrays by instead focusing on a
binary search approach. By setting initial bounds for the search range and
checking whether the midpoint satisfies the problem’s constraints, we can
iteratively adjust the range until we find the smallest valid maximum integer.

### Step-by-Step Algorithm:

1. **Initialization of Variables**:

   - Set the lower bound (`low`) to the sum of `uniqueCnt1` and `uniqueCnt2` to
     guarantee that there are enough integers available to meet the minimum
     requirements for both sets.
   - Set the upper bound (`high`) to the product of `divisor1`, `divisor2`,
     `uniqueCnt1`, and `uniqueCnt2`. This acts as a generous upper bound for the
     binary search.

2. **Calculate Least Common Multiple (LCM)**:

   - Compute the LCM of `divisor1` and `divisor2` using the greatest common
     divisor (GCD). This helps us determine how many numbers are divisible by
     both divisors, which is crucial for avoiding overlaps between the two sets.

3. **Performing Binary Search**:
   - Perform a binary search within the range between `low` and `high`. In each
     iteration:
     - Compute the midpoint (`mid`) of the current range.
     - Calculate how many numbers up to `mid` are divisible by both `divisor1`
       and `divisor2` (`countBoth`).
4. **Validity Check**:

   - Check if the current midpoint can accommodate both sets:
     - Ensure there are enough integers left after removing those divisible by
       both divisors to fill both sets.
     - Ensure that the remaining numbers not divisible by `divisor1` are
       sufficient to fill `uniqueCnt1`.
     - Ensure that the remaining numbers not divisible by `divisor2` are
       sufficient to fill `uniqueCnt2`.

5. **Adjusting Search Bounds**:

   - If the current midpoint satisfies all conditions, reduce the upper bound
     (`high`) to `mid - 1` to continue searching for a smaller maximum.
   - If the conditions are not met, increase the lower bound (`low`) to
     `mid + 1` to eliminate smaller values.

6. **Result Determination**:
   - The loop continues until `low` exceeds `high`. At this point, the smallest
     valid maximum integer is stored in `low`, which becomes the result.

### Algorithm Walkthrough with Example:

Let’s take the input: `uniqueCnt1 = 3`, `divisor1 = 3`, `uniqueCnt2 = 4`,
`divisor2 = 4`.

1. **Initialization**:

   - Set `low = 7` and `high = 144`. Calculate `lcm = 12`.

2. **Binary Search**:

   **Iteration 1**:

   - `mid = 75`, `countBoth = 6`.
   - Check conditions:
     - `mid - countBoth = 69` (Condition 1 satisfied).
     - `mid - (mid / 3) = 50` (Condition 2 satisfied).
     - `mid - (mid / 4) = 57` (Condition 3 satisfied).
   - Conditions met, update `high = 74`.

   **Iteration 2**:

   - `mid = 40`, `countBoth = 3`.
   - Check conditions:
     - `mid - countBoth = 37` (Condition 1 satisfied).
     - `mid - (mid / 3) = 27` (Condition 2 satisfied).
     - `mid - (mid / 4) = 30` (Condition 3 satisfied).
   - Conditions met, update `high = 39`.

   **Iteration 3**:

   - `mid = 23`, `countBoth = 1`.
   - Check conditions:
     - `mid - countBoth = 22` (Condition 1 satisfied).
     - `mid - (mid / 3) = 15` (Condition 2 satisfied).
     - `mid - (mid / 4) = 17` (Condition 3 satisfied).
   - Conditions met, update `high = 22`.

   **Iteration 4**:

   - `mid = 14`, `countBoth = 1`.
   - Check conditions:
     - `mid - countBoth = 13` (Condition 1 satisfied).
     - `mid - (mid / 3) = 9` (Condition 2 satisfied).
     - `mid - (mid / 4) = 10` (Condition 3 satisfied).
   - Conditions met, update `high = 13`.

   **Iteration 5**:

   - `mid = 10`, `countBoth = 0`.
   - Check conditions:
     - `mid - countBoth = 10` (Condition 1 satisfied).
     - `mid - (mid / 3) = 7` (Condition 2 satisfied).
     - `mid - (mid / 4) = 8` (Condition 3 satisfied).
   - Conditions met, update `high = 9`.

   **Iteration 6**:

   - `mid = 8`, `countBoth = 0`.
   - Check conditions:
     - `mid - countBoth = 8` (Condition 1 satisfied).
     - `mid - (mid / 3) = 6` (Condition 2 satisfied).
     - `mid - (mid / 4) = 6` (Condition 3 satisfied).
   - Conditions met, update `high = 7`.

   **Iteration 7**:

   - `mid = 7`, `countBoth = 0`.
   - Check conditions:
     - `mid - countBoth = 7` (Condition 1 satisfied).
     - `mid - (mid / 3) = 5` (Condition 2 satisfied).
     - `mid - (mid / 4) = 5` (Condition 3 satisfied).
   - Conditions met, update `high = 6`.

3. **Conclusion**:
   - After the 7th iteration, the binary search terminates, and the smallest
     valid maximum integer is 7. This is the minimum maximum value that
     satisfies the problem constraints.
