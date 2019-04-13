package carbon

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	as := assert.New(t)
	times := []struct {
		layout, value, chkLayout, chkVal string
	}{
		{"2006 01", "2019 04", "01", "04"},
		{"2006-1-02", "2019-12-24", "2", "24"},
		{"2006-1-2", "2019-2-4", "01-02", "02-04"},
	}
	for _, val := range times {
		f := Parse(val.layout, val.value).Format(val.chkLayout)
		as.Equal(f, val.chkVal)
	}
}

func TestYesterday(t *testing.T) {
	as := assert.New(t)
	yesterday := Yesterday()
	yes := time.Now().Add(-time.Duration(1) * time.Hour * 24)
	as.Equal(yesterday.Day, yes.Day(), "Yesterday error.")
}

func TestTomorrow(t *testing.T) {
	tomorrow := Tomorrow()
	as := assert.New(t)
	tom := time.Now().Add(time.Duration(1) * time.Hour * 24)
	as.Equal(tom.Day(), tomorrow.Day, "Tomorrow error.")
}

func TestCarbon_Timestamp(t *testing.T) {
	ts := Now().Timestamp()
	as := assert.New(t)
	as.Equal(len(strconv.Itoa(int(ts))), 10, "Timestamp error")
}

func TestCreate(t *testing.T) {
	as := assert.New(t)
	c := Create(2018, 10, 22, 12, 12, 12, 10, time.Local)
	as.Equal(c.Timestamp(), int64(1540181532), "Create parse error. Timestamp should be equal 1540181532")
}

func TestCreateFromDate(t *testing.T) {
	cd := CreateFromDate(2018, 14, 13, time.Local)
	as := assert.New(t)
	as.Equal(cd.Year, 2019, "Year parse error.Year should be equal 2019")
	as.Equal(cd.Month.String(), "February", "Month parse error.Month should be equal February")
	as.Equal(cd.Day, 13, "Day parse error.Day should be equal 2019")
}

func TestCreateFromTime(t *testing.T) {
	as := assert.New(t)
	ct := CreateFromTime(25, 15, 12, time.Local)
	as.Equal(ct.Hour, 1, "Hour parse error.Hour should be equal 1")
	as.Equal(ct.Minute, 15, "Minute parse error.Minute should be equal 15")
	as.Equal(ct.Second, 12, "Second parse error.Second should be equal 12")
}

func TestCarbon_CountDayForMonth(t *testing.T) {
	as := assert.New(t)
	nc := Now().CountDayForMonth()
	as.Equal(nc, 30, "CountDayForMonth error.CountDayForMonth should be equal 30.")
}

func TestCarbon_AddYears(t *testing.T) {
	as := assert.New(t)
	addYears := Now().AddYears(5)
	as.Equal(addYears.Year, 2024, "addYears.Year error.addYears.Year should be equal 2024.")
}

func TestCarbon_AddYear(t *testing.T) {
	as := assert.New(t)
	addYear := Now().AddYear()
	as.Equal(addYear.Year, 2020, "addYear.Year should be equal 2020")
}

func TestCarbon_AddMonth(t *testing.T) {
	as := assert.New(t)
	addMonth := Now().AddMonth()
	as.Equal(addMonth.Month.String(), "May", "addMonth.Month should be equal May")
}

func TestCarbon_AddMonths(t *testing.T) {
	as := assert.New(t)
	addMonths := CreateFromDate(2019, 10, 10, time.Local)
	addMonths.AddMonths(2)
	as.Equal(addMonths.Month.String(), "December", "addMonths.Month should be equal December")
}
