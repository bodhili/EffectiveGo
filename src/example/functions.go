package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	whatAmI := func(a, b int) int {
		return a << b
	}
	fmt.Println(whatAmI(1, 2))

	multiple := func(a, b int) (int, int) {
		return b, a
	}

	i, _ := multiple(2, 3)
	fmt.Println(i)

	fmt.Println(multiple(1, 2))
}
