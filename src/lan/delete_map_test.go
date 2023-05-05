package lan

import "testing"

func TestDelete(t *testing.T) {
	invokes := make(map[string]string)

	invokes["k1"] = "K1"
	invokes["k2"] = "K2"

	delete(invokes, "k1")
	delete(invokes, "k3")

	t.Log(invokes)

	for k, v := range invokes {
		t.Logf("k:v = %v:%v \n", k, v)
	}
}
