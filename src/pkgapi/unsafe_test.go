package pkgapi

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	size := len(arr)
	p := uintptr(unsafe.Pointer(&arr))

	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = p
	sh.Len = size
	sh.Cap = size

	fmt.Println(data)

	runtime.KeepAlive(arr)
}
