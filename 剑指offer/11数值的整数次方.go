package 剑指offer

// 题目: 实现函数double Power(double base, int exponent),
// 求base的exponent次方. 不得使用库函数, 同时不需要考虑大数问题

func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := Power(base, exponent>>1)
	result *= result
	if exponent&0x1 == 1 {
		result *= base
	}
	return result
}
