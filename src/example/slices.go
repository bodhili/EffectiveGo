package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = string([]byte{1, 2, 3})
	s[2] = string('c')

	fmt.Println("Set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	s[0] = "aa"

	fmt.Println(s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte(1)
	buf.WriteString("a")

	v, err := buf.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

}
