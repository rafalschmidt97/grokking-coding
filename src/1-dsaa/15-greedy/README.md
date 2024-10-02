# Greedy Algorithms

> Used to solve problems in which the best choice is made at each step, and it
> finds a solution in a minimal step. This approach assumes that choosing a
> local optimum at each stage will lead to the determination of a global
> optimum.

The main goal of a greedy algorithm is to solve complex problems by breaking
them down into simpler subproblems, solving each one optimally to find a
solution that is as close as possible to the overall optimum.

> The essence of the greedy method here is making the most efficient pairing
> choice at each step, aiming for an overall optimal solution - the minimum
> number of boats to transport everyone.

## Characteristics of Greedy Method

- Local Optimization: At every step, the algorithm makes a choice that seems the
  best at that moment, aiming for local optimality.
- Irrevocability: Once a choice is made, it cannot be undone or revisited. This
  is a key characteristic that differentiates greedy methods from dynamic
  programming, where decisions can be re-evaluated.

## Components of Greedy Algorithm

- Candidate Set: The set of choices available at each step.
- Selection Function: This function helps in choosing the most promising
  candidate to be added to the current solution.
- Feasibility Function: It checks if a candidate can be used to contribute to a
  solution without violating problem constraints.
- Objective Function: This evaluates the value or quality of the solution at
  each step.
- Solution Function: It determines if a complete solution has been reached.

## Pros of Greedy Approach

- Efficiency in Time and Space: Greedy algorithms often have lower time and
  space complexities, making them fast and memory-efficient.
- Ease of Implementation: They are generally simpler and more straightforward to
  code than other complex algorithms.
- Optimal for Certain Problems: In problems with certain structures, like those
  with greedy-choice property and optimal substructure, greedy algorithms
  guarantee an optimal solution.
- Useful for Approximations: When an exact solution is not feasible, greedy
  algorithms can provide a close approximation quickly.

## Cons of Greedy Approach with Example

- Not Always Optimal: Greedy algorithms do not always yield the global optimum
  solution, especially in problems lacking a greedy-choice property.
- Shortsighted Approach: They make decisions based only on current information,
  without considering the overall problem.

## Standard Greedy Algorithms

- Kruskal’s Minimum Spanning Tree Algorithm: Builds a minimum spanning tree by
  adding the shortest edge that doesn’t form a cycle.
- Prim’s Minimum Spanning Tree Algorithm: Grows a minimum spanning tree from a
  starting vertex by adding the nearest vertex.
- Dijkstra’s Shortest Path Algorithm: Finds the shortest path from a single
  source to all other vertices in a weighted graph.
- Huffman Coding: Used for data compression, it builds a binary tree with
  minimum weighted path lengths for given characters.
- Fractional Knapsack Problem: Maximizes the total value of items in a knapsack
  without exceeding its capacity.
- Activity Selection Problem: Select the maximum number of activities that don't
  overlap in time.
- Greedy Best-First Search: Used in AI for pathfinding and graph traversal,
  prioritizing paths that seem closest to the goal.
