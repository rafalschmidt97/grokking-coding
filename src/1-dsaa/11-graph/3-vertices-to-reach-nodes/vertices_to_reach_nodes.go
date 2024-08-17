package main

import "fmt"

// To solve the problem of determining the minimum number of vertices needed to reach
// all nodes in a directed graph, we focus on the concept of "in-degree" which
// represents the number of incoming edges to a node. In a directed graph,
// if a node doesn't have any incoming edges (in-degree of 0), then it means that
// the node cannot be reached from any other node. Hence, such nodes are mandatory
// starting points to ensure that every node in the graph can be reached.
// Our algorithm thus identifies all nodes with an in-degree of 0 as they are
// potential starting points to traverse the entire graph.
//
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
