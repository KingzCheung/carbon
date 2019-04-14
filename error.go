package carbon

import "errors"

var (
	TimeParseError      = errors.New("parse time error")
	TimestampParseError = errors.New("parse timestamp error")
)
