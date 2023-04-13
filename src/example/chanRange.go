package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "a"
	queue <- "b"

	close(queue)

	for e := range queue {
		fmt.Println(e)
	}
}
