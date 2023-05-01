package leetcode

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	i, j := 1, 2
	j, i = i, j

	fmt.Println(i)
	fmt.Println(j)

	swap(i, j)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Printf("%b \n", i)
	fmt.Printf("%b \n", j)

	fmt.Println("i | j", i|j)
	fmt.Println("i ^ j", i^j)

	nums := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(only0(nums))
	fmt.Println(only1(nums))

	fmt.Printf("%b \n", 8)
}

func swap(i, j int) {
	i = i ^ j
	j = i ^ j
	i = i ^ j
}

func only0(nums []int) int {
	var ret int
	for _, v := range nums {
		ret |= v
	}

	return ret
}

func only1(nums []int) int {
	var ret int
	for _, v := range nums {
		ret ^= v
	}

	return ret
}
