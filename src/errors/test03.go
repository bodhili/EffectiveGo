package main

import "fmt"

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstname *string, lastname *string) {
	defer recoverFullName()

	if firstname == nil {
		panic("runtime error: first name cannot be nil \n")
	}

	if lastname == nil {
		panic("runtime error: last name cannot be nil \n")
	}

	fmt.Printf("%s %s\n", *firstname, *lastname)
	fmt.Println("returned normally from fullName")
}

//func main() {
//	defer fmt.Println("deferred call in main")
//
//	firstname := "tom"
//
//	fullName(&firstname, nil)
//	fmt.Println("ending...")
//}
