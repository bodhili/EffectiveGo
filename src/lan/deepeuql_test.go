package lan

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDeepEqualStrings(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	equal := reflect.DeepEqual(got, want)
	t.Log(equal)

	t.Logf("%p", &got)
	t.Logf("%p", &want)
}

func TestDeepEqualMap(t *testing.T) {
	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(c, d))

	var b []string
	var a []string = nil
	fmt.Println(reflect.DeepEqual(a, b))
}
