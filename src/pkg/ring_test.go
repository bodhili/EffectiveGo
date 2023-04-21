package pkg

import (
	"container/ring"
	"testing"
)

type Cache struct {
	data  map[string]string // 存储数据的 map
	ring  *ring.Ring        // 环形链表
	count int               // 当前缓存数量
	size  int               // 最大缓存数量
}

func NewCache(size int) *Cache {
	return &Cache{
		data:  make(map[string]string),
		ring:  ring.New(size),
		count: 0,
		size:  size,
	}
}

func (c *Cache) Get(key string) string {
	if value, ok := c.data[key]; ok {
		for i := 0; i < c.count; i++ {
			if c.ring.Value.(string) == key {
				c.ring = c.ring.Move(i)
				break
			}
			c.ring = c.ring.Next()
		}
		return value
	}
	return ""
}

func (c *Cache) Put(key, value string) {
	if c.count >= c.size {
		delete(c.data, c.ring.Value.(string))
		c.ring = c.ring.Next()
	} else {
		c.count++
	}

	c.data[key] = value
	c.ring.Value = key
	c.ring = c.ring.Next()
}

func TestRing(t *testing.T) {
	c := NewCache(3)
	c.Put("a", "1")
	c.Put("b", "2")
	c.Put("c", "3")
	c.Put("d", "4")

	e := c.Get("d")
	t.Log(e)
}
