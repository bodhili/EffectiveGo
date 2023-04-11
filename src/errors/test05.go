package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered: ", r)
	}
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func sum(a int, b int) {
	defer recovery()

	fmt.Printf("%d + %d = %d\n", a, b, a+b)

	done := make(chan bool)

	go divide(a, b, done)
	<-done
}

func main() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}
