package carbon

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Unit 单位类型
type Unit int

const (
	// Year 年
	Year Unit = iota
	// Month 月
	Month
	// Day 日
	Day
	// Hour 时
	Hour
	// Minute 分
	Minute
	// Second 秒
	Second
	// Millisecond 毫秒
	Millisecond
	// Microsecond 微秒
	Microsecond
	// Nanosecond 纳秒
	Nanosecond
	// Week 周
	Week
)
const (
	// January 一月
	January = 1 + iota
	// February 二月
	February
	// March 三月
	March
	// April 四月
	April
	// May 五月
	May
	// June 六月
	June
	// July 七月
	July
	// August 八月
	August
	// September 九月
	September
	// October 十月
	October
	// November 十一月
	November
	// December 十二月
	December
)

const (
	// Sunday 周日
	Sunday = "Sunday"
	// Monday 周一
	Monday = "Monday"
	// Tuesday 周二
	Tuesday = "Tuesday"
	// Wednesday 三
	Wednesday = "Wednesday"
	// Thursday 四
	Thursday = "Thursday"
	// Friday 五
	Friday = "Friday"
	// Saturday 六
	Saturday = "Saturday"
)

// Now 获取现在时刻时间
func Now(locale ...string) *Carbon {
	t := time.Now()
	return &Carbon{
		Year:        t.Year(),
		Month:       t.Month(),
		Day:         t.Day(),
		Hour:        t.Hour(),
		Minute:      t.Minute(),
		Second:      t.Second(),
		Millisecond: t.Nanosecond() / 1000000,
		Microsecond: t.Nanosecond() / 1000,
		Nanosecond:  t.Nanosecond(),
		Week:        t.Weekday(),
		time:        t,
	}
}

// Today 获取今天日期,时间重置为0时0分0秒
func Today() *Carbon {
	now := Now()
	date := CreateFromDate(now.Year, int(now.Month), now.Day, time.Local)
	return &Carbon{
		Year:        now.Year,
		Month:       now.Month,
		Day:         now.Day,
		Hour:        0,
		Minute:      0,
		Second:      0,
		Millisecond: 0,
		Microsecond: 0,
		Nanosecond:  0,
		Week:        now.Week,
		time:        date.time,
	}
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

// Create Create a new Carbon instance from a specific date and time.
func Create(year, month, day, hour, minute, second int, tz *time.Location) *Carbon {
	d := time.Date(year, time.Month(month), day, hour, minute, second, 0, tz)
	return &Carbon{
		Year:        d.Year(),
		Month:       d.Month(),
		Day:         d.Day(),
		Hour:        d.Hour(),
		Minute:      d.Minute(),
		Second:      d.Second(),
		Millisecond: 0,
		Microsecond: 0,
		Nanosecond:  0,
		Week:        d.Weekday(),
		time:        d,
	}
}

// CreateFromDate Create a Carbon instance from just a date.
// The time portion is set to now.
func CreateFromDate(year, month, day int, tz *time.Location) *Carbon {
	// 时，分，秒，纳秒都使用当前时间
	now := Now()
	date := time.Date(year, time.Month(month), day, now.Hour, now.Minute, now.Second, now.Nanosecond, tz)
	return &Carbon{
		Year:        date.Year(),
		Month:       date.Month(),
		Day:         date.Day(),
		Hour:        date.Hour(),
		Minute:      date.Minute(),
		Second:      date.Second(),
		Millisecond: date.Nanosecond() / 1000000,
		Microsecond: date.Nanosecond() / 1000,
		Nanosecond:  date.Nanosecond(),
		Week:        date.Weekday(),
		time:        date,
	}
}

func CreateFromGo(date time.Time) *Carbon {
	return &Carbon{
		Year:        date.Year(),
		Month:       date.Month(),
		Day:         date.Day(),
		Hour:        date.Hour(),
		Minute:      date.Minute(),
		Second:      date.Second(),
		Millisecond: date.Nanosecond() / 1000000,
		Microsecond: date.Nanosecond() / 1000,
		Nanosecond:  date.Nanosecond(),
		Week:        date.Weekday(),
		time:        date,
	}
}

// CreateFromTime Create a Carbon instance from just a time.
// The date portion is set to today.
func CreateFromTime(hour, minute, second int, tz *time.Location) *Carbon {
	// 日期使用当前时间
	now := Now()
	date := time.Date(now.Year, now.Month, now.Day, hour, minute, second, now.Nanosecond, tz)
	return &Carbon{
		Year:        date.Year(),
		Month:       date.Month(),
		Day:         date.Day(),
		Hour:        date.Hour(),
		Minute:      date.Minute(),
		Second:      date.Second(),
		Millisecond: date.Nanosecond() / 1000000,
		Microsecond: date.Nanosecond() / 1000,
		Nanosecond:  date.Nanosecond(),
		Week:        date.Weekday(),
		time:        date,
	}
}

// CreateFromTimeString 解析冒号过来的时间字符串
func CreateFromTimeString(value string, tz *time.Location) (*Carbon, error) {

	v := strings.Split(value, ":")
	if len(v) != 3 {
		return &Carbon{}, ErrTimeParse
	}
	hour, err := strconv.Atoi(v[0])
	if err != nil {
		return &Carbon{}, ErrTimeParse
	}
	minute, err := strconv.Atoi(v[1])
	if err != nil {
		return &Carbon{}, ErrTimeParse
	}
	second, err := strconv.Atoi(v[2])
	if err != nil {
		return &Carbon{}, ErrTimeParse
	}

	return CreateFromTime(hour, minute, second, tz), nil
}

// Parse 通过格式化解析时间字符串为 Carbon 类型
func Parse(layout, value string) *Carbon {
	parse, _ := time.Parse(layout, value)
	return &Carbon{
		Year:        parse.Year(),
		Month:       parse.Month(),
		Day:         parse.Day(),
		Hour:        parse.Hour(),
		Minute:      parse.Minute(),
		Second:      parse.Second(),
		Millisecond: parse.Nanosecond() / 1000000,
		Microsecond: parse.Nanosecond() / 1000,
		Nanosecond:  parse.Nanosecond(),
		Week:        parse.Weekday(),
		time:        parse,
	}
}

// ParseFromLocale 基于时区解析时间字符串为 Carbon 类型
func ParseFromLocale(layout, value string, tz *time.Location) *Carbon {
	parse, _ := time.ParseInLocation(layout, value, tz)
	return &Carbon{
		Year:        parse.Year(),
		Month:       parse.Month(),
		Day:         parse.Day(),
		Hour:        parse.Hour(),
		Minute:      parse.Minute(),
		Second:      parse.Second(),
		Millisecond: parse.Nanosecond() / 1000000,
		Microsecond: parse.Nanosecond() / 1000,
		Nanosecond:  parse.Nanosecond(),
		Week:        parse.Weekday(),
		time:        parse,
	}
}

// CreateFromFormat as same as ParseFromLocale.
func CreateFromFormat(layout, value string, tz *time.Location) *Carbon {
	return ParseFromLocale(layout, value, tz)
}

// CreateFromTimestamp 从时间戳中解析 Carbon
func CreateFromTimestamp(value int64) *Carbon {
	t := time.Unix(value, 0)
	return &Carbon{
		Year:        t.Year(),
		Month:       t.Month(),
		Day:         t.Day(),
		Hour:        t.Hour(),
		Minute:      t.Minute(),
		Second:      t.Second(),
		Millisecond: t.Nanosecond() / 1000000,
		Microsecond: t.Nanosecond() / 1000,
		Nanosecond:  t.Nanosecond(),
		Week:        t.Weekday(),
		time:        t,
	}
}

// CreateFromTimestampString 同 CreateFromTimestamp 类似，只不过参数为时间戳字符串，并返回解析错误
func CreateFromTimestampString(value string) (*Carbon, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return &Carbon{}, ErrTimestampParse
	}
	return CreateFromTimestamp(i), nil
}

// Carbon 处理时间
type Carbon struct {
	Year, Day, Hour, Minute, Second, Millisecond, Microsecond, Nanosecond int
	Month                                                                 time.Month
	Week                                                                  time.Weekday
	time                                                                  time.Time
}

// Format 通过时间格式指定格式化时间并返回
func (c *Carbon) Format(layout string) string {
	return c.time.Format(layout)
}

// Unix 获取时间戳
func (c *Carbon) Unix() int64 { return c.time.Unix() }

// Timestamp 同 Unix 获取时间戳
func (c *Carbon) Timestamp() int64 {
	return c.time.Unix()
}

// IsLeapYear 判断是不是闰年
func (c *Carbon) IsLeapYear() bool {
	return (c.Year%100 != 0 && c.Year%4 == 0) || (c.Year%400 == 0)
}

// CountDayForYear 返回一年的天数，如果是闰年则返回366天
func (c *Carbon) CountDayForYear() int {
	if c.IsLeapYear() {
		return 366
	}
	return 365
}

// CountDayForMonth 返回每个月的天数
func (c *Carbon) CountDayForMonth() int {
	month := c.Format("1")
	m, _ := strconv.Atoi(month)
	switch m {
	case January, March, May, July, August, October, December:
		return 31
	case February:
		if c.IsLeapYear() {
			return 29
		}
		return 28
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

// Add given units or interval to the current instance.
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

// func (c *Carbon)Locale(name string) *Carbon  {
//	local,err:= time.LoadLocation(name)
//	if err != nil {
//		local = time.Local
//	}
//	c.time.Location()
// }

// Sub 从当前结构体减去 value 的 unit
func (c *Carbon) Sub(unit Unit, value int) *Carbon {
	_ = c.subValToUnit(unit, value)
	return c
}

// AddYear Add one year to the instance.
func (c *Carbon) AddYear() *Carbon {
	return c.Add(Year, 1)
}

// AddYears Add years to the instance.
// $value count passed in
func (c *Carbon) AddYears(value int) *Carbon {
	return c.Add(Year, value)
}

// AddMonth Add one month to the instance.
func (c *Carbon) AddMonth() *Carbon {
	return c.Add(Month, 1)
}

// AddMonths Add months to the instance.
// value count passed in.
func (c *Carbon) AddMonths(value int) *Carbon {
	return c.Add(Month, value)
}

// AddWeek 从当前结构体加上 1 周
func (c *Carbon) AddWeek() *Carbon {
	return c.Add(Week, 1)
}

// AddWeeks 从当前结构体加上 value 周
func (c *Carbon) AddWeeks(value int) *Carbon {
	return c.Add(Week, value)
}

// AddDay Add one day to the instance.
func (c *Carbon) AddDay() *Carbon {
	return c.Add(Day, 1)
}

// AddDays 从当前结构体加上 value 天
func (c *Carbon) AddDays(value int) *Carbon {
	return c.Add(Day, value)
}

// SubDay 从当前结构体减去 1 天
func (c *Carbon) SubDay() *Carbon {
	return c.Sub(Day, 1)
}

// SubDays 从当前结构体减去 value 天
func (c *Carbon) SubDays(value int) *Carbon {
	return c.Sub(Day, value)
}

// AddHour Add one hour to the instance.
func (c *Carbon) AddHour() *Carbon {
	return c.Add(Hour, 1)
}

// AddHours 从当前结构体加上 value 小时
func (c *Carbon) AddHours(value int) *Carbon {
	return c.Add(Hour, value)
}

// SubHour 从当前结构体减去 1 小时
func (c *Carbon) SubHour() *Carbon {
	return c.Sub(Hour, 1)
}

// SubHours 从当前结构体减去 value 小时
func (c *Carbon) SubHours(value int) *Carbon {
	return c.Sub(Hour, value)
}

// AddMinute Add one minute to the instance.
func (c *Carbon) AddMinute() *Carbon {
	return c.Add(Minute, 1)
}

// AddMinutes 从当前结构体减去 value 分钟
func (c *Carbon) AddMinutes(value int) *Carbon {
	return c.Add(Minute, value)
}

// SubMinute 从当前结构体减去 1 分钟
func (c *Carbon) SubMinute() *Carbon {
	return c.Sub(Minute, 1)
}

// SubMinutes 从当前结构体减去 value 分钟
func (c *Carbon) SubMinutes(value int) *Carbon {
	return c.Sub(Minute, value)
}

// AddSecond Add one second to the instance.
func (c *Carbon) AddSecond() *Carbon {
	return c.Add(Second, 1)
}

// AddSeconds 从当前结构体加上 value 秒
func (c *Carbon) AddSeconds(value int) *Carbon {
	return c.Add(Second, value)
}

// SubSecond 从当前结构体减去 1 秒
func (c *Carbon) SubSecond() *Carbon {
	return c.Sub(Second, 1)
}

// SubSeconds 从当前结构体减去 value 秒
func (c *Carbon) SubSeconds(value int) *Carbon {
	return c.Sub(Second, value)
}

// IsSunday 判断是不是周日
func (c *Carbon) IsSunday() bool {
	return c.time.Weekday().String() == Sunday

}

// IsMonday 判断是不是周一
func (c *Carbon) IsMonday() bool {
	return c.time.Weekday().String() == Monday
}

// IsTuesday 判断是不是周二
func (c *Carbon) IsTuesday() bool {
	return c.time.Weekday().String() == Tuesday
}

// IsWednesday 判断是不是周三
func (c *Carbon) IsWednesday() bool {
	return c.time.Weekday().String() == Wednesday
}

// IsThursday 判断是不是周四
func (c *Carbon) IsThursday() bool {
	return c.time.Weekday().String() == Thursday
}

// IsFriday 判断是不是周五
func (c *Carbon) IsFriday() bool {
	return c.time.Weekday().String() == Friday
}

// IsSaturday 判断是不是周六
func (c *Carbon) IsSaturday() bool {
	return c.time.Weekday().String() == Saturday
}

// IsWeekend 判断是不是周末
func (c *Carbon) IsWeekend() bool {
	return c.IsSunday() || c.IsSaturday()
}

// IsWeekday 是否是工作日
func (c *Carbon) IsWeekday() bool {
	return !c.IsWeekend()
}

// IsCurrentYear 判断是不是今年
func (c *Carbon) IsCurrentYear() bool {
	curYear := Now().Format("01")
	return curYear == strconv.Itoa(c.Year)
}

// IsNextYear 判断是不是明天
func (c *Carbon) IsNextYear() bool {
	curYear := Now().Year
	return c.Year-curYear == 1
}

// IsLastYear 判断是不是去年
func (c *Carbon) IsLastYear() bool {
	curYear := Now().Year
	return curYear-c.Year == 1
}

// IsCurrentDay 判断是不是今天
func (c *Carbon) IsCurrentDay() bool {
	curDay := Now().Day
	return curDay == c.Day
}

// IsNextDay 判断是不是明天
func (c *Carbon) IsNextDay() bool {
	curDay := Now().Day
	return c.Day-curDay == 1
}

// IsLastDay 判断是不是昨天
func (c *Carbon) IsLastDay() bool {
	curDay := Now().Day
	return curDay-c.Day == 1
}

// IsCurrentHour 判断是否是当前小时
func (c *Carbon) IsCurrentHour() bool {
	curHour := Now().Hour
	return curHour == c.Hour
}

// IsNextHour 判断是否是下一小时
func (c *Carbon) IsNextHour() bool {
	curHour := Now().Hour

	return c.Hour-curHour == 1
}

// IsLastHour 判断是否是上一小时
func (c *Carbon) IsLastHour() bool {
	curHour := Now().Hour

	return curHour-c.Hour == 1
}

// IsCurrentWeek 判断是不是当前周
func (c *Carbon) IsCurrentWeek() bool {
	curWeek := Now().Week

	return curWeek == c.Week
}

// IsNextWeek 判断是不是下周
func (c *Carbon) IsNextWeek() bool {
	curWeek := Now().Week
	return c.Week-curWeek == 1
}

// IsLastWeek 判断是否是上周
func (c *Carbon) IsLastWeek() bool {
	curWeek := Now().Week

	return curWeek-c.Week == 1
}

// IsCurrentMinute 判断是否是当前分钟
func (c *Carbon) IsCurrentMinute() bool {
	curMinute := Now().Minute
	return curMinute == c.Minute
}

// IsNextMinute 判断是否是下一分钟
func (c *Carbon) IsNextMinute() bool {
	curMinute := Now().Minute

	return c.Minute-curMinute == 1
}

// IsLastMinute 判断是否是上一分钟
func (c *Carbon) IsLastMinute() bool {
	curMinute := Now().Minute

	return curMinute-c.Minute == 1
}

// IsCurrentSecond 判断是否是当前秒
func (c *Carbon) IsCurrentSecond() bool {
	curSec := Now().Second
	return curSec == c.Second
}

// IsNextSecond 判断是否是下一秒
func (c *Carbon) IsNextSecond() bool {
	curSec := Now().Second

	return c.Second-curSec == 1
}

// IsLastSecond 判断是否是上一秒
func (c *Carbon) IsLastSecond() bool {
	curSec := Now().Second

	return curSec-c.Second == 1
}

// IsCurrentMonth 判断是否是当前月
func (c *Carbon) IsCurrentMonth() bool {
	curMonth := Now().Month
	return curMonth == c.Month
}

// IsNextMonth 判断是否是下一个月
func (c *Carbon) IsNextMonth() bool {
	curMonth := Now().Month
	return c.Month-curMonth == 1
}

// IsLastMonth 判断是否是上一个月
func (c *Carbon) IsLastMonth() bool {
	curMonth := Now().Month

	return curMonth-c.Month == 1
}

// CurrentQuarter 返回当前季度
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

// IsCurrentQuarter 判断是否是当前季度
func (c *Carbon) IsCurrentQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter()
}

// IsNextQuarter 判断是否是下一季度
func (c *Carbon) IsNextQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter().Next()
}

// IsLastQuarter 判断是否是上一季度
func (c *Carbon) IsLastQuarter() bool {
	return c.CurrentQuarter() == Now().CurrentQuarter().Last()
}

// ToDateTimeString 返回 "2006-01-02 15:04:05" 时间格式的字符串
func (c *Carbon) ToDateTimeString() string {
	layout := "2006-01-02 15:04:05"
	return c.Format(layout)
}

// ToDateString 返回日期时间字符串
func (c *Carbon) ToDateString() string {
	layout := "2006-01-02"
	return c.Format(layout)
}

// ToTimeString 返回时间字符串
func (c *Carbon) ToTimeString() string {
	layout := "15:04:05"
	return c.Format(layout)
}

// ToFormattedDateString 返回格式化可读性的日期
func (c *Carbon) ToFormattedDateString() string {
	layout := "Jan 02,2006"
	return c.Format(layout)
}

// String as same as ToDateTimeString
func (c *Carbon) String() string {
	return c.ToDateTimeString()
}

// ToMap Conversion to Map
func (c *Carbon) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"year":        c.Year,
		"month":       c.Month,
		"day":         c.Day,
		"hour":        c.Hour,
		"minute":      c.Minute,
		"second":      c.Second,
		"millisecond": c.Millisecond,
		"microsecond": c.Microsecond,
		"nanosecond":  c.Nanosecond,
		"week":        c.Week,
	}
}

// Comparison

// EqualTo 比较两个时间是否一样
func (c *Carbon) EqualTo(carbon *Carbon) bool {
	if c.Timestamp() == carbon.Timestamp() {
		return true
	}
	return false
}

// NotEqualTo 比较两个时间是否不一样
func (c *Carbon) NotEqualTo(carbon *Carbon) bool {

	if c.Timestamp() != carbon.Timestamp() {
		return true
	}
	return false
}

// GreaterThan 比较时间是否比目标大
func (c *Carbon) GreaterThan(carbon *Carbon) bool {
	if c.Timestamp() > carbon.Timestamp() {
		return true
	}
	return false
}

// GreaterThanOrEqualTo 比较时间是否比目标大于或者等于
func (c *Carbon) GreaterThanOrEqualTo(carbon *Carbon) bool {
	if c.Timestamp() >= carbon.Timestamp() {
		return true
	}
	return false
}

// LessThan 比较时间是否比目标小
func (c *Carbon) LessThan(carbon *Carbon) bool {
	if c.Timestamp() < carbon.Timestamp() {
		return true
	}
	return false
}

// LessThanOrEqualTo 比较时间是否比目标小或者等于
func (c *Carbon) LessThanOrEqualTo(carbon *Carbon) bool {
	if c.Timestamp() <= carbon.Timestamp() {
		return true
	}
	return false
}

// Between 比较当前值是否是在 first 和second 之间
func (c *Carbon) Between(first, second *Carbon) bool {
	if c.Timestamp() < first.Timestamp() || c.Timestamp() > second.Timestamp() {
		return false
	}

	return true
}

// After 如果c代表的时间点在u之后，返回真；否则返回假。
func (c *Carbon) After(u *Carbon) bool {
	return c.time.After(u.time)
}

// Before 如果c代表的时间点在u之前，返回真；否则返回假。
func (c *Carbon) Before(u *Carbon) bool {
	return c.time.Before(u.time)
}

const (
	SecondMax = 60
	MinuteMax = SecondMax * 60 // 360,0
	HourMax   = MinuteMax * 24 // 864,00
	DayMax    = HourMax * 30   // 2,592,000
	WeekMax   = DayMax * 7     // 18,144,000
	MonthsMax = DayMax * 12
	YearMax   = MonthsMax * 9999
)

func (c *Carbon) DiffForHumans(other ...*Carbon) string {
	var o *Carbon
	var val bytes.Buffer
	var len = len(other)

	if len > 0 {
		o = other[0]
	} else {
		o = Now()
	}
	c1 := c.Timestamp()
	o1 := o.Timestamp()
	diff := o1 - c1
	if diff < 0 {
		diff = ^diff + 1
	}
	switch {
	case diff < SecondMax:
		n := diff
		val.WriteString(fmt.Sprintf("%d ", diff))
		val.WriteString("second")
		if n > 1 {
			val.WriteString("s")
		}
	case diff < MinuteMax:
		n := diff / 60
		val.WriteString(fmt.Sprintf("%d ", n))
		val.WriteString("minute")
		if n > 1 {
			val.WriteString("s")
		}
	case diff < HourMax:
		n := diff / 60 / 60
		val.WriteString(fmt.Sprintf("%d ", diff/60/60))
		val.WriteString("hour")
		if n > 1 {
			val.WriteString("s")
		}
	case diff < DayMax:
		n := diff / 60 / 60 / 24
		val.WriteString(fmt.Sprintf("%d ", n))
		val.WriteString("day")
		if n > 1 {
			val.WriteString("s")
		}
	case diff < WeekMax:
		n := diff / 60 / 60 / 24 / 7
		val.WriteString(fmt.Sprintf("%d ", n))
		val.WriteString("week")
		if n > 1 {
			val.WriteString("s")
		}
	case diff < MonthsMax:
		n := diff / 60 / 60 / 24 / 30
		val.WriteString(fmt.Sprintf("%d ", n))
		val.WriteString("month")
		if n > 1 {
			val.WriteString("s")
		}
	default:
		n := diff / 60 / 60 / 24 / 30 / 365
		val.WriteString(fmt.Sprintf("%d ", n))
		val.WriteString("years")
		if n > 1 {
			val.WriteString("s")
		}
	}

	if o1 < c1 {
		if len > 0 {
			val.WriteString(" after")
		} else {
			val.WriteString(" from now")
		}
	} else {
		if len > 0 {
			val.WriteString(" before")
		} else {
			val.WriteString(" ago")
		}
	}

	return val.String()
}
