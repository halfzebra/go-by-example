package main

import "fmt"

func lowerAndHigher(x int) (int, int) {
	return x - 1, x + 1
}

func main() {
	a, b := lowerAndHigher(3)
	fmt.Println(a, b)

	// can reassign
	_, b = lowerAndHigher(5)
	fmt.Println(b)
}
