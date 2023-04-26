package pkgapi

import (
	"sync"
	"testing"
)

type MyPoolObject struct {
	Name string
	Age  int
}

func TestSyncPool(t *testing.T) {
	var pool = &sync.Pool{
		New: func() interface{} {
			return &MyPoolObject{}
		},
	}

	o1 := pool.Get().(*MyPoolObject)
	o1.Name = "Alice"
	o1.Age = 18

	t.Logf("o1:[%+v] \n", o1)

	o2 := pool.Get().(*MyPoolObject)
	t.Logf("o2:[%+v] \n", o2)

	clean(&o1)
	clean(&o2)

	pool.Put(o1)
	pool.Put(o2)

	o3 := pool.Get().(*MyPoolObject)
	t.Logf("o3:[%+v] \n", o3)
}

func clean(obj **MyPoolObject) {
	*obj = &MyPoolObject{}
}

// clean(obj **MyPoolObject) 注意区别，sync.Pool
//func clean0(obj *MyPoolObject) {
//	*obj = MyPoolObject{}
//}
