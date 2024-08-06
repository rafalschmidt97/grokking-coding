package main

import (
	"fmt"
)

// TODO: why I cannot share the graph structure in the same package
// and getting redeclared in this block and undefined graph map in the same time.

// Time complexity: In the worst case, DFS once visits all nodes and edges in
// the graph. For a graph with V vertices (nodes) and E edges, the time
// complexity of DFS is O(V+E)
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
func (g *GraphMapDFS) DFS(start int) {
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
	graph := GraphMapDFS{
		adjacencyList: make(map[int][]int),
	}
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	fmt.Print("DFS Traversal starting from vertex 0: ")
	graph.DFS(0)
}

// Boilerplate

type GraphMapDFS struct {
	adjacencyList map[int][]int
}

func (g *GraphMapDFS) AddEdge(vertex1, vertex2 int) {
	if _, exists := g.adjacencyList[vertex1]; !exists {
		g.adjacencyList[vertex1] = []int{vertex2}
	} else {
		g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	}

	if _, exists := g.adjacencyList[vertex2]; !exists {
		g.adjacencyList[vertex2] = []int{vertex1}
	} else {
		g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
	}
}
