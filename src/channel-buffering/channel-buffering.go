package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	fmt.Println(messages)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
