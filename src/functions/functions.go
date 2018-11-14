package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addTree(a, b, c int) int {
	return a + b + c
}

func main() {
	res := add(1, 3)
	fmt.Println("res:", res)

	res = addTree(1, 3, 4)
	fmt.Println("1 + 3 + 4 =", res)
}
