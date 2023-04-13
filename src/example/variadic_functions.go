package main

import "fmt"

func sun(nums ...int) { // var []int  slice定义
	fmt.Println(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sun(1)
	sun(1, 2)
	sun()

	nums := []int{1, 2, 3}
	sun(nums...)
}
