package pkgapi

import (
	"bytes"
	"testing"
)

func TestEqual(t *testing.T) {
	s0, s1 := newSlices(3, 3)

	equal := bytes.Equal(s0, s1)

	if equal {
		t.Log("slice equals")
	} else {
		panic("not equals")
	}
}

func TestCompare(t *testing.T) {
	s0, s1 := newSlices(3, 3)

	c0 := bytes.Compare(s0, s1)
	if c0 == 0 {
		t.Log("compare =")
	}

	s00 := []byte{0, 2, 1}
	c00 := bytes.Compare(s00, s1)
	if c00 == 1 {
		t.Log("compare >")
	}

	s11 := []byte{0, 2, 1}
	c11 := bytes.Compare(s0, s11)
	if c11 == -1 {
		t.Log("compare <")
	}
}

func newSlices(a byte, b byte) (s0 []byte, s1 []byte) {
	s0 = make([]byte, 0, a)
	s1 = make([]byte, 0, b)

	var i byte = 0
	for ; i < a; i++ {
		s0 = append(s0, i)
	}

	var j byte = 0
	for ; j < b; j++ {
		s1 = append(s1, j)
	}
	return
}

func TestExplode(t *testing.T) {
	b0 := []byte{2, 3, 4, 5, 2, 3, 4, 6}
	b1 := []byte{2, 3, 4}

	count := bytes.Count(b0, b1)
	t.Log(count)

	s0 := "ABC@ABC"
	s1 := "A#A"

	sb0 := []byte(s0)
	sb1 := []byte(s1)

	bytes.Count(sb0, sb1)
	t.Log(count)
}

func TestCutPrefix(t *testing.T) {
	s0 := "ABC@ABC"
	s1 := "A"

	sb0 := []byte(s0)
	sb1 := []byte(s1)

	after, found := bytes.CutPrefix(sb0, sb1)
	if found {
		t.Log(string(after))
	}
}

func TestCutSuffix(t *testing.T) {
	s0 := "ABC@ABC"
	s1 := "ABC"

	sb0 := []byte(s0)
	sb1 := []byte(s1)

	before, found := bytes.CutSuffix(sb0, sb1)
	if found {
		t.Log(string(before))
	}
}

func TestCut(t *testing.T) {
	s0 := "ABC@CBA"
	s1 := "@"

	sb0 := []byte(s0)
	sb1 := []byte(s1)

	before, after, found := bytes.Cut(sb0, sb1)
	if found {
		t.Log(string(before))
		t.Log(string(after))
	}
}

func TestIndex(t *testing.T) {
	s0 := "ABC@CBA"
	s1 := "#"

	sb0 := []byte(s0)
	sb1 := []byte(s1)

	index := bytes.Index(sb0, sb1)
	t.Log(index)
}

func TestReplaceAll(t *testing.T) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	old := []byte("fox")
	news := []byte("cat")

	replaced := bytes.ReplaceAll(s, old, news)

	t.Log(string(replaced))
}

func TestTrimSpace(t *testing.T) {
	b := []byte("  hello, world!   ")
	t.Logf("Original bytes: %q\n", b)

	trimmed := bytes.TrimSpace(b)
	t.Log(trimmed)
}
