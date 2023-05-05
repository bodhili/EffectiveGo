package lan

import "testing"

func TestMakeNewSlice(t *testing.T) {
	s0 := make([]string, 0, 10)

	var s1 *[]string
	s1 = new([]string)

	s0 = append(s0, "Hello")
	s1 = &s0

	t.Log(s0)
	t.Log(*s1)
	t.Log()
}

func TestMakeNewMap(t *testing.T) {
	m0 := make(map[string]string)
	m0["k1"] = "K1"

	var m1 *map[string]string
	m1 = new(map[string]string)
	*m1 = make(map[string]string)
	(*m1)["k2"] = "K2"

	t.Log(m0)
	t.Log(*m1)
}
