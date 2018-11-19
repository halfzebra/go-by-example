package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return r.width*2 + r.height*2
}

func main() {
	a := rect{4, 8}
	fmt.Println("a area:", a.area())
	fmt.Println("a perim:", a.perimeter())

	// pointer
	b := &a
	fmt.Println("b area:", b.area())
	fmt.Println("b perim:", b.perimeter())
}
