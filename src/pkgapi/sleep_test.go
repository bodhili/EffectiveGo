package pkgapi

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	d := 500 * time.Millisecond

	time.Sleep(d)
}

// timer 一次性延迟执行，并非周期性
func TestTimer(t *testing.T) {
	d := 2 * time.Second

	timer1 := time.NewTimer(d)
	<-timer1.C
	t.Log("timer1 expired")

	timer2 := time.NewTimer(d)
	go func() {
		<-timer2.C
		t.Log("timer2 expired")
	}()

	stop := timer2.Stop() // 提前终止不会执行 <-timer.C 过期处理 t.Log("timer2 expired")
	if stop {
		t.Log("timer2 stop")
	}

	timer2.Reset(d) // 执行 <-timer.C 过期处理  t.Log("timer2 expired")

	time.Sleep(5 * time.Second)
}

func TestTimerFor(t *testing.T) {
	d := 20 * time.Millisecond

	timer1 := time.NewTimer(d)
	defer timer1.Stop()

	for {
		select {
		case <-timer1.C:
			t.Log("t")
			return
		default:
			t.Log("default")
		}
	}
}

func TestAfter(t *testing.T) {
	t.Logf("now:[%v]\n", time.Now().String())
	d := 2 * time.Second
	timer := time.NewTimer(d)

	c := time.After(5 * time.Second)
	timer.C = c
	<-timer.C
	t.Logf("timer expired, time:[%v] \n", time.Now().String())

	time.Sleep(10 * time.Second)
}

func TestAfterFunc(t *testing.T) {
	t.Logf("now:[%v]\n", time.Now().String())
	d := 2 * time.Second

	time.AfterFunc(d, func() {
		t.Logf("time now:[%v]\n", time.Now().String())
	})

	t.Logf("timer expired, time:[%v] \n", time.Now().String())
	t.Logf("timer expired, time:[%v] \n", time.Now().String())
	t.Logf("timer expired, time:[%v] \n", time.Now().String())

	time.Sleep(10 * time.Second)
}
