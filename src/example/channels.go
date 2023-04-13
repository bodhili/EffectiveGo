package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	defer close(messages)

	fmt.Printf("%T \n", messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))

	go func() {
		messages <- "ping"

		fmt.Println(len(messages))
		fmt.Println(cap(messages))
	}()

	msg := <-messages // 阻塞等待, goroutine 直接通信，可使用chan
	fmt.Println(msg)
}
