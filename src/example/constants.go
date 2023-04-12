package main

import (
	"fmt"
	"math"
)

const name string = "tom" // const name = "tom"

func main() {
	fmt.Println(name)

	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	//	n = 5

}
