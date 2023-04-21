package pkg

import (
	"bytes"
	"testing"
)

func TestInitBuffer(t *testing.T) {
	var buf bytes.Buffer
	buf.Write([]byte("hello world"))

	readByte, err := buf.ReadByte()
	if err == nil {
		t.Log(string(readByte))
	}

	buf.Truncate(4)

	t.Log(&buf)
}

func TestNewBuffer(t *testing.T) {
	buffer := bytes.NewBuffer([]byte("hello"))
	t.Log(buffer.Len())
	t.Log(buffer.Cap())

	_, _ = buffer.ReadByte()

	buffer.Reset()
	t.Log(buffer.Len())
	t.Log(buffer.Cap())

	readByte, _ := buffer.ReadByte()
	t.Log(string(readByte))
}
