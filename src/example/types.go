package main

import "fmt"

type T struct {
	Name string
}

// byte        alias for uint8
// rune        alias for int32 (char 'a')

// int8 8代表的是bit位,8个bit位，相当于一个字节
// int 自定匹配操作系统位数
// uintptr （存储指针类型，字节大小和操作系统有关，64位: 8字节 ）  an unsigned integer large enough to store the uninterpreted bits of a pointer value
// <nil> 是指向底层数据结构的一个指针

func main() {

	var x interface{}
	var x1 any
	var i int
	var s struct{}
	fmt.Println(s)

	fmt.Println(i)

	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", x1)

	fmt.Println(x)
	fmt.Println(x1)
	fmt.Println(x1 == x)
	fmt.Println(x1 == nil)

	var t *T
	fmt.Printf("%T \n", x1)

	fmt.Println(t)
	fmt.Println(x1 == nil)

	x = t
	fmt.Println(x)
}
