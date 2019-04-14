package carbon

type Quarter int8

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
