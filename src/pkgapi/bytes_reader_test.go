package pkgapi

import (
	"bytes"
	"testing"
)

func TestBytesReader(t *testing.T) {
	s := "ABC@CBA"
	b := []byte(s)
	reader := bytes.NewReader(b)

	seek, err2 := reader.Seek(2, 0)
	t.Log(seek)
	t.Log(err2)

	t.Log(reader.Size())
	t.Log(reader.Len())
	t.Log(reader.ReadByte())

	_, _ = reader.Seek(2, 1)

	ch, size, err := reader.ReadRune()
	t.Log(ch)
	t.Log(size)
	t.Log(err)
}
