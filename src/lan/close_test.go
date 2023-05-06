package lan

import (
	"testing"
)

func TestCloseChanNotBuf(t *testing.T) {
	c := make(chan string) // 不带缓冲区,deadlock

	defer close(c)
	c <- "A"

	var p string = <-c
	t.Log(p)
}

func TestCloseChanWithBuf(t *testing.T) {
	c := make(chan string, 1) // 带缓冲区, not deadlock

	defer close(c)
	c <- "A"

	var p string = <-c
	t.Log(p)
}

func TestCloseChanWithMultiGoroutine(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		t.Log("sent 1")
		ch <- 2
		t.Log("sent 2")
		ch <- 3
		t.Log("sent 3")
	}()

	t.Log("received", <-ch)
	t.Log("received", <-ch)
	t.Log("received", <-ch)
}
