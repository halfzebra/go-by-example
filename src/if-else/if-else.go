package main

import "fmt"

func main() {
	if true {
		fmt.Println("hello")
	}

	if num := 0; num%2 == 0 {
		fmt.Println(num, "is divisible by 2")
	} else if num%2 != 0 {
		fmt.Println(num, "is not divisible by 2")
	} else {
		fmt.Println("impossible")
	}
}
