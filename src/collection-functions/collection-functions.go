package main

import (
	"fmt"
	"strings"
)

func Index(haystack []string, needle string) int {
	for index, char := range haystack {
		if char == needle {
			return index
		}
	}

	return -1
}

func Includes(haystack []string, needle string) bool {
	return Index(haystack, needle) != -1
}

func Any(haystack []string, predicate func(string) bool) bool {
	for _, value := range haystack {
		if predicate(value) {
			return true
		}
	}
	return false
}

func All(haystack []string, predicate func(string) bool) bool {
	for _, value := range haystack {
		if predicate(value) {
			return false
		}
	}
	return true
}

func Filter(array []string, predicate func(string) bool) []string {
	result := make([]string, 0)
	for _, value := range array {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func Map(array []string, fn func(string) string) []string {
	result := make([]string, len(array))
	for index, value := range array {
		result[index] = fn(value)
	}
	return result
}

func main() {
	fmt.Println("Collection functions")
	fruits := []string{"apple", "orange", "pear"}
	fmt.Println("Index of pear", Index(fruits, "pear"))

	fmt.Println("Includes potato", Includes(fruits, "potato"))

	fmt.Println("Any is longer than 6", Any(fruits, func(value string) bool {
		return len(value) >= 6
	}))

	fmt.Println("All are not starting with b", All(fruits, func(value string) bool {
		return strings.HasPrefix(value, "b")
	}))

	fmt.Println("Filter", Filter(fruits, func(value string) bool {
		return strings.HasSuffix(value, "e")
	}))

	fmt.Println("Map to uppercase", Map(fruits, func(value string) string {
		return strings.ToUpper(value)
	}))
}
