package main

import "fmt"

// It is impossible to use "const" since array in Go is mutable by it's nature

func main() {
	var a [5]int
	var b [5]string
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("length of b:", len(b))

	c := [2]int{1, 2}
	fmt.Println(c)

	var matrix [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println(matrix)

	var matrix2 = [2][2]int{{14, 15}, {17, 18}}
	fmt.Println(matrix2)
}
