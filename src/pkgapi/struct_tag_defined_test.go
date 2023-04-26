package pkgapi

import (
	"reflect"
	"testing"
)

type User struct {
	Name string `json:"name" db:"user_name"`
}

func TestStructDefinedTags(t *testing.T) {
	user := &User{
		Name: "tom",
	}

	r := reflect.TypeOf(*user)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		t.Logf("%s: json=%s, db=%s\n", field.Name, field.Tag.Get("json"), field.Tag.Get("db"))
	}
}
