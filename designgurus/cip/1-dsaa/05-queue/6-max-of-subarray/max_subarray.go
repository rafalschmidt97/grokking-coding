package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func maxOfSubarray(input []int, k int) []int {
	queue := []int{}

	iv := [2]int{0, input[0]} // index and value tuple
	for i, v := range input {
		if v >= iv[1] || i > iv[0]+k {
			iv = [2]int{i, v}
		}

		if i+1 >= k {
			queue = append(queue, iv[1])
		}
	}

	return queue
}

func main() {
	input := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	k := 3
	fmt.Println(maxOfSubarray(input, k))
}
