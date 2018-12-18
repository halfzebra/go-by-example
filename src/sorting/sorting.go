package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 5}
	sort.Ints(ints)
	fmt.Println("Intes: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
