package main

import (
	"fmt"
	"runtime/debug"
)

func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

//func main() {
//	invalidSliceAccess()
//	fmt.Println("normally returned from main")
//}
