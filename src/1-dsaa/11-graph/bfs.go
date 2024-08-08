package main

import (
	"fmt"
)

// Time Complexity:
//
// Visiting a vertex takes O(1) time as we dequeue it from the queue in constant time.
// Exploring the neighbors of a vertex takes O(1) time per neighbor, as we have to
// traverse its adjacency list once.
// In the worst case, we visit all the vertices at least once, which takes O(V) time.
// Additionally, for each vertex, we explore all its neighbors once, which takes O(E)
// time in total (sum of the sizes of all adjacency lists).
// Hence, the overall time complexity of BFS is: O(V + E)
//
// Space Complexity:
//
// The space required to store the graph using an adjacency list representation
// is O(V + E), as we need to store each vertex and its corresponding edges.
// The space required for the queue in BFS is O(V) in the worst case,
// as all the vertices can be in the queue at once.
// Since the space occupied by the queue is dominant in the overall
// space complexity, the space complexity of BFS is: O(V)
func (g *GraphMapBFS) BFS(start int) {
	visited := make([]bool, len(g.adjacencyList))
	queue := []int{}

	// 1. Initialization: choose a starting node and push it to the queue
	queue = append(queue, start)
	// 2. Visit the current node
	visited[start] = true

	for len(queue) > 0 { // 6. Termination
		current := queue[0]     // 3. Peak a node
		queue = queue[1:]       // 3. Pop a node
		fmt.Print(current, " ") // 4.1. Process the node

		for _, neighbor := range g.adjacencyList[current] { // Process unvisited neighbor
			if !visited[neighbor] { // 5. Backstrack until unvisited found
				queue = append(queue, neighbor) // push
				visited[neighbor] = true        // mark as visited
			}
		}
	}
}

func mainBFS() {
	graph := GraphMapBFS{
		adjacencyList: make(map[int][]int),
	}
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	fmt.Print("BFS Traversal starting from vertex 0: ")
	graph.BFS(0)
}

// Boilerplate

type GraphMapBFS struct {
	adjacencyList map[int][]int
}

func (g *GraphMapBFS) AddEdge(vertex1, vertex2 int) {
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
