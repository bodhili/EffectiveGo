package pkg

import (
	"testing"
	"time"
)

// 周期性执行任务
func TestTicker(t *testing.T) {
	c := time.Tick(100 * time.Millisecond)

	ticker := time.NewTicker(1 * time.Second)
	ticker.C = c
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.Log("heart is living...")
			return
		default:
			t.Log("default")
		}
	}
}
