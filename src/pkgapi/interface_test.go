package pkgapi

import "testing"

type Animal interface {
	Speak() string
}

type Cat struct {
	Name string
	Age  int
}

type Pig struct{}

func (p *Pig) Speak() string {
	return "pig"
}

func (c *Cat) Speak() string {
	return "Meow"
}

func Change(a Animal) {
	a = &Cat{}
}

func Change0(a *Animal) {
	*a = &Pig{}
}

func TestInterface(t *testing.T) {
	var a Animal = &Cat{
		Name: "tom",
		Age:  10,
	}

	// 注意和上面的区别
	//cat := &Cat{
	//	Name: "tom",
	//	Age:  10,
	//}

	Change(a)
	Change0(&a)

	t.Log(a.Speak())
}
