package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) push(t T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: t}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: t}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) get() []T {
	var elmes []T

	for e := lst.head; e != nil; e = e.next {
		elmes = append(elmes, e.val)
	}
	return elmes
}

func main() {
	var m = map[int]string{1: "1", 2: "2"}
	fmt.Println(MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.push(10)
	lst.push(20)
	lst.push(30)

	fmt.Println(lst.get())

}
