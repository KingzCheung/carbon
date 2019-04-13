package carbon

import (
	"errors"
	"strconv"
	"time"
)

type Unit int
type Quarter int8

const (
	Year Unit = iota
	Month
	Day
	Hour
	Minute
	Second
	Millisecond
	Nanosecond
	Week
)
const (
	January = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

const (
	Sunday    = "Sunday"
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
)

//Now 获取现在时刻时间
func Now() *Carbon {
	t := time.Now()
	return &Carbon{
		Year:        t.Year(),
		Month:       t.Month(),
		Day:         t.Day(),
		Hour:        t.Hour(),
		Minute:      t.Minute(),
		Second:      t.Second(),
		Millisecond: t.Nanosecond() / 1000,
		Nanosecond:  t.Nanosecond(),
		Week:        t.Weekday(),
		time:        t,
	}
}

func (q Quarter) Next() Quarter {
	if q == 4 {
		return 1
	}
	return q + 1
}
func (q Quarter) Last() Quarter {
	if q == 1 {
		return 4
	}
	return q - 1
}

// Tomorrow 获取明天的时间
func Tomorrow() *Carbon {
	now := Now()
	now.AddDay()
	return now
}

// Yesterday Create a Carbon instance for yesterday.
func Yesterday() *Carbon {
	now := Now()
	now.SubDay()
	return now
}

//Create Create a new Carbon instance from a specific date and time.
func Create(year, month, day, hour, minute, second, nanosecond int, tz *time.Location) *Carbon {
	d := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, tz)
	return &Carbon{
		Year:        d.Year(),
		Month:       d.Month(),
		Day:         d.Day(),
		Hour:        d.Hour(),
		Minute:      d.Minute(),
		Second:      d.Second(),
		Millisecond: d.Nanosecond() / 1000,
		Nanosecond:  d.Nanosecond(),
		Week:        d.Weekday(),
		time:        d,
	}
}

//CreateFromDate Create a Carbon instance from just a date.
// The time portion is set to now.
func CreateFromDate(year, month, day int, tz *time.Location) *Carbon {
	//时，分，秒，纳秒都使用当前时间
	now := Now()
	date := time.Date(year, time.Month(month), day, now.Hour, now.Minute, now.Second, now.Nanosecond, tz)
	return &Carbon{
		Year:        date.Year(),
		Month:       date.Month(),
		Day:         date.Day(),
		Hour:        date.Hour(),
		Minute:      date.Minute(),
		Second:      date.Second(),
		Millisecond: date.Nanosecond() / 1000,
		Nanosecond:  date.Nanosecond(),
		Week:        date.Weekday(),
		time:        date,
	}
}

// CreateFromTime Create a Carbon instance from just a time.
// The date portion is set to today.
func CreateFromTime(hour, minute, second int, tz *time.Location) *Carbon {
	//日期使用当前时间
	now := Now()
	date := time.Date(now.Year, now.Month, now.Day, hour, minute, second, now.Nanosecond, tz)
	return &Carbon{
		Year:        date.Year(),
		Month:       date.Month(),
		Day:         date.Day(),
		Hour:        date.Hour(),
		Minute:      date.Minute(),
		Second:      date.Second(),
		Millisecond: date.Nanosecond() / 1000,
		Nanosecond:  date.Nanosecond(),
		Week:        date.Weekday(),
		time:        date,
	}
}

func Parse(layout, value string) *Carbon {
	parse, _ := time.Parse(layout, value)
	return &Carbon{
		Year:        parse.Year(),
		Month:       parse.Month(),
		Day:         parse.Day(),
		Hour:        parse.Hour(),
		Minute:      parse.Minute(),
		Second:      parse.Second(),
		Millisecond: parse.Nanosecond() / 1000,
		Nanosecond:  parse.Nanosecond(),
		Week:        parse.Weekday(),
		time:        parse,
	}
}

func ParseFromLocale(layout, value string, tz *time.Location) *Carbon {
	parse, _ := time.ParseInLocation(layout, value, tz)
	return &Carbon{
		Year:        parse.Year(),
		Month:       parse.Month(),
		Day:         parse.Day(),
		Hour:        parse.Hour(),
		Minute:      parse.Minute(),
		Second:      parse.Second(),
		Millisecond: parse.Nanosecond() / 1000,
		Nanosecond:  parse.Nanosecond(),
		Week:        parse.Weekday(),
		time:        parse,
	}
}

//Carbon
type Carbon struct {
	Year, Day, Hour, Minute, Second, Millisecond, Nanosecond int
	Month                                                    time.Month
	Week                                                     time.Weekday
	time                                                     time.Time
}

func (c *Carbon) Format(layout string) string {
	return c.time.Format(layout)
}

//Timestamp 获取时间戳
func (c *Carbon) Unix() int64 { return c.time.Unix() }
func (c *Carbon) Timestamp() int64 {
	return c.time.Unix()
}

func (c *Carbon) IsLeapYear() bool {
	return (c.Year%100 != 0 && c.Year%4 == 0) || (c.Year%400 == 0)
}

func (c *Carbon) CountDayForYear() int {
	if c.IsLeapYear() {
		return 366
	} else {
		return 365
	}
}

//CountDayForMonth 返回每个月的天数
func (c *Carbon) CountDayForMonth() int {
	month := c.Format("1")
	m, _ := strconv.Atoi(month)
	switch m {
	case January, March, May, July, August, October, December:
		return 31
	case February:
		if c.IsLeapYear() {
			return 29
		} else {
			return 28
		}
	case April, June, September, November:
		return 30
	default:
		return 30

	}
}

func (c *Carbon) addValToUnit(unit Unit, value int) error {
	switch unit {
	case Year:
		c.Year += value
		c.time = c.time.Add(time.Duration(value) * time.Hour * 24 * time.Duration(c.CountDayForYear()))
	case Month:
		c.Month += time.Month(value)
		c.time = c.time.Add(time.Duration(value) * time.Hour * 24 * time.Duration(c.CountDayForMonth()))
	case Day:
		c.Day += value
		c.time = c.time.Add(time.Duration(value) * time.Hour * 24)
	case Hour:
		c.Hour += value
		c.time = c.time.Add(time.Duration(value) * time.Hour)
	case Minute:
		c.Minute += value
		c.time = c.time.Add(time.Duration(value) * time.Minute)
	case Second:
		c.Second += value
		c.time = c.time.Add(time.Duration(value) * time.Second)
	case Millisecond:
		c.Millisecond += value
		c.time = c.time.Add(time.Duration(value) * time.Nanosecond * 1000)
	case Nanosecond:
		c.Nanosecond += value
		c.time = c.time.Add(time.Duration(value) * time.Nanosecond)
	case Week:
		c.Week += time.Weekday(value)
		c.time = c.time.Add(time.Duration(value) * 7 * 24 * time.Hour)
	default:
		return errors.New("添加类型错误")
	}
	return nil
}

//Add given units or interval to the current instance.
func (c *Carbon) Add(unit Unit, value int) *Carbon {
	_ = c.addValToUnit(unit, value)
	return c
}
func (c *Carbon) subValToUnit(unit Unit, value int) error {
	switch unit {
	case Year:
		c.Year -= value
		c.time = c.time.Add(-time.Duration(value) * time.Hour * 24 * time.Duration(c.CountDayForYear()))
	case Month:
		c.Month -= time.Month(value)
		c.time = c.time.Add(-time.Duration(value) * time.Hour * 24 * time.Duration(c.CountDayForMonth()))
	case Day:
		c.Day -= value
		c.time = c.time.Add(-time.Duration(value) * time.Hour * 24)
	case Hour:
		c.Hour -= value
		c.time = c.time.Add(-time.Duration(value) * time.Hour)
	case Minute:
		c.Minute -= value
		c.time = c.time.Add(-time.Duration(value) * time.Minute)
	case Second:
		c.Second -= value
		c.time = c.time.Add(-time.Duration(value) * time.Second)
	case Millisecond:
		c.Millisecond -= value
		c.time = c.time.Add(-time.Duration(value) * time.Nanosecond * 1000)
	case Nanosecond:
		c.Nanosecond -= value
		c.time = c.time.Add(-time.Duration(value) * time.Nanosecond)
	case Week:
		c.Week -= time.Weekday(value)
		c.time = c.time.Add(-time.Duration(value) * 7 * 24 * time.Hour)
	default:
		return errors.New("添加类型错误")
	}
	return nil
}

func (c *Carbon) Sub(unit Unit, value int) *Carbon {
	_ = c.subValToUnit(unit, value)
	return c
}

//AddYear Add one year to the instance.
func (c *Carbon) AddYear() *Carbon {
	return c.Add(Year, 1)
}

//AddYears Add years to the instance.
//$value count passed in
func (c *Carbon) AddYears(value int) *Carbon {
	return c.Add(Year, value)
}

//AddMonth Add one month to the instance.
func (c *Carbon) AddMonth() *Carbon {
	return c.Add(Month, 1)
}

//AddMonths Add months to the instance.
//value count passed in.
func (c *Carbon) AddMonths(value int) *Carbon {
	return c.Add(Month, value)
}

func (c *Carbon) AddWeek() *Carbon {
	return c.Add(Week, 1)
}

func (c *Carbon) AddWeeks(value int) *Carbon {
	return c.Add(Week, value)
}

//AddDay Add one day to the instance.
func (c *Carbon) AddDay() *Carbon {
	return c.Add(Day, 1)
}

func (c *Carbon) AddDays(value int) *Carbon {
	return c.Add(Day, value)
}

func (c *Carbon) SubDay() *Carbon {
	return c.Sub(Day, 1)
}

func (c *Carbon) SubDays(value int) *Carbon {
	return c.Sub(Day, value)
}

//AddHour Add one hour to the instance.
func (c *Carbon) AddHour() *Carbon {
	return c.Add(Hour, 1)
}

func (c *Carbon) AddHours(value int) *Carbon {
	return c.Add(Hour, value)
}

func (c *Carbon) SubHour() *Carbon {
	return c.Sub(Hour, 1)
}

func (c *Carbon) SubHours(value int) *Carbon {
	return c.Sub(Hour, value)
}

//AddMinute Add one minute to the instance.
func (c *Carbon) AddMinute() *Carbon {
	return c.Add(Minute, 1)
}

func (c *Carbon) AddMinutes(value int) *Carbon {
	return c.Add(Minute, value)
}

func (c *Carbon) SubMinute() *Carbon {
	return c.Sub(Minute, 1)
}

func (c *Carbon) SubMinutes(value int) *Carbon {
	return c.Sub(Minute, value)
}

//AddSecond Add one second to the instance.
func (c *Carbon) AddSecond() *Carbon {
	return c.Add(Second, 1)
}

func (c *Carbon) AddSeconds(value int) *Carbon {
	return c.Add(Second, value)
}

func (c *Carbon) SubSecond() *Carbon {
	return c.Sub(Second, 1)
}

func (c *Carbon) SubSeconds(value int) *Carbon {
	return c.Sub(Second, value)
}

//is

func (c *Carbon) IsSunday(value int) bool {
	return c.time.Weekday().String() == Sunday

}
func (c *Carbon) IsMonday(value int) bool {
	return c.time.Weekday().String() == Monday
}
func (c *Carbon) IsTuesday(value int) bool {
	return c.time.Weekday().String() == Tuesday
}
func (c *Carbon) IsWednesday(value int) bool {
	return c.time.Weekday().String() == Wednesday
}
func (c *Carbon) IsThursday(value int) bool {
	return c.time.Weekday().String() == Thursday
}
func (c *Carbon) IsFriday(value int) bool {
	return c.time.Weekday().String() == Friday
}
func (c *Carbon) IsSaturday(value int) bool {
	return c.time.Weekday().String() == Saturday
}
func (c *Carbon) IsCurrentYear(value int) bool {
	curYear := Now().Format("01")
	return curYear == strconv.Itoa(c.Year)
}

func (c *Carbon) IsNextYear(value int) bool {
	curYear := Now().Year
	return c.Year-curYear == 1
}

func (c *Carbon) IsLastYear(value int) bool {
	curYear := Now().Year
	return curYear-c.Year == 1
}

func (c *Carbon) IsCurrentDay(value int) bool {
	curDay := Now().Day
	return curDay == c.Day
}

func (c *Carbon) IsNextDay(value int) bool {
	curDay := Now().Day
	return c.Day-curDay == 1
}
func (c *Carbon) IsLastDay(value int) bool {
	curDay := Now().Day
	return curDay-c.Day == 1
}

func (c *Carbon) IsCurrentHour(value int) bool {
	curHour := Now().Hour
	return curHour == c.Hour
}

func (c *Carbon) IsNextHour(value int) bool {
	curHour := Now().Hour

	return c.Hour-curHour == 1
}
func (c *Carbon) IsLastHour(value int) bool {
	curHour := Now().Hour

	return curHour-c.Hour == 1
}

func (c *Carbon) IsCurrentWeek(value int) bool {
	curWeek := Now().Week

	return curWeek == c.Week
}

func (c *Carbon) IsNextWeek(value int) bool {
	curWeek := Now().Week
	return c.Week-curWeek == 1
}
func (c *Carbon) IsLastWeek(value int) bool {
	curWeek := Now().Week

	return curWeek-c.Week == 1
}

func (c *Carbon) IsCurrentMinute(value int) bool {
	curMinute := Now().Minute
	return curMinute == c.Minute
}

func (c *Carbon) IsNextMinute(value int) bool {
	curMinute := Now().Minute

	return c.Minute-curMinute == 1
}
func (c *Carbon) IsLastMinute(value int) bool {
	curMinute := Now().Minute

	return curMinute-c.Minute == 1
}

func (c *Carbon) IsCurrentSecond(value int) bool {
	curSec := Now().Second
	return curSec == c.Second
}

func (c *Carbon) IsNextSecond(value int) bool {
	curSec := Now().Second

	return c.Second-curSec == 1
}
func (c *Carbon) IsLastSecond(value int) bool {
	curSec := Now().Second

	return curSec-c.Second == 1
}

func (c *Carbon) IsCurrentMonth(value int) bool {
	curMonth := Now().Month
	return curMonth == c.Month
}

func (c *Carbon) IsNextMonth(value int) bool {
	curMonth := Now().Month
	return c.Month-curMonth == 1
}
func (c *Carbon) IsLastMonth(value int) bool {
	curMonth := Now().Month

	return curMonth-c.Month == 1
}

func (c *Carbon) CurrentQuarter() Quarter {
	switch {
	case 1 <= c.Month && c.Month <= 3:
		return 1
	case 4 <= c.Month && c.Month <= 6:
		return 2
	case 7 <= c.Month && c.Month <= 9:
		return 3
	case 10 <= c.Month && c.Month <= 12:
		return 4
	default:
		return 0
	}
}

func (c *Carbon) IsCurrentQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter()
}

func (c *Carbon) IsNextQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter().Next()
}
func (c *Carbon) IsLastQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter().Last()
}
