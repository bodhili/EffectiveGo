package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{
		name: name, age: age,
	}
	return &p
}

func main() {
	p := newPerson("tom", 43) // type: *main.person
	fmt.Printf("%T \n", p)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	fmt.Println(p.name == "tom")

	p1 := p
	p1.age = 78

	fmt.Println(p)
	fmt.Println(p1)

	sp := person{
		"Tom", 79,
	} // type : main.person

	fmt.Printf("%T \n", sp)
	fmt.Printf("%T \n", p1)

	p0 := &sp
	p0.name = "JO"
	fmt.Printf("%T \n", p0)
}
