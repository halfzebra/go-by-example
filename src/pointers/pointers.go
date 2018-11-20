package main

import (
	"fmt"
)

func asssignZero(i int) {
	i = 0
}

func asssignZeroByPointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	asssignZero(i)
	fmt.Println("asssignZero(i = 1): ", i)
	fmt.Println("pointer to", &i)
}
