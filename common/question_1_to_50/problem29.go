package common

import "math"

// https://leetcode.cn/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	flag := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flag = -1
	}
	res := 0
	n1 := int64(math.Abs(float64(dividend)))
	n2 := int64(math.Abs(float64(divisor)))
	for n1 >= n2 {
		cnt := 0
		for n1 >= (n2 << cnt) {
			cnt++
		}
		n1 -= n2 << (cnt - 1)
		res += 1 << (cnt - 1)
	}
	if flag == -1 {
		res = -res
	}
	return res
}
