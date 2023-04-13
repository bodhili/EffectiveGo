package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base // b base 区别  TODO ?
	str  string
}

func main() {
	co := container{
		base: base{
			1,
		},
		str: name,
	}
	fmt.Println(co)
	fmt.Println(co.base.num)

	type describer interface {
		describe() string
	}

	var d describer = co // co.b
	fmt.Println(d.describe())
}
