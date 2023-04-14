package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}

	a1 := a[0:2]
	fmt.Println(a1)
	fmt.Printf("a1 len: %d, cap: %d \n", len(a1), cap(a1))

	a2 := a1[:]
	fmt.Println(a2)
	fmt.Printf("a2 len: %d, cap: %d \n", len(a2), cap(a2))

	a3 := a2[:3]
	fmt.Println(a3)
	fmt.Printf("a3 len: %d, cap: %d \n", len(a3), cap(a3))

	a4 := a3[3:] // <nil> slice
	fmt.Println(a4)
	fmt.Printf("a4 len: %d, cap: %d \n", len(a4), cap(a4))

	a5 := a4[:] // <nil> slice
	fmt.Println(a5)
	fmt.Printf("a5 len: %d, cap: %d \n", len(a5), cap(a5))

	a6 := a5[:1] // [i:j] i 从底层数组或者切片开始访问的位置, j 指定切片的长度 (非空切片，则j = 可访问底层数组长度), 如果是空数组或者空切片，则 i = j = 0
	fmt.Println(a6)
	fmt.Printf("a6 len: %d, cap: %d \n", len(a6), cap(a6))

	s := make([]int, 0, 1)
	fmt.Println(s)

	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
