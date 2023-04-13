package main

import "fmt"

func main() {
	messages1 := make(chan string, 2)

	defer close(messages1)

	messages1 <- "buffered"
	messages1 <- "channel"

	fmt.Println(<-messages1)
	fmt.Println(<-messages1)

}
