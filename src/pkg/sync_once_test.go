package pkg

import (
	"sync"
	"testing"
)

// 使用 sync.Once 可以避免在多个 goroutine 中重复执行某个初始化函数的情况，从而避免资源浪费和数据竞争
func TestSyncOnce(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		t.Log("once again")
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(func() {
			t.Log("not once.")
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(func() {
			t.Log("not once..")
		})
	}()

	wg.Wait()
}
