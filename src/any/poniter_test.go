package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	n := 10
	fmt.Println(read(&n))
}

func read(p *int) (v int) {
	v = *p
	return
}

func newInt() (p *int) {
	var n int
	return &n
}

func TestPanicShutdown(t *testing.T) {
	fmt.Println("test starting...")
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		panic("shutdown child goroutine")
	}()

	wg.Wait()
	fmt.Println("test ending...")
}

func TestPointerCovert(t *testing.T) {
	p := 100
	v := convert(&p)
	fmt.Println(v)

	v0 := convert0(&p)
	fmt.Println(v0)

}

func convert(p *int) (v int32) {
	q := (*int32)(unsafe.Pointer(p))
	*q = 10
	v = *q
	return
}

type Nil struct {
	name string
}

func TestStructSize(t *testing.T) {
	size := unsafe.Sizeof(Nil{})
	t.Log(size)
}

func convert0(p *int) int32 {
	q := (*int32)(unsafe.Pointer(p))
	*q = 10
	return *q
}

func coverts(s []byte) string {
	return *((*string)(unsafe.Pointer(&s)))
}
