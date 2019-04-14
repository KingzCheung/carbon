package carbon

//Quarter 季度类型
type Quarter int8

//Next 返回下一季度
func (q Quarter) Next() Quarter {
	if q == 4 {
		return 1
	}
	return q + 1
}

//Last 返回上一个季度
func (q Quarter) Last() Quarter {
	if q == 1 {
		return 4
	}
	return q - 1
}
