package lan

import "testing"

func TestStringNotNilAndIsEmpty(t *testing.T) {
	// var name string = nil   not nil

	var name string = "" // A string may be empty, but not nil
	t.Log(name)

	var id uintptr = 10
	t.Log(id)
}
