package main

import "fmt"

// Time complexity: O(e+n)
// Space complexity: O(n)
func verticesToReach(nodes int, edges [][]int) []int {
	inDegree := make([]int, nodes)

	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	result := make([]int, 0)
	for i := 0; i < nodes; i++ {
		if inDegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
	}
	nodes := 6
	fmt.Println(verticesToReach(nodes, edges))
}
