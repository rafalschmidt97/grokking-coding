# Bucket sort

> Also known as bin sort, it distributes elements into various 'buckets' which
> are then sorted using another sort, typically insertion sort.

Bucket sort, also known as bin sort, is an efficient sorting algorithm that
distributes elements of an array into a number of buckets. Each bucket is then
sorted individually, either using a different sorting algorithm or by
recursively applying the bucket sort. This method is useful when the input is
uniformly distributed over a range.

## Step-by-Step Algorithm

1. Create Buckets:
   1. Determine the number of buckets, `k`, typically equal to the square root
      of the number of elements in the array.
   1. Create an array of buckets where each bucket is a list that will hold
      elements of the array.
1. Distribute Elements:
   1. For each element in the array, find the appropriate bucket based on its
      value. This can be calculated using a function like
      `index = floor(value * k / (maxValue + 1))`.
   1. Insert the element into its corresponding bucket.
1. Sort Each Bucket:
   1. Sort the elements in each bucket. This can be done using a simple sorting
      algorithm like insertion sort for efficiency because each bucket is
      expected to be small.
1. Concatenate Buckets:
   1. Once all buckets are sorted, concatenate them back into the original array
      in order. This effectively compiles the sorted elements into a single
      sorted array.
