package main

import "fmt"

// returns index if exits, otherwise -1
// time complexity: O(logn)
//
// assumes that the arr is sorted = [2, 3, 4, 5, 10, 40]; l=0; h=5
// formula for middle = low + (high - low)/2
//
// for x=10
//  1. m: 0 + (5-0)/2 = 2; l=3
//  1. m: 3 + (5-3)/2 = 4; done
//
// for x=2
//  1. m: 0 + (5-0)/2 = 2; h=1
//  1. m: 0 + (1-0)/2 = 0; done
//
// for x=3
//  1. m: 0 + (5-0)/2 = 2; h=1
//  1. m: 0 + (1-0)/2 = 0; l=1
//  1. m: 1 + (1-1)/2 = 1; done
func binarySearch(array []int, x int) (index int) {
	lowIndex := 0
	highIndex := len(array) - 1 // -1 for empty list

	for lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2

		if array[midIndex] < x {
			lowIndex = midIndex + 1
			fmt.Printf("m: %v (l: %v, h: %v)\n", midIndex, lowIndex, highIndex)
		} else if array[midIndex] > x {
			highIndex = midIndex - 1
			fmt.Printf("m: %v (l: %v, h: %v)\n", midIndex, lowIndex, highIndex)
		} else { // equal
			fmt.Printf("m: %v (l: %v, h: %v)\n", midIndex, lowIndex, highIndex)
			return midIndex
		}
	}

	return -1
}

func main() {
	array := []int{2, 3, 4, 5, 10, 40}
	x := 3
	fmt.Println(binarySearch(array, x))
}
