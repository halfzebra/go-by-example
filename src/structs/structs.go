package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 16})
	fmt.Println(person{name: "Alice", age: 17})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Jane", age: 32})

	s := person{name: "Sean", age: 13}
	fmt.Println(s.name)
	sp := &s

	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age, s.age)
}
