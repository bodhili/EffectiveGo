package pkgapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestTypeKind(t *testing.T) {
	for _, v := range []any{"hi", 42, func() {}} {
		tp := reflect.TypeOf(v)
		t.Log(tp)
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			t.Log(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			t.Log(v.Int())
		case reflect.Func:
			t.Logf("func:[%v]", v.Interface())
		case reflect.Interface:
			t.Logf("interface:[%v]", v.Interface())
		default:
			t.Logf("unhandled kind %v", v.Kind())
		}
	}
}

func TestReflectStruct(t *testing.T) {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	sr := reflect.TypeOf(s)
	f, ok := sr.FieldByName("F")

	if ok {
		color := f.Tag.Get("color")
		species := f.Tag.Get("species")

		t.Logf("color:[%v] species:[%v] \n", color, species)
	}
}

func TestModifyStructValue(t *testing.T) {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{
		F: "tom",
	}

	v := reflect.ValueOf(&s).Elem()
	f := v.FieldByName("F")
	f.SetString("FF")

	t.Log(s)
}

func TestReflectLookup(t *testing.T) {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	sr := reflect.TypeOf(s)
	for i := 0; i < sr.NumField(); i++ {
		field := sr.Field(i)
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				t.Log("blank")
			} else {
				t.Log(alias)
			}
		} else {
			t.Log("not specified")
		}
	}
}

func TestReflectStructOf(t *testing.T) {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
}

func TestReflectOf(t *testing.T) {
	type user struct {
		firstname string
		lastname  string
	}

	u := user{
		firstname: "John",
		lastname:  "Doe",
	}

	s := reflect.ValueOf(u)
	t.Logf("Name:[%v]", s.FieldByName("firstname"))
}

func TestReflectOf0(t *testing.T) {
	type user struct {
		firstName string
		age       int
		lastName  string
	}

	type data struct {
		user
		firstName string
		lastName  string
	}

	u := data{
		user:      user{"Embedded John", 10, "Embedded Doe"},
		firstName: "John",
		lastName:  "Doe",
	}

	s := reflect.ValueOf(u)
	f1 := s.FieldByIndex([]int{0, 2})
	f2 := s.FieldByName("lastName")
	t.Log("embedded last name:", f1)
	t.Log("embedded last name:", f2)
}

func TestReflectNilTypeOf(t *testing.T) {
	writeType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil))

	t.Log(fileType.Implements(writeType))
}

func TestReflectDeepEqual(t *testing.T) {
	type user struct {
		Name string
		Age  int
		Die  bool
	}

	u1 := &user{
		Name: "tom",
		Age:  18,
		Die:  true,
	}

	u2 := &user{
		Name: "tom",
		Age:  18,
		Die:  true,
	}

	ifOrNot := reflect.DeepEqual(*u1, *u2)
	t.Log(ifOrNot)
	t.Log(*u1 == *u2)
	t.Log(u1 == u2)

	t.Log(reflect.DeepEqual(u1, u2))
}
