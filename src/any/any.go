package main

import "fmt"

func f(name string, data interface{}) { // 等价于 func f(name string, data any) {
	switch data.(type) {
	case string:
		fmt.Println(data.(string))
	case int:
		fmt.Println(data.(int))
	default:
		fmt.Println("do nothing...")
	}
}

func main() {
	f("tom", "1")
	f("Tom", 2)
}
