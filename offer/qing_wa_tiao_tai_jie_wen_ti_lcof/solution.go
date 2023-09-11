package offer

// 剑指 Offer 10- II. 青蛙跳台阶问题 https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
	a := 1
	b := 1
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	}
	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		if c >= 1000000007 {
			c %= 1000000007
		}
		a = b
		b = c
	}
	return c
}
