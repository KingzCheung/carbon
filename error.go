package carbon

import "errors"

var (
	//ErrTimeParse 解析时间错误
	ErrTimeParse = errors.New("parse time error")
	//ErrTimestampParse 解析时间戳错误
	ErrTimestampParse = errors.New("parse timestamp error")
)
