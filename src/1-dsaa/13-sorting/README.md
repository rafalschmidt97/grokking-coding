# Sorting

## Terminologies in Sorting

- Stability: A sorting algorithm is stable if it preserves the relative order of
  equivalent elements from the original list.
- In-place Sorting: An in-place sorting algorithm rearranges the elements within
  the array, using only a small, constant amount of additional storage space.
- Comparison Sort: A type of sort that determines the sorted order based on
  comparisons between the input elements.
- Non-comparison Sort: Sorting techniques that do not rely on element comparison
  and typically work with integer keys by grouping and counting elements.

## Not covered

- Heap sort - comparison-based sorting algorithm that uses a binary heap data
  structure to efficiently sort elements. It is an in-place algorithm with O(n
  log n) time complexity and works by repeatedly extracting the maximum (or
  minimum) element from a heap and placing it at the end of the array.

## Fun fact about GoLang

The sort.Slice function, uses a variant of the `Quicksort` algorithm called
`pdqsort` (Pattern-Defeating Quicksort) for sorting slices. It is a hybrid
sorting algorithm that combines the following techniques:

- Quicksort: The core algorithm is based on Quicksort, which is a
  divide-and-conquer algorithm that works by selecting a pivot and partitioning
  the array into two subarrays, then recursively sorting the subarrays.

- Heap Sort: In cases where Quicksort would degenerate into a worst-case O(n²)
  performance (e.g., when the input is already sorted or nearly sorted), pdqsort
  switches to Heap Sort, which has guaranteed O(n log n) performance.

- Insertion Sort: For small arrays, pdqsort uses Insertion Sort, which is
  efficient for small inputs due to its low overhead.
