package pkgapi

import (
	"sync"
	"testing"
)

type CyclicBarrier struct {
	n         int
	count     int
	mutex     sync.Mutex
	cond      *sync.Cond
	waitGroup *sync.WaitGroup
	f         func()
}

func NewCyclicBarrier(n int, f func()) *CyclicBarrier {
	cb := &CyclicBarrier{
		n:         n,
		count:     0,
		mutex:     sync.Mutex{},
		waitGroup: &sync.WaitGroup{},
		f:         f,
	}
	cb.cond = sync.NewCond(&cb.mutex)
	return cb
}

func (cb *CyclicBarrier) Await() {
	cb.mutex.Lock()
	cb.count++
	if cb.count < cb.n {
		cb.cond.Wait()
	} else {
		cb.count = 0
		cb.waitGroup.Add(cb.n)
		cb.cond.Broadcast()
	}
	cb.mutex.Unlock()
	cb.waitGroup.Done()
}

func (cb *CyclicBarrier) Wait() {
	cb.waitGroup.Wait()
}

func (cb *CyclicBarrier) Done() {
	defer cb.f()
}

func TestCyclicBarrier(t *testing.T) {
	barrier := NewCyclicBarrier(1, func() {
		t.Log("call back")
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer barrier.Done()

		barrier.Wait()
	}()

	wg.Wait()
}
