package main

import "container/list"

func maxOfSubarrayOriginal(arr []int, k int) []int {
	queue := list.New()
	result := []int{}

	for i := 0; i < len(arr); i++ {
		// Remove elements which are out of the current window
		for queue.Len() > 0 && queue.Front().Value.(int) < i-k+1 {
			queue.Remove(queue.Front())
		}

		// Remove all elements smaller than the currently being added element
		// (remove useless elements)
		for queue.Len() > 0 && arr[i] >= arr[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		// Add the current element's index at the rear of the deque
		queue.PushBack(i)

		if i >= k-1 {
			// When the window size reaches k, add the maximum element to the result
			result = append(result, arr[queue.Front().Value.(int)])
		}
	}

	return result
}
