package pkgapi

import (
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	s := "hello word"
	clone := strings.Clone(s)

	t.Log(s == clone)

	s1 := "hello"
	t.Log(strings.Compare(s, s1))

	var sb strings.Builder

	sa, err := sb.WriteString("A")
	if err == nil {
		t.Log(sa)
	}

	sbb, err := sb.WriteString("B")
	if err == nil {
		t.Log(sbb)
	}

	t.Log(&sb)
}
