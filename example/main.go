package main

import (
	"fmt"
	"github.com/kingzcheung/carbon"
	"time"
)

func main() {
	//获取现在时间
	now := carbon.Now()
	// 返回toString
	fmt.Println(now) //2019-04-14 16:09:20
	//获取昨天时间
	//yesterday := carbon.Yesterday()
	//获取明天时间
	//tomorrow := carbon.Tomorrow()
	//获取当前时间戳
	fmt.Println(now.Timestamp())
	//格式时间
	fmt.Println(now.Format("2006-01-02 15:04:05 pm"))
	//解析时间
	fmt.Println(carbon.Parse("2006-01-02 15:04:05", "2019-01-01 21:12:22").Format("2006-01-02 15:04:05 pm"))

	//从时间戳中解析
	carbon.CreateFromTimestamp(12434535453)
	//从时间戳字符串中解析
	_, _ = carbon.CreateFromTimestampString("12434535453")
	//从时间中解析
	carbon.CreateFromTime(12, 12, 12, time.Local)
	//从时间字符串中解析
	_, _ = carbon.CreateFromTimeString("12:12:59", time.Local)
	//从日期中解析
	carbon.CreateFromDate(2019, 4, 12, time.Local)

	//判断
	carbon.Now().IsLeapYear()       // 判断是不是闰年
	carbon.Now().IsCurrentQuarter() //判断是不是当季
	carbon.Now().IsCurrentDay()     //判断是不是今天
	carbon.Now().IsFriday()         //判断是不是周五
	carbon.Now().IsWeekend()        //是不是周末
	carbon.Now().IsWeekday()        //是不是工作日
	carbon.Now().IsCurrentHour()    //是不是当前小时
	carbon.Now().IsLastDay()        //是不是昨天
	carbon.Now().IsNextWeek()       //是不是下周
	//...

	//偏移时间
	carbon.Now().AddYear() //添加一年

	carbon.Now().AddMonth()           //添加一个月
	carbon.Now().Add(carbon.Month, 1) //同上

	carbon.Now().AddDays(5) //添加5天

	carbon.Now().Sub(carbon.Day, 5) //减去5天
	carbon.Now().SubDays(5)         //同上

	//...

	//格式化时间返回
	carbon.Now().ToDateTimeString()      //返回 "2006-01-02 15:04:05" 格式时间字符串
	carbon.Now().ToDateString()          //返回 "2006-01-02" 格式时间字符串
	carbon.Now().ToFormattedDateString() //返回 "Jan 02,2006" 格式字符串

	//比较前后
	carbon.Now().After(carbon.Yesterday())
	carbon.Now().Before(carbon.Yesterday())

}
