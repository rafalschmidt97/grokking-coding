package main

import "fmt"

// Time complexity: O(n^2)
// Space complexity: O(n)
func numberOfProvinces(adjacencyList [][]int) int {
	provinces := 0
	visited := make([]bool, len(adjacencyList))

	for i := 0; i < len(adjacencyList); i++ {
		if !visited[i] {
			dfs(adjacencyList, visited, i)
			provinces++
		}
	}

	return provinces
}

func dfs(adjacencyList [][]int, visited []bool, i int) {
	for j := 0; j < len(adjacencyList); j++ {
		if adjacencyList[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(adjacencyList, visited, j)
		}
	}
}

func main() {
	input := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(numberOfProvinces(input))
}

// TODO: Not working.
//
// func numberOfProvincesV1(adjacencyList [][]int) int {
// 	provinces := 0
// 	visited := make([]bool, len(adjacencyList))
//
// 	for start := 0; start < len(adjacencyList); start++ {
// 		if !visited[start] {
// 			stack := []int{}
// 			stack = append(stack, start)
// 			visited[start] = true
//
// 			for len(stack) > 0 {
// 				current := stack[len(stack)-1]
// 				stack = stack[:len(stack)-1]
//
// 				for _, neighbor := range adjacencyList[current] {
// 					if !visited[neighbor] && adjacencyList[current][neighbor] == 1 {
// 						stack = append(stack, neighbor)
// 						visited[neighbor] = true
// 					}
// 				}
// 			}
//
// 			provinces++
// 		}
// 	}
//
// 	return provinces
// }

// func dfsV1(adjacencyList [][]int, visited []bool, node int) {
// 	visited[node] = true
//
// 	for _, neighbor := range adjacencyList[node] {
// 		if !visited[neighbor] {
// 			dfs(adjacencyList, visited, neighbor)
// 		}
// 	}
// }
