# Carbon

![GitHub](https://img.shields.io/github/license/kingzcheung/carbon.svg)  [![godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://github.com/KingzCheung/carbon)  [![Build Status](https://travis-ci.org/KingzCheung/carbon.svg?branch=master)](https://travis-ci.org/KingzCheung/carbon)

### 安装

```
go get github.com/KingzCheung/carbon
```

### 用法

```go
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
```