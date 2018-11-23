package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("got message", msg)
	default:
		fmt.Println("no messages")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent", msg)
	default:
		fmt.Println("message was not sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)

	case msg := <-signals:
		fmt.Println("received signal", msg)
	default:
		fmt.Println("no activity")
	}
}
