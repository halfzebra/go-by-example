package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("sync")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous goroutine")

	fmt.Scanln()
	fmt.Println("done")
}
