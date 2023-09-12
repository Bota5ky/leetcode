package factorial_zeros_lcci

// 面试题 16.05. 阶乘尾数 https://leetcode.cn/problems/factorial-zeros-lcci/
func trailingZeroes(n int) int {
	i := 5
	cnt := 0
	for i <= n {
		cnt += n / i
		i *= 5
	}
	return cnt
}
