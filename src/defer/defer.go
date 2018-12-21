package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer fmt.Println("Done writing to file!")
	fmt.Println("Writing to file...")
	fmt.Fprintln(f, "Hello from Go!")
}
