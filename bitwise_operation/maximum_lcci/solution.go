package maximum_lcci

// 面试题 16.07. 最大数值 https://leetcode.cn/problems/maximum-lcci/
func maximum(a int, b int) int {
	// abs(x) = ( x ^ (x >> 7)) - ( x >> 7 )
	// max = ( a+b + |a-b| ) / 2
	val := a - b
	return (a + b + ((val ^ (val >> 63)) - (val >> 63))) / 2
}
