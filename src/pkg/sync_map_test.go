package pkg

import (
	"sync"
	"testing"
)

// [sync.Map] safely in goroutines
// m := make(map[string]string) is not safely in goroutines
func TestNewSyncMap(t *testing.T) {
	var m sync.Map

	value, ok := m.Load("A")
	t.Log(ok)
	if ok {
		t.Log(value)
	}
}

func TestSyncMapLoadAndStore(t *testing.T) {
	var m sync.Map
	actual, loaded := m.LoadOrStore("A", "a")
	t.Log(loaded)
	t.Log(actual)

	value, l := m.LoadAndDelete("A")
	t.Log(value)
	t.Log(l)
	m.LoadOrStore("A", "B")

	deleted := m.CompareAndDelete("A", "a")
	t.Log(deleted)

	//for k, v := range m {}  it's not ok

	m.Range(func(key, value any) bool {
		t.Logf("key:[%v] v:[%v]", key, value)
		return true
	})

	swap := m.CompareAndSwap("A", "a", "C")
	t.Log(swap)
	andSwap := m.CompareAndSwap("A", "B", "C")
	t.Log(andSwap)

	previous, b := m.Swap("A", "D")
	t.Log(b)
	t.Log(previous)
}
