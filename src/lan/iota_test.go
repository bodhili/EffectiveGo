package lan

// iota 重新计算表达式
const (
	ONE = iota
	TWO
	THREE
	_
	FOUR
	FIVE   = iota * 2 // 5 * 2
	SIX               // 6 *2
	SEVERN = iota
	TEN
)
