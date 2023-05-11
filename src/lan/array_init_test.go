package lan

import (
	"testing"
)

func TestInitArray(t *testing.T) {
	var a [10]int
	a[0] = 1
	t.Log(a)
}
