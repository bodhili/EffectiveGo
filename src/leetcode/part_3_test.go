package leetcode

import (
	"testing"
)

func TestSwap(t *testing.T) {
	i, j := 1, 2
	j, i = i, j

	t.Log(i)
	t.Log(j)

	swap(i, j)
	t.Log(i)
	t.Log(j)

	t.Logf("%b \n", i)
	t.Logf("%b \n", j)

	t.Log("i | j", i|j)
	t.Log("i ^ j", i^j)

	nums := []int{1, 2, 3, 4, 3, 2, 1}
	t.Log(only0(nums))
	t.Log(only1(nums))

	t.Logf("%b \n", 8)
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
