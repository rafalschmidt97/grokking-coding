# Knapsack problem

The **Knapsack problem** is a classic optimization problem in computer science
and combinatorial optimization. It involves selecting a subset of items with
given weights and values to maximize the total value without exceeding a
specified weight capacity.

### Problem Description

You are given:

- A set of `n` items, each with a weight (`w[i]`) and a value (`v[i]`).
- A knapsack that can carry a maximum weight of `W`.

The objective is to determine the items to include in the knapsack so that the
**total weight** does not exceed `W` and the **total value** is maximized.

The most common form is the **0/1 Knapsack problem**, where each item can either
be included or excluded.

### Dynamic Programming Approach

The **Dynamic Programming (DP)** approach is one of the most efficient solutions
for solving the 0/1 Knapsack problem. Here’s how you can implement it in Go:

#### Step-by-Step Explanation

1. Create a two-dimensional array `dp` where `dp[i][w]` represents the maximum
   value that can be obtained using the first `i` items and a knapsack with
   weight capacity `w`.
2. Populate the DP table:
   - If the item can fit (`w[i] <= w`), you have two options: either include it
     or not. Take the maximum value.
   - If the item can't fit, the value is the same as without the item.

#### Go Code Example

Here's an implementation of the 0/1 Knapsack problem using Dynamic Programming
in Go:

```go
package main

import "fmt"

// Knapsack function to solve the 0/1 Knapsack problem
func knapsack(values []int, weights []int, capacity int) int {
    n := len(values)
    dp := make([][]int, n+1)

    // Initialize DP table
    for i := range dp {
        dp[i] = make([]int, capacity+1)
    }

    // Populate the DP table
    for i := 1; i <= n; i++ {
        for w := 0; w <= capacity; w++ {
            if weights[i-1] <= w {
                dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
            } else {
                dp[i][w] = dp[i-1][w]
            }
        }
    }

    // The bottom-right corner of the table contains the result
    return dp[n][capacity]
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    values := []int{60, 100, 120}
    weights := []int{10, 20, 30}
    capacity := 50

    result := knapsack(values, weights, capacity)
    fmt.Printf("Maximum value in Knapsack = %d\n", result)
}
```

### Explanation

1. **Inputs**:
   - `values`: Array containing values of the items.
   - `weights`: Array containing weights of the items.
   - `capacity`: Maximum weight the knapsack can hold.
2. **DP Table**:
   - The `dp` table has dimensions `(n+1) x (capacity+1)`, where `n` is the
     number of items.
   - `dp[i][w]` stores the maximum value obtainable using the first `i` items
     and weight `w`.
3. **Filling the DP Table**:
   - If `weights[i-1] <= w`, choose the maximum between including the item or
     not.
   - Otherwise, carry forward the value without including the current item.

### Complexity

- **Time Complexity**: `O(n * W)` where `n` is the number of items and `W` is
  the capacity.
- **Space Complexity**: `O(n * W)` due to the DP table. It can be optimized to
  `O(W)` using a 1D array if needed.

### Optimized Approach

Instead of using a 2D array, you can use a 1D array to save space. Here's a
quick look at how to do that:

```go
func knapsackOptimized(values []int, weights []int, capacity int) int {
    n := len(values)
    dp := make([]int, capacity+1)

    // Populate the DP array
    for i := 0; i < n; i++ {
        for w := capacity; w >= weights[i]; w-- {
            dp[w] = max(dp[w], dp[w-weights[i]]+values[i])
        }
    }

    return dp[capacity]
}
```

This optimization works because each row of the DP table only depends on the
previous row, so we can use a 1D array to reduce space complexity to `O(W)`.

### Summary

- The **Knapsack problem** is a fundamental optimization problem that can be
  solved efficiently using **Dynamic Programming**.
- In Go, the problem can be implemented using either a 2D DP table or a
  space-optimized 1D array.
- The above implementations will help you solve the 0/1 Knapsack problem for
  small to medium-sized inputs effectively.

If you have more questions or want further explanation on any part of the
implementation, feel free to ask!

### Divide and Conquer Approach for 0/1 Knapsack

The Divide and Conquer strategy for the **0/1 Knapsack problem** involves
solving the problem recursively by breaking it down into smaller subproblems.
For each item, you have two choices:

1. **Include the item**: Add the value of the item to the result and decrease
   the capacity by the weight of the item.
2. **Exclude the item**: Move to the next item without adding anything to the
   value.

The final answer is the maximum value obtained between including or excluding
each item.

#### Recursive Formulation

Given a list of items and a capacity `W`:

- If an item’s weight is greater than `W`, you cannot include it.
- Otherwise, you recursively decide whether to include the item or not.

The recurrence relation can be represented as:

\[ V(i, w) = \max(V(i-1, w), v_i + V(i-1, w - w_i)) \]

Where:

- \( V(i, w) \) is the maximum value achievable with `i` items and weight `w`.
- `v_i` and `w_i` are the value and weight of item `i`.

### Go Code Example Using Divide and Conquer (Recursion)

Below is a Go implementation of the **0/1 Knapsack** using recursion:

```go
package main

import "fmt"

// knapsackRecursive solves the Knapsack problem using a divide and conquer approach (recursion)
func knapsackRecursive(values []int, weights []int, capacity int, n int) int {
    // Base case: no items left or capacity is 0
    if n == 0 || capacity == 0 {
        return 0
    }

    // If the weight of the nth item is more than the capacity, it cannot be included
    if weights[n-1] > capacity {
        return knapsackRecursive(values, weights, capacity, n-1)
    } else {
        // Maximize the value by including or excluding the nth item
        includeItem := values[n-1] + knapsackRecursive(values, weights, capacity-weights[n-1], n-1)
        excludeItem := knapsackRecursive(values, weights, capacity, n-1)
        return max(includeItem, excludeItem)
    }
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    values := []int{60, 100, 120}
    weights := []int{10, 20, 30}
    capacity := 50
    n := len(values)

    result := knapsackRecursive(values, weights, capacity, n)
    fmt.Printf("Maximum value in Knapsack = %d\n", result)
}
```

### Explanation

1. **Recursive Base Cases**:
   - When `n == 0` (no items left) or `capacity == 0` (no space left), the value
     is `0`.
2. **Recursive Decision**:
   - If the current item’s weight exceeds the current capacity, it cannot be
     included.
   - Otherwise, decide whether to include or exclude the item.
   - Use the maximum value obtained by including or excluding the item.
3. **Helper Function**:
   - The `max()` function helps determine which choice yields a higher value.

### Time Complexity

- The **time complexity** of the recursive approach is `O(2^n)`, where `n` is
  the number of items.
- This is because, for each item, the function makes two recursive calls,
  resulting in exponential growth.
- **Space Complexity**: `O(n)` due to the recursion stack.

### Optimizing Recursion with Memoization

The recursive solution for the Knapsack problem has **overlapping subproblems**,
meaning that many subproblems are solved multiple times. We can improve the
efficiency by using **Memoization** (storing previously computed results):

#### Memoized Version

To optimize the recursive solution, we use a **2D array** to store results for
subproblems:

```go
package main

import "fmt"

// knapsackMemo solves the Knapsack problem using memoization (top-down dynamic programming)
func knapsackMemo(values []int, weights []int, capacity int, n int, memo [][]int) int {
    // Base case: no items left or capacity is 0
    if n == 0 || capacity == 0 {
        return 0
    }

    // Check if result already exists in memo
    if memo[n][capacity] != -1 {
        return memo[n][capacity]
    }

    // If the weight of the nth item is more than the capacity, it cannot be included
    if weights[n-1] > capacity {
        memo[n][capacity] = knapsackMemo(values, weights, capacity, n-1, memo)
    } else {
        // Maximize the value by including or excluding the nth item
        includeItem := values[n-1] + knapsackMemo(values, weights, capacity-weights[n-1], n-1, memo)
        excludeItem := knapsackMemo(values, weights, capacity, n-1, memo)
        memo[n][capacity] = max(includeItem, excludeItem)
    }

    return memo[n][capacity]
}

// Helper function to initialize the memo table
func initializeMemo(n int, capacity int) [][]int {
    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, capacity+1)
        for j := range memo[i] {
            memo[i][j] = -1 // Initialize with -1 (indicating not computed)
        }
    }
    return memo
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    values := []int{60, 100, 120}
    weights := []int{10, 20, 30}
    capacity := 50
    n := len(values)

    // Initialize memo table
    memo := initializeMemo(n, capacity)

    result := knapsackMemo(values, weights, capacity, n, memo)
    fmt.Printf("Maximum value in Knapsack = %d\n", result)
}
```

### Optimized Complexity

- **Time Complexity**: `O(n * W)`, where `n` is the number of items and `W` is
  the capacity. Memoization ensures that each subproblem is only solved once.
- **Space Complexity**: `O(n * W)` for storing the memo table.

### Summary

- The **Divide and Conquer** approach using **recursion** is intuitive but can
  be computationally expensive due to redundant calculations.
- Using **memoization** significantly improves the performance by storing
  results of subproblems, thus reducing the time complexity.
- For practical use with large inputs, the **Dynamic Programming** approach
  (bottom-up) or **memoization** approach (top-down) is recommended over plain
  recursion.

If you have more questions about the Knapsack problem or need additional
explanation, feel free to ask!
