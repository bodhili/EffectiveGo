package pkgapi

import (
	"container/list"
	"testing"
)

func TestNewList(t *testing.T) {
	l := list.List{}
	t.Logf("l len:[%d], l.element:[%v]\n", l.Len(), l.Back())
}

func TestListPP(t *testing.T) {
	l := list.New()

	back := l.PushBack(4)
	front := l.PushFront(1)

	_ = l.InsertBefore(3, back)
	_ = l.InsertAfter(2, front)

	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}
