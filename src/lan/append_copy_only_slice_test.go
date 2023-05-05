package lan

import "testing"

// [0, 2)
func TestAppend(t *testing.T) {
	names := make([]string, 0, 10)
	names = append(names, "Tom")
	names = append(names, "Jimmy")

	_ = append([]byte("hello"), "names"...)
	_ = append(names, names...) // 区别于 names = append(names, names...)

	t.Log(names)
}

func TestCopy(t *testing.T) {
	names := make([]string, 0, 10)
	names = append(names, "Tom")
	names = append(names, "Jimmy")

	var copes []string // nil  不能copy给 nil 切片

	count := copy(copes, names)
	t.Log(count)

	var cs []string = []string{"Hello"} //  目的地切片len不够，则只能copy len长度部分，并且会覆盖原值
	ct := copy(cs, names)
	t.Log(ct)

	lcs := make([]string, 4, 100)
	lct := copy(lcs, names)
	t.Log(lct)
}
