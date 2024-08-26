package main

type Solution struct{}

func (s *Solution) Merge(arr []int, l int, m int, r int) {
	/**
	 * Merges two subarrays of arr[].
	 * First subarray is arr[l..m]
	 * Second subarray is arr[m+1..r]
	 */
	// Calculate the sizes of two subarrays to be merged
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays
	copy(L, arr[l:m+1])
	copy(R, arr[m+1:r+1])

	// Merge the temporary arrays back into arr[l..r]
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[], if any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func (s *Solution) Sort(arr []int, l int, r int) {
	/**
	 * Main function that sorts arr[l..r] using merge()
	 */
	if l < r {
		// Find the middle point to divide the array into two halves
		m := l + (r-l)/2

		// Recursively sort the first and second halves
		s.Sort(arr, l, m)
		s.Sort(arr, m+1, r)

		// Merge the sorted halves
		s.Merge(arr, l, m, r)
	}
}
