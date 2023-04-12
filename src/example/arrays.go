package main

import "fmt"

// var 定义变量 和 := 定义的区别
// var 初始化为 zero-valued，:= 不初始化为zero-valued (简明声明，需要同时定义变量名称和变量初始化值)
func main() {
	a1 := [...]int{}
	fmt.Println(a1)      // []
	fmt.Println(len(a1)) // 0

	// 	a1[0] = 1  Invalid array index '0' (out of bounds for the 0-element array)

	var b [5]int
	fmt.Println("emp: ", b)

	b[4] = 100
	fmt.Println("set: ", b)
	fmt.Println("get: ", b[4])
	fmt.Println("len: ", len(b))
	fmt.Println("cap: ", cap(b))

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	var twoD [2][3]int // 2行 3列
	fmt.Println(twoD)

	for i := 0; i < 2; i++ { // 行
		for j := 0; j < 3; j++ { // 列
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
