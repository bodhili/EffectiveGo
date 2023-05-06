package lan

import (
	"errors"
	"testing"
)

const (
	OOPS      = "oops"
	ErrorExit = "error, exit"
	RECOVER   = "recover:"
)

// panic 终止程序, errors 不会终止程序
func TestOnlyPanic(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)
		if i == 3 {
			panic(ErrorExit)
		}
	}

	t.Log(OOPS)
}

func TestPanicAndRecover(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Log(RECOVER, e)
		}
	}()

	for i := 0; i < 10; i++ {
		if i == 3 {
			panic(ErrorExit)
		}
		t.Log(i)
	}

	t.Log(OOPS)
}

var (
	err = errors.New("not panic")
)

func TestNotPanic(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 3 {
			t.Log(err)
			continue
		}
		t.Log(i)
	}

	t.Log(OOPS)
}
