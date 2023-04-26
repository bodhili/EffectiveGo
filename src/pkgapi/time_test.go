package pkgapi

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

func TestTimeFunc(t *testing.T) {
	before := time.Now()
	after := time.Now().Add(1)

	t.Log(before.Before(after))
	t.Log(before.After(after))
	t.Log(before.Equal(after))
	t.Log(before.Compare(after))

	t.Log(after.Year())
	t.Log(after.Month())
	t.Log(after.Day())
	t.Log(after.YearDay())
	t.Log(after.Hour())
	t.Log(after.Minute())
	t.Log(after.Second())
	t.Log(after.Nanosecond())
	t.Log(after.Weekday())
	t.Log(after.ISOWeek())
	t.Log(after.IsZero())
	t.Log(after.Truncate(-1))
	t.Log(after)
}
