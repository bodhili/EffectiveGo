package pkgapi

import "testing"

func TestNilInvoke(t *testing.T) {
	var nilInvoke NilInvokeIfOrNot // no pointer type with ni initialization is ok, if it's pointer type and is not ok
	nilInvoke.Name = "tom"

	t.Logf("invoke name:[%v]", nilInvoke.Name)

	cr := make(chan any, 1)

	var np *NilInvokeIfOrNot
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				cr <- r
				t.Error(r) // if it's pointer type with no initialization, and runtime error: invalid memory address or nil pointer dereference
			}
		}()

		(*np).Name = "jimmy"
	}()

	select {
	case err := <-cr:
		t.Errorf("Recovered err is: %v", err)
	default:
		t.Log("No panic")
	}
}

type NilInvokeIfOrNot struct {
	Name string
}
