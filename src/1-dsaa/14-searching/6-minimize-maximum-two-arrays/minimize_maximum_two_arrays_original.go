package main

// Time complexity: O(log(low-high))
// Space complexity: O(1)
func minimizeMaximumOriginal(
	divisor1, divisor2, uniqueCnt1, uniqueCnt2 int,
) int {
	low, high := uniqueCnt1+uniqueCnt2,
		divisor1*divisor2*uniqueCnt1*uniqueCnt2 // Setting initial search bounds

	// how many numbers are divisible
	lcm := (divisor1 * divisor2) / gcd(divisor1, divisor2) // Calculate the least common multiple (LCM)

	for low <= high {
		mid := (low + high) / 2
		countBoth := mid / lcm // Numbers divisible by both divisor1 and divisor2
		if mid-countBoth >= uniqueCnt1+uniqueCnt2 && mid-(mid/divisor1) >= uniqueCnt1 && mid-(mid/divisor2) >= uniqueCnt2 {
			high = mid - 1 // Adjust high to find smaller max
		} else {
			low = mid + 1 // Adjust low to find feasible max
		}
	}
	return low // The minimum possible maximum value that satisfies the conditions
}

// greatest common divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
