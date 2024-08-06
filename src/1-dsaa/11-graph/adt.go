package main

import "fmt"

// For this course, we assume that graphs don't have any
// extra data or values associated with it. Therefore, we will add
// an empty vector or list as a value attribute for the new graph node or vertex.
type GraphMap struct {
	adjacencyList map[int][]int
}

type Pair struct {
	first  int
	second int
}

func NewGraph() *GraphMap {
	return &GraphMap{
		adjacencyList: make(map[int][]int),
	}
}

func (g *GraphMap) AddVertex(vertex int) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = make([]int, 0)
	}
}

// Remove a vertex from the graph.
func (g *GraphMap) RemoveVertex(vertex int) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		return
	}
	delete(g.adjacencyList, vertex)
	for _, neighbors := range g.adjacencyList {
		for i, v := range neighbors {
			if v == vertex {
				// Remove the vertex from the neighbors list
				neighbors = append(neighbors[:i], neighbors[i+1:]...)
			}
		}
	}
}

// Add an edge between two vertices.
func (g *GraphMap) AddEdge(vertex1, vertex2 int) {
	// If vertex1 is not in the adjacency list, create a new entry and add vertex2.
	if _, exists := g.adjacencyList[vertex1]; !exists {
		g.adjacencyList[vertex1] = []int{vertex2}
	} else {
		// Otherwise, append vertex2 to the existing list for vertex1.
		g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	}

	// If vertex2 is not in the adjacency list, create a new entry and add vertex1.
	if _, exists := g.adjacencyList[vertex2]; !exists {
		g.adjacencyList[vertex2] = []int{vertex1}
	} else {
		// Otherwise, append vertex1 to the existing list for vertex2.
		g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
	}
}

// Remove an edge between two vertices.
func (g *GraphMap) RemoveEdge(vertex1, vertex2 int) {
	for i, v := range g.adjacencyList[vertex1] {
		if v == vertex2 {
			// Remove vertex2 from the neighbors of vertex1
			g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1][:i], g.adjacencyList[vertex1][i+1:]...)
		}
	}
	for i, v := range g.adjacencyList[vertex2] {
		if v == vertex1 {
			// Remove vertex1 from the neighbors of vertex2
			g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2][:i], g.adjacencyList[vertex2][i+1:]...)
		}
	}
}

// Get a list of all vertices.
func (g *GraphMap) GetVertices() []int {
	vertices := make([]int, 0, len(g.adjacencyList))
	for vertex := range g.adjacencyList {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// Check if two vertices are adjacent.
func (g *GraphMap) IsAdjacent(vertex1, vertex2 int) bool {
	neighbors := g.adjacencyList[vertex1]
	for _, neighbor := range neighbors {
		if neighbor == vertex2 {
			return true
		}
	}
	return false
}

// Get the total number of vertices.
func (g *GraphMap) GetVertexCount() int {
	return len(g.adjacencyList)
}

// Get the total number of edges.
func (g *GraphMap) GetEdgeCount() int {
	count := 0
	for _, neighbors := range g.adjacencyList {
		count += len(neighbors)
	}
	return count / 2
}

// Get a list of all edges.
func (g *GraphMap) GetEdges() []Pair {
	edges := make([]Pair, 0)
	for vertex1, neighbors := range g.adjacencyList {
		for _, vertex2 := range neighbors {
			if vertex1 < vertex2 {
				edges = append(edges, Pair{vertex1, vertex2})
			}
		}
	}
	return edges
}

// Get a list of neighbors of a given vertex.
func (g *GraphMap) GetNeighbors(vertex int) []int {
	return g.adjacencyList[vertex]
}

func mainAdt() {
	graph := GraphMap{
		adjacencyList: make(map[int][]int),
	}
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)

	fmt.Print("Vertices: ")
	for _, vertex := range graph.GetVertices() {
		fmt.Printf("%d ", vertex)
	}
	fmt.Println()

	fmt.Print("Edges: ")
	for _, edge := range graph.GetEdges() {
		fmt.Printf("(%d, %d) ", edge.first, edge.second)
	}
	fmt.Println()

	fmt.Print("Neighbors of vertex 1: ")
	for _, neighbor := range graph.GetNeighbors(1) {
		fmt.Printf("%d ", neighbor)
	}
	fmt.Println()

	fmt.Println("Is vertex 1 adjacent to vertex 2? ", graph.IsAdjacent(1, 2))

	graph.RemoveEdge(1, 2)
	graph.RemoveVertex(3)

	fmt.Println("After removing some edges and vertices:")

	fmt.Print("Vertices: ")
	for _, vertex := range graph.GetVertices() {
		fmt.Printf("%d ", vertex)
	}
	fmt.Println()

	fmt.Print("Edges: ")
	for _, edge := range graph.GetEdges() {
		fmt.Printf("(%d, %d) ", edge.first, edge.second)
	}
	fmt.Println()
}
