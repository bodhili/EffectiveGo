package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	defer close(done)

	fmt.Println(len(done))
	fmt.Println(cap(done))

	go worker(done)

	<-done // 阻塞到chan 写入
}
