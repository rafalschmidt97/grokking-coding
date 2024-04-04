package main

import "fmt"

func template(x int) int {
	return x
}

func main() {
	x := 1
	fmt.Println(template(x))
}
