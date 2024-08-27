package main

import "fmt"

func radixSort(array []int) []int {
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := radixSort(inputArray)
	fmt.Println(sortedArray)
}
