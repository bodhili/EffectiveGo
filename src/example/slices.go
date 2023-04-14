package main

import (
	"fmt"
)

func main() {
	//s := make([]string, 3)
	//fmt.Println("emp: ", s)
	//
	//s[0] = "a"
	//s[1] = string([]byte{1, 2, 3})
	//s[2] = string('c')
	//
	//fmt.Println("Set: ", s)
	//fmt.Println("get: ", s[2])
	//fmt.Println("len: ", len(s))
	//
	//s = append(s, "d")
	//s = append(s, "e", "f")
	//fmt.Println(s)
	//
	//c := make([]string, len(s))
	//copy(c, s)
	//
	//fmt.Println(c)
	//
	//s[0] = "aa"
	//
	//fmt.Println(s)
	//fmt.Println(c)
	//
	//l := s[2:5]
	//fmt.Println("sl1:", l)
	//
	//l = s[:5]
	//fmt.Println("sl2:", l)
	//
	//l = s[2:]
	//fmt.Println("sl3:", l)
	//
	//t := []string{"g", "h", "i"}
	//fmt.Println("dcl:", t)
	//
	//twoD := make([][]int, 3)
	//for i := 0; i < 3; i++ {
	//	innerLen := i + 1
	//	twoD[i] = make([]int, innerLen)
	//	for j := 0; j < innerLen; j++ {
	//		twoD[i][j] = i + j
	//	}
	//}
	//fmt.Println("2d: ", twoD)
	//
	//buf := bytes.NewBuffer([]byte{})
	//buf.WriteByte(1)
	//buf.WriteString("a")
	//
	//v, err := buf.ReadByte()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(v)

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)

	s1 := s[:5]
	fmt.Println(s1)
	fmt.Println(cap(s1))

	s2 := s1[3:]
	fmt.Println(s2)
	fmt.Println(cap(s2))

	s3 := s2[2:]
	fmt.Println(s3)
	fmt.Println(cap(s3))

	/*
		*  当你使用 s4 := s3[:2] 这种方式从一个空切片 s3 中创建一个新的切片 s4 时，Go 会创建一个新的长度为 2，容量为 0 的切片 s4，底层数组指向 s3 的底层数组。由于 s3 的长度为 0，所以 s4 可以从第 0 个位置开始访问底层数组的前两个元素，因此能够访问到底层数组的值。
		   而当你使用 s5 := s3[:] 这种方式从一个空切片 s3 中创建一个新的切片 s5 时，Go 会创建一个新的长度和容量都为 0 的切片 s5，底层数组指向 s3 的底层数组。由于 s3 的长度和容量都为 0，所以 s5 不能访问底层数组中的任何元素，因此仍然是一个空切片
	*/
	s4 := s3[:2]
	fmt.Println(s4)
	fmt.Println(cap(s4))

	s5 := s3[:] // notes: s3切片为空和不为空的区别
	fmt.Println(s5)
	fmt.Println(cap(s5))

	b := []int{1, 2}
	fmt.Println(b)

	fmt.Println(b[:])
}
