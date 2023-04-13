package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var x float64 = 3.4
	r := reflect.ValueOf(x)

	fmt.Println(r.Type())
	fmt.Println(r.Float())

	v := reflect.ValueOf(&x).Elem()
	v.SetFloat(7.1)
	fmt.Println(v.Float())

	p := Person{
		"Alice", 90,
	}

	vv := reflect.ValueOf(p)
	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		fmt.Printf("Field %d: %s = %v\n", i, field.Type(), field.Interface())
	}

	f := reflect.ValueOf(add)
	args := []reflect.Value{reflect.ValueOf(3), reflect.ValueOf(4)}

	res := f.Call(args)[0].Int
	fmt.Println("result:", res)
}

func add(a, b int) int {
	return a + b
}
