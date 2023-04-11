package main

import (
	"os"
)

var user = os.Getenv("tom")

//func main() {
//	defer func() {
//		fmt.Println("ending...")
//	}()
//
//	if user == "" {
//		panic("No value for $USER")
//	}
//
//	fmt.Println(user)
//}
