package main

import (
	"fmt"
)

type GraphIndex struct {
	adjacencyList [][]int
}

func NewGraphDFS(size int) *GraphIndex {
	g := &GraphIndex{
		adjacencyList: make([][]int, 0),
	}
	for i := 0; i < size; i++ {
		g.adjacencyList = append(g.adjacencyList, []int{})
	}
	return g
}

func (g *GraphIndex) addEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u) // For an undirected graph
}

// Time complexity: In the worst case, DFS once visits all nodes and edges in
// the graph. For a graph with V vertices (nodes) and E edges, the time
// complexity of DFS is  O(V+E)
// - Visiting a node (marking it as visited and processing it) takes O(1) time.
// - Exploring all neighbors of a node takes O(d) time, where 'd' is the average
// degree of nodes in the graph. In the worst case, 'd' can be as
// high as V-1 (complete graph).
// - So, the time complexity can be approximated as O(V) for exploring
// all neighbors of one node.
//
// Space complexity: In the worst case, the maximum depth of the recursion
// stack (or the maximum number of nodes stored in the stack) is the height
// of the deepest branch of the graph. For a graph with a single connected component,
// this height can be O(V-1) (when all nodes are connected in a straight line).
// The space complexity of the recursion stack in the worst case is O(V).
// Additionally, if an explicit stack is used, its space complexity would also
// be O(V) in the worst case.
func (g *GraphIndex) DFS(start int) {
	visited := make([]bool, len(g.adjacencyList))
	stack := []int{}

	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(current, " ")

		for _, neighbor := range g.adjacencyList[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	graph := NewGraphDFS(7)
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(2, 6)

	fmt.Print("DFS Traversal starting from vertex 0: ")
	graph.DFS(0)
}
