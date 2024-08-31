package main

import (
	"fmt"
	"math"
	"sort"
)

// Time Complexity: The average case is O(n + n^2 / k + k), where k is the
// number of buckets. If k approaches n, the time complexity nears .
//
// Space Complexity: O(n+k), due to the space needed for the buckets and
// the elements they contain.
func bucketSort(arr []int) []int {
	n := len(arr)

	// Determine the number of buckets, typically equal to the square root of the number of elements
	k := int(math.Sqrt(float64(n)))
	buckets := make([][]int, k)

	// Find the maximum value in the array to scale the bucket index
	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	// Add elements into the appropriate buckets
	for _, value := range arr {
		// Calculate bucket index
		bucketIndex := int(math.Floor(float64(value) * float64(k) / float64(maxValue+1)))
		buckets[bucketIndex] = append(buckets[bucketIndex], value)
	}

	// Sort each bucket
	for i := 0; i < k; i++ {
		sort.Slice(buckets[i], func(j, l int) bool {
			return buckets[i][j] < buckets[i][l]
		})
	}

	// Concatenate all buckets into arr
	index := 0
	for _, bucket := range buckets {
		for _, elem := range bucket {
			arr[index] = elem
			index++
		}
	}

	return arr
}

func main() {
	inputArray := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(inputArray)
	sortedArray := bucketSort(inputArray)
	fmt.Println(sortedArray)
}
