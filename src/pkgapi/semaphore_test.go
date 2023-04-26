package pkgapi

import (
	"sync"
	"testing"
)

type Semaphore struct {
	capacity int
	acquired int
	mutex    sync.Mutex
	cond     *sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
	semaphore := &Semaphore{
		acquired: 0,
		capacity: permits,
		mutex:    sync.Mutex{},
	}

	semaphore.cond = sync.NewCond(&semaphore.mutex)
	return semaphore
}

func (semaphore *Semaphore) Acquire() {
	semaphore.mutex.Lock()
	defer semaphore.mutex.Unlock()

	for semaphore.acquired == semaphore.capacity { // 存在惊群效应问题（仅此为练习golang语言使用）
		semaphore.cond.Wait()
	}

	semaphore.acquired++
}

func (semaphore *Semaphore) Release() {
	semaphore.mutex.Lock()
	defer semaphore.mutex.Unlock()

	if semaphore.acquired > 0 {
		semaphore.acquired--
	}
	semaphore.cond.Broadcast()
}

func TestSemaphore(t *testing.T) {
	semaphore := NewSemaphore(5)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		semaphore.Acquire()
		defer wg.Done()
		defer semaphore.Release()

		t.Log("acquire successfully...")
	}()

	wg.Wait()
}
