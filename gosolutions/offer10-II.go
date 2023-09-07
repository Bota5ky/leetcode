package gosolutions

// f(n)=f(n-1)+f(n-2)
// f(1)=1;f(2)=2;3;5;8;13;21
// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
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
