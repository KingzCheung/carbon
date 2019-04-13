package main

import (
	"carbon"
	"fmt"
)

func main() {
	//获取现在时间
	now := carbon.Now()
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
}
