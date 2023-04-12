package main

import "fmt"

type T struct{}

var x interface{}
var v *T
var a any

// byte        alias for uint8
// rune        alias for int32

func main() {
	ints := new([100]int)[0:50]

	fmt.Println(ints)

	i, err := x.(int)
	if !err {
		fmt.Println(err)
	}

	fmt.Println(i)
	fmt.Println(v)
	fmt.Println(x)
	fmt.Println(a)
}
