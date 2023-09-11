package maximum_lcci

// https://leetcode-cn.com/problems/maximum-lcci/
func maximum(a int, b int) int {
	// abs(x) = ( x ^ (x >> 7)) - ( x >> 7 )
	// max = ( a+b + |a-b| ) / 2
	val := a - b
	return (a + b + ((val ^ (val >> 63)) - (val >> 63))) / 2
}
