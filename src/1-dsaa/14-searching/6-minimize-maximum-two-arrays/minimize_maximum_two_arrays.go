package main

import "fmt"

func minimizeMaximum(divisor1, divisor2, uniqueCnt1, uniqueCnt2 int) int {
	counter1, counter2 := 0, 0
	largest1, largest2 := 0, 0

	for i := 1; counter1 < uniqueCnt1 || counter2 < uniqueCnt2; i++ {
		if i%divisor1 != 0 && counter1 < uniqueCnt1 {
			largest1 = i
			counter1++
			continue
		}

		if i%divisor2 != 0 && counter2 < uniqueCnt2 {
			largest2 = i
			counter2++
			continue
		}
	}

	if largest1 > largest2 {
		return largest1
	} else {
		return largest2
	}
}

func main() {
	uniqueCnt1 := 2
	divisor1 := 2
	uniqueCnt2 := 2
	divisor2 := 3
	fmt.Println(minimizeMaximum(divisor1, divisor2, uniqueCnt1, uniqueCnt2))
}
