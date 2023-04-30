package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func TestStoreArraysIntByLong(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}

	numLongs := int(math.Ceil(float64(len(arr)) / 2.0))
	nums := make([]int64, numLongs)

	for i, v := range arr {
		index := i / 2
		offset := (i % 2) * 32
		storeLongValue(nums, index, offset, v)
	}
	fmt.Println(nums)

	for i := 0; i < len(arr); i++ {
		index := i / 2
		offset := (i % 2) * 32
		v := getLongValue(nums, index, offset)
		fmt.Println(v)
	}
}

func storeLongValue(nums []int64, index int, offset int, value int) {
	nums[index] |= int64(value) << offset
}

func getLongValue(nums []int64, index int, offset int) int {
	return int(nums[index] >> offset & 0xFFFFFFFF)
}
