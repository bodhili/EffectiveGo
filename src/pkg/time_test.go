package pkg

import (
	"testing"
	"time"
)

func TestTimeString(t *testing.T) {
	now := time.Now()
	t.Log(now.String())
	t.Log(now.GoString())
	t.Log(now.Format(time.DateTime))
}

func TestTimeParse(t *testing.T) {
	layout := time.DateTime
	st := "2023-04-24 10:30:00"

	parse, err := time.Parse(layout, st)
	if err == nil {
		t.Log(parse)
	} else {
		t.Error(err)
	}
}
func TestTimeParseInLocation(t *testing.T) {
	layout := time.DateTime
	st := "2023-04-24 10:30:00"

	location := time.Local
	location, _ = time.LoadLocation("Local")
	location, _ = time.LoadLocation("UTC")

	inLocation, err := time.ParseInLocation(layout, st, location)
	if err == nil {
		t.Log(inLocation)
	} else {
		t.Error(err)
	}
}
