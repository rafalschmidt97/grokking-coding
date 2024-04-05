package main

import "fmt"

// The sequence follows the rule that each number is equal to the sum
// of the preceding two numbers. The Fibonacci sequence begins with
// the following 14 integers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233
//
// Time complexity: O(2^n), because each call to fibonacci(n) results
// in two additional recursive calls to fibonacci(n-1) and fibonacci(n-2).
//
// Space complexity: O(n), because of height of the call stack (tree) and
// only descend one of the recursions at a time.
func fibonacciSequence(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return fibonacciSequence(n-1) + fibonacciSequence(n-2)
}

func main() {
	n := 10
	fmt.Println(fibonacciSequence(n))
}
