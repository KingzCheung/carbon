package carbon

import "errors"

var (
	ErrTimeParse      = errors.New("parse time error")
	ErrTimestampParse = errors.New("parse timestamp error")
)
