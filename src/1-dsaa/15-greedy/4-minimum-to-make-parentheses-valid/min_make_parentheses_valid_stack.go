package main

// Time complexity: O(n)
// Space complexity: O(1)
func minMakeParenthesesValid(input string) int {
	balance, counter := 0, 0
	for _, c := range input {
		if c == '(' {
			balance++ // Increment for each opening parenthesis
		} else {
			if balance > 0 {
				balance-- // Match closing parenthesis with an opening one
			} else {
				counter++ // Increment counter for an unmatched closing parenthesis
			}
		}
	}
	// The sum of counter and balance gives the total number of parentheses needed.
	return counter + balance
}
