package leetcode

//https://leetcode-cn.com/problems/factorial-zeros-lcci/
func trailingZeroes(n int) int {
	i := 5
	cnt := 0
	for i <= n {
		cnt += n / i
		i *= 5
	}
	return cnt
}
