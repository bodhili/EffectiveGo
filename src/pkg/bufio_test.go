package pkg

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestNewBufIOReader(t *testing.T) {
	rs := bufio.Reader{}
	t.Logf("rs is: %+v", rs)

	var reader = bufio.NewReader(strings.NewReader("Hello Reader"))
	t.Logf("reader struct is:[%+v] \n", reader)

	s, err := reader.ReadString('e')
	if err == nil {
		t.Logf("read string is:[%v] \n", s)
	}

	bytes := make([]byte, 10, 10)
	n, err := reader.Read(bytes)
	if err == nil {
		t.Logf("avaible bytes:[%d], string:[%v]", n, string(bytes))
	}
}

func TestNewReaderSize(t *testing.T) {
	s := "Hello New Reader"
	stringReader := strings.NewReader(s)
	reader := bufio.NewReaderSize(stringReader, 1024*5)

	if reader != nil {
		t.Logf("reader type is:[%T]", reader)
		t.Logf("reader:[%+v]", *reader)

		size := reader.Size()
		t.Logf("size:[%d]", size)
	}
}

func TestDiscard(t *testing.T) {
	s := "Hello New Reader"
	stringReader := strings.NewReader(s)
	reader := bufio.NewReaderSize(stringReader, 1024*5)

	var p = make([]byte, 20, 20)

	_, err := reader.Read(p)
	_, _ = reader.ReadByte()
	if reader != nil && err == nil {
		discarded, _ := reader.Discard(5)
		t.Logf("discard byte size:[%d]", discarded)
	}
}

func TestWriter(t *testing.T) {
	w := new(bytes.Buffer)
	writer := bufio.NewWriter(w)

	if writer != nil {
		size, err := writer.WriteString("fill []buf")
		if err == nil {
			t.Logf("write buf size:[%d] \n", size)
		}
	}
}
