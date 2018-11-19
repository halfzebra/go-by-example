package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

func (c circle) area() float64 {
	return math.Phi * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Phi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	rectangeA := rectangle{width: 6, height: 5}
	circleB := circle{radius: 16}

	measure(rectangeA)
	measure(circleB)
}
