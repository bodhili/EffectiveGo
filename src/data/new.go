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

type File struct {
	fd      int
	nepipe  int
	name    string
	dirinfo any
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

func NewFile1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, 0, name, nil}
	f = File{
		fd:      fd,
		name:    name,
		nepipe:  0,
		dirinfo: nil,
	} // difference ways
	return &f
}
