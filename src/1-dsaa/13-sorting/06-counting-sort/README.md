# Counting sort

> An integer sorting algorithm that operates with key assumption about the range
> of the input data.

Counting sort calculates the number of occurrences of each distinct element in
the array to sort the array. It then uses arithmetic to determine the positions
of each element in the output sequence. This sort is efficient when the range of
input data is not significantly greater than the number of objects to be sorted.

## Step-by-Step Algorithm

1. Determine Range:
   1. Identify the Maximum Value: Find the largest value in the input array to
      determine the range of possible values.
   1. Create Count Array: Establish an array, `count`, with a size of `max + 1`
      (to account for zero indexing) initialized to zero.
1. Count Occurrences:
   1. Traverse Input Array: For each element `x` in the input array, increment
      `count[x]`.
   1. Purpose: This step tallies the number of times each value appears in the
      input array.
1. Accumulate Counts:
   1. Modify Count Array: Transform each index in the `count` array to be the
      sum of the previous 1.indices.
   1. Key Operation: `count[i] += count[i - 1]`.
   1. Outcome: After this step, `count[i]` tells the number of elements less
      than or equal to i.
1. Build the Output Array:
   1. Initialize Output Array: Create an array `output` that will store the
      sorted elements.
   1. Place Elements Correctly: Iterate through the input array from the last
      element to the first (to maintain stability):
      1. Place the element in the correct position in the `output` array:
         `output[count[arr[i]] - 1] = arr[i]`.
      1. Decrement the count in the `count` array: `count[arr[i]]--`.
      1. Highlight: This placement ensures elements are sorted and that the sort
         is stable (maintains the relative order of duplicate values).
1. Copy to Original Array:
   1. Final Step: Copy the sorted elements from the `output` array back to the
      original input array.
