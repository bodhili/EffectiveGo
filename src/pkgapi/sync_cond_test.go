package pkgapi

import (
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		cond.Wait()
		mu.Unlock()

		t.Log("goroutine 1: condition is ready")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		mu.Lock()
		cond.Signal()
		cond.Broadcast()
		mu.Unlock()

		t.Log("goroutine 2: condition is ready")
	}()

	wg.Wait()
}

func TestRWMutex(t *testing.T) {
	var rwLock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		rwLock.RLock()

		t.Log("R Lock")

		rwLock.RUnlock()
	}()

	wg.Wait()
}
