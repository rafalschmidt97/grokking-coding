package main

// Problem: single removal logic.
func validPalindromeByRemovingLeft(input string) bool {
	fi := 0
	li := len(input) - 1
	removed := 0

	for fi < li {
		if input[fi] != input[li] {
			fi++
			removed++
		}

		if removed > 1 {
			return false
		}

		fi++
		li--
	}

	return true
}
