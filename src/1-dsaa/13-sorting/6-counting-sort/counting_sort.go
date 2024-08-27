package main

import "fmt"

func countingSort(array []int) []int {
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := countingSort(inputArray)
	fmt.Println(sortedArray)
}
