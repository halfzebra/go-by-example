package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[0])
	fmt.Println("length:", len(s))

	fmt.Println("append:", append(s, "d"))
	fmt.Println("original slice didn't change:", s)

	s = append(s, "d", "e")
	fmt.Println("appended multiple values", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied to a new slice \"c\"", c)

	l := s[2:4]
	fmt.Println("sliced 2:5", l)

	l = s[1:]
	fmt.Println("sliced 2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	nestedSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		nestedSlice[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			nestedSlice[i][j] = i + j
		}
	}
	fmt.Println(nestedSlice)
}
