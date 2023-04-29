package leetcode

import "testing"

func TestPartOne(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	sum := twoSum0(nums, 9)

	t.Log(sum)
}

func twoSum0(nums []int, target int) []int {
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		t, ok := temp[v]
		if ok {
			return []int{t, i}
		}

		temp[nums[i]] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, v := range nums {
		if t, ok := temp[target-nums[i]]; ok {
			return []int{i, t}
		}
		temp[v] = i
	}
	return nil
}
