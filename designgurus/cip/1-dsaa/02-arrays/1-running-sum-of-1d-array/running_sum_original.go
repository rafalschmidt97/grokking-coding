package main

// steps:
// 1. Initialize a Variable for Running Sum: Start by setting a variable (let's call it runningSum) to 0. This variable will keep track of the cumulative sum as you iterate through the array.
// 2. Iterate Through the Array: Loop through each element of the array. You can use a standard for loop for this purpose.
// 3. Update Running Sum: In each iteration, add the current element's value to runningSum. This step updates the running total with the value of the current element.
// 4. Replace Current Element: After updating the runningSum, replace the current element in the array with the value of runningSum. This effectively changes each element to the sum of all preceding elements including itself.
// 5. Continue Until the End of the Array: Continue steps 3 and 4 for each element in the array. By the end of the loop, each array element will have been replaced with the cumulative sum up to that point.
// 6. Return the Modified Array: After completing the iteration through the array, return the modified array. This array now contains the running sums instead of the original values.
//
// Time Complexity:
// The algorithm essentially involves a single loop that iterates through the input array to compute the running sum. Here's the breakdown:
// O(n): We iterate through all (n) elements of the input array exactly once. Inside the loop, we perform a constant amount of work (a sum and an assignment). Hence, the time complexity is linear.
// Space Complexity:
// O(1): The algorithm uses a constant amount of extra space. The result array does not count towards the space complexity since it's the expected output. However, no additional data structures grow with the input size, meaning that the algorithm uses a constant amount of additional memory.
func runningSumOriginal(nums []int) []int {
	// Check if the array is nil or has no elements and return an empty array if true
	if len(nums) == 0 {
		return []int{}
	}

	// Initialize an array to store the running sum
	result := make([]int, len(nums))
	result[0] = nums[0]

	// Loop through the array starting from index 1, adding the previous sum to the current element
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}

	// Return the array containing the running sum
	return result
}
