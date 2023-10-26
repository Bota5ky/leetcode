package power_of_two

// 231. 2 的幂 https://leetcode.cn/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
