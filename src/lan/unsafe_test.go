package lan

import (
	"testing"
	"unsafe"
)

type Empty struct{}

func TestUnsafeSizeOf(t *testing.T) {
	t.Log(unsafe.Sizeof(float64(0))) //8
	t.Log(unsafe.Sizeof(int16(0)))   //2

	arr := [...]int{1, 2, 3}
	t.Log(unsafe.Sizeof(arr)) //24

	arr1 := [...]int{}
	t.Log(unsafe.Sizeof(arr1))  // 0
	t.Log(unsafe.Sizeof(&arr1)) // 8

	// map size: 8
	m := make(map[struct{}]struct{})
	t.Log(unsafe.Sizeof(m))

	m1 := make(map[string]string)
	t.Log(unsafe.Sizeof(m1))

	m2 := make(map[string]int)
	t.Log(unsafe.Sizeof(m2))

	s := make([]int, 0, 10)
	s = append(s, 1)
	t.Log(unsafe.Sizeof(s))

	s1 := make([]string, 0, 10)
	t.Log(unsafe.Sizeof(s1))

	s2 := make([]string, 0, 5)
	t.Log(unsafe.Sizeof(s2))
	t.Log(unsafe.Sizeof(&s2))

	var s3 []int
	t.Log(unsafe.Sizeof(s3))

	t.Log(unsafe.Sizeof(Empty{}))
	t.Log(unsafe.Sizeof(&Empty{}))

	t.Log(unsafe.Sizeof(any(0))) // 16

	t.Log(Empty{} == Empty{})
	t.Log(&Empty{} == &Empty{})
}

var X struct {
	a bool  // 1
	b int16 // 2
	c []int // 24
}

func TestOff(t *testing.T) {
	t.Log(unsafe.Sizeof(X))

	t.Log(unsafe.Alignof(X))
	t.Log(unsafe.Alignof(X.a))
	t.Log(unsafe.Alignof(X.b))
	t.Log(unsafe.Alignof(X.c))

	t.Log(unsafe.Offsetof(X.a))
	t.Log(unsafe.Offsetof(X.b))
	t.Log(unsafe.Offsetof(X.c))
}

func TestUnsafePointer(t *testing.T) {
	ptr := 1
	pointer := unsafe.Pointer(&ptr)
	t.Log(pointer)
	t.Log(&ptr)

	t.Logf("%#016x\n", Float64bits(2.1))
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
