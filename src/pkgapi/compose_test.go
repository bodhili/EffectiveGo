package pkgapi

import "testing"

type Number interface {
	~int | int8 | int16 | float64
}

func Add[T Number](a T, b T) T {
	return a + b
}

type P int

func TestCompose(t *testing.T) {
	ri := Add(1, 2)
	rf := Add(1.1, 2.2)

	t.Log(ri)
	t.Log(rf)

	p1 := P(1)
	p2 := P(2)

	_ = Add(p1, p2)
}

type Custom interface {
	int | []byte | []rune
}

type Consumer[T Custom] struct {
	ID   int
	Name string
	Data T
}

func TestCustom(t *testing.T) {
	c := Consumer[int]{
		ID:   1,
		Name: "tom",
		Data: 1,
	}

	m := make(CustomMap[int, string])
	m[1] = "a"

	t.Logf("%+v", m)
	t.Logf("%+v", c)
}

type CustomMap[T comparable, V comparable] map[T]V
