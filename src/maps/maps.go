package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 12
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)
	_, prs := m["k2"]
	fmt.Println(prs)

	// shorthand
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}
