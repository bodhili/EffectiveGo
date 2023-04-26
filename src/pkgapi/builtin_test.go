package pkgapi

import (
	"math/rand"
	"testing"
	"time"
)

type Session struct{}

type channel chan Session

type Cookie struct {
	Id      string
	Timeout time.Time
}

type cookie interface {
	SetCookie(id string)
	GetCookieId() string
	timeout(timeout time.Time)
}

func (c *Cookie) SetCookie(id string) {
	c.Id = id
}

func (c *Cookie) GetCookieId() string {
	return c.Id
}

func (c *Cookie) timeout(timeout time.Time) {
	c.Timeout = timeout
}

func TestPolymorphism(t *testing.T) {
	var defined cookie = &Cookie{
		Id:      string(rune(rand.Intn(1 << 8))),
		Timeout: time.Now(),
	}

	defined.SetCookie(string(rune(1)))
	defined.timeout(time.Date(2023, 4, 18, 10, 10, 10, 9, time.Local))

	if defined.GetCookieId() != string(rune(1)) {
		t.Errorf("Cookie id is not equalivet.")
	}

	t.Logf("cookie:%+v", defined)
}

func TestChannel(t *testing.T) {
	c := make(channel) // c := make(chan Session)

	session := Session{}
	go func() {
		c <- session
		defer close(c)
	}()

	var s Session
	s = <-c

	if s != session {
		t.Errorf("Chan type[alias: channel] is error.")
	}
}

func TestBool(t *testing.T) {
	v := testBool(true)
	if v != true {
		t.Errorf("bool type[TRUE] is test error.")
	}

	b := testBool(false)
	if b != false {
		t.Errorf("bool type[FALSE] is test error.")
	}
}

func testBool(b bool) (v bool) {
	v = b
	return
}

func TestByte(t *testing.T) {
	var char byte = 'a'
	var char1 int = 'a'

	if int(char) != char1 {
		t.Error("Byte type is error.")
	}
}

func TestComparable(t *testing.T) {
	c := make(chan bool)

	go func() {
		c <- true
		defer close(c)
	}()
	r := <-c

	if r != true {
		t.Error("")
	}
}

func TestSliceAppend(t *testing.T) {
	s := []int{1, 2, 3}
	s = append(s, 4)

	if len(s) != 4 && cap(s) != 4 {
		t.Error("Slice append is error.")
	}

	s = make([]int, 0, 10)
	if len(s) != 0 && cap(s) != 10 {
		t.Errorf("Slice append is error. len: %d, cap:%d", len(s), cap(s))
	}
}

func TestMakeChan(t *testing.T) {
	c := make(chan string)
	defer close(c)

	go func() {
		c <- "A"
	}()

	s := <-c

	if s != "A" {
		t.Error("S is not A.")
	}
}

func TestMapMake(t *testing.T) {
	m := make(map[string]int)

	t.Logf("len(map): %d", len(m))

	m["A"] = 1

	if len(m) != 1 {
		t.Error("Make map is error.")
	}

	mp := new(map[string]int) // pointer map
	t.Logf("len(map): %d", len(*mp))
	for k, v := range *mp {
		t.Logf("map key: %v, map value: %v", k, v)
	}

	v, ok := m["A"]
	if ok {
		t.Logf("v: %v", v)
	} else {
		t.Log(ok)
	}

	i, ok1 := (*mp)["a"]
	if ok1 {
		t.Logf("v: %v", i)
	} else {
		t.Log(ok1)
	}
}
