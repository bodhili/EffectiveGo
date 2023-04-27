package lan

import "testing"

func TestSlice(t *testing.T) {
	months := [...]string{1: "January", 2: "December"} // 下标：0 自动被初始化为 zero 值

	t.Log(len(months)) // 3
	for i, month := range months {
		t.Logf("i:[%d] v:[%v]", i, month)
	}
}

func TestReverse(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])

	t.Log(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
