package main

import (
	"bytes"
	"fmt"
	"sync"
)

const (
	name = iota + 1
	age
	id
	die = 90
	_
	others = iota + 1
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	p := new(SyncedBuffer)
	var v SyncedBuffer

	fmt.Println(others)

	fmt.Println(p)
	fmt.Println(v)
}
