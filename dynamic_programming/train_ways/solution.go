package qing_wa_tiao_tai_jie_wen_ti_lcof

// LCR 127. 跳跃训练 https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func trainWays(n int) int {
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
