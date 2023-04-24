package pkg

import (
	"testing"
)

// golang map构建set集合时，使用空结构体作为value，比用true作为value值 ，占用的空间更少
type Empty struct{}

func TestEmpty(t *testing.T) {
	m := make(map[string]Empty)

	empty := Empty{}

	m["A"] = empty
	m["B"] = empty
	m["B"] = empty

	for k, v := range m {
		t.Logf("key:[%v] v:[%v] \n", k, v)
	}

	e, ok := m["c"]
	t.Log(ok)
	if ok {
		t.Log(e)
	}

	eb, okb := m["B"]
	t.Log(okb)
	if okb {
		t.Log(eb)
		t.Log(eb == empty)
	}
}
