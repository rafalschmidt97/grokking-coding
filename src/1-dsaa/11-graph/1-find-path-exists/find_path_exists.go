package main

import "fmt"

// Time complexity: O(V+E), combination (E), traversal (VE)
// Space complexity: O(V+E), representation (VE), visited (V), stack (V)
func findIfPathExists(nodes int, edges [][]int, start int, end int) bool {
	adjacencyList := make([][]int, nodes)

	// Map edges to adjacency list
	for i := 0; i < nodes; i++ {
		adjacencyList[i] = []int{}
	}
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0]) // Because it's an undirected graph
	}

	return DFSWithTerminationOnEnd(adjacencyList, start, end)
}

func DFSWithTerminationOnEnd(adjacencyList [][]int, start int, end int) bool {
	visited := make([]bool, len(adjacencyList))
	stack := []int{}

	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current == end {
			return true // found the path to end
		}

		for _, neighbor := range adjacencyList[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false
}

// func DFSWithTerminationOnEndRecursive(adjacencyList [][]int, visited []bool, node int, end int) bool {
// 	if node == end {
// 		return true // found the path to end
// 	}
// 	visited[node] = true
//
// 	for _, neighbor := range adjacencyList[node] {
// 		if !visited[neighbor] &&
// 			DFSWithTerminationOnEndRecursive(
// 				adjacencyList, visited, neighbor, end) {
// 			return true
// 		}
// 	}
//
// 	return false
// }

func main() {
	nodes := 4
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3},
	}
	start := 0
	end := 3
	fmt.Println(findIfPathExists(nodes, edges, start, end))
}
