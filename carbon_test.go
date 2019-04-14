package carbon

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestToday(t *testing.T) {
	as := assert.New(t)
	today := Today()
	now := Now()
	as.Equal(today.Year, now.Year, "Today error. Today.Year should be equal Now().Year")
	as.Equal(today.Month, now.Month, "Today error. Today.Month should be equal Now().Month")
	as.Equal(today.Day, now.Day, "Today error. Today.Day should be equal Now().Day")
	as.Equal(today.Hour, 0, "Today error.Today.Hour should be equal zero")
	as.Equal(today.Minute, 0, "Today error.Today.Minute should be equal zero")
	as.Equal(today.Second, 0, "Today error.Today.Second should be equal zero")
	as.Equal(today.Millisecond, 0, "Today error.Today.Millisecond should be equal zero")
	as.Equal(today.Nanosecond, 0, "Today error.Today.Nanosecond should be equal zero")
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
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		as.Error(err)
	}
	c := Create(2018, 10, 22, 12, 12, 12, 10, tz)
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

func TestParseFromLocale(t *testing.T) {
	type args struct {
		layout string
		value  string
		tz     *time.Location
	}
	london, _ := time.LoadLocation("Europe/London")
	tests := []struct {
		name string
		args args
		want *Carbon
	}{
		{"2006 01 02", args{"2006 01 02", "2019 04 04", time.Local}, &Carbon{Year: 2019, Month: time.Month(04), Day: 4}},
		{"2006 01 02", args{"2006 01 02", "2019 04 04", london}, &Carbon{Year: 2019, Month: time.Month(04), Day: 4}},
	}
	as := assert.New(t)
	for _, tt := range tests {
		got := ParseFromLocale(tt.args.layout, tt.args.value, tt.args.tz)
		as.Equal(got.Day, tt.want.Day, "ParseFromLocale().Day should be equal %d,but got %d", tt.want.Day, got.Day)
	}
}

func TestCreateFromTimeString(t *testing.T) {
	type args struct {
		value string
		tz    *time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    *Carbon
		wantErr bool
	}{
		{"test1", args{"12:12:59", time.Local}, CreateFromTime(12, 12, 59, time.Local), false},
		{"test2", args{"12::12:59", time.Local}, CreateFromTime(0, 0, 0, time.Local), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFromTimeString(tt.args.value, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFromTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Hour, tt.want.Hour) {
				t.Errorf("CreateFromTimeString() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Minute, tt.want.Minute) {
				t.Errorf("CreateFromTimeString() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Second, tt.want.Second) {
				t.Errorf("CreateFromTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFromTimestampString(t *testing.T) {
	type args struct {
		value string
	}
	london, _ := time.LoadLocation("Europe/London")
	tests := []struct {
		name    string
		args    args
		want    *Carbon
		wantErr bool
	}{
		{"zero time", args{"0"}, CreateFromDate(1970, 1, 1, london), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFromTimestampString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFromTimestampString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Year, tt.want.Year) {
				t.Errorf("CreateFromTimestampString() = %v, want %v", got, tt.want)
			}
		})
	}
}
